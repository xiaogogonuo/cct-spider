package sci

import (
	"encoding/json"
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/target/model"
	"github.com/xiaogogonuo/cct-spider/internal/target/pkg/downloader"
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
	"strings"
)

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

// APISCITarget 卓创资讯指标接口
var APISCITarget = "https://index.sci99.com/api/zh-cn/dataitem/datavalue"

// SpiderSCITargetCOI 爬取卓创资讯的原油价格指数
// 适用指标：
// - 原油价格指数
//   • 页面展示接口：https://index.sci99.com/channel/product/path2/原油价格指数/2.html
//   • 数据获取接口：https://index.sci99.com/api/zh-cn/dataitem/datavalue
func SpiderSCITargetCOI() (responses []model.Response) {
	var coiPostData = PostData{
		HY:    "能源",
		Level: "1",
		Path1: "能源价格指数",
		Path2: "石油价格指数",
		Path3: "",
		Path4: "",
		Type:  "2",
	}
	return spiderSCI(coiPostData)
}

// SpiderSCITargetPII 爬取卓创资讯的造纸行业价格指数
// 适用指标：
// - 造纸行业价格指数
//   • 页面展示接口：https://index.sci99.com/channel/product/hy/造纸/2.html
//   • 数据获取接口：https://index.sci99.com/api/zh-cn/dataitem/datavalue
func SpiderSCITargetPII() (responses []model.Response) {
	var coiPostData = PostData{
		HY:    "造纸",
		Level: "0",
		Path1: "造纸行业价格指数",
		Path2: "",
		Path3: "",
		Path4: "",
		Type:  "2",
	}
	return spiderSCI(coiPostData)
}

// spiderSCI 爬取卓创资讯的任意指数
func spiderSCI(postData PostData) (responses []model.Response) {
	reader, _ := json.Marshal(postData)
	header := map[string]string{"Content-Type": "application/json"}
	body, err := downloader.Post(APISCITarget, reader, header)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	var sci sciTarget
	if err = json.Unmarshal(body, &sci); err != nil {
		logger.Error(err.Error())
		return
	}
	for _, front := range sci.List {
		var response model.Response
		response.TargetValue = fmt.Sprintf("%.2f", front.MDataValue)
		response.Date = strings.ReplaceAll(front.DataDate, "/", "")
		responses = append(responses, response)
	}
	return
}

/* PostData
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
