package v1

import (
	"github.com/gin-gonic/gin"
	"gohub/pkg/auth"
	"gohub/pkg/response"
)

type UsersController struct {
	BaseAPIController
}

//

// @Summary 获取当前用户
// @Produce  json
// @Tags 用户

// @Success 200 {string} json "{"code":200,"data":true,"msg":"success"}"
// @Router /v1/user [post]
func (ctrl *UsersController) CurrentUser(c *gin.Context) {
	userModel := auth.CurrentUser(c)
	response.Data(c, userModel)
}
