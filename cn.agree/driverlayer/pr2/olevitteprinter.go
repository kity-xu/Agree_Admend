package pr2

import (
	"cn.agree/driverlayer"
	"cn.agree/hardware"
	"cn.agree/utils"
	"encoding/hex"
	"errors"
	"io"
	"strconv"
)

var PR2PARALARGERTHAN255 = errors.New("Para can not larger than 255")

//OKI打印机通用实现
type OlevittePrinter struct {
	w            io.ReadWriteCloser
	dotinstace   int
	lineinterval float32
	colinterval  float32
}

func (jst *OlevittePrinter) Initdriver(pin *driverlayer.DriverArg) {
	jst.dotinstace = 240
}

func (jst *OlevittePrinter) Deinit(pin *driverlayer.DriverArg) error {
	return nil
}

func (jst *OlevittePrinter) GetFactoryName() string {
	return "OlivittePr2"
}

func (jst *OlevittePrinter) BeginPrintJob(pin *driverlayer.DriverArg, timeout int) error {
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

func (jst *OlevittePrinter) checkJobInPrint() bool {
	if jst.w != nil {
		return true
	}
	return false
}

/*******************************xuxiaodong新增函数接口于2016-01-20******************************************/
//设置汉字方式。
func (jst *OkiPrinter) SetChinese() error {
	utils.Trace("OkiPrinter received SetChinese request")
	s, _ := hex.DecodeString("1C26")
	utils.Trace("OkiPrinter send SetChinese pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//取消中文方式
func (jst *OkiPrinter) CnacelChinese() error {
	utils.Trace("OkiPrinter received CnacelChinese request")
	s, _ := hex.DecodeString("1C2E")
	utils.Trace("OkiPrinter send CnacelChinese pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//设置全角ACSII方式
func (jst *OkiPrinter) SetSBC_case() error {
	utils.Trace("OkiPrinter received SetSBC_case request")
	s, _ := hex.DecodeString("1C6B")
	utils.Trace("OkiPrinter send SetSBC_case pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//取消全角ASCII
func (jst *OkiPrinter) CancelSBC_case() error {
	utils.Trace("OkiPrinter received CancelSBC_case request")
	s, _ := hex.DecodeString("1C67")
	utils.Trace("OkiPrinter send CancelSBC_case pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//设置汉字字体。0、1、2、3分别是：宋、仿、楷、繁。
func (jst *OkiPrinter) SetChineseTypeface() error {
	utils.Trace("OkiPrinter received SetChineseTypeface request")
	s, _ := hex.DecodeString("1C7430")
	utils.Trace("OkiPrinter send SetChineseTypeface pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//特殊货币符号打印
func (jst *OkiPrinter) Currency_symbol() error {
	return nil
}

//纸张大小。 原命令 1B-51-3x-3x-3x-1B-5A。打印机中有纸 命令不生效
func (jst *OkiPrinter) PaperSize() error {
	return nil
}

//左边界。   000为打印机默认边界
func (jst *OlevittePrinter) SetLeftMargin(pos float32) error {
	utils.Trace("OlevittePrinter received SetLeftMargin request")
	var line int
	line = int(pos / jst.colinterval)
	utils.Trace("SetLeftMargin:left margin is %d letters", line)
	if line > 255 {
		return PR2PARALARGERTHAN255
	}
	a := strconv.Itoa(line / 100)
	b := strconv.Itoa(line / 10 % 10)
	c := strconv.Itoa(line % 10)
	s, _ := hex.DecodeString("1b4a3" + a + "3" + b + "3" + c)
	utils.Trace("OlevittePrinter send SetLeftMargin pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//pr2在放入纸或存折之前可以使用的指令
//参数单位：mm
//命令参数：行
//纸长度7mm-500mm
//小于255行
func (jst *OlevittePrinter) SetPageHeight(pos float32) error {
	utils.Trace("OlevittePrinter received SetPageHeight request")
	line := int(pos / jst.lineinterval)
	utils.Trace("SetPageHeight:page height is %d row", line)
	if line > 255 {
		return PR2PARALARGERTHAN255
	}
	a := strconv.Itoa(line / 100)
	b := strconv.Itoa(line / 10 % 10)
	c := strconv.Itoa(line % 10)
	s, _ := hex.DecodeString("1b513" + a + "3" + b + "3" + c + "1b5a")
	utils.Trace("OlevittePrinter send SetLeftMargin pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//pr2在放入纸或存折之前可以使用的指令
//参数单位：mm
//命令参数：1/240inch
//命令参数不小于70，且不小于设备配置
func (jst *OlevittePrinter) SetTop(pos float32) error {
	utils.Trace("OlevittePrinter received SetTop request")
	line := int(utils.ConvertFrommmToInch(pos) * float32(240))
	utils.Trace("SetTop:top is %d mm", line)
	if line < 70 {
		return nil
	}
	a := strconv.Itoa(line / 100)
	b := strconv.Itoa(line / 10 % 10)
	c := strconv.Itoa(line % 10)
	s, _ := hex.DecodeString("1b543" + a + "3" + b + "3" + c)
	utils.Trace("OlevittePrinter send SetTop pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//pr2在放入纸或存折之前可以使用的指令
//指令是底部若干行不可打印
//参数单位：mm
//命令参数：行
func (jst *OlevittePrinter) SetPageBottom(posh float32, posb float32) error {
	pos := posb
	utils.Trace("OlevittePrinter received SetPageBottom request")
	line := int(pos / jst.lineinterval)
	utils.Trace("SetPageBottom:bottom is %d row", line)
	if line > 255 {
		return PR2PARALARGERTHAN255
	}
	a := strconv.Itoa(line / 100)
	b := strconv.Itoa(line / 10 % 10)
	c := strconv.Itoa(line % 10)
	s, _ := hex.DecodeString("1b4d3" + a + "3" + b + "3" + c)
	utils.Trace("OlevittePrinter send SetPageBottom pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//var l int = int(utils.ConvertFrommmToInch(pos) * float32(180))
//设置行据。 1B-26-30-30  nn/240或nn/216英寸，此处为缺省值1/6英寸
func (jst *OkiPrinter) SetLineInterval(pos float32) error {

	return nil
}

//0<linefactor<2
// TODO 指令错的
func (jst *OlevittePrinter) SetLineInterval(linefactor float32) error {
	utils.Debug("OlevittePrinter received SetLineInterval request")
	if linefactor < 0 || linefactor > 2 {
		return LINEINTERVALTOOLARGE
	}
	var line int
	line = int(linefactor * 120)
	realinterval := strconv.Itoa(line)
	s, _ := hex.DecodeString("1b253900" + realinterval)
	utils.Debug("OlevittePrinter send SetLineInterval pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}

}

//设置字距。 1B3C 为10字符/英寸的字距（英文）； 1B3D 12/英寸；1B3E 16.6/英寸；1B-61-3x 可变字距
func (jst *OkiPrinter) Word_spacing() error {
	return nil
}

//设置字形（包括草稿、高速草稿、斜体、斜体草稿及各个国家的字符集）。1B-52-30-30-30 （000为草稿）
func (jst *OkiPrinter) SetFont() error {
	return nil
}

//倍宽
func (jst *OlevittePrinter) DoublePrintWidth() error {
	utils.Debug("OlevittePrinter received DoublePrintWidth request")
	s, _ := hex.DecodeString("1b33")
	utils.Debug("OlevittePrinter send DoublePrintWidth pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//取消倍宽
func (jst *OlevittePrinter) CancelDoublePrintWidth() error {
	utils.Debug("OlevittePrinter received CancelDoublePrintWidth request")
	s, _ := hex.DecodeString("1b34")
	utils.Debug("OlevittePrinter send CancelDoublePrintWidth pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//SetDouble_width_high。 倍宽倍高打印
func (jst *OlevittePrinter) SetDouble_width_high() error {
	utils.Debug("OlevittePrinter received SetDouble_width_high request")
	s, _ := hex.DecodeString("1b64")
	utils.Debug("OlevittePrinter send SetDouble_width_high pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//Cancel_width_high。 取消倍宽倍高打印
func (jst *OlevittePrinter) Cancel_width_high() error {
	utils.Debug("OlevittePrinter received Cancel_width_high request")
	s, _ := hex.DecodeString("1b65")
	utils.Debug("OlevittePrinter send Cancel_width_high pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//三倍宽打印
func (jst *OlevittePrinter) SetThreeWidth() error {
	utils.Debug("OlevittePrinter received SetThreeWidth request")
	s, _ := hex.DecodeString("1b68")
	utils.Debug("OlevittePrinter send SetThreeWidth pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//取消三倍宽打印
func (jst *OlevittePrinter) CancelThreeWidth() error {
	utils.Debug("OlevittePrinter received CancelThreeWidth request")
	s, _ := hex.DecodeString("1b6a")
	utils.Debug("OlevittePrinter send CancelThreeWidth pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//倍高
func (jst *OkiPrinter) DoublePrintHigh() error {
	utils.Debug("OlevittePrinter received DoublePrintHigh request")
	s, _ := hex.DecodeString("1b7731")
	utils.Debug("OlevittePrinter send DoublePrintHigh pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//三倍高
func (jst *OkiPrinter) SetThreeHigh() error {
	utils.Debug("OlevittePrinter received SetThreeHigh request")
	s, _ := hex.DecodeString("1b7732")
	utils.Debug("OlevittePrinter send SetThreeHigh pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//取消纵向扩展（即三倍高、倍高）
func (jst *OkiPrinter) CancelLengthways() error {
	utils.Debug("OlevittePrinter received CancelLengthways request")
	s, _ := hex.DecodeString("1b7730")
	utils.Debug("OlevittePrinter send CancelLengthways pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//设置加重打印（即黑体）
func (jst *OlevittePrinter) SetBold() error {
	utils.Debug("OlevittePrinter received SetBold request")
	s, _ := hex.DecodeString("1b28")
	utils.Debug("OlevittePrinter send SetBold pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//取消黑体
func (jst *OlevittePrinter) CancelBold() error {
	utils.Debug("OlevittePrinter received CancelBold request")
	s, _ := hex.DecodeString("1b29")
	utils.Debug("OlevittePrinter send CancelBold pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//背景打印。最后一位：0取消；1逆打印；2密网打印；3梳网打印
func (jst *OkiPrinter) BackgroundPrint() error {
	utils.Debug("OlevittePrinter received BackgroundPrint request")
	s, _ := hex.DecodeString("1C2831")
	utils.Debug("OlevittePrinter send BackgroundPrint pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//下划线
func (jst *OlevittePrinter) SetUnderline() error {
	utils.Debug("OlevittePrinter received SetUnderline request")
	s, _ := hex.DecodeString("1b2a30")
	utils.Debug("OlevittePrinter send SetUnderline pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//上划线
func (jst *OlevittePrinter) SetSuperline() error {
	utils.Trace("OlevittePrinter received SetSubline request")
	s, _ := hex.DecodeString("1b2a33")
	utils.Trace("OlevittePrinter send SetSubline pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//设置上下划线
func (jst *OlevittePrinter) SetSuperAndUnderline() error {
	utils.Trace("OlevittePrinter received SetSuperAndUnderline request")
	s, _ := hex.DecodeString("1b2a34")
	utils.Trace("OlevittePrinter send SetSuperAndUnderline pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//取消上/下划线
func (jst *OlevittePrinter) CancelSuperUnderline() error {
	utils.Debug("OlevittePrinter received CancelSuperOrUnderline request")
	s, _ := hex.DecodeString("1b2b")
	utils.Debug("OlevittePrinter send CancelSuperOrUnderline pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//设置上标打印
func (jst *OkiPrinter) SetSuperscript() error {
	utils.Debug("OlevittePrinter received SetSuperscript request")
	s, _ := hex.DecodeString("1B6030")
	utils.Debug("OlevittePrinter send SetSuperscript pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//设置下标打印
func (jst *OkiPrinter) SetSubscript() error {
	utils.Debug("OlevittePrinter received SetSubscript request")
	s, _ := hex.DecodeString("1B6031")
	utils.Debug("OlevittePrinter send SetSubscript pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//清除上/下标打印
func (jst *OkiPrinter) CancelSuperSubscript() error {
	utils.Debug("OlevittePrinter received CancelSuperSubscript request")
	s, _ := hex.DecodeString("1B7B")
	utils.Debug("OlevittePrinter send CancelSuperSubscript pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//设定存折打印机
func (jst *OkiPrinter) SetBankbookPrinter() error {
	utils.Debug("OlevittePrinter received SetBankbookPrinter request")
	s, _ := hex.DecodeString("1B5335")
	utils.Debug("OlevittePrinter send SetBankbookPrinter pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//设定流水打印机
func (jst *OkiPrinter) SetGlidePrinter() error {
	utils.Debug("OlevittePrinter received SetGlidePrinter request")
	s, _ := hex.DecodeString("1B5331")
	utils.Debug("OlevittePrinter send SetGlidePrinter pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//换行
func (jst *OlevittePrinter) ChangeLine() error {
	utils.Debug("OlevittePrinter received ChangeLine request")
	s, _ := hex.DecodeString("0A20")
	utils.Debug("OlevittePrinter send ChangeLine pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//退纸
func (jst *OlevittePrinter) EjectPaper() error {
	utils.Trace("OlevittePrinter received EjectPaper request")
	s, _ := hex.DecodeString("1b4f")
	utils.Trace("OlevittePrinter send EjectPaper pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//换页
func (jst *OlevittePrinter) ChangePage() error {
	utils.Debug("OlevittePrinter received ChangePage request")
	s, _ := hex.DecodeString("0c")
	utils.Debug("OlevittePrinter send ChangePage pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//回车
func (jst *OlevittePrinter) Carriage() error {
	utils.Trace("OlevittePrinter received Carriage request")
	s, _ := hex.DecodeString("0d20")
	utils.Trace("OlevittePrinter send Carriage pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//水平制表
func (js *OkiPrinter) HorizontalTab() error {
	utils.Trace("OlevittePrinter received HorizontalTab request")
	s, _ := hex.DecodeString("09")
	utils.Trace("OlevittePrinter send HorizontalTab pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//反向换行
func (jst *OkiPrinter) ReverseChangeLine() error {
	utils.Trace("OlevittePrinter received ReverseChangeLine request")
	s, _ := hex.DecodeString("1b37")
	utils.Trace("OlevittePrinter send ReverseChangeLine pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//add by yangxiaolong
//绝对水平定位。 打印头左右移动
//单位:字符(一个英文字符为1mm)
func (jst *OlevittePrinter) SeekLinePos(incrementpos float32, currpos float32) error {
	utils.Debug("OlevittePrinter received SeekLinePos request")
	if currpos == 0 {
		utils.Trace("currpos is zero ,skip")
		return nil
	}
	line := int(incrementpos + currpos)
	if currpos > 0 {
		a := strconv.Itoa(line / 100)
		b := strconv.Itoa(line / 10 % 10)
		c := strconv.Itoa(line % 10)
		s, _ := hex.DecodeString("1b483" + a + "3" + b + "3" + c)
		if jst.checkJobInPrint() {
			return driverlayer.WriteData(jst.w, s)
		} else {
			return PR2NOTINJOB
		}
	} else {
		return PR2NOTIMPLEMENTED
	}
}

//绝对垂直定位。 打印头下移
//假设一行高3mm
//单位：行
func (jst *OlevittePrinter) AdvanceLine(pos float32) error {
	utils.Trace("OlevittePrinter received AdvanceLine request,pos [%f]", pos)
	if pos == 0 {
		utils.Trace("pos is zero ,skip")
		return nil
	}
	line := int(pos / 3) //3mm
	if pos > 0 {
		a := strconv.Itoa(line / 100)
		b := strconv.Itoa(line / 10 % 10)
		c := strconv.Itoa(line % 10)
		s, _ := hex.DecodeString("1b4c3" + a + "3" + b + "3" + c + "20")
		utils.Trace("OlevittePrinter send SeekColPos pi %x", s)
		if jst.checkJobInPrint() {
			driverlayer.WriteData(jst.w, s)
		} else {
			return PR2NOTINJOB
		}
	} else {
		return PR2NOTINJOB
	}

	return nil
}

//相对垂直定位
func (jst *OkiPrinter) RelativeVerticalSite(pos float32) error {
	utils.Trace("OlevittePrinter received RelativeVerticalSite request,pos [%f]", pos)
	if pos == 0 {
		utils.Trace("pos is zero ,skip")
		return nil
	}
	line := int(pos / 3) //3mm
	if pos > 0 {
		a := strconv.Itoa(line / 100)
		b := strconv.Itoa(line / 10 % 10)
		c := strconv.Itoa(line % 10)
		s, _ := hex.DecodeString("1b493" + a + "3" + b + "3" + c + "20")
		utils.Trace("OlevittePrinter send SeekColPos pi %x", s)
		if jst.checkJobInPrint() {
			driverlayer.WriteData(jst.w, s)
		} else {
			return PR2NOTINJOB
		}
	} else {
		return PR2NOTINJOB
	}

	return nil
}

//改变仿真方式到oki
func (jst *OkiPrinter) ChangeToOKI() error {
	utils.Trace("OlevittePrinter received ChangeToOKI request")
	s, _ := hex.DecodeString("1B5E31")
	utils.Trace("OlevittePrinter send ChangeToOKI pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//选择字符集。1B5B303030 为国际字符集
func (jst *OkiPrinter) CharacterGather() error {
	return nil
}

//响铃
func (jst *OkiPrinter) Diabolo() error {
	utils.Trace("OlevittePrinter received Diabolo request")
	s, _ := hex.DecodeString("07")
	utils.Trace("OlevittePrinter send Diabolo pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//清除打印缓冲区及其中的命令,清除含有ESC系列的代码。
func (jst *OkiPrinter) Cleanup() error {
	utils.Trace("OlevittePrinter received Cleanup request")
	s, _ := hex.DecodeString("7f")
	utils.Trace("OlevittePrinter send Cleanup pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//用户自定义下载命令集  略...

//图形打印方式
func (jst *OkiPrinter) SetGraphicPrint() error {
	return nil
}

//取消图形打印方式
func (jst *OkiPrinter) CancelGraphicPrint() error {
	return nil
}

//设置图形打印的相对水平位置
func (jst *OkiPrinter) RelativeHorizontalSite() error {
	return nil
}

//打印针脚设置
func (jst *OkiPrinter) CortrolPrinter_Spicula() error {
	return nil
}

//打印机初始识别请求。 由主机发出,要求打印机回答它的 ID 标识。。
func (jst *OkiPrinter) PrinterInitIdentifyRequest() error {
	utils.Trace("OlevittePrinter received PrinterInitIdentifyRequest request")
	s, _ := hex.DecodeString("1b5a")
	utils.Trace("OlevittePrinter send PrinterInitIdentifyRequest pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//打印机识别。 由打印机发出,作为对命令 ESC Z 的应答
func (jst *OkiPrinter) PrinterIdentify() error {
	utils.Trace("OlevittePrinter received PrinterIdentify request")
	s, _ := hex.DecodeString("1B2F37")
	utils.Trace("OlevittePrinter send PrinterIdentify pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//打印机配置请求。命令由主机发出,要求打印机回答其基本配置。
func (jst *OkiPrinter) PrinterConfigureRequest() error {
	utils.Trace("OlevittePrinter received PrinterConfigureRequest request")
	s, _ := hex.DecodeString("1B69")
	utils.Trace("OlevittePrinter send PrinterConfigureRequest pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//打印机配置 1B-70-xx-xx-xx--xx。 见pr2文档
func (jst *OkiPrinter) PrinterConfigure() error {
	return nil
}

//打印机状态请求。由主机发出,要求打印机传送同步状态码回答主机。
func (jst *OkiPrinter) PrinterStateRequest() error {
	utils.Trace("OlevittePrinter received PrinterStateRequest request")
	s, _ := hex.DecodeString("1B6A")
	utils.Trace("OlevittePrinter send PrinterStateRequest pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//1B-72-xx. 打印机同步状态
func (jst *OkiPrinter) PrinterSynchronizationState() error {
	return nil
}

//打印介质请求状态。 由主机发出,要求打印机回答关于打印介质当前所处位置的状态。
func (jst *OkiPrinter) PrinterPaperStateRequest() error {
	utils.Trace("OlevittePrinter received PrinterPaperStateRequest request")
	s, _ := hex.DecodeString("1B2042")
	utils.Trace("OlevittePrinter send PrinterPaperStateRequest pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//打印介质同步状态 1B-42-35-xx-1B-5A . 打印介质同步状态
func (jst *OkiPrinter) PrinterPaperSynchronizationState() error {
	return nil
}

//SET-UP 配置请求。 由主机发出,要求打印机回答它自身的 SET-UP 配置
func (jst *OkiPrinter) SER_UP_ConfiguerRequest() error {
	utils.Trace("OlevittePrinter received SER_UP_ConfiguerRequest request")
	s, _ := hex.DecodeString("1B2061")
	utils.Trace("OlevittePrinter send SER_UP_ConfiguerRequest pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//清楚错误状态
func (jst *OkiPrinter) CleanErrorState() error {
	utils.Trace("OlevittePrinter received CleanErrorState request")
	s, _ := hex.DecodeString("1B6C")
	utils.Trace("OlevittePrinter send CleanErrorState pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//选择Olivetti对话进程  控制的数据交换进程(仅在 OLIVETTI 仿真方式时才接受)
func (jst *OkiPrinter) ChooseOlivettiProcess() error {
	utils.Trace("OlevittePrinter received ChooseOlivettiProcess request")
	s, _ := hex.DecodeString("1B6E")
	utils.Trace("OlevittePrinter send ChooseOlivettiProcess pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//总清。 对打印机的固件和硬件进行复位
func (jst *OkiPrinter) CleanAll() error {
	utils.Trace("OlevittePrinter received CleanAll request")
	s, _ := hex.DecodeString("1B30")
	utils.Trace("OlevittePrinter send CleanAll pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//操作员请求 1B-55-xx. 通过设定参数"xx",确定操作员请求并打开或关闭相应的"STATION"灯
func (jst *OkiPrinter) OperatorRequest() error {
	return nil
}

//操作员应答 1B-72-xx。 由打印机发出,回答操作员请求,可以按以下两情况进行:
func (jst *OkiPrinter) OperatorResponsion() error {
	return nil
}

//赋予自动操作。设定打印机,使印机在打印介质一旦被插入进纸器后,无须按面板上的任何按键就能自动操作员请求命令给以回
func (jst *OkiPrinter) PrinterAutomatismControl() error {
	utils.Trace("OlevittePrinter received PrinterAutomatismControl request")
	s, _ := hex.DecodeString("1B2E")
	utils.Trace("OlevittePrinter send PrinterAutomatismControl pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//赋予手动操作
func (jst *OkiPrinter) PrinterManualControl() error {
	utils.Trace("OlevittePrinter received PrinterManualControl request")
	s, _ := hex.DecodeString("1B5F")
	utils.Trace("OlevittePrinter send PrinterManualControl pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

/********************************************OVER********************************************************/

//打印头上下移动
//绝对垂直定位
//单位:行
func (jst *OlevittePrinter) SeekColPos(incrementpos float32, currpos float32) error {
	utils.Debug("OlevittePrinter received SeekColPos request")
	if currpos == 0 {
		utils.Trace("currpos is zero ,skip")
		return nil
	}
	pos := incrementpos + currpos
	return jst.AdvanceLine(pos)
}

//0<line<165
//TODO 不懂，需要改
func (jst *OlevittePrinter) SkipLine(line int) error {
	utils.Debug("OlevittePrinter received SkipLine request")
	if line > 165 || line < 0 {
		return LINESKIPTOOLARGE
	}
	linef := line / 10
	linee := line % 10

	reallinef := strconv.Itoa(linef)
	reallinee := strconv.Itoa(linee)
	s, _ := hex.DecodeString("1b0b" + reallinef + reallinee) //垂直tab？
	utils.Debug("OlevittePrinter send SkipLine pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

func (jst *OlevittePrinter) CancelBold() error {
	utils.Debug("OlevittePrinter received CancelBold request")
	s, _ := hex.DecodeString("1b29")
	utils.Debug("OlevittePrinter send CancelBold pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

func (jst *OlevittePrinter) OutputString(s string) error {
	utils.Debug("OlevittePrinter received OutputString request")
	ds, _ := utils.TransGBKFromUTF8(s)
	utils.Debug("OlevittePrinter send OutputString pi %x", ds)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, ds)
	} else {
		return PR2NOTINJOB
	}
}

func (jst *OlevittePrinter) OutputByte(s []byte) error {
	utils.Debug("OlevittePrinter received OutputByte request")
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

func (jst *OlevittePrinter) EndPrinterJob() {
	if jst.w != nil {
		jst.w.Close()
		jst.w = nil
	}
}

func (jst *OlevittePrinter) Init() error {
	utils.Debug("OlevittePrinter received Init request")
	s, _ := hex.DecodeString("ff1b6c")
	utils.Debug("OlevittePrinter send Init pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}
func (jst *OlevittePrinter) AdvanceOneLine() error {
	utils.Debug("OlevittePrinter received Initfirstline request")
	s, _ := hex.DecodeString("0d0a20")
	utils.Debug("OlevittePrinter send Initfirstline pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//在一行内设置打印头水平方向绝对位置
func (jst *OlevittePrinter) SeekRealLinePos(pos float32) error {
	if pos == 0 {
		utils.Trace("pos is zero ,skip")
		return nil
	}
	return jst.SeekLinePos(0, pos)
}

//设置打印头竖直方向绝对位置
func (jst *OlevittePrinter) SeekRealColPos(pos float32) error {
	if pos == 0 {
		utils.Trace("pos is zero ,skip")
		return nil
	}
	return jst.SeekColPos(0, pos)
}

//TODO
func (jst *OlevittePrinter) InjectPaper() error {
	utils.Trace("OlevittePrinter received InjectPaper request")
	if jst.checkJobInPrint() {
		return nil
	} else {
		return PR2NOTINJOB
	}
}

//设置字符间距
//单位mm
func (jst *OlevittePrinter) SetColInterval(w float32) error {
	utils.Trace("OlevittePrinter received SetColInterval request")
	var line int
	line = int(utils.ConvertFrommmToInch(w) * float32(jst.dotinstace))
	utils.Trace("SetColInterval:col interval is %d/240 inch", line)
	a := strconv.Itoa(line / 10)
	b := strconv.Itoa(line % 10)
	s, _ := hex.DecodeString("1c533" + a + "3" + b)
	utils.Trace("OlevittePrinter send SetColInterval pi %x", s)
	if jst.checkJobInPrint() {
		return driverlayer.WriteData(jst.w, s)
	} else {
		return PR2NOTINJOB
	}
}

//TODO
func (jst *OlevittePrinter) SetPageWidth(pos float32) error {
	utils.Trace("OlevittePrinter received SetPageWidth request")
	if jst.checkJobInPrint() {
		return nil
	} else {
		return PR2NOTINJOB
	}
}

//TODO
//没有这个
func (jst *OlevittePrinter) SetRightMargin(pos float32) error {
	utils.Trace("OlevittePrinter received SetRightMargin request")
	if jst.checkJobInPrint() {
		return nil
	} else {
		return PR2NOTINJOB
	}
}
