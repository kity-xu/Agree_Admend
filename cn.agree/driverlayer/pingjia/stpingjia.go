package pingjia

import (
	"cn.agree/driverlayer"
	"cn.agree/utils"
	"encoding/hex"
	"errors"
)

var PINGJIARESULT_NOT_FORMULAED = errors.New("pingjia result not formulated")
var PINGJIARESULT_EXCEED_MAXTIME = errors.New("pingjia result exceed maxtime")

//升腾评价器
type STPingjia struct {
}

func (jst *STPingjia) Initdriver(pin *driverlayer.DriverArg) {
}

func (jst *STPingjia) Deinit(pin *driverlayer.DriverArg) error {
	return nil
}

func (jst *STPingjia) GetFactoryName() string {
	return "HT"
}

func (jst *STPingjia) StartEsitimate(pin *driverlayer.DriverArg, timeout int) (string, error) {
	utils.Debug("STPingjia received StartEsitimate request")
	s, _ := hex.DecodeString("1b5b3452")
	t, _ := hex.DecodeString("1b5b3443")
	utils.Debug("STPingjia send StartEsitimate pi %x", s)

	b, e := driverlayer.WritePortAndReadWithLen(pin.Port, pin.Baud, s, 3, timeout, t)
	if e == nil {
		utils.Debug("received data %x", b)
	}
	switch string(b) {
	case "B1F":
		return "非常满意", nil
	case "B2F":
		return "满意", nil
	case "B3F":
		return "不满意", nil
	case "":
		return "超时", PINGJIARESULT_EXCEED_MAXTIME
	}
	utils.Error("unexpected pingjia value %s", string(b))
	return "", PINGJIARESULT_NOT_FORMULAED
}

func (jst *STPingjia) CancelEsitimate(pin *driverlayer.DriverArg, timeout int) error {
	utils.Debug("STPingjia received CancelEsitimate request")
	t, _ := hex.DecodeString("1b5b3443")
	utils.Debug("STPingjia send CancelEsitimate pi %x", t)

	e := driverlayer.WritePortData(pin.Port, pin.Baud, t)

	utils.Debug("STPingjia CancelEsitimate")
	return e
}

func (jst *STPingjia) Reset(pin *driverlayer.DriverArg, timeout int) error {
	return nil
}
