package irc

import (
	"encoding/json"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/api"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/model"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/pkg/downloader"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/pkg/safeguard"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/poster"
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
)

// PayLoad 银保监会POST方法请求体结构
type PayLoad struct {
	ItemName   string `json:"itemName"`
	KeyWords   string `json:"keyWords"`
	MainType   string `json:"mainType"`
	PageNo     int    `json:"pageNo"`
	PageSize   string `json:"pageSize"`
	SearchType string `json:"serchType"`
	Title      string `json:"title"`
	Type       string `json:"type"`
}

// 东方财富-拆借利率
// 适用：
// - 商业银行不良贷款余额季度
func ircEconomic(ic *model.IndexConfig) (buffers []*model.Buffer) {
	for i := 1; ; i++ {
		payload := PayLoad{
			KeyWords:   ic.TargetNameSpider,
			PageNo:     i,
			PageSize:   "10",
			SearchType: "1",
			Title:      ic.TargetNameSpider,
		}
		m, _ := json.Marshal(payload)
		header := map[string]string{"Content-Type": "application/json"}
		body, err := downloader.Post(api.NPLoan, m, header)

		if err != nil {
			if !safeguard.IsNetworkNormal() {
				logger.Error("请检查服务器的网络是否能联通外网")
				return
			}
			logger.Error(err.Error())
			go poster.Poster(ic)
			return
		}
		var s Irc
		if err = json.Unmarshal(body, &s); err != nil {
			logger.Error(err.Error())
			go poster.Poster(ic)
			return
		}
		if len(s.Data.Lists) == 0 {
			break
		}

		buffers = append(buffers, getIrcBuffer(s)...)
	}

	return
}
