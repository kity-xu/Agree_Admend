// 深圳九思泰达密码键盘的实现
//
package pin

import (
	"cn.agree/driverlayer"
	"cn.agree/utils"
	"encoding/hex"
)

//密码键盘通用实现
type PinGeneral struct {
}

func (jst *PinGeneral) Initdriver(pin *driverlayer.DriverArg) {

}

func (jst *PinGeneral) Deinit(pin *driverlayer.DriverArg) error {
	return nil
}

func (jst *PinGeneral) GetFactoryName() string {
	return "通用驱动"
}

func (jst *PinGeneral) Readonce(pin *driverlayer.DriverArg, timeout int) ([]byte, error) {
	utils.Debug("PinGeneral received Readonce request")
	s, _ := hex.DecodeString("1b49")
	t, _ := hex.DecodeString("83")
	utils.Debug("PinGeneral send readonce pi %x", s)
	b, e := driverlayer.WritePortAndReadWithTerminator(pin.Port, pin.Baud, s, []byte{0x03}, timeout, t)
	if e == nil {
		utils.Debug("begin process received data %x", b)
		b = b[1 : len(b)-1]
		utils.Debug("after process received data %x", b)
	}
	return b, e
}

func (jst *PinGeneral) Readtwice(pin *driverlayer.DriverArg, timeout int) ([]byte, error) {
	utils.Debug("PinGeneral received Readtwice request")
	s, _ := hex.DecodeString("1b45")
	t, _ := hex.DecodeString("83")
	utils.Debug("PinGeneral send readtwice pi %x", s)

	b, e := driverlayer.WritePortAndReadWithTerminator(pin.Port, pin.Baud, s, []byte{0x03}, timeout, t)
	if e == nil {
		utils.Debug("begin process received data %x", b)
		b = b[1 : len(b)-1]
		utils.Debug("after process received data %x", b)
	}
	return b, e
}

func (jst *PinGeneral) Reset(pin *driverlayer.DriverArg, timeout int) error {
	utils.Debug("PinGeneral received Reset request")
	return nil
}
