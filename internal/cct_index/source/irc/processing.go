package irc

import (
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/model"
	"strings"
)

func getIrcBuffer(irc Irc) (buffers []*model.Buffer) {
	for _, list := range irc.Data.Lists {
		if strings.Contains(list.DocTitle, "季度") {
			if len(list.Keyword) > 1 {
				buffer := &model.Buffer{}
				buffer.Date = irc.GetDate(list.DocTitle)
				buffer.TargetValue = irc.GetTargetValue(list.Keyword[1])
				buffers = append(buffers, buffer)
			}
		}
	}
	return
}
