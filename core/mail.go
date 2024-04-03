package core

import (
	"gopkg.in/gomail.v2"
)

func SendToMail(host string, port int, to, user, password string, subject, body string) error {
	// message := "hello"
	m := gomail.NewMessage()
	m.SetHeader("From", user)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	//发送的附件
	//m.Attach("/tmp/sendEmail/FilterLog.2017-07-06.csv")

	// d := gomail.NewDialer("smtp.exmail.qq.com", 465, "wangdayong@conflux.fun", "77585243@Wang") // 发送邮件服务器、端口、发件人账号、发件人密码
	// if err := d.DialAndSend(m); err != nil {
	// 	log.Println("发送失败", err)
	// 	return
	// }

	d := gomail.NewDialer(host, port, user, password) // 发送邮件服务器、端口、发件人账号、发件人密码
	return d.DialAndSend(m)
}
