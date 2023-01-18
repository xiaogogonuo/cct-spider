package fx

import (
	"fmt"
	"strconv"
	"time"
)

type ExchangeSpecial struct {
	S string   `json:"s"` // 状态
	T []string `json:"t"` // 当前时间(暂时不用)
	C []string `json:"c"` // 最新价
	O []string `json:"o"` // 开盘价(暂时不用)
	H []string `json:"h"` // 最高
	L []string `json:"l"` // 最低
	P []string `json:"p"` // 昨收
}

func (es ExchangeSpecial) StatusValid() bool {
	return es.S == "ok"
}

func (es ExchangeSpecial) DataValid() bool {
	return es.C != nil && len(es.C) > 0 && es.C[0] != "" &&
		es.H != nil && len(es.H) > 0 && es.H[0] != "" &&
		es.L != nil && len(es.L) > 0 && es.L[0] != "" &&
		es.P != nil && len(es.P) > 0 && es.P[0] != ""
}

func (es ExchangeSpecial) priceFormat(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

func (es ExchangeSpecial) formatC() (float64, error) {
	return es.priceFormat(es.C[0])
}

func (es ExchangeSpecial) formatH() (float64, error) {
	return es.priceFormat(es.H[0])
}

func (es ExchangeSpecial) formatL() (float64, error) {
	return es.priceFormat(es.L[0])
}

func (es ExchangeSpecial) formatP() (float64, error) {
	return es.priceFormat(es.P[0])
}

func (es ExchangeSpecial) Handler() (data string, err error) {
	c, err := es.formatC()
	if err != nil {
		return
	}
	h, err := es.formatH()
	if err != nil {
		return
	}
	l, err := es.formatL()
	if err != nil {
		return
	}
	p, err := es.formatP()
	if err != nil {
		return
	}
	zd := c - p
	zdf := zd / p * 100
	data = fmt.Sprintf("%.2f,%.2f,%.2f", c, zd, zdf) + "%," +
		fmt.Sprintf("%.2f,%.2f,%.2f", h, l, p)
	tm := time.Now().Format("2006-01-02 15:04:05")
	data = data + "," + tm
	return
}
