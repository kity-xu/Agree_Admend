package main

import (
	"cn.agree/crossframe"
	"cn.agree/protocol"
	"cn.agree/utils"
	"flag"
)

var logl *string = flag.String("log", "development", "process log level")

func main() {
	flag.Parse()

	//开启调试服务
	utils.Info("%s", "begin start local debug service....")
	go protocol.StartLocalDebugServer()
	utils.Info("%s", "end start local debug service success....")

	//初始化日志系统,因为要连接调试端口,所以放在debug服务后面
	switch *logl {
	case "development":
		utils.InitializeLogging(1)
		break
	case "production":
		utils.InitializeLogging(2)
		break
	default:
		utils.InitializeLogging(2)
		break
	}

	utils.Info("%s", "begin start remote debug service....")
	go protocol.StartRemoteDebugServer()
	utils.Info("%s", "end start remote debug service success....")

	defer utils.Flush()

	utils.Info("%s", "begin to parse all drivers...")
	//读入配置文件
	err := protocol.ParserDrivers()
	if err != nil {
		utils.Critical("parse xml file,error is : %s", err.Error())
		return
	}

	utils.Debug("find all devices category:%+v", protocol.AllDevice)
	utils.Info("%s", "end parse all drivers suceess....")

	utils.Info("%s", "begin parse use drivers....")

	err = protocol.ParseUseDriver()
	if err != nil {
		utils.Critical("parser use file error,error is:%s", err.Error())
		return
	}
	utils.Debug("find all user driver: %+v", protocol.CurDrivers)
	utils.Info("%s", "end parse use drivers success....")

	//add by yangxiaolong
	utils.Info("%s", "begin parse connect method...")
	/*err = protocol.ParserConnMethod()
	  if err != nil {
	     utils.Critical("parser connect method file error,error is:%s", err.Error())
	     return
	  }
	  utils.Debug("find method: %s, find ip address: %s", protocol.Connect.Method, protocol.Connect.Address)
	  utils.Info("%s", "end parse connect method success...")*/

	if protocol.Connect.Method == "WIFI" {
       wserver := protocol.Wificonn{}
       wserver.Addr = protocol.Connect.Address
       err = wserver.StartWifiServer()
       if err != nil {
       	  
       }    	
    }

	//开启跨域代理服务
	utils.Info("begin start proxy service...")
	go crossframe.CrossFrameProxyServer()

	//启动硬件服务
	utils.Info("%s", "begin start driver service....")
	err = protocol.StartDriverService(protocol.CurDrivers.Device)
	if err != nil {
		utils.Critical("failed to start device ,error is %s", err.Error())
		return
	}

	utils.Info("%s", "end start driver service success....")

	//开启json服务
	jserver := protocol.JsonDriverServer{}
	jserver.Port = 8888
	utils.Info("begin start json service....")
	err = jserver.StartServer()
	if err != nil {
		utils.Error("detect error,error is %s", err.Error())
		return
	}

	utils.Info("%s", "end start json service success....")

	//写入所有的log日志
	utils.Flush()
}
