package poster

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/api"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/model"
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
	"github.com/xiaogogonuo/cct-spider/pkg/mail"
	"strings"
)

const (
	user = "xiaogogonuo@163.com"
	pass = "JDAOREDDCXYMAXXQ"
	nick = "lujiawei"
)

var (
	receivers = []string{"xiaogogonuo@163.com"}
	subject   = "城通爬虫系统警报"
)

var e = mail.NewEmail163(user, pass, nick)

func GDPTBPoster() {
	body := `
<html>
    <body>
		<div>数据源名称：东方财富</div>
		<div>指标名称：国内生产总值同比增长</div>
		<div>指标代码：HG00001</div>
		<div>原接口：%s</div>
		<div>返回的数据格式如下：</div>
        <p>
			{
				"version": "d07e8629dcb9b3c9ecc3eecb895f4600",
				"result": {
					"pages": 1,
					"data": [
						{
							"REPORT_DATE": "2022-09-01 00:00:00",
							"TIME": "2022年第1-3季度",
							"DOMESTICL_PRODUCT_BASE": 870269,
							"FIRST_PRODUCT_BASE": 54779.1,
							"SECOND_PRODUCT_BASE": 350189.5,
							"THIRD_PRODUCT_BASE": 465300.4,
							"SUM_SAME": 3,
							"FIRST_SAME": 4.2,
							"SECOND_SAME": 3.9,
							"THIRD_SAME": 2.3
						}
					],
					"count": 1
				},
				"success": true,
				"message": "ok",
				"code": 0
			}
        </p>
		<div>请检查接口的返回数据格式</div>
	</body>
</html>
	`
	body = fmt.Sprintf(body, api.GDPTb)
	if err := e.Send(receivers, subject, body); err != nil {
		logger.Error(err.Error())
	}
}

const template = `
<html>
    <body>
		<div>数据源：%s</div>
		<div>指标名称：%s</div>
		<div>指标代码：%s</div>
		<div>原接口：%s 返回的数据格式应该如下所示：</div>
        <p>
			{
				"version": "434f9c6126190079ad8950890e8546c7",
				"result": {
					"pages": 1,
					"data": [
						%s
					],
					"count": 100
				},
				"success": true,
				"message": "ok",
				"code": 0
			}
        </p>
		<div>请检查接口的返回数据格式是否同上案例一致</div>
	</body>
</html>
`

const (
	IAVResponse = `
		{
			"REPORT_DATE": "2022-11-01 00:00:00",
			"TIME": "2022年11月份",
			"BASE_SAME": 2.2,
			"BASE_ACCUMULATE": 3.8
		}
`

	XFPResponse = `
		{
			"REPORT_DATE": "2022-11-01 00:00:00",
			"TIME": "2022年11月份",
			"RetailTotal": 38615,
			"RetailTotalSame": -5.9,
			"RetailTotalSequential": -4.11214025,
			"RetailTotalAccumulate": 399190,
			"RetailAccumulateSame": -0.1
		}
`

	M2Response = `
		{
			"REPORT_DATE": "2022-12-01 00:00:00",
			"TIME": "2022年12月份",
			"BASIC_CURRENCY": 2664300,
			"BASIC_CURRENCY_SAME": 11.8,
			"BASIC_CURRENCY_SEQUENTIAL": 0.65324762,
			"CURRENCY": 671700,
			"CURRENCY_SAME": 3.7,
			"CURRENCY_SEQUENTIAL": 0.69821477,
			"FREE_CASH": 104700,
			"FREE_CASH_SAME": 15.3,
			"FREE_CASH_SEQUENTIAL": 4.97280332
		}
`

	CPIResponse = `
		{
			"REPORT_DATE"": "2022-12-01 00:00:00",
			"TIME"": "2022年12月份",
			"NATIONAL_SAME":1.8,
			"NATIONAL_BASE":101.8,
			"NATIONAL_SEQUENTIAL":0,
			"NATIONAL_ACCUMULATE":102,
			"CITY_SAME":1.8,
			"CITY_BASE":101.8,
			"CITY_SEQUENTIAL":0,
			"CITY_ACCUMULATE":102,
			"RURAL_SAME":1.8,
			"RURAL_BASE":101.8,
			"RURAL_SEQUENTIAL":-0.2,
			"RURAL_ACCUMULATE":102
		}
`

	PMIResponse = `
		{
			"REPORT_DATE"": "2022-12-01 00:00:00",
			"TIME"": "2022年12月份",
			"MAKE_INDEX":47,
			"MAKE_SAME":-6.56063618,
			"NMAKE_INDEX":41.6,
			"NMAKE_SAME":-21.0626186
		}
`

	PPIResponse = `
		{
			"REPORT_DATE"": "2022-12-01 00:00:00",
			"TIME"": "2022年12月份",
			"BASE":99.3,
			"BASE_SAME":-0.7,
			"BASE_ACCUMULATE":104.1
		}
`

	ZBJResponse = `
		{
			"REPORT_DATE":"2022-11-25 00:00:00",
			"PUBLISH_DATE":"2022年11月25日",
			"TRADE_DATE":"2022年12月05日",	
			"INTEREST_RATE_BB":11.25,
			"INTEREST_RATE_BA":11,
			"CHANGE_RATE_B":-0.25,
			"INTEREST_RATE_SB":8.25,
			"INTEREST_RATE_SA":8,
			"CHANGE_RATE_S":-0.25,
			"NEXT_SH_RATE":-0.746195658552,
			"NEXT_SZ_RATE":-0.689504468976,
			"REMARK":"xxx"
		}
`

	JCKResponse = `
		{
			"REPORT_DATE": "2022-12-01 00:00:00",
			"TIME": "2022年12月份",
			"EXIT_BASE": 306078800,
			"IMPORT_BASE": 228066200,
			"EXIT_BASE_SAME": -9.9,
			"IMPORT_BASE_SAME": -7.5,
			"EXIT_BASE_SEQUENTIAL": 3.57969629,
			"IMPORT_BASE_SEQUENTIAL": 0.80158083,
			"EXIT_ACCUMULATE": 3593601500,
			"IMPORT_ACCUMULATE": 2715998800,
			"EXIT_ACCUMULATE_SAME": 7,
			"IMPORT_ACCUMULATE_SAME": 1.1
		}
`

	LLResponse = `
		{
			"REPORT_DATE": "2015-10-24 00:00:00",
			"PUBLISH_DATE": "2015-10-23 00:00:00",
			"DEPOSIT_RATE_BB": 1.75,
			"DEPOSIT_RATE_BA": 1.5,
			"DEPOSIT_RATE_B": -0.25,
			"LOAN_RATE_SB": 4.6,
			"LOAN_RATE_SA": 4.35,
			"LOAN_RATE_S": -0.25,
			"NEXT_SH_RATE": 1.010410302187,
			"NEXT_SZ_RATE": 1.581269520584
		}
`

	WHResponse = `
		{
			"REPORT_DATE": "2022-12-01 00:00:00",
			"TIME": "2022年12月份",
			"GOLD_RESERVES": 1172.35,
			"GOLD_RESERVES_SAME": 3.63314917,
			"GOLD_RESERVES_SEQUENTIAL": 5.00223914,
			"FOREX": 31276.91,
			"FOREX_SAME": -3.76826907,
			"FOREX_SEQUENTIAL": 0.32728274
		}
`

	LendingRateResponse = `
		{
			"REPORT_DATE": "2023-01-16 00:00:00",
			"IR_RATE": 1.568,
			"CHANGE_RATE": 32.6
		}
`

	IndustryResponse = `
		{
            "REPORT_DATE": "2004-12-06 00:00:00",
            "INDICATOR_VALUE": 6208,
            "CHANGE_RATE": 0.58327933,
            "CHANGERATE_3M": 54.65869457,
            "CHANGERATE_6M": 94.91365777,
            "CHANGERATE_1Y": 37.49723145,
            "CHANGERATE_2Y": 285.11166253,
            "CHANGERATE_3Y": 612.74397245
        }
`
)

func Poster(ic *model.IndexConfig) {
	var body string
	switch ic.TargetCode {
	case "HG00016", "HG00017":
		body = fmt.Sprintf(template, ic.DataSourceName, ic.TargetName, ic.TargetCode, api.IAV, IAVResponse)
	case "HG00027", "HG00028", "HG00029", "HG00030":
		body = fmt.Sprintf(template, ic.DataSourceName, ic.TargetName, ic.TargetCode, api.XFP, XFPResponse)
	case "HG00006", "HG00007":
		body = fmt.Sprintf(template, ic.DataSourceName, ic.TargetName, ic.TargetCode, api.M2, M2Response)
	case "HG00003", "HG00004", "HG00088":
		body = fmt.Sprintf(template, ic.DataSourceName, ic.TargetName, ic.TargetCode, api.CPI, CPIResponse)
	case "HG00020":
		body = fmt.Sprintf(template, ic.DataSourceName, ic.TargetName, ic.TargetCode, api.PMI, PMIResponse)
	case "HG00023", "HG00089":
		body = fmt.Sprintf(template, ic.DataSourceName, ic.TargetName, ic.TargetCode, api.PPI, PPIResponse)
	case "HG00066":
		body = fmt.Sprintf(template, ic.DataSourceName, ic.TargetName, ic.TargetCode, api.ZBJ, ZBJResponse)
	case "HG00065":
		body = fmt.Sprintf(template, ic.DataSourceName, ic.TargetName, ic.TargetCode, api.JCK, JCKResponse)
	case "HY00007", "HY00011":
		body = fmt.Sprintf(template, ic.DataSourceName, ic.TargetName, ic.TargetCode, api.LL, LLResponse)
	case "HG00090", "HG00091":
		body = fmt.Sprintf(template, ic.DataSourceName, ic.TargetName, ic.TargetCode, api.WH, WHResponse)
	case "HY00006", "HY00012", "HY00013", "HY00014", "HY00015", "HY00016", "HY00017", "HY00018":
		url := strings.ReplaceAll(api.LendingRate, "#", ic.SourceTargetCodeSpider)
		body = fmt.Sprintf(template, ic.DataSourceName, ic.TargetName, ic.TargetCode, url, LendingRateResponse)
	case "HY00001", "HY00002", "HY00003", "HG00110", "HG00111":
		body = fmt.Sprintf(template, ic.DataSourceName, ic.TargetName, ic.TargetCode, api.Industry, IndustryResponse)
	}
	if err := e.Send(receivers, subject, body); err != nil {
		logger.Error(err.Error())
	}
}
