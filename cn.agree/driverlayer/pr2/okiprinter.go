package pr2

import (
	"cn.agree/driverlayer"
	"cn.agree/hardware"
	"cn.agree/utils"
	"encoding/hex"
	"errors"
	"io"
)

var PR2ALREADYINJOB = errors.New("pr2 printer is in a job")
var PR2NOTINJOB = errors.New("pr2 not in a job,please acquire a job")
var LINEINTERVALTOOLARGE = errors.New("pr2 line interval is in 0<linefactor<2")
var LINESKIPTOOLARGE = errors.New("pr2 line skip maxinum value 165")
var PR2NOTIMPLEMENTED = errors.New("not implemented")
var PR2TOPTOOSMALL = errors.New("pr2's top margin is not smaller than 6.35mm")
var PR2RIGHTTOOLARGE = errors.New("pr2's right margin is larger than 238mm")

//OKI打印机通用实现
type OkiPrinter struct {
	w            io.ReadWriteCloser
	dotinstace   int
	lineinterval float32
	colinterval  float32
}

//保存接口
func (jst *OkiPrinter) Initdriver(pin *driverlayer.DriverArg) {
	jst.dotinstace = 180
}

//关闭接口
func (jst *OkiPrinter) Deinit(pin *driverlayer.DriverArg) error {
	return nil
}

func (jst *OkiPrinter) GetFactoryName() string {
	return "OkiPr2"
}

func (jst *OkiPrinter) BeginPrintJob(pin *driverlayer.DriverArg, timeout int) error {
	if jst.w != nil {
		return PR2ALREADYINJOB
	}
	w, err := hardware.GetPortInstance(pin.Port, pin.Baud)
	if err != nil {
		return err
	}
	jst.w = w
	return nil
}

func (jst *OkiPrinter) checkJobInPrint() bool {
	if jst.w != nil {
		return true
	}
	return false
}

/*******************************xuxiaodong整理、新增函数接口于2016-01-20*************************************/
//设备控制1
func (jst *OkiPrinter) DeviceControl1() error {
	utils.Trace("OkiPrinter received DeviceControl1 request")
	s, _ := hex.DecodeString("11")
	utils.Trace("OkiPrinter send DeviceControl1 pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//设备控制3
func (jst *OkiPrinter) DeviceControl3() error {
	utils.Trace("OkiPrinter received DeviceControl3 request")
	s, _ := hex.DecodeString("13")
	utils.Trace("OkiPrinter send DeviceControl3 pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//清除打印机缓冲
func (jst *OkiPrinter) CleanPrinterRom() error {
	utils.Trace("OkiPrinter received CleanPrinterRom request")
	s, _ := hex.DecodeString("18")
	utils.Trace("OkiPrinter send CleanPrinterRom pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//设置SetShift_jis 方式
func (jst *OkiPrinter) SetShift_jis() error {
	utils.Trace("OkiPrinter received SetShift_jis request")
	s, _ := hex.DecodeString("1B6B")
	utils.Trace("OkiPrinter send SetShift_jis pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//取消SetShift_jis 方式
func (jst *OkiPrinter) CancelShift_jis() error {
	utils.Trace("OkiPrinter received CancelShift_jis request")
	s, _ := hex.DecodeString("1B6C")
	utils.Trace("OkiPrinter send CancelShift_jis pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//换行
func (jst *OkiPrinter) ChangeLine() error {
	utils.Trace("OkiPrinter received ChangeLine request")
	s, _ := hex.DecodeString("0A20")
	utils.Trace("OkiPrinter send ChangeLine pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//进纸一行
func (jst *OkiPrinter) AdvanceOneLine() error {
	utils.Debug("OkiPrinter received Initfirstline request")
	s, _ := hex.DecodeString("0d0a20")
	utils.Debug("OkiPrinter send Initfirstline pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//设置1/6英寸行据
func (jst *OkiPrinter) Rowledge_6LPI() error {
	utils.Trace("OkiPrinter received Rowledge_6LPI request")
	s, _ := hex.DecodeString("1B36")
	utils.Trace("OkiPrinter send Rowledge_6LPI pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//设置1/8英寸行据
func (jst *OkiPrinter) Rowledge_8LPI() error {
	utils.Trace("OkiPrinter received Rowledge_8LPI request")
	s, _ := hex.DecodeString("1B38")
	utils.Trace("OkiPrinter send Rowledge_8LPI pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//设置n/120英寸行据。1B-25-39-n1-n2（n1为00,n2表示：00-FF/120）
func (jst *OkiPrinter) SetLineInterval(lineinterval float32) error {
	utils.Trace("OkiPrinter received SetLineInterval request")

	var line int
	line = int(utils.ConvertFrommmToInch(lineinterval) * float32(120))
	if line > 255 {
		return LINEINTERVALTOOLARGE
	}
	utils.Trace("SetLineInterval:line interval is %d/120 inch", line)
	s, _ := hex.DecodeString("1b2539" + utils.FormatInt(line, 4, "0", 16))
	utils.Trace("OkiPrinter send SetLineInterval pi %x", s)
	if jst.checkJobInPrint() {
		jst.lineinterval = lineinterval
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}

}

//直接进纸若干行。1B-0B-n1-n2（行数=n1×10+n2）
func (jst *OkiPrinter) AdvanceLine(pos uint) error {
	utils.Trace("OkiPrinter received PaperIn_some request")
	if pos < 10 {
		s, _ := hex.DecodeString("1B0B0" + utils.FormatInt(pos, 2, "0", 16))
	}
	if 9 < pos {
		s, _ := hex.DecodeString("1B0B" + utils.FormatInt(pos, 4, "0", 16))
	}
	utils.Trace("OkiPrinter send PaperIn_some pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//直接进纸n/120英寸。1B-25-35-n（01 =< n <= FF）
func (jst *OkiPrinter) SeekColPos(pos float32, currpos float32) error {
	if pos == 0 {
		utils.Trace("pos is zero ,skip")
		return nil
	}

	inch := int(utils.ConvertFrommmToInch(pos) * float32(120))
	remain := inch
	for {
		//oki的最大位移为255/120 inch
		if remain > 255 {
			s, _ := hex.DecodeString("1b2535ff")
			utils.Trace("OkiPrinter send AdvanceLine pi %x", s)
			if jst.checkJobInPrint() {
				driverlayer.WriteData(jst.w, s)
			} else {
				return PR2NOTINJOB
			}
			remain -= 255
		} else {
			var d uint8 = uint8(remain) & 0xff
			s, _ := hex.DecodeString("1b2535" + utils.FormatInt(int(d), 2, "0", 16) + "20")
			utils.Trace("OkiPrinter send AdvanceLine pi %x", s)
			if jst.checkJobInPrint() {
				return driverlayer.WriteData(jst.w, s)
			} else {
				return PR2NOTINJOB
			}
			break
		}
	}
	return nil
}

//单位：毫米.
//注：由于OKI并没有设置头部的指令。(仅有的一个也只能设置在6.35毫米处).因此要在打印一页纸前调用
//此指令用进纸指令代替
func (jst *OkiPrinter) SetTop(pos float32) error {
	utils.Trace("OkiPrinter received SetTop request")
	if pos < 6.35 {
		return PR2TOPTOOSMALL
	}
	var reallen = pos - 6.35
	utils.Trace("OkiPrinter not have settop instructon,advance line %f ", reallen)
	return jst.SeekColPos(reallen, 6.35)
}

//设置页长。毫米
func (jst *OkiPrinter) SetPageHeight(pos float32) error {
	utils.Trace("OkiPrinter received SetPageHeight request")
	var d int = int(pos / jst.lineinterval)
	var md int = d / 16
	var ld int = d % 16

	b := "1b46" + utils.FormatInt(md, 2, "0", 16) + utils.FormatInt(ld, 2, "0", 16)
	s, _ := hex.DecodeString(b)
	utils.Trace("OkiPrinter send SetPageHeight pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//换页
func (jst *OkiPrinter) EjectPaper() error {
	utils.Trace("OkiPrinter received EjectPaper request")
	s, _ := hex.DecodeString("0c")
	utils.Trace("OkiPrinter send EjectPaper pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//回车
func (jst *OkiPrinter) Carriage() error {
	utils.Trace("OkiPrinter received Carriage request")
	s, _ := hex.DecodeString("0d20")
	utils.Trace("OkiPrinter send Carriage pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}

}

//退格
func (jst *OkiPrinter) Backspace() error {
	utils.Trace("OkiPrinter received Backspace request")
	s, _ := hex.DecodeString("08")
	utils.Trace("OkiPrinter send Backspace pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//设置左边界     单位：毫米
func (jst *OkiPrinter) SetLeftMargin(pos float32) error {
	utils.Trace("OkiPrinter received SetLeftMargin request")
	var l int = int(utils.ConvertFrommmToInch(pos) * float32(180))
	if l > 1423 {
		utils.Error("OkiPrinter ;s right margin is not allowed to exceed than 177.8mm !!")
		return PR2RIGHTTOOLARGE
	}
	if l < 1 {
		l = 1
	}
	utils.Trace("SetLeftMargin:real left margin is %d/180 inch", l)
	b := "1b28" + utils.FormatInt(l, 4, "0", 16)
	s, _ := hex.DecodeString(b)
	utils.Trace("OkiPrinter send SetLeftMargin pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//设置有边界    单位：毫米
func (jst *OkiPrinter) SetRightMargin(pos float32) error {
	utils.Trace("OkiPrinter received SetRightMargin request")
	var l int = int(utils.ConvertFrommmToInch(pos) * float32(180))
	if l > 1690 {
		utils.Error("OkiPrinter ;s right margin is not allowed to exceed than 228.6mm !!")
		return PR2RIGHTTOOLARGE
	}
	utils.Trace("SetRightMargin:real right margin is %d/180 inch", l)
	b := "1b29" + utils.FormatInt(l, 4, "0", 16)
	s, _ := hex.DecodeString(b)
	utils.Trace("OkiPrinter send SetRightMargin pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//设置 Pica HS ANK 字符方式
func (jst *OkiPrinter) SetPica_HS_ANK() error {
	utils.Trace("OkiPrinter received SetPica_HS_ANK request")
	s, _ := hex.DecodeString("1B4E")
	utils.Trace("OkiPrinter send SetPica_HS_ANK pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//设置 Pica HD ANK 字符方式
func (jst *OkiPrinter) SetPica_HD_ANK() error {
	utils.Trace("OkiPrinter received SetPica_HD_ANK request")
	s, _ := hex.DecodeString("1B48")
	utils.Trace("OkiPrinter send SetPica_HD_ANK pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//设置 Elite HS ANK 字符方式
func (jst *OkiPrinter) SetElite_HS_ANK() error {
	utils.Trace("OkiPrinter received SetElite_HS_ANK request")
	s, _ := hex.DecodeString("1B42")
	utils.Trace("OkiPrinter send SetElite_HS_ANK pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//设置 Elite HD ANK 字符方式
func (jst *OkiPrinter) SetElite_HD_ANK() error {
	utils.Trace("OkiPrinter received SetElite_HD_ANK request")
	s, _ := hex.DecodeString("1B45")
	utils.Trace("OkiPrinter send SetElite_HD_ANK pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//设置 ANK 字符方式
func (jst *OkiPrinter) SetANK() error {
	utils.Trace("OkiPrinter received SetANK request")
	s, _ := hex.DecodeString("1B26")
	utils.Trace("OkiPrinter send SetANK pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//设置 ANK 字符片假名方式
func (jst *OkiPrinter) SetANK_Anonym() error {
	utils.Trace("OkiPrinter received SetANK_Anonym request")
	s, _ := hex.DecodeString("1B27")
	utils.Trace("OkiPrinter send SetANK_Anonym pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//设置高速打印方式. 打印机对于 HD ANK 字符,汉字间隔打印,因此打印速度大约为 HD 方式的两倍
func (jst *OkiPrinter) High_speedPrint() error {
	utils.Trace("OkiPrinter received High_speedPrint request")
	s, _ := hex.DecodeString("1B44")
	utils.Trace("OkiPrinter send High_speedPrint pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//设置高密打印方式. 打印机以高密方式打印 HD ANK 字符、汉字。
func (jst *OkiPrinter) High_densityPrint() error {
	utils.Trace("OkiPrinter received High_densityPrint request")
	s, _ := hex.DecodeString("1B49")
	utils.Trace("OkiPrinter send High_densityPrint pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//设置低噪音打印方式
func (jst *OkiPrinter) Low_noisePrint() error {
	utils.Trace("OkiPrinter received Low_noisePrint request")
	s, _ := hex.DecodeString("1B4F")
	utils.Trace("OkiPrinter send Low_noisePrint pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//设置下划线打印
func (jst *OkiPrinter) SetUnderline() error {
	utils.Trace("OkiPrinter received SetUnderline request")
	s, _ := hex.DecodeString("1b58")
	utils.Trace("OkiPrinter send SetUnderline pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//取消下划线打印
func (jst *OkiPrinter) CancelUnderline() error {
	utils.Trace("OkiPrinter received CancelUnderline request")
	s, _ := hex.DecodeString("1b59")
	utils.Trace("OkiPrinter send CancelUnderline pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//设置倍宽打印
func (jst *OkiPrinter) DoublePrintWidth() error {
	utils.Trace("OkiPrinter received DoublePrintWidth request")
	s, _ := hex.DecodeString("1b55")
	utils.Trace("OkiPrinter send DoublePrintWidth pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//取消倍宽打印
func (jst *OkiPrinter) CancelDoublePrintWidth() error {
	utils.Trace("OkiPrinter received CancelDoublePrintWidth request")
	s, _ := hex.DecodeString("1b52")
	utils.Trace("OkiPrinter send CancelDoublePrintWidth pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//设置全角汉字的双倍水平扩展打印方式
func (jst *OkiPrinter) SetSBC_caseDouble_width() error {
	utils.Trace("OkiPrinter received SetSBC_caseDouble_width request")
	s, _ := hex.DecodeString("1C70")
	utils.Trace("OkiPrinter send SetSBC_caseDouble_width pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//取消全角汉字的双倍水平扩展打印方式
func (jst *OkiPrinter) CancelSBC_caseDouble_width() error {
	utils.Trace("OkiPrinter received CancelSBC_caseDouble_width request")
	s, _ := hex.DecodeString("1C71")
	utils.Trace("OkiPrinter send CancelSBC_caseDouble_width pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//设置水平压缩打印方式(原来字符的1/2)
func (jst *OkiPrinter) SetHorizontalCompress() error {
	utils.Trace("OkiPrinter received SetHorizontalCompress request")
	s, _ := hex.DecodeString("1B3C")
	utils.Trace("OkiPrinter send SetHorizontalCompress pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

func (jst *OkiPrinter) CancelHorizontalCompress() error {
	utils.Trace("OkiPrinter received CancelHorizontalCompress request")
	s, _ := hex.DecodeString("1B3E")
	utils.Trace("OkiPrinter send CancelHorizontalCompress pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//设置倍高
func (jst *OkiPrinter) DoublePrintHeight() error {
	utils.Trace("OkiPrinter received DoublePrintHeight request")
	s, _ := hex.DecodeString("1b5b")
	utils.Trace("OkiPrinter send DoublePrintHeight pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//取消倍高
func (jst *OkiPrinter) CancelDoublePrintHeight() error {
	utils.Trace("OkiPrinter received CancelDoublePrintHeight request")
	s, _ := hex.DecodeString("1b5d")
	utils.Trace("OkiPrinter send CancelDoublePrintHeight pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//设置3倍高打印
func (jst *OkiPrinter) SetThreeHighPrint() error {
	utils.Trace("OkiPrinter received SetThreeHighPrint request")
	s, _ := hex.DecodeString("1b65")
	utils.Trace("OkiPrinter send SetThreeHighPrint pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//取消3倍高打印
func (jst *OkiPrinter) CancelThreeHighPrint() error {
	utils.Trace("OkiPrinter received CancelThreeHighPrint request")
	s, _ := hex.DecodeString("1b66")
	utils.Trace("OkiPrinter send CancelThreeHighPrint pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//设置3倍宽打印
func (jst *OkiPrinter) SetThreeWidthPrint() error {
	utils.Trace("OkiPrinter received SetThreeWidthPrint request")
	s, _ := hex.DecodeString("1b67")
	utils.Trace("OkiPrinter send SetThreeWidthPrint pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//取消3倍宽打印
func (jst *OkiPrinter) CancelThreeWidthPrint() error {
	utils.Trace("OkiPrinter received CancelTvhreeWidthPrint request")
	s, _ := hex.DecodeString("1b68")
	utils.Trace("OkiPrinter send CancelThreeWidthPrint pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//加重
func (jst *OkiPrinter) SetBold() error {
	utils.Trace("OkiPrinter received SetBold request")
	s, _ := hex.DecodeString("1b69")
	utils.Trace("OkiPrinter send SetBold pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//取消加重
func (jst *OkiPrinter) CancelBold() error {
	utils.Trace("OkiPrinter received CancelBold request")
	s, _ := hex.DecodeString("1b6a")
	utils.Trace("OkiPrinter send CancelBold pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//设置单向打印方式
func (jst *OkiPrinter) SetSinglenessOrientation() error {
	utils.Trace("OkiPrinter received SetSinglenessOrientation request")
	s, _ := hex.DecodeString("1b2555")
	utils.Trace("OkiPrinter send SetSinglenessOrientation pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//设置双向打印方式
func (jst *OkiPrinter) SetDiprosopyOrientation() error {
	utils.Trace("OkiPrinter received SetDiprosopyOrientation request")
	s, _ := hex.DecodeString("1b2542")
	utils.Trace("OkiPrinter send SetDiprosopyOrientation pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//设置重复打印方式
func (jst *OkiPrinter) SetRepeatPrint() error {
	utils.Trace("OkiPrinter received SetRepeatPrint request")
	s, _ := hex.DecodeString("1b6d")
	utils.Trace("OkiPrinter send SetRepeatPrint pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//取消重复打印方式
func (jst *OkiPrinter) CancelRepeatPrint() error {
	utils.Trace("OkiPrinter received CancelRepeatPrint request")
	s, _ := hex.DecodeString("1b6e")
	utils.Trace("OkiPrinter send CancelRepeatPrint pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//设置汉字方式。 本码设置打印机为汉字方式,如果已设置了半角字方式,半角字方式有效
func (jst *OkiPrinter) SetChinese() error {
	utils.Trace("OkiPrinter received SetChinese request")
	s, _ := hex.DecodeString("1B2440")
	utils.Trace("OkiPrinter send SetChinese pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//取消汉字方式。不取消半角方式
func (jst *OkiPrinter) CancelChinese() error {
	utils.Trace("OkiPrinter received CancelChinese request")
	s, _ := hex.DecodeString("1B2848")
	utils.Trace("OkiPrinter send CancelChinese pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//设置汉字字体。0、1、2、3分别是：宋、仿、楷、繁。
func (jst *OkiPrinter) SetChineseTypeface(font int) error {
	utils.Trace("OkiPrinter received SetChineseTypeface request")
	switch font {
	case 0:
		s, _ := hex.DecodeString("1C7630")
		break
	case 1:
		s, _ := hex.DecodeString("1C7631")
		break
	case 2:
		s, _ := hex.DecodeString("1C7632")
		break
	case 3:
		s, _ := hex.DecodeString("1C7633")
		break
	default:
		s, _ := hex.DecodeString("1C7630")
		break
	}
	utils.Trace("OkiPrinter send SetChineseTypeface pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//设置上标打印方式
func (jst *OkiPrinter) SetSuperscript() error {
	utils.Trace("OkiPrinter received SetSuperscript request")
	s, _ := hex.DecodeString("1C4E")
	utils.Trace("OkiPrinter send SetSuperscript pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//取消上标打印方式
func (jst *OkiPrinter) CancelSuperscript() error {
	utils.Trace("OkiPrinter received CancelSuperscript request")
	s, _ := hex.DecodeString("1C4F")
	utils.Trace("OkiPrinter send CancelSuperscript pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//设置下标打印
func (jst *OkiPrinter) SetSubscript() error {
	utils.Trace("OkiPrinter received SetSubscript request")
	s, _ := hex.DecodeString("1C50")
	utils.Trace("OkiPrinter send SetSubscript pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//取消下标打印
func (jst *OkiPrinter) CancelSubscript() error {
	utils.Trace("OkiPrinter received CancelSubscript request")
	s, _ := hex.DecodeString("1C51")
	utils.Trace("OkiPrinter send CancelSubscript pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//设置上下标打印方式
func (jst *OkiPrinter) SetSuper_Sub_Script() error {
	utils.Trace("OkiPrinter received SetSuper_Sub_Script request")
	s, _ := hex.DecodeString("1C52")
	utils.Trace("OkiPrinter send SetSuper_Sub_Script pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//取消上下标打印
func (jst *OkiPrinter) CancelSuper_Sub_Script() error {
	utils.Trace("OkiPrinter received CancelSuper_Sub_Script request")
	s, _ := hex.DecodeString("1C53")
	utils.Trace("OkiPrinter send CancelSuper_Sub_Script pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//设置半角字方式
func (jst *OkiPrinter) SetDBC_case() error {
	utils.Trace("OkiPrinter received SetDBC_case request")
	s, _ := hex.DecodeString("1C72")
	utils.Trace("OkiPrinter send SetDBC_case pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//取消半角字方式（即全角）
func (jst *OkiPrinter) CancelDBC_case() error {
	utils.Trace("OkiPrinter received CancelDBC_case request")
	s, _ := hex.DecodeString("1C73")
	utils.Trace("OkiPrinter send CancelDBC_case pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//设置汉字竖写方式
func (jst *OkiPrinter) SetVerticalChinese() error {
	utils.Trace("OkiPrinter received SetVerticalChinese request")
	s, _ := hex.DecodeString("1C4A")
	utils.Trace("OkiPrinter send SetVerticalChinese pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//设置汉字横向书写方式（即取消汉字竖写）
func (jst *OkiPrinter) SetHorizontalChinese() error {
	utils.Trace("OkiPrinter received SetHorizontalChinese request")
	s, _ := hex.DecodeString("1C4B")
	utils.Trace("OkiPrinter send SetHorizontalChinese pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//设置半角字合成方式。当在汉字竖向书写方式下接收到此码后,打印机将在此码后的两半角字字符合成起来作为合成字符。不持续有效
func (jst *OkiPrinter) SetDBC_caseCompose() error {
	utils.Trace("OkiPrinter received SetDBC_caseCompose request")
	s, _ := hex.DecodeString("1C5F")
	utils.Trace("OkiPrinter send SetDBC_caseCompose pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//禁止半角字字符竖写方式
func (jst *OkiPrinter) ForbidDBC_caseErect() error {
	utils.Trace("OkiPrinter received ForbidDBC_caseErect request")
	s, _ := hex.DecodeString("1C74")
	utils.Trace("OkiPrinter send ForbidDBC_caseErect pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//取消禁止半角字字符竖写方式
func (jst *OkiPrinter) CancelForbidDBC_caseErect() error {
	utils.Trace("OkiPrinter received CancelForbidDBC_caseErect request")
	s, _ := hex.DecodeString("1C75")
	utils.Trace("OkiPrinter send CancelForbidDBC_caseErect pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//设置单页纸出纸方向(30:向前出纸；31：向后出纸)
func (jst *OkiPrinter) OutPaperForward() error {
	utils.Trace("OkiPrinter received OutPaperDirection request")
	s, _ := hex.DecodeString("1B5430")
	utils.Trace("OkiPrinter send OutPaperDirection pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//设置单页纸出纸方向(30:向前出纸；31：向后出纸)
func (jst *OkiPrinter) OutPaperBackward() error {
	utils.Trace("OkiPrinter received OutPaperDirection request")
	s, _ := hex.DecodeString("1B5431")
	utils.Trace("OkiPrinter send OutPaperDirection pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//打印机初始化
func (jst *OkiPrinter) Init() error {
	utils.Debug("OlevittePrinter received initPrinter request")
	s, _ := hex.DecodeString("1042")
	utils.Debug("OlevittePrinter send initPrinter pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//改变仿真方式到 OLIVETTI 仿真方式
func (jst *OkiPrinter) ChanageToOlivetti() error {
	utils.Debug("OlevittePrinter received ChanageToOlivetti request")
	s, _ := hex.DecodeString("1B5E31")
	utils.Debug("OlevittePrinter send ChanageToOlivetti pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

func (jst *OkiPrinter) ChanageToIBM() error {
	utils.Debug("OlevittePrinter received ChanageToIBM request")
	s, _ := hex.DecodeString("1B5E30")
	utils.Debug("OlevittePrinter send ChanageToIBM pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

// //设置水平tab
// func (jst *OkiPrinter) SetHorizontalTab() error {
// 	return nil
// }

// //移至水平tab位置
// func (jst *OkiPrinter) MoveHorizontalTab() error {
// 	utils.Trace("OkiPrinter received MoveHorizontalTab request")
// 	s, _ := hex.DecodeString("09")
// 	utils.Trace("OkiPrinter send MoveHorizontalTab pi %x", s)
// 	if jst.checkJobInPrint() {
// 		return driverlayer.WriteData(jst.w, s)
// 	} else {
// 		return PR2NOTINJOB
// 	}
// }

//设置回车位置 1B-25-36-n1-n2
//var l int = int(utils.ConvertFrommmToInch(pos) * float32(180))
func (jst *OkiPrinter) SetPlaceOfEnter(pos float32) error {
	return nil
}

//图形数据传输。1B-25-31-n1-n2  本码定义图形打印数据,参数 n1 n2 定义了本码之后被传输的图形点列数目[1 列=3 字节(24点)]。
func (jst *OkiPrinter) FigureDataTransfers() error {
	return nil
}

//水平双倍扩展图形数据传输  1B-25-32-n1-n2
func (jst *OkiPrinter) HorizontalDoubleFigureDataTransfers() error {
	return nil
}

/********************************************OVER********************************************************/

//设置列距.是以列距为计算方式的(即字符间距)
//一般都是1/180
func (jst *OkiPrinter) SetColInterval(w float32) error {
	utils.Trace("OkiPrinter received SetColInterval request")

	var line int
	line = int(utils.ConvertFrommmToInch(w) * float32(jst.dotinstace)) //1/180
	if line > 255 {
		return LINEINTERVALTOOLARGE
	}
	utils.Trace("SetColInterval:col interval is %d/120 inch", line)

	realinterval := utils.FormatInt(line, 2, "0", 16)
	s, _ := hex.DecodeString("1c24" + realinterval)
	utils.Trace("OkiPrinter send SetColInterval pi %x", s)
	if jst.checkJobInPrint() {
		jst.colinterval = w
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}

}

//单位长度不一致,这个是以点距为单位的
//即：多少个字符
func (jst *OkiPrinter) SeekLinePos(pos float32, currpos float32) error {
	utils.Trace("OkiPrinter received SeekLinePos request")
	line := utils.ConvertFrommmToInch(pos) * float32(jst.dotinstace)
	return jst.SeekRealLinePos(line)
}

func (jst *OkiPrinter) OutputString(s string) error {
	utils.Trace("OkiPrinter received OutputString request")
	ds, _ := utils.TransGBKFromUTF8(s)
	utils.Trace("OkiPrinter send OutputString pi %x", ds)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, ds)
	} else {
		return PR2NOTINJOB
	}
}

func (jst *OkiPrinter) OutputByte(s []byte) error {
	utils.Trace("OkiPrinter received OutputByte request")
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

func (jst *OkiPrinter) EndPrinterJob() {
	if jst.w != nil {
		jst.w.Close()
		jst.w = nil
	}
}

//设置页底
func (jst *OkiPrinter) SetPageBottom(posh float32, posb float32) error {
	pos := posh - posb
	utils.Trace("OkiPrinter received SetPageHeight request")
	if jst.checkJobInPrint() {
		return jst.SetPageHeight(pos)
	} else {
		return PR2NOTINJOB
	}
}

//设置页宽。毫米
//TODO
func (jst *OkiPrinter) SetPageWidth(pos float32) error {
	utils.Trace("OkiPrinter received SetPageWidth request")
	if jst.checkJobInPrint() {
		return nil
	} else {
		return PR2NOTINJOB
	}
}

//在一行内设置打印头水平方向绝对位置
func (jst *OkiPrinter) SeekRealLinePos(pos float32) error {
	if pos > 0 {
		line1 := int(pos)
		line := utils.FormatInt(line1, 4, "0", 16)

		s, _ := hex.DecodeString("1b2533" + line)

		if jst.checkJobInPrint() {
			return driverlayer.WriteData(jst.w, s)
		} else {
			return PR2NOTINJOB
		}

	} else {
		return PR2NOTIMPLEMENTED
	}
}

//设置打印头竖直方向绝对位置
func (jst *OkiPrinter) SeekRealColPos(pos float32) error {
	return jst.SeekColPos(0, pos)
}
