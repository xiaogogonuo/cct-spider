package indicator

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/indicator/code/date_type_code"
	"github.com/xiaogogonuo/cct-spider/internal/indicator/code/indicator_code"
	"github.com/xiaogogonuo/cct-spider/internal/indicator/code/province_code"
	"github.com/xiaogogonuo/cct-spider/internal/indicator/pkg/date"
	"github.com/xiaogogonuo/cct-spider/internal/indicator/pkg/executor"
	"github.com/xiaogogonuo/cct-spider/internal/indicator/pkg/urllib"
	"github.com/xiaogogonuo/cct-spider/pkg/db/mysql"
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
	"strings"
)

var (
	table = "T_DMAA_BASE_TARGET_VALUE"
)

const (
	SplitYearBlock   = 3
	SplitSeasonBlock = 4
	SplitMonthBlock  = 12
)

var (
	MarcoYearSQL = "SELECT ACCT_YEAR, TARGET_VALUE FROM %s WHERE " +
		"VALUE_GUID = MD5(CONCAT(TARGET_CODE, ACCT_YEAR)) AND " +
		"TARGET_CODE = '%s'"
	MarcoSeasonSQL = "SELECT CONCAT(ACCT_YEAR, ACCT_QUARTOR), TARGET_VALUE FROM %s WHERE " +
		"VALUE_GUID = MD5(CONCAT(TARGET_CODE, ACCT_YEAR, ACCT_QUARTOR)) AND " +
		"TARGET_CODE = '%s'"
	MarcoMonthSQL = "SELECT CONCAT(ACCT_YEAR, ACCT_MONTH), TARGET_VALUE FROM %s WHERE " +
		"VALUE_GUID = MD5(CONCAT(TARGET_CODE, ACCT_YEAR, ACCT_MONTH)) AND " +
		"TARGET_CODE = '%s'"
	ProvinceYearSQL = "SELECT ACCT_YEAR, TARGET_VALUE FROM %s WHERE " +
		"VALUE_GUID = MD5(CONCAT(TARGET_CODE, ACCT_YEAR, REGION_CODE)) AND " +
		"TARGET_CODE = '%s' AND " +
		"REGION_CODE = '%s'"
	ProvinceSeasonSQL = "SELECT CONCAT(ACCT_YEAR, ACCT_QUARTOR), TARGET_VALUE FROM %s WHERE " +
		"VALUE_GUID = MD5(CONCAT(TARGET_CODE, ACCT_YEAR, ACCT_QUARTOR, REGION_CODE)) AND " +
		"TARGET_CODE = '%s' AND " +
		"REGION_CODE = '%s'"
	ProvinceMonthSQL = "SELECT CONCAT(ACCT_YEAR, ACCT_MONTH), TARGET_VALUE FROM %s WHERE " +
		"VALUE_GUID = MD5(CONCAT(TARGET_CODE, ACCT_YEAR, ACCT_MONTH, REGION_CODE)) AND " +
		"TARGET_CODE = '%s' AND " +
		"REGION_CODE = '%s'"
)

type Indicator struct {
	dateType      string
	regionCode    string
	indicatorName string
	indicatorInfo map[string]string
}

func (i Indicator) historyDate() (history map[string]struct{}, err error) {
	switch i.dateType {
	case date_type_code.MarcoAnnualDataCode, date_type_code.ProvinceAnnualDataCode:
		history, err = date.StartToCurrentYear(i.indicatorInfo)
	case date_type_code.MarcoQuarterlyDataCode, date_type_code.ProvinceQuarterlyDataCode:
		history, err = date.StartToCurrentSeason(i.indicatorInfo)
	case date_type_code.MarcoMonthlyDataCode, date_type_code.ProvinceMonthlyDataCode:
		history, err = date.StartToCurrentMonth(i.indicatorInfo)
	}
	return
}

func (i Indicator) queryRow() (row [][]string) {
	regionCode := i.regionCode
	targetCode := i.indicatorInfo["TargetCode"]
	switch i.dateType {
	case date_type_code.MarcoAnnualDataCode:
		row = mysql.Query(fmt.Sprintf(MarcoYearSQL, table, targetCode))
	case date_type_code.MarcoQuarterlyDataCode:
		row = mysql.Query(fmt.Sprintf(MarcoSeasonSQL, table, targetCode))
	case date_type_code.MarcoMonthlyDataCode:
		row = mysql.Query(fmt.Sprintf(MarcoMonthSQL, table, targetCode))
	case date_type_code.ProvinceAnnualDataCode:
		row = mysql.Query(fmt.Sprintf(ProvinceYearSQL, table, targetCode, regionCode))
	case date_type_code.ProvinceQuarterlyDataCode:
		row = mysql.Query(fmt.Sprintf(ProvinceSeasonSQL, table, targetCode, regionCode))
	case date_type_code.ProvinceMonthlyDataCode:
		row = mysql.Query(fmt.Sprintf(ProvinceMonthSQL, table, targetCode, regionCode))
	}
	return
}

func (i Indicator) spiltDiff(diff []string) (diffBlock [][]string) {
	switch i.dateType {
	case date_type_code.MarcoAnnualDataCode, date_type_code.ProvinceAnnualDataCode:
		diffBlock = date.SplitDiff(diff, SplitYearBlock)
	case date_type_code.MarcoQuarterlyDataCode, date_type_code.ProvinceQuarterlyDataCode:
		diffBlock = date.SplitDiff(diff, SplitSeasonBlock)
	case date_type_code.MarcoMonthlyDataCode, date_type_code.ProvinceMonthlyDataCode:
		diffBlock = date.SplitDiff(diff, SplitMonthBlock)
	}
	return
}

func (i Indicator) urlParam(dateRegion string) (param urllib.Param) {
	switch i.dateType {
	case date_type_code.MarcoAnnualDataCode:
		param = urllib.MacroYear(dateRegion)
	case date_type_code.MarcoQuarterlyDataCode:
		param = urllib.MacroQuarter(dateRegion)
	case date_type_code.MarcoMonthlyDataCode:
		param = urllib.MacroMonth(dateRegion)
	case date_type_code.ProvinceAnnualDataCode:
		param = urllib.ProvinceYear(dateRegion, i.regionCode)
	case date_type_code.ProvinceQuarterlyDataCode:
		param = urllib.ProvinceQuarter(dateRegion, i.regionCode)
	case date_type_code.ProvinceMonthlyDataCode:
		param = urllib.ProvinceMonth(dateRegion, i.regionCode)
	}
	return
}

func (i Indicator) scheduler() {
	indicatorData := make([][]string, 0)
	history, err := i.historyDate()
	if err != nil {
		return
	}
	row := i.queryRow()
	diff := date.Diff(row, history)
	if len(diff) == 0 {
		return
	}
	for _, v := range i.spiltDiff(diff) {
		dateRegion := strings.Join(v, ",")
		response := executor.Execute(i.urlParam(dateRegion), i.dateType, i.indicatorInfo["SourceTargetCode"])
		indicatorData = append(indicatorData, response...)
	}
	constructor := Constructor{
		RegionCode: i.regionCode,
		IndicatorName: i.indicatorName,
		IndicatorInfo: i.indicatorInfo,
		IndicatorData: indicatorData,
	}
	if len(constructor.IndicatorData) == 0 {
		return
	}
	logger.Info(i.indicatorName, logger.Field("crawl rows: ", len(constructor.IndicatorData)))
	Dump(constructor.Construct())
	return
}

func (i Indicator) marco(dateType string) {
	i.dateType = dateType
	i.scheduler()
}

func (i Indicator) province(dateType, regionCode string) {
	i.dateType = dateType
	i.regionCode = regionCode
	i.scheduler()
}

func Start() {
	for indicatorName, indicatorInfo := range indicator_code.IndexMap {
		logger.Info(indicatorName)
		dateType := indicatorInfo["DateType"]
		areaType := indicatorInfo["AreaType"]
		indicator := Indicator{
			indicatorName: indicatorName,
			indicatorInfo: indicatorInfo,
		}
		if areaType == indicator_code.Marco && dateType == indicator_code.Annual {
			indicator.marco(date_type_code.MarcoAnnualDataCode)
		}
		if areaType == indicator_code.Marco && dateType == indicator_code.Quarterly {
			indicator.marco(date_type_code.MarcoQuarterlyDataCode)
		}
		if areaType == indicator_code.Marco && dateType == indicator_code.Monthly {
			indicator.marco(date_type_code.MarcoMonthlyDataCode)
		}
		if areaType == indicator_code.Marco && dateType == indicator_code.AnnualQuarterly {
			indicator.marco(date_type_code.MarcoAnnualDataCode)
			indicator.marco(date_type_code.MarcoQuarterlyDataCode)
		}
		if areaType == indicator_code.Marco && dateType == indicator_code.AnnualMonthly {
			indicator.marco(date_type_code.MarcoAnnualDataCode)
			indicator.marco(date_type_code.MarcoMonthlyDataCode)
		}
		if areaType == indicator_code.Marco && dateType == indicator_code.QuarterlyMonthly {
			indicator.marco(date_type_code.MarcoQuarterlyDataCode)
			indicator.marco(date_type_code.MarcoMonthlyDataCode)
		}
		if areaType == indicator_code.Marco && dateType == indicator_code.AnnualQuarterlyMonthly {
			indicator.marco(date_type_code.MarcoAnnualDataCode)
			indicator.marco(date_type_code.MarcoQuarterlyDataCode)
			indicator.marco(date_type_code.MarcoMonthlyDataCode)
		}
		if areaType == indicator_code.Province && dateType == indicator_code.Annual {
			for code := range province_code.CodeProvince {
				indicator.province(date_type_code.ProvinceAnnualDataCode, code)
			}
		}
		if areaType == indicator_code.Province && dateType == indicator_code.Quarterly {
			for code := range province_code.CodeProvince {
				indicator.province(date_type_code.ProvinceQuarterlyDataCode, code)
			}
		}
		if areaType == indicator_code.Province && dateType == indicator_code.Monthly {
			for code := range province_code.CodeProvince {
				indicator.province(date_type_code.ProvinceMonthlyDataCode, code)
			}
		}
		if areaType == indicator_code.Province && dateType == indicator_code.AnnualQuarterly {
			for code := range province_code.CodeProvince {
				indicator.province(date_type_code.ProvinceAnnualDataCode, code)
				indicator.province(date_type_code.ProvinceQuarterlyDataCode, code)
			}
		}
		if areaType == indicator_code.Province && dateType == indicator_code.AnnualMonthly {
			for code := range province_code.CodeProvince {
				indicator.province(date_type_code.ProvinceAnnualDataCode, code)
				indicator.province(date_type_code.ProvinceMonthlyDataCode, code)
			}
		}
		if areaType == indicator_code.Province && dateType == indicator_code.QuarterlyMonthly {
			for code := range province_code.CodeProvince {
				indicator.province(date_type_code.ProvinceQuarterlyDataCode, code)
				indicator.province(date_type_code.ProvinceMonthlyDataCode, code)
			}
		}
		if areaType == indicator_code.Province && dateType == indicator_code.AnnualQuarterlyMonthly {
			for code := range province_code.CodeProvince {
				indicator.province(date_type_code.ProvinceAnnualDataCode, code)
				indicator.province(date_type_code.ProvinceQuarterlyDataCode, code)
				indicator.province(date_type_code.ProvinceMonthlyDataCode, code)
			}
		}
		if areaType == indicator_code.MarcoProvince && dateType == indicator_code.Annual {
			indicator.marco(date_type_code.MarcoAnnualDataCode)
			for code := range province_code.CodeProvince {
				indicator.province(date_type_code.ProvinceAnnualDataCode, code)
			}
		}
		if areaType == indicator_code.MarcoProvince && dateType == indicator_code.Quarterly {
			indicator.marco(date_type_code.MarcoQuarterlyDataCode)
			for code := range province_code.CodeProvince {
				indicator.province(date_type_code.ProvinceQuarterlyDataCode, code)
			}
		}
		if areaType == indicator_code.MarcoProvince && dateType == indicator_code.Monthly {
			indicator.marco(date_type_code.MarcoMonthlyDataCode)
			for code := range province_code.CodeProvince {
				indicator.province(date_type_code.ProvinceMonthlyDataCode, code)
			}
		}
		if areaType == indicator_code.MarcoProvince && dateType == indicator_code.AnnualQuarterly {
			indicator.marco(date_type_code.MarcoAnnualDataCode)
			indicator.marco(date_type_code.MarcoQuarterlyDataCode)
			for code := range province_code.CodeProvince {
				indicator.province(date_type_code.ProvinceAnnualDataCode, code)
				indicator.province(date_type_code.ProvinceQuarterlyDataCode, code)
			}
		}
		if areaType == indicator_code.MarcoProvince && dateType == indicator_code.AnnualMonthly {
			indicator.marco(date_type_code.MarcoAnnualDataCode)
			indicator.marco(date_type_code.MarcoMonthlyDataCode)
			for code := range province_code.CodeProvince {
				indicator.province(date_type_code.ProvinceAnnualDataCode, code)
				indicator.province(date_type_code.ProvinceMonthlyDataCode, code)
			}
		}
		if areaType == indicator_code.MarcoProvince && dateType == indicator_code.QuarterlyMonthly {
			indicator.marco(date_type_code.MarcoQuarterlyDataCode)
			indicator.marco(date_type_code.MarcoMonthlyDataCode)
			for code := range province_code.CodeProvince {
				indicator.province(date_type_code.ProvinceQuarterlyDataCode, code)
				indicator.province(date_type_code.ProvinceMonthlyDataCode, code)
			}
		}
		if areaType == indicator_code.MarcoProvince && dateType == indicator_code.AnnualQuarterlyMonthly {
			indicator.marco(date_type_code.MarcoAnnualDataCode)
			indicator.marco(date_type_code.MarcoQuarterlyDataCode)
			indicator.marco(date_type_code.MarcoMonthlyDataCode)
			for code := range province_code.CodeProvince {
				indicator.province(date_type_code.ProvinceAnnualDataCode, code)
				indicator.province(date_type_code.ProvinceQuarterlyDataCode, code)
				indicator.province(date_type_code.ProvinceMonthlyDataCode, code)
			}
		}
	}
}