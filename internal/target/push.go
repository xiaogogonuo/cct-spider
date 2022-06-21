package target

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/xiaogogonuo/cct-spider/internal/target/model"
	"io"
	"net/http"
)

// WebService 远程数据库服务器接口
const WebService = "http://106.37.165.121/inf/chengtong/py/sy/baseTargetValue/saveRequest"

type RemoteServerResponse struct {
	Flag     bool   `json:"flag"`
	Status   bool   `json:"status"`
	TipsCode string `json:"tipsCode"`
	Msg      string `json:"msg"`
	PkID     string `json:"pkId"`
}

// push 将新数据推送到数据库服务器
func push(newTargets []model.DataBase) []string {
	var targetForUpdate []string
	for _, tar := range newTargets {
		if err := sendAnyData(tar); err == nil {
			targetForUpdate = append(targetForUpdate, tar.ValueGUID)
		}
	}
	return targetForUpdate
}

func sendAnyData(data ...model.DataBase) error {
	payload := map[string][]model.DataBase{"data": data}
	m, _ := json.Marshal(payload)
	req, err := http.NewRequest(http.MethodPost, WebService, bytes.NewReader(m))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var remote RemoteServerResponse
	if err := json.Unmarshal(b, &remote); err != nil {
		return err
	}
	if !remote.Status {
		return errors.New("insert error")
	}
	return nil
}
