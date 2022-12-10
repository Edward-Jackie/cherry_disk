package helper

import (
	"cherry-disk/core/define"
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
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

// Token生成方法
func TokenBuilder(id int, identity, name string, second int) (string, error) {
	uc := define.UserClaim{
		Id:       id,
		Identity: identity,
		Name:     name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * time.Duration(second)).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	tokenstr, err := token.SignedString([]byte(define.JwtKey))
	if err != nil {
		return "", err
	}
	return tokenstr, nil
}

// Token解析
func AnalyzeToken(token string) (*define.UserClaim, error) {
	u := new(define.UserClaim)
	claims, err := jwt.ParseWithClaims(token, u, func(token *jwt.Token) (interface{}, error) {
		return []byte(define.JwtKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !claims.Valid {
		return u, errors.New("Token is unavaliable")
	}
	return u, err
}
