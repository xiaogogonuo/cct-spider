package sina

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/api"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/model"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/pkg/decoder"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/pkg/downloader"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/pkg/safeguard"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/poster"
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
	"strings"
)

const Pages = 7

// 宏观经济
// 适用：
// - 地区生产总值
// - 地区居民消费价格指数
func sinaEconomic(ic *model.IndexConfig) (buffers []*model.Buffer) {
	for page := 0; ; page++ {
		if page >= Pages {
			break
		}
		url := strings.ReplaceAll(api.Marco, "#", ic.SourceTargetCodeSpider)
		url = strings.ReplaceAll(url, "$", fmt.Sprintf("%d", page*31))
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

		body, err = decoder.ToGBK(body)
		if err != nil {
			logger.Error(err.Error())
			go poster.Poster(ic)
			return
		}

		body = filter(body)

		switch ic.TargetCode {
		case "HG00002": // 地区生产总值
			buffers, err = getRegionGDPBuffer(body)
			if err != nil {
				logger.Error(err.Error())
				go poster.Poster(ic)
				return
			}
		case "HG00040": // 地区居民消费价格指数
			buffers, err = getRegionCPIBuffer(body)
			if err != nil {
				logger.Error(err.Error())
				go poster.Poster(ic)
				return
			}
		}
	}

	return
}
