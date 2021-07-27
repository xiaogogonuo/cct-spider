package ppi

import (
	"encoding/json"
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/indserver/store"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/callback"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/parse"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/request"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/response"
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

func GetPageUrlList(url string, urlChan chan<- *callback.UrlChan, wg *sync.WaitGroup) {

	defer wg.Done()
	req := request.Request{
		Url:    url,
		Method: http.MethodGet,
	}
	s, err := req.VisitString()
	if err != nil {
		return
	}
	s = s[5 : len(s)-1]
	var j store.PPIJson
	err = json.Unmarshal([]byte(s), &j)
	if err != nil {
		logger.Error(err.Error(), logger.Field("url", url))
		return
	}
	num, err := strconv.Atoi(j.Page)
	if err != nil {
		logger.Error(err.Error(), logger.Field("j.Page", j.Page))
		return
	}
	for i := 1; i <= num; i++ {
		urlChan <- &callback.UrlChan{
			Url:     strings.Replace(url, `&p=1`, fmt.Sprintf(`&p=%v`, i), -1),
			GetUrlF: GetDetailPageUrl,
		}

	}
}
func GetDetailPageUrl(url string, urlChan chan<- *callback.UrlChan, infoChan chan<- *callback.InfoChan) {

	req := request.Request{
		Url:    url,
		Method: http.MethodGet,
	}
	s, err := req.VisitString()
	if err != nil {
		return
	}
	s = strings.Replace(s[5:len(s)-1], "\t", "", -1)
	var j store.PPIJson
	err = json.Unmarshal([]byte(s), &j)
	if err != nil {
		logger.Error(err.Error(), logger.Field("url", url))
		return
	}

	for _, v := range j.Rec {
		infoChan <- &callback.InfoChan{
			Url:      `http://www.chinappi.org/` + v.Fi,
			GetInfoF: GetHtmlInfo,
		}
	}
}
func GetHtmlInfo(url string, errChan chan<- *callback.InfoChan, message chan<- *callback.Message) {

	pr := response.PR{
		Request: request.Request{
			Url:    url,
			Method: http.MethodGet,
		},
		Parse: parse.Parse{
			Source:        "中国造纸协会",
			SourceCode:    "L_00009",
			DateSelector:  ".new_note>span:nth-child(1)",
			TitleSelector: ".new_title",
			TextSelector:  ".new_content",
			DomainName:    "www.chinappi.org",
		},
	}
	message <- pr.GetHtmlInfo()

}
