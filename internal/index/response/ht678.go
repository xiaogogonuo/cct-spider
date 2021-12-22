package response

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
	"io"
	"net/http"
	"strings"
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

//func RespondHT(sourceTargetCode string)(row []Respond) {
//	row = make([]Respond, 0)
//	url := fmt.Sprintf(HT, sourceTargetCode)
//	resp, err := visitHT(url)
//	if err != nil {
//		logger.Error(err.Error())
//		return
//	}
//	re := regexp.MustCompile("var yesterdayPrice = \"(.*)\";")
//	price := re.FindAllSubmatch(resp, -1)[0][1]
//	var respond Respond
//	respond.TargetValue = string(price)
//	respond.Date = time.Now().Format("20060102")
//	row = append(row, respond)
//	return
//}

func respondHT(condition string) string {
	switch condition {
	case "fxGJZS": // 全球指数
		return "https://quote.fx678.com/exchange/GJZS"
	case "fxGJZQ": // 国际债券
		return "https://quote.fx678.com/exchange/GJZQ"
	case "fxWH": // 直盘
		return "https://quote.fx678.com/exchange/WH"
	case "fxIPE": // PIE原油
		return "https://quote.fx678.com/exchange/IPE"
	case "fxCOMEX": // 纽约金属
		return "https://quote.fx678.com/exchange/COMEX"
	case "fxLME": // 伦敦金属
		return "https://quote.fx678.com/exchange/LME"
	}
	return ""
}

func RespondHT(name, condition string) (row []Respond) {
	resp, err := visitHT(respondHT(condition))
	if err != nil {
		logger.Error(err.Error())
		return
	}
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(string(resp)))
	if err != nil {
		logger.Error(err.Error())
		return
	}
	var tmp []string
	dom.Find("table tr[id] td").Each(func(i int, selection *goquery.Selection) {
		tmp = append(tmp, strings.Trim(selection.Text(), "\n\t "))
	})
	var data string
	for i := 0; i < len(tmp)/8; i++ {
		if tmp[i*8] == name {
			data = strings.Join(tmp[i*8+1:(i+1)*8-1], ",")
			break
		}
	}
	var respond Respond
	respond.TargetValue = data
	respond.Date = time.Now().Format("20060102")
	row = append(row, respond)
	return
}

/*
下午2点半爬
日经指数：https://quote.fx678.com/symbol/NIKKI
10年期日本国债：https://quote.fx678.com/symbol/GJGB10

下午2点半爬
道琼斯工业指数：https://quote.fx678.com/symbol/DJIA
纳斯达克指数：https://quote.fx678.com/symbol/NASDAQ

下午2点半爬
斯托克600：https://quote.fx678.com/symbol/SXO
标准普尔500指数：https://quote.fx678.com/symbol/SP500
英国FTSE100：https://quote.fx678.com/symbol/FTSE
10年期德国国债：https://quote.fx678.com/symbol/GDBR10
10年期英国国债：https://quote.fx678.com/symbol/GUKG10

下午5点爬
欧元美元 ：https://quote.fx678.com/symbol/EURUSD
美元日元：https://quote.fx678.com/symbol/USDJPY
英镑美元：https://quote.fx678.com/symbol/GBPUSD
原油：https://quote.fx678.com/symbol/OILC
黄金：https://quote.fx678.com/symbol/GLNC
白银：https://quote.fx678.com/symbol/SLNC
铜：https://quote.fx678.com/symbol/LMCI
恒生指数：https://quote.fx678.com/symbol/HSI
美元指数：https://quote.fx678.com/symbol/USD
10年期美国国债：https://quote.fx678.com/symbol/USG10Y  收盘时间待定

1、中国沪深市场交易时间为：周一至五的9:30-11:30，下午13:00-15:00，周末不交易；
2、中国香港股市的交易时间为：周一至五的9:30-12:00，下午13:00-16:00，周末不交易；
3、美国股票的交易时间为：夏21:30-4:00，冬令22:30-5:00；
4、欧洲股市的交易时间为：夏15:00-23:30，冬16:00-0:30。
*/
