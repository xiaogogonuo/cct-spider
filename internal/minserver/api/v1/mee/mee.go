package mee

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
			BaseUrl:     url,
			UrlSelector: "div[class='outBox zcwj']>div>a",
		},
	}
	for _, link := range pr.GetPageUrl("href") {
		urlChan <- &callback.UrlChan{
			Url:     link,
			GetUrlF: GetSecondUrl,
		}
		//time.Sleep(time.Second*1)  // 单跑这个需要加延迟，
	}

}

func GetSecondUrl(url string, urlChan chan<- *callback.UrlChan, infoChan chan<- *callback.InfoChan) {
	s := strings.Split(url, "/")
	switch s[4] {
	case "zyygwj", "gwywj", "xzspwj":
		urlChan <- &callback.UrlChan{
			Url:     url,
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
				UrlSelector: "span[class='mobile_none']>a",
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
			PageNumSelector: ".slideTxtBoxgsf script",
		},
	}
	num := pr.GetPageNum("var countPage = [0-9]+//")
	if num == 0 {
		num = 40
	}
	for i := 1; i < num; i++ {
		urlChan <- &callback.UrlChan{
			Url:     fmt.Sprintf("%sindex_%v.shtml", url, i),
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
			BaseUrl:     url,
			UrlSelector: "#div>li>a",
		},
	}
	for _, link := range pr.GetPageUrl("href") {
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
			Source:        "生态环境部",
			SourceCode:    "WEB_01013",
			DateSelector:  ".wjkFontBox>em, .content_top_box, span[class='xqLyPc time']",
			TitleSelector: "h1, .neiright_Box>h2",
			TextSelector:  ".Custom_UnionStyle p, .Custom_UnionStyle div, .content_body_box p, .content_body_box div, .neiright_JPZ_GK_CP>p, .neiright_JPZ_GK_CP div, .TRS_Editor p .TRS_Editor div",
			DomainName:    "http://www.mee.gov.cn",
		},
	}
	message <- pr.GetHtmlInfo()

}
