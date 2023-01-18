package sci

import "github.com/xiaogogonuo/cct-spider/internal/cct_index/model"

// SpiderSCI 卓创资讯
func SpiderSCI(ic *model.IndexConfig) (buffers []*model.Buffer) {
	switch ic.TargetCode {
	// 卓创资讯 - 原油价格指数
	case "HY00004":
		buffers = sciIndustryCOI(ic)
	// 卓创资讯 - 造纸行业价格指数
	case "HY00010":
		buffers = sciIndustryPII(ic)
	}
	return
}
