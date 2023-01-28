package cct_index

import (
	"encoding/json"
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/constant"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/poster"
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
	"os/exec"
	"time"
)

type CurlResponse struct {
	TimeStamp string `json:"timestamp"`
	Status    int    `json:"status"`
	Error     string `json:"error"`
	Message   string `json:"message"`
	Path      string `json:"path"`
}

func RemoteConnect() bool {
	cmd := exec.Command("curl", constant.Service)
	out, err := cmd.Output()
	if err != nil {
		logger.Error(err.Error())
		return false
	}

	var cr CurlResponse
	if err = json.Unmarshal(out, &cr); err != nil {
		logger.Error(err.Error())
		return false
	}

	switch cr.Status {
	case 404, 500:
		return false
	}

	return true
}

// RemoteListen 每隔5分钟检测一次远程服务器接口是否健康
// 如果远程服务器接口异常，则下发邮件通知并停止应用，待远程服务器接口正常后重新启动爬虫服务
func RemoteListen(listener chan struct{}) {
	for {
		if !RemoteConnect() {
			// 下发邮件
			body := fmt.Sprintf("指标对应的远程接口连接失败：%s", constant.Service)
			if err := poster.E163.Send(poster.Receivers, poster.Subject, body); err != nil {
				logger.Error(err.Error())
			}

			listener <- struct{}{}
			return
		}
		time.Sleep(ListenInterval)
	}
}
