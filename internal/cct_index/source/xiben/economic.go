package xiben

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/api"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/model"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/pkg/downloader"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/pkg/safeguard"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/poster"
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
	"strings"
	"time"
	"unsafe"
)

func xiBenEconomic(ic *model.IndexConfig) (buffers []*model.Buffer) {
	url := strings.ReplaceAll(api.XiBen, "#", ic.SourceTargetCodeSpider)
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

	stringBody := *(*string)(unsafe.Pointer(&body))
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(stringBody))
	if err != nil {
		logger.Error(err.Error())
		go poster.Poster(ic)
		return
	}

	var table []string
	dom.Find("tbody[id='indexdetaildata'] td").Each(func(i int, selection *goquery.Selection) {
		text := selection.Text()
		table = append(table, text)
	})
	if table == nil || len(table) == 0 {
		go poster.Poster(ic)
		return
	}

	date, value, upDownValue, upDownRange := "", "", "", ""
	for index, v := range table {
		buffer := &model.Buffer{}
		switch index % 4 {
		case 0:
			date = v
		case 1:
			value = v
		case 2:
			upDownValue = v
		case 3:
			upDownRange = v
			buffer.Date = strings.ReplaceAll(date, "-", "")
			buffer.TargetValue = strings.Join([]string{
				value,
				upDownValue,
				upDownRange,
				"",
				"",
				"",
				date + " " + time.Now().Format("15:04:05"),
			}, ",")
			buffers = append(buffers, buffer)
		}
	}

	return
}
