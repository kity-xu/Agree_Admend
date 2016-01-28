package print

import (
	"cn.agree/utils"
	"errors"
	"github.com/mathume/dom"
	"strings"
)

var ATTR_NOT_FOUND = errors.New("can't find attr ")
var NOT_INT_VALUE = errors.New("text value is not integer")

func getAttr(v []dom.Attribute, name string) (string, error) {
	var t dom.Attribute
	for i := 0; i < len(v); i++ {
		t = v[i]
		if strings.EqualFold(t.Name(), name) {
			return t.Value(), nil
		}
	}
	return "", ATTR_NOT_FOUND

}

//寻找最近的Pos节点
func findparentPosNode(pp PrintInterface) PrintPosInterface {
	var ppi PrintPosInterface
	var s PrintInterface
	var found bool
	s = pp.GetParent()
	for {
		if s == nil {
			utils.Error("can't find parent node,name is %s", pp.GetName())
			return nil
		}

		ppi, found = s.(PrintPosInterface)
		if found {
			return ppi
		}

		_, found2 := s.(PrintControlInterface)
		if found2 {
			s = s.GetParent()
			continue
		} else {
			utils.Error("find uncorrect kind of element name is [%s] ", s.GetName())
			return nil
		}

	}
}

//寻找最近的名称为name的Pos节点
func findparentPosNodeWithName(pp PrintInterface, name string) PrintPosInterface {
	var ppi PrintPosInterface
	var s PrintInterface
	var found bool
	s = pp.GetParent()
	for {
		if s == nil {
			utils.Error("can't find parent node,name is %s", pp.GetName())
			return nil
		}

		ppi, found = s.(PrintPosInterface)
		if !strings.EqualFold(ppi.GetName(), name) {
			s = s.GetParent()
			continue
		}
		if found {
			return ppi
		}

		_, found2 := s.(PrintControlInterface)
		if found2 {
			s = s.GetParent()
			continue
		} else {
			utils.Error("find uncorrect kind of element name is [%s] ", s.GetName())
			return nil
		}

	}
}

//寻找子节点的字段之和
//ctype:1 寻找宽度 2:寻找高度
func findChildCalValue(pp PrintInterface, ctype int) float32 {
	var ppi PrintPosInterface
	var s PrintInterface
	var sibling PrintInterface
	var found bool
	var rvalue float32
	s = pp.GetChild()
	for {
		if s == nil {
			utils.Error("can't find child node,name is %s", pp.GetName())
			return 0
		}

		ppi, found = s.(PrintPosInterface)
		if found {
			if ctype == 1 {
				rvalue += ppi.GetRealWidth()
			} else {
				rvalue += ppi.GetRealHeight()
			}

			//计算所有相邻节点
			sibling = ppi
			for {
				sibling = sibling.GetNextSibling()
				if sibling != nil {
					rvalue += findChildCalValue(sibling, ctype)
				} else {
					break
				}
				//sibling = sibling.GetNextSibling() //it wrong changed by jinsl 20151222
			}
			return rvalue
		}

		_, found2 := s.(PrintControlInterface)
		if found2 {
			s = s.GetChild()
			continue
		} else {
			utils.Error("find uncorrect kind of element name is [%s] ", s.GetName())
			return 0
		}

	}
}

//寻找子节点的字段之和
//ctype:1 寻找宽度 2:寻找高度
func findChildMaxValue(pp PrintInterface, ctype int) float32 {
	var ppi PrintPosInterface
	var s PrintInterface
	var sibling PrintInterface
	var found bool
	var rvalue float32
	var svalue float32
	s = pp.GetChild()
	for {
		if s == nil {
			utils.Error("can't find child node,name is %s", pp.GetName())
			return 0
		}

		ppi, found = s.(PrintPosInterface)
		if found {
			if ctype == 1 {
				rvalue += ppi.GetRealWidth()
			} else {
				rvalue += ppi.GetRealHeight()
			}

			//计算所有相邻节点
			sibling = ppi
			for {
				sibling = sibling.GetNextSibling()
				if sibling != nil {
					svalue = findChildCalValue(sibling, ctype)
					if svalue > rvalue {
						rvalue = svalue
					}
				//sibling = sibling.GetNextSibling() //it wrong changed by jinsl 20151222
				} else {
					break
				}
			}
			return rvalue
		}

		_, found2 := s.(PrintControlInterface)
		if found2 {
			s = s.GetChild()
			continue
		} else {
			utils.Error("find uncorrect kind of element name is [%s] ", s.GetName())
			return 0
		}

	}
}

//根据element的值寻找text下面的值
func findTextNode(cur dom.Element) (string, error) {
	f := cur.ChildNodes()
	var i int
	var n dom.Node
	var z dom.Text
	for i = 0; i < len(f); i++ {
		n = f[i]
		switch n.Kind() {
		case dom.TextKind:
			z = n.(dom.Text)
			tmp := z.Data()
			return tmp, nil
		}
	}
	return "", NOT_FOUND_VALUE
}

//打印出PrintElement列表
func printPrintElement(p PrintInterface, ident int) {
	if p == nil {
		return
	}
	localident := ident + 2
	s := strings.Repeat("-", ident)

	utils.Debug("%sname:%+v", s, p)
	printPrintElement(p.GetChild(), localident)
	printPrintElement(p.GetNextSibling(), ident)

}

//获得字符的ascii码长度。如果是中文,则算作2个ascii码
//仅支持中文和英文。如果需要别的语言,可能还要做处理
func getAsciiLength(s string) int {
	var c int
	for _, runeValue := range s {
		if runeValue < 128 && runeValue >= 0 {
			c++ //ascii码算1
		} else {
			c = c + 2 //中文长度算作2
		}
	}
	return c
}

//获得指定长度的中文字符,返回Len+1的位置
func getStringRuneLen(s string, slen int) int {
	var c int
	for i, runeValue := range s {
		if runeValue < 128 && runeValue >= 0 {
			c++ //ascii码算1
		} else {
			c = c + 2 //中文长度算作2
		}
		if slen < c {
			return i
		}
	}
	return -1
}

//获得长度所能容纳的字符数
//mm:表示要测量的长度
//colinteral:表示单个ascii字符长度
func getAsciiLengthWithMM(mm float32, colinteral float32) int {
	return int(mm / colinteral)
}
