package target

import (
	"github.com/xiaogogonuo/cct-spider/internal/target/pkg/excel"
)

const (
	SheetName  = "Sheet1"
	ConfigFile = "诚通指标配置.xlsx"
)

// ExcelSetUp Excel初始化
func ExcelSetUp() []ExcelConfig {
	rows := excel.ReadExcel(ConfigFile, SheetName)
	var ecs []ExcelConfig
	for idx, row := range rows {
		if idx == 0 {
			continue
		}
		var ec ExcelConfig
		if row[0] == "1" {
			ec.Enable = true
		}
		ec.TargetCode = row[1]
		ec.TargetName = row[2]
		ec.TargetNameSpider = row[3]
		ec.TargetNameEn = row[4]
		ec.DataSourceCode = row[5]
		ec.DataSourceName = row[6]
		ec.SourceTargetCodeSpider = row[7]
		ec.SourceTargetCode = row[8]
		ec.UnitType = row[9]
		ec.UnitName = row[10]
		ec.PeriodType = row[11]
		ec.PeriodName = row[12]
		ec.SpiderTime = row[13]
		ec.Adapter = row[14]
		ecs = append(ecs, ec)
	}
	return ecs
}
