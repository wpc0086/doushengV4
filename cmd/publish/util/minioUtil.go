package util

import (
	"context"
	"doushengV4/pkg/consts"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
)

var (
	Client *minio.Client
	err    error
)

func InitMinio() {
	Client, err = minio.New(consts.EndPoint, &minio.Options{
		Creds:  credentials.NewStaticV4(consts.AccessKeyID, consts.SecretAccessKey, ""),
		Secure: consts.UseSSL})
	if err != nil {
		log.Fatalln("minio连接错误: ", err)
	}
}

func CreateBucket(ctx context.Context) error {
	exists, _ := Client.BucketExists(ctx, consts.BucketName)
	if exists != true {
		err := Client.MakeBucket(ctx, consts.BucketName, minio.MakeBucketOptions{Region: "cn-south-1", ObjectLocking: false})
		if err != nil {
			klog.Infof("创建bucket错误: ", err)
		}
		return err
	}
	return nil
}

func ImageToMinio(c context.Context, objectName string, filePath string) error {
	//创建桶
	err := CreateBucket(c)
	if err != nil {
		return err
	}
	_, err = Client.FPutObject(c, consts.BucketName, objectName, filePath, minio.PutObjectOptions{ContentType: consts.ImageContentType})
	if err != nil {
		klog.Infof("上传失败：", err)
		return err
	}
	return nil
}

func VideoToMinio(c context.Context, objectName string, filePath string) error {
	//创建桶
	err = CreateBucket(c)
	if err != nil {
		return err
	}
	//将文件打开
	_, err = Client.FPutObject(c, consts.BucketName, objectName, filePath, minio.PutObjectOptions{ContentType: consts.VideoContentType})
	if err != nil {
		klog.Infof("上传失败：", err)
		return err
	}
	return nil
}
