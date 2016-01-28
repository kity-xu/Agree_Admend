package print

type PrintTurnPage struct {
	PrintBasic
}

func NewPrintTurnPage() *PrintTurnPage {
	return &PrintTurnPage{PrintBasic: PrintBasic{}}
}

func (j *PrintTurnPage) Decorate(pc *PrintCurState) int64 {
	return PRINT_FORMAT_TURNPAGE
}
