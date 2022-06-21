package chinawuliu

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/xiaogogonuo/cct-spider/internal/target/model"
	"github.com/xiaogogonuo/cct-spider/internal/target/pkg/downloader"
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
	"regexp"
	"strconv"
	"strings"
	"unsafe"
)

// APIChinaWuLiuTarget 中国物流指标接口
var APIChinaWuLiuTarget = "http://www.chinawuliu.com.cn/xsyj/tjsj/"

// SpiderChinaWuLiuTarget 爬取中国物流指标
// 适用指标：
// - 中国非制造业商务活动指数月度
//   • 页面展示接口：http://www.chinawuliu.com.cn/xsyj/tjsj/
//   • 数据获取接口：http://www.chinawuliu.com.cn/xsyj/tjsj/
func SpiderChinaWuLiuTarget(targetNameSpider string) (responses []model.Response) {
	for i := 1; i <= 10; i++ {
		var url string
		if i == 1 {
			url = APIChinaWuLiuTarget
		} else {
			url = APIChinaWuLiuTarget + "index_" + strconv.Itoa(i) + ".shtml"
		}
		body, err := downloader.SimpleGet(url)
		if err != nil {
			logger.Error(err.Error())
			continue
		}
		dom, err := goquery.NewDocumentFromReader(strings.NewReader(*(*string)(unsafe.Pointer(&body))))
		if err != nil {
			logger.Error(err.Error())
			continue
		}
		selector := "div[class='col-sm-8 leftRow'] ul[class='list-box list-box--pre'] a"
		dom.Find(selector).Each(func(i int, selection *goquery.Selection) {
			text := selection.Text()
			if strings.Contains(text, targetNameSpider) {
				pattern := "([0-9]{4})年([0-9]{1,2})月份中国" + targetNameSpider + "为(.*)?%"
				reg := regexp.MustCompile(pattern)
				matched := reg.FindAllStringSubmatch(text, -1)
				if len(matched) > 0 {
					match := matched[0]
					var response model.Response
					var month string
					if len(match[2]) == 1 {
						month = "0" + match[2]
					} else {
						month = match[2]
					}
					response.TargetValue = strings.ReplaceAll(match[3], "%", "")
					response.Date = match[1] + month
					responses = append(responses, response)
				}
			}
		})
	}
	return
}
