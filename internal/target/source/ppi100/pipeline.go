package ppi100

import (
	"github.com/xiaogogonuo/cct-spider/internal/target/model"
	"strings"
)

func ppiPipeline(texts []string, dateIndex, valueIndex, targetColumns int) (responses []model.Response) {
	var response model.Response
	for idx, text := range texts {
		switch idx % targetColumns {
		case dateIndex:
			response.Date = strings.ReplaceAll(text, "-", "")
		case valueIndex:
			response.TargetValue = text
			responses = append(responses, response)
		}
	}
	return
}
