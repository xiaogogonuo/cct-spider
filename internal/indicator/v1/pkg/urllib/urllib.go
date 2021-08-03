package urllib

import (
	"fmt"
	"net/url"
)

var (
	EasyQuery = "https://data.stats.gov.cn/easyquery.htm?"
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

// marco and province share
func newParam() Param {
	return Param{
		M:              "QueryData",
		RowCode:        "zb",
		ColCode:        "sj",
		DfWdsWdCode:    "sj",
	}
}

// marco share
// 宏观
func marco(dateRegion string) Param {
	param := newParam()
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

// MacroQuarter
// 宏观季度
func MacroQuarter(dateRegion string) Param {
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
// 分省
func province(dateRegion, regionCode string) Param {
	param := newParam()
	param.WdsWdCode = "reg"
	param.WdsWdValueCode = regionCode
	param.DfWdsValueCode = dateRegion
	return param
}

// ProvinceYear
// 分省年度
func ProvinceYear(dateRegion, regionCode string) Param {
	param := province(dateRegion, regionCode)
	param.DBCode = "fsnd"
	return param
}

// ProvinceQuarter
// 分省季度
func ProvinceQuarter(dateRegion, regionCode string) Param {
	param := province(dateRegion, regionCode)
	param.DBCode = "fsjd"
	return param
}

// ProvinceMonth
// 分省月度
func ProvinceMonth(dateRegion, regionCode string) Param {
	param := province(dateRegion, regionCode)
	param.DBCode = "fsyd"
	return param
}

func (p Param) Encode() string {
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
	return EasyQuery + v
}

//m: QueryData
//dbcode: fsnd  // 分省年度
//rowcode: zb
//colcode: sj
//wds: [{"wdcode":"reg","valuecode":"110000"}]
//dfwds: [{"wdcode":"sj","valuecode":"last25"}]

//m: QueryData
//dbcode: fsyd  // 分省月度
//rowcode: zb
//colcode: sj
//wds: [{"wdcode":"reg","valuecode":"110000"}]
//dfwds: [{"wdcode":"sj","valuecode":"last24"}]

// 指标年度
//m: QueryData
//dbcode: hgnd
//rowcode: zb
//colcode: sj
//wds: []
//dfwds: [{"wdcode":"sj","valuecode":"last10"}]

// 指标月度
//m: QueryData
//dbcode: hgyd
//rowcode: zb
//colcode: sj
//wds: []
//dfwds: [{"wdcode":"sj","valuecode":"last8"}]

// 指标季度
//m: QueryData
//dbcode: hgjd
//rowcode: zb
//colcode: sj
//wds: []
//dfwds: [{"wdcode":"sj","valuecode":"last100"}]