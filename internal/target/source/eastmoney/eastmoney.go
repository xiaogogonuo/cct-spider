package eastmoney

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/xiaogogonuo/cct-spider/internal/target/model"
	"github.com/xiaogogonuo/cct-spider/internal/target/pkg/downloader"
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

// APIEastMoneyEconomicTarget 东方财富宏观指标数据接口
var APIEastMoneyEconomicTarget = "http://datainterface.eastmoney.com/EM_DataCenter/JS.aspx?type=GJZB&sty=ZGZB&p=1&ps=10000&mkt="

// SpiderEastMoneyEconomicTarget 爬取"东方财富"网站的`宏观指标`
// 适用指标：
// - 工业增加值同比增长、工业增加值累计增长
//   • 页面展示接口：https://data.eastmoney.com/cjsj/gyzjz.html
//   • 数据获取接口：http://datainterface.eastmoney.com/EM_DataCenter/JS.aspx?type=GJZB&sty=ZGZB&p=1&ps=10000&mkt=0
// - 出口当月同比增速
//   • 页面展示接口：https://data.eastmoney.com/cjsj/hgjck.html
//   • 数据获取接口：http://datainterface.eastmoney.com/EM_DataCenter/JS.aspx?type=GJZB&sty=ZGZB&p=1&ps=10000&mkt=1
// - 社会消费品零售总额当期值、社会消费品零售总额累计值、社会消费品零售总额同比增长、社会消费品零售总额累计增长
//   • 页面展示接口：https://data.eastmoney.com/cjsj/xfp.html
//   • 数据获取接口：http://datainterface.eastmoney.com/EM_DataCenter/JS.aspx?type=GJZB&sty=ZGZB&p=1&ps=10000&mkt=5
// - 货币和准货币(M2)供应量期末值、货币和准货币(M2)供应量同比增长
//   • 页面展示接口：https://data.eastmoney.com/cjsj/hbgyl.html
//   • 数据获取接口：http://datainterface.eastmoney.com/EM_DataCenter/JS.aspx?type=GJZB&sty=ZGZB&p=1&ps=10000&mkt=11
// - 贷款基准利率、存款基准利率
//   • 页面展示接口：https://data.eastmoney.com/cjsj/yhll.html
//   • 数据获取接口：http://datainterface.eastmoney.com/EM_DataCenter/JS.aspx?type=GJZB&sty=ZGZB&p=1&ps=10000&mkt=13
// - 居民消费价格指数当月
//   • 页面展示接口：https://data.eastmoney.com/cjsj/cpi.html
//   • 数据获取接口：http://datainterface.eastmoney.com/EM_DataCenter/JS.aspx?type=GJZB&sty=ZGZB&p=1&ps=10000&mkt=19
// - 国内生产总值同比增长
//   • 页面展示接口：https://data.eastmoney.com/cjsj/gdp.html
//   • 数据获取接口：http://datainterface.eastmoney.com/EM_DataCenter/JS.aspx?type=GJZB&sty=ZGZB&p=1&ps=10000&mkt=20
// - 制造业采购经理指数
//   • 页面展示接口：https://data.eastmoney.com/cjsj/pmi.html
//   • 数据获取接口：http://datainterface.eastmoney.com/EM_DataCenter/JS.aspx?type=GJZB&sty=ZGZB&p=1&ps=10000&mkt=21
// - 工业品出厂价格指数当月
//   • 页面展示接口：https://data.eastmoney.com/cjsj/ppi.html
//   • 数据获取接口：http://datainterface.eastmoney.com/EM_DataCenter/JS.aspx?type=GJZB&sty=ZGZB&p=1&ps=10000&mkt=22
// - 存款准备金率
//   • 页面展示接口：https://data.eastmoney.com/cjsj/ckzbj.html
//   • 数据获取接口：http://datainterface.eastmoney.com/EM_DataCenter/JS.aspx?type=GJZB&sty=ZGZB&p=1&ps=10000&mkt=23
func SpiderEastMoneyEconomicTarget(sourceTargetCodeSpider, targetCode string) (responses []model.Response) {
	url := APIEastMoneyEconomicTarget + sourceTargetCodeSpider
	body, err := downloader.SimpleGet(url)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	stringBody := *(*string)(unsafe.Pointer(&body))
	stringBody = stringBody[1 : len(stringBody)-1] // 去除开头的"("和结尾的")"
	var table []string                             // 表格数据
	if err := json.Unmarshal([]byte(stringBody), &table); err != nil {
		logger.Error(err.Error())
		return
	}
	switch targetCode {
	case "HY00007": // 贷款基准利率
		responses = eastMoneyEconomicTargetMostPipeline(table, 9, 5, 8)
	case "HY00011": // 存款基准利率
		responses = eastMoneyEconomicTargetMostPipeline(table, 9, 2, 8)
	case "HG00001": // 国内生产总值同比增长
		responses = eastMoneyEconomicTargetGDPPipeline(table, 0, 2)
	case "HG00004": // 居民消费价格指数当月
		responses = eastMoneyEconomicTargetMostPipeline(table, 0, 1, 6)
	case "HG00006": // 货币和准货币(M2)供应量期末值
		responses = eastMoneyEconomicTargetMostPipeline(table, 0, 1, 6)
	case "HG00007": // 货币和准货币(M2)供应量同比增长
		responses = eastMoneyEconomicTargetMostPipeline(table, 0, 2, 6)
	case "HG00016": // 工业增加值同比增长
		responses = eastMoneyEconomicTargetMostPipeline(table, 0, 1, 6)
	case "HG00017": // 工业增加值累计增长
		responses = eastMoneyEconomicTargetMostPipeline(table, 0, 2, 6)
	case "HG00020": // 制造业采购经理指数
		responses = eastMoneyEconomicTargetMostPipeline(table, 0, 1, 6)
	case "HG00023": // 工业品出厂价格指数当月
		responses = eastMoneyEconomicTargetMostPipeline(table, 0, 1, 6)
	case "HG00027": // 社会消费品零售总额当期值
		responses = eastMoneyEconomicTargetMostPipeline(table, 0, 1, 6)
	case "HG00028": // 社会消费品零售总额累计值
		responses = eastMoneyEconomicTargetMostPipeline(table, 0, 4, 6)
	case "HG00029": // 社会消费品零售总额同比增长
		responses = eastMoneyEconomicTargetMostPipeline(table, 0, 2, 6)
	case "HG00030": // 社会消费品零售总额累计增长
		responses = eastMoneyEconomicTargetMostPipeline(table, 0, 5, 6)
	case "HG00065": // 出口当月同比增速
		responses = eastMoneyEconomicTargetMostPipeline(table, 0, 2, 6)
	case "HG00066": // 存款准备金率
		responses = eastMoneyEconomicTargetMostPipeline(table, 1, 3, 6)
	}
	return
}

// APIEastMoneyEconomicTargetBOR 东方财富宏观指标`上海银行间同业拆放利率隔夜`数据接口
var APIEastMoneyEconomicTargetBOR = "https://data.eastmoney.com/shibor/shibor.aspx?%s&p=%d"

// SpiderEastMoneyEconomicTargetBOR 爬取"东方财富"网站的`宏观指标-银行间同业拆放利率`
// 适用指标：
// - 上海银行间同业拆放利率隔夜
//   • 页面展示接口：https://data.eastmoney.com/shibor/shibor.aspx?m=sh&t=99&d=99221&cu=cny&type=009016&p=1
//   • 数据获取接口：https://data.eastmoney.com/shibor/shibor.aspx?m=sh&t=99&d=99221&cu=cny&type=009016&p=1
// - 上海银行间同业拆放利率3月
//   • 页面展示接口：https://data.eastmoney.com/shibor/shibor.aspx?m=sh&t=99&d=99225&cu=cny&type=009020&p=1
//   • 数据获取接口：https://data.eastmoney.com/shibor/shibor.aspx?m=sh&t=99&d=99225&cu=cny&type=009020&p=1
// - 伦敦同业间拆借利率3月
//   • 页面展示接口：https://data.eastmoney.com/shibor/shibor.aspx?m=ld&t=96&d=99255&cu=usd&type=009006&p=1
//   • 数据获取接口：https://data.eastmoney.com/shibor/shibor.aspx?m=ld&t=96&d=99255&cu=usd&type=009006&p=1
func SpiderEastMoneyEconomicTargetBOR(sourceTargetCodeSpider string, pages int) (responses []model.Response) {
	var ra model.ResponseArray
	for page := 1; ; page++ {
		if page > pages && pages != -1 {
			break
		}
		url := fmt.Sprintf(APIEastMoneyEconomicTargetBOR, sourceTargetCodeSpider, page)
		body, err := downloader.SimpleGet(url)
		if err != nil {
			logger.Error(err.Error())
			continue
		}
		stringBody := *(*string)(unsafe.Pointer(&body))
		dom, _ := goquery.NewDocumentFromReader(strings.NewReader(stringBody))
		var tableText []string
		dom.Find("table[id='tb'] td").Each(func(i int, selection *goquery.Selection) {
			text := selection.Text()
			tableText = append(tableText, text)
		})
		if len(tableText) < 1 {
			break
		}
		var response model.ResponseDateStringValueFloat
		// 提取表格中的日期和指标值
		for idx, value := range tableText {
			if idx%3 == 0 {
				response.Date = strings.ReplaceAll(value, "-", "")
			} else if idx%3 == 1 {
				response.TargetValue, _ = strconv.ParseFloat(value, 64)
			} else {
				ra = append(ra, response)
			}
		}
	}
	responses = ManualHandleUpDown(ra)
	return
}

// APIEastMoneyEconomicTargetCHN10 东方财富宏观指标`中债10年期国债到期收益率`数据接口
var APIEastMoneyEconomicTargetCHN10 = "https://datacenter-web.eastmoney.com/api/data/get?type=RPTA_WEB_TREASURYYIELD&sty=ALL&st=SOLAR_DATE&sr=-1&p=%d&ps=500"

// SpiderEastMoneyEconomicTargetCHN10 爬取"东方财富"网站的`宏观指标-中债10年期国债到期收益率`
// 适用指标：
// - 中债10年期国债到期收益率
//   • 页面展示接口：https://data.eastmoney.com/cjsj/zmgzsyl.html
//   • 数据获取接口：https://datacenter-web.eastmoney.com/api/data/get?type=RPTA_WEB_TREASURYYIELD&sty=ALL&st=SOLAR_DATE&sr=-1&p=1&ps=500
func SpiderEastMoneyEconomicTargetCHN10(pages int) (responses []model.Response) {
	var ra model.ResponseArray
	for page := 1; ; page++ {
		if page > pages && pages != -1 {
			break
		}
		url := fmt.Sprintf(APIEastMoneyEconomicTargetCHN10, page)
		body, err := downloader.SimpleGet(url)
		if err != nil {
			logger.Error(err.Error())
			continue
		}
		stringBody := *(*string)(unsafe.Pointer(&body))
		var chn eastMoneyCHN10
		if err := json.Unmarshal([]byte(stringBody), &chn); err != nil {
			logger.Error(err.Error())
			continue
		}
		if !chn.Success {
			break
		}
		for _, data := range chn.Result.Data {
			date := strings.Split(data.SolarDate, "/")
			var (
				year  string
				month string
				day   string
			)
			year = date[0]
			if len(date[1]) == 1 {
				month = "0" + date[1]
			} else {
				month = date[1]
			}
			if len(date[2]) == 1 {
				day = "0" + date[2]
			} else {
				day = date[2]
			}
			value := data.EMM00166466
			ra = append(ra, model.ResponseDateStringValueFloat{Date: year + month + day, TargetValue: value})
		}
	}
	responses = ManualHandleUpDown(ra)
	return
}

// APIEastMoneyIndustryTarget 东方财富行业指标数据接口
var APIEastMoneyIndustryTarget = "https://datacenter-web.eastmoney.com/api/data/v1/get?sortColumns=REPORT_DATE&sortTypes=-1&pageSize=500&pageNumber=PageNumber&reportName=RPT_INDUSTRY_INDEX&columns=REPORT_DATE%2CINDICATOR_VALUE%2CCHANGE_RATE%2CCHANGERATE_3M%2CCHANGERATE_6M%2CCHANGERATE_1Y%2CCHANGERATE_2Y%2CCHANGERATE_3Y&filter=(INDICATOR_ID%3D%22SourceTargetCode%22)"

// SpiderEastMoneyIndustryTarget 爬取"东方财富"网站的`行业指标`
// 适用指标：
// - 波罗的海干散货指数
//   • 页面展示接口：https://data.eastmoney.com/cjsj/hyzs_list_EMI00107664.html
//   • 数据获取接口：https://datacenter-web.eastmoney.com/api/data/v1/get?sortColumns=REPORT_DATE&sortTypes=-1&pageSize=500&pageNumber=1&reportName=RPT_INDUSTRY_INDEX&columns=REPORT_DATE%2CINDICATOR_VALUE%2CCHANGE_RATE%2CCHANGERATE_3M%2CCHANGERATE_6M%2CCHANGERATE_1Y%2CCHANGERATE_2Y%2CCHANGERATE_3Y&filter=(INDICATOR_ID%3D%22EMI00107664%22)
// - 物流业景气指数
//   • 页面展示接口：https://data.eastmoney.com/cjsj/hyzs_list_EMI00352262.html
//   • 数据获取接口：https://datacenter-web.eastmoney.com/api/data/v1/get?sortColumns=REPORT_DATE&sortTypes=-1&pageSize=500&pageNumber=1&reportName=RPT_INDUSTRY_INDEX&columns=REPORT_DATE%2CINDICATOR_VALUE%2CCHANGE_RATE%2CCHANGERATE_3M%2CCHANGERATE_6M%2CCHANGERATE_1Y%2CCHANGERATE_2Y%2CCHANGERATE_3Y&filter=(INDICATOR_ID%3D%22EMI00352262%22)
// - 中国大宗商品指数
//   • 页面展示接口：https://data.eastmoney.com/cjsj/hyzs_list_EMI00662535.html
//   • 数据获取接口：https://datacenter-web.eastmoney.com/api/data/v1/get?sortColumns=REPORT_DATE&sortTypes=-1&pageSize=500&pageNumber=1&reportName=RPT_INDUSTRY_INDEX&columns=REPORT_DATE%2CINDICATOR_VALUE%2CCHANGE_RATE%2CCHANGERATE_3M%2CCHANGERATE_6M%2CCHANGERATE_1Y%2CCHANGERATE_2Y%2CCHANGERATE_3Y&filter=(INDICATOR_ID%3D%22EMI00662535%22)
func SpiderEastMoneyIndustryTarget(sourceTargetCodeSpider string) (responses []model.Response) {
	for page := 1; ; page++ {
		url := strings.ReplaceAll(APIEastMoneyIndustryTarget, "SourceTargetCode", sourceTargetCodeSpider)
		url = strings.ReplaceAll(url, "PageNumber", strconv.Itoa(page))
		body, err := downloader.SimpleGet(url)
		if err != nil {
			logger.Error(err.Error())
			continue
		}
		stringBody := *(*string)(unsafe.Pointer(&body))
		var it eastMoneyIndustryTarget
		if err := json.Unmarshal([]byte(stringBody), &it); err != nil {
			logger.Error(err.Error())
			continue
		}
		if !it.Success {
			break
		}
		pageResponses := eastMoneyIndustryTargetMostPipeline(it, 8)
		responses = append(responses, pageResponses...)
	}
	return
}

// APIEastMoneyQiHuoTarget 东方财富期货指标数据接口
var APIEastMoneyQiHuoTarget = "http://futsseapi.eastmoney.com/static/%s"

// SpiderEastMoneyQiHuoTarget 爬取"东方财富"网站的`期货指标`
// 适用指标：
// - 铁矿石主力合约
//   • 页面展示接口：http://quote.eastmoney.com/qihuo/IM.html
//   • 数据获取接口：http://futsseapi.eastmoney.com/static/114_im_qt
// - 沪银主力
//   • 页面展示接口：http://quote.eastmoney.com/qihuo/agm.html
//   • 数据获取接口：http://futsseapi.eastmoney.com/static/113_agm_qt
// - 沪金主力
//   • 页面展示接口：http://quote.eastmoney.com/qihuo/aum.html
//   • 数据获取接口：http://futsseapi.eastmoney.com/static/113_aum_qt
// - 沪铜主力
//   • 页面展示接口：http://quote.eastmoney.com/qihuo/cum.html
//   • 数据获取接口：http://futsseapi.eastmoney.com/static/113_cum_qt
// - 沪镍主力
//   • 页面展示接口：http://quote.eastmoney.com/qihuo/nim.html
//   • 数据获取接口：http://futsseapi.eastmoney.com/static/113_nim_qt
// - 豆油主力
//   • 页面展示接口：http://quote.eastmoney.com/qihuo/ym.html
//   • 数据获取接口：http://futsseapi.eastmoney.com/static/114_ym_qt
// - 玉米主力
//   • 页面展示接口：http://quote.eastmoney.com/qihuo/cm.html
//   • 数据获取接口：http://futsseapi.eastmoney.com/static/114_cm_qt
// - 原油主力
//   • 页面展示接口：http://quote.eastmoney.com/qihuo/scm.html
//   • 数据获取接口：http://futsseapi.eastmoney.com/static/142_scm_qt
// - COMEX黄金电子盘主力合约
//   • 页面展示接口：http://quote.eastmoney.com/globalfuture/GC00Y.html
//   • 数据获取接口：http://futsseapi.eastmoney.com/static/101_GC00Y_qt
// - COMEX白银电子盘主力合约
//   • 页面展示接口：https://quote.eastmoney.com/globalfuture/SI00Y.html
//   • 数据获取接口：http://futsseapi.eastmoney.com/static/101_SI00Y_qt
// - NYMEX轻质原油电子盘主力
//   • 页面展示接口：http://quote.eastmoney.com/globalfuture/CL00Y.html
//   • 数据获取接口：https://futsseapi.eastmoney.com/static/102_CL00Y_qt
func SpiderEastMoneyQiHuoTarget(sourceTargetCodeSpider string) (responses []model.Response) {
	url := fmt.Sprintf(APIEastMoneyQiHuoTarget, sourceTargetCodeSpider)
	body, err := downloader.SimpleGet(url)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	stringBody := *(*string)(unsafe.Pointer(&body))
	var qh eastMoneyQiHuoTarget
	if err := json.Unmarshal([]byte(stringBody), &qh); err != nil {
		logger.Error(err.Error())
		return
	}
	var response model.Response
	response.Date = time.Now().Format("20060102")
	// 对于像股票、期货这类指标，数据实时更新，所以指标值包含：最新价、涨跌额、涨跌幅、最高、最低、昨收、更新日期
	// 这7个值会组合打包成一个字符串发送给Java服务器，然后解压切分组合字符串存入对应的表字段
	value := fmt.Sprintf("%.2f,%.2f,%.2f%s,%.2f,%.2f,%.2f,%s",
		qh.Qt.P, qh.Qt.ZDE, qh.Qt.ZDF, "%", qh.Qt.H, qh.Qt.L, qh.Qt.QRSPJ,
		time.Unix(qh.Qt.UTime, 0).Format("2006-01-02 15:04:05"))
	response.TargetValue = value
	responses = append(responses, response)
	return
}

// APIEastMoneyGlobalTarget 东方财富全球指数数据接口
var APIEastMoneyGlobalTarget = "http://push2.eastmoney.com/api/qt/stock/get?invt=2&fltt=1&fields=f43%2Cf44%2Cf45%2Cf46%2Cf60%2Cf86%2Cf169%2Cf170&"

// SpiderEastMoneyGlobalTarget 爬取"东方财富"网站的`全球指数`
// 适用指标：
// - 国债指数
//   • 页面展示接口：http://quote.eastmoney.com/zs000012.html
//   • 数据获取接口：http://push2.eastmoney.com/api/qt/stock/get?invt=2&fltt=1&fields=f43%2Cf44%2Cf45%2Cf46%2Cf60%2Cf86%2Cf169%2Cf170&secid=1.000012
// - 台湾加权
//   • 页面展示接口：http://quote.eastmoney.com/gb/zsTWII.html
//   • 数据获取接口：http://push2.eastmoney.com/api/qt/stock/get?invt=2&fltt=1&fields=f43%2Cf44%2Cf45%2Cf46%2Cf60%2Cf86%2Cf169%2Cf170&secid=100.TWII
// - 韩国KOSPI
//   • 页面展示接口：http://quote.eastmoney.com/gb/zsKS11.html
//   • 数据获取接口：http://push2.eastmoney.com/api/qt/stock/get?invt=2&fltt=1&fields=f43%2Cf44%2Cf45%2Cf46%2Cf60%2Cf86%2Cf169%2Cf170&secid=100.KS11
// - 俄罗斯RTS
//   • 页面展示接口：http://quote.eastmoney.com/gb/zsRTS.html
//   • 数据获取接口：http://push2.eastmoney.com/api/qt/stock/get?invt=2&fltt=1&fields=f43%2Cf44%2Cf45%2Cf46%2Cf60%2Cf86%2Cf169%2Cf170&secid=100.RTS
// - 澳大利亚标普200
//   • 页面展示接口：http://quote.eastmoney.com/gb/zsAS51.html
//   • 数据获取接口：http://push2.eastmoney.com/api/qt/stock/get?invt=2&fltt=1&fields=f43%2Cf44%2Cf45%2Cf46%2Cf60%2Cf86%2Cf169%2Cf170&secid=100.AS51
// - 路透CRB商品指数
//   • 页面展示接口：http://quote.eastmoney.com/gb/zsCRB.html
//   • 数据获取接口：http://push2.eastmoney.com/api/qt/stock/get?invt=2&fltt=1&fields=f43%2Cf44%2Cf45%2Cf46%2Cf60%2Cf86%2Cf169%2Cf170&secid=100.CRB
// - 中证国有企业综合指数
//   • 页面展示接口：http://quote.eastmoney.com/gb/zsHSCEI.html
//   • 数据获取接口：http://push2.eastmoney.com/api/qt/stock/get?invt=2&fltt=1&fields=f43%2Cf44%2Cf45%2Cf46%2Cf60%2Cf86%2Cf169%2Cf170&secid=100.HSCEI
// - 中证央企结构调整指数
//   • 页面展示接口：http://quote.eastmoney.com/zs000926.html?=
//   • 数据获取接口：http://push2.eastmoney.com/api/qt/stock/get?invt=2&fltt=1&fields=f43%2Cf44%2Cf45%2Cf46%2Cf60%2Cf86%2Cf169%2Cf170&secid=1.000926
// - 中证国企一带一路指数
//   • 页面展示接口：http://quote.eastmoney.com/zs399991.html
//   • 数据获取接口：http://push2.eastmoney.com/api/qt/stock/get?invt=2&fltt=1&fields=f43%2Cf44%2Cf45%2Cf46%2Cf60%2Cf86%2Cf169%2Cf170&secid=0.399991
// - 中证中央企业100指数
//   • 页面展示接口：http://quote.eastmoney.com/zs000903.html?=000903
//   • 数据获取接口：http://push2.eastmoney.com/api/qt/stock/get?invt=2&fltt=1&fields=f43%2Cf44%2Cf45%2Cf46%2Cf60%2Cf86%2Cf169%2Cf170&secid=1.000903
// - 原油指数
//   • 页面展示接口：http://quote.eastmoney.com/q/159.scfi.html
//   • 数据获取接口：http://push2.eastmoney.com/api/qt/stock/get?invt=2&fltt=1&fields=f43%2Cf44%2Cf45%2Cf46%2Cf60%2Cf86%2Cf169%2Cf170&secid=159.scfi
// - AMAC造纸
//   • 页面展示接口：http://quote.eastmoney.com/zz/2.H30049.html
//   • 数据获取接口：http://push2.eastmoney.com/api/qt/stock/get?invt=2&fltt=1&fields=f43%2Cf44%2Cf45%2Cf46%2Cf60%2Cf86%2Cf169%2Cf170&secid=2.H30049
func SpiderEastMoneyGlobalTarget(sourceTargetCodeSpider string) (responses []model.Response) {
	url := APIEastMoneyGlobalTarget + sourceTargetCodeSpider
	body, err := downloader.SimpleGet(url)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	stringBody := *(*string)(unsafe.Pointer(&body))
	var gt eastMoneyGlobalTarget
	if err := json.Unmarshal([]byte(stringBody), &gt); err != nil {
		logger.Error(err.Error())
		return
	}
	var response model.Response
	response.Date = time.Unix(gt.Data.F86, 0).Format("20060102")
	response.TargetValue = strings.Join([]string{
		fmt.Sprintf("%.2f", gt.Data.F43 / 100) ,
		fmt.Sprintf("%.2f", gt.Data.F169 / 100) ,
		fmt.Sprintf("%.2f", gt.Data.F170 / 100) ,
		fmt.Sprintf("%.2f", gt.Data.F44 / 100) ,
		fmt.Sprintf("%.2f", gt.Data.F45 / 100) ,
		fmt.Sprintf("%.2f", gt.Data.F60 / 100) ,
		time.Unix(gt.Data.F86, 0).Format("2006-01-02 15:04:05"),
	}, ",")
	responses = append(responses, response)
	return
}
