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
	"sort"
	"strings"
	"time"
)

// eastMoneyLendingRate 拆借利率
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

// 补充扩展计算出拆借利率的涨跌幅等信息
func expansion(bufferTFList model.BufferTFList) (buffers []*model.Buffer) {
	for i := 0; i < len(bufferTFList)-1; i++ {
		cur := fmt.Sprintf("%.4f", bufferTFList[i+0].TargetValue) // 今天的拆借利率
		pre := fmt.Sprintf("%.4f", bufferTFList[i+1].TargetValue) // 昨天的拆借利率

		// 计算涨跌值和涨跌幅
		var (
			upDownValue = "" // 涨跌值
			upDownRange = "" // 涨跌幅
		)
		if cur == "0.0000" || pre == "0.0000" {
			upDownValue = ""
			upDownRange = ""
		} else {
			delta := bufferTFList[i].TargetValue - bufferTFList[i+1].TargetValue
			upDownValue = fmt.Sprintf("%.2f", delta)
			upDownRange = fmt.Sprintf("%.2f%s", (delta)/bufferTFList[i+1].TargetValue*100, "%")
		}

		buffer := &model.Buffer{}
		buffer.Date = bufferTFList[i].Date.Format("20060102")
		buffer.TargetValue = strings.Join([]string{
			cur,         // 现价
			upDownValue, // 涨跌
			upDownRange, // 涨跌幅
			"",          // 最高
			"",          // 最低
			pre,         // 昨收
			bufferTFList[i].Date.Format("2006-01-02 03:04:05"), // 更新时间
		}, ",")
		buffers = append(buffers, buffer)
	}
	return
}
