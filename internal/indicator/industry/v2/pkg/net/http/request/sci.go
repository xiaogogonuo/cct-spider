package request

import (
	"bytes"
	"encoding/json"
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
	"io/ioutil"
	"net/http"
)

const (
	sciURL = "https://index.sci99.com/api/zh-cn/dataitem/datavalue"
)

type PostData struct {
	HY    string `json:"hy"`
	Level string `json:"level"`
	Path1 string `json:"path1"`
	Path2 string `json:"path2"`
	Path3 string `json:"path3"`
	Path4 string `json:"path4"`
	Type  string `json:"type"`
}

// VisitSCI 卓创资讯
func VisitSCI(pd PostData) (respBytes []byte, err error) {
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
	respBytes, err = ioutil.ReadAll(resp.Body)
	//byte数组直接转成string，优化内存
	//str := (*string)(unsafe.Pointer(&respBytes))
	return
}

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