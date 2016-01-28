package print

import (
	"cn.agree/utils"
)

//换行
type PrintRect struct {
	PrintBasicExt
}

func NewPrintRect() *PrintRect {
	return &PrintRect{PrintBasicExt: PrintBasicExt{}}
}

//计算规则：
//进行换行处理
func (j *PrintRect) CalcExtentValue(pc *PrintCurState) {
	pc.CurY += pc.LineInterval
	pc.CurX = 0
	//寻找最高高度
	s := findChildMaxValue(j, 2)

	utils.Trace("find rect max height [%f]", s)

	j.SetRealHeight(s)
	j.SetRealWidth(j.GetWidth())
}

func (j *PrintRect) HeightChanged(pc *PrintCurState, height int64) {

}
