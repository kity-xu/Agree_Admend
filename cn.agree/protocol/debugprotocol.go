package protocol

import (
	"bytes"
	"cn.agree/utils"
	"net/http"
	"text/template"
)

//所有设备
var ad = template.Must(template.New("alldevice").Parse(`<?xml version='1.0' encoding='utf-8'?>
		<root>
  			<res>
				<oper>success</oper>
  			</res>
  		<content>
  		{{range .}}<Device Category='{{.Category}}' State='{{.State}}' Drivername='{{.FactoryName}}' />
  		{{end}}
  		</content>
  		</root>
	`))

//所有运行设备
var rad = template.Must(template.New("alldevice").Parse(`<?xml version='1.0' encoding='utf-8'?>
		<root>
  			<res>
				<oper>success</oper>
  			</res>
  		<content>
  		{{range .}}{{if ne .State 0}}<Device Category='{{.Category}}' status='{{.State}}' drivername='{{.FactoryName}}' />{{end}}
  		{{end}}
  		</content>
  		</root>
	`))

//debug服务
type DebugProtocol struct {
}

//获取所有设备
func (p *DebugProtocol) GetAllDevice(r *http.Request, arg *JsonGeneralInputAgr, res *JsonResponse) error {
	utils.Debug("%s", "received DebugProtocol GetAllDevice request")
	var buf bytes.Buffer
	ad.Execute(&buf, getAllDevice())
	utils.Debug("get all device: %s", buf.String())
	res.ResMsg = buf.String()
	utils.Debug("%s", "end DebugProtocol GetAllDevice request")
	return nil
}

//获取所有运行设备
func (p *DebugProtocol) GetUsableDevice(r *http.Request, arg *JsonGeneralInputAgr, res *JsonResponse) error {
	utils.Debug("%s", "received DebugProtocol GetUsableDevice request")
	var buf bytes.Buffer
	rad.Execute(&buf, getAllDevice())
	utils.Debug("get all runnable device: %s", buf.String())
	res.ResMsg = buf.String()
	utils.Debug("%s", "end DebugProtocol GetUsableDevice request")
	return nil
}

//改变日志级别
func (p *DebugProtocol) ChangeLogEnv(r *http.Request, arg *JsonLogInputAgr, res *JsonResponse) error {
	utils.Debug("%s", "received DebugProtocol ChangeLogLevel request")
	utils.ChangeLogEnv(arg.Kind)
	utils.Debug("%s", "end DebugProtocol ChangeLogLevel request")
	return nil
}
