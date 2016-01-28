/**
 *徐晓东 新增于2016-1-14
 *旧的pr2指令封装在各个函数中，我们要调的时候需要查找 查看各个函数的说明， 没有一个统一直观的概念，而且代码冗余太高，不利于使用、新增和维护。
 *针于此，设置了pr2-instruction文件。这是一个常用pr2指令集，各个指令以全局变量的形式存在（并附上详细说明）。
 */

/*********************************** OLIVETTI *******************************/

package olivetti

//××××××××××××××××××   OLIVETTI 仿真方式pr2打印机命令集 ×××××××××××××××××××××//

/**
 * 1. 汉字的相关设置 （中文、拳脚、字体、特殊货币）
 */

var (
	SetChinese         string = "1C26"   //设置中文方式
	CancelChinese      string = "1C2E"   //取消中文方式
	SetSBC_case        string = "1C6B"   //设定全角ASCII
	CancelSBC_case     string = "1C67"   //取消全角ASCII
	SetChineseTypeface string = "1C7430" //最后一位 0 1 2 3 宋 仿宋 楷 繁
	Currency_symbol    string = "9A"     //$    详见pr2编程手册
)

/**
 * 2. 页面控制设置 （页面长度、左边界、页顶、页尾、行距、打印字距）
 *    在打印机域内发送指令或打印介质已存在的情况下，下面命令无效
 */
var (
	PaperSize    string = "1B513132331B5A" //纸张大小。 原命令 1B-51-3x-3x-3x-1B-5A。打印机中有纸 命令不生效
	LeftBoundary string = "1B4A303030"     //左边界。   此处000为打印机默认边界
	PaperTop     string = "1B54303030"     //页顶。    1B-54-3x-3x-3x  以十进制 nnn/340 或 nnn/216 英寸
	PaperEnd     string = "1B4D393939"     //页低。 此处为缺省值 3.1mm
	Line_width   string = "1B263030"       //行据。 nn/240或nn/216英寸，此处为缺省值1/6英寸
	Word_spacing string = "1B3C"           //字距。 此处为10字符/英寸的字距（英文）； 1B3D 12/英寸；1B3E 16.6/英寸；1B-61-3x 可变字距
)

/**
 * 3.打印控制命令（字形、倍宽、倍高、三倍、上下标、上下划线、特殊）
 *
 */
var (
	Font                    string = "1B52303030" //字形（包括草稿、高速草稿、斜体、斜体草稿及各个国家的字符集）。 此处默认为草稿
	SetDouble_width         string = "1B33"       //倍宽
	CancelDouble_width      string = "1B34"
	SetThreeWidth           string = "1C68" //三倍宽
	CancelThreeWidth        string = "1C6A"
	SetDouble_width_high    string = "1B64" //倍宽倍高
	CancelDouble_width_high string = "1B65"
	Lengthways              string = "1B7730" //纵向扩展，最后一位的含义是：0 取消；1 倍高；2 三倍高；
	SetBoldface             string = "1B28"   //黑体
	CancelBoldface          string = "1B29"
	BackgroundPrint         string = "1C2830" //背景打印。最后一位：0取消；1逆打印；2密网打印；3梳网打印
	SetTopUnderLine         string = "1B2A30" //上/下划线打印。 0下划线；3上划线；4上/下划线
	CancelTopUnderLine      string = "1B2B"   //清除上述设置。
	SetSuperSubscript       string = "1B6031" //上/下标打印。 0上标； 1下标
	CancelSuperSubscript    string = "1B7B"   //清楚上述设置
)

/**
 * 4.打印机操作命令（存折、流水、换行、换页、回车、水平制表、反向换行、绝对水平定位、相对垂直定位
 *	 绝对垂直下拉、退纸、改变仿真方式、选字符集、响铃、删除）
 */
var (
	SetBankbook            string = "1B5335"     //设定存折打印机
	SetGlide               string = "1B5331"     //设定流水打印机
	LineDown               string = "0A"         //换行
	NextPage               string = "0C"         //换页(后退纸)
	Enter                  string = "0D"         //回车
	HorizontalTab          string = "09"         //水平制表
	LineUp                 string = "1B37"       //反向换行
	AbsoluteHorizontalSite string = "1B48303030" // 绝对水平定位。 nnn相对于001向右偏移的绝对值
	RelativeVerticalSite   string = "1B49303030" // 相对垂直定位。在不改变水平位置的前提下打印介质前移动"nnn"行。所移动的实际距离取决于所选的行距。
	AbsoluteVerticalSite   string = "1B4C303030" //绝对垂直定位。在不改变水平位置的前提下打印介质向前(或向后)移动到由"nnn"定义的行上,所移动的实际距离取决于所选的行距。
	PaperOut               string = "1B4F"       //退纸（若无纸可退，仍会操作不产生错误）
	ChangeToIBM            string = "1B5E30"     //选择IBM仿真方式
	ChangeToOKI            string = "1B5E31"     //选择OKI仿真方式
	CharacterGather        string = "1B5B303030" //此处为国际字符集
	Diabolo                string = "07"         //响铃
	Cleanup                string = "7F"         //清除打印缓冲区及其中的命令,清除含有ESC系列的代码。
)

/**
 * 5.图形打印方式（图形打印、取消图形打印、相对水平位置、打印针控制）
 */
var (
	SetGraphicPrint        string = ""         //图形打印方式
	CancelGraphicPrint     string = ""         //取消图形打印
	RelativeHorizontalSite string = ""         //相对水平位置
	CortrolPrinter_Spicula string = "1B214730" // 0:8针图形打印模式；1：24针。pr2缺省值为8
)

/**
 * 6.用户自定义字符集下载命令。
 */
var (
// ...
)

/**
 * 7.打印机控制命令
 */
var (
	PrinterInitIdentifyRequest       string = "1B5A"   //打印机初始识别请求。 由主机发出,要求打印机回答它的 ID 标识。
	PrinterIdentify                  string = "1B2F37" //打印机识别。 由打印机发出,作为对命令 ESC Z 的应答
	PrinterConfigureRequest          string = "1B69"   //打印机配置请求。命令由主机发出,要求打印机回答其基本配置。
	PrinterConfigure                 string = ""       //打印机配置 1B-70-xx-xx-xx--xx。 见pr2文档
	PrinterStateRequest              string = "1B6A"   //打印机状态请求。由主机发出,要求打印机传送同步状态码回答主机。
	PrinterSynchronizationState      string = ""       //1B-72-xx. 打印机同步状态
	PrinterPaperStateRequest         string = "1B2042" //打印介质请求状态。 由主机发出,要求打印机回答关于打印介质当前所处位置的状态。
	PrinterPaperSynchronizationState string = ""       //打印介质同步状态 1B-42-35-xx-1B-5A . 打印介质同步状态
	SER_UP_ConfiguerRequest          string = "1B2061" //SET-UP 配置请求。 由主机发出,要求打印机回答它自身的 SET-UP 配置
	CleanErrorState                  string = "1B6C"   //清楚错误状态
	ChooseOlivettiProcess            string = "1B6E"   //选择Olivetti对话进程  控制的数据交换进程(仅在 OLIVETTI 仿真方式时才接受)
	CleanAll                         string = "1B30"   //总清。 对打印机的固件和硬件进行复位
	OperatorRequest                  string = ""       //操作员请求 1B-55-xx. 通过设定参数"xx",确定操作员请求并打开或关闭相应的"STATION"灯
	OperatorResponsion               string = ""       //操作员应答 1B-72-xx。 由打印机发出,回答操作员请求,可以按以下两情况进行:
	PrinterAutomatismControl         string = "1B2E"   //赋予自动操作。设定打印机,使印机在打印介质一旦被插入进纸器后,无须按面板上的任何按键就能自动操作员请求命令给以回
	PrinterManualControl             string = "1B5F"   //赋予手动操作
)

/*********************************** OVER ***********************************************************/
