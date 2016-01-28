package utils

import (
	"strconv"
	"strings"
)

//把整数转换成字符串
//a:表示需要转换的整数
//padlen:需要的长度。如果转换后的长度大于padlen,则不进行字符串切割
//b:需要填充的字符
func FormatInt(a int, padlen int, b string, base int) string {
	var s string
	if base == 10 {
		s = strconv.Itoa(a)
	} else if base == 16 {
		s = strconv.FormatInt(int64(a), 16)
	} else {
		return ""
	}

	if len(s) < padlen {
		return (strings.Repeat(b, padlen-len(s))) + s
	}
	return s
}
