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
		ec.TargetCode = row[1]
		ec.TargetName = row[2]
		ec.TargetNameSpider = row[3]
		ec.TargetNameEn = row[4]
		ec.DataSourceCode = row[5]
		ec.DataSourceName = row[6]
		ec.SourceTargetCodeSpider = row[7]
		ec.SourceTargetCode = row[8]
		ec.IsQuantity = row[9]
		ec.UnitType = row[10]
		ec.UnitName = row[11]
		ec.PeriodType = row[12]
		ec.PeriodName = row[13]
		ec.SpiderTime = row[14]
		ec.Adapter = row[15]
		if row[16] == "是" {
			ec.Enable = true
		}
		ecs = append(ecs, ec)
	}
	return ecs
}
