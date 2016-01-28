package print

type PrintSupline struct {
	PrintBasic
}

func NewPrintSupline() *PrintSupline {
	return &PrintSupline{PrintBasic: PrintBasic{}}
}

func (j *PrintSupline) Decorate(pc *PrintCurState) int64 {
	return PRINT_FORMAT_SUP
}
