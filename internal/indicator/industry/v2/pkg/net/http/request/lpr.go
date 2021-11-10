package request

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
	"io"
	"net/http"
	"time"
)

const lpr = "http://www.chinamoney.com.cn/ags/ms/cm-u-bk-currency/SddsIntrRatePlRatHis?lang=CN&startDate=1997-10-23&endDate=%s&pageNum=1&pageSize=10000"

func today() string {
	return time.Now().Format("2006-01-02")
}

func LPR() (respBytes []byte, err error) {
	resp, err := http.Post(fmt.Sprintf(lpr, today()), "application/json", nil)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	defer resp.Body.Close()
	respBytes, err = io.ReadAll(resp.Body)
	return
}
