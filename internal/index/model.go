package index

import (
	"reflect"
	"strings"
)

type Field struct {
	ValueGUID        string `db:"VALUE_GUID" json:"ValueGUID"`                // 指标值ID
	TargetGUID       string `db:"TARGET_GUID" json:"TargetGUID"`              // 指标ID
	TargetCode       string `db:"TARGET_CODE" json:"TargetCode"`              // 指标代码
	TargetName       string `db:"TARGET_NAME" json:"TargetName"`              // 指标名称
	TargetNameEN     string `db:"TARGET_NAME_EN" json:"TargetNameEN"`         // 指标英文名称
	DataSourceCode   string `db:"DATA_SOURCE_CODE" json:"DataSourceCode"`     // 数据来源代码
	DataSourceName   string `db:"DATA_SOURCE_NAME" json:"DataSourceName"`     // 数据来源名称
	SourceTargetCode string `db:"SOURCE_TARGET_CODE" json:"SourceTargetCode"` // 来源系统指标代码
	RegionCode       string `db:"REGION_CODE" json:"RegionCode"`              // 统计地区
	RegionName       string `db:"REGION_NAME" json:"RegionName"`              // 统计地区说明
	IsQuantity       string `db:"IS_QUANTITY" json:"IsQuantity"`              // 是否定量
	UnitType         string `db:"UNIT_TYPE" json:"UnitType"`                  // 计量单位类型
	UnitName         string `db:"UNIT_NAME" json:"UnitName"`                  // 计量单位名称
	PeriodType       string `db:"PERIOD_TYPE" json:"PeriodType"`              // 计量单位类型
	PeriodName       string `db:"PERIOD_NAME" json:"PeriodName"`              // 计量单位名称
	AcctYear         string `db:"ACCT_YEAR" json:"AcctYear"`                  // 年
	AcctSeason       string `db:"ACCT_QUARTOR" json:"AcctSeason"`             // 季
	AcctMonth        string `db:"ACCT_MONTH" json:"AcctMonth"`                // 月
	AcctDate         string `db:"ACCT_DATE" json:"AcctDate"`                  // 日
	TargetValue      string `db:"TARGET_VALUE" json:"TargetValue"`            // 指标值
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
