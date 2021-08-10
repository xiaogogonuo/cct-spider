package indicator

import (
	"github.com/xiaogogonuo/cct-spider/internal/indicator/v2/code"
	"github.com/xiaogogonuo/cct-spider/internal/indicator/v2/pkg/response"
	"github.com/xiaogogonuo/cct-spider/pkg/encrypt/md5"
)

type Constructor struct {
	Period        string
	IndicatorName string
	IndicatorInfo map[string]string
	Respond       []*response.Respond
}

// VALUE_GUID计算公式
// 全国：
// - 年度：md5(TARGET_CODE + ACCT_YEAR)
// - 季度：md5(TARGET_CODE + ACCT_YEAR + ACCT_QUARTOR)
// - 月度：md5(TARGET_CODE + ACCT_YEAR + ACCT_MONTH)
// 地区：
// - 年度：md5(TARGET_CODE + ACCT_YEAR + REGION_CODE)
// - 季度：md5(TARGET_CODE + ACCT_YEAR + ACCT_QUARTOR + REGION_CODE)
// - 月度：md5(TARGET_CODE + ACCT_YEAR + ACCT_MONTH + REGION_CODE)

func (c Constructor) Construct() (data []response.Field) {
	for _, res := range c.Respond {
		f := &response.Field{}
		f.ValueGUID = md5.MD5(c.IndicatorInfo["TargetCode"] + res.Date + res.ProvinceCode)
		f.TargetGUID = md5.MD5(c.IndicatorInfo["TargetCode"])
		f.TargetCode = c.IndicatorInfo["TargetCode"]
		f.TargetName = c.IndicatorName
		f.TargetNameEN = c.IndicatorInfo["TargetNameEN"]
		f.DataSourceCode = c.IndicatorInfo["DataSourceCode"]
		f.DataSourceName = c.IndicatorInfo["DataSourceName"]
		f.SourceTargetCode = c.IndicatorInfo["SourceTargetCode"]
		f.RegionCode = res.ProvinceCode
		f.RegionName = code.CodeProvince[res.ProvinceCode]
		f.IsQuantity = c.IndicatorInfo["IsQuantity"]
		f.UnitType = c.IndicatorInfo["UnitType"]
		f.UnitName = c.IndicatorInfo["UnitName"]
		f.PeriodType = c.Period
		f.PeriodName = code.PeriodTypeName[c.Period]
		f.TargetValue = res.Data
		acct(res.Date, c.Period, f)
		data = append(data, *f)
	}
	return
}

func acct(date, period string, f *response.Field) {
	switch period {
	case code.PeriodTypeYear:
		f.AcctYear = date
	case code.PeriodTypeSeason, code.PeriodTypeMonth:
		f.AcctYear = date[:4]
		f.AcctSeason = date[4:6]
	case code.PeriodTypeDay:
		f.AcctYear = date[:4]
		f.AcctMonth = date[4:6]
		f.AcctDate = date[6:8]
	}
}