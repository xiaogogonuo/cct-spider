package eastmoney

import (
	"encoding/json"
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/api"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/model"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/pkg/downloader"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/pkg/safeguard"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/poster"
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
	"strings"
	"time"
)

// 东方财富-全球指数
// 适用：
// - 国债指数
// - 台湾加权
// - 韩国KOSPI
// - 俄罗斯RTS
// - 澳大利亚标普200
// - 路透CRB商品指数
// - 原油指数
// - AMAC造纸
// - 中证国有企业综合指数
// - 中证央企结构调整指数
// - 中证国企一带一路指数
// - 中证中央企业100指数
// - 中证央企创新驱动指数
// - 中证国有企业改革指数
// - 中证中央企业综合指数
// - 银华央企结构调整ETF
// - 博时央企结构调整ETF
// - 华夏央企结构调整ETF
func eastMoneyGlobal(ic *model.IndexConfig) (buffers []*model.Buffer) {
	url := strings.ReplaceAll(api.Global, "#", ic.SourceTargetCodeSpider)
	body, err := downloader.SimpleGet(url)

	if err != nil {
		if !safeguard.IsNetworkNormal() {
			logger.Error("请检查服务器的网络是否能联通外网")
			return
		}
		logger.Error(err.Error())
		go poster.Poster(ic)
		return
	}

	var s Global
	if err = json.Unmarshal(body, &s); err != nil {
		logger.Error(err.Error())
		go poster.Poster(ic)
		return
	}

	buffer := &model.Buffer{}
	buffer.Date = time.Unix(s.Data.F86, 0).Format("20060102")
	buffer.TargetValue = strings.Join([]string{
		fmt.Sprintf("%.2f", s.Data.F43/100),                    // 现价
		fmt.Sprintf("%.2f", s.Data.F169/100),                   // 涨跌额
		fmt.Sprintf("%.2f%s", s.Data.F170/100, "%"),            // 涨跌幅
		fmt.Sprintf("%.2f", s.Data.F44/100),                    // 最高
		fmt.Sprintf("%.2f", s.Data.F45/100),                    // 最低
		fmt.Sprintf("%.2f", s.Data.F60/100),                    // 昨收
		time.Unix(s.Data.F86, 0).Format("2006-01-02 15:04:05"), // 更新时间
	}, ",")
	buffers = append(buffers, buffer)

	return
}
