package utils

import (
	"bytes"
)

//获得
func IndexByteRune(buf []byte, r rune, occ int) int {
	var curocc = 0
	var curindex = -1
	var culindex = -1
	for curocc < occ {
		curindex = bytes.IndexRune(buf[culindex+1:], r)
		if curindex == -1 {
			return -1
		}
		culindex = culindex + curindex + 1
		curocc++
	}
	return culindex
}
