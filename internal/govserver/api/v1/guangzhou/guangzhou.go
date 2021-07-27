package guangzhou

import (
	"encoding/json"
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/govserver/store"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/callback"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/parse"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/request"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/response"
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
	"net/http"
	"sync"
)

func GetPageUrlList(url string, urlChan chan<- *callback.UrlChan, wg *sync.WaitGroup) {

	defer wg.Done()
	baseUrl := `http://www.gd.gov.cn/gkmlpt/api/all/5?page=%v&sid=2`
	req := request.Request{
		Url:    url,
		Method: http.MethodGet,
	}
	b, err := req.Visit()
	if err != nil {
		return
	}
	var total store.GZTotal
	err = json.Unmarshal(b, &total)
	if err != nil {
		logger.Error(err.Error(), logger.Field("url", url))
		return
	}
	for i := 1; i <= total.Total/100; i++ {
		urlChan <- &callback.UrlChan{
			Url:     fmt.Sprintf(baseUrl, i),
			GetUrlF: GetDetailPageUrl,
		}

	}
}

func GetDetailPageUrl(url string, urlChan chan<- *callback.UrlChan, infoChan chan<- *callback.InfoChan) {

	req := request.Request{
		Url:    url,
		Method: http.MethodGet,
	}
	b, err := req.Visit()
	if err != nil {
		return
	}
	var articles store.GZArticles
	err = json.Unmarshal(b, &articles)
	if err != nil {
		logger.Error(err.Error(), logger.Field("url", url))
		return
	}

	for _, v := range articles.Articles {
		if v.ExpiredTime != 0 {
			logger.Info("to exceed the time limit", logger.Field("url", v.Url))
			continue
		}
		infoChan <- &callback.InfoChan{
			Url:      v.Url,
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
			Source:        "广州市人民政府",
			SourceCode:    "WEB_01187",
			DateSelector:  "div[class='classify']>table>tbody>:nth-child(4)>td[class='td-value']>span",
			TitleSelector: ".content>h1",
			TextSelector:  ".article-content",
			DomainName:    "www.gd.gov.cn",
		},
	}
	message <- pr.GetHtmlInfo()

}
