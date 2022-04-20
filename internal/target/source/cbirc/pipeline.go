package cbirc

import (
	"github.com/xiaogogonuo/cct-spider/internal/target/model"
	"strings"
)

// cBIRCTargetPipeline 中国银行保险监督管理委员会数据处理管道
func cBIRCTargetPipeline(target cBIRCTarget) (responses []model.Response) {
	for _, list := range target.Data.Lists {
		if strings.Contains(list.DocTitle, "季度") {
			if len(list.Keyword) > 1 {
				var response model.Response
				response.Date = target.GetDate(list.DocTitle)
				response.TargetValue = target.GetTargetValue(list.Keyword[1])
				responses = append(responses, response)
			}
		}
	}
	return
}
