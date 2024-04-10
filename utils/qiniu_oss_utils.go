package utils

import (
	"bytes"
	"context"
	"github.com/google/uuid"
	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/storage"
	"log"
)

const (
	accessKey = "I-_nCAfs3VhrHfsJM3vk0JyaZS3OR6EEZfKUCGuV"
	secretKey = "d2YB50to5XqF4HJHpy0HBz_uEsmMjzrnWMBB_E_L"
	bucket    = "file-magic"
)

// 流文件上传

func OssFileUpload(stream []byte) (hash string, err error) {
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

	dataLen := int64(len(stream))

	err = formUploader.Put(context.Background(), &ret, upToken, uuid.New().String(), bytes.NewReader(stream), dataLen, &putExtra)
	if err != nil {
		log.Println("OssFileUpload Fail", err)
		return
	}
	log.Printf("文件名key:%s, 文件内容Hash:%s\n", ret.Key, ret.Hash)
	return ret.Hash, nil
}
