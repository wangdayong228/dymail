package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/gomail.v2"
)

func TestSendTencent(t *testing.T) {

	// args := os.Args //获取用户输入的所有参数
	// if args == nil {
	// 	return
	// }
	// message := args[0]
	message := "hello"
	m := gomail.NewMessage()
	m.SetHeader("From", "wangdayong230@163.com")
	m.SetHeader("To", "wangdayong228@163.com")
	m.SetHeader("Subject", "邮件主题")
	m.SetBody("text/html", message)
	//发送的附件
	//m.Attach("/tmp/sendEmail/FilterLog.2017-07-06.csv")

	// d := gomail.NewDialer("smtp.exmail.qq.com", 465, "wangdayong@conflux.fun", "77585243@Wang") // 发送邮件服务器、端口、发件人账号、发件人密码
	// if err := d.DialAndSend(m); err != nil {
	// 	log.Println("发送失败", err)
	// 	return
	// }

	d := gomail.NewDialer("smtp.163.com", 25, "wangdayong230@163.com", "IOUYPEAXZAJDQHDK") // 发送邮件服务器、端口、发件人账号、发件人密码
	err := d.DialAndSend(m)
	assert.NoError(t, err)
}
