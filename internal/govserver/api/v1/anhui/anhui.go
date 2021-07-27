package anhui

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
	baseUrl := `https://www.ah.gov.cn/site/label/8888?IsAjax=1&labelName=maYaService&pageSize=13&pageIndex=%v&domain=http%%3A%%2F%%2F192.168.60.62%%3A8090%%2Fxxgk%%2Fsitesearch%%3FpageSize%%3D13%%26page%%3D%v%%26pageIndex%%3D%v%%26scopeType%%3D1%%26matchType%%3D1%%26rangeType%%3D1%%26classname%%3D%%26documentType%%3D%%26dsId%%3Dwww.ah.gov.cn%%26q%%3D%%26t%%3D1627027993128&file=%%2Fahxxgk%%2Fxxgk%%2FpublicInfoList_new2020_ah_node6`

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
	num := pr.GetPageNum("total:[0-9]+")
	if num == 0 {
		num = 20
	}
	for i := 1; i < num/13; i++ {
		urlChan <- &callback.UrlChan{
			Url:     fmt.Sprintf(baseUrl, i, i, i),
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
			UrlSelector: ".xxgk_nav_con a",
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
			Source:        "安徽省人民政府",
			SourceCode:    "WEB_01112",
			DateSelector:  "#zcjdDiv>table>tbody>tr:nth-child(3)>td[class='pmingcheng1']",
			TitleSelector: "h1",
			TextSelector:  "div[class='wzcon j-fontContent']",
			DomainName:    "www.ah.gov.cn/",
		},
	}
	message <- pr.GetHtmlInfo()
}
