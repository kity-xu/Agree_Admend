package print

import (
	"cn.agree/utils"
)

type PrintTd struct {
	PrintBasicExt
}

func NewPrintTd() *PrintTd {
	return &PrintTd{PrintBasicExt: PrintBasicExt{}}
}

//计算规则：
//表示一个Table中的CELL,只需移动X坐标即可
//对于下层来说。
func (j *PrintTd) CalcExtentValue(pc *PrintCurState) {
	//移动X坐标,使其等于宽度
	pc.CurX = j.GetRealLeft() + j.GetWidth()
	//重新设置Y值坐标,使其等于进入前的坐标
	pc.CurY = j.GetRealTop()

	//寻找最高高度
	s := findChildMaxValue(j, 2)

	utils.Trace("find td max height [%f]", s)

	j.SetRealHeight(s)
	j.SetRealWidth(j.GetWidth())

}

func (j *PrintTd) HeightChanged(pc *PrintCurState, height int64) {

}
