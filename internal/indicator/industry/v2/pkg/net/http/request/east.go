package request

import (
	"io/ioutil"
	"net/http"
)

const (
	eastURLP1 = "http://dcfm.eastmoney.com/em_mutisvcexpandinterface/api/js/get?" +
		"type=HYZS_PageStock&st=DATADATE&sr=-1&filter=(ID%3D%27"
	eastURLP2 = "%27)&token="
)

// TODO: scale
var token = "894050c76af8597a853f5b408b759f5d"

// VisitEastMoney 东方财富
func VisitEastMoney(targetCode string) (b []byte, err error) {
	url := eastURLP1 + targetCode + eastURLP2 + token
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	b, err = ioutil.ReadAll(resp.Body)
	return
}