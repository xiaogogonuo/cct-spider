package response

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
	"io"
	"net/http"
	"regexp"
	"time"
)

// 汇通财经

const HT = "https://quote.fx678.com/symbol/%s"

func visitHT(url string) (respBytes []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	respBytes, err = io.ReadAll(resp.Body)
	return
}

func RespondHT(sourceTargetCode string)(row []Respond) {
	row = make([]Respond, 0)
	url := fmt.Sprintf(HT, sourceTargetCode)
	resp, err := visitHT(url)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	re := regexp.MustCompile("var yesterdayPrice = \"(.*)\";")
	price := re.FindAllSubmatch(resp, -1)[0][1]
	var respond Respond
	respond.TargetValue = string(price)
	respond.Date = time.Now().Format("20060102")
	row = append(row, respond)
	return
}

/*
日经指数：https://quote.fx678.com/symbol/NIKKI
斯托克600: https://quote.fx678.com/symbol/SXO
道琼斯工业指数：https://quote.fx678.com/symbol/DJIA
标准普尔500指数：https://quote.fx678.com/symbol/SP500
欧元美元 ：https://quote.fx678.com/symbol/EURUSD
美元日元：https://quote.fx678.com/symbol/USDJPY
英镑美元：https://quote.fx678.com/symbol/GBPUSD
原油：https://quote.fx678.com/symbol/OILC
黄金：https://quote.fx678.com/symbol/GLNC
白银：https://quote.fx678.com/symbol/SLNC
铜：https://quote.fx678.com/symbol/LMCI
10年期美国国债：https://quote.fx678.com/symbol/USG10Y
10年期德国国债：https://quote.fx678.com/symbol/GDBR10
10年期英国国债：https://quote.fx678.com/symbol/GUKG10
10年期日本国债：https://quote.fx678.com/symbol/GJGB10
 */