package remote

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/constant"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/model"
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
	"io"
	"net/http"
)

type ServerResponse struct {
	Flag     bool   `json:"flag"`
	Status   bool   `json:"status"`
	TipsCode string `json:"tipsCode"`
	Msg      string `json:"msg"`
	PkID     string `json:"pkId"`
}

// Push 由于并发发送数据服务器无法支撑，因此一条一条的往服务器推送数据
func Push(indexes []*model.Index) (newIndexes []string) {
	for _, index := range indexes {
		if err := send(*index); err == nil {
			newIndexes = append(newIndexes, index.ValueGUID)
		}
	}
	return
}

func send(index ...model.Index) (err error) {
	payload := map[string][]model.Index{"data": index}
	m, _ := json.Marshal(payload)
	req, _ := http.NewRequest(http.MethodPost, constant.Service, bytes.NewReader(m))
	req.Header.Add("Content-Type", "application/json")
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	var sr ServerResponse
	if err := json.Unmarshal(b, &sr); err != nil {
		logger.Error(err.Error())
		return err
	}
	if !sr.Status {
		return errors.New("insert fail")
	}
	return
}
