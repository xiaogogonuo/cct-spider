package response

// 卓创资讯

type StructSCI struct {
	List []Data `json:"List"`
}

type Data struct {
	MDataValue float64 `json:"MDataValue"`
	DataDate   string  `json:"DataDate"`
}

// 东方财富

type StructEastMoney struct {
	Result  R    `json:"result"`
	Success bool `json:"success"`
}

type R struct {
	Data []D `json:"data"`
}

type D struct {
	ReportDate     string  `json:"REPORT_DATE"`
	IndicatorValue float64 `json:"INDICATOR_VALUE"`
}

// 外汇交易中心

type StructLPR struct {
	Records []LPR `json:"records"`
}

type LPR struct {
	DateString string `json:"dateString"`
	LoanRate   string `json:"loanRate"`
}
