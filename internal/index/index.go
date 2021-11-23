package index

import (
	"encoding/json"
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/index/response"
	"github.com/xiaogogonuo/cct-spider/pkg/db/mysql"
	"github.com/xiaogogonuo/cct-spider/pkg/encrypt/md5"
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
)

var Table = "T_DMAA_BASE_TARGET_VALUE"

type Config struct {
	Name                  string `json:"Name"`
	Case                  string `json:"Case"`
	TargetNameEN          string `json:"TargetNameEN"`
	TargetCode            string `json:"TargetCode"`
	DataSourceCode        string `json:"DataSourceCode"`
	DataSourceName        string `json:"DataSourceName"`
	SourceTargetCode      string `json:"SourceTargetCode"`
	SourceTargetCodeTable string `json:"SourceTargetCodeTable"`
	IsQuantity            string `json:"IsQuantity"`
	UnitType              string `json:"UnitType"`
	UnitName              string `json:"UnitName"`
	PeriodType            string `json:"PeriodType"`
	PeriodName            string `json:"PeriodName"`
	HY                    string `json:"HY"`
	Level                 string `json:"Level"`
	Path1                 string `json:"Path1"`
	Path2                 string `json:"Path2"`
	Path3                 string `json:"Path3"`
	Path4                 string `json:"Path4"`
	Type                  string `json:"Type"`
}

// downloaded 查询已经下载过的日期
func (c Config) downloaded() (row [][]string) {
	switch c.PeriodName {
	case "年":
	case "季":
		sql := fmt.Sprintf("SELECT CONCAT(ACCT_YEAR, ACCT_QUARTOR) FROM %s WHERE TARGET_CODE = '%s'", Table, c.TargetCode)
		row = mysql.Query(sql)
	case "月":
		sql := fmt.Sprintf("SELECT CONCAT(ACCT_YEAR, ACCT_MONTH) FROM %s WHERE TARGET_CODE = '%s'", Table, c.TargetCode)
		row = mysql.Query(sql)
	case "日":
		sql := fmt.Sprintf("SELECT ACCT_DATE FROM %s WHERE TARGET_CODE = '%s'", Table, c.TargetCode)
		row = mysql.Query(sql)
	}
	return
}

// routingDistribution 路由分发到对应的网址
func (c Config) routingDistribution() (rowRespond []response.Respond) {
	switch c.Case {
	case "eastMoneyHG":
		rowRespond = response.RespondMacroIndex(c.SourceTargetCode, c.TargetCode)
	case "eastMoneyHY":
		rowRespond = response.RespondIndustryIndex(c.SourceTargetCode)
	case "eastMoneySHIBOR":
		rowRespond = response.RespondShiBor()
	case "sina":
		rowRespond = response.RespondSina(c.SourceTargetCode)
	case "ifeng":
		rowRespond = response.RespondTBI()
	case "sci":
		pd := response.PostData{
			HY:    c.HY,
			Level: c.Level,
			Path1: c.Path1,
			Path2: c.Path2,
			Path3: c.Path3,
			Path4: c.Path4,
			Type:  c.Type,
		}
		rowRespond = response.RespondSCI(pd)
	}
	return
}

func (c Config) difference(rowDate [][]string, rowRespond []response.Respond) (diffRespond []response.Respond) {
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

func (c Config) construct(rowRespond []response.Respond) (data []Field) {
	for _, respond := range rowRespond {
		f := &Field{}
		f.ValueGUID = md5.MD5(c.TargetCode + respond.Date)
		f.TargetGUID = md5.MD5(c.TargetCode)
		f.TargetCode = c.TargetCode
		f.TargetName = c.Name
		f.TargetNameEN = c.TargetNameEN
		f.DataSourceCode = c.DataSourceCode
		f.DataSourceName = c.DataSourceName
		f.SourceTargetCode = c.SourceTargetCodeTable
		f.IsQuantity = c.IsQuantity
		f.UnitType = c.UnitType
		f.UnitName = c.UnitName
		f.PeriodType = c.PeriodType
		f.PeriodName = c.PeriodName
		f.TargetValue = respond.TargetValue
		acct(respond.Date, f.PeriodType, f)
		data = append(data, *f)
	}
	return
}

func acct(date, period string, f *Field) {
	switch period {
	case PeriodTypeYear:
		f.AcctYear = date
	case PeriodTypeMonth:
		f.AcctYear = date[:4]
		f.AcctMonth = date[4:6]
	case PeriodTypeSeason:
		f.AcctYear = date[:4]
		f.AcctSeason = date[4:6]
	case PeriodTypeDay:
		f.AcctYear = date[:4]
		f.AcctMonth = date[4:6]
		f.AcctDate = date
	}
}

type Configs []Config

func RunIndex() {
	var configs Configs
	if err := json.Unmarshal([]byte(ConfigString), &configs); err != nil {
		logger.Fatal(err.Error())
		return
	}
	for _, config := range configs {
		rowDate := config.downloaded()
		rowRespond := config.routingDistribution()
		diffRespond := config.difference(rowDate, rowRespond)
		if diffRespond == nil || len(diffRespond) == 0 {
			logger.Info(fmt.Sprintf("%s has no data to update", config.Name))
			continue
		}
		data := config.construct(diffRespond)
		logger.Info(config.Name, logger.Field("updating rows: ", len(data)))
		batchDump(data)
	}
}
