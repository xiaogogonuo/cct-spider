package wuliu

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/api"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/model"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/pkg/downloader"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/pkg/safeguard"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/poster"
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
	"regexp"
	"strconv"
	"strings"
	"unsafe"
)

func wuLiuEconomic(ic *model.IndexConfig) (buffers []*model.Buffer) {
	for i := 1; i <= 10; i++ {
		var url string
		if i == 1 {
			url = api.WuLiu
		} else {
			url = api.WuLiu + "index_" + strconv.Itoa(i) + ".shtml"
		}
		body, err := downloader.SimpleGet(url)

		if err != nil {
			if !safeguard.IsNetworkNormal() {
				logger.Error("请检查服务器的网络是否能联通外网")
				return
			}
			logger.Error(err.Error())
			go poster.Poster(ic)
			return
		}

		dom, err := goquery.NewDocumentFromReader(strings.NewReader(*(*string)(unsafe.Pointer(&body))))
		if err != nil {
			logger.Error(err.Error())
			go poster.Poster(ic)
			return
		}

		selector := "div[class='col-sm-8 leftRow'] ul[class='list-box list-box--pre'] a"
		dom.Find(selector).Each(func(i int, selection *goquery.Selection) {
			text := selection.Text()
			if strings.Contains(text, ic.TargetNameSpider) {
				pattern := "([0-9]{4})年([0-9]{1,2})月份中国" + ic.TargetNameSpider + "为(.*)?%"
				reg := regexp.MustCompile(pattern)
				matched := reg.FindAllStringSubmatch(text, -1)
				if len(matched) > 0 {
					match := matched[0]
					buffer := &model.Buffer{}
					var month string
					if len(match[2]) == 1 {
						month = "0" + match[2]
					} else {
						month = match[2]
					}
					buffer.TargetValue = strings.ReplaceAll(match[3], "%", "")
					buffer.Date = match[1] + month
					buffers = append(buffers, buffer)
				}
			}
		})
	}

	return
}
