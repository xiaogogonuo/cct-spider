package callback

import (
	"sync"
)

type UrlChan struct {
	Url     string
	GetUrlF func(url string, urlChan chan<- *UrlChan, infoChan chan<- *InfoChan)
}

type InfoChan struct {
	Url      string
	GetInfoF func(url string, infoChan chan<- *InfoChan, message chan<- *Message)
}

type getUrlFunc func(url string, urlChan chan<- *UrlChan, infoChan chan<- *InfoChan)
type getInfoFunc func(url string, infoChan chan<- *InfoChan, message chan<- *Message)

func (f getUrlFunc) callUrl(url string, urlChan chan *UrlChan, infoChan chan *InfoChan) {
	f(url, urlChan, infoChan)
}

func (f getInfoFunc) callInfo(url string, infoChan chan *InfoChan, message chan *Message) {
	f(url, infoChan, message)
}

type callUrl interface {
	callUrl(string, chan *UrlChan, chan *InfoChan)
}

type callInfo interface {
	callInfo(string, chan *InfoChan, chan *Message)
}

func callback1(url string, c callUrl, urlChan chan *UrlChan, infoChan chan *InfoChan) {
	c.callUrl(url, urlChan, infoChan)
}

func callback2(url string, c callInfo, infoChan chan *InfoChan, message chan *Message) {
	c.callInfo(url, infoChan, message)
}

func (u UrlChan) GetUrlFunc(urlChan chan *UrlChan, infoChan chan *InfoChan, wg *sync.WaitGroup) {
	defer wg.Done()
	callback1(u.Url, getUrlFunc(u.GetUrlF), urlChan, infoChan)

}

func (u InfoChan) GetInfoFunc(infoChan chan *InfoChan, message chan *Message, wg *sync.WaitGroup) {
	defer wg.Done()
	callback2(u.Url, getInfoFunc(u.GetInfoF), infoChan, message)
}

type Message struct {
	Id         string
	Url        string
	Title      string
	Summary    string
	Source     string
	SourceCode string
	Date       string
}

type SqlValues struct {
	NEWS_GUID        string `json:"newsGuid"`
	NEWS_TITLE       string `json:"newsTitle"`
	NEWS_TITLE_EN    string `json:"newsTitleEn"`
	NEWS_TS          string `json:"newsTs"`
	NEWS_URL         string `json:"newsUrl"`
	NEWS_SOURCE      string `json:"newsSource"`
	NEWS_SOURCE_CODE string `json:"newsSourceCode"`
	NEWS_SUMMARY     string `json:"newsSummary"`
	POLICY_TYPE      string `json:"policyType"`
	POLICY_TYPE_NAME string `json:"policyTypeName"`
	REGION_CODE      string `json:"regionCode"`
	REGION_NAME      string `json:"regionName"`
	IS_CONTROL       string `json:"isControl"`
	IS_INVEST        string `json:"isInvest"`
	IS_DEPOSIT       string `json:"isDeposit"`
	IS_FUND          string `json:"isFund"`
	IS_STOCK         string `json:"isStock"`
	IS_FINANCE       string `json:"isFinance"`
	IS_INDUSTRY      string `json:"isIndustry"`
	IS_CAPITAL       string `json:"isCapital"`
	NEWS_GYS_CODE    string `json:"newsGysCode"`
	NEWS_GYS_NAME    string `json:"newsGysName"`
	NEWS_ID          int    `json:"newsId"`
}
