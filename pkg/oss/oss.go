package oss

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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
func (b *BUCKET) UploadPart(fileName, filePath string, c *gin.Context) (string, error) {

	// 填写Object完整路径。Object完整路径中不能包含Bucket名称。
	objectName := "gohub/uploads/" + time.Now().Format("2006-01-02") + "/" + fileName
	// 填写本地文件的完整路径。如果未指定本地路径，则默认从示例程序所属项目对应本地路径中上传文件。
	locaFilename := filePath

	// 将本地文件分片，且分片数量指定为3。
	chunks, err := oss.SplitFileByPartNum(locaFilename, 3)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	fd, err := os.Open(locaFilename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	defer fd.Close()

	// 指定过期时间。
	expires := time.Date(2049, time.January, 10, 23, 0, 0, 0, time.UTC)
	// 如果需要在初始化分片时设置请求头，请参考以下示例代码。
	options := []oss.Option{
		oss.MetadataDirective(oss.MetaReplace),
		oss.Expires(expires),
		// 指定该Object被下载时的网页缓存行为。
		// oss.CacheControl("no-cache"),
		// 指定该Object被下载时的名称。
		// oss.ContentDisposition("attachment;filename=FileName.txt"),
		// 指定该Object的内容编码格式。
		// oss.ContentEncoding("gzip"),
		// 指定对返回的Key进行编码，目前支持URL编码。
		// oss.EncodingType("url"),
		// 指定Object的存储类型。
		// oss.ObjectStorageClass(oss.StorageStandard),
	}

	// 步骤1：初始化一个分片上传事件。
	imur, err := b.Bucket.InitiateMultipartUpload(objectName, options...)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	// 步骤2：上传分片。
	var parts []oss.UploadPart
	for _, chunk := range chunks {
		fd.Seek(chunk.Offset, os.SEEK_SET)
		// 调用UploadPart方法上传每个分片。
		part, err := b.Bucket.UploadPart(imur, fd, chunk.Size, chunk.Number)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(-1)
		}
		parts = append(parts, part)
	}

	// 指定Object的读写权限为私有，默认为继承Bucket的读写权限。
	objectAcl := oss.ObjectACL(oss.ACLPublicReadWrite)
	//objectAcl := oss.ObjectACL(oss.ACLPrivate)

	// 步骤3：完成分片上传。
	_, err = b.Bucket.CompleteMultipartUpload(imur, parts, objectAcl)
	//cmur, err := b.Bucket.CompleteMultipartUpload(imur, parts, objectAcl)
	if err != nil {
		response.Abort500(c)
		os.Exit(-1)
		return "", err

	}
	//fmt.Println("cmur:", cmur)
	return objectName, nil
}

// 定义进度变更事件处理函数。
func (listener *OssProgressListener) ProgressChanged(event *oss.ProgressEvent) {
	switch event.EventType {
	case oss.TransferStartedEvent:
		logger.Info("oss", zap.String("上传进度日志", fmt.Sprintf("上传进度 Started, ConsumedBytes: %d, TotalBytes %d.\n",
			event.ConsumedBytes, event.TotalBytes)))
	case oss.TransferDataEvent:
		fmt.Printf("\r上传进度 Data, ConsumedBytes: %d, TotalBytes %d, %d%%.",
			event.ConsumedBytes, event.TotalBytes, event.ConsumedBytes*100/event.TotalBytes)
	case oss.TransferCompletedEvent:
		logger.Info("oss", zap.String("上传进度日志", fmt.Sprintf("\n上传进度 Completed, ConsumedBytes: %d, TotalBytes %d.\n",
			event.ConsumedBytes, event.TotalBytes)))
	case oss.TransferFailedEvent:
		logger.Info("oss", zap.String("上传进度日志", fmt.Sprintf("\n上传进度 Failed, ConsumedBytes: %d, TotalBytes %d.\n",
			event.ConsumedBytes, event.TotalBytes)))
	default:
	}
}
