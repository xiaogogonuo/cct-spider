package nrc

import "github.com/xiaogogonuo/cct-spider/internal/cct_index/model"

// SpiderNrc 中华人民共和国国家发展和改革委员会
func SpiderNrc(ic *model.IndexConfig) (buffers []*model.Buffer) {
	switch ic.TargetCode {
	case "HG00118", "HG00119", "HG00120":
		buffers = nrcFinance(ic)
	}
	return
}
