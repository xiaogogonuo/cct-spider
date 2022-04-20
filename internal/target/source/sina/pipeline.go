package sina

import (
	"encoding/json"
	"github.com/xiaogogonuo/cct-spider/internal/target/model"
	"github.com/xiaogogonuo/cct-spider/internal/target/pkg/province"
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
	"strings"
	"unsafe"
)

// filter 过滤器，对新浪财经宏观经济接口返回的字符串做清洗，提取出json形式的数据供后端反序列化成结构体
func filter(body []byte) []byte {
	s := *(*string)(unsafe.Pointer(&body))
	index := strings.Index(s, "config")
	if index == -1 {
		return nil
	}
	config := s[index-1 : len(s)-3]
	config = strings.ReplaceAll(config, "Cmap", `"xxx"`)
	config = strings.ReplaceAll(config, "C", `"C"`)
	config = strings.ReplaceAll(config, "mapid", `"mapid"`)
	config = strings.ReplaceAll(config, "mapname", `"mapname"`)
	config = strings.ReplaceAll(config, "all", `"all"`)
	config = strings.ReplaceAll(config, "data", `"data"`)
	config = strings.ReplaceAll(config, "count", `"count"`)
	config = strings.ReplaceAll(config, "index", `"index"`)
	config = strings.ReplaceAll(config, "title", `"title"`)
	config = strings.ReplaceAll(config, "config", `"config"`)
	config = strings.ReplaceAll(config, "except", `"except"`)
	config = strings.ReplaceAll(config, "querylist", `"querylist"`)
	config = strings.ReplaceAll(config, "conditions", `"conditions"`)
	config = strings.ReplaceAll(config, "defaultItems", `"defaultItems"`)
	config = strings.ReplaceAll(config, "'", `"`)
	return []byte(config)
}

// sinaEconomicTargetRegionGDPPipeline 新浪财经宏观指标`地方生产总值`数据处理管道
func sinaEconomicTargetRegionGDPPipeline(body []byte) (responses []model.Response) {
	var regionGDP sinaTargetRegionGDP
	if err := json.Unmarshal(body, &regionGDP); err != nil {
		logger.Error(err.Error())
		return
	}
	for _, data := range regionGDP.Data {
		var response model.Response
		if data[1] == "全国" || strings.Contains(data[1], "地区") {
			continue
		}
		response.Date = data[0]
		response.TargetValue = data[2]
		response.RegionName = province.UnifyProvinceName(data[1])
		response.RegionCode = province.ChinaProvinceName2Code(response.RegionName)
		responses = append(responses, response)
	}
	return
}

// sinaEconomicTargetRegionCPIPipeline 新浪财经宏观指标`地区居民消费价格指数`数据处理管道
func sinaEconomicTargetRegionCPIPipeline(body []byte) (responses []model.Response) {
	var regionCPI sinaTargetRegionCPI
	if err := json.Unmarshal(body, &regionCPI); err != nil {
		logger.Error(err.Error())
		return
	}
	lj := regionCPI.Data["累计"]
	if len(lj) == 0 {
		return
	}
	for _, data := range lj {
		if data[1] == "全国" {
			continue
		}
		var response model.Response
		date := strings.Split(data[0], ".")
		year := date[0]
		month := ""
		if len(date[1]) == 1 {
			month = "0" + date[1]
		} else {
			month = date[1]
		}
		response.Date = year + month
		response.TargetValue = data[2]
		response.RegionName = province.UnifyProvinceName(data[1])
		response.RegionCode = province.ChinaProvinceName2Code(response.RegionName)
		responses = append(responses, response)
	}
	return
}
