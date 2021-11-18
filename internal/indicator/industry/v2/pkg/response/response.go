package response

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/xiaogogonuo/cct-spider/internal/indicator/industry/v2/pkg/net/http/request"
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
	"regexp"
	"strings"
	"time"
	"unsafe"
)

type Respond struct {
	Date        string
	TargetValue string
}

func GetIavTB(sourceTargetCode string) (row []Respond) {
	return
}

func GetEastMoney(sourceTargetCode string) (row []Respond) {
	row = make([]Respond, 0)
	for i := 1; ; i++ {
		b, err := request.EastMoneyHY(sourceTargetCode, i)
		if err != nil {
			logger.Error(err.Error())
			continue
		}
		var obj StructEastMoney
		if err = json.Unmarshal(b, &obj); err != nil {
			logger.Error(err.Error())
			continue
		}
		if !obj.Success {
			break
		}
		for _, data := range obj.Result.Data {
			var respond Respond
			respond.Date = strings.ReplaceAll(data.ReportDate, "-", "")[:8]
			respond.TargetValue = fmt.Sprintf("%.0f", data.IndicatorValue)
			row = append(row, respond)
		}
	}
	return
}

func GetSCI(pd request.PostData) (row []Respond) {
	row = make([]Respond, 0)
	b, err := request.SCI(pd)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	var frontEnd StructSCI
	if err = json.Unmarshal(b, &frontEnd); err != nil {
		logger.Error(err.Error())
		return
	}
	for _, front := range frontEnd.List {
		var respond Respond
		respond.TargetValue = fmt.Sprintf("%.2f", front.MDataValue)
		respond.Date = strings.ReplaceAll(front.DataDate, "/", "")
		row = append(row, respond)
	}
	return
}

func GetSina(sourceTargetCode string) (row []Respond) {
	row = make([]Respond, 0)
	b, err := request.Sina(sourceTargetCode)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	str := (*string)(unsafe.Pointer(&b))
	reg := regexp.MustCompile(`".*"`)
	all := reg.FindString(*str)
	all = strings.ReplaceAll(all, `"`, "")
	allArray := strings.Split(all, "|")
	for _, v := range allArray {
		var respond Respond
		vs := strings.Split(v, ",")
		respond.TargetValue = vs[3]
		respond.Date = strings.ReplaceAll(vs[0], "-", "")
		row = append(row, respond)
	}
	return
}

func GetShiBor() (row []Respond) {
	row = make([]Respond, 0)
	for i := 1; ; i++ {
		b, err := request.ShiBor(i)
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
		var respond Respond
		for idx, value := range tableText {
			if idx%3 == 0 {
				respond.Date = strings.ReplaceAll(value, "-", "")
			} else if idx%3 == 1 {
				respond.TargetValue = value
			} else {
				row = append(row, respond)
				respond = Respond{}
			}
		}
	}
	return
}

// GetTBI 国债指数
// 起始日期：2010-01-04
func GetTBI() (row []Respond) {
	row = make([]Respond, 0)
	for i := 2014; i <= time.Now().Year(); i++ {
		b, err := request.TBI(i)
		if err != nil {
			logger.Error(err.Error())
			continue
		}
		dom, _ := goquery.NewDocumentFromReader(strings.NewReader(string(b)))
		var tableText []string
		dom.Find("div[class='tab01'] td").Each(func(i int, selection *goquery.Selection) {
			text := selection.Text()
			tableText = append(tableText, text)
		})
		if len(tableText) < 1 {
			break
		}
		var respond Respond
		for idx, value := range tableText {
			if idx%9 == 0 {
				respond.Date = strings.ReplaceAll(value, "-", "")
			} else if idx%9 == 4 {
				respond.TargetValue = value
			} else if idx%9 == 8 {
				row = append(row, respond)
				respond = Respond{}
			}
		}
	}
	return
}

// GetLPR 贷款基准利率
func GetLPR() (row []Respond) {
	row = make([]Respond, 0)
	b, err := request.LPR()
	if err != nil {
		logger.Error(err.Error())
		return
	}
	var lpr StructLPR
	err = json.Unmarshal(b, &lpr)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	for _, record := range lpr.Records {
		var respond Respond
		respond.Date = strings.ReplaceAll(record.DateString, "-", "")
		respond.TargetValue = record.LoanRate
		row = append(row, respond)
	}
	return
}
