package poster

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/api"
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
	"github.com/xiaogogonuo/cct-spider/pkg/mail"
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

func IAVPoster() {
	body := `
<html>
    <body>
		<div>数据源名称：东方财富</div>
		<div>指标名称：工业增加值同比增长、工业增加值同比增长</div>
		<div>指标代码：HG00016、HG00017</div>
		<div>原接口：%s</div>
		<div>返回的数据格式如下：</div>
        <p>
			{
				"version": "434f9c6126190079ad8950890e8546c7",
				"result": {
					"pages": 164,
					"data": [
						{
							"REPORT_DATE": "2022-11-01 00:00:00",
							"TIME": "2022年11月份",
							"BASE_SAME": 2.2,
							"BASE_ACCUMULATE": 3.8
						}
					],
					"count": 164
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
	body = fmt.Sprintf(body, api.IAV)
	if err := e.Send(receivers, subject, body); err != nil {
		logger.Error(err.Error())
	}
}

func XFPPoster() {
	body := `
<html>
    <body>
		<div>数据源名称：东方财富</div>
		<div>指标名称：社会消费品零售总额当期值、社会消费品零售总额累计值、社会消费品零售总额同比增长、社会消费品零售总额累计增长</div>
		<div>指标代码：HG00027、HG00028、HG00029、HG00030</div>
		<div>原接口：%s</div>
		<div>返回的数据格式如下：</div>
        <p>
			{
				"version": "434f9c6126190079ad8950890e8546c7",
				"result": {
					"pages": 164,
					"data": [
						{
							"REPORT_DATE": "2022-11-01 00:00:00",
							"TIME": "2022年11月份",
							"RetailTotal": 38615,
							"RetailTotalSame": -5.9,
							"RetailTotalSequential": -4.11214025,
							"RetailTotalAccumulate": 399190,
							"RetailAccumulateSame": -0.1
						}
					],
					"count": 164
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
	body = fmt.Sprintf(body, api.XFP)
	if err := e.Send(receivers, subject, body); err != nil {
		logger.Error(err.Error())
	}
}

func M2Poster() {
	body := `
<html>
    <body>
		<div>数据源名称：东方财富</div>
		<div>指标名称：货币和准货币(M2)供应量期末值、货币和准货币(M2)供应量同比增长</div>
		<div>指标代码：HG00006、HG00007</div>
		<div>原接口：%s</div>
		<div>返回的数据格式如下：</div>
        <p>
			{
				"version": "434f9c6126190079ad8950890e8546c7",
				"result": {
					"pages": 164,
					"data": [
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
					],
					"count": 164
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
	body = fmt.Sprintf(body, api.M2)
	if err := e.Send(receivers, subject, body); err != nil {
		logger.Error(err.Error())
	}
}

func CPIPoster() {
	body := `
<html>
    <body>
		<div>数据源名称：东方财富</div>
		<div>指标名称：CPI同比增速月度</div>
		<div>指标代码：HG00088</div>
		<div>原接口：%s</div>
		<div>返回的数据格式如下：</div>
        <p>
			{
				"version": "434f9c6126190079ad8950890e8546c7",
				"result": {
					"pages": 164,
					"data": [
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
					],
					"count": 164
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
	body = fmt.Sprintf(body, api.CPI)
	if err := e.Send(receivers, subject, body); err != nil {
		logger.Error(err.Error())
	}
}

func PMIPoster() {
	body := `
<html>
    <body>
		<div>数据源名称：东方财富</div>
		<div>指标名称：制造业采购经理指数</div>
		<div>指标代码：HG00020</div>
		<div>原接口：%s</div>
		<div>返回的数据格式如下：</div>
        <p>
			{
				"version": "434f9c6126190079ad8950890e8546c7",
				"result": {
					"pages": 1,
					"data": [
						{
							"REPORT_DATE"": "2022-12-01 00:00:00",
							"TIME"": "2022年12月份",
							"MAKE_INDEX":47,
							"MAKE_SAME":-6.56063618,
							"NMAKE_INDEX":41.6,
							"NMAKE_SAME":-21.0626186
						}
					],
					"count": 180
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
	body = fmt.Sprintf(body, api.PMI)
	if err := e.Send(receivers, subject, body); err != nil {
		logger.Error(err.Error())
	}
}

func PPIPoster() {
	body := `
<html>
    <body>
		<div>数据源名称：东方财富</div>
		<div>指标名称：工业品出厂价格指数当月、PPI同比增速月度</div>
		<div>指标代码：HG00023、HG00089</div>
		<div>原接口：%s</div>
		<div>返回的数据格式如下：</div>
        <p>
			{
				"version": "434f9c6126190079ad8950890e8546c7",
				"result": {
					"pages": 1,
					"data": [
						{
							"REPORT_DATE"": "2022-12-01 00:00:00",
							"TIME"": "2022年12月份",
							"BASE":99.3,
							"BASE_SAME":-0.7,
							"BASE_ACCUMULATE":104.1
						}
					],
					"count": 204
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
	body = fmt.Sprintf(body, api.PPI)
	if err := e.Send(receivers, subject, body); err != nil {
		logger.Error(err.Error())
	}
}

func ZBJPoster() {
	body := `
<html>
    <body>
		<div>数据源名称：东方财富</div>
		<div>指标名称：存款准备金率</div>
		<div>指标代码：HG00066</div>
		<div>原接口：%s</div>
		<div>返回的数据格式如下：</div>
        <p>
			{
				"version": "434f9c6126190079ad8950890e8546c7",
				"result": {
					"pages": 1,
					"data": [
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
					],
					"count": 53
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
	body = fmt.Sprintf(body, api.ZBJ)
	if err := e.Send(receivers, subject, body); err != nil {
		logger.Error(err.Error())
	}
}

func JCKPoster() {
	body := `
<html>
    <body>
		<div>数据源名称：东方财富</div>
		<div>指标名称：出口当月同比增速</div>
		<div>指标代码：HG00065</div>
		<div>原接口：%s</div>
		<div>返回的数据格式如下：</div>
        <p>
			{
				"version": "434f9c6126190079ad8950890e8546c7",
				"result": {
					"pages": 1,
					"data": [
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
					],
					"count": 180
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
	body = fmt.Sprintf(body, api.JCK)
	if err := e.Send(receivers, subject, body); err != nil {
		logger.Error(err.Error())
	}
}

func LLPoster() {
	body := `
<html>
    <body>
		<div>数据源名称：东方财富</div>
		<div>指标名称：贷款基准利率、存款基准利率</div>
		<div>指标代码：HY00007、HY00011</div>
		<div>原接口：%s</div>
		<div>返回的数据格式如下：</div>
        <p>
			{
				"version": "434f9c6126190079ad8950890e8546c7",
				"result": {
					"pages": 1,
					"data": [
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
					],
					"count": 27
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
	body = fmt.Sprintf(body, api.LL)
	if err := e.Send(receivers, subject, body); err != nil {
		logger.Error(err.Error())
	}
}

func WHPoster() {
	body := `
<html>
    <body>
		<div>数据源名称：东方财富</div>
		<div>指标名称：外汇储备、外汇储备同比增速</div>
		<div>指标代码：HG00090、HG00091</div>
		<div>原接口：%s</div>
		<div>返回的数据格式如下：</div>
        <p>
			{
				"version": "434f9c6126190079ad8950890e8546c7",
				"result": {
					"pages": 1,
					"data": [
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
					],
					"count": 180
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
	body = fmt.Sprintf(body, api.WH)
	if err := e.Send(receivers, subject, body); err != nil {
		logger.Error(err.Error())
	}
}
