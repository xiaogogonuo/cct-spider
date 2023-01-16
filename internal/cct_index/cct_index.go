package cct_index

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
	"github.com/xiaogogonuo/cct-spider/pkg/mail"
	"time"
)

func spider() {
	fmt.Println("spider ing")
	time.Sleep(time.Second)
}

func RemoteConnect() bool {
	fmt.Println("try to connect remote server")
	return true
}

// RemoteListen 每隔5秒检测一次远程服务器接口是否健康
// 如果远程服务器接口异常，则下发邮件通知并停止应用，待远程服务器接口正常后重新启动
func RemoteListen(listener chan struct{}) {
	for {
		if !RemoteConnect() {
			// 下发邮件
			e := mail.NewEmail163("xiaogogonuo@163.com", "JDAOREDDCXYMAXXQ", "lujiawei")
			receivers := []string{"xiaogogonuo@163.com"}
			subject := "城通爬虫系统警报"
			body := `
			指标对应的远程接口连接失败：http://106.37.165.121/inf/chengtong/py/sy/baseTargetValue/saveRequest
			`
			if err := e.Send(receivers, subject, body); err != nil {
				logger.Error(err.Error())
			}

			listener <- struct{}{}
			return
		}
		time.Sleep(time.Second * 5)
	}
}

func RunApplication() {
	listener := make(chan struct{})
	go RemoteListen(listener)
	for {
		select {
		case <-listener:
			return
		default:
			spider()
		}
	}
}
