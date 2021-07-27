package store

type ZJData struct {
	Url string `json:"url"`
}

type ZJJson struct {
	Total int      `json:"total"`
	Data  []ZJData `json:"data"`
}
