package auth

import (
	"github.com/gin-gonic/gin"
	v1 "gohub/app/http/controllers/api/v1"
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

}
