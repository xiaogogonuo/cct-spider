package sasac

type GroupNews struct {
	NewsGuid       string   `json:"newsGuid"`
	NewsTitle      string   `json:"newsTitle"`
	NewsTs         string   `json:"newsTs"`
	NewsUrl        string   `json:"newsUrl"`
	NewsSource     string   `json:"newsSource"`     // 国务院国有资产监督管理委员会
	NewsSourceCode string   `json:"newsSourceCode"` // WEB_01024
	NewsSummary    string   `json:"newsSummary"`
	PolicyType     string   `json:"policyType"`     // 10
	PolicyTypeName string   `json:"policyTypeName"` // 国家政策
	NewsGysCode    string   `json:"newsGysCode"`    // 90
	NewsGysName    string   `json:"newsGysName"`    // 爬虫
	NewsId         int      `json:"newsId"`         // 0
	Image          [][]byte `json:"image"`
}
