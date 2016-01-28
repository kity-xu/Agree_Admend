package protocol

import (
	"cn.agree/driverlayer"
	"cn.agree/utils"
	"net/http"
)

//调用winscard.进行操作
type IdcProtocol struct {
}

//获得二代证的实现
func getIDCImpl() (driverlayer.IDCReader, *driverlayer.DriverArg, error) {
	drvbase, drvarg, err := getDeviceImp(IDR)
	if err != nil {
		return nil, nil, err
	}

	b, err := MarkDeviceUse(IDR)
	if err != nil || b == false {
		return nil, nil, err
	}
	return drvbase.(driverlayer.IDCReader), drvarg, nil
}

func (p *IdcProtocol) ReadData(r *http.Request, arg *JsonGeneralInputAgr, res *JsonResponse) error {
	utils.Debug("received Idccard ReadData request,arg is : %+v", *arg)

	f, a, err := getIDCImpl()
	if err != nil {
		utils.Error("get Idc impl error: %s", err.Error())
		res.Code = driverlayer.DEVICE_IN_USE
		res.ErrMsg = err.Error()
		utils.Debug("end Idc ReadData request")
		return nil
	}

	defer MarkDeviceIDLE(IDR)
	s, err := f.ReadData(a, arg.Timeout)
	if err != nil {
		utils.Error("invoke Idccard read error:%s", err.Error())
		res.Code = driverlayer.DEVICE_OPER_ERROR
		res.ErrMsg = ""
		res.ResMsg = string(s)
	} else {
		res.Code = driverlayer.DEVICE_OPER_SUCCESS
		res.ErrMsg = ""
		res.ResMsg = string(s)
		utils.Debug("IdcProtocol ReadData success")
	}
	utils.Debug("end Idccard ReadData request")

	return nil
}
