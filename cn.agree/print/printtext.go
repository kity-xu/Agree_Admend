package print

import (
	"cn.agree/utils"
)

//打印的原子元素
type PrintText struct {
	PrintBasicExt
	Con string
}

func NewPrintText(w, h int) *PrintText {
	return &PrintText{PrintBasicExt: PrintBasicExt{}}
}

//计算所占长度和宽度
//计算是否是ascii码,如果是,则记作1
//如果打印出来的是一个矩形,则最后的改变的PrintCurState的值为:
//Y:指打印的最后一行的位置(并不换行) X:指的是打印的矩形的最右边
//如：打印一行1111
//如果传入的pc.CurX = 10,pc.CurY =1.则在计算后.pc.Curx=10+(1111长度),pc.CurY=1
//还有一种情况,就是特效的影响。如：倍高和倍宽。
//在这种情况下:每个字符占平时的一倍高和宽。需要在这里面一并计算
func (j *PrintText) CalcExtentValue(pc *PrintCurState) {
	//首先找到父pos节点
	var (
		maxx    float32
		count   int
		pname   string
		hfactor int //是否有高度加成
		wfactor int //是否有宽度加成
	)
	hfactor = 1
	wfactor = 1
	par := findparentPosNode(j)
	if par == nil {
		utils.Error("can't find text's parent node,text is %s", j.Con)
		return
	}
	w := par.GetWidth()
	pname = par.GetName()
	//倍高倍宽
	if TestPrintInsFlag(&pc.Flags, PRINT_FORMAT_DOUBLEHEIGHT) {
		hfactor = 2
	}
	if TestPrintInsFlag(&pc.Flags, PRINT_FORMAT_DOUBLEWIDTH) {
		wfactor = 2
	}

	//一行能容纳的字符数。为了统一处理,一律用作双数
	cpline := getAsciiLengthWithMM(float32(w), pc.ColInterval)
	if cpline%2 == 1 {
		cpline = cpline - 1
	}
	utils.Debug("colinterval is [%f],parent width is [%f],calculated ascii len is [%d]", pc.ColInterval, w, cpline)

	tmplength := getAsciiLength(j.Con)
	tmplength = tmplength * wfactor //需要计算倍高倍宽的影响
	utils.Debug("text is [%s], ascii len [%d]", j.Con, tmplength)

	pl := &PrintLine{PrX: pc.CurX, PrY: pc.CurY, Source: pname}
	var startIndex = 0
	pl.PrNext = nil
	var curpl *PrintLine
	curpl = pl
	var findex int

	//框子的长度大于字符串的长度
	//如果是需要下移的,则需要计算长度
	for {
		if cpline >= tmplength {
			//首先需要测试下字体效果
			if TestPrintInsFlag(&pc.Flags, PRINT_FORMAT_FONT_LEFT) {
				wfactor = 2
			}
			//中间对齐
			if TestPrintInsFlag(&pc.Flags, PRINT_FORMAT_FONT_MIDDLE) {
				curpl.PrX += float32((cpline - tmplength) / 2)
			}
			//右对齐
			if TestPrintInsFlag(&pc.Flags, PRINT_FORMAT_FONT_RIGHT) {
				curpl.PrX += float32(cpline - tmplength)
			}

			curpl.PrContent = j.Con[startIndex:]
			curpl.clen = tmplength
			curpl.PrNext = nil
			if maxx == 0 {
				maxx = float32(tmplength) * pc.ColInterval
			}
			//移动列坐标

			pc.CurX += maxx
			pl.Flags = pc.Flags //设置标志位

			//设置实际的高和宽
			j.SetRealWidth(maxx)
			j.SetRealHeight(float32(count+1) * pc.LineInterval)
			j.pl = pl
			return
		} else {
			tmplength -= cpline
			if maxx == 0 {
				maxx = float32(cpline) * pc.ColInterval
			}
			findex = getStringRuneLen(string(j.Con[startIndex:]), cpline/wfactor)
			if findex == -1 {
				curpl.PrContent = j.Con[startIndex:]
			} else {
				curpl.PrContent = j.Con[startIndex : findex*(count+1)]
			}

			//正好分配完成,直接返回即可
			if tmplength == 0 {
				pc.CurX += maxx

				//设置实际的高和宽
				j.SetRealWidth(maxx)
				j.SetRealHeight(float32(count+1) * pc.LineInterval)
				pl.Flags = pc.Flags
				j.pl = pl
			}
			//这步不太合理,应该是先校验,然后再进行计算
			plt := &PrintLine{PrX: pc.CurX, PrY: pc.CurY, Source: pname}
			plt.Flags = pc.Flags
			curpl.PrNext = plt
			curpl.clen = cpline
			curpl = plt
			//移动列坐标

			curpl.PrY += (pc.LineInterval * float32(hfactor))
			pc.CurY += (pc.LineInterval * float32(hfactor))
			count++

			startIndex += findex
		}
	}

}

func (j *PrintText) HeightChanged(pc *PrintCurState, height int64) {

}
