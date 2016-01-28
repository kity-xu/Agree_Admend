package print

type PrintDoubleHeight struct {
	PrintBasic
}

func NewPrintDoubleHeight() *PrintDoubleHeight {
	return &PrintDoubleHeight{PrintBasic: PrintBasic{}}
}

func (j *PrintDoubleHeight) Decorate(pc *PrintCurState) int64 {
	return PRINT_FORMAT_DOUBLEHEIGHT
}
