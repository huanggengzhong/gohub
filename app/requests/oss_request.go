package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"mime/multipart"
)

type OssUpdateFileRequest struct {
	File *multipart.FileHeader `valid:"file" form:"file"`
}

func OssUpdateFile(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{

		"file:file": []string{"required"},
	}
	messages := govalidator.MapData{
		"file:file": []string{
			"required:文件不能为空",
		},
	}

	return validateFile(c, data, rules, messages)
}
