package protocol

import (
	"cn.agree/utils"
	"io"
	"net"
	"strings"
	"sync"
)

const BUFF_SIZE = 1024

var buff = make([]byte, BUFF_SIZE)
var rcvbuffer = make([]byte, BUFF_SIZE)

type connlist struct {
	rw      sync.RWMutex
	conlist []*net.TCPConn
}

var con connlist

func addConnection(tcpConn *net.TCPConn) {
	con.rw.Lock()
	defer con.rw.Unlock()
	con.conlist = append(con.conlist, tcpConn)
}

func removeConnection(tcpConn *net.TCPConn) {
	con.rw.Lock()
	defer con.rw.Unlock()
	var i = 0
	for ; i < len(con.conlist); i++ {
		if strings.EqualFold(con.conlist[i].RemoteAddr().String(), tcpConn.RemoteAddr().String()) {
			utils.Debug("clear remote connection", tcpConn.RemoteAddr().String())
			con.conlist[i] = nil
			copy(con.conlist[:i], con.conlist[i+1:])
			con.conlist = con.conlist[0 : len(con.conlist)-1]
			return
		}
	}
}

//本地调试信息
func handleLocalConn(tcpConn *net.TCPConn) {
	if tcpConn == nil {
		return
	}
	var v *net.TCPConn
	var ll int
	var err error
	utils.Info("received local debug client:%s", tcpConn.LocalAddr().String())
	for {
		ll, err = tcpConn.Read(rcvbuffer)
		if err == io.EOF {
			tcpConn.Close()
			return
		} else {
			con.rw.RLock()
			for _, v = range con.conlist {
				v.Write(rcvbuffer[0:ll])
			}
			con.rw.RUnlock()
		}
	}
	tcpConn.Close()
}

//远程调试信息
func handleRemoteConn(tcpConn *net.TCPConn) {
	if tcpConn == nil {
		return
	}
	addConnection(tcpConn)
	utils.Info("received remote debug client:%s", tcpConn.LocalAddr().String())
	tcpConn.Write([]byte("welcome to debug service...."))
	for {
		_, err := tcpConn.Read(buff)
		if err == io.EOF {
			utils.Info("remote disconnect:%s", tcpConn.RemoteAddr().String())
			removeConnection(tcpConn)
			tcpConn.Close()
			return
		}
	}
}

//开启调试服务
func StartLocalDebugServer() {
	port := "8887"

	tcpAddr, err1 := net.ResolveTCPAddr("tcp4", "0.0.0.0:"+port)
	if err1 != nil {
		utils.Error("start resolving socket error:%s", err1.Error())
		return
	}

	tcpListener, err2 := net.ListenTCP("tcp", tcpAddr) //监听
	if err2 != nil {
		utils.Error("start listening debug socket error:%s", err2.Error())
		return
	} else {
		utils.Info("debug server start listening ip 0.0.0.0 port %s", port)
	}

	defer tcpListener.Close()
	for {
		tcpConn, err3 := tcpListener.AcceptTCP()
		if err3 != nil {
			utils.Error("establish tcp connection error:%s", err3.Error())
			continue
		}
		go handleLocalConn(tcpConn) //起一个goroutine处理
	}

}

//开启调试服务
func StartRemoteDebugServer() {
	port := "6666"

	tcpAddr, err1 := net.ResolveTCPAddr("tcp4", "0.0.0.0:"+port)
	if err1 != nil {
		utils.Error("start resolving socket error:%s", err1.Error())
		return
	}

	tcpListener, err2 := net.ListenTCP("tcp", tcpAddr) //监听
	if err2 != nil {
		utils.Error("start listening debug socket error:%s", err2.Error())
		return
	} else {
		utils.Info("debug server start listening ip 0.0.0.0 port %s", port)
	}

	defer tcpListener.Close()
	for {
		tcpConn, err3 := tcpListener.AcceptTCP()
		if err3 != nil {
			utils.Error("establish tcp connection error:%s", err3.Error())
			continue
		}
		go handleRemoteConn(tcpConn) //起一个goroutine处理
	}

}
