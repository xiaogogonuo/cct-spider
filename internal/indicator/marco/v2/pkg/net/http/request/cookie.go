package request

import (
	"crypto/tls"
	"fmt"
	"net/http"
)

const CookieSource = "https://data.stats.gov.cn/easyquery.htm?cn=%s&zb=%s"

// GetCookie return indicator cookie
// cn: date type, zb: indicator
func GetCookie(cn, zb string) (c []*http.Cookie, err error) {
	url := fmt.Sprintf(CookieSource, cn, zb)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return
	}
	req.Header.Add("User-Agent", UserAgent)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	for _, cookie := range resp.Cookies() {
		c = append(c, &http.Cookie{
			Name:  cookie.Name,
			Value: cookie.Value,
		})
	}
	return
}