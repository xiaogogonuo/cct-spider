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
// http://finance.sina.com.cn/mac/#price-0-0-31-2

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
var sinaRegionGDPURL = "https://quotes.sina.cn/mac/api/jsonp_v3.php/SINAREMOTECALLCALLBACK1638751228372/" +
	"MacPage_Service.get_pagedata?cate=nation&event=7&from=1&num=%d&condition="

func visitSinaRegionGDP() (respBytes []byte, err error) {
	queryRow := (time.Now().Year() - 1992 + 1) * 31
	resp, err := http.Get(fmt.Sprintf(sinaRegionGDPURL, queryRow))
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
		return
	}
	b, err = GBKToUTF8(b)
	if err != nil {
		return
	}
	br := selectJson(b)
	var regionGDP RegionGDP
	err = json.Unmarshal(br, &regionGDP)
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

// CPI 新浪财经居民消费价格指数接口返回的结构体
type CPI struct {
	Data [][]string `json:"data"`
}

var sinaCPIURL = "https://quotes.sina.cn/mac/api/jsonp_v3.php/SINAREMOTECALLCALLBACK1639556513165/" +
	"MacPage_Service.get_pagedata?cate=price&event=0&from=%d&num=31&condition="

func visitSinaCPI(url string) (respBytes []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	defer resp.Body.Close()
	respBytes, err = io.ReadAll(resp.Body)
	return
}

// RespondSinaCPI 返回新浪财经的居民消费价格指数
//// 适用指标：居民消费价格指数
func RespondSinaCPI() (row []Respond) {
	row = make([]Respond, 0)
	for i := 0; ; i++ {
		_from := 31 * i
		url := fmt.Sprintf(sinaCPIURL, _from)
		b, err := visitSinaCPI(url)
		if err != nil {
			continue
		}
		b, err = GBKToUTF8(b)
		if err != nil {
			continue
		}
		br := selectJson(b)
		var cpi CPI
		if err := json.Unmarshal(br, &cpi); err != nil {
			logger.Error(err.Error())
			continue
		}
		if len(cpi.Data) == 0 || cpi.Data == nil {
			break
		}
		for _, data := range cpi.Data {
			var respond Respond
			date := strings.Split(data[0], ".")
			year := date[0]
			month := ""
			if len(date[1]) == 1 {
				month = "0" + date[1]
			} else {
				month = date[1]
			}
			respond.Date = year + month
			respond.TargetValue = data[1]
			row = append(row, respond)
		}
	}
	return
}

// 新浪财经地区居民消费价格指数
var sinaRegionCPIURL = "https://quotes.sina.cn/mac/api/jsonp_v3.php/SINAREMOTECALLCALLBACK1638839653350/" +
	"MacPage_Service.get_pagedata?cate=price&event=2&from=%d&num=31&condition="

type RegionCPI struct {
	Data map[string][][]string `json:"data"`
}

func visitSinaRegionCPI(url string) (respBytes []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	defer resp.Body.Close()
	respBytes, err = io.ReadAll(resp.Body)
	return
}

// RespondSinaRegionCPI 返回新浪财经的地区居民消费价格指数
//// 适用指标：地区居民消费价格指数
func RespondSinaRegionCPI() (row []Respond) {
	row = make([]Respond, 0)
	for i := 0; ; i++ {
		_from := 31 * i
		url := fmt.Sprintf(sinaRegionCPIURL, _from)
		b, err := visitSinaRegionCPI(url)
		if err != nil {
			continue
		}
		b, err = GBKToUTF8(b)
		if err != nil {
			continue
		}
		br := selectJson(b)
		var regionCPI RegionCPI
		if err := json.Unmarshal(br, &regionCPI); err != nil {
			logger.Error(err.Error())
			continue
		}
		// 取累计字段
		lj := regionCPI.Data["累计"]
		if len(lj) == 0 {
			break
		}
		for _, data := range lj {
			if data[1] == "全国" {
				continue
			}
			var respond Respond
			date := strings.Split(data[0], ".")
			year := date[0]
			month := ""
			if len(date[1]) == 1 {
				month = "0" + date[1]
			} else {
				month = date[1]
			}
			respond.Date = year + month
			respond.TargetValue = data[2]
			respond.RegionName = addProvinceName(data[1])
			respond.RegionCode = provinceNameCode(respond.RegionName)
			row = append(row, respond)
		}
	}
	return
}

/* 香港恒生指数
页面展示接口：http://stock.finance.sina.com.cn/hkstock/quotes/HSI.html
数据抓包接口：http://hq.sinajs.cn/rn=1639970108245&list=rt_hkHSI
 */

var HsiAPI = "http://hq.sinajs.cn/rn=1639970108245&list=rt_hkHSI"

func visitHSI(url string) (respBytes []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	defer resp.Body.Close()
	respBytes, err = io.ReadAll(resp.Body)
	return
}

func RespondHSI() (row []Respond) {
	body, err := visitHSI(HsiAPI)
	if err != nil {
		return
	}
	data := strings.Split(string(body), ",")
	var respond Respond
	respond.TargetValue = data[3]
	respond.Date = time.Now().Format("20060102")
	row = append(row, respond)
	return
}

/* 美元兑人民币
页面展示接口：https://finance.sina.com.cn/money/forex/hq/USDCNY.shtml
数据抓包接口：https://hq.sinajs.cn/rn=1639973351984list=fx_susdcny
 */

var UsdCnyApi = "https://hq.sinajs.cn/rn=1639973351984list=fx_susdcny"

func visitUSDCNY(url string) (respBytes []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	defer resp.Body.Close()
	respBytes, err = io.ReadAll(resp.Body)
	return
}

func RespondUSDCNY() (row []Respond) {
	body, err := visitUSDCNY(UsdCnyApi)
	if err != nil {
		return
	}
	data := strings.Split(string(body), ",")
	var respond Respond
	respond.TargetValue = data[3]
	respond.Date = time.Now().Format("20060102")
	row = append(row, respond)
	return
}

/* 人民币兑港币汇率
页面展示接口：https://finance.sina.com.cn/money/forex/hq/HKDCNY.shtml
数据抓包接口：https://hq.sinajs.cn/rn=1639973901502list=fx_shkdcny
 */

var HkdCnyApi = "https://hq.sinajs.cn/rn=1639973901502list=fx_shkdcny"

func visitHKDCNY(url string) (respBytes []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	defer resp.Body.Close()
	respBytes, err = io.ReadAll(resp.Body)
	return
}

func RespondHKDCNY() (row []Respond) {
	body, err := visitUSDCNY(HkdCnyApi)
	if err != nil {
		return
	}
	data := strings.Split(string(body), ",")
	var respond Respond
	respond.TargetValue = data[3]
	respond.Date = time.Now().Format("20060102")
	row = append(row, respond)
	return
}


func selectJson(b []byte) (br []byte) {
	s := string(b)
	index := strings.Index(s, "config")
	if index == -1 {
		return
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
	br = []byte(config)
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
