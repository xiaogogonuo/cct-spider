package tianjin

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
			PageNumSelector: "body script",
		},
	}
	num := pr.GetPageNum("parseInt\\('[0-9]+}'\\)")
	if num == 0 {
		num = 70
	}
	for i := 1; i < num; i++ {
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
			UrlSelector: "div[class='list-item clear']>a",
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
			Source:        "天津市人民政府",
			SourceCode:    "WEB_01373",
			DateSelector:  "#content_fbrq",
			TitleSelector: "div[class='sx-item item-all']>div[class='sx-con']",
			TextSelector:  "div[class*='view TRS_UEDITOR']",
			DomainName:    "www.tj.gov.cn",
		},
	}
	message <- pr.GetHtmlInfo()

}
