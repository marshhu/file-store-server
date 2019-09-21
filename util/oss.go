package util

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/marshhu/file-store-server/conf"
	"io"
	"log"
)

var ossCli *oss.Client

// Client : 创建oss client对象
func Client() *oss.Client {
	if ossCli != nil {
		return ossCli
	}
	ossCli, err := oss.New("http://"+ conf.OssSetting.OSSEndpoint,
		conf.OssSetting.OSSAccessKeyID, conf.OssSetting.OSSAccessKeySecret)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return ossCli
}

// Bucket : 获取bucket存储空间
func Bucket() *oss.Bucket {
	cli := Client()
	if cli != nil {
		bucket, err := cli.Bucket(conf.OssSetting.OSSBucket)
		if err != nil {
			log.Fatal(err)
			return nil
		}
		return bucket
	}
	return nil
}

// GetSignURL : 获取临时签名访问URL
func GetSignURL(objName string) string {
	signedURL, err := Bucket().SignURL(objName, oss.HTTPGet, 3600)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	return signedURL
}

//获取公共访问URL
func GetPublicURL(objName string) string{
	return  fmt.Sprintf("http://%s.%s/%s",conf.OssSetting.OSSBucket,conf.OssSetting.OSSEndpoint,objName)
}

func PutObjectToOSS(objectKey string,reader io.Reader)(string,error){
	bucket := Bucket()
	err := bucket.PutObject(objectKey, reader)
	if err != nil{
		return "",err
	}
	return GetPublicURL(objectKey),nil  //访问公共URL
}