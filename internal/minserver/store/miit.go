package store


type DetailsMiit struct {
	Url string `json:"url"`
}

type GroupData struct {
	DetailsMiit `json:"data"`
}

type DataResults struct {
	GroupData []GroupData `json:"groupData"`
}

type SearchResult struct {
	DataResults []DataResults `json:"dataResults"`

	Total       int           `json:"total"`
}

type DataMiit struct {
	SearchResult `json:"searchResult"`
}

type JsonMiit struct {
	DataMiit `json:"data"`
}

var BaseUrl = "https://www.miit.gov.cn"