package response

import (
	"reflect"
	"strings"
)

type Response struct {
	ReturnData ReturnData `json:"returndata"`
}

type ReturnData struct {
	DataNodes []Node `json:"datanodes"`
}

type Node struct {
	Data Data
	Code string `json:"code"`
}

type Data struct {
	Data    float64 `json:"data"`
	StrData string  `json:"strdata"`
}

type TargetValue struct {
	ValueGUID        string `db:"VALUE_GUID"`         // 指标值ID
	TargetGUID       string `db:"TARGET_GUID"`        // 指标ID
	TargetCode       string `db:"TARGET_CODE"`        // 指标代码
	TargetName       string `db:"TARGET_NAME"`        // 指标名称
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

func (tv TargetValue) GetNumFields() int  {
	r := reflect.ValueOf(tv)
	return r.NumField()
}

func (tv TargetValue) GetValues() (values []interface{}) {
	r := reflect.ValueOf(tv)
	for i := 0; i < r.NumField(); i++ {
		values = append(values, r.Field(i).Interface())
	}
	return
}

func (tv TargetValue) GetFields() (fields []string) {
	r := reflect.TypeOf(tv)
	for i := 0; i < r.NumField(); i++ {
		field := r.Field(i).Tag.Get("db")
		fields = append(fields, field)
	}
	return
}

var TargetValueFieldNum int
var TargetValueFieldStr string
var TargetValueFieldSVS string

func init() {
	tv := TargetValue{}
	TargetValueFieldNum = tv.GetNumFields()
	fields := tv.GetFields()
	TargetValueFieldStr = strings.Join(fields, ",")
	var container []string
	for _, filed := range fields {
		container = append(container, filed + " = VALUES(" + filed + ")")
	}
	TargetValueFieldSVS = strings.Join(container, ",")
}