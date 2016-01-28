package protocol

import (
	"cn.agree/driverlayer"
	"cn.agree/utils"
	"net/http"
)

//调用winscard.进行操作
type IcProtocol struct {
}

//获得IC卡的实现
func getICImpl() (driverlayer.ICReader, *driverlayer.DriverArg, error) {
	drvbase, drvarg, err := getDeviceImp(IC)
	if err != nil {
		return nil, nil, err
	}

	b, err := MarkDeviceUse(IC)
	if err != nil || b == false {
		return nil, nil, err
	}

	return drvbase.(driverlayer.ICReader), drvarg, nil
}

//下电
func (p *IcProtocol) PowerOff(r *http.Request, arg *JsonGeneralInputAgr, res *JsonResponse) error {
	utils.Debug("received iccard PowerOff request,arg is :%+v", *arg)

	f, a, err := getICImpl()

	if err != nil {
		utils.Error("get ic impl error: %s", err.Error())
		res.Code = driverlayer.DEVICE_IN_USE
		res.ErrMsg = err.Error()
		utils.Debug("end ic PowerOff request")
		return nil
	}

	defer MarkDeviceIDLE(IC)
	err = f.PowerOff(a, arg.Timeout)
	if err != nil {
		utils.Error("invoke iccard PowerOff error:%s", err.Error())
		res.Code = driverlayer.DEVICE_OPER_ERROR
		res.ErrMsg = ""
		res.ResMsg = string(err.Error())
	} else {
		res.Code = driverlayer.DEVICE_OPER_SUCCESS
		res.ErrMsg = ""
		utils.Debug("received PowerOff read")
	}
	utils.Debug("end iccard PowerOff request")

	return nil
}

//下电
func (p *IcProtocol) GetICCardInfo(r *http.Request, arg *JsonICInputAgr, res *JsonResponse) error {
	utils.Debug("received iccard GetICCardInfo request,arg is :%+v", *arg)
	var out []byte
	f, a, err := getICImpl()

	if err != nil {
		utils.Error("get ic impl error: %s", err.Error())
		res.Code = driverlayer.DEVICE_IN_USE
		res.ErrMsg = err.Error()
		utils.Debug("end ic GetICCardInfo request")
		return nil
	}

	defer MarkDeviceIDLE(IC)
	out, err = f.GetICCardInfo(a, arg.Timeout, []byte(arg.TagList), []byte(arg.LpicAppData))
	if err != nil {
		utils.Error("invoke iccard GetICCardInfo error:%s", err.Error())
		res.Code = driverlayer.DEVICE_OPER_ERROR
		res.ErrMsg = string(err.Error())
	} else {
		res.Code = driverlayer.DEVICE_OPER_SUCCESS
		res.ResMsg = string(out)
		utils.Debug("received GetICCardInfo read")
	}
	utils.Debug("end iccard GetICCardInfo request")

	return nil
}

//从ic卡获取arqc
func (p *IcProtocol) GenARQC(r *http.Request, arg *JsonICInputAgr, res *JsonResponse) error {
	utils.Debug("received iccard GenARQC request,arg is :%+v", *arg)
	var out []byte
	f, a, err := getICImpl()

	if err != nil {
		utils.Error("get ic impl error: %s", err.Error())
		res.Code = driverlayer.DEVICE_IN_USE
		res.ErrMsg = err.Error()
		utils.Debug("end ic GenARQC request")
		return nil
	}

	defer MarkDeviceIDLE(IC)
	out, err = f.GenARQC(a, arg.Timeout, []byte(arg.TagList), []byte(arg.LpicAppData))
	if err != nil {
		utils.Error("invoke iccard GenARQC error:%s", err.Error())
		res.Code = driverlayer.DEVICE_OPER_ERROR
		res.ErrMsg = string(err.Error())
	} else {
		res.Code = driverlayer.DEVICE_OPER_SUCCESS
		res.ResMsg = string(out)
		utils.Debug("received GenARQC read")
	}
	utils.Debug("end iccard GenARQC request")

	return nil
}

//获得交易详细
func (p *IcProtocol) GetTransDetail(r *http.Request, arg *JsonGetDetailInputAgr, res *JsonResponse) error {
	utils.Debug("received iccard GenARQC request,arg is :%+v", *arg)
	var out []byte
	f, a, err := getICImpl()

	if err != nil {
		utils.Error("get ic impl error: %s", err.Error())
		res.Code = driverlayer.DEVICE_IN_USE
		res.ErrMsg = err.Error()
		utils.Debug("end ic GetTransDetail request")
		return nil
	}

	defer MarkDeviceIDLE(IC)
	out, err = f.GetTransDetail(a, arg.Timeout, []byte(arg.Path))
	if err != nil {
		utils.Error("invoke iccard GetTransDetail error:%s", err.Error())
		res.Code = driverlayer.DEVICE_OPER_ERROR
		res.ErrMsg = string(err.Error())
	} else {
		res.Code = driverlayer.DEVICE_OPER_SUCCESS
		res.ResMsg = string(out)
		utils.Debug("received GetTransDetail read")
	}
	utils.Debug("end iccard GetTransDetail request")

	return nil
}

//获得交易详细
func (p *IcProtocol) CtrScriptData(r *http.Request, arg *JsonScriptDetailInputAgr, res *JsonResponse) error {
	utils.Debug("received iccard CtrScriptData request,arg is :%+v", *arg)
	var out []byte
	f, a, err := getICImpl()

	if err != nil {
		utils.Error("get ic impl error: %s", err.Error())
		res.Code = driverlayer.DEVICE_IN_USE
		res.ErrMsg = err.Error()
		utils.Debug("end ic CtrScriptData request")
		return nil
	}

	defer MarkDeviceIDLE(IC)
	out, err = f.CtrScriptData(a, arg.Timeout, []byte(arg.TagList), []byte(arg.LpicAppData), []byte(arg.ARPC))
	if err != nil {
		utils.Error("invoke iccard CtrScriptData error:%s", err.Error())
		res.Code = driverlayer.DEVICE_OPER_ERROR
		res.ErrMsg = string(err.Error())
	} else {
		res.Code = driverlayer.DEVICE_OPER_SUCCESS
		res.ResMsg = string(out)
		utils.Debug("received CtrScriptData read")
	}
	utils.Debug("end iccard CtrScriptData request")

	return nil
}
