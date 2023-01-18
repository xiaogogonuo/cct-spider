package fx

import "github.com/xiaogogonuo/cct-spider/internal/cct_index/model"

// SpiderFx 汇通财经
func SpiderFx(ic *model.IndexConfig) (buffers []*model.Buffer) {
	switch ic.TargetCode {
	// 汇通财经 - 斯托克600
	case "HG00042":
		buffers = fxExchangeSpecial(ic)
	default:
		buffers = fxExchange(ic)
	}
	return
}
