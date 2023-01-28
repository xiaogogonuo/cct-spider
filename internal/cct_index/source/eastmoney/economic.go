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
	"regexp"
	"sort"
	"time"
)

// 宏观指数

// 国内生产总值
func eastMoneyEconomicGDP(ic *model.IndexConfig) (buffers []*model.Buffer) {
	body, err := downloader.SimpleGet(api.GDP)
	if err != nil {
		if !safeguard.IsNetworkNormal() {
			logger.Error("请检查服务器的网络是否能联通外网")
			return
		}
		logger.Error(err.Error())
		go poster.Poster(ic)
		return
	}

	var gdp GDP
	if err = json.Unmarshal(body, &gdp); err != nil {
		logger.Error(err.Error())
		go poster.Poster(ic)
		return
	}
	if len(gdp.Result.Data) == 0 {
		go poster.Poster(ic)
		return
	}

	switch ic.TargetCode {
	case "HG00001":
		buffers = processGDPTB(gdp, ic)
	case "HG00098":
		buffers = processGDPHB(gdp, ic)
	}

	return
}

// 工业增加值
func eastMoneyEconomicIAV(ic *model.IndexConfig) (buffers []*model.Buffer) {
	body, err := downloader.SimpleGet(api.IAV)

	if err != nil {
		if !safeguard.IsNetworkNormal() {
			logger.Error("请检查服务器的网络是否能联通外网")
			return
		}
		logger.Error(err.Error())
		go poster.Poster(ic)
		return
	}

	var s IAV
	if err = json.Unmarshal(body, &s); err != nil {
		logger.Error(err.Error())
		go poster.Poster(ic)
		return
	}
	if len(s.Result.Data) == 0 {
		go poster.Poster(ic)
		return
	}

	for _, v := range s.Result.Data {
		buffer := &model.Buffer{}
		t, err := time.Parse("2006-01-02 03:04:05", v.ReportDate)
		if err != nil {
			logger.Error(err.Error())
			go poster.Poster(ic)
			break
		}
		buffer.Date = fmt.Sprintf("%d%02d", t.Year(), t.Month())
		switch ic.TargetCode {
		case "HG00016": // 工业增加值同比增长
			buffer.TargetValue = arithmetic.TrimZero(fmt.Sprintf("%.2f", v.BaseSame))
		case "HG00017": // 工业增加值累计增长
			buffer.TargetValue = arithmetic.TrimZero(fmt.Sprintf("%.2f", v.BaseAccumulate))
		}
		buffers = append(buffers, buffer)
	}

	return
}

// 社会消费品零售总额
func eastMoneyEconomicXFP(ic *model.IndexConfig) (buffers []*model.Buffer) {
	body, err := downloader.SimpleGet(api.XFP)

	if err != nil {
		if !safeguard.IsNetworkNormal() {
			logger.Error("请检查服务器的网络是否能联通外网")
			return
		}
		logger.Error(err.Error())
		go poster.Poster(ic)
		return
	}

	var s XFP
	if err = json.Unmarshal(body, &s); err != nil {
		logger.Error(err.Error())
		go poster.Poster(ic)
		return
	}
	if len(s.Result.Data) == 0 {
		go poster.Poster(ic)
		return
	}

	for _, v := range s.Result.Data {
		buffer := &model.Buffer{}
		t, err := time.Parse("2006-01-02 03:04:05", v.ReportDate)
		if err != nil {
			logger.Error(err.Error())
			go poster.Poster(ic)
			break
		}
		buffer.Date = fmt.Sprintf("%d%02d", t.Year(), t.Month())
		switch ic.TargetCode {
		case "HG00027": // 社会消费品零售总额当期值
			buffer.TargetValue = arithmetic.TrimZero(fmt.Sprintf("%.2f", v.RetailTotal))
		case "HG00028": // 社会消费品零售总额累计值
			buffer.TargetValue = arithmetic.TrimZero(fmt.Sprintf("%.2f", v.RetailTotalAccumulate))
		case "HG00029": // 社会消费品零售总额同比增长
			buffer.TargetValue = arithmetic.TrimZero(fmt.Sprintf("%.2f", v.RetailTotalSame))
		case "HG00030": // 社会消费品零售总额累计增长
			buffer.TargetValue = arithmetic.TrimZero(fmt.Sprintf("%.2f", v.RetailAccumulateSame))
		}
		buffers = append(buffers, buffer)
	}

	return
}

// 货币和准货币(M2)供应量
func eastMoneyEconomicM2(ic *model.IndexConfig) (buffers []*model.Buffer) {
	body, err := downloader.SimpleGet(api.M2)

	if err != nil {
		if !safeguard.IsNetworkNormal() {
			logger.Error("请检查服务器的网络是否能联通外网")
			return
		}
		logger.Error(err.Error())
		go poster.Poster(ic)
		return
	}

	var s M2
	if err = json.Unmarshal(body, &s); err != nil {
		logger.Error(err.Error())
		go poster.Poster(ic)
		return
	}
	if len(s.Result.Data) == 0 {
		go poster.Poster(ic)
		return
	}

	for _, v := range s.Result.Data {
		buffer := &model.Buffer{}
		t, err := time.Parse("2006-01-02 03:04:05", v.ReportDate)
		if err != nil {
			logger.Error(err.Error())
			go poster.Poster(ic)
			break
		}
		buffer.Date = fmt.Sprintf("%d%02d", t.Year(), t.Month())
		switch ic.TargetCode {
		case "HG00006": // 货币和准货币(M2)供应量期末值
			buffer.TargetValue = arithmetic.TrimZero(fmt.Sprintf("%.2f", v.BasicCurrency))
		case "HG00007": // 货币和准货币(M2)供应量同比增长
			buffer.TargetValue = arithmetic.TrimZero(fmt.Sprintf("%.2f", v.BasicCurrencySame))
		}
		buffers = append(buffers, buffer)
	}

	return
}

// 居民消费价格指数
func eastMoneyEconomicCPI(ic *model.IndexConfig) (buffers []*model.Buffer) {
	body, err := downloader.SimpleGet(api.CPI)

	if err != nil {
		if !safeguard.IsNetworkNormal() {
			logger.Error("请检查服务器的网络是否能联通外网")
			return
		}
		logger.Error(err.Error())
		go poster.Poster(ic)
		return
	}

	var s CPI
	if err = json.Unmarshal(body, &s); err != nil {
		logger.Error(err.Error())
		go poster.Poster(ic)
		return
	}
	if len(s.Result.Data) == 0 {
		go poster.Poster(ic)
		return
	}

	for _, v := range s.Result.Data {
		buffer := &model.Buffer{}
		t, err := time.Parse("2006-01-02 03:04:05", v.ReportDate)
		if err != nil {
			logger.Error(err.Error())
			go poster.Poster(ic)
			break
		}
		buffer.Date = fmt.Sprintf("%d%02d", t.Year(), t.Month())
		switch ic.TargetCode {
		case "HG00003", "HG00004": // 居民消费价格指数、居民消费价格指数当月  (两者是一致的，有点多余了)
			buffer.TargetValue = arithmetic.TrimZero(fmt.Sprintf("%.2f", v.NationalBase))
		case "HG00088": // CPI同比增速月度
			buffer.TargetValue = arithmetic.TrimZero(fmt.Sprintf("%.2f", v.NationalSame))
		}
		buffers = append(buffers, buffer)
	}

	return
}

// 采购经理人指数
func eastMoneyEconomicPMI(ic *model.IndexConfig) (buffers []*model.Buffer) {
	body, err := downloader.SimpleGet(api.PMI)

	if err != nil {
		if !safeguard.IsNetworkNormal() {
			logger.Error("请检查服务器的网络是否能联通外网")
			return
		}
		logger.Error(err.Error())
		go poster.Poster(ic)
		return
	}

	var s PMI
	if err = json.Unmarshal(body, &s); err != nil {
		logger.Error(err.Error())
		go poster.Poster(ic)
		return
	}
	if len(s.Result.Data) == 0 {
		go poster.Poster(ic)
		return
	}

	for _, v := range s.Result.Data {
		buffer := &model.Buffer{}
		t, err := time.Parse("2006-01-02 03:04:05", v.ReportDate)
		if err != nil {
			logger.Error(err.Error())
			go poster.Poster(ic)
			break
		}
		buffer.Date = fmt.Sprintf("%d%02d", t.Year(), t.Month())
		switch ic.TargetCode {
		case "HG00020": // 制造业采购经理指数
			buffer.TargetValue = arithmetic.TrimZero(fmt.Sprintf("%.2f", v.MakeIndex))
		}
		buffers = append(buffers, buffer)
	}

	return
}

// 工业品出厂价格指数
func eastMoneyEconomicPPI(ic *model.IndexConfig) (buffers []*model.Buffer) {
	body, err := downloader.SimpleGet(api.PPI)

	if err != nil {
		if !safeguard.IsNetworkNormal() {
			logger.Error("请检查服务器的网络是否能联通外网")
			return
		}
		logger.Error(err.Error())
		go poster.Poster(ic)
		return
	}

	var s PPI
	if err = json.Unmarshal(body, &s); err != nil {
		logger.Error(err.Error())
		go poster.Poster(ic)
		return
	}
	if len(s.Result.Data) == 0 {
		go poster.Poster(ic)
		return
	}

	for _, v := range s.Result.Data {
		buffer := &model.Buffer{}
		t, err := time.Parse("2006-01-02 03:04:05", v.ReportDate)
		if err != nil {
			logger.Error(err.Error())
			go poster.Poster(ic)
			break
		}
		buffer.Date = fmt.Sprintf("%d%02d", t.Year(), t.Month())
		switch ic.TargetCode {
		case "HG00023": // 工业品出厂价格指数当月
			buffer.TargetValue = arithmetic.TrimZero(fmt.Sprintf("%.2f", v.Base))
		case "HG00089": // PPI同比增速月度
			buffer.TargetValue = arithmetic.TrimZero(fmt.Sprintf("%.2f", v.BaseSame))
		}
		buffers = append(buffers, buffer)
	}

	return
}

// 存款准备金率
func eastMoneyEconomicZBJ(ic *model.IndexConfig) (buffers []*model.Buffer) {
	body, err := downloader.SimpleGet(api.ZBJ)

	if err != nil {
		if !safeguard.IsNetworkNormal() {
			logger.Error("请检查服务器的网络是否能联通外网")
			return
		}
		logger.Error(err.Error())
		go poster.Poster(ic)
		return
	}

	var s ZBJ
	if err = json.Unmarshal(body, &s); err != nil {
		logger.Error(err.Error())
		go poster.Poster(ic)
		return
	}
	if len(s.Result.Data) == 0 {
		go poster.Poster(ic)
		return
	}

	for _, v := range s.Result.Data {
		trade := regexp.MustCompile("[0-9]+").FindAllString(v.TradeDate, -1)
		if len(trade) != 3 {
			logger.Error("字段TradeDate对应的生效时间不符合以下形式: 2022年12月05日")
			go poster.Poster(ic)
			break
		}

		buffer := &model.Buffer{}
		buffer.Date = trade[0] + trade[1] + trade[2]

		switch ic.TargetCode {
		case "HG00066": // 存款准备金率
			buffer.TargetValue = arithmetic.TrimZero(fmt.Sprintf("%.2f", v.InterestRateBA))
		}
		buffers = append(buffers, buffer)
	}
	return
}

// 海关进出口
func eastMoneyEconomicJCK(ic *model.IndexConfig) (buffers []*model.Buffer) {
	body, err := downloader.SimpleGet(api.JCK)

	if err != nil {
		if !safeguard.IsNetworkNormal() {
			logger.Error("请检查服务器的网络是否能联通外网")
			return
		}
		logger.Error(err.Error())
		go poster.Poster(ic)
		return
	}

	var s JCK
	if err = json.Unmarshal(body, &s); err != nil {
		logger.Error(err.Error())
		go poster.Poster(ic)
		return
	}
	if len(s.Result.Data) == 0 {
		go poster.Poster(ic)
		return
	}

	for _, v := range s.Result.Data {
		buffer := &model.Buffer{}
		t, err := time.Parse("2006-01-02 03:04:05", v.ReportDate)
		if err != nil {
			logger.Error(err.Error())
			go poster.Poster(ic)
			break
		}
		buffer.Date = fmt.Sprintf("%d%02d", t.Year(), t.Month())
		switch ic.TargetCode {
		case "HG00065": // 出口当月同比增速
			buffer.TargetValue = arithmetic.TrimZero(fmt.Sprintf("%.2f", v.ExitBaseSame))
		}
		buffers = append(buffers, buffer)
	}

	return
}

// 利率调整
func eastMoneyEconomicLL(ic *model.IndexConfig) (buffers []*model.Buffer) {
	body, err := downloader.SimpleGet(api.LL)

	if err != nil {
		if !safeguard.IsNetworkNormal() {
			logger.Error("请检查服务器的网络是否能联通外网")
			return
		}
		logger.Error(err.Error())
		go poster.Poster(ic)
		return
	}

	var s LL
	if err = json.Unmarshal(body, &s); err != nil {
		logger.Error(err.Error())
		go poster.Poster(ic)
		return
	}
	if len(s.Result.Data) == 0 {
		go poster.Poster(ic)
		return
	}

	for _, v := range s.Result.Data {
		buffer := &model.Buffer{}
		t, err := time.Parse("2006-01-02 03:04:05", v.ReportDate)
		if err != nil {
			logger.Error(err.Error())
			go poster.Poster(ic)
			break
		}
		buffer.Date = fmt.Sprintf("%d%02d%02d", t.Year(), t.Month(), t.Day())
		switch ic.TargetCode {
		case "HY00007": // 贷款基准利率
			buffer.TargetValue = arithmetic.TrimZero(fmt.Sprintf("%.2f", v.LoadRateSA))
		case "HY00011": // 存款基准利率
			buffer.TargetValue = arithmetic.TrimZero(fmt.Sprintf("%.2f", v.DepositRateBA))
		}
		buffers = append(buffers, buffer)
	}

	return
}

// 外汇和黄金储备
func eastMoneyEconomicWH(ic *model.IndexConfig) (buffers []*model.Buffer) {
	body, err := downloader.SimpleGet(api.WH)

	if err != nil {
		if !safeguard.IsNetworkNormal() {
			logger.Error("请检查服务器的网络是否能联通外网")
			return
		}
		logger.Error(err.Error())
		go poster.Poster(ic)
		return
	}

	var s WH
	if err = json.Unmarshal(body, &s); err != nil {
		logger.Error(err.Error())
		go poster.Poster(ic)
		return
	}
	if len(s.Result.Data) == 0 {
		go poster.Poster(ic)
		return
	}

	for _, v := range s.Result.Data {
		buffer := &model.Buffer{}
		t, err := time.Parse("2006-01-02 03:04:05", v.ReportDate)
		if err != nil {
			logger.Error(err.Error())
			go poster.Poster(ic)
			break
		}
		buffer.Date = fmt.Sprintf("%d%02d", t.Year(), t.Month())
		switch ic.TargetCode {
		case "HG00090": // 外汇储备
			buffer.TargetValue = arithmetic.TrimZero(fmt.Sprintf("%.2f", v.Forex))
		case "HG00091": // 外汇储备同比增速
			buffer.TargetValue = arithmetic.TrimZero(fmt.Sprintf("%.2f", v.ForexSame))
		}
		buffers = append(buffers, buffer)
	}

	return
}

// 中美国债收益率
func eastMoneyEconomicNationalDebt(ic *model.IndexConfig) (buffers []*model.Buffer) {
	body, err := downloader.SimpleGet(api.NationalDebt)

	if err != nil {
		if !safeguard.IsNetworkNormal() {
			logger.Error("请检查服务器的网络是否能联通外网")
			return
		}
		logger.Error(err.Error())
		go poster.Poster(ic)
		return
	}

	var s NationalDebt
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
		t, err := time.Parse("2006-01-02 03:04:05", v.SolarDate)
		if err != nil {
			logger.Error(err.Error())
			go poster.Poster(ic)
			break
		}
		bufferTF := model.BufferTF{}
		bufferTF.Date = t
		switch ic.TargetCode {
		case "HG00062": // 中国：国债收益率 - 10年
			bufferTF.TargetValue = v.EMM00166466
		}
		bufferTFList = append(bufferTFList, bufferTF)
	}
	sort.Sort(bufferTFList)
	buffers = expansion(bufferTFList)

	return
}
