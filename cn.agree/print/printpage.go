package print

//首先是Root节点,然后是Page节点。
//因为Page节点首先是可以确认的。最多后面的Page节点和前面的Page大小不一致

//Page节点
type PrintPage struct {
	PrintBasicExt
}

//Page节点元素
func NewPrintPage() *PrintPage {
	return &PrintPage{PrintBasicExt: PrintBasicExt{}}
}

func (j *PrintPage) CalcExtentValue(pc *PrintCurState) {
}

func (j *PrintPage) HeightChanged(pc *PrintCurState, height int64) {

}

//可从Left,Top这两个参数推出其余的参数
func (j *PrintPage) InitPageConfig(ps *PrintSetting) {
	j.Width = ps.ps.Width
	j.Height = ps.ps.Height
	j.Left = "0"
	j.Top = "0"
	j.Width = ps.ps.Width
	j.Height = ps.ps.Height
	j.Name = "page"
}

//打印的数据结构,包括数据和设置
type PrintFacade struct {
	S *PrintSetting
	V PrintInterface
}
