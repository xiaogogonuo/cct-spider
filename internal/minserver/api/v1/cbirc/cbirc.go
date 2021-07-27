package cbirc

import (
	"encoding/json"
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/minserver/store"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/callback"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/parse"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/request"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/urlprocess"
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
	"net/http"
	"strings"
	"sync"
)

func GetPageUrlList(url string, infoChan chan<- *callback.InfoChan, wg *sync.WaitGroup) {
	defer wg.Done()
	num := 1
	if urlprocess.GetParseQuery(url, "itemId") == "928" {
		num = 100
	}
	for i := 1; i <= num; i++ {
		req := request.Request{
			Url:    fmt.Sprintf("%s%v", url, i),
			Method: http.MethodGet,
		}
		b, err := req.Visit()
		if err != nil {
			logger.Error(err.Error(), logger.Field("url", url))
			return
		}
		var j store.JsonCbirc
		err = json.Unmarshal(b, &j)
		if err != nil {
			logger.Error(err.Error())
		}
		if len(j.Data.Rows) == 0 {
			break
		}
		for _, v := range j.Data.Rows {
			//fmt.Printf(store.DetailUrl, v.DocId)
			infoChan <- &callback.InfoChan{
				Url:      fmt.Sprintf(store.DetailUrl, v.DocId),
				GetInfoF: GetHtmlInfo,
			}
		}
	}
}

func GetHtmlInfo(url string, errChan chan<- *callback.InfoChan, message chan<- *callback.Message) {
	infoMap := make(map[string]string)
	req := request.Request{
		Url:    url,
		Method: http.MethodGet,
	}
	b, err := req.Visit()
	if err != nil {
		return
	}
	var j store.JsonDetailsCbirc
	err = json.Unmarshal(b, &j)
	if err != nil {
		logger.Error(err.Error())
	}

	p := parse.Parse{
		Html:         j.DocClob,
		TextSelector: "p",
	}
	_, data, _ := p.GetTextByParseHtml()
	infoMap[j.DocTitle] = strings.Join(data, "")
	date := strings.Replace(strings.Split(j.DocDate, " ")[0], "-", "", -1)
	message <- &callback.Message{
		Url:        fmt.Sprintf(store.PageUrl, j.DocId),
		Title:      strings.Replace(j.DocTitle, `'`, `"`, -1),
		Summary:    strings.Replace(strings.Join(data, ""), `'`, `"`, -1),
		Source:     "中国银行保险监督管理委员会",
		SourceCode: "WEB_01045",
		Date:       date,
	}
}
