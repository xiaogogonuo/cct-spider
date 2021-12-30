package response

import (
	"encoding/json"
	"fmt"
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
	"io"
	"net/http"
	"time"
)

//

/* 人民币指数
页面展示接口：http://www.cnindex.com.cn/module/index-detail-CNYX.html?act_menu=1
数据抓包接口：http://hq.cnindex.com.cn/market/market/getIndexRealTimeData?indexCode=CNYX
*/

const CNYXURL = "http://hq.cnindex.com.cn/market/market/getIndexRealTimeData?indexCode=CNYX"

type CNYX struct {
	Code int `json:"code"`
	Data struct {
		Data [][]interface{} `json:"data"`
	} `json:"data"`
}

func visitCNYX(url string) (body []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	defer resp.Body.Close()
	body, err = io.ReadAll(resp.Body)
	return
}

func RespondCNYX() (row []Respond) {
	body, err := visitCNYX(CNYXURL)
	if err != nil {
		return
	}
	var cnyx CNYX
	err = json.Unmarshal(body, &cnyx)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	data := cnyx.Data.Data
	lastOne := data[len(data)-1] // 倒数第一个数据
	lastTwo := data[len(data)-2] // 倒数第二个数据
	var timeStamp float64
	var v []interface{}
	if lastOne[1] != nil {
		timeStamp = lastOne[0].(float64)
		v = lastOne
	} else {
		timeStamp = lastTwo[0].(float64)
		v = lastTwo
	}
	tm := time.Unix(int64(timeStamp)/1000, 0)
	var respond Respond
	respond.Date = tm.Format("20060102")
	row = append(row, respond)
	respond.TargetValue = fmt.Sprintf("%.2f,%.2f,%.2f%s,%.2f,%.2f,%.2f,%s",
		v[1].(float64),                   // 现价
		v[6].(float64),                   // 涨跌
		v[7].(float64),                   // 涨跌幅
		"%",
		v[2].(float64),                   // 最高
		v[4].(float64),                   // 最低
		v[11].(float64),                  // 昨收
		tm.Format("2006-01-02 15:04:05"), // 更新日期
	)
	row = append(row, respond)
	return
}
