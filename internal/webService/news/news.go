package news

import (
	"bytes"
	"encoding/json"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/callback"
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
	"net/http"
)

var (
	url = "http://127.0.0.1:8080/post"
	method = http.MethodPost
)



func HandlerNews(info []callback.SqlValues) {
	m, _ := json.Marshal(info)

	req, err := http.NewRequest(method, url, bytes.NewReader(m))
	if err != nil{
		logger.Error(err.Error())
	}
	if err != nil {
		logger.Error(err.Error())
		return
	}
	req.Header.Add("Content-Type", "application/json")
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	defer resp.Body.Close()

}
