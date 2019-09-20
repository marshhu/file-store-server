package util

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io"
	"log"
)

var ossCli *oss.Client

const(
	OSSEndpoint string = "oss-cn-shenzhen.aliyuncs.com"
	OSSBucket string = "ma-image"
	OSSAccessKeyID string = "1111"
	OSSAccessKeySecret string = "1111"
)

// Client : 创建oss client对象
func Client() *oss.Client {
	if ossCli != nil {
		return ossCli
	}
	ossCli, err := oss.New("http://"+ OSSEndpoint,
		OSSAccessKeyID, OSSAccessKeySecret)
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
		bucket, err := cli.Bucket(OSSBucket)
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
	return  fmt.Sprintf("http://%s.%s/%s",OSSBucket,OSSEndpoint,objName)
}

func PutObjectToOSS(objectKey string,reader io.Reader)(string,error){
	bucket := Bucket()
	err := bucket.PutObject(objectKey, reader)
	if err != nil{
		return "",err
	}
	return GetPublicURL(objectKey),nil  //访问公共URL
}