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
	Url        string
	Title      string
	Summary    string
	Source     string
	SourceCode string
	Date       string
}

type SqlValues struct {
	NEWS_GUID        string `json:"NEWS_GUID"`
	NEWS_TITLE       string `json:"NEWS_TITLE"`
	NEWS_TITLE_EN    string `json:"NEWS_TITLE_EN"`
	NEWS_TS          string `json:"NEWS_TS"`
	NEWS_URL         string `json:"NEWS_URL"`
	NEWS_SOURCE      string `json:"NEWS_SOURCE"`
	NEWS_SOURCE_CODE string `json:"NEWS_SOURCE_CODE"`
	NEWS_SUMMARY     string `json:"NEWS_SUMMARY"`
	POLICY_TYPE      string `json:"POLICY_TYPE"`
	POLICY_TYPE_NAME string `json:"POLICY_TYPE_NAME"`
	REGION_CODE      string `json:"REGION_CODE"`
	REGION_NAME      string `json:"REGION_NAME"`
	IS_CONTROL       string `json:"IS_CONTROL"`
	IS_INVEST        string `json:"IS_INVEST"`
	IS_DEPOSIT       string `json:"IS_DEPOSIT"`
	IS_FUND          string `json:"IS_FUND"`
	IS_STOCK         string `json:"IS_STOCK"`
	IS_FINANCE       string `json:"IS_FINANCE"`
	IS_INDUSTRY      string `json:"IS_INDUSTRY"`
	IS_CAPITAL       string `json:"IS_CAPITAL"`
	NEWS_GYS_CODE    string `json:"NEWS_GYS_CODE"`
	NEWS_GYS_NAME    string `json:"NEWS_GYS_NAME"`
	NEWS_ID          int    `json:"NEWS_ID"`
}
