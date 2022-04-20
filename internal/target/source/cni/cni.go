package cni

import (
	"encoding/json"
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/target/model"
	"github.com/xiaogogonuo/cct-spider/internal/target/pkg/downloader"
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
	"time"
	"unsafe"
)

var APICniCNYXTarget = "http://hq.cnindex.com.cn/market/market/getIndexRealTimeData?indexCode=CNYX"

// SpiderCniCNYXTarget 爬取国证指数网站的人民币汇率
// 适用指标：
// - 人民币汇率
//   • 页面展示接口：http://www.cnindex.com.cn/module/index-detail-CNYX.html?act_menu=1
//   • 数据获取接口：http://hq.cnindex.com.cn/market/market/getIndexRealTimeData?indexCode=CNYX
func SpiderCniCNYXTarget() (responses []model.Response) {
	body, err := downloader.SimpleGet(APICniCNYXTarget)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	stringBody := *(*string)(unsafe.Pointer(&body))
	var ct cniCNYXTarget
	if err := json.Unmarshal([]byte(stringBody), &ct); err != nil {
		logger.Error(err.Error())
		return
	}
	data := ct.Data.Data
	lastOne := data[len(data)-1] // 倒数第一个数据
	lastTwo := data[len(data)-2] // 倒数第二个数据
	var timeStamp float64
	var v []interface{}
	if lastOne[1] != nil {
		timeStamp = lastOne[0].(float64)
		v = lastOne
	} else {
		timeStamp = lastTwo[0].(float64)
		v = lastTwo
	}
	tm := time.Unix(int64(timeStamp)/1000, 0)
	var response model.Response
	response.Date = tm.Format("20060102")
	response.TargetValue = fmt.Sprintf("%.2f,%.2f,%.2f%s,%.2f,%.2f,%.2f,%s",
		v[1].(float64), // 现价
		v[6].(float64), // 涨跌
		v[7].(float64), // 涨跌幅
		"%",
		v[2].(float64),                   // 最高
		v[4].(float64),                   // 最低
		v[11].(float64),                  // 昨收
		tm.Format("2006-01-02 15:04:05"), // 更新日期
	)
	responses = append(responses, response)
	return
}
