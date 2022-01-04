package response

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
	"io"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"
)

// 东方财富

/* 中国10年期国债到期收益率
页面展示接口：https://data.eastmoney.com/cjsj/zmgzsyl.html
数据抓包接口：https://datacenter-web.eastmoney.com/api/data/get?type=RPTA_WEB_TREASURYYIELD&sty=ALL&st=SOLAR_DATE&sr=-1&p=1&ps=500
*/

const GCHN10URL = "https://datacenter-web.eastmoney.com/api/data/get?type=RPTA_WEB_TREASURYYIELD&sty=ALL&st=SOLAR_DATE&sr=-1&p=%d&ps=500"

type GCHN10 struct {
	Success bool `json:"success"`
	Result  struct {
		Data []struct {
			SOLAR_DATE  string  `json:"SOLAR_DATE"`
			EMM00166466 float64 `json:"EMM00166466"`
		} `json:"data"`
	} `json:"result"`
}

type Record []RecordSt

type RecordSt struct {
	Date  string
	Value float64
}

func (r Record) Len() int {
	return len(r)
}

func (r Record) Less(i, j int) bool {
	return r[i].Date > r[j].Date // 降序
}

func (r Record) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func visitGCHN10URL(url string) (body []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	defer resp.Body.Close()
	body, err = io.ReadAll(resp.Body)
	return
}

func RespondGCHN10() (row []Respond) {
	var record Record
	for i := 1; ; i++ {
		body, err := visitGCHN10URL(fmt.Sprintf(GCHN10URL, i))
		if err != nil {
			continue
		}
		var s GCHN10
		err = json.Unmarshal(body, &s)
		if err != nil {
			logger.Error(err.Error())
			continue
		}
		if !s.Success {
			break
		}
		for _, data := range s.Result.Data {
			date := strings.Split(data.SOLAR_DATE, "/")
			var (
				year  string
				month string
				day   string
			)
			year = date[0]
			if len(date[1]) == 1 {
				month = "0" + date[1]
			} else {
				month = date[1]
			}
			if len(date[2]) == 1 {
				day = "0" + date[2]
			} else {
				day = date[2]
			}
			value := data.EMM00166466
			record = append(record, RecordSt{year + month + day, value})
		}
	}
	sort.Sort(record)
	for i := 0; i < len(record)-1; i++ {
		var respond Respond
		respond.Date = record[i].Date
		updateTime := respond.Date[:4] + "-" + respond.Date[4:6] + "-" + respond.Date[6:8] + " " +
			strconv.FormatInt(int64(time.Now().Hour()), 10) + ":" +
			strconv.FormatInt(int64(time.Now().Minute()), 10) + ":" +
			strconv.FormatInt(int64(time.Now().Second()), 10)
		todayValue := fmt.Sprintf("%.4f", record[i].Value)
		if todayValue == "0.0000" {
			todayValue = ""
		}
		yesterdayValue := fmt.Sprintf("%.4f", record[i+1].Value)
		if yesterdayValue == "0.0000" {
			yesterdayValue = ""
		}
		var upDown, upDownPercent string
		if todayValue == "" || yesterdayValue == "" {
			upDown = ""
			upDownPercent = ""
		} else {
			upDown = fmt.Sprintf("%.4f", record[i].Value-record[i+1].Value)
			upDownPercent = fmt.Sprintf("%.4f", (record[i].Value-record[i+1].Value)/record[i+1].Value)
		}
		respond.TargetValue = strings.Join([]string{
			todayValue,     // 现价
			upDown,         // 涨跌
			upDownPercent,  // 涨跌幅
			"",             // 最高
			"",             // 最低
			yesterdayValue, // 昨收
			updateTime,     // 更新时间
		}, ",")
		row = append(row, respond)
	}
	return
}

/* SHIBOR隔夜
页面展示接口：https://data.eastmoney.com/shibor/shibor.aspx?m=sh&t=99&d=99221&cu=cny&type=009016
数据抓包接口：https://data.eastmoney.com/shibor/shibor.aspx?m=sh&t=99&d=99221&cu=cny&type=009016
*/

// Shibor同业间拆借（隔夜）利率
const shiBorURL = "https://data.eastmoney.com/shibor/shibor.aspx?m=sh&t=99&d=99221&cu=cny&type=009016&p=%d"

// visitShiBor 请求东方财富接口
// 适用指标：上海银行间同业拆放利率
func visitShiBor(page int) (respBytes []byte, err error) {
	url := fmt.Sprintf(shiBorURL, page)
	resp, err := http.Get(url)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	defer resp.Body.Close()
	respBytes, err = io.ReadAll(resp.Body)
	return
}

// RespondShiBor 返回东方财富的数据
// 适用指标：上海银行间同业拆放利率
func RespondShiBor(stopPage int) (row []Respond) {
	row = make([]Respond, 0)
	var record Record
	for i := 1; ; i++ {
		if i > stopPage + 1 && stopPage != -1 {
			break
		}
		b, err := visitShiBor(i)
		if err != nil {
			logger.Error(err.Error())
			continue
		}
		dom, _ := goquery.NewDocumentFromReader(strings.NewReader(string(b)))
		var tableText []string
		dom.Find("table[id='tb'] td").Each(func(i int, selection *goquery.Selection) {
			text := selection.Text()
			tableText = append(tableText, text)
		})
		if len(tableText) < 1 {
			break
		}
		var st RecordSt
		for idx, value := range tableText {
			if idx%3 == 0 {
				st.Date = strings.ReplaceAll(value, "-", "")
			} else if idx%3 == 1 {
				st.Value, _ = strconv.ParseFloat(value, 64)
			} else {
				record = append(record, st)
			}
		}
	}
	sort.Sort(record)
	for i := 0; i < len(record)-1; i++ {
		var respond Respond
		respond.Date = record[i].Date
		updateTime := respond.Date[:4] + "-" + respond.Date[4:6] + "-" + respond.Date[6:8] + " " +
			strconv.FormatInt(int64(time.Now().Hour()), 10) + ":" +
			strconv.FormatInt(int64(time.Now().Minute()), 10) + ":" +
			strconv.FormatInt(int64(time.Now().Second()), 10)
		todayValue := fmt.Sprintf("%.4f", record[i].Value)
		if todayValue == "0.0000" {
			todayValue = ""
		}
		yesterdayValue := fmt.Sprintf("%.4f", record[i+1].Value)
		if yesterdayValue == "0.0000" {
			yesterdayValue = ""
		}
		var upDown, upDownPercent string
		if todayValue == "" || yesterdayValue == "" {
			upDown = ""
			upDownPercent = ""
		} else {
			upDown = fmt.Sprintf("%.4f", record[i].Value-record[i+1].Value)
			upDownPercent = fmt.Sprintf("%.4f", (record[i].Value-record[i+1].Value)/record[i+1].Value)
		}
		respond.TargetValue = strings.Join([]string{
			todayValue,     // 现价
			upDown,         // 涨跌
			upDownPercent,  // 涨跌幅
			"",             // 最高
			"",             // 最低
			yesterdayValue, // 昨收
			updateTime,     // 更新时间
		}, ",")
		row = append(row, respond)
	}
	return
}

//func RespondShiBor() (row []Respond) {
//	row = make([]Respond, 0)
//	for i := 1; ; i++ {
//		b, err := visitShiBor(i)
//		if err != nil {
//			logger.Error(err.Error())
//			continue
//		}
//		dom, _ := goquery.NewDocumentFromReader(strings.NewReader(string(b)))
//		var tableText []string
//		dom.Find("table[id='tb'] td").Each(func(i int, selection *goquery.Selection) {
//			text := selection.Text()
//			tableText = append(tableText, text)
//		})
//		if len(tableText) < 1 {
//			break
//		}
//		var respond Respond
//		for idx, value := range tableText {
//			if idx%3 == 0 {
//				respond.Date = strings.ReplaceAll(value, "-", "")
//			} else if idx%3 == 1 {
//				respond.TargetValue = value
//			} else {
//				row = append(row, respond)
//				respond = Respond{}
//			}
//		}
//	}
//	return
//}

/* 铁矿石主力合约
页面展示接口：http://quote.eastmoney.com/qihuo/IM.html
数据抓包接口：https://futsseapi.eastmoney.com/static/114_im_qt
*/

var ImApi = "https://futsseapi.eastmoney.com/static/114_im_qt"

type IM struct {
	QT Detail `json:"qt"`
}

type Detail struct {
	P     float64 `json:"p"`     // 最新价
	ZDE   float64 `json:"zde"`   // 涨跌额
	ZDF   float64 `json:"zdf"`   // 涨跌幅
	H     float64 `json:"h"`     // 最高
	L     float64 `json:"l"`     // 最低
	QRSPJ float64 `json:"qrspj"` // 昨收
}

func visitIM(url string) (respBytes []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	defer resp.Body.Close()
	respBytes, err = io.ReadAll(resp.Body)
	return
}

func RespondIM() (row []Respond) {
	body, err := visitIM(ImApi)
	if err != nil {
		return
	}
	var im IM
	if err := json.Unmarshal(body, &im); err != nil {
		logger.Error(err.Error())
		return
	}
	var respond Respond
	respond.Date = time.Now().Format("20060102")
	value := fmt.Sprintf("%.2f,%.2f,%.2f%s,%.2f,%.2f,%.2f,%s", im.QT.P, im.QT.ZDE, im.QT.ZDF, "%", im.QT.H, im.QT.L, im.QT.QRSPJ,
		time.Now().Format("2006-01-02 15:04:05"))
	respond.TargetValue = value
	row = append(row, respond)
	return
}

// StructIndustryIndex 东方财富行业指标接口返回的数据结构
type StructIndustryIndex struct {
	Result  R    `json:"result"`
	Success bool `json:"success"`
}

type R struct {
	Data []D `json:"data"`
}

type D struct {
	ReportDate     string  `json:"REPORT_DATE"`
	IndicatorValue float64 `json:"INDICATOR_VALUE"`
}

// 东方财富行业指标的URL
// https://datacenter-web.eastmoney.com/api/data/v1/get?sortColumns=REPORT_DATE&sortTypes=-1&pageSize=500&pageNumber=17&reportName=RPT_INDUSTRY_INDEX&columns=REPORT_DATE%2CINDICATOR_VALUE&filter=(INDICATOR_ID%3D%22EMI00107664%22)
const (
	u1HY = "https://datacenter-web.eastmoney.com/api/data/v1/get?sortColumns=REPORT_DATE&sortTypes=-1&pageSize=500&pageNumber="
	u2HY = "&reportName=RPT_INDUSTRY_INDEX&columns=REPORT_DATE%2CINDICATOR_VALUE&filter=(INDICATOR_ID%3D%22"
	u3HY = "%22)"
)

// visitIndustryIndex 请求东方财富行业数据接口
// 适用指标：中国大宗商品指数、物流业景气指数、波罗的海干散货指数
func visitIndustryIndex(sourceTargetCode string, page int) (b []byte, err error) {
	url := u1HY + strconv.Itoa(page) + u2HY + sourceTargetCode + u3HY
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	b, err = io.ReadAll(resp.Body)
	return
}

// RespondIndustryIndex 返回东方财富行业指标数据
// 适用指标：中国大宗商品指数、物流业景气指数、波罗的海干散货指数
func RespondIndustryIndex(sourceTargetCode string) (row []Respond) {
	row = make([]Respond, 0)
	for i := 1; ; i++ {
		b, err := visitIndustryIndex(sourceTargetCode, i)
		if err != nil {
			logger.Error(err.Error())
			continue
		}
		var s StructIndustryIndex
		if err = json.Unmarshal(b, &s); err != nil {
			logger.Error(err.Error())
			continue
		}
		if !s.Success {
			break
		}
		for _, data := range s.Result.Data {
			var respond Respond
			respond.Date = strings.ReplaceAll(data.ReportDate, "-", "")[:8]
			respond.TargetValue = fmt.Sprintf("%.0f", data.IndicatorValue)
			row = append(row, respond)
		}
	}
	return
}

// 东方财富宏观指标的URL
const macroURL = "http://datainterface.eastmoney.com/EM_DataCenter/JS.aspx?type=GJZB&sty=ZGZB&p=2&ps=1000000&mkt=%s"

// visitMacroIndex 请求东方财富宏观数据接口
// 适用指标：工业增加值增长、社会消费品零售总额、货币供应量、居民消费价格指数(CPI)、国内生产总值(GDP)
// 采购经理人指数(PMI)、工业品出厂价格指数(PPI)
func visitMacroIndex(sourceTargetCode string) (b []byte, err error) {
	url := fmt.Sprintf(macroURL, sourceTargetCode)
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	b, err = io.ReadAll(resp.Body)
	return
}

// RespondMacroIndex 返回东方财富宏观指标数据
// 适用指标：工业增加值增长、社会消费品零售总额、货币供应量、居民消费价格指数(CPI)、国内生产总值(GDP)
// 采购经理人指数(PMI)、工业品出厂价格指数(PPI)、贷款基准利率LPR
func RespondMacroIndex(sourceTargetCode, targetCode string) (row []Respond) {
	row = make([]Respond, 0)
	b, err := visitMacroIndex(sourceTargetCode)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	stringB := string(b)
	stringB = stringB[1 : len(stringB)-1]
	switch targetCode {
	case "HY00007": // 贷款基准利率
		row = lprIndustryIndexApiSimilarMarcoIndex(stringB, 9, 5)
	case "HG00016": // 工业增加值同比增长
		row = marcoPipe(stringB, 0, 1)
	case "HG00017": // 工业增加值累计增长
		row = marcoPipe(stringB, 0, 2)
	case "HG00027": // 社会消费品零售总额当期值
		row = marcoPipe(stringB, 0, 1)
	case "HG00028": // 社会消费品零售总额累计值
		row = marcoPipe(stringB, 0, 4)
	case "HG00029": // 社会消费品零售总额同比增长
		row = marcoPipe(stringB, 0, 2)
	case "HG00030": // 社会消费品零售总额累计增长
		row = marcoPipe(stringB, 0, 5)
	case "HG00006": // 货币和准货币(M2)供应量期末值
		row = marcoPipe(stringB, 0, 1)
	case "HG00007": // 货币和准货币(M2)供应量同比增长
		row = marcoPipe(stringB, 0, 2)
	case "HG00004": // 居民消费价格指数当月，对应国家统计局的居民消费价格指数(上年同月=100)
		row = marcoPipe(stringB, 0, 1)
	case "HG00020": // 制造业采购经理指数
		row = marcoPipe(stringB, 0, 1)
	case "HG00023": // 工业品出厂价格指数当月，对应国家统计局的工业生产者出厂价格指数(上年同月=100)
		row = marcoPipe(stringB, 0, 1)
	case "HG00001": // 国内生产总值同比增长(季度)
		row = marcoPipeMonth2Season(stringB, 0, 2)
	}
	return
}

// 贷款基准利率在表设计时，当作行业指标，但是它的接口跟宏观指标一致，同时它是以天为单位，而宏观指标是以月为单位
// 故提取前8位，其余部分与宏观指标逻辑一致
func lprIndustryIndexApiSimilarMarcoIndex(s string, dateIndex, valueIndex int) (row []Respond) {
	var data []string
	if err := json.Unmarshal([]byte(s), &data); err != nil {
		logger.Error(err.Error())
		return
	}
	for _, d := range data {
		var respond Respond
		ds := strings.Split(d, ",")
		respond.Date = strings.ReplaceAll(ds[dateIndex], "-", "")[:8] // 提取前8位
		respond.TargetValue = ds[valueIndex]
		row = append(row, respond)
	}
	return
}

func marcoPipe(s string, dateIndex, valueIndex int) (row []Respond) {
	var data []string
	if err := json.Unmarshal([]byte(s), &data); err != nil {
		logger.Error(err.Error())
		return
	}
	for _, d := range data {
		var respond Respond
		ds := strings.Split(d, ",")
		respond.Date = strings.ReplaceAll(ds[dateIndex], "-", "")[:6] // 提取前6位
		respond.TargetValue = ds[valueIndex]
		row = append(row, respond)
	}
	return
}

func marcoPipeMonth2Season(s string, dateIndex, valueIndex int) (row []Respond) {
	var data []string
	if err := json.Unmarshal([]byte(s), &data); err != nil {
		logger.Error(err.Error())
		return
	}
	for _, d := range data {
		var respond Respond
		ds := strings.Split(d, ",")
		date := strings.ReplaceAll(ds[dateIndex], "-", "")
		switch date[4:6] {
		case "01", "02", "03":
			respond.Date = date[:4] + "Q1"
		case "04", "05", "06":
			respond.Date = date[:4] + "Q2"
		case "07", "08", "09":
			respond.Date = date[:4] + "Q3"
		case "10", "11", "12":
			respond.Date = date[:4] + "Q4"
		}
		respond.TargetValue = ds[valueIndex]
		row = append(row, respond)
	}
	return
}
