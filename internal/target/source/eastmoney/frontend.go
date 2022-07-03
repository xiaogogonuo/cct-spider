package eastmoney

// eastMoneyIndustryTarget 东方财富行业指标接口返回的数据结构
type eastMoneyIndustryTarget struct {
	Success bool `json:"success"`
	Result  struct {
		Data []struct {
			ReportDate     string  `json:"REPORT_DATE"`     // 日期
			IndicatorValue float64 `json:"INDICATOR_VALUE"` // 最新值
			ChangeRate     float64 `json:"CHANGE_RATE"`     // 涨跌幅
			ChangeRate3M   float64 `json:"CHANGERATE_3M"`   // 近3月涨跌幅
			ChangeRate6M   float64 `json:"CHANGERATE_6M"`   // 近6月涨跌幅
			ChangeRate1Y   float64 `json:"CHANGERATE_1Y"`   // 近1年涨跌幅
			ChangeRate2Y   float64 `json:"CHANGERATE_2Y"`   // 近2年涨跌幅
			ChangeRate3Y   float64 `json:"CHANGERATE_3Y"`   // 近3年涨跌幅
		} `json:"data"`
	} `json:"result"`
}

// eastMoneyQiHuoTarget 东方财富期货指标接口返回的数据结构
type eastMoneyQiHuoTarget struct {
	Qt struct {
		P     float64 `json:"p"`     // 最新价
		ZDE   float64 `json:"zde"`   // 涨跌额
		ZDF   float64 `json:"zdf"`   // 涨跌幅
		H     float64 `json:"h"`     // 最高
		L     float64 `json:"l"`     // 最低
		QRSPJ float64 `json:"qrspj"` // 昨收
		UTime int64   `json:"utime"` // 更新日期
	} `json:"qt"`
}

// eastMoneyCHN10 东方财富中债10年期国债到期收益率接口返回的数据结构
type eastMoneyCHN10 struct {
	Success bool `json:"success"`
	Result  struct {
		Data []struct {
			SolarDate   string  `json:"SOLAR_DATE"`  // 日期
			EMM00166462 float64 `json:"EMM00166462"` // 5年
			EMM00166466 float64 `json:"EMM00166466"` // 10年
			EMM00166469 float64 `json:"EMM00166469"` // 30年
		} `json:"data"`
	} `json:"result"`
}

// eastMoneyGlobalTarget 东方财富全球指数接口返回的数据结构
type eastMoneyGlobalTarget struct {
	Data struct {
		F43  float64 `json:"f43"`  // 现价
		F44  float64 `json:"f44"`  // 最高
		F45  float64 `json:"f45"`  // 最低
		F46  float64 `json:"f46"`  // 今开(暂时不用)
		F60  float64 `json:"f60"`  // 昨收
		F86  int64   `json:"f86"`  // 更新时间
		F169 float64 `json:"f169"` // 涨跌
		F170 float64 `json:"f170"` // 涨跌幅
	} `json:"data"`
}

// eastMoneyBOR 东方财富拆借利率接口返回的数据结构
type eastMoneyBOR struct {
	Result struct{
		Data []struct{
			REPORT_DATE string `json:"REPORT_DATE"`
			REPORT_PERIOD string `json:"REPORT_PERIOD"`
			IR_RATE float64 `json:"IR_RATE"`
			CHANGE_RATE float64 `json:"CHANGE_RATE"`
			INDICATOR_ID string `json:"INDICATOR_ID"`
			LATEST_RECORD float64 `json:"LATEST_RECORD"`
			MARKET string `json:"MARKET"`
			MARKET_CODE string `json:"MARKET_CODE"`
			CURRENCY string `json:"CURRENCY"`
			CURRENCY_CODE string `json:"CURRENCY_CODE"`
		} `json:"data"`
	} `json:"result"`
	Success bool `json:"success"`
}
