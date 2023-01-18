package sina

import "github.com/xiaogogonuo/cct-spider/internal/cct_index/model"

// SpiderSina 新浪财经
func SpiderSina(ic *model.IndexConfig) (buffers []*model.Buffer) {
	switch ic.TargetCode {
	// 新浪财经 - 汇率
	case "HY00008", "HG00086", "HG00087":
		buffers = sinaForex(ic)
		// 新浪财经 - 宏观经济
	case "HG00002", "HG00040":
		buffers = sinaEconomic(ic)
	}
	return
}
