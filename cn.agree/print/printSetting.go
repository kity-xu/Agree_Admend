package print

//打印的设置
type PrintSetting struct {
	ps *PageSetting
	ss *OtherSetting
}

//纸张的设置
type PageSetting struct {
	Width        float32
	Height       float32
	Leftmargin   float32
	Rightmargin  float32
	Topmargin    float32
	Bottommargin float32
	LineInterval float32
	ColInterval  float32
}

//换页时的设置
type OtherSetting struct {
	Cpi float32
}

type PageRoot struct {
}
