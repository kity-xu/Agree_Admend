package protocol

import (
	"cn.agree/driverlayer"
	"cn.agree/utils"
	"net/http"
)

type PingjiaProtocol struct {
}

//获得指纹的实现
func getPingjiaImpl() (driverlayer.IPingjia, *driverlayer.DriverArg, error) {
	drvbase, drvarg, err := getDeviceImp(PINGJIA)
	if err != nil {
		return nil, nil, err
	}

	b, err := MarkDeviceUse(PINGJIA)
	if err != nil || b == false {
		return nil, nil, err
	}

	return drvbase.(driverlayer.IPingjia), drvarg, nil
}

func (p *PingjiaProtocol) StartEstimate(r *http.Request, arg *JsonGeneralInputAgr, res *JsonResponse) error {
	utils.Debug("received PingjiaProtocol StartEstimate request,arg is : %+v", *arg)

	f, a, err := getPingjiaImpl()
	if err != nil {
		utils.Error("get Pingjia impl error: %s", err.Error())
		res.Code = driverlayer.DEVICE_IN_USE
		res.ErrMsg = err.Error()
		utils.Debug("end fin StartEstimate request")
		return nil
	}

	defer MarkDeviceIDLE(PINGJIA)
	s, err := f.StartEsitimate(a, arg.Timeout)
	if err != nil {
		utils.Error("invoke StartEstimate read error: %s", err.Error())
		res.Code = driverlayer.DEVICE_OPER_ERROR
		res.ErrMsg = err.Error()
	} else {
		res.Code = driverlayer.DEVICE_OPER_SUCCESS
		res.ResMsg = string(s)
		utils.Debug("Pingjia StartEstimate request success,reteive data %s", string(s))

	}
	utils.Debug("end PingjiaProtocol StartEstimate request")
	return nil
}
