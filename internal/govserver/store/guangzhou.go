package store

type GZTotal struct {
	Total int `json:"total"`
}

type Articles struct {
	Url         string `json:"url"`
	ExpiredTime int    `json:"expired_time"`
}

type GZArticles struct {
	Articles []Articles `json:"articles"`
}
