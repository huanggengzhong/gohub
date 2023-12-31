package auth

import (
	"github.com/gin-gonic/gin"
	v1 "gohub/app/http/controllers/api/v1"
	"gohub/app/requests"
	"gohub/pkg/auth"
	"gohub/pkg/jwt"
	"gohub/pkg/response"
)

// 登录控制器
type LoginController struct {
	v1.BaseAPIController
}

// @Summary 使用手机验证码进行登录
// @Produce  json
// @Tags 授权
// @Param phone query string true "手机号"
// @Param verify_code query string true "短信验证码"
// @Success 200 {string} json "{"code":200,"data":true,"msg":"success"}"
// @Router /v1/auth/login/using-phone [post]
func (lc *LoginController) LoginByPhone(c *gin.Context) {
	//表单校验
	request := requests.LoginByPhoneRequest{}
	if ok := requests.Validate(c, &request, requests.LoginByPhone); !ok {
		return
	}
	//登录
	user, err := auth.AuthLoginByPhone(request.Phone)
	if err != nil {
		response.Error(c, err, "账号或者密码错误")
	} else {
		token := jwt.NewJWT().IssueToken(user.GetStringID(), user.Name)
		response.JSON(c, gin.H{
			"code":    200,
			"token":   token,
			"message": "success",
		})
	}
}

// @Summary 刷新令牌(延长时间)
// @Produce  json
// @Tags 授权
// @Success 200 {string} json "{"code":200,"data":true,"msg":"success"}"
// @Router /v1/auth/login/refresh-token [post]
func (lc *LoginController) RefreshToken(c *gin.Context) {
	token, err := jwt.NewJWT().RefreshToken(c)
	if err != nil {
		response.Error(c, err, "令牌刷新失败")
	} else {
		response.JSON(c, gin.H{
			"token": token,
		})
	}
}

// @Summary 使用密码登录
// @Produce  json
// @Tags 授权
// @Param login_id query string true "填手机号/email/用户名"
// @Param password query string true "密码"
// @Param captcha_id query string true "图形验证码id"
// @Param captcha_answer query string true "图形验证码"
// @Success 200 {string} json "{"code":200,"data":true,"msg":"success"}"
// @Router /v1/auth/login/using-password [post]
func (lc *LoginController) LoginByPassword(c *gin.Context) {
	//表单验证
	request := requests.LoginByPasswordRequest{}
	if ok := requests.Validate(c, &request, requests.LoginByPassword); !ok {
		return
	}
	//登录
	user, err := auth.Attempt(request.LoginID, request.Password)
	if err != nil {
		response.Unauthorized(c, "账号或者密码错误")
	} else {
		token := jwt.NewJWT().IssueToken(user.GetStringID(), user.Name)
		response.JSON(c, gin.H{
			"code":    200,
			"token":   token,
			"message": "success",
		})
	}

}
