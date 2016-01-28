package print

type PrintDoubleWidth struct {
	PrintBasic
}

func NewPrintDoubleWidth() *PrintDoubleWidth {
	return &PrintDoubleWidth{PrintBasic: PrintBasic{}}
}

func (j *PrintDoubleWidth) Decorate(pc *PrintCurState) int64 {
	return PRINT_FORMAT_DOUBLEWIDTH
}
