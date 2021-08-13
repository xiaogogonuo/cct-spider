package response

type FrontEastMoney []EastMoney

type EastMoney struct {
	Date     string  `json:"DATADATE"`
	Value    float64 `json:"VALUE"`
}


type FrontSCI struct {
	List []Data `json:"List"`
}

type Data struct {
	MDataValue float64 `json:"MDataValue"`
	DataDate   string  `json:"DataDate"`
}