package code

import "fmt"

const (
	BDIName = "波罗的海指数"
	LPIName = "物流业景气指数"
	CCIName = "中国大宗商品指数"
	PIIName = "造纸行业价格指数"
	COIName = "原油价格指数"
	RMBName = "人民币汇率"
	USDName = "美元指数"
	SHIName = "上海银行间同业拆放利率"
	TBIName = "国债指数"
	LPRName = "贷款基准利率"
)

var Indicator = make(map[string]map[string]string)

func setUpIndicator() {
	// 波罗的海干散货指数
	Indicator[BDIName] = make(map[string]string)
	Indicator[BDIName]["Flag"] = "eastmoney"
	Indicator[BDIName]["TargetNameEN"] = "BDI"
	Indicator[BDIName]["TargetCode"] = "HY00003"
	Indicator[BDIName]["DataSourceCode"] = "eastmoney"
	Indicator[BDIName]["DataSourceName"] = "东方财富"
	Indicator[BDIName]["SourceTargetCode"] = "EMI00107664"
	Indicator[BDIName]["IsQuantity"] = "Y"
	Indicator[BDIName]["UnitType"] = ""
	Indicator[BDIName]["UnitName"] = ""
	Indicator[BDIName]["PeriodType"] = fmt.Sprintf("%s", PeriodTypeDay)
	Indicator[BDIName]["PeriodName"] = fmt.Sprintf("%s", PeriodNameDay)

	// 物流景气指数
	Indicator[LPIName] = make(map[string]string)
	Indicator[LPIName]["Flag"] = "eastmoney"
	Indicator[LPIName]["TargetNameEN"] = "LPI"
	Indicator[LPIName]["TargetCode"] = "HY00001"
	Indicator[LPIName]["DataSourceCode"] = "eastmoney"
	Indicator[LPIName]["DataSourceName"] = "东方财富"
	Indicator[LPIName]["SourceTargetCode"] = "EMI00352262"
	Indicator[LPIName]["IsQuantity"] = "Y"
	Indicator[LPIName]["UnitType"] = ""
	Indicator[LPIName]["UnitName"] = ""
	Indicator[LPIName]["PeriodType"] = fmt.Sprintf("%s", PeriodTypeDay)
	Indicator[LPIName]["PeriodName"] = fmt.Sprintf("%s", PeriodNameDay)

	// 中国大宗商品指数
	Indicator[CCIName] = make(map[string]string)
	Indicator[CCIName]["Flag"] = "eastmoney"
	Indicator[CCIName]["TargetNameEN"] = "CCI"
	Indicator[CCIName]["TargetCode"] = "HY00002"
	Indicator[CCIName]["DataSourceCode"] = "eastmoney"
	Indicator[CCIName]["DataSourceName"] = "东方财富"
	Indicator[CCIName]["SourceTargetCode"] = "EMI00662535"
	Indicator[CCIName]["IsQuantity"] = "Y"
	Indicator[CCIName]["UnitType"] = ""
	Indicator[CCIName]["UnitName"] = ""
	Indicator[CCIName]["PeriodType"] = fmt.Sprintf("%s", PeriodTypeDay)
	Indicator[CCIName]["PeriodName"] = fmt.Sprintf("%s", PeriodNameDay)

	// 上海银行间同业拆放利率
	Indicator[SHIName] = make(map[string]string)
	Indicator[SHIName]["Flag"] = "shi"
	Indicator[SHIName]["TargetNameEN"] = "SHIBOR"
	Indicator[SHIName]["TargetCode"] = "HY00006"
	Indicator[SHIName]["DataSourceCode"] = "eastmoney"
	Indicator[SHIName]["DataSourceName"] = "东方财富"
	Indicator[SHIName]["SourceTargetCode"] = "EMI99221"
	Indicator[SHIName]["IsQuantity"] = "Y"
	Indicator[SHIName]["UnitType"] = UnitTypeP
	Indicator[SHIName]["UnitName"] = UnitNameP
	Indicator[SHIName]["PeriodType"] = fmt.Sprintf("%s", PeriodTypeDay)
	Indicator[SHIName]["PeriodName"] = fmt.Sprintf("%s", PeriodNameDay)

	// 造纸行业价格指数
	Indicator[PIIName] = make(map[string]string)
	Indicator[PIIName]["Flag"] = "sci"
	Indicator[PIIName]["TargetNameEN"] = "PII"
	Indicator[PIIName]["TargetCode"] = "HY00010"
	Indicator[PIIName]["DataSourceCode"] = "sci"
	Indicator[PIIName]["DataSourceName"] = "卓创资讯"
	Indicator[PIIName]["SourceTargetCode"] = "SCI195"
	Indicator[PIIName]["IsQuantity"] = "Y"
	Indicator[PIIName]["UnitType"] = ""
	Indicator[PIIName]["UnitName"] = ""
	Indicator[PIIName]["PeriodType"] = fmt.Sprintf("%s", PeriodTypeDay)
	Indicator[PIIName]["PeriodName"] = fmt.Sprintf("%s", PeriodNameDay)
	Indicator[PIIName]["HY"] = "造纸"
	Indicator[PIIName]["Level"] = "0"
	Indicator[PIIName]["Path1"] = "造纸行业价格指数"
	Indicator[PIIName]["Path2"] = ""
	Indicator[PIIName]["Path3"] = ""
	Indicator[PIIName]["Path4"] = ""
	Indicator[PIIName]["Type"] = "2"

	// 原油价格指数
	Indicator[COIName] = make(map[string]string)
	Indicator[COIName]["Flag"] = "sci"
	Indicator[COIName]["TargetNameEN"] = "COI"
	Indicator[COIName]["TargetCode"] = "HY00004"
	Indicator[COIName]["DataSourceCode"] = "sci"
	Indicator[COIName]["DataSourceName"] = "卓创资讯"
	Indicator[COIName]["SourceTargetCode"] = "SCI195"
	Indicator[COIName]["IsQuantity"] = "Y"
	Indicator[COIName]["UnitType"] = ""
	Indicator[COIName]["UnitName"] = ""
	Indicator[COIName]["PeriodType"] = fmt.Sprintf("%s", PeriodTypeDay)
	Indicator[COIName]["PeriodName"] = fmt.Sprintf("%s", PeriodNameDay)
	Indicator[COIName]["HY"] = "能源"
	Indicator[COIName]["Level"] = "2"
	Indicator[COIName]["Path1"] = "能源价格指数"
	Indicator[COIName]["Path2"] = "石油价格指数"
	Indicator[COIName]["Path3"] = "原油价格指数"
	Indicator[COIName]["Path4"] = ""
	Indicator[COIName]["Type"] = "2"

	// 人民币汇率
	Indicator[RMBName] = make(map[string]string)
	Indicator[RMBName]["Flag"] = "sina"
	Indicator[RMBName]["TargetNameEN"] = "RMBExchangeRate"
	Indicator[RMBName]["TargetCode"] = "HY00008"
	Indicator[RMBName]["DataSourceCode"] = "sina"
	Indicator[RMBName]["DataSourceName"] = "新浪财经"
	Indicator[RMBName]["SourceTargetCode"] = "CNYUSD"
	Indicator[RMBName]["IsQuantity"] = "Y"
	Indicator[RMBName]["UnitType"] = ""
	Indicator[RMBName]["UnitName"] = ""
	Indicator[RMBName]["PeriodType"] = fmt.Sprintf("%s", PeriodTypeDay)
	Indicator[RMBName]["PeriodName"] = fmt.Sprintf("%s", PeriodNameDay)

	// 美元指数
	Indicator[USDName] = make(map[string]string)
	Indicator[USDName]["Flag"] = "sina"
	Indicator[USDName]["TargetNameEN"] = "USDX"
	Indicator[USDName]["TargetCode"] = "HY00009"
	Indicator[USDName]["DataSourceCode"] = "sina"
	Indicator[USDName]["DataSourceName"] = "新浪财经"
	Indicator[USDName]["SourceTargetCode"] = "DINIW"
	Indicator[USDName]["IsQuantity"] = "Y"
	Indicator[USDName]["UnitType"] = ""
	Indicator[USDName]["UnitName"] = ""
	Indicator[USDName]["PeriodType"] = fmt.Sprintf("%s", PeriodTypeDay)
	Indicator[USDName]["PeriodName"] = fmt.Sprintf("%s", PeriodNameDay)

	// 国债指数
	Indicator[TBIName] = make(map[string]string)
	Indicator[TBIName]["Flag"] = "tbi"
	Indicator[TBIName]["TargetNameEN"] = "TBI"
	Indicator[TBIName]["TargetCode"] = "HY00005"
	Indicator[TBIName]["DataSourceCode"] = "ifeng"
	Indicator[TBIName]["DataSourceName"] = "凤凰网财经"
	Indicator[TBIName]["SourceTargetCode"] = "sh000012"
	Indicator[TBIName]["IsQuantity"] = "Y"
	Indicator[TBIName]["UnitType"] = ""
	Indicator[TBIName]["UnitName"] = ""
	Indicator[TBIName]["PeriodType"] = fmt.Sprintf("%s", PeriodTypeDay)
	Indicator[TBIName]["PeriodName"] = fmt.Sprintf("%s", PeriodNameDay)

	// 贷款基准利率
	Indicator[LPRName] = make(map[string]string)
	Indicator[LPRName]["Flag"] = "lpr"
	Indicator[LPRName]["TargetNameEN"] = "LPR"
	Indicator[LPRName]["TargetCode"] = "HY00007"
	Indicator[LPRName]["DataSourceCode"] = "yinhang123"
	Indicator[LPRName]["DataSourceName"] = "银行信息港"
	Indicator[LPRName]["SourceTargetCode"] = "1254559"
	Indicator[LPRName]["IsQuantity"] = "Y"
	Indicator[LPRName]["UnitType"] = UnitTypeP
	Indicator[LPRName]["UnitName"] = UnitNameP
	Indicator[LPRName]["PeriodType"] = fmt.Sprintf("%s", PeriodTypeDay)
	Indicator[LPRName]["PeriodName"] = fmt.Sprintf("%s", PeriodNameDay)
}

/*
HY00001,物流业景气指数,LPI ok
HY00002,中国大宗商品指数,CCI ok
HY00003,波罗的海指数,BDI ok
HY00004,原油价格指数,COI ok
HY00005,国债指数,TBI ok
HY00006,上海银行间同业拆放利率,SHIBOR ok
HY00007,贷款基准利率,LPR ok
HY00008,人民币汇率,RMBExchangeRate ok
HY00009,美元指数,USDX ok
HY00010,造纸行业价格指数,PII ok
 */