package print

//可打印元素的基础接口
type PrintInterface interface {
	AddChild(PrintInterface)
	GetChild() PrintInterface
	SetParent(PrintInterface)
	GetParent() PrintInterface
	GetNextSibling() PrintInterface
	AddNextSibling(PrintInterface)
	GetName() string
	SetName(string)
}

//控制元素
type PrintControlInterface interface {
	PrintInterface
	//传入当前的打印头,基本数据结构。返回一个标志位
	Decorate(pc *PrintCurState) int64
}

//定位元素
type PrintPosInterface interface {
	PrintInterface
	GetWidth() float32
	SetWidth(s float32)
	SetHeight(s float32)
	GetHeight() float32

	//传过来的报文中的值,未进行解析
	SetPos(s string)
	GetPos() string
	SetTop(s string)
	GetTop() string
	SetLeft(s string)
	GetLeft() string
	SetAlign(s string)
	GetAlign() string

	//运行中的值.如：本来元素是相对排列的,需要动态计算。
	//这些就是计算后的值
	SetRealTop(s float32)
	GetRealTop() float32
	SetRealLeft(s float32)
	GetRealLeft() float32
	SetRealWidth(s float32)
	GetRealWidth() float32
	SetRealHeight(s float32)
	GetRealHeight() float32

	//传入当前打印头,基本数据结构。返回一串打印
	CalcExtentValue(pc *PrintCurState)

	//当高度发生变化时,所采取的动作
	HeightChanged(pc *PrintCurState, height int64)

	//返回生成的PrintLine
	GetPrintLine() *PrintLine
}

//元素的基本信息
type PrintBasic struct {
	Name  string
	Value string
	Kind  string //用来指示所寄宿的宿主类型
	//子指针
	child PrintInterface
	//兄弟指针
	nextsibling PrintInterface
	//父节点
	parent PrintInterface
}

//获得子节点第一个元素
func (j *PrintBasic) GetChild() PrintInterface {
	return j.child
}

func (j *PrintBasic) GetParent() PrintInterface {
	return j.parent
}

func (j *PrintBasic) SetParent(s PrintInterface) {
	j.parent = s
}

//增加子节点
//放到最后一个节点上
func (j *PrintBasic) AddChild(s PrintInterface) {
	if j.child == nil {
		j.child = s
	} else {
		t := j.child
		for t.GetNextSibling() != nil {
			t = t.GetNextSibling()
			continue
		}
		t.AddNextSibling(s)
	}
}

//获得兄弟节点
func (j *PrintBasic) GetNextSibling() PrintInterface {
	return j.nextsibling
}

//增加兄弟节点
func (j *PrintBasic) AddNextSibling(s PrintInterface) {
	if j.nextsibling == nil {
		j.nextsibling = s
	} else {
		t := j.nextsibling
		for t.GetNextSibling() != nil {
			t = t.GetNextSibling()
			continue
		}
		t.AddNextSibling(s)
	}
}

func (j *PrintBasic) GetName() string {
	return j.Name
}

func (j *PrintBasic) SetName(s string) {
	j.Name = s
}

//另外的域
type PrintBasicPos struct {
	Width      float32
	Height     float32
	RealTop    float32
	RealLeft   float32
	RealWidth  float32
	RealHeight float32
	Left       string
	Top        string
	Pos        string //定位方式
	Align      string //对齐方式
	pl         *PrintLine
}

func (j *PrintBasicPos) GetWidth() float32 {
	return j.Width
}

func (j *PrintBasicPos) SetWidth(s float32) {
	j.Width = s
}

func (j *PrintBasicPos) GetHeight() float32 {
	return j.Height
}

func (j *PrintBasicPos) SetHeight(s float32) {
	j.Height = s
}

func (j *PrintBasicPos) GetPos() string {
	return j.Pos
}

func (j *PrintBasicPos) SetPos(s string) {
	j.Pos = s
}

func (j *PrintBasicPos) GetTop() string {
	return j.Top
}

func (j *PrintBasicPos) SetTop(s string) {
	j.Top = s
}

func (j *PrintBasicPos) GetLeft() string {
	return j.Left
}

func (j *PrintBasicPos) SetLeft(s string) {
	j.Left = s
}

func (j *PrintBasicPos) GetAlign() string {
	return j.Align
}

func (j *PrintBasicPos) SetAlign(s string) {
	j.Align = s
}

func (j *PrintBasicPos) SetRealTop(s float32) {
	j.RealTop = s
}

func (j *PrintBasicPos) GetRealTop() float32 {
	return j.RealTop
}

func (j *PrintBasicPos) SetRealLeft(s float32) {
	j.RealLeft = s
}

func (j *PrintBasicPos) GetRealLeft() float32 {
	return j.RealLeft
}

func (j *PrintBasicPos) SetRealWidth(s float32) {
	j.RealWidth = s
}

func (j *PrintBasicPos) GetRealWidth() float32 {
	return j.RealWidth
}

func (j *PrintBasicPos) SetRealHeight(s float32) {
	j.RealHeight = s
}

func (j *PrintBasicPos) GetRealHeight() float32 {
	return j.RealHeight
}

func (j *PrintBasicPos) GetPrintLine() *PrintLine {
	return j.pl
}

//复合结构
type PrintBasicExt struct {
	PrintBasic
	PrintBasicPos
}
