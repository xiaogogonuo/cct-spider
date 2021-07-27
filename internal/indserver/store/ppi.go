package store

type PPIRec struct {
	Fi string `json:"fi"`
}

type PPIJson struct {
	Page string   `json:"page"`
	Rec  []PPIRec `json:"rec"`
}
