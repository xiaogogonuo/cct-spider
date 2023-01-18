package eastmoney

import (
	"encoding/json"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/api"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/model"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/pkg/downloader"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/pkg/safeguard"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/poster"
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
	"sort"
	"strings"
	"time"
)

// 东方财富-拆借利率
// 适用：
// - 上海银行同业拆借市场-Shibor人民币-隔夜
// - 上海银行同业拆借市场-Shibor人民币-1周
// - 上海银行同业拆借市场-Shibor人民币-1月
// - 上海银行同业拆借市场-Shibor人民币-3月
// - 上海银行同业拆借市场-Shibor人民币-1年
// - 伦敦银行同业拆借市场-Libor美元-隔夜(O/N)
// - 伦敦银行同业拆借市场-Libor美元-1月
// - 伦敦银行同业拆借市场-Libor美元-3月
func eastMoneyLendingRate(ic *model.IndexConfig) (buffers []*model.Buffer) {
	body, err := downloader.SimpleGet(strings.ReplaceAll(api.LendingRate, "#", ic.SourceTargetCodeSpider))

	if err != nil {
		if !safeguard.IsNetworkNormal() {
			logger.Error("请检查服务器的网络是否能联通外网")
			return
		}
		logger.Error(err.Error())
		go poster.Poster(ic)
		return
	}

	var s LendingRate
	if err = json.Unmarshal(body, &s); err != nil {
		logger.Error(err.Error())
		go poster.Poster(ic)
		return
	}
	if len(s.Result.Data) == 0 {
		go poster.Poster(ic)
		return
	}

	var bufferTFList model.BufferTFList
	for _, v := range s.Result.Data {
		bufferTF := model.BufferTF{}
		t, err := time.Parse("2006-01-02 03:04:05", v.ReportDate)
		if err != nil {
			logger.Error(err.Error())
			go poster.Poster(ic)
			break
		}
		bufferTF.Date = t
		bufferTF.TargetValue = v.IrRate
		bufferTFList = append(bufferTFList, bufferTF)
	}
	sort.Sort(bufferTFList)
	buffers = expansion(bufferTFList)
	return
}
