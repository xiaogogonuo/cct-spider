package downloader

import (
	"bytes"
	"io"
	"net/http"
)

// SimpleGet 最简单的GET请求
func SimpleGet(url string) (body []byte, err error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return
	}
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()
	return io.ReadAll(res.Body)
}

// Get 带请求头的GET请求
func Get(url string, header map[string]string) (body []byte, err error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return
	}
	for k, v := range header {
		req.Header.Add(k, v)
	}
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()
	return io.ReadAll(res.Body)
}

// Post 带请求体的POST方法
func Post(url string, reader []byte, header map[string]string) (body []byte, err error) {
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(reader))
	if err != nil {
		return
	}
	for k, v := range header {
		req.Header.Add(k, v)
	}
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}
