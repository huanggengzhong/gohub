package upload

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"mime/multipart"
	"os"
	"time"
)

type FileUploadAndDownloadService struct{}

func UploadFile(header *multipart.FileHeader) {

	// 创建OSSClient实例。
	client, err := oss.New("oss-cn-shenzhen.aliyuncs.com", "key", "hah")
	if err != nil {
		fmt.Println("Error2:", err)
		os.Exit(-1)
	}
	// 获取存储空间。
	bucket, err := client.Bucket("sliver-com")
	if err != nil {
		fmt.Println("Error3:", err)
		os.Exit(-1)
	}
	//本地文件

	//云路径
	yunFileTmpPath := "gohub/uploads/" + time.Now().Format("2006-01-02") + "/" + "499550.jpeg"
	fd, err := os.Open("/Users/luogengzhong/Documents/study/02-GOPROJECT/gohub/pkg/upload/499550.jpeg")
	if err != nil {
		fmt.Println("Error4:", err)
		os.Exit(-1)
	}
	defer fd.Close()
	// 上传文件流。
	err = bucket.PutObject(yunFileTmpPath, fd)
	if err != nil {
		fmt.Println("Error5:", err)
		os.Exit(-1)
	}
}
