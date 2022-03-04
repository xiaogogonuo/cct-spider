package cctgroup

type GroupNews struct {
	NewsGuid       string   `json:"newsGuid"`       // 新闻主键
	NewsTitle      string   `json:"newsTitle"`      // 新闻标题
	NewsTs         string   `json:"newsTs"`         // 新闻日期
	NewsUrl        string   `json:"newsUrl"`        // 新闻链接
	NewsSource     string   `json:"newsSource"`     // 中国诚通控股集团有限公司
	NewsSourceCode string   `json:"newsSourceCode"` // WEB_02019
	NewsSummary    string   `json:"newsSummary"`    // 新闻正文
	PolicyType     string   `json:"policyType"`     // 10
	PolicyTypeName string   `json:"policyTypeName"` // 国家政策
	NewsGysCode    string   `json:"newsGysCode"`    // 90
	NewsGysName    string   `json:"newsGysName"`    // 爬虫
	NewsId         int      `json:"newsId"`         // 0
	Image          [][]byte `json:"image"`          // 新闻图片
}
