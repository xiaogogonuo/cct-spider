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
			record := tmp[i*8+1:(i+1)*8]
			record[len(record)-1] = time.Now().Format("2006-01-02") + " " + record[len(record)-1]
			data = strings.Join(record, ",")
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
恒生指数：https://quote.fx678.com/symbol/HSI 09:25-16:00

日经指数：https://quote.fx678.com/symbol/NIKKI  08:00-14:00
日债10年收益率：https://quote.fx678.com/symbol/GJGB10 08:05-16:00

道琼斯工业指数：https://quote.fx678.com/symbol/DJIA 22:35-04:15
纳斯达克指数：https://quote.fx678.com/symbol/NASDAQ 22:35-04:15
标准普尔500指数：https://quote.fx678.com/symbol/SP500 22:35-04:15

斯托克600：https://quote.fx678.com/symbol/SXO 16:00-01:00
英国FTSE100：https://quote.fx678.com/symbol/FTSE 16:20-23:45
德债10年收益率：https://quote.fx678.com/symbol/GDBR10 15:20-01:00
英债10年收益率：https://quote.fx678.com/symbol/GUKG10 16:05-01:00

欧元美元 ：https://quote.fx678.com/symbol/EURUSD 24小时
美元日元：https://quote.fx678.com/symbol/USDJPY 24小时
英镑美元：https://quote.fx678.com/symbol/GBPUSD 24小时
原油：https://quote.fx678.com/symbol/OILC 24小时
黄金：https://quote.fx678.com/symbol/GLNC  24小时
白银：https://quote.fx678.com/symbol/SLNC  24小时
铜：https://quote.fx678.com/symbol/LMCI  24小时
美元指数：https://quote.fx678.com/symbol/USD 24小时
美债10年收益率：https://quote.fx678.com/symbol/USG10Y 24小时

1、中国沪深市场交易时间为：周一至五的9:30-11:30，下午13:00-15:00，周末不交易；
2、中国香港股市的交易时间为：周一至五的9:30-12:00，下午13:00-16:00，周末不交易；
3、美国股票的交易时间为：夏21:30-4:00，冬令22:30-5:00；
4、欧洲股市的交易时间为：夏15:00-23:30，冬16:00-0:30。
5、港股开盘和收盘时间是早上9点半-12点,下午12点半-4点的
*/
