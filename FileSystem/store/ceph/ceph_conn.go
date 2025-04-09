package ceph

import (
	"gopkg.in/amz.v1/aws"
	"gopkg.in/amz.v1/s3"
)

var cephConn *s3.S3

// GetCephConnection : 获取ceph连接
func GetCephConnection() *s3.S3 {
	if cephConn != nil {
		return cephConn
	}
	//初始化ceph的一些信息
	auth := aws.Auth{
		AccessKey: "YCU14JD9KCHYIZXISGPT",
		SecretKey: "iDEzHFaaAsspIOAqttRho0EH8mJUYmnCkoqRLbzj",
	}

	curRegion := aws.Region{
		Name:                 "default",
		EC2Endpoint:          "http://49.232.28.223:7480",
		S3Endpoint:           "http://49.232.28.223:7480",
		S3BucketEndpoint:     "",
		S3LocationConstraint: false,
		S3LowercaseBucket:    false,
		Sign:                 aws.SignV2,
	}
	//创建S3类型的连接
	return s3.New(auth, curRegion)
}

// GetCephBucket : 获取指定的bucket对象
func GetCephBucket(bucket string) *s3.Bucket {
	conn := GetCephConnection()
	return conn.Bucket(bucket)
}
