package xiben

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/xiaogogonuo/cct-spider/internal/target/model"
	"github.com/xiaogogonuo/cct-spider/internal/target/pkg/downloader"
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
	"strings"
	"time"
	"unsafe"
)

// APIXiBenTarget 西本新干线指标接口
var APIXiBenTarget = "http://www.96369.net/indices/%s"

// SpiderXiBenTarget 爬取西本新干线指标
// 适用指标：
// - 固定资产投资完成额累计同比增速月度
//   • 页面展示接口：http://www.96369.net/indices/171
//   • 数据获取接口：http://www.96369.net/indices/171
// - 新增人民币贷款
//   • 页面展示接口：http://www.96369.net/indices/174
//   • 数据获取接口：http://www.96369.net/indices/174
func SpiderXiBenTarget(sourceTargetCodeSpider string) (responses []model.Response) {
	url := fmt.Sprintf(APIXiBenTarget, sourceTargetCodeSpider)
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
	dom.Find("table[class='mod_tab'] td").Each(func(i int, selection *goquery.Selection) {
		text := selection.Text()
		table = append(table, text)
	})
	var response model.Response
	date, value, riseFall, riseFallPercent := "", "", "", ""
	for index, v := range table {
		switch index % 4 {
		case 0:
			date = v
		case 1:
			value = v
		case 2:
			riseFall = v
		case 3:
			riseFallPercent = v
			response.Date = strings.ReplaceAll(date, "-","")
			response.TargetValue = strings.Join([]string{
				value,
				riseFall,
				riseFallPercent,
				"",
				"",
				"",
				date + " " + time.Now().Format("15:04:05"),
			}, ",")
			responses = append(responses, response)
		}
	}
	return
}
