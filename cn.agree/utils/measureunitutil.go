package utils

//从英尺换算为毫米
func ConvertFromInchTomm(f float32) float32 {
	return f * 25.4
}

//从毫米换算为英尺
func ConvertFrommmToInch(f float32) float32 {
	return f * 0.03938
}
