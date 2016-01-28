package driverlayer

import (
	"bytes"
	"cn.agree/hardware"
	"cn.agree/utils"
	"errors"
	"io"
	"net"
	"strconv"
	"time"
)

var handle net.Conn //add by yangxiaolong

var methodIsWifi bool = false

var ErrIOTimeOut = errors.New("operation exceed desired time")

func GetPortDescription(port string) (string, int, error) {
	return hardware.GetPortDescription(port)
}

//往接口中写入数据,读取数据。返回条件
// timeout:单位为纳秒
//1 :读取到数据
//2 :流结束
//3 :读取错误
//4 :超时
//5 :超时后执行的指令
func WritePortAndRead(port string, baud int, writedata []byte, readlen int, timeout int, ains []byte) ([]byte, error) {
	utils.Trace("received write port request,port is %s", port)
	w, err := hardware.GetPortInstance(port, baud)

	if err != nil {
		return nil, err
	}

	defer w.Close()

	b, s := WriteAndRead(w, writedata, readlen, timeout)
	if s == ErrIOTimeOut && ains != nil && len(ains) > 0 {
		utils.Info("detected device oper exceed maxinum time,write instructon : %x", ains)
		WriteData(w, ains)
	}
	return b, s
}

//往接口中写入数据,读取数据。返回条件
//1 读到指定readlen长度
//2 流结束
//3 读取错误
//4 超时
//5 超市后执行的指令
func WritePortAndReadWithLen(port string, baud int, writedata []byte, readlen int, timeout int, ains []byte) ([]byte, error) {
	utils.Trace("received WritePortAndReadWithLen,port is %s", port)
	w, err := hardware.GetPortInstance(port, baud)
	if err != nil {
		return nil, err
	}
	defer w.Close()

	b, s := WriteAndReadLen(w, writedata, readlen, timeout)
	if s == ErrIOTimeOut && ains != nil && len(ains) > 0 {
		utils.Info("detected device oper exceed maxinum time,write instructon : %x", ains)
		WriteData(w, ains)
	}
	return b, s
}

//往接口中写入数据,读取数据。返回条件
//1 读到ter结束符后返回.如果ter后有多余字符,将被截断
//2 流结束
//3 读取错误
//4 超时
//5 超时后执行的指令
func WritePortAndReadWithTerminator(port string, baud int, writedata []byte, ter []byte, timeout int, ains []byte) ([]byte, error) {
	utils.Trace("received WritePortAndReadWithTerminator,port is %s", port)
	w, err := hardware.GetPortInstance(port, baud)
	if err != nil {
		return nil, err
	}
	defer w.Close()
	b, s := WriteAndReadWithTer(w, writedata, ter, timeout)
	if s == ErrIOTimeOut && ains != nil && len(ains) > 0 {
		utils.Info("detected device oper exceed maxinum time,write instructon : %x", ains)
		WriteData(w, ains)
	}
	return b, s
}

//从io接口中读取任意长度数据
func WriteAndRead(rw io.ReadWriter, writedata []byte, readlen int, timeout int) ([]byte, error) {
	p := []byte(writedata)
	_, err := rw.Write(p)
	if err != nil {
		return nil, err
	}

	return ReadData(rw, readlen, timeout)
}

//向io接口写入数据
func WriteData(rw io.ReadWriter, writedata []byte) error {
	p := []byte(writedata)
	var err error
	if methodIsWifi {
       _, err = handle.Write(p)
	}else{
	   _, err = rw.Write(p)
    }

	if err != nil {
		return err
	}

	if err != nil {
		return err
	}

	return nil
}

//向io接口写入数据
func WritePortData(port string, baud int, writedata []byte) error {
	w, err := hardware.GetPortInstance(port, baud)
	if err != nil {
		return err
	}
	defer w.Close()
	err = WriteData(w, writedata)

	return err
}

//从io接口中读取指定长度数据
func WriteAndReadLen(rw io.ReadWriter, writedata []byte, readlen int, timeout int) ([]byte, error) {
	p := []byte(writedata)
	_, err := rw.Write(p)
	if err != nil {
		return nil, err
	}
	return ReadWithLen(rw, readlen, timeout)
}

//从io接口中读取指定结束符数据
func WriteAndReadWithTer(rw io.ReadWriter, writedata []byte, ter []byte, timeout int) ([]byte, error) {
	p := []byte(writedata)
	_, err := rw.Write(p)
	if err != nil {
		return nil, err
	}
	return ReadWithTer(rw, ter, timeout)
}

//从指定接口中读取数据,遇到分隔符后返回
//此算法的byte比较效率不高,适用于读取少量数据
func ReadWithTer(r io.Reader, ter []byte, timeout int) ([]byte, error) {
	utils.Trace("received read terminator task,%s", ter)
	rawdata := make([]byte, 0, 50)
	readdata := make([]byte, 10, 10)

	timed, _ := time.ParseDuration(strconv.Itoa(timeout) + "ms")

	t := time.NewTimer(timed)
	defer t.Stop()

	type res struct {
		buf []byte
		err error
	}

	chani := make(chan res, 1)

	//启动一个goroutine,执行读取命令

	go func() {
		for {
			n, err2 := r.Read(readdata)
			if err2 == io.EOF {
				if n > 0 {
					rawdata = append(rawdata, readdata[0:n]...)
					ind := bytes.Index(rawdata, ter)
					//读取正常数据
					if ind != -1 {
						chani <- res{buf: rawdata[:ind+len(ter)], err: io.EOF}
						return
					}

				}
				chani <- res{buf: rawdata[:len(rawdata)], err: io.EOF}
				return

			}
			if err2 != nil {
				chani <- res{buf: rawdata[:len(rawdata)], err: err2}
				return
			}

			rawdata = append(rawdata, readdata[0:n]...)

			ind := bytes.Index(rawdata, ter)
			//读取正常数据
			if ind != -1 {
				chani <- res{buf: rawdata[:ind+len(ter)], err: err2}
				return
			}
		}
	}()

	select {
	case <-t.C:
		return nil, ErrIOTimeOut
	case r := <-chani:
		return r.buf, r.err
	}

}

//从io接口中读取指定长度数据,长度不够,不会返回
func ReadWithLen(rw io.Reader, readlen int, timeout int) ([]byte, error) {
	utils.Trace("received read fixed len task,%d", readlen)
	if readlen == 0 {
		return nil, nil
	}
	readdata := make([]byte, readlen, readlen)

	var curlen = 0

	timed, _ := time.ParseDuration(strconv.Itoa(timeout) + "ms")

	t := time.NewTimer(timed)
	defer t.Stop()

	type res struct {
		buf []byte
		err error
	}

	chani := make(chan res, 1)

	//启动一个goroutine,执行读取命令

	go func() {
		for curlen < readlen {
			n, err2 := rw.Read(readdata[curlen:])
			if err2 == io.EOF {
				chani <- res{buf: readdata[:curlen+n], err: io.EOF}
				return
			}
			if err2 != nil {
				chani <- res{buf: readdata[:curlen], err: err2}
				return
			}

			curlen = curlen + n
			if curlen == readlen {
				chani <- res{buf: readdata, err: nil}
			}
		}
	}()

	select {
	case <-t.C:
		err := ErrIOTimeOut
		return nil, err
	case r := <-chani:
		return r.buf, r.err
	}

}

//读取指定长度内容,立即返回
func ReadData(rw io.Reader, readlen int, timeout int) ([]byte, error) {
	readdata := make([]byte, readlen, readlen)

	timed, _ := time.ParseDuration(strconv.Itoa(timeout) + "ms")

	t := time.NewTimer(timed)
	defer t.Stop()

	type res struct {
		buf []byte
		err error
	}

	chani := make(chan res, 1)

	//启动一个goroutine,执行读取命令

	go func() {
		n, err2 := rw.Read(readdata)
		if err2 == io.EOF {
			chani <- res{buf: readdata[0:n], err: io.EOF}
			return
		}
		if err2 != nil {
			chani <- res{buf: readdata[0:0], err: err2}
			return
		}
		chani <- res{buf: readdata[0:n], err: nil}
		return
	}()

	select {
	case <-t.C:
		return nil, ErrIOTimeOut
	case r := <-chani:
		return r.buf, r.err
	}

}

//add by yangxiaolong
func SaveWifiConn(conn net.Conn) {
	handle = conn
	methodIsWifi = true
	return
}
