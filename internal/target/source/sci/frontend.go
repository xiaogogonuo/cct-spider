package sci

// sciTarget 卓创资讯接口返回的数据结构
type sciTarget struct {
	List []struct{
		MDataValue float64 `json:"MDataValue"`
		DataDate   string  `json:"DataDate"`
	} `json:"List"`
}
