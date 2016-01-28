package protocol

import (
	"cn.agree/driverlayer"
	"cn.agree/utils"
	"net/http"
)

type MsfProtocol struct {
}

//获得MSF的实现
func getMSFImpl() (driverlayer.IMsf, *driverlayer.DriverArg, error) {
	drvbase, drvarg, err := getDeviceImp(MSF)
	if err != nil {
		return nil, nil, err
	}

	b, err := MarkDeviceUse(MSF)
	if err != nil || b == false {
		return nil, nil, err
	}

	return drvbase.(driverlayer.IMsf), drvarg, nil
}

func (p *MsfProtocol) Read(r *http.Request, arg *JsonGeneralReadInputAgr, res *JsonResponse) error {
	utils.Debug("received MsfProtocol read request,arg is : %+v", *arg)
	f, a, err := getMSFImpl()
	if err != nil {
		utils.Error("get msf impl error: %s", err.Error())
		res.Code = driverlayer.DEVICE_IN_USE
		res.ErrMsg = err.Error()
		utils.Debug("end MsfProtocol Read request")

		return nil
	}
	defer MarkDeviceIDLE(MSF)
	s, err := f.Read(a, arg.Readtype, arg.Timeout)
	if err != nil {
		utils.Error("invoke msf read error:%s", err.Error())
		res.Code = driverlayer.DEVICE_OPER_ERROR
		res.ResMsg = string(s)
	} else {
		res.Code = driverlayer.DEVICE_OPER_SUCCESS
		res.ResMsg = string(s)
		utils.Debug("MsfProtocol read success %s", string(s))
	}
	utils.Debug("end MsfProtocol read request")
	return nil
}
