package sarm

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
	num := 20
	pr := response.PR{
		Request: request.Request{
			Url:    url,
			Method: http.MethodGet,
		},
		Parse: parse.Parse{
			PageNumSelector: "head>script",
		},
	}
	// var m_nRecordCount = "133"  var m_nPageSize = 20;
	countR := "var m_nRecordCount = \"[0-9]+\""
	sizeR := "var m_nPageSize = [0-9]+;"
	count, size := pr.GetCountAndSize(countR, sizeR)
	//fmt.Println(count, size)
	if count != 0 && size != 0 {
		num = count/size + 1
	}
	for i := 1; i < num; i++ {
		urlChan <- &callback.UrlChan{
			Url:     fmt.Sprintf("%s_%v.html", url[:len(url)-len(".html")], i),
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
			UrlSelector: ".Three_zhnlist_02 a",
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
			Source:        "国家市场监督管理总局",
			SourceCode:    "WEB_00306",
			DateSelector:  "li[class='Three_xilan01_02 Three_xilan01_0201']",
			TitleSelector: ".xilanboxbg td[colspan='2'] li[class='Three_xilan01_02 Three_xilan01_0201']",
			TextSelector:  ".Three_xilan_07 p",
			DomainName:    "http://gkml.samr.gov.cn/",
		},
	}
	message <- pr.GetHtmlInfo()
}
