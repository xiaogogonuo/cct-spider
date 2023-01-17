package eastmoney

import (
	"encoding/json"
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/api"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/model"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/pkg/arithmetic"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/pkg/downloader"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/pkg/safeguard"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/poster"
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
	"strings"
	"time"
)

// 东方财富-行业指数
// 适用指标：物流业景气指数、中国大宗商品指数、波罗的海干散货指数
// 行业指数接口每天记录都返回相同的两条，对其进行过滤
func eastMoneyIndustry(ic *model.IndexConfig) (buffers []*model.Buffer) {
	url := strings.ReplaceAll(api.Industry, "#", ic.SourceTargetCodeSpider)
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

	var s Industry
	if err = json.Unmarshal(body, &s); err != nil {
		logger.Error(err.Error())
		go poster.Poster(ic)
		return
	}
	if len(s.Result.Data) == 0 {
		go poster.Poster(ic)
		return
	}

	filter := make(map[string]struct{})
	for _, v := range s.Result.Data {
		buffer := &model.Buffer{}
		t, err := time.Parse("2006-01-02 03:04:05", v.ReportDate)
		if err != nil {
			logger.Error(err.Error())
			go poster.Poster(ic)
			break
		}
		if _, ok := filter[v.ReportDate]; ok {
			continue
		}
		filter[v.ReportDate] = struct{}{}
		buffer.Date = fmt.Sprintf("%d%02d%02d", t.Year(), t.Month(), t.Day())
		buffer.TargetValue = arithmetic.TrimZero(fmt.Sprintf("%.2f", v.IndicatorValue))
		buffers = append(buffers, buffer)
	}
	return
}
