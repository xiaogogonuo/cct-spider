package eastmoney

import (
	"encoding/json"
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/api"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/model"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/pkg/downloader"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/poster"
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
	"time"
)

// SpiderEastMoneyEconomicTarget 爬取"东方财富"网站的经济数据
func SpiderEastMoneyEconomicTarget(targetCode string) (buffers []*model.Buffer) {
	switch targetCode {
	case "HG00001":
		buffers = eastMoneyEconomicTargetGDPTB()
	case "HG00016", "HG00017":
		buffers = eastMoneyEconomicTargetIAV(targetCode)
	case "HG00027", "HG00028", "HG00029", "HG00030":
		buffers = eastMoneyEconomicTargetXFP(targetCode)
	}
	return
}

// 国内生产总值同比增长
func eastMoneyEconomicTargetGDPTB() (buffers []*model.Buffer) {
	body, err := downloader.SimpleGet(api.GDPTb)
	var gdp GDP
	if err = json.Unmarshal(body, &gdp); err != nil {
		logger.Error(err.Error())
		go poster.GDPTBPoster()
		return
	}
	if len(gdp.Result.Data) == 0 {
		go poster.GDPTBPoster()
		return
	}

	for _, v := range gdp.Result.Data {
		buffer := &model.Buffer{}
		t, err := time.Parse("2006-01-02 03:04:05", v.ReportDate)
		if err != nil {
			logger.Error(err.Error())
			go poster.GDPTBPoster()
			break
		}
		buffer.Date = fmt.Sprintf("%d", t.Year())
		switch int(t.Month()) {
		case 3:
			buffer.Date += "Q1"
		case 6:
			buffer.Date += "Q2"
		case 9:
			buffer.Date += "Q3"
		case 12:
			buffer.Date += "Q4"
		}
		buffer.TargetValue = fmt.Sprintf("%.1f", v.SumSame)
		buffers = append(buffers, buffer)
	}
	return
}

// 工业增加值
func eastMoneyEconomicTargetIAV(targetCode string) (buffers []*model.Buffer) {
	body, err := downloader.SimpleGet(api.IAV)
	var s IAV
	if err = json.Unmarshal(body, &s); err != nil {
		logger.Error(err.Error())
		go poster.IAVPoster()
		return
	}
	if len(s.Result.Data) == 0 {
		go poster.IAVPoster()
		return
	}

	for _, v := range s.Result.Data {
		buffer := &model.Buffer{}
		t, err := time.Parse("2006-01-02 03:04:05", v.ReportDate)
		if err != nil {
			logger.Error(err.Error())
			go poster.IAVPoster()
			break
		}
		buffer.Date = fmt.Sprintf("%d%02d", t.Year(), t.Month())
		switch targetCode {
		case "HG00016": // 工业增加值同比增长
			buffer.TargetValue = fmt.Sprintf("%.1f", v.BaseSame)
		case "HG00017": // 工业增加值累计增长
			buffer.TargetValue = fmt.Sprintf("%.1f", v.BaseAccumulate)
		}
		buffers = append(buffers, buffer)
	}
	return
}

// 社会消费品零售总额
func eastMoneyEconomicTargetXFP(targetCode string) (buffers []*model.Buffer) {
	body, err := downloader.SimpleGet(api.XFP)
	var s XFP
	if err = json.Unmarshal(body, &s); err != nil {
		logger.Error(err.Error())
		go poster.XFPPoster()
		return
	}
	if len(s.Result.Data) == 0 {
		go poster.XFPPoster()
		return
	}

	for _, v := range s.Result.Data {
		buffer := &model.Buffer{}
		t, err := time.Parse("2006-01-02 03:04:05", v.ReportDate)
		if err != nil {
			logger.Error(err.Error())
			go poster.XFPPoster()
			break
		}
		buffer.Date = fmt.Sprintf("%d%02d", t.Year(), t.Month())
		switch targetCode {
		case "HG00027": // 社会消费品零售总额当期值
			buffer.TargetValue = fmt.Sprintf("%.1f", v.RetailTotal)
		case "HG00028": // 社会消费品零售总额累计值
			buffer.TargetValue = fmt.Sprintf("%.1f", v.RetailTotalAccumulate)
		case "HG00029": // 社会消费品零售总额同比增长
			buffer.TargetValue = fmt.Sprintf("%.1f", v.RetailTotalSame)
		case "HG00030": // 社会消费品零售总额累计增长
			buffer.TargetValue = fmt.Sprintf("%.1f", v.RetailAccumulateSame)
		}
		buffers = append(buffers, buffer)
	}
	return
}

// 货币和准货币(M2)供应量
func eastMoneyEconomicTargetM2(targetCode string) (buffers []*model.Buffer) {
	body, err := downloader.SimpleGet(api.M2)
	var s M2
	if err = json.Unmarshal(body, &s); err != nil {
		logger.Error(err.Error())
		go poster.M2Poster()
		return
	}
	if len(s.Result.Data) == 0 {
		go poster.M2Poster()
		return
	}

	for _, v := range s.Result.Data {
		buffer := &model.Buffer{}
		t, err := time.Parse("2006-01-02 03:04:05", v.ReportDate)
		if err != nil {
			logger.Error(err.Error())
			go poster.M2Poster()
			break
		}
		buffer.Date = fmt.Sprintf("%d%02d", t.Year(), t.Month())
		switch targetCode {
		case "HG00006": // 货币和准货币(M2)供应量期末值
			buffer.TargetValue = fmt.Sprintf("%.1f", v.BasicCurrency)
		case "HG00007": // 货币和准货币(M2)供应量同比增长
			buffer.TargetValue = fmt.Sprintf("%.1f", v.BasicCurrencySame)
		}
		buffers = append(buffers, buffer)
	}
	return
}
