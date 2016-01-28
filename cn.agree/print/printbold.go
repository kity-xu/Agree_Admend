package print

type PrintBold struct {
	PrintBasic
}

func NewPrintBold() *PrintBold {
	return &PrintBold{PrintBasic: PrintBasic{}}
}

//返回加粗标志
func (j *PrintBold) Decorate(pc *PrintCurState) int64 {
	return PRINT_FORMAT_BOLD
}
