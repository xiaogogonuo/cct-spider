package poster

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/api"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/model"
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
	"github.com/xiaogogonuo/cct-spider/pkg/mail"
	"strings"
	"time"
)

const (
	user = "xiaogogonuo@163.com"
	pass = "JDAOREDDCXYMAXXQ"
	nick = "lujiawei"
)

var (
	Receivers = []string{"xiaogogonuo@163.com"}
	Subject   = "城通爬虫系统警报"
)

var E163 = mail.NewEmail163(user, pass, nick)

const Template = `
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

const TemplateGlobal = `
<html>
    <body>
		<div>数据源：%s</div>
		<div>指标名称：%s</div>
		<div>指标代码：%s</div>
		<div>原接口：%s 返回的数据格式应该如下所示：</div>
        <p>
			%s
        </p>
		<div>请检查接口的返回数据格式是否同上案例一致</div>
	</body>
</html>
`

const TemplateExchange = `
<html>
    <body>
		<div>数据源：%s</div>
		<div>指标名称：%s</div>
		<div>指标代码：%s</div>
		<div>原接口：%s 返回的数据格式应该如下所示：</div>
			<table class="market_tab_big foreign" id="GJZS">
			
				<tr>
					<th width="330px">名称</th>
					<th width="117px">最新价</th>
					<th width="117px">涨跌</th>
					<th width="117px">涨跌幅</th>
					<th width="117px">最高</th>
					<th width="117px">最低</th>
					<th width="117px">昨收</th>
					<th width="117px">更新时间</th>
				</tr>
			
				<tr id="NIKKI">
					<td><a class="mar_name" href="https://quote.fx678.com/symbol/NIKKI">日经225</a></td>
					<td>26791.12</td>
					<td><span class="arrow_red">652.44</span></td>
					<td class="red_tab">2.50</td>
					<td>26816.68</td>
					<td>26211.84</td>
					<td>26138.68</td>
					<td>14:00:02</td>
				</tr>
			
			</table>
		<div>请检查接口的返回数据格式是否同上案例一致</div>
	</body>
</html>
`

const TemplateXiBen = `
<html>
    <body>
		<div>数据源：%s</div>
		<div>指标名称：%s</div>
		<div>指标代码：%s</div>
		<div>原接口：%s 返回的数据格式应该如下所示：</div>
			<tbody id="indexdetaildata">
				<tr>
					<td>2022-12-31</td>
					<td>14000.00</td>
					<td class=" rise">1900.00</td>
					<td class=" rise">15.70</td>
				</tr>
				<tr>
					<td>2022-11-30</td>
					<td>12100.00</td>
					<td class=" rise">5948.00</td>
					<td class=" rise">96.68</td>
				</tr>
			</tbody>
		<div>请检查接口的返回数据格式是否同上案例一致</div>
	</body>
</html>
`

const TemplateWuLiu = `
<html>
    <body>
		<div>数据源：中国物流</div>
		<div>指标名称：中国非制造业商务活动指数</div>
		<div>指标代码：HG00094</div>
		<div>原接口：http://www.chinawuliu.com.cn/xsyj/tjsj/ 返回的数据格式应该如下所示：</div>
			<div class="col-sm-8 leftRow">
				<div class="media-body media-body-inner ">
					<ul class="list-box list-box--pre">
						<li>
							<a href="xxx" title="xxx">2022年12月份中国物流业景气指数为46%</a>
						</li>
					</ul>
				</div>
			</div>
		<div>请检查接口的返回数据格式是否同上案例一致</div>
	</body>
</html>
`

const TemplateNrc = `
<html>
    <body>
		<div>数据源：%s</div>
		<div>指标名称：%s</div>
		<div>指标代码：%s</div>
		<div>原接口：https://www.ndrc.gov.cn/fggz/fgzh/gnjjjc/hbjr/202212/t20221229_1344648.html 返回的数据格式应该如下所示：</div>
			<div class="article_con article_con_title">
				<div class=TRS_Editor>
					<div class="Custom_UnionStyle">
						<p>xxx</p>
					</div>
				</div>
			</div>
		<div>请检查接口的返回数据格式是否同上案例一致</div>
	</body>
</html>
`

const (
	GDPResponse = `
		{
			"REPORT_DATE": "2022-12-01 00:00:00",
			"TIME": "2022年第1-4季度",
			"DOMESTICL_PRODUCT_BASE": 1210207,
			"FIRST_PRODUCT_BASE": 88345,
			"SECOND_PRODUCT_BASE": 483164,
			"THIRD_PRODUCT_BASE: 638698,
			"SUM_SAME": 3,
			"FIRST_SAME": 4.1,
			"SECOND_SAME": 3.8,
			"THIRD_SAME": 2.3
		}
`

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

	GlobalResponse = `
	{
		"rc": 0,
		"rt": 4,
		"svr": 182994189,
		"lt": 1,
		"full": 1,
		"dlmkts": "",
		"data": {
			"f43": 19874,
			"f44": 19878,
			"f45": 19874,
			"f46": 19878,
			"f60": 19877,
			"f86": 1673942372,
			"f169": -3,
			"f170": -2
		}
	}
`

	NationalDebtResponse = `
		{
			"SOLAR_DATE": "2023-01-17 00:00:00",
			"EMM00588704": 2.4013,
			"EMM00166462": 2.7346,
			"EMM00166466": 2.9017,
			"EMM00166469": 3.2531,
			"EMM01276014": 0.5004,
			"EMM00000024": null,
			"EMG00001306": 4.18,
			"EMG00001308": 3.6,
			"EMG00001310": 3.53,
			"EMG00001312": 3.64,
			"EMG01339436": -0.65,
			"EMG00159635": null
}
`

	ForexResponse = `var hq_str_fx_s^="13:33:05,0.1475,0.1475,0.1478,3,0.1475,0.1476,0.1473,0.1475,人民币兑美元即期汇率,-0.1,-0.0002,0.002034,x,0.1586,0.1365,x,2023-01-18";`

	SinaMarcoRegionGDPResponse = `
		{
			config: {
				all: [],
				defaultItems: [2],
				except: "231",
				index: 2,
				querylist: []
			},
			count: "6326",
			data: [
				["2019.10", "新疆", "92.20", null, null, null, null, null, null, null, null, null],
				["2019.10", "宁夏", "96.70", null, null, null, null, null, null, null, null, null]
			]
		}
`

	SinaMarcoRegionCPIResponse = `
		{
			config: {
				all: [],
				index: 2,
				defaultItems: [],
				C: "F005",
				Cmap: [],
				except: "100000",
				querylist: [],
			count: "11022",
			data: {
				'累计': [
					["2022.12", "新疆", "101.80"],
					["2022.12", "宁夏", "102.30"]
				]
			}
		}
`

	SCIResponse = `
		{
			"DataItem": {},
			"List": [
				{
					"DataName": "",
					"DIID": 58222,
					"DataTypeID": 201,
					"MDataValue": 1255.65,
					"Change": 0.15,
					"ChangeRate": 0.01,
					"DataDate": "2023/01/17",
					"ChangeValue": null,
					"ChangePercent": null,
					"Level": 0,
					"Fpname": null,
					"IsHasExplain": false
				}
			],
			"News": {}
		}
`

	ExchangeSpecialResponse = `
		{
			"s": "ok",
			"t": ["1674033450", "1674033450"],
			"c": ["457.0300", "457.0300"],
			"o": ["456.5100", "456.5100"],
			"h": ["457.9400", "457.9400"],
			"l": ["456.5100", "456.5100"],
			"p": ["456.4600", "456.4600"]
		}
`

	IrcResponse = `
		{
			"rptCode": 200,
			"msg": "Success",
			"data": {
				"total": 42,
				"lists": [
					{
						"id": "3627778",
						"docId": "1090231",
						"itemId": "995",
						"itemName": "监管动态",
						"keyword": [
							"监管动态",
							"一、上海银行业资产负债情况\n2022年11月，上海辖内银行业金融机构本外币总资产余额23.29万亿元，同比增长11.93%。
							商业银行本外币总资产余额20.36万亿元，同比增长11.62%，其中大型商业银行本外币资产余额8.28万亿元，同比增长11.90%；
							股份制商业银行本外币资产余额5.76万亿元，同比增长14.34%。\n2022年11月，上海辖内银行业金融机构本外币总负债余额
							22.28万亿元，同比增长12.08%。商业银行本外币总负债余额19.72万亿元，同比增长11.82%，其中大型商业银行本外币负
							债余额8.18万亿元，同比增长11.99%；股份制商业银行本外币负债余额5.73万亿元，同比增长14.54%。\n二、上海银行业
							资产质量情况\n2022年11月，上海辖内银行业金融机构本外币不良贷款余额795.38亿元，不良贷款率0.77%；其中，商业银
							行本外币不良贷款余额565.11亿元，不良贷款率0.67%。\n三、上海保险业主要经营数据\n2022年1—11月，上海辖内保险公
							司原保险保费收入累计1918亿元，其中财产险公司原保险保费收入605亿元，人身险公司原保险保费收入1313亿元。\n2022
							年1—11月，上海辖内保险公司原保险赔付支出累计598亿元，其中财产险公司原保险赔付支出301亿元，人身险公司原保险赔付
							支出297亿元。\n2022年1—11月，上海辖内保险公司保单件数当年累计319170万件，其中财产险公司316665万件，人身险
							公司2505万件。\n\n注：因部分保险机构目前处于风险处置阶段，从2021年6月起，汇总数据口径暂不包括这部分机构。,
							"上海银保监局发布2022年11月辖内银行业保险业主要监管指标数据情况"
						]
				],
				"checkpage": 4,
				"agencyOrgs": [],
				"items": []
			}
		}
`
)

func Poster(ic *model.IndexConfig) {
	var body string
	switch ic.TargetCode {
	case "HG00001", "HG00098":
		body = fmt.Sprintf(Template, ic.DataSourceName, ic.TargetName, ic.TargetCode, api.GDP, GDPResponse)
	case "HG00016", "HG00017":
		body = fmt.Sprintf(Template, ic.DataSourceName, ic.TargetName, ic.TargetCode, api.IAV, IAVResponse)
	case "HG00027", "HG00028", "HG00029", "HG00030":
		body = fmt.Sprintf(Template, ic.DataSourceName, ic.TargetName, ic.TargetCode, api.XFP, XFPResponse)
	case "HG00006", "HG00007":
		body = fmt.Sprintf(Template, ic.DataSourceName, ic.TargetName, ic.TargetCode, api.M2, M2Response)
	case "HG00003", "HG00004", "HG00088":
		body = fmt.Sprintf(Template, ic.DataSourceName, ic.TargetName, ic.TargetCode, api.CPI, CPIResponse)
	case "HG00020":
		body = fmt.Sprintf(Template, ic.DataSourceName, ic.TargetName, ic.TargetCode, api.PMI, PMIResponse)
	case "HG00023", "HG00089":
		body = fmt.Sprintf(Template, ic.DataSourceName, ic.TargetName, ic.TargetCode, api.PPI, PPIResponse)
	case "HG00066":
		body = fmt.Sprintf(Template, ic.DataSourceName, ic.TargetName, ic.TargetCode, api.ZBJ, ZBJResponse)
	case "HG00065":
		body = fmt.Sprintf(Template, ic.DataSourceName, ic.TargetName, ic.TargetCode, api.JCK, JCKResponse)
	case "HY00007", "HY00011":
		body = fmt.Sprintf(Template, ic.DataSourceName, ic.TargetName, ic.TargetCode, api.LL, LLResponse)
	case "HG00090", "HG00091":
		body = fmt.Sprintf(Template, ic.DataSourceName, ic.TargetName, ic.TargetCode, api.WH, WHResponse)
	case "HY00006", "HY00012", "HY00013", "HY00014", "HY00015", "HY00016", "HY00017", "HY00018":
		url := strings.ReplaceAll(api.LendingRate, "#", ic.SourceTargetCodeSpider)
		body = fmt.Sprintf(Template, ic.DataSourceName, ic.TargetName, ic.TargetCode, url, LendingRateResponse)
	case "HY00001", "HY00002", "HY00003", "HG00110", "HG00111":
		body = fmt.Sprintf(Template, ic.DataSourceName, ic.TargetName, ic.TargetCode, api.Industry, IndustryResponse)
	case "HY00005", "HG00099", "HG00100", "HG00101", "HG00102", "HG00103", "HG00104", "HG00105", "HG00106", "HG00107",
		"HG00108", "HG00109", "HG00112", "HG00113", "HG00114", "HG00115", "HG00116", "HG00117":
		url := strings.ReplaceAll(api.Global, "#", ic.SourceTargetCodeSpider)
		body = fmt.Sprintf(TemplateGlobal, ic.DataSourceName, ic.TargetName, ic.TargetCode, url, GlobalResponse)
	case "HG00062":
		body = fmt.Sprintf(Template, ic.DataSourceName, ic.TargetName, ic.TargetCode, api.NationalDebt, NationalDebtResponse)
	case "HY00008", "HG00086", "HG00087":
		url := strings.ReplaceAll(api.Forex, "#", strings.ToLower(ic.SourceTargetCodeSpider))
		url = strings.ReplaceAll(url, "$", fmt.Sprintf("%d", time.Now().UnixNano()/1e6))
		body = fmt.Sprintf(TemplateGlobal, ic.DataSourceName, ic.TargetName, ic.TargetCode, url, ForexResponse)
		body = strings.ReplaceAll(body, "^", strings.ToLower(ic.SourceTargetCodeSpider))
	case "HG00002":
		url := strings.ReplaceAll(api.Marco, "#", ic.SourceTargetCodeSpider)
		url = strings.ReplaceAll(url, "$", "0")
		body = fmt.Sprintf(TemplateGlobal, ic.DataSourceName, ic.TargetName, ic.TargetCode, url, SinaMarcoRegionGDPResponse)
	case "HG00040":
		url := strings.ReplaceAll(api.Marco, "#", ic.SourceTargetCodeSpider)
		url = strings.ReplaceAll(url, "$", "0")
		body = fmt.Sprintf(TemplateGlobal, ic.DataSourceName, ic.TargetName, ic.TargetCode, url, SinaMarcoRegionCPIResponse)
	case "HY00004", "HY00010":
		body = fmt.Sprintf(TemplateGlobal, ic.DataSourceName, ic.TargetName, ic.TargetCode, api.SCI, SCIResponse)
	case "HG00042":
		u := strings.ReplaceAll(api.FxExchangeSpecial, "#", ic.SourceTargetCodeSpider)
		u = strings.ReplaceAll(u, "$", ic.SourceTargetCode)
		body = fmt.Sprintf(TemplateGlobal, ic.DataSourceName, ic.TargetName, ic.TargetCode, u, ExchangeSpecialResponse)
	case "HG00092", "HG00093":
		url := strings.ReplaceAll(api.XiBen, "#", ic.SourceTargetCodeSpider)
		body = fmt.Sprintf(TemplateXiBen, ic.DataSourceName, ic.TargetName, ic.TargetCode, url)
	case "HG00094":
		body = TemplateWuLiu
	case "HG00095":
		body = fmt.Sprintf(TemplateGlobal, ic.DataSourceName, ic.TargetName, ic.TargetCode, api.NPLoan, IrcResponse)
	case "HG00118", "HG00119", "HG00120":
		body = fmt.Sprintf(TemplateNrc, ic.DataSourceName, ic.TargetName, ic.TargetCode)
	default:
		url := strings.ReplaceAll(api.FxExchange, "#", ic.SourceTargetCodeSpider)
		body = fmt.Sprintf(TemplateExchange, ic.DataSourceName, ic.TargetName, ic.TargetCode, url)
	}
	if err := E163.Send(Receivers, Subject, body); err != nil {
		logger.Error(err.Error())
	}
}
