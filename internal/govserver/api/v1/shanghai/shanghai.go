package shanghai

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/callback"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/parse"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/request"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/response"
	"net/http"
	"sync"
)

func GetPageUrlList(url string, urlChan chan<- *callback.UrlChan, wg *sync.WaitGroup) {

	defer wg.Done()
	urlChan <- &callback.UrlChan{
		Url:     url,
		GetUrlF: GetDetailPageUrl,
	}
	pr := response.PR{
		Request: request.Request{
			Url:    url,
			Method: http.MethodGet,
		},
		Parse: parse.Parse{
			PageNumSelector: "ul[class='pagination']>script",
		},
	}
	num := pr.GetPageNum("totalPage: [0-9]+,")
	if num == 0 {
		num = 1500
	}
	for i := 2; i < num; i++ {
		urlChan <- &callback.UrlChan{
			Url:     fmt.Sprintf("%sindex_%v.html", url, i),
			GetUrlF: GetDetailPageUrl,
		}
	}
}

func GetDetailPageUrl(url string, urlChan chan<- *callback.UrlChan, infoChan chan<- *callback.InfoChan) {

	pr := response.PR{
		Request: request.Request{
			Url:    url,
			Method: http.MethodGet,
		},
		Parse: parse.Parse{
			UrlSelector: "ul[class='tadaty-list uli14 nowrapli list-date']>li>a",
			BaseUrl:     url,
		},
	}
	urlList := pr.GetPageUrl("href")
	for _, link := range urlList {
		infoChan <- &callback.InfoChan{
			Url:      link,
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
			Source:        "上海市人民政府",
			SourceCode:    "WEB_01348",
			DateSelector:  ".PBtime>p",
			TitleSelector: "title",
			TextSelector:  "#ivs_content",
			DomainName:    "www.shanghai.gov.cn",
		},
	}
	message <- pr.GetHtmlInfo()

}
