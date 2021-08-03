package indicator_code

const (
	Marco         = "0" // 宏观
	Province      = "1" // 分省
	MarcoProvince = "2" // 宏观、分省
)

const (
	Annual                 = "0" // 年度
	Quarterly              = "1" // 季度
	Monthly                = "2" // 月度
	Weekly                 = "3" // 周度
	Daily                  = "4" // 日度
	AnnualQuarterly        = "5" // 年度、季度
	AnnualMonthly          = "6" // 年度、月度
	QuarterlyMonthly       = "7" // 季度、月度
	AnnualQuarterlyMonthly = "8" // 年度、季度、月度
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

// IndexMap record index name mapping outer code and inner code
var IndexMap map[string]map[string]string

func setUpIndexMap() {
	IndexMap = make(map[string]map[string]string)

	// 居民消费价格指数(上年=100)
	IndexMap[CPIName] = make(map[string]string)
	IndexMap[CPIName]["TargetNameEN"] = "CPI"
	IndexMap[CPIName]["TargetCode"] = "HG00003"
	IndexMap[CPIName]["DataSourceCode"] = "stat"
	IndexMap[CPIName]["DataSourceName"] = "国家统计局"
	IndexMap[CPIName]["SourceTargetCode"] = "A090101"
	IndexMap[CPIName]["IsQuantity"] = "Y"
	IndexMap[CPIName]["UnitType"] = ""
	IndexMap[CPIName]["UnitName"] = ""
	IndexMap[CPIName]["PeriodType"] = PeriodTypeYear
	IndexMap[CPIName]["PeriodName"] = PeriodNameYear
	IndexMap[CPIName]["StartYear"] = "1951"
	IndexMap[CPIName]["StartSeason"] = ""
	IndexMap[CPIName]["StartMonth"] = ""
	IndexMap[CPIName]["DateType"] = Annual
	IndexMap[CPIName]["AreaType"] = Province

	// 居民消费价格指数(上月=100)
	//IndexMap[CPI2Name] = make(map[string]string)
	//IndexMap[CPI2Name]["TargetCode"] = "HG00037"
	//IndexMap[CPI2Name]["DataSourceCode"] = "stat"
	//IndexMap[CPI2Name]["DataSourceName"] = "国家统计局"
	//IndexMap[CPI2Name]["SourceTargetCode"] = "A01030101"
	//IndexMap[CPI2Name]["IsQuantity"] = "Y"
	//IndexMap[CPI2Name]["UnitType"] = ""
	//IndexMap[CPI2Name]["UnitName"] = ""
	//IndexMap[CPI2Name]["PeriodType"] = PeriodTypeMonth
	//IndexMap[CPI2Name]["PeriodName"] = PeriodNameMonth
	//IndexMap[CPI2Name]["StartYear"] = "2016"
	//IndexMap[CPI2Name]["StartSeason"] = ""
	//IndexMap[CPI2Name]["StartMonth"] = "01"
	//IndexMap[CPI2Name]["DateType"] = Monthly
	//IndexMap[CPI2Name]["AreaType"] = Province

	// 居民消费价格指数(上年同月=100)
	IndexMap[CPI3Name] = make(map[string]string)
	IndexMap[CPI3Name]["TargetNameEN"] = "CPI3"
	IndexMap[CPI3Name]["TargetCode"] = "HG00004"
	IndexMap[CPI3Name]["DataSourceCode"] = "stat"
	IndexMap[CPI3Name]["DataSourceName"] = "国家统计局"
	IndexMap[CPI3Name]["SourceTargetCode"] = "A01010101"
	IndexMap[CPI3Name]["IsQuantity"] = "Y"
	IndexMap[CPI3Name]["UnitType"] = ""
	IndexMap[CPI3Name]["UnitName"] = ""
	IndexMap[CPI3Name]["PeriodType"] = PeriodTypeMonth
	IndexMap[CPI3Name]["PeriodName"] = PeriodNameMonth
	IndexMap[CPI3Name]["StartYear"] = "2016"
	IndexMap[CPI3Name]["StartSeason"] = ""
	IndexMap[CPI3Name]["StartMonth"] = "01"
	IndexMap[CPI3Name]["DateType"] = Monthly
	IndexMap[CPI3Name]["AreaType"] = Province

	// 居民消费价格指数(上年同期=100)
	//IndexMap[CPI4Name] = make(map[string]string)
	//IndexMap[CPI4Name]["TargetCode"] = "HG00038"
	//IndexMap[CPI4Name]["DataSourceCode"] = "stat"
	//IndexMap[CPI4Name]["DataSourceName"] = "国家统计局"
	//IndexMap[CPI4Name]["SourceTargetCode"] = "A01020101"
	//IndexMap[CPI4Name]["IsQuantity"] = "Y"
	//IndexMap[CPI4Name]["UnitType"] = ""
	//IndexMap[CPI4Name]["UnitName"] = ""
	//IndexMap[CPI4Name]["PeriodType"] = PeriodTypeMonth
	//IndexMap[CPI4Name]["PeriodName"] = PeriodNameMonth
	//IndexMap[CPI4Name]["StartYear"] = "2016"
	//IndexMap[CPI4Name]["StartSeason"] = ""
	//IndexMap[CPI4Name]["StartMonth"] = "01"
	//IndexMap[CPI4Name]["DateType"] = Monthly
	//IndexMap[CPI4Name]["AreaType"] = Province

	// 货币和准货币(M2)供应量同比增长率
	//IndexMap[CQCName] = make(map[string]string)
	//IndexMap[CQCName]["TargetCode"] = "HG00005"
	//IndexMap[CQCName]["DataSourceCode"] = "stat"
	//IndexMap[CQCName]["DataSourceName"] = "国家统计局"
	//IndexMap[CQCName]["SourceTargetCode"] = "A0L0309"
	//IndexMap[CQCName]["IsQuantity"] = "Y"
	//IndexMap[CQCName]["UnitType"] = UnitTypeP
	//IndexMap[CQCName]["UnitName"] = UnitNameP
	//IndexMap[CQCName]["PeriodType"] = PeriodTypeYear
	//IndexMap[CQCName]["PeriodName"] = PeriodNameYear
	//IndexMap[CQCName]["StartYear"] = "1991"
	//IndexMap[CQCName]["StartSeason"] = ""
	//IndexMap[CQCName]["StartMonth"] = ""
	//IndexMap[CQCName]["DateType"] = Annual
	//IndexMap[CQCName]["AreaType"] = Marco

	// 货币和准货币(M2)供应量_期末值(亿元)
	//IndexMap[CQC1Name] = make(map[string]string)
	//IndexMap[CQC1Name]["TargetCode"] = "HG00006"
	//IndexMap[CQC1Name]["DataSourceCode"] = "stat"
	//IndexMap[CQC1Name]["DataSourceName"] = "国家统计局"
	//IndexMap[CQC1Name]["SourceTargetCode"] = "A0D0101"
	//IndexMap[CQC1Name]["IsQuantity"] = "Y"
	//IndexMap[CQC1Name]["UnitType"] = UnitTypeE
	//IndexMap[CQC1Name]["UnitName"] = UnitNameE
	//IndexMap[CQC1Name]["PeriodType"] = PeriodTypeMonth
	//IndexMap[CQC1Name]["PeriodName"] = PeriodNameMonth
	//IndexMap[CQC1Name]["StartYear"] = "1999"
	//IndexMap[CQC1Name]["StartSeason"] = ""
	//IndexMap[CQC1Name]["StartMonth"] = "12"
	//IndexMap[CQC1Name]["DateType"] = Monthly
	//IndexMap[CQC1Name]["AreaType"] = Marco

	// 货币和准货币(M2)供应量_同比增长
	IndexMap[CQC2Name] = make(map[string]string)
	IndexMap[CQC2Name]["TargetNameEN"] = "M2同比"
	IndexMap[CQC2Name]["TargetCode"] = "HG00007"
	IndexMap[CQC2Name]["DataSourceCode"] = "stat"
	IndexMap[CQC2Name]["DataSourceName"] = "国家统计局"
	IndexMap[CQC2Name]["SourceTargetCode"] = "A0D0102"
	IndexMap[CQC2Name]["IsQuantity"] = "Y"
	IndexMap[CQC2Name]["UnitType"] = UnitTypeP
	IndexMap[CQC2Name]["UnitName"] = UnitNameP
	IndexMap[CQC2Name]["PeriodType"] = PeriodTypeMonth
	IndexMap[CQC2Name]["PeriodName"] = PeriodNameMonth
	IndexMap[CQC2Name]["StartYear"] = "1999"
	IndexMap[CQC2Name]["StartSeason"] = ""
	IndexMap[CQC2Name]["StartMonth"] = "12"
	IndexMap[CQC2Name]["DateType"] = Monthly
	IndexMap[CQC2Name]["AreaType"] = Marco

	// 固定资产投资增速(固定资产投资额_累计增长)
	IndexMap[FAI4Name] = make(map[string]string)
	IndexMap[FAI4Name]["TargetNameEN"] = "固定资产投资增速"
	IndexMap[FAI4Name]["TargetCode"] = "HG00039"
	IndexMap[FAI4Name]["DataSourceCode"] = "stat"
	IndexMap[FAI4Name]["DataSourceName"] = "国家统计局"
	IndexMap[FAI4Name]["SourceTargetCode"] = "A040102"
	IndexMap[FAI4Name]["IsQuantity"] = "Y"
	IndexMap[FAI4Name]["UnitType"] = UnitTypeP
	IndexMap[FAI4Name]["UnitName"] = UnitNameP
	IndexMap[FAI4Name]["PeriodType"] = PeriodTypeMonth
	IndexMap[FAI4Name]["PeriodName"] = PeriodNameMonth
	IndexMap[FAI4Name]["StartYear"] = "1998"
	IndexMap[FAI4Name]["StartSeason"] = ""
	IndexMap[FAI4Name]["StartMonth"] = "02"
	IndexMap[FAI4Name]["DateType"] = Monthly
	IndexMap[FAI4Name]["AreaType"] = Marco

	// 固定资产投资价格指数(上年=100)
	//IndexMap[FAIName] = make(map[string]string)
	//IndexMap[FAIName]["TargetCode"] = "HG00008"
	//IndexMap[FAIName]["DataSourceCode"] = "stat"
	//IndexMap[FAIName]["DataSourceName"] = "国家统计局"
	//IndexMap[FAIName]["SourceTargetCode"] = "A090107"
	//IndexMap[FAIName]["IsQuantity"] = "Y"
	//IndexMap[FAIName]["UnitType"] = ""
	//IndexMap[FAIName]["UnitName"] = ""
	//IndexMap[FAIName]["PeriodType"] = PeriodTypeYear
	//IndexMap[FAIName]["PeriodName"] = PeriodNameYear
	//IndexMap[FAIName]["StartYear"] = "1990"
	//IndexMap[FAIName]["StartSeason"] = ""
	//IndexMap[FAIName]["StartMonth"] = ""
	//IndexMap[FAIName]["DateType"] = Annual
	//IndexMap[FAIName]["AreaType"] = Marco

	// 固定资产投资价格指数_当季值(上年同季=100)
	//IndexMap[FAI2Name] = make(map[string]string)
	//IndexMap[FAI2Name]["TargetCode"] = "HG00009"
	//IndexMap[FAI2Name]["DataSourceCode"] = "stat"
	//IndexMap[FAI2Name]["DataSourceName"] = "国家统计局"
	//IndexMap[FAI2Name]["SourceTargetCode"] = "A060201"
	//IndexMap[FAI2Name]["IsQuantity"] = "Y"
	//IndexMap[FAI2Name]["UnitType"] = ""
	//IndexMap[FAI2Name]["UnitName"] = ""
	//IndexMap[FAI2Name]["PeriodType"] = PeriodTypeSeason
	//IndexMap[FAI2Name]["PeriodName"] = PeriodNameSeason
	//IndexMap[FAI2Name]["StartYear"] = "1998"
	//IndexMap[FAI2Name]["StartSeason"] = "B"
	//IndexMap[FAI2Name]["StartMonth"] = ""
	//IndexMap[FAI2Name]["DateType"] = Quarterly
	//IndexMap[FAI2Name]["AreaType"] = Marco

	// 固定资产投资价格指数_累计值(上年同期=100)
	//IndexMap[FAI3Name] = make(map[string]string)
	//IndexMap[FAI3Name]["TargetCode"] = "HG00010"
	//IndexMap[FAI3Name]["DataSourceCode"] = "stat"
	//IndexMap[FAI3Name]["DataSourceName"] = "国家统计局"
	//IndexMap[FAI3Name]["SourceTargetCode"] = "A060301"
	//IndexMap[FAI3Name]["IsQuantity"] = "Y"
	//IndexMap[FAI3Name]["UnitType"] = ""
	//IndexMap[FAI3Name]["UnitName"] = ""
	//IndexMap[FAI3Name]["PeriodType"] = PeriodTypeSeason
	//IndexMap[FAI3Name]["PeriodName"] = PeriodNameSeason
	//IndexMap[FAI3Name]["StartYear"] = "2007"
	//IndexMap[FAI3Name]["StartSeason"] = "A"
	//IndexMap[FAI3Name]["StartMonth"] = ""
	//IndexMap[FAI3Name]["DateType"] = Quarterly
	//IndexMap[FAI3Name]["AreaType"] = Marco

	// 国内生产总值
	//IndexMap[GDPName] = make(map[string]string)
	//IndexMap[GDPName]["TargetCode"] = "HG00001"
	//IndexMap[GDPName]["DataSourceCode"] = "stat"
	//IndexMap[GDPName]["DataSourceName"] = "国家统计局"
	//IndexMap[GDPName]["SourceTargetCode"] = "A020102"
	//IndexMap[GDPName]["IsQuantity"] = "Y"
	//IndexMap[GDPName]["UnitType"] = UnitTypeE
	//IndexMap[GDPName]["UnitName"] = UnitNameE
	//IndexMap[GDPName]["PeriodType"] = PeriodTypeYear
	//IndexMap[GDPName]["PeriodName"] = PeriodNameYear
	//IndexMap[GDPName]["StartYear"] = "1952"
	//IndexMap[GDPName]["StartSeason"] = ""
	//IndexMap[GDPName]["StartMonth"] = ""
	//IndexMap[GDPName]["DateType"] = Annual
	//IndexMap[GDPName]["AreaType"] = Marco

	// 国内生产总值增长
	IndexMap[GDP3Name] = make(map[string]string)
	IndexMap[GDP3Name]["TargetNameEN"] = "GDP"
	IndexMap[GDP3Name]["TargetCode"] = "HG00001"
	IndexMap[GDP3Name]["DataSourceCode"] = "stat"
	IndexMap[GDP3Name]["DataSourceName"] = "国家统计局"
	IndexMap[GDP3Name]["SourceTargetCode"] = "A020801"
	IndexMap[GDP3Name]["IsQuantity"] = "Y"
	IndexMap[GDP3Name]["UnitType"] = UnitTypeP
	IndexMap[GDP3Name]["UnitName"] = UnitNameP
	IndexMap[GDP3Name]["PeriodType"] = PeriodTypeYear
	IndexMap[GDP3Name]["PeriodName"] = PeriodNameYear
	IndexMap[GDP3Name]["StartYear"] = "1978"
	IndexMap[GDP3Name]["StartSeason"] = ""
	IndexMap[GDP3Name]["StartMonth"] = ""
	IndexMap[GDP3Name]["DateType"] = Annual
	IndexMap[GDP3Name]["AreaType"] = Marco

	// 地区生产总值
	IndexMap[GDPRName] = make(map[string]string)
	IndexMap[GDPRName]["TargetNameEN"] = "RegionGDP"
	IndexMap[GDPRName]["TargetCode"] = "HG00002"
	IndexMap[GDPRName]["DataSourceCode"] = "stat"
	IndexMap[GDPRName]["DataSourceName"] = "国家统计局"
	IndexMap[GDPRName]["SourceTargetCode"] = "A020101"
	IndexMap[GDPRName]["IsQuantity"] = "Y"
	IndexMap[GDPRName]["UnitType"] = UnitTypeE
	IndexMap[GDPRName]["UnitName"] = UnitNameE
	IndexMap[GDPRName]["PeriodType"] = PeriodTypeYear
	IndexMap[GDPRName]["PeriodName"] = PeriodNameYear
	IndexMap[GDPRName]["StartYear"] = "1992"
	IndexMap[GDPRName]["StartSeason"] = ""
	IndexMap[GDPRName]["StartMonth"] = ""
	IndexMap[GDPRName]["DateType"] = Annual
	IndexMap[GDPRName]["AreaType"] = Province

	// 国内生产总值_当季值
	//IndexMap[GDP1Name] = make(map[string]string)
	//IndexMap[GDP1Name]["TargetCode"] = "HG00035"
	//IndexMap[GDP1Name]["DataSourceCode"] = "stat"
	//IndexMap[GDP1Name]["DataSourceName"] = "国家统计局"
	//IndexMap[GDP1Name]["SourceTargetCode"] = "A010101"
	//IndexMap[GDP1Name]["IsQuantity"] = "Y"
	//IndexMap[GDP1Name]["UnitType"] = UnitTypeE
	//IndexMap[GDP1Name]["UnitName"] = UnitNameE
	//IndexMap[GDP1Name]["PeriodType"] = PeriodTypeSeason
	//IndexMap[GDP1Name]["PeriodName"] = PeriodNameSeason
	//IndexMap[GDP1Name]["StartYear"] = "1992"
	//IndexMap[GDP1Name]["StartSeason"] = "A"
	//IndexMap[GDP1Name]["StartMonth"] = ""
	//IndexMap[GDP1Name]["DateType"] = Quarterly
	//IndexMap[GDP1Name]["AreaType"] = Marco

	// 国内生产总值_累计值
	//IndexMap[GDP2Name] = make(map[string]string)
	//IndexMap[GDP2Name]["TargetCode"] = "HG00036"
	//IndexMap[GDP2Name]["DataSourceCode"] = "stat"
	//IndexMap[GDP2Name]["DataSourceName"] = "国家统计局"
	//IndexMap[GDP2Name]["SourceTargetCode"] = "A010102"
	//IndexMap[GDP2Name]["IsQuantity"] = "Y"
	//IndexMap[GDP2Name]["UnitType"] = UnitTypeE
	//IndexMap[GDP2Name]["UnitName"] = UnitNameE
	//IndexMap[GDP2Name]["PeriodType"] = PeriodTypeSeason
	//IndexMap[GDP2Name]["PeriodName"] = PeriodNameSeason
	//IndexMap[GDP2Name]["StartYear"] = "1992"
	//IndexMap[GDP2Name]["StartSeason"] = "A"
	//IndexMap[GDP2Name]["StartMonth"] = ""
	//IndexMap[GDP2Name]["DateType"] = Quarterly
	//IndexMap[GDP2Name]["AreaType"] = Marco

	// 居民人均消费支出
	//IndexMap[HCEName] = make(map[string]string)
	//IndexMap[HCEName]["TargetCode"] = "HG00011"
	//IndexMap[HCEName]["DataSourceCode"] = "stat"
	//IndexMap[HCEName]["DataSourceName"] = "国家统计局"
	//IndexMap[HCEName]["SourceTargetCode"] = "A0A0107"
	//IndexMap[HCEName]["IsQuantity"] = "Y"
	//IndexMap[HCEName]["UnitType"] = UnitTypeY
	//IndexMap[HCEName]["UnitName"] = UnitNameY
	//IndexMap[HCEName]["PeriodType"] = PeriodTypeYear
	//IndexMap[HCEName]["PeriodName"] = PeriodNameYear
	//IndexMap[HCEName]["StartYear"] = "2013"
	//IndexMap[HCEName]["StartSeason"] = ""
	//IndexMap[HCEName]["StartMonth"] = ""
	//IndexMap[HCEName]["DateType"] = Annual
	//IndexMap[HCEName]["AreaType"] = Marco

	// 居民人均消费支出_同比增长
	//IndexMap[HCE1Name] = make(map[string]string)
	//IndexMap[HCE1Name]["TargetCode"] = "HG00012"
	//IndexMap[HCE1Name]["DataSourceCode"] = "stat"
	//IndexMap[HCE1Name]["DataSourceName"] = "国家统计局"
	//IndexMap[HCE1Name]["SourceTargetCode"] = "A0A0108"
	//IndexMap[HCE1Name]["IsQuantity"] = "Y"
	//IndexMap[HCE1Name]["UnitType"] = UnitTypeP
	//IndexMap[HCE1Name]["UnitName"] = UnitNameP
	//IndexMap[HCE1Name]["PeriodType"] = PeriodTypeYear
	//IndexMap[HCE1Name]["PeriodName"] = PeriodNameYear
	//IndexMap[HCE1Name]["StartYear"] = "2014"
	//IndexMap[HCE1Name]["StartSeason"] = ""
	//IndexMap[HCE1Name]["StartMonth"] = ""
	//IndexMap[HCE1Name]["DateType"] = Annual
	//IndexMap[HCE1Name]["AreaType"] = Marco

	// 居民人均消费支出_累计值
	//IndexMap[HCE2Name] = make(map[string]string)
	//IndexMap[HCE2Name]["TargetCode"] = "HG00013"
	//IndexMap[HCE2Name]["DataSourceCode"] = "stat"
	//IndexMap[HCE2Name]["DataSourceName"] = "国家统计局"
	//IndexMap[HCE2Name]["SourceTargetCode"] = "A050109"
	//IndexMap[HCE2Name]["IsQuantity"] = "Y"
	//IndexMap[HCE2Name]["UnitType"] = UnitTypeY
	//IndexMap[HCE2Name]["UnitName"] = UnitNameY
	//IndexMap[HCE2Name]["PeriodType"] = PeriodTypeSeason
	//IndexMap[HCE2Name]["PeriodName"] = PeriodNameSeason
	//IndexMap[HCE2Name]["StartYear"] = "2013"
	//IndexMap[HCE2Name]["StartSeason"] = "A"
	//IndexMap[HCE2Name]["StartMonth"] = ""
	//IndexMap[HCE2Name]["DateType"] = Quarterly
	//IndexMap[HCE2Name]["AreaType"] = Marco

	// 居民人均消费支出_累计增长
	//IndexMap[HCE3Name] = make(map[string]string)
	//IndexMap[HCE3Name]["TargetCode"] = "HG00014"
	//IndexMap[HCE3Name]["DataSourceCode"] = "stat"
	//IndexMap[HCE3Name]["DataSourceName"] = "国家统计局"
	//IndexMap[HCE3Name]["SourceTargetCode"] = "A05010A"
	//IndexMap[HCE3Name]["IsQuantity"] = "Y"
	//IndexMap[HCE3Name]["UnitType"] = UnitTypeP
	//IndexMap[HCE3Name]["UnitName"] = UnitNameP
	//IndexMap[HCE3Name]["PeriodType"] = PeriodTypeSeason
	//IndexMap[HCE3Name]["PeriodName"] = PeriodNameSeason
	//IndexMap[HCE3Name]["StartYear"] = "2014"
	//IndexMap[HCE3Name]["StartSeason"] = "A"
	//IndexMap[HCE3Name]["StartMonth"] = ""
	//IndexMap[HCE3Name]["DateType"] = Quarterly
	//IndexMap[HCE3Name]["AreaType"] = Marco

	// 工业增加值
	//IndexMap[IAVName] = make(map[string]string)
	//IndexMap[IAVName]["TargetCode"] = "HG00015"
	//IndexMap[IAVName]["DataSourceCode"] = "stat"
	//IndexMap[IAVName]["DataSourceName"] = "国家统计局"
	//IndexMap[IAVName]["SourceTargetCode"] = "A020403"
	//IndexMap[IAVName]["IsQuantity"] = "Y"
	//IndexMap[IAVName]["UnitType"] = UnitTypeE
	//IndexMap[IAVName]["UnitName"] = UnitNameE
	//IndexMap[IAVName]["PeriodType"] = PeriodTypeYear
	//IndexMap[IAVName]["PeriodName"] = PeriodNameYear
	//IndexMap[IAVName]["StartYear"] = "1952"
	//IndexMap[IAVName]["StartSeason"] = ""
	//IndexMap[IAVName]["StartMonth"] = ""
	//IndexMap[IAVName]["DateType"] = Annual
	//IndexMap[IAVName]["AreaType"] = Marco

	// 工业增加值_同比增长
	IndexMap[IAV1Name] = make(map[string]string)
	IndexMap[IAV1Name]["TargetNameEN"] = "IAV同比"
	IndexMap[IAV1Name]["TargetCode"] = "HG00016"
	IndexMap[IAV1Name]["DataSourceCode"] = "stat"
	IndexMap[IAV1Name]["DataSourceName"] = "国家统计局"
	IndexMap[IAV1Name]["SourceTargetCode"] = "A020101"
	IndexMap[IAV1Name]["IsQuantity"] = "Y"
	IndexMap[IAV1Name]["UnitType"] = UnitTypeP
	IndexMap[IAV1Name]["UnitName"] = UnitNameP
	IndexMap[IAV1Name]["PeriodType"] = PeriodTypeMonth
	IndexMap[IAV1Name]["PeriodName"] = PeriodNameMonth
	IndexMap[IAV1Name]["StartYear"] = "1998"
	IndexMap[IAV1Name]["StartSeason"] = ""
	IndexMap[IAV1Name]["StartMonth"] = "07"
	IndexMap[IAV1Name]["DateType"] = Monthly
	IndexMap[IAV1Name]["AreaType"] = Marco

	// 工业增加值_累计增长
	//IndexMap[IAV2Name] = make(map[string]string)
	//IndexMap[IAV2Name]["TargetCode"] = "HG00017"
	//IndexMap[IAV2Name]["DataSourceCode"] = "stat"
	//IndexMap[IAV2Name]["DataSourceName"] = "国家统计局"
	//IndexMap[IAV2Name]["SourceTargetCode"] = "A020102"
	//IndexMap[IAV2Name]["IsQuantity"] = "Y"
	//IndexMap[IAV2Name]["UnitType"] = UnitTypeP
	//IndexMap[IAV2Name]["UnitName"] = UnitNameP
	//IndexMap[IAV2Name]["PeriodType"] = PeriodTypeMonth
	//IndexMap[IAV2Name]["PeriodName"] = PeriodNameMonth
	//IndexMap[IAV2Name]["StartYear"] = "1998"
	//IndexMap[IAV2Name]["StartSeason"] = ""
	//IndexMap[IAV2Name]["StartMonth"] = "07"
	//IndexMap[IAV2Name]["DateType"] = Monthly
	//IndexMap[IAV2Name]["AreaType"] = Marco

	// 工业增加值_当季值
	//IndexMap[IAV3Name] = make(map[string]string)
	//IndexMap[IAV3Name]["TargetCode"] = "HG00018"
	//IndexMap[IAV3Name]["DataSourceCode"] = "stat"
	//IndexMap[IAV3Name]["DataSourceName"] = "国家统计局"
	//IndexMap[IAV3Name]["SourceTargetCode"] = "A01010B"
	//IndexMap[IAV3Name]["IsQuantity"] = "Y"
	//IndexMap[IAV3Name]["UnitType"] = UnitTypeE
	//IndexMap[IAV3Name]["UnitName"] = UnitNameE
	//IndexMap[IAV3Name]["PeriodType"] = PeriodTypeSeason
	//IndexMap[IAV3Name]["PeriodName"] = PeriodNameSeason
	//IndexMap[IAV3Name]["StartYear"] = "1992"
	//IndexMap[IAV3Name]["StartSeason"] = "A"
	//IndexMap[IAV3Name]["StartMonth"] = ""
	//IndexMap[IAV3Name]["DateType"] = Quarterly
	//IndexMap[IAV3Name]["AreaType"] = Marco

	// 工业增加值_累计值
	//IndexMap[IAV4Name] = make(map[string]string)
	//IndexMap[IAV4Name]["TargetCode"] = "HG00019"
	//IndexMap[IAV4Name]["DataSourceCode"] = "stat"
	//IndexMap[IAV4Name]["DataSourceName"] = "国家统计局"
	//IndexMap[IAV4Name]["SourceTargetCode"] = "A01010C"
	//IndexMap[IAV4Name]["IsQuantity"] = "Y"
	//IndexMap[IAV4Name]["UnitType"] = UnitTypeE
	//IndexMap[IAV4Name]["UnitName"] = UnitNameE
	//IndexMap[IAV4Name]["PeriodType"] = PeriodTypeSeason
	//IndexMap[IAV4Name]["PeriodName"] = PeriodNameSeason
	//IndexMap[IAV4Name]["StartYear"] = "1992"
	//IndexMap[IAV4Name]["StartSeason"] = "A"
	//IndexMap[IAV4Name]["StartMonth"] = ""
	//IndexMap[IAV4Name]["DateType"] = Quarterly
	//IndexMap[IAV4Name]["AreaType"] = Marco

	// 制造业采购经理指数
	IndexMap[PMIName] = make(map[string]string)
	IndexMap[PMIName]["TargetNameEN"] = "PMI"
	IndexMap[PMIName]["TargetCode"] = "HG00020"
	IndexMap[PMIName]["DataSourceCode"] = "stat"
	IndexMap[PMIName]["DataSourceName"] = "国家统计局"
	IndexMap[PMIName]["SourceTargetCode"] = "A0B0101"
	IndexMap[PMIName]["IsQuantity"] = "Y"
	IndexMap[PMIName]["UnitType"] = UnitTypeP
	IndexMap[PMIName]["UnitName"] = UnitNameP
	IndexMap[PMIName]["PeriodType"] = PeriodTypeMonth
	IndexMap[PMIName]["PeriodName"] = PeriodNameMonth
	IndexMap[PMIName]["StartYear"] = "2005"
	IndexMap[PMIName]["StartSeason"] = ""
	IndexMap[PMIName]["StartMonth"] = "01"
	IndexMap[PMIName]["DateType"] = Monthly
	IndexMap[PMIName]["AreaType"] = Marco

	// 工业生产者出厂价格指数(上年=100)
	//IndexMap[PPIName] = make(map[string]string)
	//IndexMap[PPIName]["TargetCode"] = "HG00021"
	//IndexMap[PPIName]["DataSourceCode"] = "stat"
	//IndexMap[PPIName]["DataSourceName"] = "国家统计局"
	//IndexMap[PPIName]["SourceTargetCode"] = "A090105"
	//IndexMap[PPIName]["IsQuantity"] = "Y"
	//IndexMap[PPIName]["UnitType"] = ""
	//IndexMap[PPIName]["UnitName"] = ""
	//IndexMap[PPIName]["PeriodType"] = PeriodTypeYear
	//IndexMap[PPIName]["PeriodName"] = PeriodNameYear
	//IndexMap[PPIName]["StartYear"] = "1978"
	//IndexMap[PPIName]["StartSeason"] = ""
	//IndexMap[PPIName]["StartMonth"] = ""
	//IndexMap[PPIName]["DateType"] = Annual
	//IndexMap[PPIName]["AreaType"] = Marco

	// 工业生产者出厂价格指数(上月=100)
	//IndexMap[PPI2Name] = make(map[string]string)
	//IndexMap[PPI2Name]["TargetCode"] = "HG00022"
	//IndexMap[PPI2Name]["DataSourceCode"] = "stat"
	//IndexMap[PPI2Name]["DataSourceName"] = "国家统计局"
	//IndexMap[PPI2Name]["SourceTargetCode"] = "A01080701"
	//IndexMap[PPI2Name]["IsQuantity"] = "Y"
	//IndexMap[PPI2Name]["UnitType"] = ""
	//IndexMap[PPI2Name]["UnitName"] = ""
	//IndexMap[PPI2Name]["PeriodType"] = PeriodTypeMonth
	//IndexMap[PPI2Name]["PeriodName"] = PeriodNameMonth
	//IndexMap[PPI2Name]["StartYear"] = "2011"
	//IndexMap[PPI2Name]["StartSeason"] = ""
	//IndexMap[PPI2Name]["StartMonth"] = "01"
	//IndexMap[PPI2Name]["DateType"] = Monthly
	//IndexMap[PPI2Name]["AreaType"] = Marco

	// 工业生产者出厂价格指数(上年同月=100)
	IndexMap[PPI3Name] = make(map[string]string)
	IndexMap[PPI3Name]["TargetNameEN"] = "PPI月度同比"
	IndexMap[PPI3Name]["TargetCode"] = "HG00023"
	IndexMap[PPI3Name]["DataSourceCode"] = "stat"
	IndexMap[PPI3Name]["DataSourceName"] = "国家统计局"
	IndexMap[PPI3Name]["SourceTargetCode"] = "A01080101"
	IndexMap[PPI3Name]["IsQuantity"] = "Y"
	IndexMap[PPI3Name]["UnitType"] = ""
	IndexMap[PPI3Name]["UnitName"] = ""
	IndexMap[PPI3Name]["PeriodType"] = PeriodTypeMonth
	IndexMap[PPI3Name]["PeriodName"] = PeriodNameMonth
	IndexMap[PPI3Name]["StartYear"] = "1996"
	IndexMap[PPI3Name]["StartSeason"] = ""
	IndexMap[PPI3Name]["StartMonth"] = "10"
	IndexMap[PPI3Name]["DateType"] = Monthly
	IndexMap[PPI3Name]["AreaType"] = Marco

	// 工业生产者出厂价格指数(上年同期=100)
	//IndexMap[PPI4Name] = make(map[string]string)
	//IndexMap[PPI4Name]["TargetCode"] = "HG00024"
	//IndexMap[PPI4Name]["DataSourceCode"] = "stat"
	//IndexMap[PPI4Name]["DataSourceName"] = "国家统计局"
	//IndexMap[PPI4Name]["SourceTargetCode"] = "A01080401"
	//IndexMap[PPI4Name]["IsQuantity"] = "Y"
	//IndexMap[PPI4Name]["UnitType"] = ""
	//IndexMap[PPI4Name]["UnitName"] = ""
	//IndexMap[PPI4Name]["PeriodType"] = PeriodTypeMonth
	//IndexMap[PPI4Name]["PeriodName"] = PeriodNameMonth
	//IndexMap[PPI4Name]["StartYear"] = "2011"
	//IndexMap[PPI4Name]["StartSeason"] = ""
	//IndexMap[PPI4Name]["StartMonth"] = "01"
	//IndexMap[PPI4Name]["DateType"] = Monthly
	//IndexMap[PPI4Name]["AreaType"] = Marco

	// 居民消费水平
	IndexMap[RCLName] = make(map[string]string)
	IndexMap[RCLName]["TargetNameEN"] = "RCL"
	IndexMap[RCLName]["TargetCode"] = "HG00025"
	IndexMap[RCLName]["DataSourceCode"] = "stat"
	IndexMap[RCLName]["DataSourceName"] = "国家统计局"
	IndexMap[RCLName]["SourceTargetCode"] = "A020501"
	IndexMap[RCLName]["IsQuantity"] = "Y"
	IndexMap[RCLName]["UnitType"] = UnitTypeY
	IndexMap[RCLName]["UnitName"] = UnitNameY
	IndexMap[RCLName]["PeriodType"] = PeriodTypeYear
	IndexMap[RCLName]["PeriodName"] = PeriodNameYear
	IndexMap[RCLName]["StartYear"] = "1992"
	IndexMap[RCLName]["StartSeason"] = ""
	IndexMap[RCLName]["StartMonth"] = ""
	IndexMap[RCLName]["DateType"] = Annual
	IndexMap[RCLName]["AreaType"] = Province

	// 社会消费品零售总额
	//IndexMap[SCGName] = make(map[string]string)
	//IndexMap[SCGName]["TargetCode"] = "HG00026"
	//IndexMap[SCGName]["DataSourceCode"] = "stat"
	//IndexMap[SCGName]["DataSourceName"] = "国家统计局"
	//IndexMap[SCGName]["SourceTargetCode"] = "A0H01"
	//IndexMap[SCGName]["IsQuantity"] = "Y"
	//IndexMap[SCGName]["UnitType"] = UnitTypeE
	//IndexMap[SCGName]["UnitName"] = UnitNameE
	//IndexMap[SCGName]["PeriodType"] = PeriodTypeYear
	//IndexMap[SCGName]["PeriodName"] = PeriodNameYear
	//IndexMap[SCGName]["StartYear"] = "1952"
	//IndexMap[SCGName]["StartSeason"] = ""
	//IndexMap[SCGName]["StartMonth"] = ""
	//IndexMap[SCGName]["DateType"] = Annual
	//IndexMap[SCGName]["AreaType"] = Marco

	// 社会消费品零售总额_当期值
	//IndexMap[SCG1Name] = make(map[string]string)
	//IndexMap[SCG1Name]["TargetCode"] = "HG00027"
	//IndexMap[SCG1Name]["DataSourceCode"] = "stat"
	//IndexMap[SCG1Name]["DataSourceName"] = "国家统计局"
	//IndexMap[SCG1Name]["SourceTargetCode"] = "A070101"
	//IndexMap[SCG1Name]["IsQuantity"] = "Y"
	//IndexMap[SCG1Name]["UnitType"] = UnitTypeE
	//IndexMap[SCG1Name]["UnitName"] = UnitNameE
	//IndexMap[SCG1Name]["PeriodType"] = PeriodTypeMonth
	//IndexMap[SCG1Name]["PeriodName"] = PeriodNameMonth
	//IndexMap[SCG1Name]["StartYear"] = "1984"
	//IndexMap[SCG1Name]["StartSeason"] = ""
	//IndexMap[SCG1Name]["StartMonth"] = "01"
	//IndexMap[SCG1Name]["DateType"] = Monthly
	//IndexMap[SCG1Name]["AreaType"] = Marco

	// 社会消费品零售总额_累计值
	//IndexMap[SCG2Name] = make(map[string]string)
	//IndexMap[SCG2Name]["TargetCode"] = "HG00028"
	//IndexMap[SCG2Name]["DataSourceCode"] = "stat"
	//IndexMap[SCG2Name]["DataSourceName"] = "国家统计局"
	//IndexMap[SCG2Name]["SourceTargetCode"] = "A070102"
	//IndexMap[SCG2Name]["IsQuantity"] = "Y"
	//IndexMap[SCG2Name]["UnitType"] = UnitTypeE
	//IndexMap[SCG2Name]["UnitName"] = UnitNameE
	//IndexMap[SCG2Name]["PeriodType"] = PeriodTypeMonth
	//IndexMap[SCG2Name]["PeriodName"] = PeriodNameMonth
	//IndexMap[SCG2Name]["StartYear"] = "2000"
	//IndexMap[SCG2Name]["StartSeason"] = ""
	//IndexMap[SCG2Name]["StartMonth"] = "01"
	//IndexMap[SCG2Name]["DateType"] = Monthly
	//IndexMap[SCG2Name]["AreaType"] = Marco

	// 社会消费品零售总额_同比增长
	IndexMap[SCG3Name] = make(map[string]string)
	IndexMap[SCG3Name]["TargetNameEN"] = "SCG同比"
	IndexMap[SCG3Name]["TargetCode"] = "HG00029"
	IndexMap[SCG3Name]["DataSourceCode"] = "stat"
	IndexMap[SCG3Name]["DataSourceName"] = "国家统计局"
	IndexMap[SCG3Name]["SourceTargetCode"] = "A070103"
	IndexMap[SCG3Name]["IsQuantity"] = "Y"
	IndexMap[SCG3Name]["UnitType"] = UnitTypeP
	IndexMap[SCG3Name]["UnitName"] = UnitNameP
	IndexMap[SCG3Name]["PeriodType"] = PeriodTypeMonth
	IndexMap[SCG3Name]["PeriodName"] = PeriodNameMonth
	IndexMap[SCG3Name]["StartYear"] = "2000"
	IndexMap[SCG3Name]["StartSeason"] = ""
	IndexMap[SCG3Name]["StartMonth"] = "01"
	IndexMap[SCG3Name]["DateType"] = Monthly
	IndexMap[SCG3Name]["AreaType"] = Marco

	// 社会消费品零售总额_累计增长
	//IndexMap[SCG4Name] = make(map[string]string)
	//IndexMap[SCG4Name]["TargetCode"] = "HG00030"
	//IndexMap[SCG4Name]["DataSourceCode"] = "stat"
	//IndexMap[SCG4Name]["DataSourceName"] = "国家统计局"
	//IndexMap[SCG4Name]["SourceTargetCode"] = "A070104"
	//IndexMap[SCG4Name]["IsQuantity"] = "Y"
	//IndexMap[SCG4Name]["UnitType"] = UnitTypeP
	//IndexMap[SCG4Name]["UnitName"] = UnitNameP
	//IndexMap[SCG4Name]["PeriodType"] = PeriodTypeMonth
	//IndexMap[SCG4Name]["PeriodName"] = PeriodNameMonth
	//IndexMap[SCG4Name]["StartYear"] = "2000"
	//IndexMap[SCG4Name]["StartSeason"] = ""
	//IndexMap[SCG4Name]["StartMonth"] = "01"
	//IndexMap[SCG4Name]["DateType"] = Monthly
	//IndexMap[SCG4Name]["AreaType"] = Marco

	// 城镇居民人均可支配收入
	//IndexMap[URIName] = make(map[string]string)
	//IndexMap[URIName]["outerCode"] = "A0A0103"
	//IndexMap[URIName]["innerCode"] = "HG00031"
	//IndexMap[URIName]["startYear"] = "2013"

	// 城镇居民人均可支配收入_同比增长
	//IndexMap[URI1Name] = make(map[string]string)
	//IndexMap[URI1Name]["outerCode"] = "A0A0104"
	//IndexMap[URI1Name]["innerCode"] = "HG00032"
	//IndexMap[URI1Name]["startYear"] = "2014"

	// 城镇居民人均可支配收入_累计值
	//IndexMap[URI2Name] = make(map[string]string)
	//IndexMap[URI2Name]["outerCode"] = "A050201"
	//IndexMap[URI2Name]["innerCode"] = "HG00033"
	//IndexMap[URI2Name]["startYear"] = "2013"

	// 城镇居民人均可支配收入_累计增长
	//IndexMap[URI3Name] = make(map[string]string)
	//IndexMap[URI3Name]["outerCode"] = "A050202"
	//IndexMap[URI3Name]["innerCode"] = "HG00034"
	//IndexMap[URI3Name]["startYear"] = "2014"
}

func init() {
	setUpIndexMap()
}
