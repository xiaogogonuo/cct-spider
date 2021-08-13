package response

import (
	"encoding/json"
	"github.com/xiaogogonuo/cct-spider/internal/indicator/marco/v2/pkg/net/http/request"
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
	"strings"
)

type Respond struct {
	HasData       bool
	Date          string
	Data          string
	ProvinceCode  string
	IndicatorCode string
}

func replaceSeason(date string) string {
	date = strings.ReplaceAll(date, "A", "Q1")
	date = strings.ReplaceAll(date, "B", "Q2")
	date = strings.ReplaceAll(date, "C", "Q3")
	date = strings.ReplaceAll(date, "D", "Q4")
	return date
}

func (r *Respond) coder(code string) {
	for _, c := range strings.Split(code, "_") {
		cc := strings.Split(c, ".")
		switch cc[0] {
		case "zb":
			r.IndicatorCode = cc[1]
		case "reg":
			r.ProvinceCode = cc[1]
		case "sj":
			r.Date = replaceSeason(cc[1])
		}
	}
}

func Crawl(cn, zb string, param request.Param) (row []*Respond) {
	resBody, err := request.Request(cn, zb, param)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	var frontResponse Response
	if err = json.Unmarshal(resBody, &frontResponse); err != nil {
		logger.Error(err.Error())
		return
	}
	nodes := frontResponse.ReturnData.DataNodes
	for _, node := range nodes {
		respond := &Respond{}
		respond.coder(node.Code)
	    respond.Data = node.Data.StrData
	    respond.HasData = node.Data.HasData
	    if !respond.HasData {
	    	continue
		}
	    row = append(row, respond)
	}
	return
}