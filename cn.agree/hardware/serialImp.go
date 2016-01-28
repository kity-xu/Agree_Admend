package hardware

import (
	"cn.agree/utils"
	"errors"
	"github.com/efarres/goparallel"
	"github.com/tarm/goserial"
	"io"
	"strconv"
	"strings"
)

var DEVICE_NOT_IMPLEMENTED = errors.New("device operation not implemented")

//传入端口代码,获得操作流
//一般用于程序自己操作端口
//支持com口,并口,转换口
//com口和并口操作类似
func GetPortInstance(port string, baud int) (io.ReadWriteCloser, error) {
	utils.Trace("received Get Port request,port is %s,baud is %d", port, baud)
	if strings.HasPrefix(port, "COM") {
		c := &serial.Config{Name: port, Baud: baud}
		utils.Trace("achieve com port : %s,baud : %d", port, baud)
		return serial.OpenPort(c)
	} else if strings.HasPrefix(port, "LPT") {
		c := &parallel.Config{Name: port}
		utils.Trace("achieve parallel port : %s,baud : %d", port, baud)
		return parallel.OpenPort(c)
	} else if strings.HasPrefix(port, "USB") {
		utils.Error("yet haven't implementd usb,port is %s ", port)
		return nil, DEVICE_NOT_IMPLEMENTED
	} else {
		utils.Error("non recognized port %s ", port)
		return nil, DEVICE_NOT_IMPLEMENTED
	}

}

func GetPortDescription(port string) (string, int, error) {
	if strings.HasPrefix(port, "COM") {
		sport := strings.TrimLeft(port, "COM")
		port, err := strconv.Atoi(sport)
		if err != nil {
			return "", -1, err
		} else {
			return "COM", port, nil
		}
	} else if strings.HasPrefix(port, "USB") {
		sport := strings.TrimLeft(port, "USB")
		if strings.EqualFold(sport, "") {
			return "USB", 0, nil
		}
		port, err := strconv.Atoi(sport)
		if err != nil {
			return "", -1, err
		} else {
			return "USB", port, nil
		}
	}
	return "", -1, DEVICE_NOT_IMPLEMENTED
}
