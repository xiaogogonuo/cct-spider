package cook

import (
	"crypto/tls"
	"fmt"
	"net/http"
)

// Cookie
// "https://data.stats.gov.cn/easyquery.htm?cn=C01&zb=A020102"
// cn代表数据类型，C01代表年度数据
// zb代表指标，A020102是GDP
func Cookie(cn, zb string) (c []*http.Cookie) {
	url := fmt.Sprintf("https://data.stats.gov.cn/easyquery.htm?cn=%s&zb=%s", cn, zb)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get(url)
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
