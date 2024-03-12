package v1

import (
	"gohub/app/models/log"
	"gohub/app/requests"
	"gohub/pkg/response"

	"github.com/gin-gonic/gin"
)

type LogController struct {
	BaseAPIController
}

// @Summary 增加日志
// @Produce  json
// @Tags 数据库日志
// @Param message query string true "内容"
// @Success 200 {string} json "{"code":200,"data":true,"msg":"success"}"
// @Router /v1/log/add [post]
func (t *LogController) Add(c *gin.Context) {
	request := requests.LogRequest{}
	if ok := requests.Validate(c, &request, requests.LogSave); !ok {
		return
	}
	logModel := log.Log{
		Message: request.Message,
	}
	logModel.Create()
	if logModel.ID > 0 {
		response.Created(c, logModel)
	} else {
		response.Abort500(c, "创建失败")
	}
}
