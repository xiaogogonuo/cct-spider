package request

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
	"io"
	"net/http"
)

// http://app.finance.ifeng.com/hq/stock_daily.php?code=sh000012

const tbi = "http://app.finance.ifeng.com/hq/stock_daily.php?code=sh000012&begin_day=%d-01-01&end_day=%d-12-31"

func TBI(year int) (respBytes []byte, err error) {
	url := fmt.Sprintf(tbi, year, year)
	resp, err := http.Get(url)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	defer resp.Body.Close()
	respBytes, err = io.ReadAll(resp.Body)
	return
}
