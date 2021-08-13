package response

type Response struct {
	ReturnData ReturnData `json:"returndata"`
}

type ReturnData struct {
	DataNodes []Node `json:"datanodes"`
}

type Node struct {
	Data D
	Code string `json:"code"`
}

type D struct {
	HasData bool    `json:"hasdata"`
	StrData string  `json:"strdata"`
}