package stats

import "github.com/xiaogogonuo/cct-spider/internal/cct_index/model"

// SpiderStats 国家统计局
func SpiderStats(ic *model.IndexConfig) (buffers []*model.Buffer) {
	switch ic.TargetCode {
	case "HG00096", "HG00097":
		buffers = industrialOutput(ic)
	}
	return
}
