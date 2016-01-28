package print

const (
	PRINT_FORMAT_BOLD         = 1 << iota //黑体
	PRINT_FORMAT_UNDERLINE                //下划线
	PRINT_FORMAT_SUP                      //上标
	PRINT_FORMAT_SUB                      //下标
	PRINT_FORMAT_DOUBLEHEIGHT             //倍高
	PRINT_FORMAT_DOUBLEWIDTH              //倍宽
	PRINT_FORMAT_TRIPLEHEIGHT             //三倍高
	PRINT_FORMAT_TRIPLEWIDTH              //三倍宽
	PRINT_FORMAT_TURNPAGE                 //分页
	PRINT_FORMAT_FONT_LEFT                //靠左对齐
	PRINT_FORMAT_FONT_RIGHT               //靠右对齐
	PRINT_FORMAT_FONT_MIDDLE              //中间对齐
	PRINT_MAX_FLAG
)

//记录打印头的位置,当前的设置等(是否加粗之类)
type PrintCurState struct {
	CurX         float32
	CurY         float32
	CurPage      int
	PageWidth    float32
	PageHeight   float32
	LeftMargin   float32
	LineInterval float32
	ColInterval  float32
	RightMargin  float32
	TopMargin    float32
	BottomMargin float32
	Flags        int64
}

//最后形成的一行的结构
type PrintLine struct {
	Source    string     //表示其来源
	PrX       float32    //要打印的x坐标
	PrY       float32    //要打印的y坐标
	PrContent string     //要打印的内容
	clen      int        //字符串的长度
	Flags     int64      //打印标志
	PrNext    *PrintLine //下一个元素
}

//置位操作
func SetPrintInsFlag(flags *int64, flag int64) {
	*flags = (*flags) | flag
}

//清位操作
func ClearPrintInsFlag(flags *int64, flag int64) {
	*flags = (*flags) & ^flag
}

//测试位操作
func TestPrintInsFlag(flags *int64, flag int64) bool {
	return ((*flags) & flag) == flag
}
