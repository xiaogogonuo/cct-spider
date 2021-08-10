package indicator_code

import "fmt"

const (
	Marco    = "0" // 宏观
	Province = "1" // 分省
)

// 计量单位
const (
	UnitTypeG = "10"
	UnitNameG = "个"
	UnitTypeR = "20"
	UnitNameR = "人"
	UnitTypeY = "30"
	UnitNameY = "元"
	UnitTypeW = "31"
	UnitNameW = "万元"
	UnitTypeE = "32"
	UnitNameE = "亿元"
	UnitTypeJ = "40"
	UnitNameJ = "家"
	UnitTypeP = "50"
	UnitNameP = "%"
)

// 指标周期
const (
	PeriodTypeYear      = "10"
	PeriodNameYear      = "年"
	PeriodTypeSeason    = "20"
	PeriodNameSeason    = "季"
	PeriodTypeMonth     = "30"
	PeriodNameMonth     = "月"
	PeriodTypeWeek      = "40"
	PeriodNameWeek      = "周"
	PeriodTypeDay       = "50"
	PeriodNameDay       = "日"
	PeriodTypeReal      = "60"
	PeriodNameReal      = "实时"
	PeriodTypeUnCertain = "90"
	PeriodNameUnCertain = "不定期"
)

var PeriodTypeName = make(map[string]string)
var PeriodNameType = make(map[string]string)

func setUpPeriod() {
	PeriodTypeName[PeriodTypeYear] = PeriodNameYear
	PeriodTypeName[PeriodTypeSeason] = PeriodNameSeason
	PeriodTypeName[PeriodTypeMonth] = PeriodNameMonth
	PeriodTypeName[PeriodTypeWeek] = PeriodNameWeek
	PeriodTypeName[PeriodTypeDay] = PeriodNameDay
	PeriodTypeName[PeriodTypeReal] = PeriodNameReal
	PeriodTypeName[PeriodTypeUnCertain] = PeriodNameUnCertain

	PeriodNameType[PeriodNameYear] = PeriodTypeYear
	PeriodNameType[PeriodNameSeason] = PeriodTypeSeason
	PeriodNameType[PeriodNameMonth] = PeriodTypeMonth
	PeriodNameType[PeriodNameWeek] = PeriodTypeWeek
	PeriodNameType[PeriodNameDay] = PeriodTypeDay
	PeriodNameType[PeriodNameReal] = PeriodTypeReal
	PeriodNameType[PeriodNameUnCertain] = PeriodTypeUnCertain
}

// 指标名称(国家统计局和系统内部统一)
const (
	CPIName  = "居民消费价格指数(上年=100)"
	CPI1Name = "居民消费价格指数(1978=100)"
	CPI2Name = "居民消费价格指数(上月=100)"
	CPI3Name = "居民消费价格指数(上年同月=100)"
	CPI4Name = "居民消费价格指数(上年同期=100)"
	CQCName  = "货币和准货币(M2)供应量同比增长率"
	CQC1Name = "货币和准货币(M2)供应量_期末值"
	CQC2Name = "货币和准货币(M2)供应量_同比增长"
	FAIName  = "固定资产投资价格指数(上年=100)"
	FAI1Name = "固定资产投资价格指数(1990=100)"
	FAI2Name = "固定资产投资价格指数_当季值(上年同季=100)"
	FAI3Name = "固定资产投资价格指数_累计值(上年同期=100)"
	FAI4Name = "固定资产投资增速"
	GDPName  = "国内生产总值"
	GDPRName = "地区生产总值"
	GDP1Name = "国内生产总值_当季值"
	GDP2Name = "国内生产总值_累计值"
	GDP3Name = "国内生产总值增长"
	HCEName  = "居民人均消费支出"
	HCE1Name = "居民人均消费支出_同比增长"
	HCE2Name = "居民人均消费支出_累计值"
	HCE3Name = "居民人均消费支出_累计增长"
	IAVName  = "工业增加值"
	IAV1Name = "工业增加值_同比增长"
	IAV2Name = "工业增加值_累计增长"
	IAV3Name = "工业增加值_当季值"
	IAV4Name = "工业增加值_累计值"
	PMIName  = "制造业采购经理指数"
	PPIName  = "工业生产者出厂价格指数(上年=100)"
	PPI1Name = "工业生产者出厂价格指数(1985=100)"
	PPI2Name = "工业生产者出厂价格指数(上月=100)"
	PPI3Name = "工业生产者出厂价格指数(上年同月=100)"
	PPI4Name = "工业生产者出厂价格指数(上年同期=100)"
	RCLName  = "居民消费水平"
	SCGName  = "社会消费品零售总额"
	SCG1Name = "社会消费品零售总额_当期值"
	SCG2Name = "社会消费品零售总额_累计值"
	SCG3Name = "社会消费品零售总额_同比增长"
	SCG4Name = "社会消费品零售总额_累计增长"
	URIName  = "城镇居民人均可支配收入"
	URI1Name = "城镇居民人均可支配收入_同比增长"
	URI2Name = "城镇居民人均可支配收入_累计值"
	URI3Name = "城镇居民人均可支配收入_累计增长"
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
	Indicator[CPIName]["AreaType"] = fmt.Sprintf("%s", Province)

	// 居民消费价格指数(上月=100)
	//Indicator[CPI2Name] = make(map[string]string)
	//Indicator[CPI2Name]["TargetCode"] = "HG00037"
	//Indicator[CPI2Name]["DataSourceCode"] = "stat"
	//Indicator[CPI2Name]["DataSourceName"] = "国家统计局"
	//Indicator[CPI2Name]["SourceTargetCode"] = "A01030101"
	//Indicator[CPI2Name]["IsQuantity"] = "Y"
	//Indicator[CPI2Name]["UnitType"] = ""
	//Indicator[CPI2Name]["UnitName"] = ""
	//Indicator[CPI2Name]["PeriodType"] = PeriodTypeMonth
	//Indicator[CPI2Name]["PeriodName"] = PeriodNameMonth
	//Indicator[CPI2Name]["StartYear"] = "2016"
	//Indicator[CPI2Name]["StartSeason"] = ""
	//Indicator[CPI2Name]["StartMonth"] = "01"
	//Indicator[CPI2Name]["DateType"] = Monthly
	//Indicator[CPI2Name]["AreaType"] = Province

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
	Indicator[CPI3Name]["AreaType"] = fmt.Sprintf("%s", Province)

	// 居民消费价格指数(上年同期=100)
	//Indicator[CPI4Name] = make(map[string]string)
	//Indicator[CPI4Name]["TargetCode"] = "HG00038"
	//Indicator[CPI4Name]["DataSourceCode"] = "stat"
	//Indicator[CPI4Name]["DataSourceName"] = "国家统计局"
	//Indicator[CPI4Name]["SourceTargetCode"] = "A01020101"
	//Indicator[CPI4Name]["IsQuantity"] = "Y"
	//Indicator[CPI4Name]["UnitType"] = ""
	//Indicator[CPI4Name]["UnitName"] = ""
	//Indicator[CPI4Name]["PeriodType"] = PeriodTypeMonth
	//Indicator[CPI4Name]["PeriodName"] = PeriodNameMonth
	//Indicator[CPI4Name]["StartYear"] = "2016"
	//Indicator[CPI4Name]["StartSeason"] = ""
	//Indicator[CPI4Name]["StartMonth"] = "01"
	//Indicator[CPI4Name]["DateType"] = Monthly
	//Indicator[CPI4Name]["AreaType"] = Province

	// 货币和准货币(M2)供应量同比增长率
	//Indicator[CQCName] = make(map[string]string)
	//Indicator[CQCName]["TargetCode"] = "HG00005"
	//Indicator[CQCName]["DataSourceCode"] = "stat"
	//Indicator[CQCName]["DataSourceName"] = "国家统计局"
	//Indicator[CQCName]["SourceTargetCode"] = "A0L0309"
	//Indicator[CQCName]["IsQuantity"] = "Y"
	//Indicator[CQCName]["UnitType"] = UnitTypeP
	//Indicator[CQCName]["UnitName"] = UnitNameP
	//Indicator[CQCName]["PeriodType"] = PeriodTypeYear
	//Indicator[CQCName]["PeriodName"] = PeriodNameYear
	//Indicator[CQCName]["StartYear"] = "1991"
	//Indicator[CQCName]["StartSeason"] = ""
	//Indicator[CQCName]["StartMonth"] = ""
	//Indicator[CQCName]["DateType"] = Annual
	//Indicator[CQCName]["AreaType"] = Marco

	// 货币和准货币(M2)供应量_期末值(亿元)
	//Indicator[CQC1Name] = make(map[string]string)
	//Indicator[CQC1Name]["TargetCode"] = "HG00006"
	//Indicator[CQC1Name]["DataSourceCode"] = "stat"
	//Indicator[CQC1Name]["DataSourceName"] = "国家统计局"
	//Indicator[CQC1Name]["SourceTargetCode"] = "A0D0101"
	//Indicator[CQC1Name]["IsQuantity"] = "Y"
	//Indicator[CQC1Name]["UnitType"] = UnitTypeE
	//Indicator[CQC1Name]["UnitName"] = UnitNameE
	//Indicator[CQC1Name]["PeriodType"] = PeriodTypeMonth
	//Indicator[CQC1Name]["PeriodName"] = PeriodNameMonth
	//Indicator[CQC1Name]["StartYear"] = "1999"
	//Indicator[CQC1Name]["StartSeason"] = ""
	//Indicator[CQC1Name]["StartMonth"] = "12"
	//Indicator[CQC1Name]["DateType"] = Monthly
	//Indicator[CQC1Name]["AreaType"] = Marco

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
	Indicator[CQC2Name]["AreaType"] = fmt.Sprintf("%s", Marco)

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
	Indicator[FAI4Name]["AreaType"] = fmt.Sprintf("%s", Marco)

	// 固定资产投资价格指数(上年=100)
	//Indicator[FAIName] = make(map[string]string)
	//Indicator[FAIName]["TargetCode"] = "HG00008"
	//Indicator[FAIName]["DataSourceCode"] = "stat"
	//Indicator[FAIName]["DataSourceName"] = "国家统计局"
	//Indicator[FAIName]["SourceTargetCode"] = "A090107"
	//Indicator[FAIName]["IsQuantity"] = "Y"
	//Indicator[FAIName]["UnitType"] = ""
	//Indicator[FAIName]["UnitName"] = ""
	//Indicator[FAIName]["PeriodType"] = PeriodTypeYear
	//Indicator[FAIName]["PeriodName"] = PeriodNameYear
	//Indicator[FAIName]["StartYear"] = "1990"
	//Indicator[FAIName]["StartSeason"] = ""
	//Indicator[FAIName]["StartMonth"] = ""
	//Indicator[FAIName]["DateType"] = Annual
	//Indicator[FAIName]["AreaType"] = Marco

	// 固定资产投资价格指数_当季值(上年同季=100)
	//Indicator[FAI2Name] = make(map[string]string)
	//Indicator[FAI2Name]["TargetCode"] = "HG00009"
	//Indicator[FAI2Name]["DataSourceCode"] = "stat"
	//Indicator[FAI2Name]["DataSourceName"] = "国家统计局"
	//Indicator[FAI2Name]["SourceTargetCode"] = "A060201"
	//Indicator[FAI2Name]["IsQuantity"] = "Y"
	//Indicator[FAI2Name]["UnitType"] = ""
	//Indicator[FAI2Name]["UnitName"] = ""
	//Indicator[FAI2Name]["PeriodType"] = PeriodTypeSeason
	//Indicator[FAI2Name]["PeriodName"] = PeriodNameSeason
	//Indicator[FAI2Name]["StartYear"] = "1998"
	//Indicator[FAI2Name]["StartSeason"] = "B"
	//Indicator[FAI2Name]["StartMonth"] = ""
	//Indicator[FAI2Name]["DateType"] = Quarterly
	//Indicator[FAI2Name]["AreaType"] = Marco

	// 固定资产投资价格指数_累计值(上年同期=100)
	//Indicator[FAI3Name] = make(map[string]string)
	//Indicator[FAI3Name]["TargetCode"] = "HG00010"
	//Indicator[FAI3Name]["DataSourceCode"] = "stat"
	//Indicator[FAI3Name]["DataSourceName"] = "国家统计局"
	//Indicator[FAI3Name]["SourceTargetCode"] = "A060301"
	//Indicator[FAI3Name]["IsQuantity"] = "Y"
	//Indicator[FAI3Name]["UnitType"] = ""
	//Indicator[FAI3Name]["UnitName"] = ""
	//Indicator[FAI3Name]["PeriodType"] = PeriodTypeSeason
	//Indicator[FAI3Name]["PeriodName"] = PeriodNameSeason
	//Indicator[FAI3Name]["StartYear"] = "2007"
	//Indicator[FAI3Name]["StartSeason"] = "A"
	//Indicator[FAI3Name]["StartMonth"] = ""
	//Indicator[FAI3Name]["DateType"] = Quarterly
	//Indicator[FAI3Name]["AreaType"] = Marco

	// 国内生产总值
	//Indicator[GDPName] = make(map[string]string)
	//Indicator[GDPName]["TargetCode"] = "HG00001"
	//Indicator[GDPName]["DataSourceCode"] = "stat"
	//Indicator[GDPName]["DataSourceName"] = "国家统计局"
	//Indicator[GDPName]["SourceTargetCode"] = "A020102"
	//Indicator[GDPName]["IsQuantity"] = "Y"
	//Indicator[GDPName]["UnitType"] = UnitTypeE
	//Indicator[GDPName]["UnitName"] = UnitNameE
	//Indicator[GDPName]["PeriodType"] = PeriodTypeYear
	//Indicator[GDPName]["PeriodName"] = PeriodNameYear
	//Indicator[GDPName]["StartYear"] = "1952"
	//Indicator[GDPName]["StartSeason"] = ""
	//Indicator[GDPName]["StartMonth"] = ""
	//Indicator[GDPName]["DateType"] = Annual
	//Indicator[GDPName]["AreaType"] = Marco

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
	Indicator[GDP3Name]["AreaType"] = fmt.Sprintf("%s", Marco)

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
	Indicator[GDPRName]["AreaType"] = fmt.Sprintf("%s", Province)

	// 国内生产总值_当季值
	//Indicator[GDP1Name] = make(map[string]string)
	//Indicator[GDP1Name]["TargetCode"] = "HG00035"
	//Indicator[GDP1Name]["DataSourceCode"] = "stat"
	//Indicator[GDP1Name]["DataSourceName"] = "国家统计局"
	//Indicator[GDP1Name]["SourceTargetCode"] = "A010101"
	//Indicator[GDP1Name]["IsQuantity"] = "Y"
	//Indicator[GDP1Name]["UnitType"] = UnitTypeE
	//Indicator[GDP1Name]["UnitName"] = UnitNameE
	//Indicator[GDP1Name]["PeriodType"] = PeriodTypeSeason
	//Indicator[GDP1Name]["PeriodName"] = PeriodNameSeason
	//Indicator[GDP1Name]["StartYear"] = "1992"
	//Indicator[GDP1Name]["StartSeason"] = "A"
	//Indicator[GDP1Name]["StartMonth"] = ""
	//Indicator[GDP1Name]["DateType"] = Quarterly
	//Indicator[GDP1Name]["AreaType"] = Marco

	// 国内生产总值_累计值
	//Indicator[GDP2Name] = make(map[string]string)
	//Indicator[GDP2Name]["TargetCode"] = "HG00036"
	//Indicator[GDP2Name]["DataSourceCode"] = "stat"
	//Indicator[GDP2Name]["DataSourceName"] = "国家统计局"
	//Indicator[GDP2Name]["SourceTargetCode"] = "A010102"
	//Indicator[GDP2Name]["IsQuantity"] = "Y"
	//Indicator[GDP2Name]["UnitType"] = UnitTypeE
	//Indicator[GDP2Name]["UnitName"] = UnitNameE
	//Indicator[GDP2Name]["PeriodType"] = PeriodTypeSeason
	//Indicator[GDP2Name]["PeriodName"] = PeriodNameSeason
	//Indicator[GDP2Name]["StartYear"] = "1992"
	//Indicator[GDP2Name]["StartSeason"] = "A"
	//Indicator[GDP2Name]["StartMonth"] = ""
	//Indicator[GDP2Name]["DateType"] = Quarterly
	//Indicator[GDP2Name]["AreaType"] = Marco

	// 居民人均消费支出
	//Indicator[HCEName] = make(map[string]string)
	//Indicator[HCEName]["TargetCode"] = "HG00011"
	//Indicator[HCEName]["DataSourceCode"] = "stat"
	//Indicator[HCEName]["DataSourceName"] = "国家统计局"
	//Indicator[HCEName]["SourceTargetCode"] = "A0A0107"
	//Indicator[HCEName]["IsQuantity"] = "Y"
	//Indicator[HCEName]["UnitType"] = UnitTypeY
	//Indicator[HCEName]["UnitName"] = UnitNameY
	//Indicator[HCEName]["PeriodType"] = PeriodTypeYear
	//Indicator[HCEName]["PeriodName"] = PeriodNameYear
	//Indicator[HCEName]["StartYear"] = "2013"
	//Indicator[HCEName]["StartSeason"] = ""
	//Indicator[HCEName]["StartMonth"] = ""
	//Indicator[HCEName]["DateType"] = Annual
	//Indicator[HCEName]["AreaType"] = Marco

	// 居民人均消费支出_同比增长
	//Indicator[HCE1Name] = make(map[string]string)
	//Indicator[HCE1Name]["TargetCode"] = "HG00012"
	//Indicator[HCE1Name]["DataSourceCode"] = "stat"
	//Indicator[HCE1Name]["DataSourceName"] = "国家统计局"
	//Indicator[HCE1Name]["SourceTargetCode"] = "A0A0108"
	//Indicator[HCE1Name]["IsQuantity"] = "Y"
	//Indicator[HCE1Name]["UnitType"] = UnitTypeP
	//Indicator[HCE1Name]["UnitName"] = UnitNameP
	//Indicator[HCE1Name]["PeriodType"] = PeriodTypeYear
	//Indicator[HCE1Name]["PeriodName"] = PeriodNameYear
	//Indicator[HCE1Name]["StartYear"] = "2014"
	//Indicator[HCE1Name]["StartSeason"] = ""
	//Indicator[HCE1Name]["StartMonth"] = ""
	//Indicator[HCE1Name]["DateType"] = Annual
	//Indicator[HCE1Name]["AreaType"] = Marco

	// 居民人均消费支出_累计值
	//Indicator[HCE2Name] = make(map[string]string)
	//Indicator[HCE2Name]["TargetCode"] = "HG00013"
	//Indicator[HCE2Name]["DataSourceCode"] = "stat"
	//Indicator[HCE2Name]["DataSourceName"] = "国家统计局"
	//Indicator[HCE2Name]["SourceTargetCode"] = "A050109"
	//Indicator[HCE2Name]["IsQuantity"] = "Y"
	//Indicator[HCE2Name]["UnitType"] = UnitTypeY
	//Indicator[HCE2Name]["UnitName"] = UnitNameY
	//Indicator[HCE2Name]["PeriodType"] = PeriodTypeSeason
	//Indicator[HCE2Name]["PeriodName"] = PeriodNameSeason
	//Indicator[HCE2Name]["StartYear"] = "2013"
	//Indicator[HCE2Name]["StartSeason"] = "A"
	//Indicator[HCE2Name]["StartMonth"] = ""
	//Indicator[HCE2Name]["DateType"] = Quarterly
	//Indicator[HCE2Name]["AreaType"] = Marco

	// 居民人均消费支出_累计增长
	//Indicator[HCE3Name] = make(map[string]string)
	//Indicator[HCE3Name]["TargetCode"] = "HG00014"
	//Indicator[HCE3Name]["DataSourceCode"] = "stat"
	//Indicator[HCE3Name]["DataSourceName"] = "国家统计局"
	//Indicator[HCE3Name]["SourceTargetCode"] = "A05010A"
	//Indicator[HCE3Name]["IsQuantity"] = "Y"
	//Indicator[HCE3Name]["UnitType"] = UnitTypeP
	//Indicator[HCE3Name]["UnitName"] = UnitNameP
	//Indicator[HCE3Name]["PeriodType"] = PeriodTypeSeason
	//Indicator[HCE3Name]["PeriodName"] = PeriodNameSeason
	//Indicator[HCE3Name]["StartYear"] = "2014"
	//Indicator[HCE3Name]["StartSeason"] = "A"
	//Indicator[HCE3Name]["StartMonth"] = ""
	//Indicator[HCE3Name]["DateType"] = Quarterly
	//Indicator[HCE3Name]["AreaType"] = Marco

	// 工业增加值
	//Indicator[IAVName] = make(map[string]string)
	//Indicator[IAVName]["TargetCode"] = "HG00015"
	//Indicator[IAVName]["DataSourceCode"] = "stat"
	//Indicator[IAVName]["DataSourceName"] = "国家统计局"
	//Indicator[IAVName]["SourceTargetCode"] = "A020403"
	//Indicator[IAVName]["IsQuantity"] = "Y"
	//Indicator[IAVName]["UnitType"] = UnitTypeE
	//Indicator[IAVName]["UnitName"] = UnitNameE
	//Indicator[IAVName]["PeriodType"] = PeriodTypeYear
	//Indicator[IAVName]["PeriodName"] = PeriodNameYear
	//Indicator[IAVName]["StartYear"] = "1952"
	//Indicator[IAVName]["StartSeason"] = ""
	//Indicator[IAVName]["StartMonth"] = ""
	//Indicator[IAVName]["DateType"] = Annual
	//Indicator[IAVName]["AreaType"] = Marco

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
	Indicator[IAV1Name]["AreaType"] = fmt.Sprintf("%s", Marco)

	// 工业增加值_累计增长
	//Indicator[IAV2Name] = make(map[string]string)
	//Indicator[IAV2Name]["TargetCode"] = "HG00017"
	//Indicator[IAV2Name]["DataSourceCode"] = "stat"
	//Indicator[IAV2Name]["DataSourceName"] = "国家统计局"
	//Indicator[IAV2Name]["SourceTargetCode"] = "A020102"
	//Indicator[IAV2Name]["IsQuantity"] = "Y"
	//Indicator[IAV2Name]["UnitType"] = UnitTypeP
	//Indicator[IAV2Name]["UnitName"] = UnitNameP
	//Indicator[IAV2Name]["PeriodType"] = PeriodTypeMonth
	//Indicator[IAV2Name]["PeriodName"] = PeriodNameMonth
	//Indicator[IAV2Name]["StartYear"] = "1998"
	//Indicator[IAV2Name]["StartSeason"] = ""
	//Indicator[IAV2Name]["StartMonth"] = "07"
	//Indicator[IAV2Name]["DateType"] = Monthly
	//Indicator[IAV2Name]["AreaType"] = Marco

	// 工业增加值_当季值
	//Indicator[IAV3Name] = make(map[string]string)
	//Indicator[IAV3Name]["TargetCode"] = "HG00018"
	//Indicator[IAV3Name]["DataSourceCode"] = "stat"
	//Indicator[IAV3Name]["DataSourceName"] = "国家统计局"
	//Indicator[IAV3Name]["SourceTargetCode"] = "A01010B"
	//Indicator[IAV3Name]["IsQuantity"] = "Y"
	//Indicator[IAV3Name]["UnitType"] = UnitTypeE
	//Indicator[IAV3Name]["UnitName"] = UnitNameE
	//Indicator[IAV3Name]["PeriodType"] = PeriodTypeSeason
	//Indicator[IAV3Name]["PeriodName"] = PeriodNameSeason
	//Indicator[IAV3Name]["StartYear"] = "1992"
	//Indicator[IAV3Name]["StartSeason"] = "A"
	//Indicator[IAV3Name]["StartMonth"] = ""
	//Indicator[IAV3Name]["DateType"] = Quarterly
	//Indicator[IAV3Name]["AreaType"] = Marco

	// 工业增加值_累计值
	//Indicator[IAV4Name] = make(map[string]string)
	//Indicator[IAV4Name]["TargetCode"] = "HG00019"
	//Indicator[IAV4Name]["DataSourceCode"] = "stat"
	//Indicator[IAV4Name]["DataSourceName"] = "国家统计局"
	//Indicator[IAV4Name]["SourceTargetCode"] = "A01010C"
	//Indicator[IAV4Name]["IsQuantity"] = "Y"
	//Indicator[IAV4Name]["UnitType"] = UnitTypeE
	//Indicator[IAV4Name]["UnitName"] = UnitNameE
	//Indicator[IAV4Name]["PeriodType"] = PeriodTypeSeason
	//Indicator[IAV4Name]["PeriodName"] = PeriodNameSeason
	//Indicator[IAV4Name]["StartYear"] = "1992"
	//Indicator[IAV4Name]["StartSeason"] = "A"
	//Indicator[IAV4Name]["StartMonth"] = ""
	//Indicator[IAV4Name]["DateType"] = Quarterly
	//Indicator[IAV4Name]["AreaType"] = Marco

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
	Indicator[PMIName]["AreaType"] = fmt.Sprintf("%s", Marco)

	// 工业生产者出厂价格指数(上年=100)
	//Indicator[PPIName] = make(map[string]string)
	//Indicator[PPIName]["TargetCode"] = "HG00021"
	//Indicator[PPIName]["DataSourceCode"] = "stat"
	//Indicator[PPIName]["DataSourceName"] = "国家统计局"
	//Indicator[PPIName]["SourceTargetCode"] = "A090105"
	//Indicator[PPIName]["IsQuantity"] = "Y"
	//Indicator[PPIName]["UnitType"] = ""
	//Indicator[PPIName]["UnitName"] = ""
	//Indicator[PPIName]["PeriodType"] = PeriodTypeYear
	//Indicator[PPIName]["PeriodName"] = PeriodNameYear
	//Indicator[PPIName]["StartYear"] = "1978"
	//Indicator[PPIName]["StartSeason"] = ""
	//Indicator[PPIName]["StartMonth"] = ""
	//Indicator[PPIName]["DateType"] = Annual
	//Indicator[PPIName]["AreaType"] = Marco

	// 工业生产者出厂价格指数(上月=100)
	//Indicator[PPI2Name] = make(map[string]string)
	//Indicator[PPI2Name]["TargetCode"] = "HG00022"
	//Indicator[PPI2Name]["DataSourceCode"] = "stat"
	//Indicator[PPI2Name]["DataSourceName"] = "国家统计局"
	//Indicator[PPI2Name]["SourceTargetCode"] = "A01080701"
	//Indicator[PPI2Name]["IsQuantity"] = "Y"
	//Indicator[PPI2Name]["UnitType"] = ""
	//Indicator[PPI2Name]["UnitName"] = ""
	//Indicator[PPI2Name]["PeriodType"] = PeriodTypeMonth
	//Indicator[PPI2Name]["PeriodName"] = PeriodNameMonth
	//Indicator[PPI2Name]["StartYear"] = "2011"
	//Indicator[PPI2Name]["StartSeason"] = ""
	//Indicator[PPI2Name]["StartMonth"] = "01"
	//Indicator[PPI2Name]["DateType"] = Monthly
	//Indicator[PPI2Name]["AreaType"] = Marco

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
	Indicator[PPI3Name]["AreaType"] = fmt.Sprintf("%s", Marco)

	// 工业生产者出厂价格指数(上年同期=100)
	//Indicator[PPI4Name] = make(map[string]string)
	//Indicator[PPI4Name]["TargetCode"] = "HG00024"
	//Indicator[PPI4Name]["DataSourceCode"] = "stat"
	//Indicator[PPI4Name]["DataSourceName"] = "国家统计局"
	//Indicator[PPI4Name]["SourceTargetCode"] = "A01080401"
	//Indicator[PPI4Name]["IsQuantity"] = "Y"
	//Indicator[PPI4Name]["UnitType"] = ""
	//Indicator[PPI4Name]["UnitName"] = ""
	//Indicator[PPI4Name]["PeriodType"] = PeriodTypeMonth
	//Indicator[PPI4Name]["PeriodName"] = PeriodNameMonth
	//Indicator[PPI4Name]["StartYear"] = "2011"
	//Indicator[PPI4Name]["StartSeason"] = ""
	//Indicator[PPI4Name]["StartMonth"] = "01"
	//Indicator[PPI4Name]["DateType"] = Monthly
	//Indicator[PPI4Name]["AreaType"] = Marco

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
	Indicator[RCLName]["AreaType"] = fmt.Sprintf("%s", Province)

	// 社会消费品零售总额
	//Indicator[SCGName] = make(map[string]string)
	//Indicator[SCGName]["TargetCode"] = "HG00026"
	//Indicator[SCGName]["DataSourceCode"] = "stat"
	//Indicator[SCGName]["DataSourceName"] = "国家统计局"
	//Indicator[SCGName]["SourceTargetCode"] = "A0H01"
	//Indicator[SCGName]["IsQuantity"] = "Y"
	//Indicator[SCGName]["UnitType"] = UnitTypeE
	//Indicator[SCGName]["UnitName"] = UnitNameE
	//Indicator[SCGName]["PeriodType"] = PeriodTypeYear
	//Indicator[SCGName]["PeriodName"] = PeriodNameYear
	//Indicator[SCGName]["StartYear"] = "1952"
	//Indicator[SCGName]["StartSeason"] = ""
	//Indicator[SCGName]["StartMonth"] = ""
	//Indicator[SCGName]["DateType"] = Annual
	//Indicator[SCGName]["AreaType"] = Marco

	// 社会消费品零售总额_当期值
	//Indicator[SCG1Name] = make(map[string]string)
	//Indicator[SCG1Name]["TargetCode"] = "HG00027"
	//Indicator[SCG1Name]["DataSourceCode"] = "stat"
	//Indicator[SCG1Name]["DataSourceName"] = "国家统计局"
	//Indicator[SCG1Name]["SourceTargetCode"] = "A070101"
	//Indicator[SCG1Name]["IsQuantity"] = "Y"
	//Indicator[SCG1Name]["UnitType"] = UnitTypeE
	//Indicator[SCG1Name]["UnitName"] = UnitNameE
	//Indicator[SCG1Name]["PeriodType"] = PeriodTypeMonth
	//Indicator[SCG1Name]["PeriodName"] = PeriodNameMonth
	//Indicator[SCG1Name]["StartYear"] = "1984"
	//Indicator[SCG1Name]["StartSeason"] = ""
	//Indicator[SCG1Name]["StartMonth"] = "01"
	//Indicator[SCG1Name]["DateType"] = Monthly
	//Indicator[SCG1Name]["AreaType"] = Marco

	// 社会消费品零售总额_累计值
	//Indicator[SCG2Name] = make(map[string]string)
	//Indicator[SCG2Name]["TargetCode"] = "HG00028"
	//Indicator[SCG2Name]["DataSourceCode"] = "stat"
	//Indicator[SCG2Name]["DataSourceName"] = "国家统计局"
	//Indicator[SCG2Name]["SourceTargetCode"] = "A070102"
	//Indicator[SCG2Name]["IsQuantity"] = "Y"
	//Indicator[SCG2Name]["UnitType"] = UnitTypeE
	//Indicator[SCG2Name]["UnitName"] = UnitNameE
	//Indicator[SCG2Name]["PeriodType"] = PeriodTypeMonth
	//Indicator[SCG2Name]["PeriodName"] = PeriodNameMonth
	//Indicator[SCG2Name]["StartYear"] = "2000"
	//Indicator[SCG2Name]["StartSeason"] = ""
	//Indicator[SCG2Name]["StartMonth"] = "01"
	//Indicator[SCG2Name]["DateType"] = Monthly
	//Indicator[SCG2Name]["AreaType"] = Marco

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
	Indicator[SCG3Name]["AreaType"] = fmt.Sprintf("%s", Marco)

	// 社会消费品零售总额_累计增长
	//Indicator[SCG4Name] = make(map[string]string)
	//Indicator[SCG4Name]["TargetCode"] = "HG00030"
	//Indicator[SCG4Name]["DataSourceCode"] = "stat"
	//Indicator[SCG4Name]["DataSourceName"] = "国家统计局"
	//Indicator[SCG4Name]["SourceTargetCode"] = "A070104"
	//Indicator[SCG4Name]["IsQuantity"] = "Y"
	//Indicator[SCG4Name]["UnitType"] = UnitTypeP
	//Indicator[SCG4Name]["UnitName"] = UnitNameP
	//Indicator[SCG4Name]["PeriodType"] = PeriodTypeMonth
	//Indicator[SCG4Name]["PeriodName"] = PeriodNameMonth
	//Indicator[SCG4Name]["StartYear"] = "2000"
	//Indicator[SCG4Name]["StartSeason"] = ""
	//Indicator[SCG4Name]["StartMonth"] = "01"
	//Indicator[SCG4Name]["DateType"] = Monthly
	//Indicator[SCG4Name]["AreaType"] = Marco

	// 城镇居民人均可支配收入
	//Indicator[URIName] = make(map[string]string)
	//Indicator[URIName]["outerCode"] = "A0A0103"
	//Indicator[URIName]["innerCode"] = "HG00031"
	//Indicator[URIName]["startYear"] = "2013"

	// 城镇居民人均可支配收入_同比增长
	//Indicator[URI1Name] = make(map[string]string)
	//Indicator[URI1Name]["outerCode"] = "A0A0104"
	//Indicator[URI1Name]["innerCode"] = "HG00032"
	//Indicator[URI1Name]["startYear"] = "2014"

	// 城镇居民人均可支配收入_累计值
	//Indicator[URI2Name] = make(map[string]string)
	//Indicator[URI2Name]["outerCode"] = "A050201"
	//Indicator[URI2Name]["innerCode"] = "HG00033"
	//Indicator[URI2Name]["startYear"] = "2013"

	// 城镇居民人均可支配收入_累计增长
	//Indicator[URI3Name] = make(map[string]string)
	//Indicator[URI3Name]["outerCode"] = "A050202"
	//Indicator[URI3Name]["innerCode"] = "HG00034"
	//Indicator[URI3Name]["startYear"] = "2014"
}

func init() {
	setUpPeriod()
	setUpIndicator()
}