package index

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/index/response"
	"github.com/xiaogogonuo/cct-spider/pkg/db/mysql"
	"github.com/xiaogogonuo/cct-spider/pkg/encrypt/md5"
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
	"io"
	"math"
	"net/http"
	"strings"
	"sync"
	"time"
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
	RunTime               string `json:"RunTime"`
}

// downloaded 查询已经下载过的日期
func (c Config) downloaded() (row [][]string) {
	switch c.PeriodName {
	case "年":
		sql := fmt.Sprintf("SELECT ACCT_YEAR FROM %s WHERE TARGET_CODE = '%s'", Table, c.TargetCode)
		row = mysql.Query(sql)
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

// regionYearDownloaded 查询地区生产总值已经下载过的年度日期
func (c Config) regionYearDownloaded() (row [][]string) {
	sql := fmt.Sprintf("SELECT CONCAT(ACCT_YEAR, REGION_CODE) FROM %s WHERE TARGET_CODE = '%s'", Table, c.TargetCode)
	row = mysql.Query(sql)
	return
}

// regionMonthDownloaded 查询地区生产总值已经下载过的月度日期
func (c Config) regionMonthDownloaded() (row [][]string) {
	sql := fmt.Sprintf("SELECT CONCAT(ACCT_YEAR, ACCT_MONTH, REGION_CODE) FROM %s WHERE TARGET_CODE = '%s'", Table, c.TargetCode)
	row = mysql.Query(sql)
	return
}

// routingDistribution 路由分发到对应的网址
func (c Config) routingDistribution() (rowRespond []response.Respond) {
	switch c.Case {
	//case "sinaHSI":
	//	rowRespond = response.RespondHSI()
	//case "sinaUSDCNY":
	//	rowRespond = response.RespondUSDCNY()
	//case "sinaHDKCNY":
	//	rowRespond = response.RespondHKDCNY()
	//case "eastMoneyIM":
	//	rowRespond = response.RespondIM()
	//case "eastMoneyGCHN10":
	//	rowRespond = response.RespondGCHN10()
	//case "eastMoneyHG":
	//	rowRespond = response.RespondMacroIndex(c.SourceTargetCode, c.TargetCode)
	//case "eastMoneyHY":
	//	rowRespond = response.RespondIndustryIndex(c.SourceTargetCode)
	//case "eastMoneySHIBOR":
	//	rowRespond = response.RespondShiBor(1) // 第一次-1，后面设置为1
	//case "sina":
	//	rowRespond = response.RespondSina(c.SourceTargetCode)
	//case "sinaRegionGDP":
	//	rowRespond = response.RespondSinaRegionGDP()
	//case "sinaRegionCPI":
	//	rowRespond = response.RespondSinaRegionCPI()
	//case "sinaCPI":
	//	rowRespond = response.RespondSinaCPI()
	//case "ifeng":
	//	rowRespond = response.RespondTBI()
	//case "fxGJZS", "fxGJZQ", "fxWH", "fxWHMP", "fxIPE", "fxCOMEX", "fxLME":
	//	rowRespond = response.RespondHT(c.Name, c.Case)
	case "fxLibor": // 美元Libor隔夜
	  rowRespond = response.RespondUSLibor(c.TargetCode)
	//case "cni":
	//	rowRespond = response.RespondCNYX()
	//case "sci":
	//	pd := response.PostData{
	//		HY:    c.HY,
	//		Level: c.Level,
	//		Path1: c.Path1,
	//		Path2: c.Path2,
	//		Path3: c.Path3,
	//		Path4: c.Path4,
	//		Type:  c.Type,
	//	}
	//	rowRespond = response.RespondSCI(pd)
	}
	return
}

func (c Config) difference(rowDate [][]string, rowRespond []response.Respond) (diffRespond []response.Respond) {
	ch := make(map[string]struct{})
	for _, v := range rowDate {
		ch[v[0]] = struct{}{}
	}
	for _, row := range rowRespond {
		if _, ok := ch[row.Date+row.RegionCode]; !ok {
			diffRespond = append(diffRespond, row)
		}
	}
	return
}

func (c Config) construct(rowRespond []response.Respond, all bool) (data []Field) {
	for _, respond := range rowRespond {
		f := &Field{}
		f.ValueGUID = md5.MD5(c.TargetCode + respond.Date + respond.RegionCode)
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
		f.RegionCode = respond.RegionCode
		f.RegionName = respond.RegionName
		f.PeriodType = c.PeriodType
		f.PeriodName = c.PeriodName
		switch all {
		case true:
			f.TargetValue = respond.TargetValue
		case false:
			f.TargetValue = strings.Split(respond.TargetValue, ",")[0]
		}
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

var webService = "http://106.37.165.121/inf/dm/aa/baseTargetValue/saveRequest"

func RunIndex() {
	var wg sync.WaitGroup
	var configs Configs
	if err := json.Unmarshal([]byte(ConfigString), &configs); err != nil {
		logger.Fatal(err.Error())
		return
	}
	for _, config := range configs {
		curTime := time.Now()
		runTimes := strings.Split(config.RunTime, "~")
		runTimeL, _ := time.Parse("15:04", runTimes[0])
		runTimeR, _ := time.Parse("15:04", runTimes[1])
		// 24小时刷新数据的用另一套逻辑，此处跳过
		if runTimeR.Hour()-runTimeL.Hour() == 0 {
			continue
		}
		if curTime.Hour() < runTimeL.Hour() || curTime.Hour() > runTimeR.Hour() {
			continue
		}
		if curTime.Hour() == runTimeL.Hour() && curTime.Minute() < runTimeL.Minute() {
			continue
		}
		if curTime.Hour() == runTimeR.Hour() && curTime.Minute() > runTimeR.Minute() {
			continue
		}
		var rowDate [][]string
		switch config.Case {
		case "sinaRegionGDP":
			rowDate = config.regionYearDownloaded()
		case "sinaRegionCPI":
			rowDate = config.regionMonthDownloaded()
		default:
			rowDate = config.downloaded()
		}
		rowRespond := config.routingDistribution()
		diffRespond := config.difference(rowDate, rowRespond)
		if diffRespond == nil || len(diffRespond) == 0 {
			continue
		}
		data := config.construct(diffRespond, false) // 推送到本地数据库
		java := config.construct(diffRespond, true)  // 推送给java服务器
		logger.Info(config.Name, logger.Field("updating rows: ", len(data)))
		epoch := int(math.Ceil(float64(len(data)) / float64(batchSize)))
		wg.Add(epoch + 1)
		go batchDump(data, &wg)
		for i := 0; i < epoch; i++ {
			if batchSize*(i+1) < len(data) {
				batchData := java[i*batchSize : (i+1)*batchSize]
				go send(webService, batchData, &wg)
			} else {
				batchData := java[i*batchSize:]
				go send(webService, batchData, &wg)
			}
		}
	}
	wg.Wait()
}

func send(api string, data []Field, wg *sync.WaitGroup) {
	defer wg.Done()
	postData := map[string][]Field{"data": data}
	m, _ := json.Marshal(postData)
	req, err := http.NewRequest(http.MethodPost, api, bytes.NewReader(m))
	if err != nil {
		logger.Error(err.Error())
		return
	}
	req.Header.Add("Content-Type", "application/json")
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	defer resp.Body.Close()
	b, _ := io.ReadAll(resp.Body)
	logger.Info(string(b))
}
