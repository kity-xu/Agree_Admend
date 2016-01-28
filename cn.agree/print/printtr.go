package print

import (
	"cn.agree/utils"
)

type PrintTr struct {
	PrintBasicExt
}

func NewPrintTr() *PrintTr {
	return &PrintTr{PrintBasicExt: PrintBasicExt{}}
}

//计算规则：
//表示一个Table中的一行,移动X坐标为起始位置,移动Y坐标为下一行
//Tr的设置：X坐标设为Table的起始坐标
//          Y坐标移动一个间距
//这里需要注意：很有可能会出现Td一行排列不下,进而往下移的情况。
//长度为最高的Td的长度。如果有需要,调整td的长度
func (j *PrintTr) CalcExtentValue(pc *PrintCurState) {
	//寻找最高高度
	s := findChildMaxValue(j, 2)
	//寻找累计宽度
	t := findChildCalValue(j, 1)

	utils.Trace("find tr max height [%f]", s)

	par := findparentPosNodeWithName(j, "Table")
	if par == nil {
		utils.Error("can't find Tr's parent node,maybe error")
		return
	}

	table, found2 := par.(*PrintTable)
	if !found2 {
		utils.Error("table element can't convert to PrintTable,maybe error")
		return
	}
	pc.CurX = par.GetRealLeft()
	j.SetRealHeight(s)
	j.SetRealWidth(t)

	//加上表格宽度
	pc.CurY += (s + table.LineSeperate)

}

func (j *PrintTr) HeightChanged(pc *PrintCurState, height int64) {

}
