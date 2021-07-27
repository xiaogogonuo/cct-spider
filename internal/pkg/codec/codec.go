package codec

import (
	"bytes"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"unicode/utf8"
)

func isUtf8(s []byte) bool {
	return utf8.Valid(s)
}

func IsGbk(data []byte) bool {
	if isUtf8(data) {
		return false
	}
	for i := 0; i < len(data); {
		if data[i] <= 0x7f {
			i++
			continue
		}
		if data[i] >= 0x81 &&
			data[i] <= 0xfe &&
			data[i+1] >= 0x40 &&
			data[i+1] <= 0xfe &&
			data[i+1] != 0xf7 {
			i += 2
			continue
		}
		return false

	}

	return true
}

func GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}

	return d, nil
}