package print

//表格
type PrintTable struct {
	PrintBasicExt
	LineSeperate float32 //行之间的宽度
}

func NewPrintTable() *PrintTable {
	return &PrintTable{PrintBasicExt: PrintBasicExt{}}
}

//计算规则：
//表示一整个Table
//X坐标变为0.Y向下移动一个距离
//因为Tr已经移动过一个距离,这边就不处理了
func (j *PrintTable) CalcExtentValue(pc *PrintCurState) {
}

func (j *PrintTable) HeightChanged(pc *PrintCurState, height int64) {

}
