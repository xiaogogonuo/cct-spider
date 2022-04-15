package model

import "github.com/xiaogogonuo/cct-spider/internal/target/constant"

// DataBase 保存记录到数据库的字段
type DataBase struct {
	ValueGUID        string `db:"VALUE_GUID" json:"valueGuid"`                // 指标值ID
	TargetGUID       string `db:"TARGET_GUID" json:"targetGuid"`              // 指标ID
	TargetCode       string `db:"TARGET_CODE" json:"targetCode"`              // 指标代码
	TargetName       string `db:"TARGET_NAME" json:"targetName"`              // 指标名称
	TargetNameEN     string `db:"TARGET_NAME_EN" json:"targetNameEn"`         // 指标英文名称
	DataSourceCode   string `db:"DATA_SOURCE_CODE" json:"dataSourceCode"`     // 数据来源代码
	DataSourceName   string `db:"DATA_SOURCE_NAME" json:"dataSourceName"`     // 数据来源名称
	SourceTargetCode string `db:"SOURCE_TARGET_CODE" json:"sourceTargetCode"` // 来源系统指标代码
	RegionCode       string `db:"REGION_CODE" json:"regionCode"`              // 统计地区
	RegionName       string `db:"REGION_NAME" json:"regionName"`              // 统计地区说明
	IsQuantity       string `db:"IS_QUANTITY" json:"isQuantity"`              // 是否定量
	UnitType         string `db:"UNIT_TYPE" json:"unitType"`                  // 计量单位类型
	UnitName         string `db:"UNIT_NAME" json:"unitName"`                  // 计量单位名称
	PeriodType       string `db:"PERIOD_TYPE" json:"periodType"`              // 计量单位类型
	PeriodName       string `db:"PERIOD_NAME" json:"periodName"`              // 计量单位名称
	AcctYear         string `db:"ACCT_YEAR" json:"acctYear"`                  // 年：2022
	AcctSeason       string `db:"ACCT_QUARTOR" json:"acctQuartor"`            // 季：Q1、Q2、Q3、Q4
	AcctMonth        string `db:"ACCT_MONTH" json:"acctMonth"`                // 月：01、02、...、11、12
	AcctDate         string `db:"ACCT_DATE" json:"acctDate"`                  // 日：20220414、20220415
	TargetValue      string `db:"TARGET_VALUE" json:"targetValue"`            // 指标值
}

// UpdateAcct 更新日期
func (d *DataBase) UpdateAcct(date, periodType string) {
	switch periodType {
	case constant.PeriodTypeYear:
		d.AcctYear = date
	case constant.PeriodTypeSeason:
		d.AcctYear = date[:4]
		d.AcctSeason = date[4:6]
	case constant.PeriodTypeMonth:
		d.AcctYear = date[:4]
		d.AcctMonth = date[4:6]
	case constant.PeriodTypeDay:
		d.AcctYear = date[:4]
		d.AcctMonth = date[4:6]
		d.AcctDate = date
	}
}
