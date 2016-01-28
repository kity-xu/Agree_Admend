package protocol

import (
	"cn.agree/driverlayer"
	"cn.agree/utils"
	"net/http"
)

//调用winscard.进行操作
type FinProtocol struct {
}

//获得指纹的实现
func getFingerImpl() (driverlayer.IFinger, *driverlayer.DriverArg, error) {
	drvbase, drvarg, err := getDeviceImp(FP)
	if err != nil {
		return nil, nil, err
	}

	b, err := MarkDeviceUse(FP)
	if err != nil || b == false {
		return nil, nil, err
	}

	return drvbase.(driverlayer.IFinger), drvarg, nil
}

func (p *FinProtocol) GetRegisterFg(r *http.Request, arg *JsonGeneralInputAgr, res *JsonResponse) error {
	utils.Debug("received FinProtocol GetFingerPrinter request,arg is : %+v", *arg)

	f, a, err := getFingerImpl()
	if err != nil {
		utils.Error("get fin impl error: %s", err.Error())
		res.Code = driverlayer.DEVICE_IN_USE
		res.ErrMsg = err.Error()
		utils.Debug("end fin GetFingerPrinter request")
		return nil
	}

	defer MarkDeviceIDLE(FP)
	s, err := f.GetRegisterFg(a, arg.Timeout)
	if err != nil {
		utils.Error("invoke GetFingerPrinter read error: %s", err.Error())
		res.Code = driverlayer.DEVICE_OPER_ERROR
		res.ErrMsg = err.Error()
	} else {
		res.Code = driverlayer.DEVICE_OPER_SUCCESS
		res.ResMsg = string(s)
		utils.Debug("pin Readonce request success,reteive data %s", string(s))

	}
	utils.Debug("end FinProtocol GetFingerPrinter request")
	return nil
}

func (p *FinProtocol) GetValidateFg(r *http.Request, arg *JsonGeneralInputAgr, res *JsonResponse) error {
	utils.Debug("received FinProtocol GetValidateFg request,arg is : %+v", *arg)

	f, a, err := getFingerImpl()
	if err != nil {
		utils.Error("get fin impl error: %s", err.Error())
		res.Code = driverlayer.DEVICE_IN_USE
		res.ErrMsg = err.Error()
		utils.Debug("end fin GetValidateFg request")
		return nil
	}

	defer MarkDeviceIDLE(FP)
	s, err := f.GetValidateFg(a, arg.Timeout)
	if err != nil {
		utils.Error("invoke GetValidateFg read error: %s", err.Error())
		res.Code = driverlayer.DEVICE_OPER_ERROR
		res.ErrMsg = err.Error()
	} else {
		res.Code = driverlayer.DEVICE_OPER_SUCCESS
		res.ResMsg = string(s)
		utils.Debug("fin GetValidateFg request success,reteive data %s", string(s))

	}
	utils.Debug("end FinProtocol GetValidateFg request")
	return nil
}

func (p *FinProtocol) MatchFinger(r *http.Request, arg *JsonFingerMatchInputAgr, res *JsonResponse) error {
	utils.Debug("received FinProtocol MatchFinger request,arg is : %+v", *arg)
	f, a, err := getFingerImpl()
	if err != nil {
		utils.Error("get fin impl error: %s", err.Error())
		res.Code = driverlayer.DEVICE_IN_USE
		res.ErrMsg = err.Error()
		utils.Debug("end fin GetValidateFg request")
		return nil
	}

	defer MarkDeviceIDLE(FP)
	s := f.MatchFinger(a, arg.Timeout, []byte(arg.Reg), []byte(arg.Vad))
	if s >= 0 {
		utils.Debug("match finger success")
	} else {
		utils.Debug("match finger error")
	}
	res.Code = driverlayer.DEVICE_OPER_SUCCESS
	res.ErrMsg = ""
	if s >= 0 {
		res.ResMsg = "true"
	} else {
		res.ResMsg = "false"
	}
	utils.Debug("end FinProtocol MatchFinger request")

	return nil
}
