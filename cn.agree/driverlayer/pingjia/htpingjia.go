package pingjia

import (
	"cn.agree/dllinterop"
	"cn.agree/driverlayer"
	"syscall"
	"unsafe"
)

//杭州中正
type HTPingjia struct {
	htdll         syscall.Handle
	startEstimate uintptr
}

func (jst *HTPingjia) Initdriver(pin *driverlayer.DriverArg) {
	var err error
	jst.htdll, err = syscall.LoadLibrary("estimator_ht.dll")
	if err != nil {
		panic("LoadLibrary " + err.Error())
	}
	jst.startEstimate, err = syscall.GetProcAddress(jst.htdll, "StartEstimate")
}

func (jst *HTPingjia) Deinit(pin *driverlayer.DriverArg) error {
	syscall.FreeLibrary(jst.htdll)
	return nil
}

func (jst *HTPingjia) GetFactoryName() string {
	return "HT"
}

func (jst *HTPingjia) StartEsitimate(pin *driverlayer.DriverArg, timeout int) (int, error) {
	var res []int
	dllinterop.CallProc(jst.startEstimate, uintptr(unsafe.Pointer(&pin.Port)),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("9600,n,8,1"))),
		uintptr(unsafe.Pointer(&timeout)), uintptr(unsafe.Pointer(&res)))
	return res[0], nil
}

func (jst *HTPingjia) CancelEsitimate(pin *driverlayer.DriverArg, timeout int) error {
	return nil
}

func (jst *HTPingjia) Reset(pin *driverlayer.DriverArg, timeout int) (int, error) {
	var res []int
	dllinterop.CallProc(jst.startEstimate, uintptr(unsafe.Pointer(&pin.Port)),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("9600,n,8,1"))),
		uintptr(unsafe.Pointer(&timeout)), uintptr(unsafe.Pointer(&res)))
	return res[0], nil
}
