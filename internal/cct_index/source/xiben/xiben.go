package xiben

import "github.com/xiaogogonuo/cct-spider/internal/cct_index/model"

// SpiderXiBen 西本资讯
func SpiderXiBen(ic *model.IndexConfig) (buffers []*model.Buffer) {
	switch ic.TargetCode {
	case "HG00092", "HG00093":
		buffers = xiBenEconomic(ic)
	}
	return
}
