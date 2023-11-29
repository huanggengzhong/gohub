package v1

import (
	"github.com/gin-gonic/gin"
	"gohub/app/requests"
	"gohub/pkg/config"
	"gohub/pkg/file"
	"gohub/pkg/oss"
	"gohub/pkg/response"
)

type OssController struct {
	BaseAPIController
}

// @Summary 上传文件
// @Produce  json
// @Tags 通用
// @Param file query string true "上传文件的key"
// @Success 200 {string} json "{"code":200,"data":true,"msg":"success"}"
// @Router /v1/oss/upload [post]
func (t *OssController) Upload(c *gin.Context) {
	request := requests.OssUpdateFileRequest{}
	if ok := requests.Validate(c, &request, requests.OssUpdateFile); !ok {
		return
	}
	//保存头像
	fileName, filePath, err := file.SaveUploadFile(c, request.File)
	if err != nil {
		response.Abort500(c, "文件上传失败")
		return
	}
	yunFilePath, err := oss.NewBUCKET().UploadFile(fileName, filePath, c)
	if err != nil {
		response.Abort500(c, "文件上传失败")
	}
	response.Data(c, config.Get("oss.oss_aliyun_base_url")+"/"+yunFilePath)
}
