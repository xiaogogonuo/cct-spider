package province

import "strings"

// UnifyProvinceName 统一化省份名
func UnifyProvinceName(province string) (name string) {
	if strings.Contains(province, "北京") {
		return "北京市"
	}
	if strings.Contains(province, "天津") {
		return "天津市"
	}
	if strings.Contains(province, "上海") {
		return "上海市"
	}
	if strings.Contains(province, "重庆") {
		return "重庆市"
	}
	if strings.Contains(province, "广西") {
		return "广西壮族自治区"
	}
	if strings.Contains(province, "西藏") {
		return "西藏自治区"
	}
	if strings.Contains(province, "内蒙古") {
		return "内蒙古自治区"
	}
	if strings.Contains(province, "宁夏") {
		return "宁夏回族自治区"
	}
	if strings.Contains(province, "新疆") {
		return "新疆维吾尔族自治区"
	}
	if strings.Contains(province, "省") {
		return province
	}
	return province + "省"
}

// ChinaProvinceName2Code 中国省份-地区编码转换
func ChinaProvinceName2Code(province string) (code string) {
	switch province {
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
