package test

import (
	"cherry-disk/core/define"
	"context"
	"net/http"
	"net/url"
	"testing"

	"github.com/tencentyun/cos-go-sdk-v5"
)

func TestFileUploadByFilePath(t *testing.T) {
	u, _ := url.Parse("https://cherryneko-1312558494.cos.ap-nanjing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.TencentSecretID,
			SecretKey: define.TencentSecretKEY,
		},
	})

	key := "cherry-disk/neko.jpg"

	_, _, err := client.Object.Upload(
		context.Background(), key, "./img/neko.jpg", nil,
	)
	if err != nil {
		panic(err)
	}
}
