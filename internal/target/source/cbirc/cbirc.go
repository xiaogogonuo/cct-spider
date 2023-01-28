package cbirc

import (
	"encoding/json"
	"github.com/xiaogogonuo/cct-spider/internal/target/model"
	"github.com/xiaogogonuo/cct-spider/internal/target/pkg/downloader"
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
)

// APICBIRCTarget 银保监会指标接口
var APICBIRCTarget = "http://www.irc.gov.cn/cbircweb/solr/totalStaSerch"

// SpiderCBIRCTarget 爬取银保监会指标
// 适用指标：
// - 商业银行不良贷款余额季度
//   • 页面展示接口：http://www.cbirc.gov.cn/cn/view/pages/index/jiansuo.html?keyWords=银行业保险业主要监管指标数据
//   • 数据获取接口：http://www.cbirc.gov.cn/cbircweb/solr/totalStaSerch
func SpiderCBIRCTarget(targetNameSpider string) (responses []model.Response) {
	for i := 1; ; i++ {
		payload := PayLoad{
			KeyWords:   targetNameSpider,
			PageNo:     i,
			PageSize:   "10",
			SearchType: "1",
			Title:      targetNameSpider,
		}
		m, _ := json.Marshal(payload)
		header := map[string]string{"Content-Type": "application/json"}
		body, err := downloader.Post(APICBIRCTarget, m, header)
		if err != nil {
			logger.Error(err.Error())
			continue
		}
		var front cBIRCTarget
		if err := json.Unmarshal(body, &front); err != nil {
			logger.Error(err.Error())
			continue
		}
		if len(front.Data.Lists) == 0 {
			break
		}
		_responses := cBIRCTargetPipeline(front)
		responses = append(responses, _responses...)
	}
	return
}
