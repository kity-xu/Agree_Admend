package print

import ()

//Root节点
type PrintRoot struct {
	PrintBasicExt
}

//Root节点元素
func NewPrintRoot() *PrintRoot {
	return &PrintRoot{PrintBasicExt: PrintBasicExt{}}
}

func (j *PrintRoot) CalcExtentValue(pc *PrintCurState) {
}

func (j *PrintRoot) HeightChanged(pc *PrintCurState, height int64) {

}
