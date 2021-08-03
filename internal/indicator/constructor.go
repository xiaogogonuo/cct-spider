package indicator

import (
	"github.com/xiaogogonuo/cct-spider/internal/indicator/code/indicator_code"
	"github.com/xiaogogonuo/cct-spider/internal/indicator/code/province_code"
	"github.com/xiaogogonuo/cct-spider/internal/indicator/pkg/response"
	"github.com/xiaogogonuo/cct-spider/pkg/encrypt/md5"
)

type Constructor struct {
	RegionCode    string            // 地区代码
	IndicatorName string            // 指标名称
	IndicatorInfo map[string]string // 指标详情
	IndicatorData [][]string        // 爬取数据
}

func acct(date, dateType string, targetValue *response.TargetValue) {
	switch dateType {
	case indicator_code.Annual:
		targetValue.AcctYear = date
	case indicator_code.Quarterly:
		targetValue.AcctYear = date[:4]
		targetValue.AcctSeason = date[4:6]
	case indicator_code.Monthly:
		targetValue.AcctYear = date[:4]
		targetValue.AcctMonth = date[4:6]
	case indicator_code.Daily:
		targetValue.AcctYear = date[:4]
		targetValue.AcctMonth = date[4:6]
		targetValue.AcctDate = date[6:8]
	}
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

func (c Constructor ) Construct() (data []response.TargetValue) {
	for _, indicator := range c.IndicatorData {
		tv := &response.TargetValue{}
		tv.ValueGUID = md5.MD5(c.IndicatorInfo["TargetCode"] + indicator[0] + c.RegionCode)
		tv.TargetGUID = md5.MD5(c.IndicatorInfo["TargetCode"])
		tv.TargetCode = c.IndicatorInfo["TargetCode"]
		tv.TargetName = c.IndicatorName
		tv.TargetNameEN = c.IndicatorInfo["TargetNameEN"]
		tv.DataSourceCode = c.IndicatorInfo["DataSourceCode"]
		tv.DataSourceName = c.IndicatorInfo["DataSourceName"]
		tv.SourceTargetCode = c.IndicatorInfo["SourceTargetCode"]
		tv.RegionCode = c.RegionCode
		tv.RegionName = province_code.CodeProvince[c.RegionCode]
		tv.IsQuantity = c.IndicatorInfo["IsQuantity"]
		tv.UnitType = c.IndicatorInfo["UnitType"]
		tv.UnitName = c.IndicatorInfo["UnitName"]
		tv.PeriodType = c.IndicatorInfo["PeriodType"]
		tv.PeriodName = c.IndicatorInfo["PeriodName"]
		tv.TargetValue = indicator[1]
		acct(indicator[0], c.IndicatorInfo["DateType"], tv)
		data = append(data, *tv)
	}
	return
}