package sina

// sinaTargetRegionGDP 新浪财经宏观指标地区GDP接口返回的数据结构
type sinaTargetRegionGDP struct {
	Data [][]string `json:"data"`
}

// sinaTargetRegionCPI 新浪财经宏观指标地区CPI接口返回的数据结构
type sinaTargetRegionCPI struct {
	Data map[string][][]string `json:"data"`
}
