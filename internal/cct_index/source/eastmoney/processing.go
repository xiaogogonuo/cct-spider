package eastmoney

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/model"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/pkg/arithmetic"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/poster"
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
	"sort"
	"strings"
	"time"
)

// 数据加工处理

// 处理GDP同比
func processGDPTB(gdp GDP, ic *model.IndexConfig) (buffers []*model.Buffer) {
	for _, v := range gdp.Result.Data {
		buffer := &model.Buffer{}
		t, err := time.Parse("2006-01-02 03:04:05", v.ReportDate)
		if err != nil {
			logger.Error(err.Error())
			go poster.Poster(ic)
			break
		}
		year := fmt.Sprintf("%d", t.Year())
		seas := "Q" + fmt.Sprintf("%d", t.Month()/3)
		buffer.Date = year + seas
		buffer.TargetValue = arithmetic.TrimZero(fmt.Sprintf("%.2f", v.SumSame))
		buffers = append(buffers, buffer)
	}
	return
}

// 处理GDP环比
func processGDPHB(gdp GDP, ic *model.IndexConfig) (buffers []*model.Buffer) {
	var bufferTFList model.BufferTFList
	for _, v := range gdp.Result.Data {
		bufferTF := model.BufferTF{}
		t, err := time.Parse("2006-01-02 03:04:05", v.ReportDate)
		if err != nil {
			logger.Error(err.Error())
			go poster.Poster(ic)
			break
		}
		bufferTF.Date = t
		bufferTF.TargetValue = v.DomesticProductBase
		bufferTFList = append(bufferTFList, bufferTF)
	}
	sort.Sort(bufferTFList)

	for i := 0; i < len(bufferTFList)-2; i++ {

		theRightNowDate := bufferTFList[i+0].Date
		backOneSkipDate := bufferTFList[i+1].Date
		backTwoSkipDate := bufferTFList[i+2].Date
		switch theRightNowDate.Month() {
		case 3:
			rnd1 := theRightNowDate.AddDate(0, -3, 0)
			rnd2 := theRightNowDate.AddDate(0, -6, 0)
			if rnd1.Year() == backOneSkipDate.Year() && rnd1.Month() == backOneSkipDate.Month() &&
				rnd2.Year() == backTwoSkipDate.Year() && rnd2.Month() == backTwoSkipDate.Month() {
				tv := bufferTFList[i+1].TargetValue - bufferTFList[i+2].TargetValue
				if tv > 0 {
					buffer := &model.Buffer{}
					buffer.Date = fmt.Sprintf("%d", theRightNowDate.Year()) + "Q1"
					hb := (bufferTFList[i].TargetValue - tv) / tv * 100
					buffer.TargetValue = arithmetic.TrimZero(fmt.Sprintf("%.2f", hb))
					buffers = append(buffers, buffer)
				}
			}
		case 6:
			rnd1 := theRightNowDate.AddDate(0, -3, 0)
			if rnd1.Year() == backOneSkipDate.Year() && rnd1.Month() == backOneSkipDate.Month() {
				tv := bufferTFList[i+1].TargetValue
				if tv > 0 {
					buffer := &model.Buffer{}
					buffer.Date = fmt.Sprintf("%d", theRightNowDate.Year()) + "Q2"
					hb := (bufferTFList[i].TargetValue - tv) / tv * 100
					buffer.TargetValue = arithmetic.TrimZero(fmt.Sprintf("%.2f", hb))
					buffers = append(buffers, buffer)
				}
			}
		case 9, 12:
			rnd1 := theRightNowDate.AddDate(0, -3, 0)
			rnd2 := theRightNowDate.AddDate(0, -6, 0)
			if rnd1.Year() == backOneSkipDate.Year() && rnd1.Month() == backOneSkipDate.Month() &&
				rnd2.Year() == backTwoSkipDate.Year() && rnd2.Month() == backTwoSkipDate.Month() {
				tv1 := bufferTFList[i+0].TargetValue - bufferTFList[i+1].TargetValue
				tv2 := bufferTFList[i+1].TargetValue - bufferTFList[i+2].TargetValue
				if tv2 > 0 {
					buffer := &model.Buffer{}
					year := fmt.Sprintf("%d", theRightNowDate.Year())
					seas := "Q" + fmt.Sprintf("%d", theRightNowDate.Month()/3)
					buffer.Date = year + seas
					hb := (tv1 - tv2) / tv2 * 100
					buffer.TargetValue = arithmetic.TrimZero(fmt.Sprintf("%.2f", hb))
					buffers = append(buffers, buffer)
				}
			}
		}
	}
	return
}

// 补充扩展计算出拆借利率的涨跌幅等信息
func expansion(bufferTFList model.BufferTFList) (buffers []*model.Buffer) {
	for i := 0; i < len(bufferTFList)-1; i++ {
		cur := fmt.Sprintf("%.4f", bufferTFList[i+0].TargetValue) // 今天的拆借利率
		pre := fmt.Sprintf("%.4f", bufferTFList[i+1].TargetValue) // 昨天的拆借利率

		// 计算涨跌值和涨跌幅
		var (
			upDownValue = "" // 涨跌额
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
			upDownValue, // 涨跌额
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
