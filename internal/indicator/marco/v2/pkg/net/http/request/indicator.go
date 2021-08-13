package request

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
)

// GetIndicator return indicator data
func GetIndicator(url string, cookie []*http.Cookie) (b []byte, err error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return
	}
	req.Header.Add("User-Agent", UserAgent)
	for _, c := range cookie {
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