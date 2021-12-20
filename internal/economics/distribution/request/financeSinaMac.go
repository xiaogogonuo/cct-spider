package request

import "fmt"

const (
	BaseFinanceSinaURL = "https://quotes.sina.cn/mac/api/jsonp_v3.php/SINAREMOTECALLCALLBACK%s/MacPage_Service.get_pagedata?cate=%s&%s&from=%d&num=31&condition="
)

type FinanceSinaURL struct {
	From     int
	Cate     string
	Event    string
	CallBack string
}

func (f FinanceSinaURL) ToURL() string {
	return fmt.Sprintf(BaseFinanceSinaURL, f.CallBack, f.Cate, f.Event, f.From)
}

func NewFinanceSinaURL() *FinanceSinaURL {
	return &FinanceSinaURL{
		From: 0,
		Cate: "boom",
		CallBack: "1639826157401",
	}
}
