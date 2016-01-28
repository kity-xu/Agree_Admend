/*************************************************************************
   > File Name: oki.go
   > Author: xuxiaodong
   > Mail: cherish_xulang@163.com
   > Created Time: 2016年01月19日 星期二 15时44分19秒
*************************************************************************/

package oki

//××××××××××××××××××   OKI 仿真方式pr2打印机命令集 ×××××××××××××××××××××//

/**
 * 1. 打印机硬件控制码（设备控制1、3；清楚缓存；确定打印操作结束；设置SHIFT JIS方式；取消SHIFT JIS方式）
 */
var (
	DeviceControl1  string = "11"   //设备控制1. 本码将打印机设置为 SELECT(联机)状态。当接收到本码或在 DESELECT(脱机)状态下按下“LOCAL”键,打印机被设置为 SELECT(联机)状态。
	DeviceControl3  string = "13"   //设备控制3. 本码取消打印机的 SELECT 状态,并设置为 DESELECT(脱机)状态,当接收到本码或在 SELECT状态下按下 LOCAL 键,打印机被设置为 DESELECT(脱机)状态
	CleanPrinterRom string = "18"   //清除打印机缓冲。 数据被清除,水平扩展方式也被清除(仅 ANK 方式),所有其它方式有保持不变,已收到的确定打印开始的数据不清除
	SetShift_jis    string = "1B6B" //设置SHIFT JIS方式。在这种方式下,接收到的所有第八位为 1 的数据视为汉字码,本方式可由 ESC l 或 ESC $ @取消 。在本方式下,汉字字符宽度为 26,半角字宽度为 13。
	CancelShift_jis string = "1B6C" //取消SHIFT JIS方式。
)

/**
 * 2.垂直方向控制码（换行、设置行据、直接进纸、页长、页顶、换页）
 */
var (
	NextLine      string = "0A"         //换行
	Rowledge_6LPI string = "1B36"       //设置1/6行据
	Rowledge_8LPI string = "1B38"       //设置1/8行据
	Rowledge_some string = "1B25390001" //设置n/120英寸行据。1B-25-39-n1-n2（n1为00,n2表示：00-FF/120）
	PaperIn_some  string = "1B0B01"     //直接进纸若干行。1B-0B-n1-n2（行数=n1×10+n2）
	PaperIn_n120  string = "1B253501"   //直接进纸n/120英寸。1B-25-35-n（01 =< n <= FF）
	PaperLength   string = "1B4601"     //设置页长度。1B-46-n1-n2(起始行 = n1 × 10 + n2)
	PaperTop      string = "1B35"       //设置页顶
	NextPaper     string = "0C"         //换页
)

/**
 * 2.水平方向控制码（回车、打印头左移、退格、设置回车位置、打印头右移、设置左界、设置右界、设置水平tab、移至水平tab）
 */
var (
	Enter               string = "0D"         //回车
	Left_movePrintHead  string = "1B25340001" //打印头左移。1B-25-34-n1-n2
	Backspace           string = "08"         //退格
	SetPlaceOfEnter     string = "1B25360001" //设置回车位置。1B-25-36-n1-n2
	Right_movePrintHead string = "1B25330001" //打印头右移。1B-25-33-n1-n2
	LeftBoundary        string = "1B280001"   //设置左界
	RightBoundary       string = "1B29010E"   //设置右界
	SetHorizontalTab    string = "1B4C"       //设置水平tab
	MoveHorizontalTab   string = "09"         //移至水平tab
)

/**
 * 3.选择字符特性控制码（Pica HS ANK、Pica HD ANK、Elite HS ANK、Elite HD ANK、ANK字符平假名ANK方式、ANK字符平假名方式、图形数据传输、水平双倍扩展图形数据传输）
 */

var (
	Pica_HS_ANK                         string = "1B4E"       //设置 Pica HS ANK 字符方式
	Pica_HD_ANK                         string = "1B48"       //设置 Pica HD ANK 字符方式
	Elite_HS_ANK                        string = "1B42"       //设置 Elite HS ANK 字符方式
	Elite_HD_ANK                        string = "1B45"       //设置 Elite HD ANK 字符方式
	ANK                                 string = "1B26"       //设置 ANK 字符方式
	ANK_Anonym                          string = "1B27"       //设置 ANK 字符片假名方式
	FigureDataTransfers                 string = "1B25310001" //图形数据传输。本码定义图形打印数据,参数 n1 n2 定义了本码之后被传输的图形点列数目[1 列=3 字节(24点)]。
	HorizontalDoubleFigureDataTransfers string = "1B25320001" //水平双倍扩展图形数据传输
	High_speedPrint                     string = "1B44"       //设置高速打印方式. 打印机对于 HD ANK 字符,汉字间隔打印,因此打印速度大约为 HD 方式的两倍
	High_densityPrint                   string = "1B49"       //设置高密打印方式. 打印机以高密方式打印 HD ANK 字符、汉字。
	Low_noisePrint                      string = "1B4F"       //设置低噪音打印方式
	SetUnderLine                        string = "1B58"       //设置下划线打印方式
	CancelUnderLine                     string = "1B59"       //取消下划线打印方式
	SetDouble_width                     string = "1B55"       //设置字符的双倍水平扩展打印方式
	CancelDouble_width                  string = "1B52"       //取消字符的双倍水平扩展打印方式
	SetSBC_caseDouble_width             string = "1C70"       //设置全角汉字的双倍水平扩展打印方式
	CancelSBC_caseDouble_width          string = "1C71"       //取消全角汉字的双倍水平扩展打印方式
	SetHorizontalCompress               string = "1B3C"       //设置水平压缩打印方式(原来字符的1/2)
	CancelHorizontalCompress            string = "1B3E"       //取消水平压缩打印方式
	SetDouble_high                      string = "1B5B"       //设置垂直扩展打印方式
	CancelDouble_high                   string = "1B5D"       //取消垂直扩展打印方式
	SetThreeHigh                        string = "1B65"       //设置 3 倍高度打印方式
	CancelThreeHigh                     string = "1B66"       //取消 3 倍高度打印方式
	SetThreeWidth                       string = "1B67"       //设置 3 倍宽度打印方式
	CancelThreeWidth                    string = "1B68"       //取消 3 倍宽度打印方式
	SetBlod                             string = "1B69"       //设置加重打印方式
	CancelBlod                          string = "1B6A"       //取消加重打印方式
	SetSinglenessOrientation            string = "1B2555"     //设置单向打印方式
	SetDiprosopyOrientation             string = "1B2542"     //设置双向打印方式
	SetRepeatPrint                      string = "1B6D"       //设置重复打印方式
	CancelRepeatPrint                   string = "1B6E"       //取消重复打印方式
)

/**
 * 4.汉字特性码（方式、字体、上下标、半角、外部传输字模、字符间距、汉字竖写方式、汉字横写方式、半角字合成、禁止半角字符竖写、取消禁止）
 */
var (
	SetChinese                string = "1B2440" //设置汉字方式。 本码设置打印机为汉字方式,如果已设置了半角字方式,半角字方式有效
	CancelChinese             string = "1B2848" //取消汉字方式。不取消半角方式
	SetChineseTypeface        string = "1C7630" //设置汉字字体。0、1、2、3分别是：宋、仿、楷、繁。
	SetSuperscript            string = "1C4E"   //设置上标打印方式
	CancelSuperscript         string = "1C4F"   //取消上标打印方式
	SetSubscript              string = "1C50"   //设置下标打印方式
	CancelSubscript           string = "1C51"   //取消下标打印方式
	SetSuper_Sub_Script       string = "1C52"   //设置上下标打印方式
	CancelSuper_Sub_Script    string = "1C53"   //取消上下标打印方式
	SetDBC_case               string = "1C72"   //设置半角字方式
	CancelDBC_case            string = "1C73"   //取消半角字方式
	SetVerticalChinese        string = "1C4A"   //设置汉字竖写方式
	SetHorizontalChinese      string = "1C4B"   //设置汉字横向书写方式（即取消汉字竖写）
	SetDBC_caseCompose        string = "1C5F"   //设置半角字合成方式。当在汉字竖向书写方式下接收到此码后,打印机将在此码后的两半角字字符合成起来作为合成字符。不持续有效
	ForbidDBC_caseErect       string = "1C74"   //禁止半角字字符竖写方式
	CancelForbidDBC_caseErect string = "1C75"   //取消禁止半角字竖写方式
)

/**
 * 5.OKI 5530SC 新增加的控制命令
 */
var (
	OutPaperDirection string = "1B5430" //设置单页纸出纸方向(30:向前出纸；31：向后出纸)
	InitPrintrer      string = "1042"   //打印机初始化
)

/**
 * 6.改变仿真方式命令
 * 在下列情况之一,命令的发送才会有效:
 *.打印机处于开机状态
 *.打印缓冲区为空区
 *.从行清零以后
 *.逻辑硬件复位之后
 *若打印缓冲区不是处于空亲状态,那么接收到此命令后将会先打印出缓冲区的全部字符,此命令才会生效。
 */
var (
	ChanageToOlivetti string = "1B5E31" //改变仿真方式到 OLIVETTI 仿真方式
	ChanageToIBM      string = "1B5E30" //改变仿真方式到 IBM 仿真方式
)
