package print

type PrintSubline struct {
	PrintBasic
}

func NewPrintSubline() *PrintSubline {
	return &PrintSubline{PrintBasic: PrintBasic{}}
}

func (j *PrintSubline) Decorate(pc *PrintCurState) int64 {
	return PRINT_FORMAT_SUB
}
