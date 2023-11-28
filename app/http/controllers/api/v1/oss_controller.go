package v1

import (
	"github.com/gin-gonic/gin"
	"gohub/pkg/upload"
)

type OssController struct {
	BaseAPIController
}

// @Summary 上传文件
// @Produce  json
// @Tags 通用
// @Param file query string true "key"
// @Success 200 {string} json "{"code":200,"data":true,"msg":"success"}"
// @Router /v1/oss/upload [post]
func (t *OssController) Upload(c *gin.Context) {
	_, header, _ := c.Request.FormFile("file")

	upload.UploadFile(header)
}
