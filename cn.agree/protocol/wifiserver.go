package protocol

import (
	"cn.agree/driverlayer"
	"cn.agree/utils"
	"net"
)

//this file add by yangxiaolong
type Wificonn struct {
	Net  string
	Addr string
}

func (server *Wificonn) StartWifiServer() error {
	conn, err := net.Dial("tcp", server.Addr)
     if err != nil {
     	utils.Error("start wifi service:%s", err.Error())
     	return err
     }
     driverlayer.SaveWifiConn(conn)
	 return nil
}
