package helper

import (
	"bytes"
	"cherry-disk/core/common"
	"cherry-disk/core/define"
	"context"
	"crypto/md5"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/smtp"
	"net/url"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/tencentyun/cos-go-sdk-v5"

	"github.com/dgrijalva/jwt-go"
	"github.com/jordan-wright/email"
	uuid "github.com/satori/go.uuid"
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

func MailSendCode(mail, code string) error {
	e := email.NewEmail()
	e.From = "Cherry Disk <iamyunjielai@163.com>"
	e.To = []string{mail}
	e.Subject = "Cherry Disk 验证码发送测试"
	e.HTML = []byte("你的验证码为：<h1>" + code + "</h1>")
	err := e.SendWithTLS("smtp.163.com:465", smtp.PlainAuth("", "iamyunjielai@163.com", common.Cfg.MailPwd, "smtp.163.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})
	if err != nil {
		return err
	}
	return nil
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

// 上传文件-腾讯云
func UploadTX(r *http.Request) (string, error) {
	u, _ := url.Parse(define.CosBucket)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  common.Cfg.Bucket.TencentSecretId,
			SecretKey: common.Cfg.Bucket.TencentSecretKey,
		},
	})

	file, fileHeader, err := r.FormFile("file")
	key := "cherry-disk/" + UUID() + path.Ext(fileHeader.Filename)

	_, err = client.Object.Put(
		context.Background(), key, file, nil,
	)

	if err != nil {
		panic(err)
	}

	return define.CosBucket + "/" + key, nil
}

func CosPartUpload(r *http.Request) (string, error) {
	bk, err := url.Parse(define.CosBucket)
	if err != nil {
		return "", err
	}
	b := &cos.BaseURL{BucketURL: bk}
	client := cos.NewClient(b, &http.Client{Transport: &cos.AuthorizationTransport{
		SecretID:  common.Cfg.Bucket.TencentSecretId,
		SecretKey: common.Cfg.Bucket.TencentSecretKey,
	}})

	key := r.PostForm.Get("key")
	uploadID := r.PostForm.Get("upload_id")
	partNum, err := strconv.Atoi(r.PostForm.Get("part_num"))
	if err != nil {
		return "", err
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		return "", err
	}

	buf := bytes.NewBuffer(nil)
	io.Copy(buf, file)

	resp, err := client.Object.UploadPart(
		context.Background(), key, uploadID, partNum, bytes.NewReader(buf.Bytes()), nil,
	)
	if err != nil {
		return "", err
	}

	return strings.Trim(resp.Header.Get("ETag"), "\""), nil
}

func CosPartUploadComplete(key, uploadId string, co []cos.Object) error {
	u, _ := url.Parse(define.CosBucket)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  common.Cfg.Bucket.TencentSecretId,
			SecretKey: common.Cfg.Bucket.TencentSecretKey,
		},
	})

	ops := &cos.CompleteMultipartUploadOptions{}
	ops.Parts = append(ops.Parts, co...)
	_, _, err := client.Object.CompleteMultipartUpload(
		context.Background(), key, uploadId, ops,
	)
	return err
}
