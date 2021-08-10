package request

import (
	"fmt"
	"net/url"
)

const (
	UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.101 Safari/537.36"
)

type Param struct {
	M              string
	DBCode         string
	RowCode        string
	ColCode        string
	WdsWdCode      string
	WdsWdValueCode string
	DfWdsWdCode    string
	DfWdsValueCode string
}

func (p Param) ParamEncode() string {
	value := url.Values{}
	value.Set("m", p.M)
	value.Set("dbcode", p.DBCode)
	value.Set("rowcode", p.RowCode)
	value.Set("colcode", p.ColCode)
	if p.WdsWdCode != "" || p.WdsWdValueCode != "" {
		value.Set("wds", fmt.Sprintf(`[{"wdcode":"%s","valuecode":"%s"}]`, p.WdsWdCode, p.WdsWdValueCode))
	} else {
		value.Set("wds", `[]`)
	}
	value.Set("dfwds", fmt.Sprintf(`[{"wdcode":"%s","valuecode":"%s"}]`, p.DfWdsWdCode, p.DfWdsValueCode))
	v := value.Encode()
	return v
}

// marco and province share
func newParam() Param {
	return Param{
		M:              "QueryData",
		ColCode:        "sj",
		DfWdsWdCode:    "sj",
	}
}

// marco share
// 宏观共享字段
func marco(dateRegion string) Param {
	param := newParam()
	param.RowCode = "zb"
	param.DfWdsValueCode = dateRegion
	return param
}

// MacroYear
// 宏观年度
func MacroYear(dateRegion string) Param {
	param := marco(dateRegion)
	param.DBCode = "hgnd"
	return param
}

// MacroSeason
// 宏观季度
func MacroSeason(dateRegion string) Param {
	param := marco(dateRegion)
	param.DBCode = "hgjd"
	return param
}

// MacroMonth
// 宏观月度
func MacroMonth(dateRegion string) Param {
	param := marco(dateRegion)
	param.DBCode = "hgyd"
	return param
}

// province share
// 分省共享字段
func province(dateRegion, indicatorCode string) Param {
	param := newParam()
	param.RowCode = "reg"
	param.WdsWdCode = "zb"
	param.WdsWdValueCode = indicatorCode
	param.DfWdsValueCode = dateRegion
	return param
}

// ProvinceYear
// 分省年度
func ProvinceYear(dateRegion, indicatorCode string) Param {
	param := province(dateRegion, indicatorCode)
	param.DBCode = "fsnd"
	return param
}

// ProvinceSeason
// 分省季度
func ProvinceSeason(dateRegion, indicatorCode string) Param {
	param := province(dateRegion, indicatorCode)
	param.DBCode = "fsjd"
	return param
}

// ProvinceMonth
// 分省月度
func ProvinceMonth(dateRegion, indicatorCode string) Param {
	param := province(dateRegion, indicatorCode)
	param.DBCode = "fsyd"
	return param
}

// 分省月度
//m: QueryData
//dbcode: fsyd
//rowcode: reg
//colcode: sj
//wds: [{"wdcode":"zb","valuecode":"A01010101"}]
//dfwds: [{"wdcode":"sj","valuecode":"201601-202106"}]

// 分省年度
//m: QueryData
//dbcode: fsnd
//rowcode: reg
//colcode: sj
//wds: [{"wdcode":"zb","valuecode":"A090101"}]
//dfwds: [{"wdcode":"sj","valuecode":"2018-2020"}]

// 宏观月度
//m: QueryData
//dbcode: hgyd
//rowcode: zb
//colcode: sj
//wds: []
//dfwds: [{"wdcode":"sj","valuecode":"199912-202106"}]

// 宏观季度
//m: QueryData
//dbcode: hgjd
//rowcode: zb
//colcode: sj
//wds: []
//dfwds: [{"wdcode":"sj","valuecode":"2020,2021"}]