package idccard

import (
	"bytes"
	"cn.agree/dllinterop"
	"cn.agree/driverlayer"
	"cn.agree/utils"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"syscall"
	"unsafe"
)

//磁条卡com口通用实现
type STIdr struct {
	termdll syscall.Handle
	read    uintptr
}

func (jst *STIdr) Initdriver(pin *driverlayer.DriverArg) {
	var err error
	jst.termdll, err = syscall.LoadLibrary("third_party/shengteng/idcheck_shengteng.dll")
	if err != nil {
		panic("load library error:" + err.Error())
	}

	jst.read, err = syscall.GetProcAddress(jst.termdll, "drive")

}

func (jst *STIdr) Deinit(pin *driverlayer.DriverArg) error {
	syscall.FreeLibrary(jst.termdll)
	return nil
}

func (jst *STIdr) GetFactoryName() string {
	return "升腾"
}

//调用terms.dll
func (jst *STIdr) ReadData(pin *driverlayer.DriverArg, timeout int) ([]byte, error) {
	utils.Debug("receive STIdr ReadData request")

	var term_type int = 0
	var func_id int = 2
	var buf_size int = 1024
	var empty int = 0

	t := make([]byte, 1024)
	tty_name := make([]byte, 1)
	tty_name[0] = 0
	var func_arg_pointer []*byte
	func_arg_pointer = make([]*byte, 1)
	func_arg_pointer[0] = &([]byte("data/"))[0]

	_, port, _ := driverlayer.GetPortDescription(pin.Port)

	r, _, _ := dllinterop.CallProc(jst.read, uintptr(term_type), uintptr(unsafe.Pointer(&tty_name[0])),
		uintptr(port), uintptr(timeout), uintptr(func_id),
		uintptr(empty), uintptr(unsafe.Pointer(&func_arg_pointer[0])), uintptr(unsafe.Pointer(&t[0])), uintptr(buf_size))

	if int(r) == 0 {
		out, err := utils.TransUTF8FromCode(t, utils.GBK)
		if err != nil {
			return nil, err
		}
		utils.Debug("end STIdr ReadData request %s", string(out))
		empocc := utils.IndexByteRune(out, 0, 1)
		if empocc != -1 {
			out = out[0:empocc]
		}

		res := bytes.NewBufferString("")
		res.Write(out)
		firocc := utils.IndexByteRune(out, ',', 5)
		secocc := utils.IndexByteRune(out, ',', 6)
		if firocc == -1 || secocc == -1 {
			utils.Error("STIdr ReadData data error  %s", string(out))
			return nil, errors.New("not formatted data")
		}
		res.WriteByte('|')

		rb, ferr := ioutil.ReadFile(fmt.Sprintf(`data/%s%s`, string(out[firocc+1:secocc]), ".bmp"))

		if ferr != nil {
			utils.Error("can't read file  %s", ferr.Error())
			return nil, errors.New("can't read file")
		}

		var b bytes.Buffer

		we := base64.NewEncoder(base64.StdEncoding, &b)
		we.Write(rb)
		we.Close()

		res.WriteString(string(b.Bytes()))
		utils.Debug("process STIdr ReadData data %s", res.String())
		return res.Bytes(), nil
	} else {
		utils.Error("end STIdr ReadData request error %s", int(r))
		return nil, errors.New(fmt.Sprintf("%d", int(r)))
	}

}
