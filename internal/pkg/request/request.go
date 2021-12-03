package request

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/codec"
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type Request struct {
	Url     string
	Method  string
	Timeout time.Duration
	Body    io.Reader
	Param   map[string]string
	Header  map[string]string
	Cookies struct {
		StrCookie  string
		HttpCookie []*http.Cookie
	}
}

func (r *Request) Visit() (b []byte, err error) {
	resp, err := r.request()
	if err != nil {
		return
	}
	defer resp.Body.Close()
	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error(err.Error(), logger.Field("url", r.Url))
		return
	}
	if codec.IsGbk(b) {
		b, err = codec.GbkToUtf8(b)
		if err != nil {
			logger.Error(err.Error(), logger.Field("url", r.Url))
			return
		}
	}
	return
}

func (r *Request) VisitString() (s string, err error) {
	resp, err := r.request()
	if err != nil {
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error(err.Error(), logger.Field("url", r.Url))
		return
	}
	if codec.IsGbk(b) {
		b, err = codec.GbkToUtf8(b)
		if err != nil {
			logger.Error(err.Error(), logger.Field("url", r.Url))
			return
		}
	}
	s = string(b)
	return
}

func (r *Request) GetCookie(name string) (ck string) {
	var cks []string
	for _, cookie := range r.Cookies.HttpCookie {
		if cookie.Name == name {
			cks = append(cks, cookie.String())
		}
	}
	ck = strings.Join(cks, "")
	return
}

func (r *Request) request() (resp *http.Response, err error) {
	req, err := http.NewRequest(r.Method, r.Url, r.Body)
	if err != nil {
		logger.Error(err.Error(), logger.Field("url", r.Url))
		return
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.106 Safari/537.36")
	q := req.URL.Query()
	for key, val := range r.Param {
		q.Add(key, val)
	}

	req.URL.RawQuery = q.Encode()
	for k, v := range r.Header {
		req.Header.Set(k, v)
	}
	if r.Cookies.StrCookie != ""{
		req.Header.Set("Cookie", r.Cookies.StrCookie)
	}
	client := &http.Client{Timeout: r.Timeout}

	resp, err = client.Do(req)
	if err != nil {
		logger.Error(err.Error(), logger.Field("url", r.Url))
		return
	}
	if resp.StatusCode != http.StatusOK && resp.StatusCode != 521 {
		logger.Warn(fmt.Sprintf("StatusCode: %v, Url: %s", resp.StatusCode, r.Url))
		return
	}
	r.Cookies.HttpCookie = resp.Cookies()
	return
}
