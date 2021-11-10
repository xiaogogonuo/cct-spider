package request

import (
	"io/ioutil"
	"net/http"
	"strconv"
)

// https://datacenter-web.eastmoney.com/api/data/v1/get?sortColumns=REPORT_DATE&sortTypes=-1&pageSize=500&pageNumber=17&reportName=RPT_INDUSTRY_INDEX&columns=REPORT_DATE%2CINDICATOR_VALUE&filter=(INDICATOR_ID%3D%22EMI00107664%22)
const (
	u1 = "https://datacenter-web.eastmoney.com/api/data/v1/get?sortColumns=REPORT_DATE&sortTypes=-1&pageSize=500&pageNumber="
	u2 = "&reportName=RPT_INDUSTRY_INDEX&columns=REPORT_DATE%2CINDICATOR_VALUE&filter=(INDICATOR_ID%3D%22"
	u3 = "%22)"
)

// EastMoney 东方财富
func EastMoney(sourceTargetCode string, page int) (b []byte, err error) {
	url := u1 + strconv.Itoa(page) + u2 + sourceTargetCode + u3
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	b, err = ioutil.ReadAll(resp.Body)
	return
}
