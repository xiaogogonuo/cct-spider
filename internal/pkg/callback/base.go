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
	NEWS_GUID        string
	NEWS_TITLE       string
	NEWS_TITLE_EN    string
	NEWS_TS          string
	NEWS_URL         string
	NEWS_SOURCE      string
	NEWS_SOURCE_CODE string
	NEWS_SUMMARY     string
	POLICY_TYPE      string
	POLICY_TYPE_NAME string
	REGION_CODE      string
	REGION_NAME      string
	IS_CONTROL       string
	IS_INVEST        string
	IS_DEPOSIT       string
	IS_FUND          string
	IS_STOCK         string
	IS_FINANCE       string
	IS_INDUSTRY      string
	IS_CAPITAL       string
	NEWS_GYS_CODE    string
	NEWS_GYS_NAME    string
	NEWS_ID          int
}
