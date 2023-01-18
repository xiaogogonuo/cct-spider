package sci

import (
	"encoding/json"
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/api"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/model"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/pkg/arithmetic"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/pkg/downloader"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/pkg/safeguard"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/poster"
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

func sciIndustryCOI(ic *model.IndexConfig) (buffers []*model.Buffer) {
	var pd = PostData{
		HY:    "能源",
		Level: "1",
		Path1: "能源价格指数",
		Path2: "石油价格指数",
		Path3: "",
		Path4: "",
		Type:  "2",
	}
	return spiderSCI(ic, pd)
}

func sciIndustryPII(ic *model.IndexConfig) (buffers []*model.Buffer) {
	var pd = PostData{
		HY:    "造纸",
		Level: "0",
		Path1: "造纸行业价格指数",
		Path2: "",
		Path3: "",
		Path4: "",
		Type:  "2",
	}
	return spiderSCI(ic, pd)
}

// spiderSCI 爬取卓创资讯的任意指数
func spiderSCI(ic *model.IndexConfig, postData PostData) (buffers []*model.Buffer) {
	reader, _ := json.Marshal(postData)
	header := map[string]string{"Content-Type": "application/json"}
	body, err := downloader.Post(api.SCI, reader, header)

	if err != nil {
		if !safeguard.IsNetworkNormal() {
			logger.Error("请检查服务器的网络是否能联通外网")
			return
		}
		logger.Error(err.Error())
		go poster.Poster(ic)
		return
	}

	var sci SCI
	if err = json.Unmarshal(body, &sci); err != nil {
		logger.Error(err.Error())
		go poster.Poster(ic)
		return
	}

	for _, front := range sci.List {
		buffer := &model.Buffer{}
		buffer.Date = strings.ReplaceAll(front.DataDate, "/", "")
		buffer.TargetValue = arithmetic.TrimZero(fmt.Sprintf("%.2f", front.MDataValue))
		buffers = append(buffers, buffer)
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
