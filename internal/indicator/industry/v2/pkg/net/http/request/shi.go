package request

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
	"io"
	"net/http"
)

const shi = "http://data.eastmoney.com/shibor/shibor.aspx?m=sh&t=99&d=99228&cu=cny&type=009023&p=%d"

// ShiBor 上海银行间同业拆放利率
func ShiBor(page int) (respBytes []byte, err error) {
	url := fmt.Sprintf(shi, page)
	resp, err := http.Get(url)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	defer resp.Body.Close()
	respBytes, err = io.ReadAll(resp.Body)
	return
}
