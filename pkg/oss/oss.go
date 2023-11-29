package oss

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
	"gohub/pkg/config"
	"gohub/pkg/logger"
	"gohub/pkg/response"
	"os"
	"sync"
	"time"
)

type BUCKET struct {
	Bucket *oss.Bucket
}

// 定义进度条监听器。
type OssProgressListener struct {
}

// once单例模式
var once sync.Once

// 内部使用的单例
var internalBUCKET *BUCKET

// 单例
func NewBUCKET() *BUCKET {
	once.Do(func() {
		// 创建OSSClient实例。
		client, err := oss.New(config.Get("oss.oss_aliyun_endpoint"), config.Get("oss.oss_aliyun_accesskeyid"), config.Get("oss.oss_aliyun_accesskeysecret"))
		if err != nil {
			logger.ErrorJSON("oss", "创建client", err)
			os.Exit(-1)
		}
		// 获取存储空间。
		bucket, err := client.Bucket(config.Get("oss.oss_aliyun_bucketname"))
		if err != nil {
			logger.ErrorJSON("oss", "创建Bucket", err)
			os.Exit(-1)
		}

		internalBUCKET = &BUCKET{
			Bucket: bucket,
		}
	})
	return internalBUCKET
}

func (b *BUCKET) UploadFile(fileName, filePath string, c *gin.Context) (string, error) {

	//云路径
	yunFileTmpPath := "gohub/uploads/" + time.Now().Format("2006-01-02") + "/" + fileName
	// 上传文件
	err := b.Bucket.PutObjectFromFile(yunFileTmpPath, filePath, oss.Progress(&OssProgressListener{}))
	if err != nil {
		response.Abort500(c)
		os.Exit(-1)
		return "", nil
	}
	return yunFileTmpPath, nil
}

// 定义进度变更事件处理函数。
func (listener *OssProgressListener) ProgressChanged(event *oss.ProgressEvent) {
	switch event.EventType {
	case oss.TransferStartedEvent:
		fmt.Printf("上传进度 Started, ConsumedBytes: %d, TotalBytes %d.\n",
			event.ConsumedBytes, event.TotalBytes)
	case oss.TransferDataEvent:
		fmt.Printf("\r上传进度 Data, ConsumedBytes: %d, TotalBytes %d, %d%%.",
			event.ConsumedBytes, event.TotalBytes, event.ConsumedBytes*100/event.TotalBytes)
	case oss.TransferCompletedEvent:
		fmt.Printf("\n上传进度 Completed, ConsumedBytes: %d, TotalBytes %d.\n",
			event.ConsumedBytes, event.TotalBytes)
	case oss.TransferFailedEvent:
		fmt.Printf("\n上传进度 Failed, ConsumedBytes: %d, TotalBytes %d.\n",
			event.ConsumedBytes, event.TotalBytes)
	default:
	}
}
