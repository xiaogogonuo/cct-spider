package fx

import (
	"encoding/json"
	"fmt"
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

// 适用：
// - 美元人民币、
// - 港元人民币
// - 欧元人民币
// - 欧元美元
// - 美元日元
// - 英镑美元
// - 美元指数
// - 布伦特原油连续
// - 纽约黄金连续
// - 纽约白银连续
// - LME铜
// - LME镍
// - LME铝
// - 美玉米连续
// - 美黄豆连续
// - 美债10年收益率
// - 日债10年收益率
// - 德债10年收益率
// - 英债10年收益率
// - 日经225
// - 英国FTSE100
// - 德国DAX30
// - 法国CAC40
// - 意大利MIB
// - 加拿大SP/TSX
// - 纳斯达克指数
// - 道琼斯工业指数
// - 标普500
// -恒生指数
func fxExchange(ic *model.IndexConfig) (buffers []*model.Buffer) {
	url := strings.ReplaceAll(api.FxExchange, "#", ic.SourceTargetCodeSpider)
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
	dom.Find("table tr[id] td").Each(func(i int, selection *goquery.Selection) {
		table = append(table, strings.Trim(selection.Text(), "\n\t "))
	})
	if table == nil || len(table) == 0 {
		go poster.Poster(ic)
		return
	}

	var data string
	for i := 0; i < len(table)/8; i++ {
		if table[i*8] == ic.TargetNameSpider {
			record := table[i*8+1 : (i+1)*8]
			record[len(record)-1] = time.Now().Format("2006-01-02") + " " + record[len(record)-1]
			data = strings.Join(record, ",")
			break
		}
	}

	buffer := &model.Buffer{}
	buffer.TargetValue = data
	buffer.Date = time.Now().Format("20060102")
	buffers = append(buffers, buffer)

	return
}

func fxExchangeSpecial(ic *model.IndexConfig) (buffers []*model.Buffer) {
	url := strings.ReplaceAll(api.FxExchangeSpecial, "#", ic.SourceTargetCodeSpider)
	url = strings.ReplaceAll(url, "$", ic.SourceTargetCode)
	body, err := downloader.Get(url, map[string]string{
		"Referer":    api.FxExchangeSpecialRefer,
		"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.54 Safari/537.36",
	})

	if err != nil {
		if !safeguard.IsNetworkNormal() {
			logger.Error("请检查服务器的网络是否能联通外网")
			return
		}
		logger.Error(err.Error())
		go poster.Poster(ic)
		return
	}

	var es ExchangeSpecial
	if err = json.Unmarshal(body, &es); err != nil {
		logger.Error(err.Error())
		go poster.Poster(ic)
		return
	}

	if !es.StatusValid() {
		logger.Error(fmt.Sprintf("%s status invalid", ic.SourceTargetCode))
		go poster.Poster(ic)
		return
	}
	if !es.DataValid() {
		logger.Error(fmt.Sprintf("%s data invalid", ic.SourceTargetCode))
		go poster.Poster(ic)
		return
	}

	data, err := es.Handler()
	if err != nil {
		logger.Error(fmt.Sprintf("%s data invalid", ic.SourceTargetCode))
		go poster.Poster(ic)
		return
	}

	buffer := &model.Buffer{}
	buffer.Date = time.Now().Format("20060102")
	buffer.TargetValue = data
	buffers = append(buffers, buffer)

	return
}
