package sina

import (
	"encoding/json"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/model"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/pkg/province"
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

func getRegionGDPBuffer(body []byte) (buffers []*model.Buffer, err error) {
	var s RegionGDP
	if err = json.Unmarshal(body, &s); err != nil {
		return
	}
	for _, data := range s.Data {
		if data[1] == "全国" || strings.Contains(data[1], "地区") {
			continue
		}
		buffer := &model.Buffer{}
		buffer.Date = data[0]
		buffer.TargetValue = data[2]
		buffer.RegionName = province.UnifyProvinceName(data[1])
		buffer.RegionCode = province.ChinaProvinceName2Code(buffer.RegionName)
		buffers = append(buffers, buffer)
	}

	return
}

func getRegionCPIBuffer(body []byte) (buffers []*model.Buffer, err error) {
	var s RegionCPI
	if err = json.Unmarshal(body, &s); err != nil {
		return
	}
	lj := s.Data["累计"]
	if len(lj) == 0 {
		return
	}
	for _, data := range lj {
		if data[1] == "全国" {
			continue
		}
		buffer := &model.Buffer{}
		date := strings.Split(data[0], ".")
		year := date[0]
		month := ""
		if len(date[1]) == 1 {
			month = "0" + date[1]
		} else {
			month = date[1]
		}
		buffer.Date = year + month
		buffer.TargetValue = data[2]
		buffer.RegionName = province.UnifyProvinceName(data[1])
		buffer.RegionCode = province.ChinaProvinceName2Code(buffer.RegionName)
		buffers = append(buffers, buffer)
	}
	return
}
