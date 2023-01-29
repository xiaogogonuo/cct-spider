package stats

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/api"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/model"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/pkg/downloader"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/pkg/safeguard"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/poster"
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
	"regexp"
	"strings"
	"time"
	"unsafe"
)

const (
	PageNum = 1
	PerPage = 30
)

type Detail struct {
	Href     string
	Title    string
	DateTime time.Time
}

// 适用：
// - 规模以上工业增加值当月同比增速
// - 国有及国有控股企业工业增加值同比增速
func industrialOutput(ic *model.IndexConfig) (buffers []*model.Buffer) {
	var details []Detail
	only := make(map[string]struct{})
	atLeast := 0
	for i := 1; i <= PageNum; i++ {
		url := strings.ReplaceAll(api.IndustrialOutput, "#", fmt.Sprintf("%d", i))
		url = strings.ReplaceAll(url, "$", fmt.Sprintf("%d", PerPage))
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

		// var urlstr = 'http://www.stats.gov.cn/tjsj/zxfb/202211/t20221115_1890235.html';
		pattern := `var urlstr = '(http://www.stats.gov.cn/.*?)';`
		re := regexp.MustCompile(pattern)
		matched := re.FindAllStringSubmatch(stringBody, -1)
		if matched == nil || len(matched) == 0 {
			go poster.Poster(ic)
			return
		}
		for _, match := range matched {
			if match == nil || len(match) == 0 {
				continue
			}
			href := match[1]
			if href == "" {
				continue
			}
			atLeast += 1
			if _, ok := only[href]; !ok {
				only[href] = struct{}{}
				details = append(details, Detail{Href: href})
			}
		}
	}

	if atLeast == 0 {
		go poster.Poster(ic)
		return
	}

	buffers = travelDetail(ic, details)

	return
}

func travelDetail(ic *model.IndexConfig, details []Detail) (buffers []*model.Buffer) {
	for _, detail := range details {
		idx := strings.Index(detail.Href, "zxfb")
		date := detail.Href[idx+5 : idx+11]
		t, err := time.Parse("200601", date)
		if err != nil {
			continue
		}
		if t.Month() == 3 {
			continue
		}
		detail.DateTime = t.AddDate(0, -1, 0)

		body, err := downloader.SimpleGet(detail.Href)

		if err != nil {
			if !safeguard.IsNetworkNormal() {
				logger.Error("请检查服务器的网络是否能联通外网")
				return
			}
			logger.Error(err.Error())
			go poster.Poster(ic)
			return
		}

		dom, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
		if err != nil {
			logger.Error(err.Error())
			go poster.Poster(ic)
			return
		}

		var paragraph []string
		dom.Find("div[class='TRS_Editor']").Each(func(i int, selection *goquery.Selection) {
			paragraph = append(paragraph, selection.Text())
		})
		paragraphs := strings.Join(paragraph, "")
		paragraphs = strings.ReplaceAll(paragraphs, " ", "")
		paragraphs = strings.ReplaceAll(paragraphs, ",", "，")
		paragraphs = strings.ReplaceAll(paragraphs, "，", "")
		paragraphs = strings.ReplaceAll(paragraphs, "实际", "")

		var pattern string
		switch ic.TargetCode {
		case "HG00096":
			pattern = `(-|—{0,1})([0-9]+)月份规模以上工业增加值同比(增长|下降)([0-9]*\.[0-9]+|[0-9]+)%?`
		case "HG00097":
			pattern = `([0-9]+)月份(国有|国有及国有)控股企业增加值同比(增长|下降|持平)([0-9]*\.[0-9]+|[0-9]*)%?`
		}
		re := regexp.MustCompile(pattern)
		matched := re.FindAllStringSubmatch(paragraphs, -1)

		var buffer *model.Buffer
		switch ic.TargetCode {
		case "HG00096":
			buffer = parseIndustriesAboveDesignatedSize(detail, matched)
		case "HG00097":
			buffer = parseStateOwnedAndHoldingEnterprises(detail, matched)
		}

		if buffer != nil {
			buffers = append(buffers, buffer)
		}
	}

	return
}

// 解析规模以上工业增加值
func parseIndustriesAboveDesignatedSize(detail Detail, matched [][]string) (buffer *model.Buffer) {
	if matched == nil || len(matched) == 0 {
		logger.Warn(detail.Href + "doesn't match any data!")
		return
	}
	for _, match := range matched {
		if match[1] == "" {
			v := match[4]
			if match[3] != "增长" {
				v = "-" + v
			}
			buffer = &model.Buffer{}
			buffer.Date = detail.DateTime.Format("200601")
			buffer.TargetValue = v
			return
		}
	}
	return
}

// 解析国有及国有控股企业工业增加值
func parseStateOwnedAndHoldingEnterprises(detail Detail, matched [][]string) (buffer *model.Buffer) {
	if matched == nil || len(matched) == 0 {
		logger.Warn(detail.Href + "doesn't match any data!")
		return
	}
	for _, match := range matched {
		v := match[4]
		switch match[3] {
		case "下降":
			v = "-" + v
		case "持平":
			v = "0"
		}

		buffer = &model.Buffer{}
		buffer.Date = detail.DateTime.Format("200601")
		buffer.TargetValue = v
		return

	}
	return
}
