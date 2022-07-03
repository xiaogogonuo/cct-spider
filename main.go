package main

import (
	"fmt"
)

func main() {
	//marketCode := "001"
	//currencyCode := "CNY"
	//indicatorID := "001"
	//_ = "https://datacenter-web.eastmoney.com/api/data/v1/get?reportName=RPT_IMP_INTRESTRATEN&" +
	//	"columns=REPORT_DATE%2CREPORT_PERIOD%2CIR_RATE%2CCHANGE_RATE%2CINDICATOR_ID%2CLATEST_RECORD%2CMARKET%2CMARKET_CODE%2CCURRENCY%2CCURRENCY_CODE&" +
	//	"filter=(MARKET_CODE%3D%22001%22)(CURRENCY_CODE%3D%22CNY%22)(INDICATOR_ID%3D%22001%22)&" +
	//	"pageNumber=1&pageSize=20&sortTypes=-1&sortColumns=REPORT_DATE"
	//
	//u := "https://datacenter-web.eastmoney.com/api/data/v1/get?reportName=RPT_IMP_INTRESTRATEN&" +
	//	"columns=REPORT_DATE%2CREPORT_PERIOD%2CIR_RATE%2CCHANGE_RATE%2CINDICATOR_ID%2CLATEST_RECORD%2CMARKET%2CMARKET_CODE%2CCURRENCY%2CCURRENCY_CODE&" +
	//	"filter=(MARKET_CODE%3D%22" + marketCode + "%22)(CURRENCY_CODE%3D%22" + currencyCode + "%22)(INDICATOR_ID%3D%22" + indicatorID + "%22)&" +
	//	"pageNumber=1&pageSize=20&sortTypes=-1&sortColumns=REPORT_DATE"

	filter := "(MARKET_CODE%3D%22001%22)(CURRENCY_CODE%3D%22CNY%22)(INDICATOR_ID%3D%22001%22)"
	x := "https://datacenter-web.eastmoney.com/api/data/v1/get?reportName=RPT_IMP_INTRESTRATEN&" +
		"columns=REPORT_DATE%2CREPORT_PERIOD%2CIR_RATE%2CCHANGE_RATE%2CINDICATOR_ID%2CLATEST_RECORD%2CMARKET%2CMARKET_CODE%2CCURRENCY%2CCURRENCY_CODE&" +
		"filter=" + filter + "&pageNumber=1&pageSize=20&sortTypes=-1&sortColumns=REPORT_DATE"
	fmt.Println(x)

}
