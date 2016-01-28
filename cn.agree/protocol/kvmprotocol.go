package protocol

import (
	"cn.agree/driverlayer"
	"cn.agree/utils"
	"net/http"
)

type KvmProtocol struct {
}

//获得MSF的实现
func getKVMImpl() (driverlayer.IKvm, *driverlayer.DriverArg, error) {
	drvbase, drvarg, err := getDeviceImp(KVM)
	if err != nil {
		return nil, nil, err
	}

	b, err := MarkDeviceUse(KVM)
	if err != nil || b == false {
		return nil, nil, err
	}

	return drvbase.(driverlayer.IKvm), drvarg, nil
}

func (p *KvmProtocol) OutSyncInner(r *http.Request, arg *JsonGeneralInputAgr, res *JsonResponse) error {
	utils.Debug("received KvmProtocol OutSyncInner request,arg is : %+v", *arg)
	f, a, err := getKVMImpl()
	if err != nil {
		utils.Error("get kvm impl error: %s", err.Error())
		res.Code = driverlayer.DEVICE_IN_USE
		res.ErrMsg = err.Error()
		utils.Debug("end KvmProtocol OutSyncInner request")
		return nil
	}
	defer MarkDeviceIDLE(KVM)
	err = f.OutSyncInner(a, arg.Timeout)
	if err != nil {
		utils.Error("invoke kvm OutSyncInner error:%s", err.Error())
		res.Code = driverlayer.DEVICE_OPER_ERROR
	} else {
		res.Code = driverlayer.DEVICE_OPER_SUCCESS
		utils.Debug("KvmProtocol OutSyncInner success ")
	}
	utils.Debug("end KvmProtocol OutSyncInner request")
	return nil
}

func (p *KvmProtocol) DeSync(r *http.Request, arg *JsonGeneralInputAgr, res *JsonResponse) error {
	utils.Debug("received KvmProtocol DeSync request,arg is : %+v", *arg)
	f, a, err := getKVMImpl()
	if err != nil {
		utils.Error("get kvm impl error: %s", err.Error())
		res.Code = driverlayer.DEVICE_IN_USE
		res.ErrMsg = err.Error()
		utils.Debug("end KvmProtocol DeSync request")
		return nil
	}
	defer MarkDeviceIDLE(KVM)
	err = f.DeSync(a, arg.Timeout)
	if err != nil {
		utils.Error("invoke kvm DeSync error:%s", err.Error())
		res.Code = driverlayer.DEVICE_OPER_ERROR
	} else {
		res.Code = driverlayer.DEVICE_OPER_SUCCESS
		utils.Debug("KvmProtocol DeSync success ")
	}
	utils.Debug("end KvmProtocol DeSync request")
	return nil
}

func (p *KvmProtocol) InnSyncOuter(r *http.Request, arg *JsonGeneralInputAgr, res *JsonResponse) error {
	utils.Debug("received KvmProtocol InnSyncOuter request,arg is : %+v", *arg)
	f, a, err := getKVMImpl()
	if err != nil {
		utils.Error("get kvm impl error: %s", err.Error())
		res.Code = driverlayer.DEVICE_IN_USE
		res.ErrMsg = err.Error()
		utils.Debug("end KvmProtocol Read request")

		return nil
	}
	defer MarkDeviceIDLE(KVM)
	err = f.InnSyncOuter(a, arg.Timeout)
	if err != nil {
		utils.Error("invoke kvm InnSyncOuter error:%s", err.Error())
		res.Code = driverlayer.DEVICE_OPER_ERROR
	} else {
		res.Code = driverlayer.DEVICE_OPER_SUCCESS
		utils.Debug("KvmProtocol InnSyncOuter success ")
	}
	utils.Debug("end KvmProtocol InnSyncOuter request")
	return nil
}
