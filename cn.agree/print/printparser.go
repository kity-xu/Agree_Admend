package print

import (
	"cn.agree/utils"
	"errors"
	"github.com/mathume/dom"
	"strconv"
	"strings"
)

var NOT_FOUND_VALUE = errors.New("text value is not integer")
var ELE_NOT_EXISTS = errors.New("node does not exist")

//解析设置
func parseSetting(cur dom.Element) *PrintSetting {
	var i int
	var n dom.Node
	var t dom.Element
	var ps = PrintSetting{}
	f := cur.ChildNodes()
	for i = 0; i < len(f); i++ {
		n = f[i]
		switch n.Kind() {
		case dom.ElementKind:
			t = n.(dom.Element)
			if strings.EqualFold("Page", t.Name()) {
				ps.ps = parsePageSetting(t)
			}
		}
	}
	return &ps
}

//解析页面设置
func parsePageSetting(cur dom.Element) *PageSetting {
	var (
		i   int
		tmp string
		v   float64
		n   dom.Node
		t   dom.Element
		err error
	)

	var ps = PageSetting{}

	f := cur.ChildNodes()
	for i = 0; i < len(f); i++ {
		n = f[i]
		switch n.Kind() {
		case dom.ElementKind:
			t = n.(dom.Element)
			tmp, err = findTextNode(t)
			if err != nil {
				utils.Error("find element %s not have text value ", t.Name())
				continue
			}
			switch t.Name() {
			case "Width":
				v, err = strconv.ParseFloat(tmp, 32)
				if err != nil {
					utils.Error("not supported width format %s", tmp)
					return nil
				}
				ps.Width = float32(v)
				break
			case "Height":
				v, err = strconv.ParseFloat(tmp, 32)
				if err != nil {
					utils.Error("not supported Height format %s", tmp)
					return nil
				}
				ps.Height = float32(v)
				break
			case "Leftmargin":
				v, err = strconv.ParseFloat(tmp, 32)
				if err != nil {
					utils.Error("not supported Leftmargin format %s", tmp)
					return nil
				}
				ps.Leftmargin = float32(v)
				break
			case "Rightmargin":
				v, err = strconv.ParseFloat(tmp, 32)
				if err != nil {
					utils.Error("not supported Rightmargin format %s", tmp)
					return nil
				}
				ps.Rightmargin = float32(v)
				break
			case "Topmargin":
				v, err = strconv.ParseFloat(tmp, 32)
				if err != nil {
					utils.Error("not supported Topmargin format %s", tmp)
					return nil
				}
				ps.Topmargin = float32(v)
				break
			case "Bottommargin":
				v, err = strconv.ParseFloat(tmp, 32)
				if err != nil {
					utils.Error("not supported Bottommargin format %s", tmp)
					return nil
				}
				ps.Bottommargin = float32(v)
				break
			case "LineInterval":
				v, err = strconv.ParseFloat(tmp, 32)
				if err != nil {
					utils.Error("not supported LineInterval format %s", tmp)
					return nil
				}
				ps.LineInterval = float32(v)
				break
			case "ColInterval":
				v, err = strconv.ParseFloat(tmp, 32)
				if err != nil {
					utils.Error("not supported ColInterval format %s", tmp)
					return nil
				}
				ps.ColInterval = float32(v)
				break
			}
			break
		}
	}
	utils.Trace("parse setting result %+v", ps)
	return &ps
}

//解析br元素,不会嵌套其它元素
func parseBrElement(cur dom.Element, p *PageSetting) *PrintBr {
	z := NewPrintBr()
	z.SetName("Br")
	z.Width = p.Width
	return &z
}

//解析Bold元素,会嵌套其它元素
//解析完后会进行计算
func parseBoldElement(cur dom.Element, p *PageSetting) *PrintBold {
	z := NewPrintBold()
	z.SetName("Bold")
	parsePrintContent(cur, z, p)
	return z
}

//解析Sup元素,会嵌套其它元素
//解析完后会进行计算
func parseSupElement(cur dom.Element, p *PageSetting) *PrintSupline {
	z := NewPrintSupline()
	z.SetName("Sup")
	parsePrintContent(cur, z, p)
	return z
}

//解析元素,会嵌套其它元素
//解析完后会进行计算
func parseSubElement(cur dom.Element, p *PageSetting) *PrintSubline {
	z := NewPrintSubline()
	z.SetName("Sub")
	parsePrintContent(cur, z, p)
	return z
}

//解析underline元素,会嵌套其它元素
//解析完后会进行计算
func parseUnderElement(cur dom.Element, p *PageSetting) *PrintUnderline {
	z := NewPrintUnderline()
	z.SetName("Underline")
	parsePrintContent(cur, z, p)
	return z
}

//解析doubleheight元素,会嵌套其它元素
//解析完后会进行计算
func parseDoubleHeightElement(cur dom.Element, p *PageSetting) *PrintDoubleHeight {
	z := NewPrintDoubleHeight()
	z.SetName("Doubleheight")
	parsePrintContent(cur, z, p)
	return z
}

//解析doublewidth元素,会嵌套其它元素
//解析完后会进行计算
func parseDoubleWidthElement(cur dom.Element, p *PageSetting) *PrintDoubleWidth {
	z := NewPrintDoubleWidth()
	z.SetName("Doublewidth")
	parsePrintContent(cur, z, p)
	return z
}

//解析turnpage元素
func parseTurnpageElement(cur dom.Element, p *PageSetting) *PrintTurnPage {
	z := NewPrintTurnPage()
	z.SetName("Turnpage")
	parsePrintContent(cur, z, p)
	return z
}

//解析Text元素,会嵌套其它元素
//解析完后会进行计算
func parseTextNode(cur dom.Text, p *PageSetting) *PrintText {
	z := NewPrintText(0, 0)
	z.SetName("Text")
	z.Con = cur.Data()
	return z
}

//解析通用的属性
func parsePosCommonAttribute(cur dom.Element, ps PrintPosInterface) {
	var tmp string
	var fval float64
	var err error
	tmp, err = getAttr(cur.Attr(), "Width")
	if err == nil {
		fval, err = strconv.ParseFloat(tmp, 32)
		if err != nil {
			utils.Error("can't format Height attribute %s,expected int value", tmp)
		}
		ps.SetWidth(float32(fval))
	}

	tmp, err = getAttr(cur.Attr(), "Height")
	if err == nil {
		fval, err = strconv.ParseFloat(tmp, 32)
		if err != nil {
			utils.Error("can't format Height attribute %s,expected int value", tmp)
		}
		ps.SetHeight(float32(fval))
	}

	tmp, err = getAttr(cur.Attr(), "Pos")
	ps.SetPos(tmp)

	tmp, err = getAttr(cur.Attr(), "Align")
	ps.SetAlign(tmp)

	tmp, err = getAttr(cur.Attr(), "Left")
	ps.SetLeft(tmp)

	tmp, err = getAttr(cur.Attr(), "Top")
	ps.SetTop(tmp)
}

//解析Rect元素
func parseRectElement(rect dom.Element, p *PageSetting) *PrintRect {
	z := NewPrintRect()
	parsePosCommonAttribute(rect, z)
	if z.GetWidth() == 0 {
		z.SetWidth(p.Width)
	}
	z.SetName("Rect")
	parsePrintContent(rect, z, p)
	return z
}

//解析Table元素
func parseTableElement(table dom.Element, p *PageSetting) *PrintTable {
	var (
		tmp  string
		err  error
		fval float64
	)
	z := NewPrintTable()
	parsePosCommonAttribute(table, z)
	if z.GetWidth() == 0 {
		z.SetWidth(p.Width)
	}
	z.SetName("Table")

	tmp, err = getAttr(table.Attr(), "LineSeperate")
	if err == nil {
		fval, err = strconv.ParseFloat(tmp, 32)
		if err != nil {
			utils.Error("can't format LineSeperate attribute %s,expected float32 value", tmp)
		}
		z.LineSeperate = float32(fval)
	}

	parsePrintContent(table, z, p)
	return z
}

//解析Tr元素
func parseTrElement(tr dom.Element, p *PageSetting) *PrintTr {
	z := NewPrintTr()
	parsePosCommonAttribute(tr, z)
	if z.GetWidth() == 0 {
		z.SetWidth(p.Width)
	}
	z.SetName("Tr")
	parsePrintContent(tr, z, p)
	return z
}

//解析Td元素
func parseTdElement(td dom.Element, p *PageSetting) *PrintTd {
	z := NewPrintTd()
	parsePosCommonAttribute(td, z)
	if z.GetWidth() == 0 {
		z.SetWidth(p.Width)
	}
	z.SetName("Td")
	parsePrintContent(td, z, p)
	return z
}

//解析需要打印的内容
//需要解析子结构,深度遍历->广度遍历
//cur:解析树中的当前节点
//par:父节点
func parsePrintContent(cur dom.Element, par PrintInterface, p *PageSetting) PrintInterface {
	var i int
	var (
		n    dom.Node
		t    dom.Element
		z    PrintInterface
		prev PrintInterface
		te   dom.Text
	)
	f := cur.ChildNodes()

	for i = 0; i < len(f); i++ {
		n = f[i]
		z = nil
		switch n.Kind() {
		case dom.ElementKind:
			t = n.(dom.Element)
			utils.Info("find element %s", t.Name())

			switch t.Name() {
			case "Br":
				z = parseBrElement(t, p)
			case "Table":
				z = parseTableElement(t, p)
			case "Rect":
				z = parseRectElement(t, p)
			case "Tr":
				z = parseTrElement(t, p)
			case "Td":
				z = parseTdElement(t, p)
			case "Bold":
				z = parseBoldElement(t, p)
			case "Sup":
				z = parseSupElement(t, p)
			case "Sub":
				z = parseSubElement(t, p)
			case "Underline":
				z = parseUnderElement(t, p)
			case "Doubleheight":
				z = parseDoubleHeightElement(t, p)
			case "Doublewidth":
				z = parseDoubleWidthElement(t, p)

			default:
				utils.Error("can't find element name: %s", t.Name())
				continue
			}

		case dom.TextKind:
			te = n.(dom.Text)
			t := strings.TrimSpace(te.Data())
			//可能是换行符
			//if t == "" || strings.HasPrefix(t, "\r") || strings.HasPrefix(t, "\n") {
			//utils.Info("detect text line character,text is %s", te.Data())
			//continue
			// changed by jinsl 20151222
			if t == "" {
				utils.Info("detect text line character,text is null")
				continue
			} else if strings.HasPrefix(t, "\r") {
				utils.Info("detect text line character,text is \\r")
				continue
			} else if strings.HasPrefix(t, "\n") {
				utils.Info("detect text line character,text is \\n")
				continue
			} else {
				utils.Info("find text node,content %s", te.Data())
				z = parseTextNode(te, p)
			}

		default:
			utils.Debug("find default node %s ,kind %d", n.String(), n.Kind())
		}
		if z == nil {
			continue
		}
		if par != nil {
			par.AddChild(z)
			z.SetParent(par)
		}
		//加入到列表
		if prev == nil {
			prev = z
		}
	}
	return prev
}

func parsePrintFormat(cur dom.Element) *PrintFacade {
	var i int
	var (
		n dom.Node
		t dom.Element
		p PrintFacade
	)
	f := cur.ChildNodes()

	for i = 0; i < len(f); i++ {
		n = f[i]
		switch n.Kind() {
		case dom.ElementKind:
			t = n.(dom.Element)
			switch t.Name() {
			case "Setting": // 解析设置
				p.S = parseSetting(t)
				break

			case "Content": //解析内容
				root := NewPrintRoot()
				root.SetName("root")
				page := NewPage(p.S)
				root.AddChild(page)
				page.SetParent(root)
				parsePrintContent(t, page, p.S.ps)
				p.V = root
				printPrintElement(p.V, 0)
				break

			default:
				utils.Error("can't find element name: %s", t.Name())
			}
			break
		}
	}
	return &p
}

//创建一个新的页
func NewPage(ps *PrintSetting) *PrintPage {
	page := NewPrintPage()
	page.InitPageConfig(ps)
	page.SetName("page")
	return page
}

func ParserPrintXml(s string) (*PrintFacade, error) {
	utils.Debug("begin parse pr2 string")
	db := dom.NewDOMBuilder(strings.NewReader(s), dom.NewDOMStore())
	d, err := db.Build()
	if err != nil {
		utils.Error("parse xml error %s", err.Error())
		return nil, err
	}
	tmp := parsePrintFormat(d.Root())
	return tmp, nil
}
