package protocol

import (
	"cn.agree/driverlayer"
	"cn.agree/print"
	"cn.agree/utils"
	"net/http"
)

type Pr2Protocol struct {
}

//获得PR2的实现
func getPr2Impl() (driverlayer.IPr2Print, *driverlayer.DriverArg, error) {
	drvbase, drvarg, err := getDeviceImp(PR2PRT)
	if err != nil {
		return nil, nil, err
	}
	b, err := MarkDeviceUse(PR2PRT)
	if err != nil || b == false {
		return nil, nil, err
	}
	return drvbase.(driverlayer.IPr2Print), drvarg, nil
}

func (p *Pr2Protocol) Print(r *http.Request, arg *JsonPr2InputAgr, res *JsonResponse) error {
	utils.Debug("received pr2 print request,arg is : %+v", *arg)

	f, s, err := getPr2Impl()
	if err != nil {
		utils.Error("get pr2 impl error: %s", err.Error())
		res.Code = driverlayer.DEVICE_IN_USE
		res.ErrMsg = err.Error()
		utils.Debug("end pr2 print request")
		return nil
	}
	defer MarkDeviceIDLE(PR2PRT)

	//打印
	print.RunPr2Engine(arg.Con, f, s)

	utils.Debug("end pr2 print request")
	return nil
}

//add by yangxiaolong
func (p *Pr2Protocol) PrintEx(r *http.Request, arg *JsonPr2InputAgr, res *JsonResponse) error {
	utils.Debug("received pr2 print request,arg is : %+v", *arg)

	f, s, err := getPr2Impl()
	if err != nil {
		utils.Error("get pr2 impl error: %s", err.Error())
		res.Code = driverlayer.DEVICE_IN_USE
		res.ErrMsg = err.Error()
		utils.Debug("end pr2 print request")
		return nil
	}
	defer MarkDeviceIDLE(PR2PRT)

	f.BeginPrintJob(s, 20)
	defer f.EndPrinterJob()

	//debug by jinsl
	//f.Init()
	f.SetTop(10)
	f.SetLeftMargin(2)
	f.OutputString(arg.Con)
	f.EjectPaper()
	res.ResMsg = "成功"
	res.ErrMsg = "0"

	utils.Debug("end pr2 print request")
	return nil
}
