// 深圳九思泰达密码键盘的实现
//
package bankcard

import (
	"cn.agree/driverlayer"
	"cn.agree/utils"
	"encoding/hex"
	"errors"
)

var CARDNO_NOT_FORMULAED = errors.New("card no format not standard")

//磁条卡com口实达实现
type SDMsf struct {
}

func (jst *SDMsf) Initdriver(pin *driverlayer.DriverArg) {

}

func (jst *SDMsf) Deinit(pin *driverlayer.DriverArg) error {
	return nil
}

func (jst *SDMsf) GetFactoryName() string {
	return "实达"
}

//向com口中写入任意数据即可
func (jst *SDMsf) Read(pin *driverlayer.DriverArg, read_type int, timeout int) ([]byte, error) {
	utils.Debug("SDMsf received read request")
	if read_type < 1 || read_type > 3 {
		utils.Error("not implementd read type %d,expected 1<=mode<=3", read_type)
		return nil, CARDNO_NOT_FORMULAED
	}
	var s []byte
	switch read_type {
	case 1:
		s, _ = hex.DecodeString("1b301b5d")
		break
	case 2:
		s, _ = hex.DecodeString("1b301b545d")
		break
	case 3:
		s, _ = hex.DecodeString("1b301b425d")
		break
	}
	b, e := driverlayer.WritePortAndReadWithTerminator(pin.Port, pin.Baud, s, []byte{0x3F}, timeout, nil)
	if e != nil {
		return nil, nil
	} else {
		var epos = 0
		if b[0] != 0x1b || b[1] != 0x73 {
			utils.Error("SDMsf data %s,left data not standard,expected %s", b, "1b73")
			return nil, CARDNO_NOT_FORMULAED
		}
		for i := 0; i < len(b); i++ {
			if b[i] == 0x41 && read_type == 3 {
				b[i] = '|'
			}
			if b[i] == 0x3f {
				epos = i
			}
		}
		b = b[2 : epos+1]

		return b, nil
	}

}
