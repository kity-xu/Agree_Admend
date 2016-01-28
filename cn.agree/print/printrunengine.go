package print

import (
	"cn.agree/driverlayer"
	"cn.agree/utils"
	"errors"
	"strconv"
	"strings"
)

var UNSUPPORTED_MEASURE_FORMAT = errors.New("unsupported measure unit")

func RunPr2Engine(s string, pr2 driverlayer.IPr2Print, pin *driverlayer.DriverArg) error {
	var pc PrintCurState
	pf, e := ParserPrintXml(s)
	if e != nil {
		utils.Error("parse xml file error %s", e.Error())
	} else {
		utils.Info("parse xml file success")
	}
	ps := pf.S
	pi := pf.V
	pc.BottomMargin = ps.ps.Bottommargin
	pc.LeftMargin = ps.ps.Leftmargin
	pc.PageHeight = ps.ps.Height
	pc.PageWidth = ps.ps.Width
	pc.TopMargin = ps.ps.Topmargin
	pc.RightMargin = ps.ps.Rightmargin
	pc.CurX = 0
	pc.CurY = 0
	pc.CurPage = 1
	pc.Flags = 0
	pc.LineInterval = ps.ps.LineInterval
	pc.ColInterval = ps.ps.ColInterval

	RecurCalcExtent(pi, &pc)
	utils.Info("RecurCalcExtent success")
	fl := GeneratePrintLine(pi)
	utils.Debug("begin print printline----------------------------")
	PrintPrintLine(fl)
	utils.Debug("end print printline----------------------------")
	pr2.BeginPrintJob(pin, 20)
	defer pr2.EndPrinterJob()
	RunRealPrinter(ps, fl, pr2)
	return nil
}

//检查是否超过了页数
func detectExceedPage() {

}

//进行坐标的转换
//如果是绝对坐标：则返回绝对坐标值
//如果是相对坐标：则返回相对于原值后的值
func convertaxistype(raw float32, newexp string, postype string) (float32, error) {
	//utils.Trace("received convertaxistype,raw :[%f],exp [%s],postype [%s]", raw, newexp, postype)
	if strings.EqualFold(postype, "") {
		return raw, nil
	}
	var sign = true
	if (newexp[0] == '+') || (newexp[0] == '-') {
		if strings.EqualFold(postype, "absolute") {
			return 0, UNSUPPORTED_MEASURE_FORMAT
		}
		if newexp[0] == '-' {
			sign = false
		}
		newexp = newexp[1 : len(newexp)-1]
	}

	s, e := strconv.ParseFloat(newexp, 32)
	if e != nil {
		if strings.HasSuffix(newexp, "mm") {
			newexp = newexp[0 : len(newexp)-2]
			s, e = strconv.ParseFloat(newexp, 32)
		} else {
			utils.Error("unsupported measure unit %s", newexp)
			return raw, UNSUPPORTED_MEASURE_FORMAT
		}
	}

	//绝对定位
	switch postype {
	case "absolute":
		return float32(s), nil
	case "relative":
		if sign {
			return (raw + float32(s)), nil
		} else {
			return (raw - float32(s)), nil
		}
	default:
		utils.Error("unsupported pos type %s", postype)
		return raw, UNSUPPORTED_MEASURE_FORMAT
	}

}

//生成PrintLine
func GeneratePrintLine(cur PrintInterface) *PrintLine {
	var s *PrintLine
	var c *PrintLine
	child := cur.GetChild()
	sibling := cur.GetNextSibling()
	//如果有子节点的,则首先计算子节点
	if child != nil {
		c = GeneratePrintLine(child)
	}

	//子节点计算完毕后,计算自身
	switch value := cur.(type) {
	case PrintControlInterface:
		break

	case PrintPosInterface:
		z := cur.(PrintPosInterface)
		s = z.GetPrintLine()
		break
		//需要进行回写处理
	default:
		utils.Error("find not correct kind element %+v", value)
		return nil
	}
	//self child rewrite by jinsl
	s = AdjustPrintLine(s, c)

	if sibling != nil { //计算相邻节点,并且获得值
		s = AdjustPrintLine(s, GeneratePrintLine(sibling))
	}

	return s
}

//递归的计算属性,所有的坐标一律以绝对坐标返回,方便处理
//然后在打印的时候进行偏移处理
func RecurCalcExtent(cur PrintInterface, pc *PrintCurState) {
	//首先计算自身的属性
	//因为打印是纵向扩展的,因此需要把纵向设大

	var eff int64 = -1
	curt, ok := cur.(PrintPosInterface)
	//utils.Debug("ok is [%s]", ok)
	if ok {
		if curt.GetHeight() == 0 {
			curt.SetHeight(pc.BottomMargin - pc.CurY)
		}

		pos := curt.GetPos()
		//记录转换后的值
		//if strings.EqualFold(pos, "absolute") || strings.EqualFold(pos, "relative") {
		pc.CurX, _ = convertaxistype(pc.CurX, curt.GetLeft(), pos)
		curt.SetRealLeft(pc.CurX)
		pc.CurY, _ = convertaxistype(pc.CurY, curt.GetTop(), pos)
		curt.SetRealTop(pc.CurY)
		al := curt.GetAlign()

		//获取特殊效果
		switch al {
		case "Middle":
			eff = PRINT_FORMAT_FONT_MIDDLE
		case "Left":
			eff = PRINT_FORMAT_FONT_LEFT
		case "Right":
			eff = PRINT_FORMAT_FONT_RIGHT

		}
		if eff != -1 && TestPrintInsFlag(&pc.Flags, eff) == false {
			utils.Debug("detect PrintPosElement,name [%s],set effect %d", cur.GetName(), eff)
			SetPrintInsFlag(&pc.Flags, eff)
		}

	}

	//装饰的效果需要在一开始就计算
	curc, ok2 := cur.(PrintControlInterface)
	//utils.Debug("ok2 is [%s]", ok2)
	if ok2 {
		eff = curc.Decorate(pc)
		//如果已经在效果中了,则不需要设置单位了
		if TestPrintInsFlag(&pc.Flags, eff) == false {
			utils.Debug("detect not decorated element,name [%s],set effect", cur.GetName())
			SetPrintInsFlag(&pc.Flags, eff)
		} else {
			utils.Debug("detect decorated element,name [%s],skip", cur.GetName())
			eff = -1
		}
	}

	child := cur.GetChild()
	utils.Debug("child is [%s]", child)
	sibling := cur.GetNextSibling()
	utils.Debug("sibling is [%s]", sibling)
	//如果有子节点的,则首先计算子节点
	if child != nil {
		RecurCalcExtent(child, pc)
		//然后计算自己
	}

	//子节点计算完毕后,计算自身
	switch value := cur.(type) {
	case PrintControlInterface:
		break
		//需要进行回写处理
	case PrintPosInterface:
		z := cur.(PrintPosInterface)
		z.CalcExtentValue(pc)
		break
		//需要进行回写处理
	default:
		utils.Error("find not correct kind element,name[%s],value [%+v]", cur.GetName(), value)
	}

	//清除效果
	if eff != -1 {
		utils.Debug("find clear flag instruction,name [%s],remove effect", cur.GetName())
		ClearPrintInsFlag(&pc.Flags, eff)
	}

	if sibling != nil { //计算相邻节点,并且获得值
		RecurCalcExtent(sibling, pc)
	}

}

//执行实际的打印操作,至此,PrintLine都是排好序的
func RunRealPrinter(ps *PrintSetting, pi *PrintLine, pr2 driverlayer.IPr2Print) {
	var pl *PrintLine
	pl = pi
	var (
		X        float32
		Y        float32
		ftmp     int64
		curflag  int64
		curbflag int64
		remain   int64
	)
	//首先是各种预设参数的打印
	pr2.Init()
	pr2.SetLineInterval(ps.ps.LineInterval)
	pr2.SetColInterval(ps.ps.ColInterval)
	pr2.SetLeftMargin(ps.ps.Leftmargin)
	pr2.SetRightMargin(ps.ps.Width - ps.ps.Rightmargin)
	pr2.SetTop(ps.ps.Topmargin)
	//	pr2.SetPageBottom(ps.ps.Height, ps.ps.Bottommargin) //
	pr2.AdvanceOneLine() //pr2需要先进纸至少一行

	//然后是打印文字
	for {
		if pl == nil {
			utils.Trace("end of PrintLine,print success")
			//清除所有标志位
			remain = PRINT_FORMAT_BOLD
			for {
				if remain >= PRINT_MAX_FLAG {
					break
				}
				curbflag = curflag & remain
				//标志改变了
				if curbflag >= 1 {
					switch curbflag {
					case PRINT_FORMAT_BOLD:
						pr2.CancelBold()
						break

					case PRINT_FORMAT_UNDERLINE:
						pr2.CancelUnderline()
						break

					case PRINT_FORMAT_SUP:
						pr2.CancelSuperline()
						break

					case PRINT_FORMAT_SUB:
						pr2.CancelSubline()
						break

					case PRINT_FORMAT_DOUBLEHEIGHT:
						pr2.CancelDoublePrintHeight()
						break

					case PRINT_FORMAT_DOUBLEWIDTH:
						pr2.CancelDoublePrintWidth()
						break

					case PRINT_FORMAT_TRIPLEHEIGHT:
					case PRINT_FORMAT_TRIPLEWIDTH:
					default:
					}
				}
				remain = remain << 1
			}
			//进行退纸操作
			pr2.EjectPaper()
			return
		}

		if Y != pl.PrY {
			//推进列
			pr2.SeekColPos(pl.PrY-Y, Y)
			pr2.Carriage()
			X = 0
			Y = pl.PrY
		}
		//进行打印
		if X != pl.PrX {
			//推进行
			pr2.SeekLinePos(pl.PrX-X, X)
			X = pl.PrX
		}

		//处理特效的事情,对当前标志和所需要的标志进行异或处理
		ftmp = curflag ^ pl.Flags
		utils.Debug("get current flag [%d]", ftmp)
		remain = PRINT_FORMAT_BOLD
		for {
			if remain >= PRINT_MAX_FLAG {
				break
			}
			curbflag = ftmp & remain
			//标志改变了
			if curbflag >= 1 {
				if (pl.Flags & curbflag) >= 1 { //表明是置位操作
					switch curbflag {

					case PRINT_FORMAT_BOLD:
						pr2.SetBold()
						break

					case PRINT_FORMAT_UNDERLINE:
						pr2.SetUnderline()
						break

					case PRINT_FORMAT_SUP:
						pr2.SetSuperline()
						break

					case PRINT_FORMAT_SUB:
						pr2.SetSubline()
						break

					case PRINT_FORMAT_DOUBLEHEIGHT:
						pr2.DoublePrintHeight()
						break

					case PRINT_FORMAT_DOUBLEWIDTH:
						pr2.DoublePrintWidth()
						break

					case PRINT_FORMAT_TRIPLEHEIGHT:
						break
					case PRINT_FORMAT_TRIPLEWIDTH:
						break
					case PRINT_FORMAT_TURNPAGE:
						pr2.EjectPaper()
					default:
					}
					SetPrintInsFlag(&curflag, curbflag)
				} else { //清除位操作
					switch curbflag {
					case PRINT_FORMAT_BOLD:
						pr2.CancelBold()
						break

					case PRINT_FORMAT_UNDERLINE:
						pr2.CancelUnderline()
						break

					case PRINT_FORMAT_SUP:
						pr2.CancelSuperline()
						break

					case PRINT_FORMAT_SUB:
						pr2.CancelSubline()
						break

					case PRINT_FORMAT_DOUBLEHEIGHT:
						pr2.CancelDoublePrintHeight()
						break

					case PRINT_FORMAT_DOUBLEWIDTH:
						pr2.CancelDoublePrintWidth()
						break

					case PRINT_FORMAT_TRIPLEHEIGHT:
					case PRINT_FORMAT_TRIPLEWIDTH:
					case PRINT_FORMAT_TURNPAGE:
					default:
					}
					ClearPrintInsFlag(&curflag, curbflag)
				}
			}

			remain = remain << 1
		}

		if pl.PrContent != "" {
			pr2.OutputString(pl.PrContent)
			X += float32(pl.clen) * ps.ps.ColInterval
			utils.Trace("After print text,X change to %f", X)
			//调整水平距离
		}
		pl = pl.PrNext
	}

}

//调整PrintLine,形成布局
//主要排序规则:
//1:行小的排在前面.prev:已经排好序的序列
//2:如果有相同行的,列小的牌前面。cur:已经排好序的序列
//类似于冒泡排序
//self prev，child cur
//prev是当前节点，只有一个对象，不是链表，cur是已经排好序的所有子节点，是链表，所以将prev插入到cur里面
//
//当前节点                     *
//子节点序列                ************************
//找位置插入
//rewrite by jinsl
func AdjustPrintLine(prev *PrintLine, cur *PrintLine) *PrintLine {
	if prev == nil {
		//只计算子节点时 jinsl
		//可以return nil jinsl
		return cur
	}
	if cur == nil {
		//只计算自身时 jinsl
		return prev
	}
	var (
		p      *PrintLine
		header *PrintLine
		prevh  *PrintLine
		q      *PrintLine
	)
	p = prev

	q = cur
	if p.PrY > q.PrY {
		header = q
	} else {
		if p.PrY == q.PrY && p.PrX > q.PrX {
			header = q
		} else { //prev在cur的最前面
			header = p
			header.PrNext = q
			return header
		}
	}
	prevh = header //只能是q
	for {
		//第一个链表已经完成,把剩下的q连接起来，即prev在cur中间
		//		if p == nil && q != nil {
		//			prevh.PrNext.PrNext = q
		//			return header
		//		}
		//第二个链表已经完成,把剩下的p连接起来，即prev在cur最后面
		if q == nil && p != nil {
			prevh.PrNext = p
			return header
		}
		//避免异常
		if q == nil && p == nil {
			return header
		}
		//正确,继续往下比较
		if p.PrY < q.PrY {
			prevh.PrNext = p
			prevh.PrNext.PrNext = q
			return header
		} else {
			if p.PrY == q.PrY { //在同一行
				//正确,继续比较
				if p.PrX > q.PrX {
					prevh = q
					q = q.PrNext
				} else {
					prevh.PrNext = p
					prevh.PrNext.PrNext = q
					return header
				}
			} else { //大于,调整顺序,移动后一个节点到前面
				prevh = q
				q = q.PrNext
			}
		}
	}
}

//打印出PrintLine
func PrintPrintLine(prev *PrintLine) {
	for {
		if prev == nil {
			return
		}
		utils.Debug("PrintLine struct:%+v\n", *prev)
		prev = prev.PrNext
	}
}
