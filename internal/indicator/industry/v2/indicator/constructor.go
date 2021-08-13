package indicator

import (
	"github.com/xiaogogonuo/cct-spider/internal/indicator/industry/v2/code"
	"github.com/xiaogogonuo/cct-spider/internal/indicator/industry/v2/pkg/response"
	"github.com/xiaogogonuo/cct-spider/pkg/encrypt/md5"
)

type constructor struct {
	indicatorName string
	indicatorInfo map[string]string
	respond       []response.Respond
}

func (c constructor) construct() (data []response.Field) {
	for _, res := range c.respond {
		f := &response.Field{}
		f.ValueGUID = md5.MD5(c.indicatorInfo["TargetCode"] + res.Date)
		f.TargetGUID = md5.MD5(c.indicatorInfo["TargetCode"])
		f.TargetCode = c.indicatorInfo["TargetCode"]
		f.TargetName = c.indicatorName
		f.TargetNameEN = c.indicatorInfo["TargetNameEN"]
		f.DataSourceCode = c.indicatorInfo["DataSourceCode"]
		f.DataSourceName = c.indicatorInfo["DataSourceName"]
		f.SourceTargetCode = c.indicatorInfo["SourceTargetCode"]
		f.IsQuantity = c.indicatorInfo["IsQuantity"]
		f.UnitType = c.indicatorInfo["UnitType"]
		f.UnitName = c.indicatorInfo["UnitName"]
		f.PeriodType = c.indicatorInfo["PeriodType"]
		f.PeriodName = c.indicatorInfo["PeriodName"]
		f.TargetValue = res.TargetValue
		acct(res.Date, f.PeriodType, f)
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
		f.AcctDate = date
	}
}