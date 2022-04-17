package target

// NewTarget 新建指标爬虫任务
func NewTarget() {
	ecs := ExcelSetUp()
	for _, ec := range ecs {
		if !ec.Enable {
			continue
		}
		switch ec.Adapter {
		case "EastMoneyEconomicTarget": // 爬取"东方财富"网站的`宏观指标`
			//responses := eastmoney.SpiderEastMoneyEconomicTarget(ec.SourceTargetCodeSpider, ec.TargetCode)
			//data := Generator(ec, responses, true)
			//fmt.Println(data)
		case "EastMoneyEconomicTargetBOR": // 爬取"东方财富"网站的`宏观指标-上海银行间同业拆放利率隔夜`
			// pages：-1 代表爬完所有页数，适用于第一次写数据库
			// pages：1  代表后期每次只爬一页，做增量爬取
			//responses := eastmoney.SpiderEastMoneyEconomicTargetBOR(ec.SourceTargetCodeSpider, 1)
			//data := Generator(ec, responses, false)
			//fmt.Println(data)
		case "EastMoneyEconomicTargetCHN10": // 爬取"东方财富"网站的`宏观指标-中债10年期国债到期收益率`
			// pages：-1 代表爬完所有页数，适用于第一次写数据库
			// pages：1  代表后期每次只爬一页，做增量爬取
			//responses := eastmoney.SpiderEastMoneyEconomicTargetCHN10(1)
			//data := Generator(ec, responses, false)
			//fmt.Println(data)
		case "EastMoneyGlobalTarget": // 爬取"东方财富"网站的`全球指数`
			//responses := eastmoney.SpiderEastMoneyGlobalTarget(ec.SourceTargetCodeSpider)
			//data := Generator(ec, responses, true)
			//fmt.Println(data)
		case "EastMoneyIndustryTarget": // 爬取"东方财富"网站的`行业指标`
			//responses := eastmoney.SpiderEastMoneyIndustryTarget(ec.SourceTargetCodeSpider)
			//data := Generator(ec, responses, true)
			//fmt.Println(data)
		case "EastMoneyQiHuoTarget": // 爬取"东方财富"网站的`期货指标`
			//responses := eastmoney.SpiderEastMoneyQiHuoTarget(ec.SourceTargetCodeSpider)
			//data := Generator(ec, responses, false)
			//fmt.Println(data)
		case "SinaTargetForex": // 爬取"新浪财经"网站的`外汇指标`
			//responses := sina.SpiderSinaTargetForex(ec.SourceTargetCodeSpider)
			//data := Generator(ec, responses, false)
			//fmt.Println(data)
		case "FxExchangeTarget": // 爬取"汇通财经"网站的外汇、债券、原油、期货、外盘、汇率等指标
			//responses := fx678.SpiderFxExchangeTarget(ec.SourceTargetCodeSpider, ec.TargetNameSpider)
			//data := Generator(ec, responses, false)
			//fmt.Println(data)
		case "SCITargetCOI": // 爬取卓创资讯的原油价格指数
			//responses := sci.SpiderSCITargetCOI()
			//data := Generator(ec, responses, false)
			//fmt.Println(data)
		case "SCITargetPII": // 爬取卓创资讯的造纸行业价格指数
			//responses := sci.SpiderSCITargetPII()
			//data := Generator(ec, responses, false)
			//fmt.Println(data)
		}
	}
}

// 对实时数据指标，发送给Java服务器的指标值格式如下：
// 最新价,涨跌,涨跌幅,最高,最低,昨收,更新时间
