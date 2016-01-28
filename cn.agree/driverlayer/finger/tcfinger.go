package finger

import (
	"cn.agree/dllinterop"
	"cn.agree/driverlayer"
	"cn.agree/utils"
	"errors"
	"syscall"
	"unsafe"
)

//升腾指纹仪
type TCFinger struct {
	stdll            syscall.Handle
	getRegisterImage uintptr
	getValidteImage  uintptr
	matchtz          uintptr
	reset            uintptr
}

var GET_FINGER_ERROR = errors.New("获取指纹失败")

func (jst *TCFinger) Initdriver(pin *driverlayer.DriverArg) {
	var err error
	jst.stdll, err = syscall.LoadLibrary("third_party/teso/TesoLive.dll")
	if err != nil {
		panic("LoadLibrary " + err.Error())
	}
	jst.getRegisterImage, err = syscall.GetProcAddress(jst.stdll, "FPIGetTemplate")
	jst.getValidteImage, err = syscall.GetProcAddress(jst.stdll, "FPIGetFeature")
	jst.matchtz, err = syscall.GetProcAddress(jst.stdll, "FPIMatch")
}

func (jst *TCFinger) Deinit(pin *driverlayer.DriverArg) error {
	syscall.FreeLibrary(jst.stdll)
	return nil
}

func (jst *TCFinger) GetFactoryName() string {
	return "天诚"
}

//获取指纹
func (jst *TCFinger) GetRegisterFg(pin *driverlayer.DriverArg, timeout int) ([]byte, error) {
	var use int = 0
	t := make([]byte, 513)
	err := make([]byte, 65)

	_, port, _ := driverlayer.GetPortDescription(pin.Port)

	r, _, _ := dllinterop.CallProc(jst.getRegisterImage,
		uintptr(port), uintptr(unsafe.Pointer(&t[0])), uintptr(unsafe.Pointer(&use)), uintptr(unsafe.Pointer(&err[0])))

	if int(r) >= 0 {
		return t[0:use], nil
	} else {
		out, _ := utils.TransUTF8FromCode(err, utils.GBK)
		return nil, errors.New(string(out))
	}
}

//注册指纹
func (jst *TCFinger) GetValidateFg(pin *driverlayer.DriverArg, timeout int) ([]byte, error) {
	var use int = 0
	t := make([]byte, 513)
	err := make([]byte, 65)

	_, port, _ := driverlayer.GetPortDescription(pin.Port)

	r, _, _ := dllinterop.CallProc(jst.getValidteImage,
		uintptr(port), uintptr(unsafe.Pointer(&t[0])), uintptr(use), uintptr(unsafe.Pointer(&err[0])))

	for i := 0; i < 512; i++ {
		if t[i] == 0 {
			use = i
			break
		}
	}
	if int(r) >= 0 {
		return t[0:use], nil
	} else {
		out, _ := utils.TransUTF8FromCode(err, utils.GBK)
		return nil, errors.New(string(out))
	}

	return t, nil
}

func (jst *TCFinger) MatchFinger(pin *driverlayer.DriverArg, timeout int, reg []byte, vad []byte) int {
	//产生内容
	var level int = 3
	r, _, _ := dllinterop.CallProc(jst.matchtz,
		uintptr(unsafe.Pointer(&reg[0])), uintptr(unsafe.Pointer(&vad[0])), uintptr(level))

	//解析生成的文件
	return int(r)
}

func (jst *TCFinger) Reset(pin *driverlayer.DriverArg, timeout int) error {
	dllinterop.CallProc(jst.reset)
	return nil
}
