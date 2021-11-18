package indicator

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/indicator/industry/v2/code"
	"github.com/xiaogogonuo/cct-spider/internal/indicator/industry/v2/pkg/net/http/request"
	"github.com/xiaogogonuo/cct-spider/internal/indicator/industry/v2/pkg/response"
	"github.com/xiaogogonuo/cct-spider/pkg/db/mysql"
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
	"sync"
)

type model struct {
	indicatorName string
	indicatorInfo map[string]string
}

func (m model) downloaded() (row [][]string) {
	row = mysql.Query(fmt.Sprintf(IndustrySQL, table, m.indicatorInfo["TargetCode"]))
	return
}

func (m model) difference(rowDate [][]string, rowRespond []response.Respond) (diffRespond []response.Respond) {
	ch := make(map[string]struct{})
	for _, v := range rowDate {
		ch[v[0]] = struct{}{}
	}
	for _, row := range rowRespond {
		if _, ok := ch[row.Date]; !ok {
			diffRespond = append(diffRespond, row)
		}
	}
	return
}

func (m model) entrypoint(wg *sync.WaitGroup) {
	defer wg.Done()
	rowDate := m.downloaded()
	var rowRespond []response.Respond
	switch m.indicatorInfo["Flag"] {
	case "eastmoney":
		rowRespond = response.GetEastMoney(m.indicatorInfo["SourceTargetCode"])
	case "IavTB":
		rowRespond = response.GetIavTB(m.indicatorInfo["SourceTargetCode"])
	case "sci":
		pd := request.PostData{
			HY:    m.indicatorInfo["HY"],
			Level: m.indicatorInfo["2"],
			Path1: m.indicatorInfo["Path1"],
			Path2: m.indicatorInfo["Path2"],
			Path3: m.indicatorInfo["Path3"],
			Path4: m.indicatorInfo["Path4"],
			Type:  m.indicatorInfo["Type"],
		}
		rowRespond = response.GetSCI(pd)
	case "sina":
		rowRespond = response.GetSina(m.indicatorInfo["SourceTargetCode"])
	case "shi":
		rowRespond = response.GetShiBor()
	case "tbi":
		rowRespond = response.GetTBI()
	case "lpr":
		rowRespond = response.GetLPR()
	}
	diffRespond := m.difference(rowDate, rowRespond)
	if diffRespond == nil || len(diffRespond) == 0 {
		logger.Info(fmt.Sprintf("%s has no data to update", m.indicatorName))
		return
	}
	c := constructor{
		indicatorName: m.indicatorName,
		indicatorInfo: m.indicatorInfo,
		respond:       diffRespond,
	}
	data := c.construct()
	logger.Info(m.indicatorName, logger.Field("updating rows: ", len(data)))
	batchDump(data)
}

func Start() {
	var wg sync.WaitGroup
	for indicatorName, indicatorInfo := range code.Indicator {
		wg.Add(1)
		go model{indicatorName, indicatorInfo}.entrypoint(&wg)
	}
	wg.Wait()
}
