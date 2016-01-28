package print

type PrintUnderline struct {
	PrintBasic
}

func NewPrintUnderline() *PrintUnderline {
	return &PrintUnderline{PrintBasic: PrintBasic{}}
}

func (j *PrintUnderline) Decorate(pc *PrintCurState) int64 {
	return PRINT_FORMAT_UNDERLINE
}
