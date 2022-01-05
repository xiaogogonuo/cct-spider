package response

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/xiaogogonuo/cct-spider/pkg/db/mysql"
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
	"io"
	"net/http"
	"regexp"
	"strconv"
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
	case "fxWHMP": // 人民币汇率中间价
		return "https://quote.fx678.com/exchange/WHMP"
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
			record := tmp[i*8+1 : (i+1)*8]
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

/* 美元Libor隔夜
页面展示接口：https://quote.fx678.com/rate/libor
数据抓包接口：https://quote.fx678.com/rate/libor
*/

const USLiborURL = "https://quote.fx678.com/rate/libor"

func visitLibor(url string) (body []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	defer resp.Body.Close()
	body, err = io.ReadAll(resp.Body)
	return
}

func RespondUSLibor(targetCode string) (row []Respond) {
	body, err := visitLibor(USLiborURL)
	if err != nil {
		return
	}
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(string(body)))
	if err != nil {
		logger.Error(err.Error())
		return
	}
	var coinDate []string
	dom.Find("h2[class='ttable_area-title']").Each(func(i int, selection *goquery.Selection) {
		coinDate = append(coinDate, strings.Trim(selection.Text(), "\n\t "))
	})
	var libor []string
	dom.Find("table[class='market_tab_gold_add2'] td").Each(func(i int, selection *goquery.Selection) {
		libor = append(libor, strings.Trim(selection.Text(), "\n\t "))
	})
	reg := regexp.MustCompile("[0-9]+")
	monthDay := reg.FindAllString(coinDate[1], -1)
	var year, month, day string
	// 如果从页面提取出来的月份大于系统月份，则代表跨年没更新
	if monthInt, _ := strconv.Atoi(monthDay[0]); monthInt > int(time.Now().Month()) {
		year = fmt.Sprintf("%d", time.Now().Year() - 1)
	} else {
		year = fmt.Sprintf("%d", time.Now().Year())
	}
	if len(monthDay[0]) == 1 {
		month = "0" + monthDay[0]
	} else {
		month = monthDay[0]
	}
	if len(monthDay[1]) == 1 {
		day = "0" + monthDay[1]
	} else {
		day = monthDay[1]
	}
	var respond Respond
	respond.Date = year + month + day
	updateTime := respond.Date[:4] + "-" + respond.Date[4:6] + "-" + respond.Date[6:8] + " " +
		strconv.FormatInt(int64(time.Now().Hour()), 10) + ":" +
		strconv.FormatInt(int64(time.Now().Minute()), 10) + ":" +
		strconv.FormatInt(int64(time.Now().Second()), 10)
	var yesterdayValue string
	// 从数据库提取历史数据
	sql := fmt.Sprintf("SELECT ACCT_DATE, TARGET_VALUE FROM t_dmaa_base_target_value WHERE TARGET_CODE = '%s' ORDER BY ACCT_DATE DESC", targetCode)
	history := mysql.Query(sql)
	if len(history) >= 1 {
		// 如果页面提取的日期与数据库最新的日期一样，则代表没有开市，无数据需要更新
		if history[0][0] == respond.Date {
			return
		}
		// 拿第一条记录的TARGET_VALUE值作为昨天的TARGET_VALUE
		yesterdayValue = history[0][1]
	}
	todayValue := libor[31]
	var upDown, upDownPercent string
	if todayValue == "" || yesterdayValue == "" {
		upDown, upDownPercent = "", ""
	} else {
		// 计算涨跌、涨跌幅
		todayValueF, _ := strconv.ParseFloat(todayValue, 64)
		yesterdayValueF, _ := strconv.ParseFloat(yesterdayValue, 64)
		upDown = fmt.Sprintf("%.2f", todayValueF-yesterdayValueF)
		upDownPercent = fmt.Sprintf("%.2f%s", (todayValueF-yesterdayValueF)/yesterdayValueF, "%")
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

美元人民币：https://quote.fx678.com/symbol/MUSD 24小时
港元人民币：https://quote.fx678.com/symbol/MHKD 24小时
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
