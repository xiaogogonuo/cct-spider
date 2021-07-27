package request

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
)

const ua = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.101 Safari/537.36"

type Request struct {
	URL    string
	Cookie []*http.Cookie
}

func (r *Request) Visit() (b []byte, err error) {
	req, err := http.NewRequest(http.MethodGet, r.URL, nil)
	if err != nil {
		return
	}
	req.Header.Add("User-Agent", ua)
	for _, c := range r.Cookie {
		req.AddCookie(c)
	}
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: transport}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	b, err = ioutil.ReadAll(resp.Body)
	return
}
