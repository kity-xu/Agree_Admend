package protocol

import (
	"cn.agree/utils"
	"github.com/gorilla/rpc/v2"
	"github.com/gorilla/rpc/v2/json2"
	"net/http"
	"strconv"
)

const JSON_RESPONSE_OK = 0

type JsonDriverServer struct {
	Port int
}

//json的返回值
type JsonResponse struct {
	Code   int
	Key    int
	ErrMsg string
	ResMsg string
}

//json的通用输入参数
type JsonGeneralInputAgr struct {
	Timeout int
}

//json的读取类型输入参数
type JsonGeneralReadInputAgr struct {
	Timeout  int
	Readtype int
}

//json的日志级别输入参数
type JsonLogInputAgr struct {
	Timeout int
	Kind    int
}

//json的设备输入参数
type JsonDevInputAgr struct {
	Timeout int
	Dev     string
}

//json的设备输入参数
type JsonConfigInputAgr struct {
	Timeout int
	Dev     string
	Content string
}

//json的更新服务器地址
type JsonUpdateInputAgr struct {
	Timeout int
	Content string
}

//json的通用输入参数
type JsonPr2InputAgr struct {
	Timeout int
	Con     string
}

//json的通用输入参数
type JsonFingerMatchInputAgr struct {
	Timeout int
	Reg     string
	Vad     string
}

//json的IC卡输入参数
type JsonICInputAgr struct {
	Timeout     int
	TagList     string
	LpicAppData string
}

//json的获得交易数据
type JsonGetDetailInputAgr struct {
	Timeout int
	Path    string
}

//json的脚本执行参数
type JsonScriptDetailInputAgr struct {
	Timeout     int
	TagList     string
	LpicAppData string
	ARPC        string
}

func (server *JsonDriverServer) StartServer() error {
	s := rpc.NewServer()
	s.RegisterCodec(json2.NewCodec(), "application/json-rpc")
	s.RegisterService(new(PinProtocol), "")
	s.RegisterService(new(FinProtocol), "")
	s.RegisterService(new(PingjiaProtocol), "")
	s.RegisterService(new(MsfProtocol), "")
	s.RegisterService(new(IcProtocol), "")
	s.RegisterService(new(IdcProtocol), "")
	s.RegisterService(new(DebugProtocol), "")
	s.RegisterService(new(ConfigProtocol), "")
	s.RegisterService(new(Pr2Protocol), "")
	http.Handle("/rpc", s)
	p := strconv.Itoa(server.Port)
	var err error
	//具有多个ip

	err = http.ListenAndServe("0.0.0.0:"+p, nil)
	utils.Info("json rpc server start listening ip 0.0.0.0 port %s", p)
	if err != nil {
		utils.Error("start listening ip 0.0.0.0 error:%s", err.Error())
	}

	return nil
}
