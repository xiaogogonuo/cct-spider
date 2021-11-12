package logistics

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
			PageNumSelector: ".pagination>li:nth-child(10)>a",
		},
	}
	num := pr.GetPageNum("[0-9]+")
	if num == 0 {
		num = 80
	}
	for i := 2; i < num; i++ {
		if i < 10{
			urlChan <- &callback.UrlChan{
				Url:     fmt.Sprintf("%sindex_%v.shtml", url, i),
				GetUrlF: GetDetailPageUrl,
			}
		} else
		{
			url = "http://www.chinawuliu.com.cn/zcms/ui/catalog/15166/pc/"
			urlChan <- &callback.UrlChan{
				Url:     fmt.Sprintf("%sindex_%v.shtml", url, i),
				GetUrlF: GetDetailPageUrl,
			}
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
			UrlSelector: "div[class='col-sm-8 leftRow'] ul[class='list-box list-box--pre']>li>a",
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
			Source:        "中国物流与信息采购联合会",
			SourceCode:    "L_00008",
			DateSelector:  ".new-time>span",
			TitleSelector: ".bg-title",
			TextSelector:  ".newText",
			DomainName:    "www.chinawuliu.com.cn",
		},
	}
	message <- pr.GetHtmlInfo()

}
