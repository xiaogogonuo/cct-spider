package miit

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/robertkrimen/otto"
	"github.com/xiaogogonuo/cct-spider/internal/minserver/store"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/callback"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/parse"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/request"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/response"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/urlprocess"
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
	"hash"
	"net/http"
	"regexp"
	"strings"
	"sync"
)

var _cookie = _getCookie()

func _getCookie() (cookie string) {
	url := "https://www.miit.gov.cn/search-front-server/api/search/info"
	req := request.Request{
		Url:    url,
		Method: http.MethodGet,
	}
	b, err := req.VisitString()
	if err != nil {
		return
	}
	reg := regexp.MustCompile(`cookie=(\(.*?\));location`)
	jslClearances := reg.FindStringSubmatch(b)
	if len(jslClearances) == 0 {
		logger.Error("first request error", logger.Field("b", b))
		return
	}
	vm := otto.New()
	v, err := vm.Run(jslClearances[1])
	if err != nil {
		logger.Error(err.Error(), logger.Field("mes", "otto run js error"))
		return
	}
	cookiePro := strings.Split(strings.Split(v.String(), "=")[1], ";")[0]
	ck := req.GetCookie("__jsluid_s")
	req.Cookies.StrCookie = fmt.Sprintf("%s; __jsl_clearance_s=%s", ck, cookiePro)

	b, err = req.VisitString()
	if err != nil {
		return
	}
	reg = regexp.MustCompile(`;go\((.*?)\)`)
	data := reg.FindStringSubmatch(b)
	if len(data) == 0 {
		logger.Warn("getCookie error", logger.Field("b", b))
		return
	}
	c := _getjsluid(data[1])
	if c == "" {
		logger.Warn("getCookie error", logger.Field("data", data))
		return
	}
	cookie = fmt.Sprintf("%s; __jsl_clearance_s=%s", ck, c)
	return
}

func _getjsluid(ck string) string {
	var data map[string]interface{}
	err := json.Unmarshal([]byte(ck), &data)
	if err != nil {
		logger.Error(err.Error(), logger.Field("mes", "[]byte -> map err"))
		return ""
	}
	chars := fmt.Sprintf("%s", data["chars"].(string))
	charsLen := len(chars)
	for i := 0; i < charsLen; i++ {
		for j := 0; j < charsLen; j++ {
			clearance := fmt.Sprintf("%s%c%c%s",
				data["bts"].([]interface{})[0].(string), chars[i], chars[j], data["bts"].([]interface{})[1].(string))
			var encrypt hash.Hash
			switch data["ha"].(string) {
			case "md5":
				encrypt = md5.New()
			case "sha1":
				encrypt = sha1.New()
			default:
				encrypt = sha256.New()
			}
			encrypt.Write([]byte(clearance))
			s := hex.EncodeToString(encrypt.Sum(nil))
			if s == data["ct"].(string) {
				return clearance
			}
		}
	}
	return ""
}

func GetPageUrlList(url string, urlChan chan<- *callback.UrlChan, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= 100; i++ {
		urlChan <- &callback.UrlChan{
			Url:     fmt.Sprintf("%s%v", url, i),
			GetUrlF: GetDetailPageUrl,
		}

	}
}

func GetDetailPageUrl(url string, urlChan chan<- *callback.UrlChan, infoChan chan<- *callback.InfoChan) {
	req := request.Request{
		Url:    url,
		Method: http.MethodGet,
	}
	req.Cookies.StrCookie = _cookie
	b, err := req.Visit()
	if err != nil{
		return
	}
	var j store.JsonMiit
	err = json.Unmarshal(b, &j)
	if err != nil {
		logger.Error(err.Error(), logger.Field("url", url))
		return
	}
	for _, v := range j.DataMiit.DataResults {
		infoChan <- &callback.InfoChan{
			Url:      urlprocess.UrlJoint(store.BaseUrl, v.GroupData[0].Url),
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
			Source:        "工业和信息化部",
			SourceCode:    "WEB_00213",
			DateSelector:  "#con_time, .xxgk-span4",
			TitleSelector: "#con_title",
			TextSelector:  "#con_con p",
			DomainName:    "https://www.miit.gov.cn/",
		},
	}
	//fmt.Println(_cookie)
	pr.Request.Cookies.StrCookie = _cookie
	message <- pr.GetHtmlInfo()
}
