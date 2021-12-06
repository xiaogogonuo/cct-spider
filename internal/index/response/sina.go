package response

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"
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
	respBytes, err = io.ReadAll(resp.Body)
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

// RegionGDP 新浪财经地区生产总值接口返回的结构体
type RegionGDP struct {
	Data [][]string `json:"data"`
}

// 新浪财经地区生产总值接口
var sinaRegionGDBURL = "https://quotes.sina.cn/mac/api/jsonp_v3.php/SINAREMOTECALLCALLBACK1638751228372/" +
	"MacPage_Service.get_pagedata?cate=nation&event=7&from=1&num=%d&condition="

func visitSinaRegionGDP() (respBytes []byte, err error) {
	queryRow := (time.Now().Year() - 1992 + 1) * 31
	resp, err := http.Get(fmt.Sprintf(sinaRegionGDBURL, queryRow))
	if err != nil {
		logger.Error(err.Error())
		return
	}
	defer resp.Body.Close()
	respBytes, err = io.ReadAll(resp.Body)
	return
}

// RespondSinaRegionGDP 返回新浪财经的地区生产总值数据
// 适用指标：地区生产总值
func RespondSinaRegionGDP() (row []Respond) {
	row = make([]Respond, 0)
	b, err := visitSinaRegionGDP()
	if err != nil {
		logger.Error(err.Error())
		return
	}
	b, err = GBKToUTF8(b)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	s := string(b)
	index := strings.Index(s, "config")
	if index == -1 {
		return
	}
	config := s[index-1 : len(s)-3]
	config = strings.ReplaceAll(config, "all", `"all"`)
	config = strings.ReplaceAll(config, "data", `"data"`)
	config = strings.ReplaceAll(config, "count", `"count"`)
	config = strings.ReplaceAll(config, "index", `"index"`)
	config = strings.ReplaceAll(config, "title", `"title"`)
	config = strings.ReplaceAll(config, "config", `"config"`)
	config = strings.ReplaceAll(config, "except", `"except"`)
	config = strings.ReplaceAll(config, "querylist", `"querylist"`)
	config = strings.ReplaceAll(config, "defaultItems", `"defaultItems"`)
	var regionGDP RegionGDP
	err = json.Unmarshal([]byte(config), &regionGDP)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	for _, data := range regionGDP.Data {
		var respond Respond
		respond.Date = data[0]
		respond.TargetValue = data[2]
		respond.RegionName = addProvinceName(data[1])
		if data[1] == "全国" || strings.Contains(data[1], "地区") {
			continue
		}
		respond.RegionCode = provinceNameCode(addProvinceName(data[1]))
		row = append(row, respond)
	}
	return
}

func GBKToUTF8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func addProvinceName(s string) (name string) {
	switch s {
	case "北京", "天津", "上海", "重庆":
		name = s + "市"
	case "广西":
		name = s + "壮族自治区"
	case "西藏", "内蒙古":
		name = s + "自治区"
	case "宁夏":
		name = s + "回族自治区"
	case "新疆":
		name = s + "维吾尔族自治区"
	default:
		name = s + "省"
	}
	return
}

func provinceNameCode(name string) (code string) {
	switch name {
	case "北京市":
		code = "110000"
	case "天津市":
		code = "120000"
	case "河北省":
		code = "130000"
	case "山西省":
		code = "140000"
	case "内蒙古自治区":
		code = "150000"
	case "辽宁省":
		code = "210000"
	case "吉林省":
		code = "220000"
	case "黑龙江省":
		code = "230000"
	case "上海市":
		code = "310000"
	case "江苏省":
		code = "320000"
	case "浙江省":
		code = "330000"
	case "安徽省":
		code = "340000"
	case "福建省":
		code = "350000"
	case "江西省":
		code = "360000"
	case "山东省":
		code = "370000"
	case "河南省":
		code = "410000"
	case "湖北省":
		code = "420000"
	case "湖南省":
		code = "430000"
	case "广东省":
		code = "440000"
	case "广西壮族自治区":
		code = "450000"
	case "海南省":
		code = "460000"
	case "重庆市":
		code = "500000"
	case "四川省":
		code = "510000"
	case "贵州省":
		code = "520000"
	case "云南省":
		code = "530000"
	case "西藏自治区":
		code = "540000"
	case "陕西省":
		code = "610000"
	case "甘肃省":
		code = "620000"
	case "青海省":
		code = "630000"
	case "宁夏回族自治区":
		code = "640000"
	case "新疆维吾尔族自治区":
		code = "650000"
	}
	return
}
