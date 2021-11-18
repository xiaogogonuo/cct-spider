package response

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
	"io"
	"net/http"
	"strings"
	"time"
)

// 凤凰网财经

const tbiURL = "http://app.finance.ifeng.com/hq/stock_daily.php?code=sh000012&begin_day=%d-01-01&end_day=%d-12-31"

// visitTBI 请求凤凰网财经的国债指数接口
func visitTBI(year int) (respBytes []byte, err error) {
	url := fmt.Sprintf(tbiURL, year, year)
	resp, err := http.Get(url)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	defer resp.Body.Close()
	respBytes, err = io.ReadAll(resp.Body)
	return
}

// RespondTBI 返回凤凰网财经的国债指数数据
func RespondTBI() (row []Respond) {
	row = make([]Respond, 0)
	for i := 2014; i <= time.Now().Year(); i++ {
		b, err := visitTBI(i)
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
