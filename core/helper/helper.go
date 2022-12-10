package helper

import (
	"cherry-disk/core/define"
	"crypto/md5"
	"fmt"
	"github.com/jordan-wright/email"
	uuid "github.com/satori/go.uuid"
	"math/rand"
	"time"
)

func Md5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

func UUID() string {
	return uuid.NewV4().String()
}

func RandCode() string {
	s := "1234567890"
	code := ""
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < define.CodeLength; i++ {
		code += string(s[rand.Intn(len(s))])
	}
	return code
}

func MailSendCode(mail, code string) {
	e := email.NewEmail()
	e.From = "Get <CherryNeko@edj.com>"
	e.To = []string{mail}
	e.Subject = "验证码发送测试"
	e.HTML = []byte("你的验证码为： <h1>" + code + "</h1>")
	//e.SendWithTLS("stmp.edj.com:666", smtp.PlainAuth(""))
}
