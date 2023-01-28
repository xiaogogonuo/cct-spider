package irc

import "github.com/xiaogogonuo/cct-spider/internal/cct_index/model"

// SpiderIrc 银保监会
func SpiderIrc(ic *model.IndexConfig) (buffers []*model.Buffer) {
	switch ic.TargetCode {
	case "HG00095":
		buffers = ircEconomic(ic)
	}
	return
}
