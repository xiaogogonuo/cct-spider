package sina

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/api"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/model"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/pkg/arithmetic"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/pkg/downloader"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/pkg/safeguard"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/poster"
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

// 汇率
// 适用：
// - 人民币汇率
// - 人民币港元
// - 人民币日元
func sinaForex(ic *model.IndexConfig) (buffers []*model.Buffer) {
	url := strings.ReplaceAll(api.Forex, "#", strings.ToLower(ic.SourceTargetCodeSpider))
	url = strings.ReplaceAll(url, "$", fmt.Sprintf("%d", time.Now().UnixNano()/1e6))
	body, err := downloader.Get(url, map[string]string{
		"Referer": strings.ReplaceAll(api.ForexReferer, "#", ic.SourceTargetCodeSpider),
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

	stringBody := *(*string)(unsafe.Pointer(&body))
	dataStartIndex := strings.Index(stringBody, "\"")
	data := stringBody[dataStartIndex+1 : len(stringBody)-3]
	dataList := strings.Split(data, ",")

	buffer := &model.Buffer{}
	buffer.Date = strings.ReplaceAll(dataList[len(dataList)-1], "-", "")
	cur := dataList[1]          // 现价
	upDownValue := dataList[11] // 涨跌额
	upDownRange := dataList[10] // 涨跌幅
	high := dataList[6]         // 最高
	low := dataList[7]          // 最低
	pre := dataList[3]          // 昨收
	curF, _ := strconv.ParseFloat(cur, 64)
	upDownValueF, _ := strconv.ParseFloat(upDownValue, 64)
	upDownRangeF, _ := strconv.ParseFloat(upDownRange, 64)
	highF, _ := strconv.ParseFloat(high, 64)
	lowF, _ := strconv.ParseFloat(low, 64)
	preF, _ := strconv.ParseFloat(pre, 64)
	buffer.TargetValue = strings.Join([]string{
		strconv.FormatFloat(arithmetic.Round(curF, 4), 'f', -1, 64),
		strconv.FormatFloat(arithmetic.Round(upDownValueF, 4), 'f', -1, 64),
		strconv.FormatFloat(arithmetic.Round(upDownRangeF, 4), 'f', -1, 64) + "%",
		strconv.FormatFloat(arithmetic.Round(highF, 4), 'f', -1, 64),
		strconv.FormatFloat(arithmetic.Round(lowF, 4), 'f', -1, 64),
		strconv.FormatFloat(arithmetic.Round(preF, 4), 'f', -1, 64),
		dataList[len(dataList)-1] + " " + dataList[0], // 更新时间
	}, ",")

	buffers = append(buffers, buffer)

	return
}
