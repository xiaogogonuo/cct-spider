package code

import "fmt"

// 指标名称(国家统计局和系统内部统一)
const (
	CPIName  = "居民消费价格指数(上年=100)"
	CPI3Name = "居民消费价格指数(上年同月=100)"
	CQC2Name = "货币和准货币(M2)供应量_同比增长"
	FAI4Name = "固定资产投资增速"
	GDPRName = "地区生产总值"
	GDP3Name = "国内生产总值增长"
	IAV1Name = "工业增加值_同比增长"
	PMIName  = "制造业采购经理指数"
	PPI3Name = "工业生产者出厂价格指数(上年同月=100)"
	RCLName  = "居民消费水平"
	SCG3Name = "社会消费品零售总额_同比增长"
)

var Indicator = make(map[string]map[string]string)

func setUpIndicator() {
	// 居民消费价格指数(上年=100)
	Indicator[CPIName] = make(map[string]string)
	Indicator[CPIName]["TargetNameEN"] = "CPI"
	Indicator[CPIName]["TargetCode"] = "HG00003"
	Indicator[CPIName]["DataSourceCode"] = "stat"
	Indicator[CPIName]["DataSourceName"] = "国家统计局"
	Indicator[CPIName]["SourceTargetCode"] = "A090101"
	Indicator[CPIName]["IsQuantity"] = "Y"
	Indicator[CPIName]["UnitType"] = ""
	Indicator[CPIName]["UnitName"] = ""
	Indicator[CPIName]["PeriodType"] = fmt.Sprintf("%s", PeriodTypeYear)
	Indicator[CPIName]["PeriodName"] = fmt.Sprintf("%s", PeriodNameYear)
	Indicator[CPIName]["StartYear"] = "1951"
	Indicator[CPIName]["StartSeason"] = ""
	Indicator[CPIName]["StartMonth"] = ""
	Indicator[CPIName]["Zone"] = fmt.Sprintf("%s", Province)

	// 居民消费价格指数(上年同月=100)
	Indicator[CPI3Name] = make(map[string]string)
	Indicator[CPI3Name]["TargetNameEN"] = "CPI3"
	Indicator[CPI3Name]["TargetCode"] = "HG00004"
	Indicator[CPI3Name]["DataSourceCode"] = "stat"
	Indicator[CPI3Name]["DataSourceName"] = "国家统计局"
	Indicator[CPI3Name]["SourceTargetCode"] = "A01010101"
	Indicator[CPI3Name]["IsQuantity"] = "Y"
	Indicator[CPI3Name]["UnitType"] = ""
	Indicator[CPI3Name]["UnitName"] = ""
	Indicator[CPI3Name]["PeriodType"] = fmt.Sprintf("%s", PeriodTypeMonth)
	Indicator[CPI3Name]["PeriodName"] = fmt.Sprintf("%s", PeriodNameMonth)
	Indicator[CPI3Name]["StartYear"] = "2016"
	Indicator[CPI3Name]["StartSeason"] = ""
	Indicator[CPI3Name]["StartMonth"] = "01"
	Indicator[CPI3Name]["Zone"] = fmt.Sprintf("%s", Province)

	// 货币和准货币(M2)供应量_同比增长
	Indicator[CQC2Name] = make(map[string]string)
	Indicator[CQC2Name]["TargetNameEN"] = "M2同比"
	Indicator[CQC2Name]["TargetCode"] = "HG00007"
	Indicator[CQC2Name]["DataSourceCode"] = "stat"
	Indicator[CQC2Name]["DataSourceName"] = "国家统计局"
	Indicator[CQC2Name]["SourceTargetCode"] = "A0D0102"
	Indicator[CQC2Name]["IsQuantity"] = "Y"
	Indicator[CQC2Name]["UnitType"] = UnitTypeP
	Indicator[CQC2Name]["UnitName"] = UnitNameP
	Indicator[CQC2Name]["PeriodType"] = fmt.Sprintf("%s", PeriodTypeMonth)
	Indicator[CQC2Name]["PeriodName"] = fmt.Sprintf("%s", PeriodNameMonth)
	Indicator[CQC2Name]["StartYear"] = "1999"
	Indicator[CQC2Name]["StartSeason"] = ""
	Indicator[CQC2Name]["StartMonth"] = "12"
	Indicator[CQC2Name]["Zone"] = fmt.Sprintf("%s", Marco)

	// 固定资产投资增速(固定资产投资额_累计增长)
	Indicator[FAI4Name] = make(map[string]string)
	Indicator[FAI4Name]["TargetNameEN"] = "固定资产投资增速"
	Indicator[FAI4Name]["TargetCode"] = "HG00039"
	Indicator[FAI4Name]["DataSourceCode"] = "stat"
	Indicator[FAI4Name]["DataSourceName"] = "国家统计局"
	Indicator[FAI4Name]["SourceTargetCode"] = "A040102"
	Indicator[FAI4Name]["IsQuantity"] = "Y"
	Indicator[FAI4Name]["UnitType"] = UnitTypeP
	Indicator[FAI4Name]["UnitName"] = UnitNameP
	Indicator[FAI4Name]["PeriodType"] = fmt.Sprintf("%s", PeriodTypeMonth)
	Indicator[FAI4Name]["PeriodName"] = fmt.Sprintf("%s", PeriodNameMonth)
	Indicator[FAI4Name]["StartYear"] = "1998"
	Indicator[FAI4Name]["StartSeason"] = ""
	Indicator[FAI4Name]["StartMonth"] = "02"
	Indicator[FAI4Name]["Zone"] = fmt.Sprintf("%s", Marco)

	// 国内生产总值增长
	Indicator[GDP3Name] = make(map[string]string)
	Indicator[GDP3Name]["TargetNameEN"] = "GDP"
	Indicator[GDP3Name]["TargetCode"] = "HG00001"
	Indicator[GDP3Name]["DataSourceCode"] = "stat"
	Indicator[GDP3Name]["DataSourceName"] = "国家统计局"
	Indicator[GDP3Name]["SourceTargetCode"] = "A020801"
	Indicator[GDP3Name]["IsQuantity"] = "Y"
	Indicator[GDP3Name]["UnitType"] = UnitTypeP
	Indicator[GDP3Name]["UnitName"] = UnitNameP
	Indicator[GDP3Name]["PeriodType"] = fmt.Sprintf("%s", PeriodTypeYear)
	Indicator[GDP3Name]["PeriodName"] = fmt.Sprintf("%s", PeriodNameYear)
	Indicator[GDP3Name]["StartYear"] = "1978"
	Indicator[GDP3Name]["StartSeason"] = ""
	Indicator[GDP3Name]["StartMonth"] = ""
	Indicator[GDP3Name]["Zone"] = fmt.Sprintf("%s", Marco)

	// 地区生产总值
	Indicator[GDPRName] = make(map[string]string)
	Indicator[GDPRName]["TargetNameEN"] = "RegionGDP"
	Indicator[GDPRName]["TargetCode"] = "HG00002"
	Indicator[GDPRName]["DataSourceCode"] = "stat"
	Indicator[GDPRName]["DataSourceName"] = "国家统计局"
	Indicator[GDPRName]["SourceTargetCode"] = "A020101"
	Indicator[GDPRName]["IsQuantity"] = "Y"
	Indicator[GDPRName]["UnitType"] = UnitTypeE
	Indicator[GDPRName]["UnitName"] = UnitNameE
	Indicator[GDPRName]["PeriodType"] = fmt.Sprintf("%s", PeriodTypeYear)
	Indicator[GDPRName]["PeriodName"] = fmt.Sprintf("%s", PeriodNameYear)
	Indicator[GDPRName]["StartYear"] = "1992"
	Indicator[GDPRName]["StartSeason"] = ""
	Indicator[GDPRName]["StartMonth"] = ""
	Indicator[GDPRName]["Zone"] = fmt.Sprintf("%s", Province)

	// 工业增加值_同比增长
	Indicator[IAV1Name] = make(map[string]string)
	Indicator[IAV1Name]["TargetNameEN"] = "IAV同比"
	Indicator[IAV1Name]["TargetCode"] = "HG00016"
	Indicator[IAV1Name]["DataSourceCode"] = "stat"
	Indicator[IAV1Name]["DataSourceName"] = "国家统计局"
	Indicator[IAV1Name]["SourceTargetCode"] = "A020101"
	Indicator[IAV1Name]["IsQuantity"] = "Y"
	Indicator[IAV1Name]["UnitType"] = UnitTypeP
	Indicator[IAV1Name]["UnitName"] = UnitNameP
	Indicator[IAV1Name]["PeriodType"] = fmt.Sprintf("%s", PeriodTypeMonth)
	Indicator[IAV1Name]["PeriodName"] = fmt.Sprintf("%s", PeriodNameMonth)
	Indicator[IAV1Name]["StartYear"] = "1998"
	Indicator[IAV1Name]["StartSeason"] = ""
	Indicator[IAV1Name]["StartMonth"] = "07"
	Indicator[IAV1Name]["Zone"] = fmt.Sprintf("%s", Marco)

	// 制造业采购经理指数
	Indicator[PMIName] = make(map[string]string)
	Indicator[PMIName]["TargetNameEN"] = "PMI"
	Indicator[PMIName]["TargetCode"] = "HG00020"
	Indicator[PMIName]["DataSourceCode"] = "stat"
	Indicator[PMIName]["DataSourceName"] = "国家统计局"
	Indicator[PMIName]["SourceTargetCode"] = "A0B0101"
	Indicator[PMIName]["IsQuantity"] = "Y"
	Indicator[PMIName]["UnitType"] = UnitTypeP
	Indicator[PMIName]["UnitName"] = UnitNameP
	Indicator[PMIName]["PeriodType"] = fmt.Sprintf("%s", PeriodTypeMonth)
	Indicator[PMIName]["PeriodName"] = fmt.Sprintf("%s", PeriodNameMonth)
	Indicator[PMIName]["StartYear"] = "2005"
	Indicator[PMIName]["StartSeason"] = ""
	Indicator[PMIName]["StartMonth"] = "01"
	Indicator[PMIName]["Zone"] = fmt.Sprintf("%s", Marco)

	// 工业生产者出厂价格指数(上年同月=100)
	Indicator[PPI3Name] = make(map[string]string)
	Indicator[PPI3Name]["TargetNameEN"] = "PPI月度同比"
	Indicator[PPI3Name]["TargetCode"] = "HG00023"
	Indicator[PPI3Name]["DataSourceCode"] = "stat"
	Indicator[PPI3Name]["DataSourceName"] = "国家统计局"
	Indicator[PPI3Name]["SourceTargetCode"] = "A01080101"
	Indicator[PPI3Name]["IsQuantity"] = "Y"
	Indicator[PPI3Name]["UnitType"] = ""
	Indicator[PPI3Name]["UnitName"] = ""
	Indicator[PPI3Name]["PeriodType"] = fmt.Sprintf("%s", PeriodTypeMonth)
	Indicator[PPI3Name]["PeriodName"] = fmt.Sprintf("%s", PeriodNameMonth)
	Indicator[PPI3Name]["StartYear"] = "1996"
	Indicator[PPI3Name]["StartSeason"] = ""
	Indicator[PPI3Name]["StartMonth"] = "10"
	Indicator[PPI3Name]["Zone"] = fmt.Sprintf("%s", Marco)

	// 居民消费水平
	Indicator[RCLName] = make(map[string]string)
	Indicator[RCLName]["TargetNameEN"] = "RCL"
	Indicator[RCLName]["TargetCode"] = "HG00025"
	Indicator[RCLName]["DataSourceCode"] = "stat"
	Indicator[RCLName]["DataSourceName"] = "国家统计局"
	Indicator[RCLName]["SourceTargetCode"] = "A020501"
	Indicator[RCLName]["IsQuantity"] = "Y"
	Indicator[RCLName]["UnitType"] = UnitTypeY
	Indicator[RCLName]["UnitName"] = UnitNameY
	Indicator[RCLName]["PeriodType"] = fmt.Sprintf("%s", PeriodTypeYear)
	Indicator[RCLName]["PeriodName"] = fmt.Sprintf("%s", PeriodNameYear)
	Indicator[RCLName]["StartYear"] = "1992"
	Indicator[RCLName]["StartSeason"] = ""
	Indicator[RCLName]["StartMonth"] = ""
	Indicator[RCLName]["Zone"] = fmt.Sprintf("%s", Province)

	// 社会消费品零售总额_同比增长
	Indicator[SCG3Name] = make(map[string]string)
	Indicator[SCG3Name]["TargetNameEN"] = "SCG同比"
	Indicator[SCG3Name]["TargetCode"] = "HG00029"
	Indicator[SCG3Name]["DataSourceCode"] = "stat"
	Indicator[SCG3Name]["DataSourceName"] = "国家统计局"
	Indicator[SCG3Name]["SourceTargetCode"] = "A070103"
	Indicator[SCG3Name]["IsQuantity"] = "Y"
	Indicator[SCG3Name]["UnitType"] = UnitTypeP
	Indicator[SCG3Name]["UnitName"] = UnitNameP
	Indicator[SCG3Name]["PeriodType"] = fmt.Sprintf("%s", PeriodTypeMonth)
	Indicator[SCG3Name]["PeriodName"] = fmt.Sprintf("%s", PeriodNameMonth)
	Indicator[SCG3Name]["StartYear"] = "2000"
	Indicator[SCG3Name]["StartSeason"] = ""
	Indicator[SCG3Name]["StartMonth"] = "01"
	Indicator[SCG3Name]["Zone"] = fmt.Sprintf("%s", Marco)
}