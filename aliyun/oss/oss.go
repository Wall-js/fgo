package oss

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io"
	"io/ioutil"
	"os"
)

func handleError(err error) {
	fmt.Println("Error:", err)
	os.Exit(-1)
}

type Client struct {
	endpoint string
	// 阿里云主账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM账号进行API访问或日常运维，请登录 https://ram.console.aliyun.com 创建RAM账号。
	accessKeyId     string
	accessKeySecret string
	bucketName      string
	objectName      string
}

func (obj *Client) PutObject(bucketName string, objectName string, reader io.Reader) { // 定义Student方法
	client, err := oss.New(obj.endpoint, obj.accessKeyId, obj.accessKeySecret)
	if err != nil {
		handleError(err)
	}
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		handleError(err)
	}
	// 上传文件。
	err = bucket.PutObject(objectName, reader)
	if err != nil {
		handleError(err)
	}
}

func (obj *Client) GetObject(bucketName string, objectName string) []byte { // 定义Student方法
	client, err := oss.New(obj.endpoint, obj.accessKeyId, obj.accessKeySecret)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	// 获取存储空间。
	bucket, err := client.Bucket("<yourBucketName>")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	// 下载文件到流。
	body, err := bucket.GetObject("<yourObjectName>")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	// 数据读取完成后，获取的流必须关闭，否则会造成连接泄漏，导致请求无连接可用，程序无法正常工作。
	defer body.Close()

	data, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	return data
}
