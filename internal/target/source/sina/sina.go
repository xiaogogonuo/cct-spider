package sina

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/target/model"
	"github.com/xiaogogonuo/cct-spider/internal/target/pkg/calculator"
	"github.com/xiaogogonuo/cct-spider/internal/target/pkg/decoder"
	"github.com/xiaogogonuo/cct-spider/internal/target/pkg/downloader"
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
	"strconv"
	"strings"
	"unsafe"
)

// APISinaTargetForex 新浪财经指标`外汇指标`数据获取接口
var APISinaTargetForex = "https://hq.sinajs.cn/rn=0list=fx_s"

// APISinaTargetForexShow 新浪财经指标`外汇指标`页面展示接口
var APISinaTargetForexShow = "https://finance.sina.com.cn/money/forex/hq/%s.shtml"

// SpiderSinaTargetForex 爬取"新浪财经"网站的`外汇指标`
// 适用指标：
// - 人民币汇率
//   • 页面展示接口：https://finance.sina.com.cn/money/forex/hq/CNYUSD.shtml
//   • 数据获取接口：https://hq.sinajs.cn/rn=0list=fx_scnyusd
//     - 数据获取接口携带的请求头
//       + Referer: https://finance.sina.com.cn/money/forex/hq/CNYUSD.shtml
// - 人民币港元
//   • 页面展示接口：https://finance.sina.com.cn/money/forex/hq/CNYHKD.shtml
//   • 数据获取接口：https://hq.sinajs.cn/rn=0list=fx_scnyhkd
//     - 数据获取接口携带的请求头
//       + Referer: https://finance.sina.com.cn/money/forex/hq/CNYHKD.shtml
// - 人民币日元
//   • 页面展示接口：https://finance.sina.com.cn/money/forex/hq/CNYJPY.shtml
//   • 数据获取接口：https://hq.sinajs.cn/rn=0list=fx_scnyjpy
//     - 数据获取接口携带的请求头
//       + Referer: https://finance.sina.com.cn/money/forex/hq/CNYJPY.shtml
func SpiderSinaTargetForex(sourceTargetCodeSpider string) (responses []model.Response) {
	url := APISinaTargetForex + strings.ToLower(sourceTargetCodeSpider)
	body, err := downloader.Get(url, map[string]string{"Referer": fmt.Sprintf(APISinaTargetForexShow, strings.ToUpper(sourceTargetCodeSpider))})
	if err != nil {
		logger.Error(err.Error())
		return
	}
	stringBody := *(*string)(unsafe.Pointer(&body))
	dataStartIndex := strings.Index(stringBody, "\"")
	data := stringBody[dataStartIndex+1 : len(stringBody)-3]
	cutData := strings.Split(data, ",")
	var response model.Response
	response.Date = strings.ReplaceAll(cutData[len(cutData)-1], "-", "")
	currentPrice := cutData[8]   // 现价
	upDown := cutData[11]        // 涨跌
	upDownPercent := cutData[10] // 涨跌幅
	high := cutData[6]           // 最高
	low := cutData[7]            // 最低
	yesterday := cutData[3]      // 昨收
	currentPriceF, _ := strconv.ParseFloat(currentPrice, 64)
	upDownF, _ := strconv.ParseFloat(upDown, 64)
	upDownPercentF, _ := strconv.ParseFloat(upDownPercent, 64)
	highF, _ := strconv.ParseFloat(high, 64)
	lowF, _ := strconv.ParseFloat(low, 64)
	yesterdayF, _ := strconv.ParseFloat(yesterday, 64)
	response.TargetValue = strings.Join([]string{
		strconv.FormatFloat(calculator.Round(currentPriceF, 4), 'f', -1, 64),
		strconv.FormatFloat(calculator.Round(upDownF, 4), 'f', -1, 64),
		strconv.FormatFloat(calculator.Round(upDownPercentF, 4), 'f', -1, 64) + "%",
		strconv.FormatFloat(calculator.Round(highF, 4), 'f', -1, 64),
		strconv.FormatFloat(calculator.Round(lowF, 4), 'f', -1, 64),
		strconv.FormatFloat(calculator.Round(yesterdayF, 4), 'f', -1, 64),
		cutData[len(cutData)-1] + " " + cutData[0], // 更新时间
	}, ",")
	responses = append(responses, response)
	return
}

// APISinaEconomicTarget 新浪财经宏观经济指标接口
var APISinaEconomicTarget = "https://quotes.sina.cn/mac/api/jsonp_v3.php/SINAREMOTECALL/MacPage_Service.get_pagedata?cate=price&%s&from=%d&num=31&condition="

// SpiderSinaEconomicTarget 爬取新浪财经宏观指标
// 适用指标：
// - 地区居民消费价格指数
//   • 页面展示接口：http://finance.sina.com.cn/mac/#price-2-0-31-1
//   • 数据获取接口：https://quotes.sina.cn/mac/api/jsonp_v3.php/SINAREMOTECALL/MacPage_Service.get_pagedata?cate=price&event=2&from=1&num=31&condition=
// - 地区生产总值
//   • 页面展示接口：http://finance.sina.com.cn/mac/#nation-7-0-31-3
//   • 数据获取接口：https://quotes.sina.cn/mac/api/jsonp_v3.php/SINAREMOTECALL/MacPage_Service.get_pagedata?cate=price&event=7&from=0&num=31&condition=
func SpiderSinaEconomicTarget(sourceTargetCodeSpider, targetCode string, pages int) (responses []model.Response) {
	for page := 0; ; page++ {
		if page >= pages && pages != -1 {
			break
		}
		url := fmt.Sprintf(APISinaEconomicTarget, sourceTargetCodeSpider, page*31)
		body, err := downloader.SimpleGet(url)
		if err != nil {
			logger.Error(err.Error())
			continue
		}
		body, err = decoder.ToGBK(body)
		if err != nil {
			logger.Error(err.Error())
			continue
		}
		body = filter(body)
		var _responses []model.Response
		switch targetCode {
		case "HG00002": // 地区生产总值
			_responses = sinaEconomicTargetRegionGDPPipeline(body)
		case "HG00040": // 地区居民消费价格指数
			_responses = sinaEconomicTargetRegionCPIPipeline(body)
		}
		responses = append(responses, _responses...)
	}
	return
}
