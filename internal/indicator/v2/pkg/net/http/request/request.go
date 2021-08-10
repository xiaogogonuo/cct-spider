package request

import (
	"errors"
)

const (
	EasyQuery = "https://data.stats.gov.cn/easyquery.htm?"
)

func Request(cn, zb string, param Param) (resBody []byte, err error) {
	cookie, err := GetCookie(cn, zb)
	if err != nil {
		return
	}
	if cookie == nil {
		err = errors.New(zb + " get cookie failed")
		return
	}
	resBody, err = GetIndicator(EasyQuery + param.ParamEncode(), cookie)
    return
}