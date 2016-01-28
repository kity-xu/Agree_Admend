package print

//换行
type PrintBr struct {
	*PrintBasicExt
}

func NewPrintBr() PrintBr {
	tmp := PrintBr{PrintBasicExt: &PrintBasicExt{}}
	return tmp
}

//所占宽度为一行
//所占长度为整个width
func (j *PrintBr) CalcExtentValue(pc *PrintCurState) {
	hfactor := 1
	if TestPrintInsFlag(&pc.Flags, PRINT_FORMAT_DOUBLEHEIGHT) {
		hfactor = 2
	}
	pc.CurY += pc.LineInterval * float32(hfactor)
	pc.CurX = 0
	j.SetRealHeight(pc.LineInterval * float32(hfactor))
	j.SetRealWidth(j.GetWidth())
	return
}

func (j *PrintBr) HeightChanged(pc *PrintCurState, height int64) {
}