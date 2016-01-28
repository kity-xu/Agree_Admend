package iccard

import (
	"cn.agree/dllinterop"
	"cn.agree/driverlayer"
	"cn.agree/utils"
	"errors"
	"fmt"
	"syscall"
	"unsafe"
)

//ic卡com口通用实现
type STICCard struct {
	termdll  syscall.Handle
	read     uintptr
	poweroff uintptr
}

func (jst *STICCard) Initdriver(pin *driverlayer.DriverArg) {
	var err error
	jst.termdll, err = syscall.LoadLibrary("third_party/shengteng/iccard_shengteng.dll")
	if err != nil {
		panic("load library error:" + err.Error())
	}
	jst.read, err = syscall.GetProcAddress(jst.termdll, "drive")

}

func (jst *STICCard) Deinit(pin *driverlayer.DriverArg) error {
	syscall.FreeLibrary(jst.termdll)
	return nil
}

func (jst *STICCard) GetFactoryName() string {
	return "通用驱动"
}

//调用winscard.dll实现
func (jst *STICCard) PowerOff(pin *driverlayer.DriverArg, timeout int) error {
	utils.Debug("receive STICCard PowerOff request")

	var term_type int = 0
	var func_id int = 4
	var buf_size int = 1024
	var func_arg int = 1

	t := make([]byte, 1024)
	tty_name := make([]byte, 1)
	tty_name[0] = 0

	var func_arg_pointer []*byte
	func_arg_pointer = make([]*byte, 1)

	porttype, port, _ := driverlayer.GetPortDescription(pin.Port)
	switch porttype {
	case "USB":
		func_arg_pointer[0] = &([]byte("U"))[0]
	case "COM":
		func_arg_pointer[0] = &([]byte("C"))[0]
	}
	utils.Debug("port is %s,%d", pin.Port, port)

	r, _, _ := dllinterop.CallProc(jst.read, uintptr(term_type), uintptr(unsafe.Pointer(&tty_name[0])),
		uintptr(port), uintptr(timeout), uintptr(func_id),
		uintptr(func_arg), uintptr(unsafe.Pointer(&func_arg_pointer[0])), uintptr(unsafe.Pointer(&t[0])), uintptr(buf_size))

	if int(r) == 0 {
		return nil
	} else {
		empocc := utils.IndexByteRune(t, 0, 1)
		if empocc != -1 {
			t = t[0:empocc]
		}
		out, _ := utils.TransUTF8FromCode(t, utils.GBK)
		utils.Error("end STICCard PowerOff request error %s", string(out))
		return errors.New(string(out))
	}

}

//获取ic卡的信息
func (jst *STICCard) GetICCardInfo(pin *driverlayer.DriverArg, timeout int, taglist []byte, lpicappdata []byte) ([]byte, error) {
	utils.Debug("receive STICCard GetICCardInfo request")

	var term_type int = 0
	var func_id int = 0
	var buf_size int = 4096
	var func_arg int = 2

	t := make([]byte, 4096)
	tty_name := make([]byte, 1)
	tty_name[0] = 0

	var func_arg_pointer []*byte
	func_arg_pointer = make([]*byte, 2)

	if taglist == nil || len(taglist) == 0 {
		taglist = make([]byte, 1)
		taglist[0] = 0
	}
	lpicappdata = make([]byte, 4096)

	porttype, port, _ := driverlayer.GetPortDescription(pin.Port)
	switch porttype {
	case "USB":
		func_arg_pointer[0] = &([]byte("U"))[0]
	case "COM":
		func_arg_pointer[0] = &([]byte("C"))[0]
	}
	//taglist
	//func_arg_pointer[3] = &lpicappdata[0]
	func_arg_pointer[1] = &taglist[0]

	r, _, _ := dllinterop.CallProc(jst.read, uintptr(term_type), uintptr(unsafe.Pointer(&tty_name[0])),
		uintptr(port), uintptr(timeout), uintptr(func_id),
		uintptr(func_arg), uintptr(unsafe.Pointer(&func_arg_pointer[0])), uintptr(unsafe.Pointer(&t[0])), uintptr(buf_size))

	if int(r) == 0 {
		var empocc int
		empocc = utils.IndexByteRune(t, 0, 1)
		if empocc != -1 {
			t = t[0:empocc]
		}

		utils.Debug("end STICCard GetICCardInfo request ,data [%s]", string(t))
		return t, nil
	} else {
		utils.Error("end STICCard GetICCardInfo request error %s", int(r))
		return nil, errors.New(fmt.Sprintf("%d", int(r)))
	}

}

//获取arqc
func (jst *STICCard) GenARQC(pin *driverlayer.DriverArg, timeout int, taglist []byte, lpicappdata []byte) ([]byte, error) {
	utils.Debug("receive STICCard GenARQC request")

	var term_type int = 0
	var func_id int = 1
	var buf_size int = 4096
	var func_arg int = 3

	t := make([]byte, 4096)
	tty_name := make([]byte, 1)
	tty_name[0] = 0

	var func_arg_pointer []*byte
	func_arg_pointer = make([]*byte, 3)

	if taglist == nil || len(taglist) == 0 {
		taglist = tty_name
	}
	if lpicappdata == nil || len(lpicappdata) == 0 {
		lpicappdata = tty_name
	}

	porttype, port, _ := driverlayer.GetPortDescription(pin.Port)
	switch porttype {
	case "USB":
		func_arg_pointer[0] = &([]byte("U"))[0]
	case "COM":
		func_arg_pointer[0] = &([]byte("C"))[0]
	}
	//taglist
	func_arg_pointer[2] = &lpicappdata[0]
	func_arg_pointer[1] = &taglist[0]

	r, _, _ := dllinterop.CallProc(jst.read, uintptr(term_type), uintptr(unsafe.Pointer(&tty_name[0])),
		uintptr(port), uintptr(timeout), uintptr(func_id),
		uintptr(func_arg), uintptr(unsafe.Pointer(&func_arg_pointer[0])), uintptr(unsafe.Pointer(&t[0])), uintptr(buf_size))

	if int(r) == 0 {
		var empocc int
		empocc = utils.IndexByteRune(t, 0, 1)
		if empocc != -1 {
			t = t[0:empocc]
		}

		empocc = utils.IndexByteRune(t, '|', 2)
		if empocc != -1 {
			t = t[empocc+1 : len(t)]
		}

		utils.Debug("end STICCard GenARQC request ,data [%s]", string(t))
		return t, nil
	} else {
		utils.Error("end STICCard GenARQC request error %s", int(r))
		return nil, errors.New(fmt.Sprintf("%d", int(r)))
	}

}

//获取arqc
func (jst *STICCard) GetTransDetail(pin *driverlayer.DriverArg, timeout int, path []byte) ([]byte, error) {
	utils.Debug("receive STICCard GetTransDetail request")

	var term_type int = 0
	var func_id int = 3
	var buf_size int = 4096
	var func_arg int = 1

	t := make([]byte, 4096)
	tty_name := make([]byte, 1)
	tty_name[0] = 0

	var func_arg_pointer []*byte
	func_arg_pointer = make([]*byte, 1)

	if path == nil || len(path) == 0 {
		path = tty_name
	}

	_, port, _ := driverlayer.GetPortDescription(pin.Port)

	//taglist
	func_arg_pointer[0] = &path[0]

	r, _, _ := dllinterop.CallProc(jst.read, uintptr(term_type), uintptr(unsafe.Pointer(&tty_name[0])),
		uintptr(port), uintptr(timeout), uintptr(func_id),
		uintptr(func_arg), uintptr(unsafe.Pointer(&func_arg_pointer[0])), uintptr(unsafe.Pointer(&t[0])), uintptr(buf_size))

	if int(r) == 0 {
		var empocc int
		empocc = utils.IndexByteRune(t, 0, 1)
		if empocc != -1 {
			t = t[0:empocc]
		}

		empocc = utils.IndexByteRune(t, '|', 2)
		if empocc != -1 {
			t = t[empocc+1 : len(t)]
		}

		utils.Debug("end STICCard GetTransDetail request ,data [%s]", string(t))
		return t, nil
	} else {
		utils.Error("end STICCard GetTransDetail request error %s", int(r))
		return nil, errors.New(fmt.Sprintf("%d", int(r)))
	}

}

func (jst *STICCard) CtrScriptData(pin *driverlayer.DriverArg, timeout int, taglist []byte, lpicappdata []byte, arpc []byte) ([]byte, error) {
	utils.Debug("receive STICCard CtrScriptData request")

	var term_type int = 0
	var func_id int = 2
	var buf_size int = 4096
	var func_arg int = 4

	t := make([]byte, 4096)
	tty_name := make([]byte, 1)
	tty_name[0] = 0

	var func_arg_pointer []*byte
	func_arg_pointer = make([]*byte, 4)

	if taglist == nil || len(taglist) == 0 {
		taglist = tty_name
	}
	if lpicappdata == nil || len(lpicappdata) == 0 {
		lpicappdata = tty_name
	}
	if arpc == nil || len(arpc) == 0 {
		arpc = tty_name
	}

	porttype, port, _ := driverlayer.GetPortDescription(pin.Port)
	switch porttype {
	case "USB":
		func_arg_pointer[0] = &([]byte("U"))[0]
	case "COM":
		func_arg_pointer[0] = &([]byte("C"))[0]
	}
	//taglist
	func_arg_pointer[2] = &lpicappdata[0]
	func_arg_pointer[1] = &taglist[0]
	func_arg_pointer[3] = &arpc[0]

	r, _, _ := dllinterop.CallProc(jst.read, uintptr(term_type), uintptr(unsafe.Pointer(&tty_name[0])),
		uintptr(port), uintptr(timeout), uintptr(func_id),
		uintptr(func_arg), uintptr(unsafe.Pointer(&func_arg_pointer[0])), uintptr(unsafe.Pointer(&t[0])), uintptr(buf_size))

	if int(r) == 0 {
		var empocc int
		empocc = utils.IndexByteRune(t, 0, 1)
		if empocc != -1 {
			t = t[0:empocc]
		}

		empocc = utils.IndexByteRune(t, '|', 2)
		if empocc != -1 {
			t = t[empocc+1 : len(t)]
		}

		utils.Debug("end STICCard GenARQC request ,data [%s]", string(t))
		return t, nil
	} else {
		utils.Error("end STICCard GenARQC request error %s", int(r))
		return nil, errors.New(fmt.Sprintf("%d", int(r)))
	}
}
