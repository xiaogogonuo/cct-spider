package response

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
	"io"
	"net/http"
	"strings"
)

// 卓创资讯

const (
	sciURL = "https://index.sci99.com/api/zh-cn/dataitem/datavalue"
)

/* Post Data
type:
 - 1: 1个月
 - 2: 3个月
 - 3: 1年
 - 4: 点击购买更多历史数据

木浆价格指数
hy: "造纸"
level: 2
path1: "造纸行业价格指数"
path2: "造纸原料价格指数"
path3: "木浆价格指数"
path4: ""
type: "2"

白卡纸价格指数
hy: "造纸"
level: 3
path1: "造纸行业价格指数"
path2: "造纸原纸价格指数"
path3: "包装用纸价格指数"
path4: "白卡纸价格指数"
type: "2"

造纸原料价格指数
hy: "造纸"
level: 1
path1: "造纸行业价格指数"
path2: "造纸原料价格指数"
path3: ""
path4: ""
type: "2"

造纸行业价格指数
hy: "造纸"
level: 0
path1: "造纸行业价格指数"
path2: ""
path3: ""
path4: ""
type: "2"
*/

// PostData 卓创资讯接口的请求入参
type PostData struct {
	HY    string `json:"hy"`
	Level string `json:"level"`
	Path1 string `json:"path1"`
	Path2 string `json:"path2"`
	Path3 string `json:"path3"`
	Path4 string `json:"path4"`
	Type  string `json:"type"`
}

// StructSCI 卓创资讯接口返回的数据结构
type StructSCI struct {
	List []SCIData `json:"List"`
}

type SCIData struct {
	MDataValue float64 `json:"MDataValue"`
	DataDate   string  `json:"DataDate"`
}

// visitSCI 请求卓创资讯接口
func visitSCI(pd PostData) (respBytes []byte, err error) {
	bt, _ := json.Marshal(pd)
	req, err := http.NewRequest(http.MethodPost, sciURL, bytes.NewReader(bt))
	if err != nil {
		logger.Error(err.Error())
		return
	}
	req.Header.Add("Content-Type", "application/json")
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	defer resp.Body.Close()
	respBytes, err = io.ReadAll(resp.Body)
	//byte数组直接转成string，优化内存
	//str := (*string)(unsafe.Pointer(&respBytes))
	return
}

// RespondSCI 返回卓创资讯响应数据
func RespondSCI(pd PostData) (row []Respond) {
	row = make([]Respond, 0)
	b, err := visitSCI(pd)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	var s StructSCI
	if err = json.Unmarshal(b, &s); err != nil {
		logger.Error(err.Error())
		return
	}
	for _, front := range s.List {
		var respond Respond
		respond.TargetValue = fmt.Sprintf("%.2f", front.MDataValue)
		respond.Date = strings.ReplaceAll(front.DataDate, "/", "")
		row = append(row, respond)
	}
	return
}
