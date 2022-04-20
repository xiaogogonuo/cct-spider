package fx678

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/xiaogogonuo/cct-spider/internal/target/model"
	"github.com/xiaogogonuo/cct-spider/internal/target/pkg/downloader"
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
	"strings"
	"time"
	"unsafe"
)

// APIFxExchangeTarget 汇通财经数据接口
var APIFxExchangeTarget = "https://quote.fx678.com/exchange/"

// SpiderFxExchangeTarget 爬取"汇通财经"网站的外汇、债券、原油、期货、外盘、汇率等指标
// 适用指标：
// - 美元人民币、港元人民币、欧元人民币
//   • 页面展示接口：https://quote.fx678.com/exchange/WHMP
//   • 数据获取接口：https://quote.fx678.com/exchange/WHMP
// - 欧元美元、美元日元、英镑美元、美元指数
//   • 页面展示接口：https://quote.fx678.com/exchange/WH
//   • 数据获取接口：https://quote.fx678.com/exchange/WH
// - 布伦特原油连续
//   • 页面展示接口：https://quote.fx678.com/exchange/MAINOIL
//   • 数据获取接口：https://quote.fx678.com/exchange/MAINOIL
// - 纽约黄金连续、纽约白银连续
//   • 页面展示接口：https://quote.fx678.com/exchange/MAINGOLD
//   • 数据获取接口：https://quote.fx678.com/exchange/MAINGOLD
// - LME铜、LME镍、LME铝
//   • 页面展示接口：https://quote.fx678.com/exchange/LME
//   • 数据获取接口：https://quote.fx678.com/exchange/LME
// - 美玉米连续、美黄豆连续
//   • 页面展示接口：https://quote.fx678.com/exchange/CBOT
//   • 数据获取接口：https://quote.fx678.com/exchange/CBOT
// - 美债10年收益率、日债10年收益率、德债10年收益率、英债10年收益率
//   • 页面展示接口：https://quote.fx678.com/exchange/GJZQ
//   • 数据获取接口：https://quote.fx678.com/exchange/GJZQ
// - 日经225、英国FTSE100、德国DAX30、法国CAC40、意大利MIB、(加拿大)S&P/TSX 60、斯托克600、纳斯达克指数、道琼斯工业指数、标普500、恒生指数
//   • 页面展示接口：https://quote.fx678.com/exchange/GJZS
//   • 数据获取接口：https://quote.fx678.com/exchange/GJZS
func SpiderFxExchangeTarget(sourceTargetCodeSpider, targetNameSpider string) (responses []model.Response) {
	url := APIFxExchangeTarget + sourceTargetCodeSpider
	body, err := downloader.SimpleGet(url)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	stringBody := *(*string)(unsafe.Pointer(&body))
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(stringBody))
	if err != nil {
		logger.Error(err.Error())
		return
	}
	var table []string
	dom.Find("table tr[id] td").Each(func(i int, selection *goquery.Selection) {
		table = append(table, strings.Trim(selection.Text(), "\n\t "))
	})
	var data string
	for i := 0; i < len(table)/8; i++ {
		if table[i*8] == targetNameSpider {
			record := table[i*8+1 : (i+1)*8]
			record[len(record)-1] = time.Now().Format("2006-01-02") + " " + record[len(record)-1]
			data = strings.Join(record, ",")
			break
		}
	}
	var response model.Response
	response.TargetValue = data
	response.Date = time.Now().Format("20060102")
	responses = append(responses, response)
	return
}
