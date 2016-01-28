package driverlayer

const (
	DEVICE_OPER_SUCCESS = iota //设备操作成功
	DEVICE_NOT_FOUND           //未找到设备
	DEVICE_OPEN_ERROR          //打开设备错误
	DEVICE_OPER_ERROR          //设备操作错误
	DEVICE_IN_USE              //设备正在使用
	DEVICE_OPER_TIMEOUT        //设备操作超时
)

//所有设备的共有接口
type DriverBase interface {
	Initdriver(pin *DriverArg)
	Deinit(pin *DriverArg) error
	GetFactoryName() string
}

type DriverExtraParam struct {
	Name  string
	Value string
}

//设备的操作结构
type DriverArg struct {
	Port        string
	ExtPort     string
	Baud        int
	FactoryName string
	ExtraParam  []DriverExtraParam
}

//密码键盘的接口
type IPin interface {
	DriverBase
	Readonce(pin *DriverArg, timeout int) ([]byte, error)  //进行语音提示,并且获取输入
	Readtwice(pin *DriverArg, timeout int) ([]byte, error) //进行语音提示,并且获取输入
	Reset(pin *DriverArg, timeout int) error               //重置密码键盘
}

//指纹操作器的接口
type IFinger interface {
	DriverBase
	GetRegisterFg(pin *DriverArg, timeout int) ([]byte, error)
	GetValidateFg(pin *DriverArg, timeout int) ([]byte, error)
	MatchFinger(pin *DriverArg, timeout int, reg []byte, vad []byte) int
	Reset(pin *DriverArg, timeout int) error //重置指纹仪
}

//身份证的接口
type IDCReader interface {
	DriverBase
	ReadData(pin *DriverArg, timeout int) ([]byte, error)
}

//磁卡接口
type IMsf interface {
	DriverBase
	Read(pin *DriverArg, readtype int, timeout int) ([]byte, error)
}

//IC卡接口
type ICReader interface {
	DriverBase
	PowerOff(pin *DriverArg, timeout int) error
	GetICCardInfo(pin *DriverArg, timeout int, taglist []byte, lpicappdata []byte) ([]byte, error)
	GenARQC(pin *DriverArg, timeout int, taglist []byte, lpicappdata []byte) ([]byte, error)
	CtrScriptData(pin *DriverArg, timeout int, taglist []byte, lpicappdata []byte, arpc []byte) ([]byte, error)
	GetTransDetail(pin *DriverArg, timeout int, path []byte) ([]byte, error)
}

//扫描仪接口
type IScan interface {
	DriverBase
	Read(pin *DriverArg, timeout int, data string) error
}

//评价器接口
type IPingjia interface {
	DriverBase
	StartEsitimate(pin *DriverArg, timeout int) (string, error)
	CancelEsitimate(pin *DriverArg, timeout int) error
	Reset(pin *DriverArg, timeout int) error //重置评价器
}

type IKvm interface {
	DriverBase
	//外屏同频内屏
	OutSyncInner(pin *DriverArg, timeout int) error
	//内屏同屏外屏
	InnSyncOuter(pin *DriverArg, timeout int) error
	//断屏
	DeSync(pin *DriverArg, timeout int) error
}

//pr2打印机接口
type IPr2Print interface {
	DriverBase

	DeviceControl1() error

	DeviceControl3() error

	SetShift_jis() error

	CancelShift_jis() error

	//清除打印机缓冲
	CleanPrinterRom() error

	//初始化
	Init() error

	//改变仿真方式到Olivetti
	ChanageToOlivetti() error

	//改变仿真方式到IBM
	ChanageToIBM() error

	/**
	 * 页面设置
	 */
	//设置行距
	Rowledge_6LPI() error            //1/6行距
	Rowledge_8LPI() error            //1/8行距
	SetLineInterval(w float32) error //n/120英寸行距

	//设置打印介质、
	ChangeLine() error                              //换行
	AdvanceOneLine() error                          //进纸一行
	AdvanceLine(pos float32) error                  //进纸若干行
	SetTop(pos float32) error                       //设置页顶。毫米
	SetPageHeight(pos float32) error                //设置页长。毫米
	SetPageBottom(posh float32, posb float32) error //设置页底。毫米
	EjectPaper() error                              //退纸（换页）
	Carriage() error                                //回车
	Backspace() error                               //退格
	SetLeftMargin(pos float32) error                //设置左边界.毫米
	SetRightMargin(pos float32) error               //设置右边界。毫米

	//上、下标打印
	SetUnderline() error           //设置下划线打印
	CancelUnderline() error        //取消下划线打印
	SetSuperscript() error         //设置上标打印
	CancelSuperscript() error      //取消上标
	SetSubscript() error           //设置下标
	CancelSubscript() error        //取消下标
	SetSuper_Sub_Script() error    //设置上下标打印方式
	CancelSuper_Sub_Script() error //取消上下标打印

	//倍宽、倍高、水平扩展
	DoublePrintWidth() error           //倍宽
	CancelDoublePrintWidth() error     //取消倍宽
	SetThreeWidthPrint() error         //设置3倍宽
	CancelThreeWidthPrint() error      //取消3倍宽
	DoublePrintHeight() error          //倍高
	CancelDoublePrintHeight() error    //取消倍高
	SetThreeHighPrint() error          //设置3倍高
	CancelThreeHighPrint() error       //取消3倍高
	SetSBC_caseDouble_width() error    //设置全角汉字的双倍水平扩展打印方式
	CancelSBC_caseDouble_width() error //取消全角汉字的双倍水平扩展打印方式
	SetHorizontalCompress() error      //设置水平压缩打印方式(原来字符的1/2)
	CancelHorizontalCompress() error   //取消水平压缩

	/**
	 * 打印头定位
	 */
	SeekRealLinePos(pos float32) error                       //在一行内设置打印头水平方向绝对位置
	SeekLinePos(incrementpos float32, currpos float32) error //在一行内设置打印头水平方向相对位置
	SeekRealColPos(pos float32) error                        //设置打印头竖直方向绝对位置
	SeekColPos(incrementpos float32, currpos float32) error  //设置打印头竖直方向相对位置（直接进纸n/120英寸）
	SetPlaceOfEnter(pos float32) error                       //设置回车位置 1B-25-36-n1-n2
	SetHorizontalTab() error                                 //设置水平tab
	MoveHorizontalTab() error                                //移至水平tab位置

	/**
	 * 打印方式
	 */
	SetBold() error                  //设置加重打印
	CancelBold() error               //取消加重打印
	SetSinglenessOrientation() error //设置单向打印方式
	SetDiprosopyOrientation() error  //设置双向打印方式
	SetRepeatPrint() error           //设置重复打印方式
	CancelRepeatPrint() error        //取消重复打印方式
	High_speedPrint() error          //设置高速打印方式. 打印机对于 HD ANK 字符,汉字间隔打印,因此打印速度大约为 HD 方式的两倍
	High_densityPrint() error        //设置高密打印方式. 打印机以高密方式打印 HD ANK 字符、汉字。
	Low_noisePrint() error           //设置低噪音打印方式

	//各种字符方式打印
	SetPica_HS_ANK() error  //设置 Pica HS ANK 字符方式
	SetPica_HD_ANK() error  //设置 Pica HD ANK 字符方式
	SetElite_HS_ANK() error //设置 Elite HS ANK 字符方式
	SetElite_HD_ANK() error //设置 Elite HD ANK 字符方式
	SetANK() error          //设置 ANK 字符方式
	SetANK_Anonym() error   //设置 ANK 字符片假名方式

	//汉字打印方式
	SetChinese() error                 //设置汉字打印
	CancelChinese() error              //取消汉字
	SetChineseTypeface(font int) error //设置汉字字体。0、1、2、3分别是：宋、仿、楷、繁。
	SetDBC_case() error                //设置半角字方式
	CancelDBC_case() error             //取消半角字方式（即全角）
	SetVerticalChinese() error         //设置汉字竖写方式
	SetHorizontalChinese() error       //设置汉字横向书写方式（即取消汉字竖写）
	SetDBC_caseCompose() error         //设置半角字合成方式。当在汉字竖向书写方式下接收到此码后,打印机将在此码后的两半角字字符合成起来作为合成字符。不持续有效
	ForbidDBC_caseErect() error        //禁止半角字字符竖写方式
	CancelForbidDBC_caseErect() error  //取消禁止半角字字符竖写方式
	OutPaperForward() error            //单页向前出纸
	OutPaperBackward() error           //单页向后出纸

	////////////////////////////////////////////////////////////////////////////////////////////////////////

	//开始打印job
	BeginPrintJob(pin *DriverArg, timeout int) error

	//设置页宽。毫米
	SetPageWidth(pos float32) error

	//打印字节内容
	OutputByte(s []byte) error

	//打印字符串,编码为GBK
	OutputString(s string) error

	//TODO subscript才是下标。。。。 by jinsl 20160119

	//设置上标打印
	SetSuperline() error

	//取消上标打印
	CancelSuperline() error

	//设置下标打印
	SetSubline() error

	//取消下标打印
	CancelSubline() error

	//设置列的间距
	SetColInterval(w float32) error

	//结束打印job
	EndPrinterJob()
}
