package executor

import (
	"encoding/json"
	"github.com/xiaogogonuo/cct-spider/internal/indicator/v1/pkg/cook"
	"github.com/xiaogogonuo/cct-spider/internal/indicator/v1/pkg/request"
	"github.com/xiaogogonuo/cct-spider/internal/indicator/v1/pkg/response"
	"github.com/xiaogogonuo/cct-spider/internal/indicator/v1/pkg/urllib"
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
	"strings"
	"time"
)

// end return the last value of a slice
func end(s []string) string {
	length := len(s)
	return s[length-1]
}

func replaceSeason(date string) string {
	date = strings.ReplaceAll(date, "A", "Q1")
	date = strings.ReplaceAll(date, "B", "Q2")
	date = strings.ReplaceAll(date, "C", "Q3")
	date = strings.ReplaceAll(date, "D", "Q4")
	return date
}

// Execute
// 查询年度指标数据的步骤
// 1、访问https://data.stats.gov.cn/easyquery.htm?cn=%s&zb=%s获取cookie，
// 其中cn代表查询类型的代码(年度、季度、月度等)，zb代表某个指标的代码(GDP、PPI、PMI等)
// 2、带着上一步产生的cookie访问URL，只需修改变量LAST-N，即查询多少年的数据
func Execute(param urllib.Param, cn, zb string) (row [][]string) {
	cookie := cook.Cookie(cn, zb)
	req := request.Request{
		URL: param.Encode(),
		Cookie: cookie,
	}
	resBody, err := req.Visit()
	time.Sleep(time.Second * 3)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	var res response.Response
	if err = json.Unmarshal(resBody, &res); err != nil {
		logger.Error(err.Error())
		return
	}
	nodes := res.ReturnData.DataNodes
	for _, node := range nodes {
		// Value := node.Data.Data
		StrValue := node.Data.StrData
		date := end(strings.Split(node.Code, "."))
		date = replaceSeason(date)
		row = append(row, []string{date, StrValue})
	}
	return
}