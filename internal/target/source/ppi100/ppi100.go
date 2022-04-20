package ppi100

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"github.com/xiaogogonuo/cct-spider/internal/target/model"
	"github.com/xiaogogonuo/cct-spider/internal/target/pkg/downloader"
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
	"net/url"
	"strconv"
)

// APIPPI100Target 宏观数据网站的宏观指标接口
var APIPPI100Target = "http://www.100ppi.com/mac/?welcome=no&f=list_data"

// SpiderPPI100Target 爬取宏观数据网站的宏观指标
// 适用指标：
// - 规模以上工业增加值当月同比增速
//   • 页面展示接口：http://www.100ppi.com/mac/world_gj---111A.html 工业增加值同比增长
//   • 数据获取接口：http://www.100ppi.com/mac/?welcome=no&f=list_data
// - 国有及国有控股企业工业增加值同比增速
//   • 页面展示接口：http://www.100ppi.com/mac/world_gj---111A.html 国有及国有控股企业同比增长
//   • 数据获取接口：http://www.100ppi.com/mac/?welcome=no&f=list_data
func SpiderPPI100Target(targetCode, sourceTargetCodeSpider string, pages int) (responses []model.Response) {
	var texts []string
	for page := 1; ; page++ {
		if page > pages && pages != -1 {
			break
		}
		values := url.Values{}
		values.Add("cid", sourceTargetCodeSpider)
		values.Add("p", strconv.Itoa(page))
		header := map[string]string{
			"Content-Type": "application/x-www-form-urlencoded; charset=UTF-8",
		}
		body, err := downloader.Post(APIPPI100Target, []byte(values.Encode()), header)
		if err != nil {
			logger.Error(err.Error())
			continue
		}
		dom, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
		if err != nil {
			logger.Error(err.Error())
			continue
		}
		dom.Find("td").Each(func(i int, selection *goquery.Selection) {
			text := selection.Text()
			texts = append(texts, text)
		})
	}
	switch targetCode {
	case "HG00096": // 规模以上工业增加值当月同比增速
		responses = append(responses, ppiPipeline(texts, 1,2, 20)...)
	case "HG00097": // 国有及国有控股企业工业增加值同比增速
		responses = append(responses, ppiPipeline(texts, 1,9, 20)...)
	}
	return
}
