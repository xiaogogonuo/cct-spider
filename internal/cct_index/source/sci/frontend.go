package sci

// 卓创资讯指标接口返回模型

type SCI struct {
	List []struct {
		MDataValue float64 `json:"MDataValue"`
		DataDate   string  `json:"DataDate"`
	} `json:"List"`
}
