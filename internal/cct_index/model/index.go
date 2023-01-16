package model

import "github.com/xiaogogonuo/cct-spider/internal/cct_index/constant"

// Index 保存记录到数据库的字段
type Index struct {
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
	// 指标值
	// 对实时数据指标，发送给Java服务器的指标值格式如下：
	// 最新价,涨跌,涨跌幅,最高,最低,昨收,更新时间
	// 1.3543,-0.0045,-0.33%,1.3604,1.3533,1.3588,2022-04-15 23:28:50
	TargetValue string `db:"TARGET_VALUE" json:"targetValue"`
}

// AcctSetter 针对非实时指标
// 传入的date应该是如下的形式之一：20230101、2023Q1、202306、2023
func (i *Index) AcctSetter(date string) {
	switch i.PeriodType {
	case constant.PeriodTypeYear:
		i.AcctYear = date[:4]
	case constant.PeriodTypeSeason:
		i.AcctYear = date[:4]
		i.AcctSeason = date[4:6]
	case constant.PeriodTypeMonth:
		i.AcctYear = date[:4]
		i.AcctMonth = date[4:6]
	case constant.PeriodTypeDay:
		i.AcctYear = date[:4]
		i.AcctMonth = date[4:6]
		i.AcctDate = date[:8]
	}
}
