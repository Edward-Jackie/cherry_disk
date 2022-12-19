package test

import (
	"cherry-disk/core/common"
	"context"
	"math"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"testing"

	"github.com/tencentyun/cos-go-sdk-v5"
)

// 文件上传测试
func TestFileUploadByFilePath(t *testing.T) {
	common.GetConfig()
	u, _ := url.Parse("https://cherryneko-1312558494.cos.ap-nanjing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  common.Cfg.Bucket.TencentSecretId,
			SecretKey: common.Cfg.Bucket.TencentSecretKey,
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

var chunkSize = 10 * 1024 * 1024 // 10MB

// 文件分片
func TestGeneralChunkFile(t *testing.T) {
	file, err := os.OpenFile("./mv/test.mp4", os.O_RDONLY, 0666)
	if err != nil {
		t.Fatal(err)
	}

	fileInfo, err := os.Stat("./mv/test.mp4")
	if err != nil {
		t.Fatal(err)
	}

	//分片的个数
	chunkNum := math.Ceil(float64(fileInfo.Size()) / float64(chunkSize))
	b := make([]byte, chunkSize)
	for i := 0; i < int(chunkNum); i++ {
		file.Seek(int64(i*chunkSize), 0)
		if int64(chunkSize) > fileInfo.Size()-int64(i*chunkSize) {
			b = make([]byte, fileInfo.Size()-int64(i*chunkSize))
		}

		file.Read(b)
		f, err := os.OpenFile("./chunks/"+strconv.Itoa(i)+".chunk", os.O_CREATE|os.O_WRONLY, os.ModePerm)
		if err != nil {
			t.Fatal(err)
		}

		f.Write(b)
		f.Close()
	}
	file.Close()
}

// 文件合并 - 校验一致性
func TestChunkFile(t *testing.T) {
	//file, err := os.OpenFile("test.mp4", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	//if err != nil {
	//	t.Fatal(err)
	//}
}
