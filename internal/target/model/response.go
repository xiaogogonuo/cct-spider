package model

// Response 暂存网页数据
type Response struct {
	Date        string
	TargetValue string
	RegionCode  string
	RegionName  string
}

type ResponseDateStringValueFloat struct {
	Date        string
	TargetValue float64
}

type ResponseArray []ResponseDateStringValueFloat

func (ra ResponseArray) Len() int {
	return len(ra)
}

func (ra ResponseArray) Less(i, j int) bool {
	return ra[i].Date > ra[j].Date // 按照日期降序
}

func (ra ResponseArray) Swap(i, j int) {
	ra[i].Date, ra[j].Date = ra[j].Date, ra[i].Date
}
