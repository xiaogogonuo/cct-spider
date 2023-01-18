package sina

// 新浪财经指标接口返回模型

// RegionGDP 地区生产总值
type RegionGDP struct {
	Data [][]string `json:"data"`
}

// RegionCPI 地区居民消费价格指数
type RegionCPI struct {
	Data map[string][][]string `json:"data"`
}
