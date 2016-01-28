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
type ZZFinger struct {
	stdll syscall.Handle
	read  uintptr
}

func (jst *ZZFinger) Initdriver(pin *driverlayer.DriverArg) {
	var err error
	jst.stdll, err = syscall.LoadLibrary("third_party/zhongzheng/FingerDLL.dll")
	if err != nil {
		panic("LoadLibrary " + err.Error())
	}
	jst.read, err = syscall.GetProcAddress(jst.stdll, "drive")
}

func (jst *ZZFinger) Deinit(pin *driverlayer.DriverArg) error {
	syscall.FreeLibrary(jst.stdll)
	return nil
}

func (jst *ZZFinger) GetFactoryName() string {
	return "中正"
}

//获取指纹
func (jst *ZZFinger) GetRegisterFg(pin *driverlayer.DriverArg, timeout int) ([]byte, error) {
	utils.Debug("receive ZZFinger GetRegisterFg request")

	var term_type int = 0
	var func_id int = 2
	var buf_size int = 1024
	var func_arg int = 1

	t := make([]byte, 1024)
	tty_name := make([]byte, 1)
	tty_name[0] = 0

	var func_arg_pointer []*byte
	func_arg_pointer = make([]*byte, 1)

	_, port, _ := driverlayer.GetPortDescription(pin.Port)

	func_arg_pointer[0] = &([]byte("X"))[0]

	r, _, _ := dllinterop.CallProc(jst.read, uintptr(term_type), uintptr(unsafe.Pointer(&tty_name[0])),
		uintptr(port), uintptr(timeout), uintptr(func_id),
		uintptr(func_arg), uintptr(unsafe.Pointer(&func_arg_pointer[0])), uintptr(unsafe.Pointer(&t[0])), uintptr(buf_size))

	if int(r) == 0 {
		return t, nil
	} else {
		empocc := utils.IndexByteRune(t, 0, 1)
		if empocc != -1 {
			t = t[0:empocc]
		}
		out, _ := utils.TransUTF8FromCode(t, utils.GBK)
		utils.Error("end ZZFinger GetRegisterFg request error %s", string(out))
		return nil, errors.New(string(out))
	}
}

//注册指纹
func (jst *ZZFinger) GetValidateFg(pin *driverlayer.DriverArg, timeout int) ([]byte, error) {
	utils.Debug("receive ZZFinger GetRegisterFg request")

	var term_type int = 0
	var func_id int = 1
	var buf_size int = 1024
	var func_arg int = 1

	t := make([]byte, 1024)
	tty_name := make([]byte, 1)
	tty_name[0] = 0

	var func_arg_pointer []*byte
	func_arg_pointer = make([]*byte, 1)

	_, port, _ := driverlayer.GetPortDescription(pin.Port)

	func_arg_pointer[0] = &([]byte("X"))[0]

	r, _, _ := dllinterop.CallProc(jst.read, uintptr(term_type), uintptr(unsafe.Pointer(&tty_name[0])),
		uintptr(port), uintptr(timeout), uintptr(func_id),
		uintptr(func_arg), uintptr(unsafe.Pointer(&func_arg_pointer[0])), uintptr(unsafe.Pointer(&t[0])), uintptr(buf_size))

	if int(r) == 0 {
		return t, nil
	} else {
		empocc := utils.IndexByteRune(t, 0, 1)
		if empocc != -1 {
			t = t[0:empocc]
		}
		out, _ := utils.TransUTF8FromCode(t, utils.GBK)
		utils.Error("end ZZFinger GetRegisterFg request error %s", string(out))
		return nil, errors.New(string(out))
	}
}

func (jst *ZZFinger) MatchFinger(pin *driverlayer.DriverArg, timeout int, reg []byte, vad []byte) int {

	//解析生成的文件
	return int(-1)
}

func (jst *ZZFinger) Reset(pin *driverlayer.DriverArg, timeout int) error {

	return nil
}
