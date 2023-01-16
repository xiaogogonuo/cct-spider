package cct_index

import (
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/model"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/pkg/excel"
	"strings"
)

const (
	SheetName = "Sheet1"
	IndexFile = "诚通指标配置.xlsx"
)

// InitIndexConfigFromExcel 从Excel读取指标配置
func InitIndexConfigFromExcel() (ics []*model.IndexConfig, err error) {
	rows, err := excel.ReadFromExcel(IndexFile, SheetName)
	if err != nil {
		return
	}
	for idx, row := range rows {
		if idx == 0 {
			continue
		}
		ic := &model.IndexConfig{}
		if strings.Trim(row[0], " ") == "1" {
			ic.Enable = true
		}
		ic.TargetCode = row[1]
		ic.TargetName = row[2]
		ic.TargetNameSpider = row[3]
		ic.TargetNameEn = row[4]
		ic.DataSourceCode = row[5]
		ic.DataSourceName = row[6]
		ic.SourceTargetCodeSpider = row[7]
		ic.SourceTargetCode = row[8]
		ic.UnitType = row[9]
		ic.UnitName = row[10]
		ic.PeriodType = row[11]
		ic.PeriodName = row[12]
		ic.SpiderTime = row[13]
		ic.Adapter = row[14]
		ics = append(ics, ic)
	}
	return
}
