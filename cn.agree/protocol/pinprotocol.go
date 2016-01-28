package protocol

import (
	"cn.agree/driverlayer"
	"cn.agree/utils"
	"net/http"
)

type PinProtocol struct {
}

//获得PIN的实现
func getPinImpl() (driverlayer.IPin, *driverlayer.DriverArg, error) {
	drvbase, drvarg, err := getDeviceImp(PIN)

	if err != nil {
		return nil, nil, err
	}

	b, err := MarkDeviceUse(PIN)
	if err != nil || b == false {
		return nil, nil, err
	}

	return drvbase.(driverlayer.IPin), drvarg, nil
}

func (p *PinProtocol) Readonce(r *http.Request, arg *JsonGeneralInputAgr, res *JsonResponse) error {
	utils.Debug("received pin readonce request,arg is : %+v", *arg)

	f, a, err := getPinImpl()
	if err != nil {
		utils.Error("get pin impl error: %s", err.Error())
		res.Code = driverlayer.DEVICE_IN_USE
		res.ErrMsg = err.Error()
		utils.Debug("end pin readonce request")
		return nil
	}

	defer MarkDeviceIDLE(PIN)
	s, err := f.Readonce(a, arg.Timeout)

	if err != nil {
		utils.Error("invoke pin Readonce error: %s", err.Error())
		res.Code = driverlayer.DEVICE_OPER_ERROR
		res.ErrMsg = err.Error()
	} else {
		res.Code = driverlayer.DEVICE_OPER_SUCCESS
		res.ResMsg = string(s)
		utils.Debug("pin Readonce request success,reteive data %s", string(s))
	}
	utils.Debug("end pin readonce request")

	return nil
}

func (p *PinProtocol) Readtwice(r *http.Request, arg *JsonGeneralInputAgr, res *JsonResponse) error {
	utils.Debug("received pin readtwice request,arg is : %+v", *arg)
	f, a, err := getPinImpl()
	if err != nil {
		utils.Error("get pin impl error: %s", err.Error())
		res.Code = driverlayer.DEVICE_IN_USE
		res.ErrMsg = err.Error()
		utils.Debug("end pin readonce request")
		return nil
	}
	defer MarkDeviceIDLE(PIN)
	s, err := f.Readtwice(a, arg.Timeout)
	if err != nil {
		utils.Error("invoke pin Readtwice error:%s", err.Error())
		res.Code = driverlayer.DEVICE_OPER_ERROR
		res.ErrMsg = err.Error()
	} else {
		res.Code = driverlayer.DEVICE_OPER_SUCCESS
		res.ResMsg = string(s)
		utils.Debug("pin readtwice request success,teteive data %s", string(s))
	}
	utils.Debug("end pin readtwice request")
	return nil
}
