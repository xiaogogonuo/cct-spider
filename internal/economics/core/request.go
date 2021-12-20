package core

import (
	"crypto/tls"
	"github.com/xiaogogonuo/cct-spider/internal/economics/pkg/configReader"
	"io"
	"net/http"
)

const (
	UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.101 Safari/537.36"
)

type Sender interface {
	Send()
}

type Request struct {
	Page      int
	URL       string
	Method    string
	Cookie    []*http.Cookie
	Header    map[string]string
	Transport *http.Transport
	Parse     func(*Response)
	Meta      configReader.EconomicsConfig
}

func (s *Request) Get() {
	req, err := http.NewRequest(http.MethodGet, s.URL, nil)
	if err != nil {
		return
	}
	for _, c := range s.Cookie {
		req.AddCookie(c)
	}
	for k, v := range s.Header {
		req.Header.Add(k, v)
	}
	client := &http.Client{Transport: s.Transport}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	s.Parse(&Response{Page: s.Page, URL: s.URL, Body: body, Meta: s.Meta})
}

func (s *Request) Post() {
	return
}

func (s *Request) Send() {
	switch s.Method {
	case http.MethodGet:
		s.Get()
	case http.MethodPost:
		s.Post()
	}
}

func NewRequest(url string) *Request {
	return &Request{
		URL: url,
		Method: http.MethodGet,
		Header: map[string]string{
			"User-Agent": UserAgent,
		},
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
}