package response

import (
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"unsafe"
)

// 新浪财经

const (
	sinaURL = "https://vip.stock.finance.sina.com.cn/forex/api/jsonp.php/_/NewForexService.getDayKLine?symbol="
)

// visitSina 请求新浪财经的接口
// 适用指标：人民币汇率、美元指数
func visitSina(symbol string) (respBytes []byte, err error) {
	resp, err := http.Get(sinaURL + symbol)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	defer resp.Body.Close()
	respBytes, err = ioutil.ReadAll(resp.Body)
	//byte数组直接转成string，优化内存
	//str := (*string)(unsafe.Pointer(&respBytes))
	//fmt.Println(*str)
	return
}

// RespondSina 返回新浪财经的数据
// 适用指标：人民币汇率、美元指数
func RespondSina(sourceTargetCode string) (row []Respond) {
	row = make([]Respond, 0)
	b, err := visitSina(sourceTargetCode)
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
