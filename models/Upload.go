package models

import (
	"context"
	"fmt"
	"ginblog/utils"
	"ginblog/utils/errmsg"
	"github.com/tencentyun/cos-go-sdk-v5"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
)

var BucketName = utils.BucketName
var SecretKey = utils.SecretKey
var SecretID = utils.SecretID
var Region = utils.Region
var client *cos.Client

func initClient() {
	if client == nil {
		u, _ := url.Parse(fmt.Sprintf("https://%s.cos.%s.myqcloud.com", BucketName, Region))
		b := &cos.BaseURL{BucketURL: u}
		client = cos.NewClient(b, &http.Client{
			Transport: &cos.AuthorizationTransport{
				SecretID:  SecretID,
				SecretKey: SecretKey,
			},
		})
	}
}

// 七牛云UpLoadFile 上传文件函数
//func UpLoadFile(file multipart.File, fileSize int64) (string, int) {
//	putPolicy := storage.PutPolicy{
//		Scope: Bucket,
//	}
//	mac := qbox.NewMac(AccessKey, SecretKey)
//	upToken := putPolicy.UploadToken(mac)
//
//	cfg := setConfig()
//
//	putExtra := storage.PutExtra{}
//
//	formUploader := storage.NewFormUploader(&cfg)
//	ret := storage.PutRet{}
//
//	err := formUploader.PutWithoutKey(context.Background(), &ret, upToken, file, fileSize, &putExtra)
//	if err != nil {
//		return "", errmsg.ERROR
//	}
//	url := ImgUrl + ret.Key
//	return url, errmsg.SUCCESS
//}
//
//func setConfig() storage.Config {
//	cfg := storage.Config{
//		Zone:          selectZone(Zone),
//		UseCdnDomains: false,
//		UseHTTPS:      false,
//	}
//	return cfg
//}
//
//func selectZone(id int) *storage.Zone {
//	switch id {
//	case 1:
//		return &storage.ZoneHuadong
//	case 2:
//		return &storage.ZoneHuabei
//	case 3:
//		return &storage.ZoneHuanan
//	default:
//		return &storage.ZoneHuadong
//	}
//}

// 腾讯云上传文件函数
func UploadFile(file *multipart.File, fileHeader *multipart.FileHeader) (string, int) {
	initClient()
	res, err := client.Object.Put(context.Background(), "/ginblog/"+fileHeader.Filename, *file, nil)
	if err != nil || res.StatusCode != http.StatusOK {
		return "", errmsg.ERROR
	}
	//返回一个文件名
	return "http://img.centyoubet.xyz/" + fileHeader.Filename, errmsg.SUCCESS
}

// 从腾讯云服务器下载文件
func DownloadFile(fileName string) ([]byte, int) {
	initClient()
	res, err := client.Object.Get(context.Background(), "/ginblog/"+fileName, nil)
	if err != nil || res.StatusCode != http.StatusOK {
		return nil, errmsg.ERROR
	}
	fileByte, _ := io.ReadAll(res.Body)
	defer res.Body.Close()
	fmt.Println("获取到的文件数据", string(fileByte))
	//将文件数据写入http的响应头
	//gin.Context{}.Request.Response.Write(bytes.NewWriter(fileByte))
	return fileByte, errmsg.SUCCESS
}
