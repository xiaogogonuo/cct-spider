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
			SumSame             float64 `json:"SUM_SAME"`               // 国内生产总值同比增长：2.5  需加%
			FirstSame           float64 `json:"FIRST_SAME"`             // 第一产业同比增长：5       需加%
			SecondSame          float64 `json:"SECOND_SAME"`            // 第二产业同比增长：3.2     需加%
			ThirdSame           float64 `json:"THIRD_SAME"`             // 第三产业同比增长：1.8     需加%
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
			RetailTotalSame       float64 `json:"RETAIL_TOTAL_SAME"`       // 当月同比增长：-5.9,
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
			BasicCurrencySame       float64 `json:"BASIC_CURRENCY_SAME"`       // 货币和准货币(M2)-同比增长：11.8,
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
