package ppi100

import "github.com/xiaogogonuo/cct-spider/internal/cct_index/model"

// SpiderPPI100 生意社
func SpiderPPI100(ic *model.IndexConfig) (buffers []*model.Buffer) {
	switch ic.TargetCode {
	case "HG00096", "HG00097":
		buffers = ppi100Economic(ic)
	}
	return
}
