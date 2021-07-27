package chongqing

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/callback"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/parse"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/request"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/response"
	"net/http"
	"strings"
	"sync"
)

func GetFirstUrl(url string, urlChan chan<- *callback.UrlChan, wg *sync.WaitGroup) {
	defer wg.Done()
	pr := response.PR{
		Request: request.Request{
			Url:    url,
			Method: http.MethodGet,
		},
		Parse: parse.Parse{
			UrlSelector: "a[class='item-top-right rt hover']",
			BaseUrl:     url,
		},
	}
	urlList := pr.GetPageUrl("href")
	for _, link := range urlList {
		urlChan <- &callback.UrlChan{
			Url:     link,
			GetUrlF: GetSecondUrl,
		}

	}
}

func GetSecondUrl(url string, urlChan chan<- *callback.UrlChan, infoChan chan<- *callback.InfoChan) {

	s := strings.Split(url, "/")
	switch s[6] {
	case "qtgw":
		urlChan <- &callback.UrlChan{
			Url:     url,
			GetUrlF: GetPageUrlList,
		}
	case "zfgz":
		urlChan <- &callback.UrlChan{
			Url:     url + "zfgz/",
			GetUrlF: GetPageUrlList,
		}
	default:
		pr := response.PR{
			Request: request.Request{
				Url:    url,
				Method: http.MethodGet,
			},
			Parse: parse.Parse{
				BaseUrl:     url,
				UrlSelector: "a[class='item-top-right rt']",
			},
		}
		for _, link := range pr.GetPageUrl("href") {
			urlChan <- &callback.UrlChan{
				Url:     link,
				GetUrlF: GetPageUrlList,
			}
		}
	}
}

func GetPageUrlList(url string, urlChan chan<- *callback.UrlChan, infoChan chan<- *callback.InfoChan) {

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
			PageNumSelector: "#page>script",
		},
	}
	num := pr.GetPageNum("createPage\\([0-9]+,")
	if num == 0 {
		num = 20
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
			UrlSelector: ".list a",
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
			Source:        "重庆市人民政府",
			SourceCode:    "WEB_01460",
			DateSelector:  "table[class='table']>tbody>tr:nth-child(4)>td:nth-child(4)",
			TitleSelector: "td[colspan='5']",
			TextSelector:  "div[class='document mt-1 mt-12']",
			DomainName:    "www.cq.gov.cn/",
		},
	}
	message <- pr.GetHtmlInfo()

}
