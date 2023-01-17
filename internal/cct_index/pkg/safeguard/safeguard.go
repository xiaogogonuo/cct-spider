package safeguard

import "net/http"

const (
	target = "https://www.baidu.com"
	failed = 5
)

func IsNetworkNormal() bool {
	c := 0
	for {
		if c >= failed {
			return false
		}
		res, err := http.Get(target)
		if err != nil {
			c++
			continue
		}
		if res.StatusCode == 200 {
			return true
		}
		c++
	}
}
