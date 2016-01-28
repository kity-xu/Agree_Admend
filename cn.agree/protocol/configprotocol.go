package protocol

import (
	"bytes"
	"cn.agree/hardware"
	"cn.agree/utils"
	"fmt"
	"github.com/mathume/dom"
	"io/ioutil"
	"net/http"
	"strings"
)

//配置服务
type ConfigProtocol struct {
}

//获取设备配置
func (p *ConfigProtocol) GetDevConfig(r *http.Request, arg *JsonDevInputAgr, res *JsonResponse) error {
	utils.Debug("received ConfigProtocol GetDevConfig request")
	fs, err := GetDrvConfig(arg.Dev)
	if err != nil {
		res.Code = 1
		res.ErrMsg = err.Error()
		utils.Error("get %s config file content error,detail is %s", arg.Dev, err.Error())
	} else {
		res.Code = 0
		res.ErrMsg = string(fs)
		utils.Debug("get config file content success :\r\n%s", string(fs))
	}
	utils.Debug("end ConfigProtocol GetDevConfig request")
	return nil
}

//获取正在使用的设备配置
func (p *ConfigProtocol) GetAllDevConfig(r *http.Request, arg *JsonGeneralInputAgr, res *JsonResponse) error {
	utils.Debug("received ConfigProtocol GetAllDevConfig request")
	fs, err := GetAllDrvConfig()
	if err != nil {
		res.Code = 1
		res.ErrMsg = err.Error()
		utils.Error("get all config file content error,detail is %s", err.Error())
	} else {
		res.Code = 0
		res.ResMsg = string(fs)
		utils.Debug("get config file content success :\r\n%s", string(fs))
	}
	utils.Debug("end ConfigProtocol GetAllDevConfig request")
	return nil
}

//获取所有可用设备
func (p *ConfigProtocol) GetDeviceList(r *http.Request, arg *JsonGeneralInputAgr, res *JsonResponse) error {
	utils.Debug("received ConfigProtocol GetDeviceList request")
	fs, err := GetAllDeviceList()
	if err != nil {
		res.Code = 1
		res.ErrMsg = err.Error()
		utils.Error("get all config file content error,detail is %s", err.Error())
	} else {
		res.Code = 0
		res.ResMsg = string(fs)
		utils.Debug("get config file content success :\r\n%s", string(fs))
	}
	utils.Debug("end ConfigProtocol GetDeviceList request")
	return nil
}

//保存单一设备配置
func (p *ConfigProtocol) SaveDevConfig(r *http.Request, arg *JsonConfigInputAgr, res *JsonResponse) error {
	utils.Debug("received ConfigProtocol SaveDevConfig request")
	err := SaveDrvConfig(arg.Dev, arg.Content)
	if err != nil {
		res.Code = 1
		res.ErrMsg = err.Error()
		utils.Error("save %s config file content error,detail is %s", arg.Dev, err.Error())
	} else {
		res.Code = 0
		utils.Debug("save config file content success")
	}
	utils.Debug("end ConfigProtocol SaveDevConfig request")
	return nil
}

//保存所有设备配置
func (p *ConfigProtocol) SaveAllDevConfig(r *http.Request, arg *JsonConfigInputAgr, res *JsonResponse) error {
	utils.Debug("received ConfigProtocol SaveAllDevConfig request")
	var (
		n  dom.Node
		t  dom.Element
		dt dom.Attribute
	)
	db := dom.NewDOMBuilder(strings.NewReader(arg.Content), dom.NewDOMStore())
	d, err2 := db.Build()
	if err2 != nil {
		res.Code = 1
		res.ErrMsg = fmt.Sprintf("parse xml file error,category :[%+v] ", err2.Error())
		utils.Error("parse xml file error,category:[%+v]", err2.Error())
		return nil
	}

	f := d.ChildNodes()

	for i := 0; i < len(f); i++ {
		n = f[i]
		switch n.Kind() {
		case dom.ElementKind:
			t = n.(dom.Element)
			switch t.Name() {
			case "Device": // 解析设置
				var v []dom.Attribute
				v = t.Attr()
				for j := 0; j < len(v); j++ {
					dt = v[j]
					if strings.EqualFold(dt.Name(), "Category") {
						err := SaveDrvConfig(dt.Value(), t.String())

						if err != nil {
							res.Code = 1
							res.ErrMsg = fmt.Sprintf("save file error,category :%s  content is : [%s]", dt.Value(), t.String())
							utils.Error("can't save category [%s] ,content is : [%s]", dt.Value(), t.String())
							return nil
						}
					}
				}
				break

			default:
				utils.Error("can't recognize element name: %s", t.Name())
			}
			break
		}
	}

	res.Code = 0

	utils.Debug("end ConfigProtocol SaveAllDevConfig request")
	return nil
}

//保存设备配置
func (p *ConfigProtocol) GetUpdateUrl(r *http.Request, arg *JsonGeneralInputAgr, res *JsonResponse) error {
	utils.Debug("received ConfigProtocol GetUpdateUrl request")

	fs, err := ioutil.ReadFile("config/updateurl.xml")
	if err != nil {
		res.Code = 1
		res.ErrMsg = err.Error()
		utils.Error("get update url config error,detail is %s", err.Error())
	} else {
		res.Code = 0
		res.ErrMsg = string(fs)
		utils.Debug("get updaterul file content success :\r\n%s", string(fs))
	}
	utils.Debug("end ConfigProtocol GetUpdateUrl request")
	return nil
}

//保存设备配置
func (p *ConfigProtocol) SaveUpdateUrl(r *http.Request, arg *JsonUpdateInputAgr, res *JsonResponse) error {
	utils.Debug("received ConfigProtocol SaveUpdateUrl request")
	err := ioutil.WriteFile("config/updateurl.xml", []byte(arg.Content), 0644)
	if err != nil {
		res.Code = 1
		res.ErrMsg = err.Error()
		utils.Error("save updateurl file content error,detail is %s", err.Error())
	} else {
		res.Code = 0
		utils.Debug("save updateurl file content success")
	}
	utils.Debug("end ConfigProtocol SaveUpdateUrl request")
	return nil
}

//获得所有硬件列表
func (p *ConfigProtocol) GetAllHardWare(r *http.Request, arg *JsonUpdateInputAgr, res *JsonResponse) error {
	utils.Debug("received ConfigProtocol GetAllHardWare request")
	//列举com口
	s := []string{"COM1", "COM2", "COM3", "COM4", "COM5", "COM6", "COM7", "COM8", "COM9", "COM10"}
	bs := bytes.NewBufferString("")
	for _, t := range s {

		w, err := hardware.GetPortInstance(t, 9600)

		if err == nil {
			bs.WriteString(t)
			bs.WriteString(",")
			w.Close()
		}
	}
	bs.WriteString("USB1")

	res.Code = 0
	res.ResMsg = bs.String()

	utils.Debug("end ConfigProtocol GetAllHardWare request,return data is [%s]", bs.String())
	return nil
}
