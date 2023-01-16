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
	case "HG00006", "HG00007":
		buffers = eastMoneyEconomicTargetM2(targetCode)
	case "HG00088":
		buffers = eastMoneyEconomicTargetCPI(targetCode)
	case "HG00020":
		buffers = eastMoneyEconomicTargetPMI(targetCode)
	case "HG00023", "HG00089":
		buffers = eastMoneyEconomicTargetPPI(targetCode)
	case "HG00066":
		buffers = eastMoneyEconomicTargetZBJ(targetCode)
	case "HG00065":
		buffers = eastMoneyEconomicTargetJCK(targetCode)
	case "HY00007", "HY00011":
		buffers = eastMoneyEconomicTargetLL(targetCode)
	case "HG00090", "HG00091":
		buffers = eastMoneyEconomicTargetWH(targetCode)
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

// TODO: 确认HG00088对应的TargetValue是否正确
// 居民消费价格指数
func eastMoneyEconomicTargetCPI(targetCode string) (buffers []*model.Buffer) {
	body, err := downloader.SimpleGet(api.CPI)
	var s CPI
	if err = json.Unmarshal(body, &s); err != nil {
		logger.Error(err.Error())
		go poster.CPIPoster()
		return
	}
	if len(s.Result.Data) == 0 {
		go poster.CPIPoster()
		return
	}

	for _, v := range s.Result.Data {
		buffer := &model.Buffer{}
		t, err := time.Parse("2006-01-02 03:04:05", v.ReportDate)
		if err != nil {
			logger.Error(err.Error())
			go poster.CPIPoster()
			break
		}
		buffer.Date = fmt.Sprintf("%d%02d", t.Year(), t.Month())
		switch targetCode {
		case "HG00088": // CPI同比增速月度
			buffer.TargetValue = fmt.Sprintf("%.1f", v.NationalSame)
		}
		buffers = append(buffers, buffer)
	}
	return
}

// TODO: 确认HG00020对应的TargetValue是否正确
// 采购经理人指数
func eastMoneyEconomicTargetPMI(targetCode string) (buffers []*model.Buffer) {
	body, err := downloader.SimpleGet(api.PMI)
	var s PMI
	if err = json.Unmarshal(body, &s); err != nil {
		logger.Error(err.Error())
		go poster.PMIPoster()
		return
	}
	if len(s.Result.Data) == 0 {
		go poster.PMIPoster()
		return
	}

	for _, v := range s.Result.Data {
		buffer := &model.Buffer{}
		t, err := time.Parse("2006-01-02 03:04:05", v.ReportDate)
		if err != nil {
			logger.Error(err.Error())
			go poster.PMIPoster()
			break
		}
		buffer.Date = fmt.Sprintf("%d%02d", t.Year(), t.Month())
		switch targetCode {
		case "HG00020": // 制造业采购经理指数
			buffer.TargetValue = fmt.Sprintf("%.1f", v.MakeIndex)
		}
		buffers = append(buffers, buffer)
	}
	return
}

// 工业品出厂价格指数
func eastMoneyEconomicTargetPPI(targetCode string) (buffers []*model.Buffer) {
	body, err := downloader.SimpleGet(api.PPI)
	var s PPI
	if err = json.Unmarshal(body, &s); err != nil {
		logger.Error(err.Error())
		go poster.PPIPoster()
		return
	}
	if len(s.Result.Data) == 0 {
		go poster.PPIPoster()
		return
	}

	for _, v := range s.Result.Data {
		buffer := &model.Buffer{}
		t, err := time.Parse("2006-01-02 03:04:05", v.ReportDate)
		if err != nil {
			logger.Error(err.Error())
			go poster.PPIPoster()
			break
		}
		buffer.Date = fmt.Sprintf("%d%02d", t.Year(), t.Month())
		switch targetCode {
		case "HG00023": // 工业品出厂价格指数当月
			buffer.TargetValue = fmt.Sprintf("%.1f", v.Base)
		case "HG00089": // PPI同比增速月度
			buffer.TargetValue = fmt.Sprintf("%.1f", v.BaseSame)
		}
		buffers = append(buffers, buffer)
	}
	return
}

// TODO: 确认指标日前使用公布时间还是生效时间 单位是天
// TODO：确认指标值使用调整前还是调整后
// 存款准备金率
func eastMoneyEconomicTargetZBJ(targetCode string) (buffers []*model.Buffer) {
	body, err := downloader.SimpleGet(api.ZBJ)
	var s ZBJ
	if err = json.Unmarshal(body, &s); err != nil {
		logger.Error(err.Error())
		go poster.ZBJPoster()
		return
	}
	if len(s.Result.Data) == 0 {
		go poster.ZBJPoster()
		return
	}

	for _, v := range s.Result.Data {
		buffer := &model.Buffer{}
		t, err := time.Parse("2006-01-02 03:04:05", v.ReportDate)
		if err != nil {
			logger.Error(err.Error())
			go poster.ZBJPoster()
			break
		}
		buffer.Date = fmt.Sprintf("%d%02d%02d", t.Year(), t.Month(), t.Day())
		switch targetCode {
		case "HG00066": // 存款准备金率
			buffer.TargetValue = fmt.Sprintf("%.1f", v.InterestRateBA)
		}
		buffers = append(buffers, buffer)
	}
	return
}

// 海关进出口
func eastMoneyEconomicTargetJCK(targetCode string) (buffers []*model.Buffer) {
	body, err := downloader.SimpleGet(api.JCK)
	var s JCK
	if err = json.Unmarshal(body, &s); err != nil {
		logger.Error(err.Error())
		go poster.JCKPoster()
		return
	}
	if len(s.Result.Data) == 0 {
		go poster.JCKPoster()
		return
	}

	for _, v := range s.Result.Data {
		buffer := &model.Buffer{}
		t, err := time.Parse("2006-01-02 03:04:05", v.ReportDate)
		if err != nil {
			logger.Error(err.Error())
			go poster.JCKPoster()
			break
		}
		buffer.Date = fmt.Sprintf("%d%02d", t.Year(), t.Month())
		switch targetCode {
		case "HG00065": // 出口当月同比增速
			buffer.TargetValue = fmt.Sprintf("%.1f", v.ExitBaseSame)
		}
		buffers = append(buffers, buffer)
	}
	return
}

// TODO: 确认指标日前使用公布时间还是生效时间 单位是天
// TODO：确认指标值使用调整前还是调整后
// 利率调整
func eastMoneyEconomicTargetLL(targetCode string) (buffers []*model.Buffer) {
	body, err := downloader.SimpleGet(api.LL)
	var s LL
	if err = json.Unmarshal(body, &s); err != nil {
		logger.Error(err.Error())
		go poster.LLPoster()
		return
	}
	if len(s.Result.Data) == 0 {
		go poster.LLPoster()
		return
	}

	for _, v := range s.Result.Data {
		buffer := &model.Buffer{}
		t, err := time.Parse("2006-01-02 03:04:05", v.ReportDate)
		if err != nil {
			logger.Error(err.Error())
			go poster.LLPoster()
			break
		}
		buffer.Date = fmt.Sprintf("%d%02d%02d", t.Year(), t.Month(), t.Day())
		switch targetCode {
		case "HY00007": // 贷款基准利率
			buffer.TargetValue = fmt.Sprintf("%.1f", v.LoadRateSA)
		case "HY00011": // 存款基准利率
			buffer.TargetValue = fmt.Sprintf("%.1f", v.DepositRateBA)
		}
		buffers = append(buffers, buffer)
	}
	return
}

// 外汇和黄金储备
func eastMoneyEconomicTargetWH(targetCode string) (buffers []*model.Buffer) {
	body, err := downloader.SimpleGet(api.WH)
	var s WH
	if err = json.Unmarshal(body, &s); err != nil {
		logger.Error(err.Error())
		go poster.WHPoster()
		return
	}
	if len(s.Result.Data) == 0 {
		go poster.WHPoster()
		return
	}

	for _, v := range s.Result.Data {
		buffer := &model.Buffer{}
		t, err := time.Parse("2006-01-02 03:04:05", v.ReportDate)
		if err != nil {
			logger.Error(err.Error())
			go poster.WHPoster()
			break
		}
		buffer.Date = fmt.Sprintf("%d%02d", t.Year(), t.Month())
		switch targetCode {
		case "HG00090": // 外汇储备
			buffer.TargetValue = fmt.Sprintf("%.1f", v.Forex)
		case "HG00091": // 外汇储备同比增速
			buffer.TargetValue = fmt.Sprintf("%.1f", v.ForexSame)
		}
		buffers = append(buffers, buffer)
	}
	return
}
