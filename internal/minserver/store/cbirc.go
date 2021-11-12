package store

type DetailsCbirc struct {
	DocId int `json:"docId"`
}

type Rows struct {
	Rows []DetailsCbirc `json:"rows"`
}

type JsonCbirc struct {
	Data Rows `json:"data"`
}

type DataCbirc struct {
	DocId    int `json:"docId"`
	DocTitle string `json:"docTitle"`
	DocDate  string `json:"publishDate"`
	DocClob  string `json:"docClob"`
}

type JsonDetailsCbirc struct {
	DataCbirc `json:"data"`
}

var (
	DetailUrl = "http://www.cbirc.gov.cn/cn/static/data/DocInfo/SelectByDocId/data_docId=%v.json"
	PageUrl = "https://www.cbirc.gov.cn/cn/view/pages/ItemDetail.html?docId=%v"
)

