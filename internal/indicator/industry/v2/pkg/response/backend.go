package response

import (
    "reflect"
    "strings"
)

type Field struct {
	ValueGUID        string `db:"VALUE_GUID"`         // 指标值ID
	TargetGUID       string `db:"TARGET_GUID"`        // 指标ID
	TargetCode       string `db:"TARGET_CODE"`        // 指标代码
	TargetName       string `db:"TARGET_NAME"`        // 指标名称
	TargetNameEN     string `db:"TARGET_NAME_EN"`     // 指标英文名称
	DataSourceCode   string `db:"DATA_SOURCE_CODE"`   // 数据来源代码
	DataSourceName   string `db:"DATA_SOURCE_NAME"`   // 数据来源名称
	SourceTargetCode string `db:"SOURCE_TARGET_CODE"` // 来源系统指标代码
	RegionCode       string `db:"REGION_CODE"`        // 统计地区
	RegionName       string `db:"REGION_NAME"`        // 统计地区说明
	IsQuantity       string `db:"IS_QUANTITY"`        // 是否定量
	UnitType         string `db:"UNIT_TYPE"`          // 计量单位类型
	UnitName         string `db:"UNIT_NAME"`          // 计量单位名称
	PeriodType       string `db:"PERIOD_TYPE"`        // 计量单位类型
	PeriodName       string `db:"PERIOD_NAME"`        // 计量单位名称
	AcctYear         string `db:"ACCT_YEAR"`          // 年
	AcctSeason       string `db:"ACCT_QUARTOR"`       // 季
	AcctMonth        string `db:"ACCT_MONTH"`         // 月
	AcctDate         string `db:"ACCT_DATE"`          // 日
	TargetValue      string `db:"TARGET_VALUE"`       // 指标值
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
		container = append(container, filed + " = VALUES(" + filed + ")")
	}
	FieldSVS = strings.Join(container, ",")
}
