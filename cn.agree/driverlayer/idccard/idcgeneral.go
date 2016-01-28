package idccard

import (
	"bytes"
	"cn.agree/dllinterop"
	"cn.agree/driverlayer"
	"cn.agree/utils"
	"code.google.com/p/go.text/encoding/unicode"
	"code.google.com/p/go.text/transform"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"syscall"
	"time"
	"unsafe"
)

//磁条卡com口通用实现
type IDCGeneral struct {
	termdll     syscall.Handle
	initcomm    uintptr
	closecomm   uintptr
	findcard    uintptr
	selectcard  uintptr
	readbasemsg uintptr
}

var wzfile = []byte("./wz.txt")
var bmfile = []byte("./bm.wlp")

func (jst *IDCGeneral) Initdriver(pin *driverlayer.DriverArg) {
	var err error
	jst.termdll, err = syscall.LoadLibrary("third_party/general/sdtapi.dll")
	if err != nil {
		panic("load library error:" + err.Error())
	}

	f := func(n string, res *uintptr) bool {
		*res, err = syscall.GetProcAddress(jst.termdll, n)
		if err != nil {
			syscall.FreeLibrary(jst.termdll)
			panic("load proc " + n + " error:" + err.Error())
		}
		return true
	}

	//获得dll的各种handle
	if f("SDT_OpenPort", &jst.initcomm) &&
		f("SDT_ClosePort", &jst.closecomm) &&
		f("SDT_StartFindIDCard", &jst.findcard) &&
		f("SDT_SelectIDCard", &jst.selectcard) &&
		f("SDT_ReadBaseMsg", &jst.readbasemsg) {
	} else {
		utils.Error("Init shensi Driver error...%s", err.Error())
	}

}

func (jst *IDCGeneral) Deinit(pin *driverlayer.DriverArg) error {
	syscall.FreeLibrary(jst.termdll)
	return nil
}

func (jst *IDCGeneral) GetFactoryName() string {
	return "通用驱动"
}

func (jst *IDCGeneral) genIDC(port string, timeout int, readtype int) ([]byte, error) {
	os.Remove(string(wzfile))
	os.Remove(string(bmfile))

	//获得端口描述
	s, p, err := driverlayer.GetPortDescription(port)
	if err != nil {
		utils.Error("receive shensi port error %s  ", port)
		return nil, err
	}
	utils.Debug("begin to open port %s %d", s, p)
	var r uintptr
	if strings.EqualFold(s, "COM") {
		r, _, _ = dllinterop.CallProc(uintptr(jst.initcomm), uintptr(p))
	} else if strings.EqualFold(s, "USB") {
		p = p + 1000
		r, _, _ = dllinterop.CallProc(uintptr(jst.initcomm), uintptr(p))
	}
	if int32(r) != 144 {
		utils.Error("shensi init com error %d", int32(r))
		return nil, err
	}

	//关闭端口
	defer dllinterop.CallProc(uintptr(jst.closecomm), uintptr(p))
	var punc int32
	var punsn int64
	step := 1
	timed, _ := time.ParseDuration(strconv.Itoa(timeout) + "ms")

	remiantime := timed * time.Millisecond

	var (
		basemsg []byte
		bmpmsg  []byte
	)

	//检查二代证的放置位置
	for {
		switch step {
		case 1:
			{
				r, _, _ = dllinterop.CallProc(uintptr(jst.findcard), uintptr(p), uintptr(unsafe.Pointer(&punc)), uintptr(0))
				utils.Debug("start find card result %d", int32(r))
				if r == 159 {
					utils.Debug("find card success")
					step = 2
				}
			}
			break
		case 2:
			{
				r, _, _ = dllinterop.CallProc(uintptr(jst.selectcard), uintptr(p), uintptr(unsafe.Pointer(&punsn)), uintptr(0))
				if r == 144 {
					utils.Debug("select card success")
					step = 3
				} else {
					utils.Debug("start select card result %d", int32(r))
				}
			}
		case 3:
			{
				var (
					wzlen int
					bmlen int
				)
				basemsg = make([]byte, 1024, 1024)
				bmpmsg = make([]byte, 1024, 1024)
				r, _, _ = dllinterop.CallProc(uintptr(jst.readbasemsg), uintptr(p),
					uintptr(unsafe.Pointer(&basemsg[0])),
					uintptr(unsafe.Pointer(&wzlen)),
					uintptr(unsafe.Pointer(&bmpmsg[0])),
					uintptr(unsafe.Pointer(&bmlen)),
					uintptr(0))
				if r == 144 {
					utils.Debug("end idc card read")
					basemsg = basemsg[0:wzlen]
					bmpmsg = bmpmsg[0:bmlen]
					goto readsuccess
				} else {
					utils.Debug("start select card result %d", int32(r))
				}
			}

		}

		remiantime = remiantime - 1000*time.Millisecond
		if remiantime <= 0 {
			break
		}
		time.Sleep(1000 * time.Millisecond)
		if remiantime <= 0 {
			return nil, driverlayer.ErrIOTimeOut
		}
	}

readsuccess:

	outinfo := make([]byte, 0, 1000)

	rInUTF8 := transform.NewReader(bytes.NewReader(basemsg[0:30]), unicode.UTF16(true, false).NewDecoder())
	out, _ := ioutil.ReadAll(rInUTF8)

	outinfo = append(outinfo, out[0:len(out)]...)
	outinfo = append(outinfo, '|')

	rInUTF8 = transform.NewReader(bytes.NewReader(basemsg[30:32]), unicode.UTF16(true, false).NewDecoder())
	out, _ = ioutil.ReadAll(rInUTF8)
	outinfo = append(outinfo, out[0:len(out)]...)
	outinfo = append(outinfo, '|')

	rInUTF8 = transform.NewReader(bytes.NewReader(basemsg[32:36]), unicode.UTF16(true, false).NewDecoder())
	out, _ = ioutil.ReadAll(rInUTF8)
	outinfo = append(outinfo, out[0:len(out)]...)
	outinfo = append(outinfo, '|')

	rInUTF8 = transform.NewReader(bytes.NewReader(basemsg[36:52]), unicode.UTF16(true, false).NewDecoder())
	out, _ = ioutil.ReadAll(rInUTF8)
	outinfo = append(outinfo, out[0:len(out)]...)
	outinfo = append(outinfo, '|')

	rInUTF8 = transform.NewReader(bytes.NewReader(basemsg[52:122]), unicode.UTF16(true, false).NewDecoder())
	out, _ = ioutil.ReadAll(rInUTF8)
	outinfo = append(outinfo, out[0:len(out)]...)
	outinfo = append(outinfo, '|')

	rInUTF8 = transform.NewReader(bytes.NewReader(basemsg[122:158]), unicode.UTF16(true, false).NewDecoder())
	out, _ = ioutil.ReadAll(rInUTF8)
	outinfo = append(outinfo, out[0:len(out)]...)
	outinfo = append(outinfo, '|')

	rInUTF8 = transform.NewReader(bytes.NewReader(basemsg[158:188]), unicode.UTF16(true, false).NewDecoder())
	out, _ = ioutil.ReadAll(rInUTF8)
	outinfo = append(outinfo, out[0:len(out)]...)
	outinfo = append(outinfo, '|')

	rInUTF8 = transform.NewReader(bytes.NewReader(basemsg[188:204]), unicode.UTF16(true, false).NewDecoder())
	out, _ = ioutil.ReadAll(rInUTF8)
	outinfo = append(outinfo, out[0:len(out)]...)
	outinfo = append(outinfo, '|')

	rInUTF8 = transform.NewReader(bytes.NewReader(basemsg[204:220]), unicode.UTF16(true, false).NewDecoder())
	out, _ = ioutil.ReadAll(rInUTF8)
	outinfo = append(outinfo, out[0:len(out)]...)

	//只读基础ixnxi
	if readtype == 1 {
		return outinfo, nil
	} else {
		outinfo = append(outinfo, '|')
		return outinfo, nil
	}
}

func (jst *IDCGeneral) ReadData(pin *driverlayer.DriverArg, timeout int) ([]byte, error) {
	utils.Debug("receive general ReadData request")
	out, err := jst.genIDC(pin.Port, timeout, 2)
	if err != nil {
		utils.Debug("end general ReadData request")
		return nil, err
	} else {
		utils.Debug("reteive id card info %s", string(out))
		utils.Debug("end general ReadData request")
		return out, nil
	}
}
