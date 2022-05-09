package fx678

import (
	"fmt"
	"strconv"
	"time"
)

// fxTarget 汇通财经接口返回的数据结构
type fxTarget struct {
	S string   `json:"s"` // 状态
	T []string `json:"t"` // 当前时间(暂时不用)
	C []string `json:"c"` // 最新价
	O []string `json:"o"` // 开盘价(暂时不用)
	H []string `json:"h"` // 最高
	L []string `json:"l"` // 最低
	P []string `json:"p"` // 昨收
}

func (ft fxTarget) StatusValid() bool {
	return ft.S == "ok"
}

func (ft fxTarget) DataValid() bool {
	return ft.C != nil && len(ft.C) > 0 && ft.C[0] != "" &&
		ft.H != nil && len(ft.H) > 0 && ft.H[0] != "" &&
		ft.L != nil && len(ft.L) > 0 && ft.L[0] != "" &&
		ft.P != nil && len(ft.P) > 0 && ft.P[0] != ""
}

func (ft fxTarget) priceFormat(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

func (ft fxTarget) formatC() (float64, error) {
	return ft.priceFormat(ft.C[0])
}

func (ft fxTarget) formatH() (float64, error) {
	return ft.priceFormat(ft.H[0])
}

func (ft fxTarget) formatL() (float64, error) {
	return ft.priceFormat(ft.L[0])
}

func (ft fxTarget) formatP() (float64, error) {
	return ft.priceFormat(ft.P[0])
}

func (ft fxTarget) Handler() (data string, err error) {
	c, err := ft.formatC()
	if err != nil {
		return
	}
	h, err := ft.formatH()
	if err != nil {
		return
	}
	l, err := ft.formatL()
	if err != nil {
		return
	}
	p, err := ft.formatP()
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