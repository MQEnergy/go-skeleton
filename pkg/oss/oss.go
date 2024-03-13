package oss

import (
	"bytes"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type Oss struct {
	client *oss.Client
	bucket *oss.Bucket
}

type Config struct {
	EndPoint     string
	AccessId     string
	AccessSecret string
	BucketName   string
}

func New(config *Config, options ...oss.ClientOption) (*Oss, error) {
	client, err := oss.New(config.EndPoint, config.AccessId, config.AccessSecret, options...)
	if err != nil {
		return nil, err
	}
	bucket, err := client.Bucket(config.BucketName)
	if err != nil {
		return nil, err
	}
	return &Oss{
		client: client,
		bucket: bucket,
	}, nil
}

// PutObject 上传字符串
func (o *Oss) PutObject(object string, content []byte) error {
	return o.bucket.PutObject(object, bytes.NewReader(content))
}

// PutObjectFromFile 上传文件
func (o *Oss) PutObjectFromFile(object, local string) error {
	return o.bucket.PutObjectFromFile(object, local)
}
