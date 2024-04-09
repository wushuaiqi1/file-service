package utils

import (
	"bytes"
	"context"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/storage"
	"log"
	"math/rand"
	"strconv"
	"time"
)

const (
	accessKey = "I-_nCAfs3VhrHfsJM3vk0JyaZS3OR6EEZfKUCGuV"
	secretKey = "d2YB50to5XqF4HJHpy0HBz_uEsmMjzrnWMBB_E_L"
	bucket    = "file-magic"
)

func Test() {
	mac := auth.New(accessKey, secretKey)
	localFile := "//Users/tal/Desktop/sami.png"
	//key := "github-x.png"

	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{
		Region:        &storage.ZoneHuabei,
		UseHTTPS:      false,
		UseCdnDomains: false,
	}

	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	// 可选配置
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo",
		},
	}
	err := formUploader.PutFile(context.Background(), &ret, upToken, strconv.FormatInt(rand.Int63(), 10), localFile, &putExtra)
	fmt.Println("11")
	if err != nil {
		fmt.Println(err)
		return
	}
	//文件名key 文件内容Hash
	fmt.Println(ret.Key, ret.Hash)
}

// 流文件上传

func FileUpload(stream []byte) (hash string, err error) {
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := auth.New(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{
		Region:        &storage.ZoneHuabei,
		UseHTTPS:      false,
		UseCdnDomains: false,
	}

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo 1 byte",
		},
	}

	//data := []byte("hello, this is qiniu cloud")
	dataLen := int64(len(stream))

	err = formUploader.Put(context.Background(), &ret, upToken, time.Now().String(), bytes.NewReader(stream), dataLen, &putExtra)
	if err != nil {
		log.Println("ssss ", err)
		return
	}
	fmt.Println(ret.Key, ret.Hash)
	return hash, nil
}
