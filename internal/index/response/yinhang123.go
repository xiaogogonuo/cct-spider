package response

import (
	"encoding/json"
	"fmt"
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
	"io"
	"net/http"
	"strings"
	"time"
)

// 外汇交易中心

const lpr = "http://www.chinamoney.com.cn/ags/ms/cm-u-bk-currency/SddsIntrRatePlRatHis?lang=CN&" +
	"startDate=1997-10-23&endDate=%s&pageNum=1&pageSize=10000"

type StructLPR struct {
	Records []LPR `json:"records"`
}

type LPR struct {
	DateString string `json:"dateString"`
	LoanRate   string `json:"loanRate"`
}

func today() string {
	return time.Now().Format("2006-01-02")
}

func visitYinHang123() (respBytes []byte, err error) {
	resp, err := http.Post(fmt.Sprintf(lpr, today()), "application/json", nil)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	defer resp.Body.Close()
	respBytes, err = io.ReadAll(resp.Body)
	return
}

func RespondYinHang() (row []Respond) {
	row = make([]Respond, 0)
	b, err := visitYinHang123()
	if err != nil {
		logger.Error(err.Error())
		return
	}
	var s StructLPR
	err = json.Unmarshal(b, &s)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	for _, record := range s.Records {
		var respond Respond
		respond.Date = strings.ReplaceAll(record.DateString, "-", "")
		respond.TargetValue = record.LoanRate
		row = append(row, respond)
	}
	return
}
