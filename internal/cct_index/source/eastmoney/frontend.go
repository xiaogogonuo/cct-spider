package eastmoney

// 东方财富指标接口返回模型

// GDP 国内生产总值
type GDP struct {
	Result struct {
		Data []struct {
			ReportDate          string  `json:"REPORT_DATE"`            // 报告日期："2022-06-01 00:00:00"
			Time                string  `json:"TIME"`                   // 季度："2022年第1-2季度"
			DomesticProductBase float64 `json:"DOMESTICL_PRODUCT_BASE"` // 国内生产总值绝对值(亿元)：562641.6
			FirstProductBase    float64 `json:"FIRST_PRODUCT_BASE"`     // 第一产业绝对值(亿元)：29137.2
			SecondProductBase   float64 `json:"SECOND_PRODUCT_BASE"`    // 第二产业绝对值(亿元)：228636.4
			ThirdProductBase    float64 `json:"THIRD_PRODUCT_BASE"`     // 第三产业绝对值(亿元)：304868
			SumSame             float64 `json:"SUM_SAME"`               // 国内生产总值同比增长：2.5
			FirstSame           float64 `json:"FIRST_SAME"`             // 第一产业同比增长：5
			SecondSame          float64 `json:"SECOND_SAME"`            // 第二产业同比增长：3.2
			ThirdSame           float64 `json:"THIRD_SAME"`             // 第三产业同比增长：1.8
		} `json:"data"`
	} `json:"result"`
}

// IAV 工业增加值
type IAV struct {
	Result struct {
		Data []struct {
			ReportDate     string  `json:"REPORT_DATE"`     // 报告日期："2022-11-01 00:00:00"
			Time           string  `json:"TIME"`            // 月份："2022年11月份"
			BaseSame       float64 `json:"BASE_SAME"`       // 同比增长：2.2  可能为null
			BaseAccumulate float64 `json:"BASE_ACCUMULATE"` // 累计增长：3.8
		} `json:"data"`
	} `json:"result"`
}

// XFP 社会消费品零售总额
type XFP struct {
	Result struct {
		Data []struct {
			ReportDate            string  `json:"REPORT_DATE"`             // 报告日期："2022-11-01 00:00:00"
			Time                  string  `json:"TIME"`                    // 月份："2022年11月份"
			RetailTotal           float64 `json:"RETAIL_TOTAL"`            // 当月(亿元)：38615
			RetailTotalSame       float64 `json:"RETAIL_TOTAL_SAME"`       // 当月同比增长：-5.9
			RetailTotalSequential float64 `json:"RETAIL_TOTAL_SEQUENTIAL"` // 环比增长：-4.11214025
			RetailTotalAccumulate float64 `json:"RETAIL_TOTAL_ACCUMULATE"` // 累计(亿元)	：399190
			RetailAccumulateSame  float64 `json:"RETAIL_ACCUMULATE_SAME"`  // 累计同比增长：-0.1
		} `json:"data"`
	} `json:"result"`
}

// M2 货币和准货币(M2)供应量
type M2 struct {
	Result struct {
		Data []struct {
			ReportDate              string  `json:"REPORT_DATE"`               // 报告日期："2022-11-01 00:00:00"
			Time                    string  `json:"TIME"`                      // 月份："2022年11月份"
			BasicCurrency           float64 `json:"BASIC_CURRENCY"`            // 货币和准货币(M2)-数量(亿元)：2664300
			BasicCurrencySame       float64 `json:"BASIC_CURRENCY_SAME"`       // 货币和准货币(M2)-同比增长：11.8
			BasicCurrencySequential float64 `json:"BASIC_CURRENCY_SEQUENTIAL"` // 货币和准货币(M2)-环比增长：0.65324762
			Currency                float64 `json:"CURRENCY"`                  // 货币(M1)-数量(亿元)：671700
			CurrencySame            float64 `json:"CURRENCY_SAME"`             // 货币(M1)-同比增长：3.7
			CurrencySequential      float64 `json:"CURRENCY_SEQUENTIAL"`       // 货币(M1)-环比增长：0.69821477
			FreeCash                float64 `json:"FREE_CASH"`                 // 流通中的现金(M0)-数量(亿元)：104700
			FreeCashSame            float64 `json:"FREE_CASH_SAME"`            // 流通中的现金(M0)-同比增长：15.3
			FreeCashSequential      float64 `json:"FREE_CASH_SEQUENTIAL"`      // 流通中的现金(M0)-环比增长：4.97280332
		} `json:"data"`
	} `json:"result"`
}

// CPI 居民消费价格指数
type CPI struct {
	Result struct {
		Data []struct {
			ReportDate         string  `json:"REPORT_DATE"`         // 报告日期："2022-12-01 00:00:00"
			Time               string  `json:"TIME"`                // 月份："2022年12月份"
			NationalSame       float64 `json:"NATIONAL_SAME"`       // 全国-同比增长：1.8
			NationalBase       float64 `json:"NATIONAL_BASE"`       // 全国-当月：101.8
			NationalSequential float64 `json:"NATIONAL_SEQUENTIAL"` // 全国-环比增长：0
			NationalAccumulate float64 `json:"NATIONAL_ACCUMULATE"` // 全国-累计：102
			CitySame           float64 `json:"CITY_SAME"`           // 城市-同比增长：1.8
			CityBase           float64 `json:"CITY_BASE"`           // 城市-当月：101.8
			CitySequential     float64 `json:"CITY_SEQUENTIAL"`     // 城市-环比增长：0
			CityAccumulate     float64 `json:"CITY_ACCUMULATE"`     // 城市-累计：102
			RuralSame          float64 `json:"RURAL_SAME"`          // 农村-同比增长：1.8
			RuralBase          float64 `json:"RURAL_BASE"`          // 农村-当月：101.8
			RuralSequential    float64 `json:"RURAL_SEQUENTIAL"`    // 农村-环比增长：-0.2
			RuralAccumulate    float64 `json:"RURAL_ACCUMULATE"`    // 农村-累计：102
		} `json:"data"`
	} `json:"result"`
}

// PMI 采购经理人指数
type PMI struct {
	Result struct {
		Data []struct {
			ReportDate string  `json:"REPORT_DATE"` // 报告日期："2022-12-01 00:00:00"
			Time       string  `json:"TIME"`        // 月份："2022年12月份"
			MakeIndex  float64 `json:"MAKE_INDEX"`  // 制造业-指数	：47
			MakeSame   float64 `json:"MAKE_SAME"`   // 制造业-同比增长：-6.56063618
			NMakeIndex float64 `json:"NMAKE_INDEX"` // 非制造业-指数：41.6
			NMakeSame  float64 `json:"NMAKE_SAME"`  // 非制造业-同比增长：-21.0626186
		} `json:"data"`
	} `json:"result"`
}

// PPI 工业品出厂价格指数
type PPI struct {
	Result struct {
		Data []struct {
			ReportDate     string  `json:"REPORT_DATE"`     // 报告日期："2022-12-01 00:00:00"
			Time           string  `json:"TIME"`            // 月份："2022年12月份"
			Base           float64 `json:"BASE"`            // 当月	：99.3
			BaseSame       float64 `json:"BASE_SAME"`       // 当月同比增长：-0.7
			BaseAccumulate float64 `json:"BASE_ACCUMULATE"` // 累计：104.1
		} `json:"data"`
	} `json:"result"`
}

// ZBJ 存款准备金率
type ZBJ struct {
	Result struct {
		Data []struct {
			ReportDate     string  `json:"REPORT_DATE"`      // 报告日期："2022-11-25 00:00:00"
			PublishDate    string  `json:"PUBLISH_DATE"`     // 公布时间："2022年11月25日"
			TradeDate      string  `json:"TRADE_DATE"`       // 生效时间："2022年12月05日"
			InterestRateBB float64 `json:"INTEREST_RATE_BB"` // 大型金融机构-调整前：11.25
			InterestRateBA float64 `json:"INTEREST_RATE_BA"` // 大型金融机构-调整后：11
			ChangeRateB    float64 `json:"CHANGE_RATE_B"`    // 大型金融机构-调整幅度：-0.25
			InterestRateSB float64 `json:"INTEREST_RATE_SB"` // 中小金融机构-调整前：8.25
			InterestRateSA float64 `json:"INTEREST_RATE_SA"` // 中小金融机构-调整前：8
			ChangeRateS    float64 `json:"CHANGE_RATE_S"`    // 中小金融机构-调整幅度：-0.25
			NextSHRate     float64 `json:"NEXT_SH_RATE"`     // 消息公布次日指数涨跌-上证：-0.746195658552
			NextSZRate     float64 `json:"NEXT_SZ_RATE"`     // 消息公布次日指数涨跌-深证：-0.689504468976
			Remark         string  `json:"REMARK"`           // 备注：
		} `json:"data"`
	} `json:"result"`
}

// JCK 海关进出口
type JCK struct {
	Result struct {
		Data []struct {
			ReportDate           string  `json:"REPORT_DATE"`            // 报告日期："2022-12-01 00:00:00"
			Time                 string  `json:"TIME"`                   // 月   份："2022年12月份"
			ExitBase             float64 `json:"EXIT_BASE"`              // 当月出口额-金额(亿美元)：306078800
			ImportBase           float64 `json:"IMPORT_BASE"`            // 当月进口额-金额(亿美元)：228066200
			ExitBaseSame         float64 `json:"EXIT_BASE_SAME"`         // 当月出口额-同比增长：-9.9
			ImportBaseSame       float64 `json:"IMPORT_BASE_SAME"`       // 当月进口额-同比增长：-7.5
			ExitBaseSequential   float64 `json:"EXIT_BASE_SEQUENTIAL"`   // 当月出口额-环比增长：3.57969629
			ImportBaseSequential float64 `json:"IMPORT_BASE_SEQUENTIAL"` // 当月进口额-环比增长：0.80158083
			ExitAccumulate       float64 `json:"EXIT_ACCUMULATE"`        // 累计出口额-金额(亿美元)：3593601500
			ImportAccumulate     float64 `json:"IMPORT_ACCUMULATE"`      // 累计进口额-金额(亿美元)：2715998800
			ExitAccumulateSame   float64 `json:"EXIT_ACCUMULATE_SAME"`   // 累计出口额-同比增长：7
			ImportAccumulateSame float64 `json:"IMPORT_ACCUMULATE_SAME"` // 累计进口额-同比增长：1.1
		} `json:"data"`
	} `json:"result"`
}

// LL 利率调整
type LL struct {
	Result struct {
		Data []struct {
			ReportDate    string  `json:"REPORT_DATE"`     // 生效时间："2015-10-24 00:00:00"
			PublishDate   string  `json:"PUBLISH_DATE"`    // 公布时间："2015-10-23 00:00:00"
			DepositRateBB float64 `json:"DEPOSIT_RATE_BB"` // 存款基准利率-调整前：1.75
			DepositRateBA float64 `json:"DEPOSIT_RATE_BA"` // 存款基准利率-调整后：1.5
			DepositRateB  float64 `json:"DEPOSIT_RATE_B"`  // 存款基准利率-调整幅度：-0.25
			LoadRateSB    float64 `json:"LOAN_RATE_SB"`    // 贷款基准利率-调整前：4.6
			LoadRateSA    float64 `json:"LOAN_RATE_SA"`    // 贷款基准利率-调整后：4.35
			LoadRateS     float64 `json:"LOAN_RATE_S"`     // 贷款基准利率-调整幅度：-0.25
			NextSHRate    float64 `json:"NEXT_SH_RATE"`    // 消息公布次日指数涨跌-上证：1.010410302187
			NextSZRate    float64 `json:"NEXT_SZ_RATE"`    // 消息公布次日指数涨跌-深证：1.581269520584
		} `json:"data"`
	} `json:"result"`
}

// WH 外汇和黄金储备
type WH struct {
	Result struct {
		Data []struct {
			ReportDate             string  `json:"REPORT_DATE"`              // 报告日期："2022-12-01 00:00:00"
			Time                   string  `json:"TIME"`                     // 月份："2022年12月份"
			GoldReserves           float64 `json:"GOLD_RESERVES"`            // 黄金储备(亿美元)-数值：1172.35
			GoldReservesSame       float64 `json:"GOLD_RESERVES_SAME"`       // 黄金储备(亿美元)-同比：3.63314917
			GoldReservesSequential float64 `json:"GOLD_RESERVES_SEQUENTIAL"` // 黄金储备(亿美元)-环比：5.00223914
			Forex                  float64 `json:"FOREX"`                    // 国家外汇储备(亿美元)-数值：31276.91
			ForexSame              float64 `json:"FOREX_SAME"`               // 国家外汇储备(亿美元)-同比：-3.76826907
			ForexSequential        float64 `json:"FOREX_SEQUENTIAL"`         // 国家外汇储备(亿美元)-环比：0.32728274
		} `json:"data"`
	} `json:"result"`
}

// LendingRate 拆借利率
type LendingRate struct {
	Result struct {
		Data []struct {
			ReportDate string  `json:"REPORT_DATE"` // 日期："2023-01-16 00:00:00"
			IrRate     float64 `json:"IR_RATE"`     // 利率(%)：1.568
			ChangeRate float64 `json:"CHANGE_RATE"` // 涨跌(BP)：32.6
		} `json:"data"`
	} `json:"result"`
}

// Industry 行业指数
type Industry struct {
	Result struct {
		Data []struct {
			ReportDate     string  `json:"REPORT_DATE"`     // 日期："2023-01-16 00:00:00"
			IndicatorValue float64 `json:"INDICATOR_VALUE"` // 最新值：946
			ChangeRate     float64 `json:"CHANGE_RATE"`     // 涨跌涨跌幅：0
			ChangeRate3M   float64 `json:"CHANGERATE_3M"`   // 近3月涨跌幅：-48.53101197
			ChangeRate6M   float64 `json:"CHANGERATE_6M"`   // 近6月涨跌幅：-56
			ChangeRate1Y   float64 `json:"CHANGERATE_1Y"`   // 近1年涨跌幅：-46.37188209
			ChangeRate2Y   float64 `json:"CHANGERATE_2Y"`   // 近2年涨跌幅：-46.06613455
			ChangeRate3Y   float64 `json:"CHANGERATE_3Y"`   // 近3年涨跌幅：23.17708333
		} `json:"data"`
	} `json:"result"`
}

// Global 全球指数
// 字段F43、F44、F45、F46、F60、F169、F170的值都需要除100
// 字段F86是unix时间戳的表示形式
type Global struct {
	Data struct {
		F43  float64 `json:"f43"`  // 现价：19874
		F44  float64 `json:"f44"`  // 最高：19878
		F45  float64 `json:"f45"`  // 最低：19874
		F46  float64 `json:"f46"`  // 今开：19878 (暂时不用)
		F60  float64 `json:"f60"`  // 昨收：19877
		F86  int64   `json:"f86"`  // 更新时间：1673942372
		F169 float64 `json:"f169"` // 涨跌额：-3
		F170 float64 `json:"f170"` // 涨跌幅：-2   单位是%
	} `json:"data"`
}

// NationalDebt 国债
type NationalDebt struct {
	Result struct {
		Data []struct {
			SolarDate   string  `json:"SOLAR_DATE"`  // 日期："2023-01-17 00:00:00"
			EMM00588704 float64 `json:"EMM00588704"` // 中国：国债收益率 - 2年：2.4013
			EMM00166462 float64 `json:"EMM00166462"` // 中国：国债收益率 - 5年：2.7346
			EMM00166466 float64 `json:"EMM00166466"` // 中国：国债收益率 - 10年：2.9017
			EMM00166469 float64 `json:"EMM00166469"` // 中国：国债收益率 - 30年：3.2531
			EMM01276014 float64 `json:"EMM01276014"` // 中国：国债收益率 - 30年-2年：0.5004
			EMM00000024 float64 `json:"EMM00000024"` // 中国GDP年增率(%)：null
			EMG00001306 float64 `json:"EMG00001306"` // 美国：国债收益率 - 2年：4.18
			EMG00001308 float64 `json:"EMG00001308"` // 美国：国债收益率 - 5年：3.6
			EMG00001310 float64 `json:"EMG00001310"` // 美国：国债收益率 - 10年：3.53
			EMG00001312 float64 `json:"EMG00001312"` // 美国：国债收益率 - 30年：3.64
			EMG01339436 float64 `json:"EMG01339436"` // 美国：国债收益率 - 10-2年：-0.65
			EMG00159635 float64 `json:"EMG00159635"` // 美国GDP年增率(%)：null
		} `json:"data"`
	} `json:"result"`
}
