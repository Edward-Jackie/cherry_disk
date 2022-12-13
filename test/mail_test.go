package test

import (
	"cherry-disk/core/define"
	"crypto/tls"
	"net/smtp"
	"testing"

	"github.com/jordan-wright/email"
)

func TestSendMail(t *testing.T) {
	e := email.NewEmail()
	e.From = "Cherry Disk <iamyunjielai@163.com>"
	e.To = []string{"1938941711@qq.com"}
	e.Subject = "Cherry Disk 验证码发送测试"
	e.HTML = []byte("你的验证码为：<h1>123456</h1>")
	err := e.SendWithTLS("smtp.163.com:465", smtp.PlainAuth("", "iamyunjielai@163.com", define.MailPWD, "smtp.163.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})
	if err != nil {
		t.Fatal(err)
	}
}
