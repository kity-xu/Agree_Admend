package iccard

import (
	"cn.agree/dllinterop"
	"cn.agree/driverlayer"
	"cn.agree/utils"
	"errors"
	"syscall"
	"unsafe"
)

//ic卡com口通用实现
type ICGeneral struct {
	winsdll  syscall.Handle
	readcard uintptr
}

const (
	SCARD_S_SUCCESS = 0
)

func (jst *ICGeneral) Initdriver(pin *driverlayer.DriverArg) {
	var err error
	jst.winsdll, err = syscall.LoadLibrary("third_party/general/iccardreader.dll")
	if err != nil {
		panic("load library error:" + err.Error())
	}

	f := func(n string, res *uintptr) bool {
		*res, err = syscall.GetProcAddress(jst.winsdll, n)
		if err != nil {
			syscall.FreeLibrary(jst.winsdll)
			panic("load proc " + n + " error:" + err.Error())
		}
		return true
	}

	//获得dll的各种handle
	if f("ReadICCard", &jst.readcard) {
	} else {
		utils.Error("Init general Driver error...%s", err.Error())
	}

}

func (jst *ICGeneral) Deinit(pin *driverlayer.DriverArg) error {
	syscall.FreeLibrary(jst.winsdll)
	return nil
}

func (jst *ICGeneral) GetFactoryName() string {
	return "通用驱动"
}

//调用winscard.dll实现
func (jst *ICGeneral) ReadCardNo(pin *driverlayer.DriverArg, timeout int) ([]byte, error) {
	utils.Debug("receive ICGeneral ReadCardNo request")
	cardno := make([]byte, 100)
	errmsg := make([]byte, 300)
	var len = 100
	r, _, _ := dllinterop.CallProc(jst.readcard, uintptr(timeout), uintptr(unsafe.Pointer(&cardno[0])),
		uintptr(unsafe.Pointer(&len)), uintptr(unsafe.Pointer(&errmsg[0])))
	if int(r) == 0 {
		utils.Debug("reteive card no %s", string(cardno[0:len]))
		cardno = cardno[0:len]
		for i, v := range cardno {
			if v == 0x3D {
				cardno = cardno[0:i]
				break
			}

		}

		utils.Debug("end ICGeneral ReadCardNo request")
		return cardno, nil
	} else {
		utils.Debug("end ICGeneral ReadCardNo request")
		return nil, errors.New(string(errmsg))
	}

}
