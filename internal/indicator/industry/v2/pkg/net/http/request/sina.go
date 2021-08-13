package request

import (
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
	"io/ioutil"
	"net/http"
)

const (
	sinaURL = "https://vip.stock.finance.sina.com.cn/forex/api/jsonp.php/_/NewForexService.getDayKLine?symbol="
)

// VisitSina 新浪财经
func VisitSina(symbol string) (respBytes []byte, err error) {
	resp, err := http.Get(sinaURL + symbol)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	respBytes, err = ioutil.ReadAll(resp.Body)
	//byte数组直接转成string，优化内存
	//str := (*string)(unsafe.Pointer(&respBytes))
	//fmt.Println(*str)
	return
}