package mail

import (
	"gopkg.in/gomail.v2"
)

// Email163 网易163邮箱
type Email163 struct {
	Host string // 邮箱服务器地址
	Port int    // 邮箱服务器端口
	User string // 邮箱发送者名字
	Pass string // 邮箱发送者密码
	Nick string // 邮箱发送者昵称
}

// Send receivers不能为空
func (e *Email163) Send(receivers []string, subject string, body string) error {
	m := gomail.NewMessage()
	// 设置发送者
	m.SetHeader("From", e.Nick+"<"+e.User+">")
	// 设置接受者
	m.SetHeader("To", receivers...)
	// 设置主题
	m.SetHeader("Subject", subject)
	// 设置正文
	m.SetBody("text/html", body)
	d := gomail.NewDialer(e.Host, e.Port, e.User, e.Pass)
	return d.DialAndSend(m)
}

func NewEmail163(user, pass, nick string) *Email163 {
	return &Email163{
		Host: "smtp.163.com",
		Port: 25,
		User: user,
		Pass: pass,
		Nick: nick,
	}
}
