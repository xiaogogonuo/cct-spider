package wuliu

import "github.com/xiaogogonuo/cct-spider/internal/cct_index/model"

// SpiderWuLiu 中国物流
func SpiderWuLiu(ic *model.IndexConfig) (buffers []*model.Buffer) {
	switch ic.TargetCode {
	case "HG00094":
		buffers = wuLiuEconomic(ic)
	}
	return
}
