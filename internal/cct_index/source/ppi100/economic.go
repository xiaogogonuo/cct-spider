package ppi100

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/api"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/model"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/pkg/downloader"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/pkg/safeguard"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/poster"
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
	"net/url"
	"strconv"
)

const Pages = 1

// 适用：
// - 规模以上工业增加值当月同比增速
// - 国有及国有控股企业工业增加值同比增速
func ppi100Economic(ic *model.IndexConfig) (buffers []*model.Buffer) {
	var texts []string
	for page := 1; ; page++ {
		if page > Pages {
			break
		}
		values := url.Values{}
		values.Add("cid", ic.SourceTargetCodeSpider)
		values.Add("p", strconv.Itoa(page))
		header := map[string]string{
			"Content-Type": "application/x-www-form-urlencoded; charset=UTF-8",
		}
		body, err := downloader.Post(api.PPI00, []byte(values.Encode()), header)

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

		dom.Find("td").Each(func(i int, selection *goquery.Selection) {
			text := selection.Text()
			texts = append(texts, text)
		})
	}

	switch ic.TargetCode {
	case "HG00096": // 规模以上工业增加值当月同比增速
		buffers = append(buffers, getPPI100Buffer(texts, 1, 2, 20)...)
	case "HG00097": // 国有及国有控股企业工业增加值同比增速
		buffers = append(buffers, getPPI100Buffer(texts, 1, 9, 20)...)
	}

	return
}
