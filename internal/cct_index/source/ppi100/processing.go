package ppi100

import (
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/model"
	"strings"
)

func getPPI100Buffer(texts []string, dateIndex, valueIndex, targetColumns int) (buffers []*model.Buffer) {
	buffer := &model.Buffer{}
	for idx, text := range texts {
		switch idx % targetColumns {
		case dateIndex:
			buffer.Date = strings.ReplaceAll(text, "-", "")
		case valueIndex:
			buffer.TargetValue = text
			buffers = append(buffers, buffer)
			buffer = &model.Buffer{}
		}
	}
	return
}
