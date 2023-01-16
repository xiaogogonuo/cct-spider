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
				version: "d07e8629dcb9b3c9ecc3eecb895f4600",
				result: {
					pages: 1,
					data: [
						{
							REPORT_DATE: "2022-09-01 00:00:00",
							TIME: "2022年第1-3季度",
							DOMESTICL_PRODUCT_BASE: 870269,
							FIRST_PRODUCT_BASE: 54779.1,
							SECOND_PRODUCT_BASE: 350189.5,
							THIRD_PRODUCT_BASE: 465300.4,
							SUM_SAME: 3,
							FIRST_SAME: 4.2,
							SECOND_SAME: 3.9,
							THIRD_SAME: 2.3
						}
					],
					count: 1
				},
				success: true,
				message: "ok",
				code: 0
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
				version: "434f9c6126190079ad8950890e8546c7",
				result: {
					pages: 164,
					data: [
						{
							REPORT_DATE: "2022-11-01 00:00:00",
							TIME: "2022年11月份",
							BASE_SAME: 2.2,
							BASE_ACCUMULATE: 3.8
						}
					],
					count: 164
				},
				success: true,
				message: "ok",
				code: 0
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
				version: "434f9c6126190079ad8950890e8546c7",
				result: {
					pages: 164,
					data: [
						{
							REPORT_DATE: "2022-11-01 00:00:00",
							TIME: "2022年11月份",
							RetailTotal: 38615,
							RetailTotalSame: -5.9,
							RetailTotalSequential: -4.11214025,
							RetailTotalAccumulate: 399190,
							RetailAccumulateSame: -0.1
						}
					],
					count: 164
				},
				success: true,
				message: "ok",
				code: 0
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
				version: "434f9c6126190079ad8950890e8546c7",
				result: {
					pages: 164,
					data: [
						{
							REPORT_DATE: "2022-12-01 00:00:00",
							TIME: "2022年12月份",
							BASIC_CURRENCY: 2664300,
							BASIC_CURRENCY_SAME: 11.8,
							BASIC_CURRENCY_SEQUENTIAL: 0.65324762,
							CURRENCY: 671700,
							CURRENCY_SAME: 3.7,
							CURRENCY_SEQUENTIAL: 0.69821477,
							FREE_CASH: 104700,
							FREE_CASH_SAME: 15.3,
							FREE_CASH_SEQUENTIAL: 4.97280332
						}
					],
					count: 164
				},
				success: true,
				message: "ok",
				code: 0
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
