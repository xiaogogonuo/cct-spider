package cni

// 国证指数网站人民币汇率返回的数据结构
type cniCNYXTarget struct {
	Code int `json:"code"`
	Data struct {
		Data [][]interface{} `json:"data"`
	} `json:"data"`
}
