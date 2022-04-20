package target

// ExcelConfig Excel配置
type ExcelConfig struct {
	Enable                 bool   // 是否启用
	TargetCode             string // 指标编码
	TargetName             string // 指标名称
	TargetNameSpider       string // 指标名称_爬虫程序
	TargetNameEn           string // 指标名称英文
	DataSourceCode         string // 数据源代码
	DataSourceName         string // 数据源名称
	SourceTargetCodeSpider string // 来源系统指标编号_爬虫程序
	SourceTargetCode       string // 来源系统指标编号
	IsQuantity             string // 是否定量
	UnitType               string // 计量单位类型
	UnitName               string // 计量单位名称
	PeriodType             string // 指标期间类型
	PeriodName             string // 指标期间名称
	SpiderTime             string // 爬虫时间
	Adapter                string // 适配
}
