package auth

import (
	"github.com/gin-gonic/gin"
	v1 "gohub/app/http/controllers/api/v1"
	"gohub/app/models/user"
	"gohub/app/requests"
	"gohub/pkg/response"
)

type PasswordController struct {
	v1.BaseAPIController
}

// @Summary 找回密码(通过手机号+短信验证码)
// @Produce  json
// @Tags 授权
// @Param phone query string true "手机号"
// @Param verify_code query string true "短信验证码"
// @Success 200 {string} json "{"code":200,"data":true,"msg":"success"}"
// @Router /v1/auth/password-reset/using-phone [post]
func (pc *PasswordController) ResetByPhone(c *gin.Context) {
	//1.验证表单
	request := requests.ResetByPhoneRequest{}
	if ok := requests.Validate(c, &request, requests.ResetByPhone); !ok {
		return
	}
	//2.更新密码
	userModel := user.GetByPhone(request.Phone)
	if userModel.ID == 0 {
		response.Abort404(c)
	} else {
		userModel.Password = request.Password
		userModel.Save()
		response.Success(c)
	}
}
