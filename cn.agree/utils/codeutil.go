package utils

import (
	"bytes"
	"code.google.com/p/go.text/encoding/simplifiedchinese"
	"code.google.com/p/go.text/transform"
	"errors"
	"io"
	"io/ioutil"
)

//设备类别
const (
	GBK = iota
	GB2312
	MAX_LANGUAGE
)

var NOT_SUPPORTED_CODE = errors.New("not supported codepage")

func TransUTF8FromCode(src []byte, code int) ([]byte, error) {
	var rInUTF8 io.Reader
	switch code {
	case GBK:
		rInUTF8 = transform.NewReader(bytes.NewReader(src), simplifiedchinese.GBK.NewDecoder())
		break
	case GB2312:
		rInUTF8 = transform.NewReader(bytes.NewReader(src), simplifiedchinese.HZGB2312.NewDecoder())
		break
	default:
		return nil, NOT_SUPPORTED_CODE
	}

	out, _ := ioutil.ReadAll(rInUTF8)
	return out, nil
}

func TransGBKFromUTF8(src string) ([]byte, error) {
	dst := make([]byte, len(src)*2)
	tr := simplifiedchinese.GB18030.NewEncoder()
	nDst, _, err := tr.Transform(dst, []byte(src), true)
	if err != nil {
		return []byte(src), err
	}
	return dst[:nDst], nil
}
