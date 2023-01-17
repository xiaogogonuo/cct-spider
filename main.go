package main

import (
	"fmt"
	"regexp"
)

func main() {
	//s := "2017-06-01 00:00:00"
	//t, err := time.Parse("2006-01-02 03:04:05", s)
	//fmt.Println(int(t.Month()))
	//fmt.Println(err)
	//x := "xxx"
	//x += "yyy"
	//fmt.Println(fmt.Sprintf("%d%02d%02d", t.Year(), t.Month(), t.Day()))

	//x := "https://datacenter-web.eastmoney.com/api/data/v1/get?reportName=RPT_IMP_INTRESTRATEN&columns=REPORT_DATE%2CIR_RATE%2CCHANGE_RATE%2C&filter=#&pageNumber=1&pageSize=3&sortTypes=-1&sortColumns=REPORT_DATE"
	//y := strings.ReplaceAll(x, "#", "(MARKET_CODE%3D%22001%22)(CURRENCY_CODE%3D%22CNY%22)(INDICATOR_ID%3D%22001%22)")
	//fmt.Println(y)
	//b, err := downloader.SimpleGet("https://datacenter-web.eastmoney.com/api/data/v1/get")
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//fmt.Println(string(b))

	s := "2020年05月15日"
	o := regexp.MustCompile("[0-9]+").FindAllString(s, -1)
	fmt.Println(o)

}

// 结构体解析前端其他结构体响应不会报错，而是返回默认零值，如何解决？

// 判断接口是否正常运行
/*
gopher@192 ~ % curl http://106.37.165.121:18088/inf/chengtong/py/sy/baseTargetValue/saveRequest
{"timestamp":"2023-01-17 00:25:13","status":400,"error":"Bad Request","exception":"org.springframework.http.converter.HttpMessageNotReadableException","message":"Required request body is missing: public com.baosight.audit.api.dm.aa.response.operation.DmaaBaseTargetValueSaveResp com.baosight.audit.site.controller.py.sy.DmaaBaseTargetValueController.save(com.baosight.audit.api.dm.aa.command.operation.DmaaBaseTargetValueSaveCmd)","path":"/chengtong/py/sy/baseTargetValue/saveRequest"}%      gopher@192 ~ %
gopher@192 ~ %
gopher@192 ~ % curl http://106.37.165.121:18088/inf/chengtong/py/sy/baseTargetValue/saveRequestx
{"timestamp":"2023-01-17 00:25:20","status":404,"error":"Not Found","message":"No message available","path":"/chengtong/py/sy/baseTargetValue/saveRequestx"}%                                                                                         gopher@192 ~ % curl http://106.37.165.121/inf/chengtong/py/sy/baseTargetValue/saveRequest
curl: (7) Failed to connect to 106.37.165.121 port 80 after 40 ms: Connection refused
*/
