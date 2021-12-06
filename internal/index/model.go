package index

import (
	"reflect"
	"strings"
)

type Field struct {
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
	AcctYear         string `db:"ACCT_YEAR" json:"acctYear"`                  // 年
	AcctSeason       string `db:"ACCT_QUARTOR" json:"acctQuartor"`            // 季
	AcctMonth        string `db:"ACCT_MONTH" json:"acctMonth"`                // 月
	AcctDate         string `db:"ACCT_DATE" json:"acctDate"`                  // 日
	TargetValue      string `db:"TARGET_VALUE" json:"targetValue"`            // 指标值
}

func (f Field) GetNumFields() int {
	r := reflect.ValueOf(f)
	return r.NumField()
}

func (f Field) GetValues() (values []interface{}) {
	r := reflect.ValueOf(f)
	for i := 0; i < r.NumField(); i++ {
		values = append(values, r.Field(i).Interface())
	}
	return
}

func (f Field) GetFields() (fields []string) {
	r := reflect.TypeOf(f)
	for i := 0; i < r.NumField(); i++ {
		field := r.Field(i).Tag.Get("db")
		fields = append(fields, field)
	}
	return
}

var FieldNum int
var FieldStr string
var FieldSVS string

func init() {
	f := Field{}
	FieldNum = f.GetNumFields()
	fields := f.GetFields()
	FieldStr = strings.Join(fields, ",")
	var container []string
	for _, filed := range fields {
		container = append(container, filed+" = VALUES("+filed+")")
	}
	FieldSVS = strings.Join(container, ",")
}
