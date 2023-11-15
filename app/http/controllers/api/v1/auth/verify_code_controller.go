package auth

import (
	"github.com/gin-gonic/gin"
	v1 "gohub/app/http/controllers/api/v1"
	"gohub/app/requests"
	"gohub/pkg/captcha"
	"gohub/pkg/logger"
	"gohub/pkg/response"
	"gohub/pkg/verifycode"
)

type VerifyCodeController struct {
	v1.BaseAPIController
}

// ShowCaptcha 获取图片验证码

// @Summary 获取图片验证码
// @Produce  json
// @Tags 授权
// @Success 200 {string} json "{"captcha_id":1,"data":""}"
// @Router /verify-codes/captcha [post]
func (vc VerifyCodeController) ShowCaptcha(c *gin.Context) {
	//生成验证码
	id, b64s, err := captcha.NewCaptcha().GenerateCaptcha()
	logger.LogIf(err) //日志
	response.JSON(c, gin.H{
		"code":       200,
		"data":       b64s,
		"captcha_id": id,
		"msg":        "success",
	})
}

// @Summary 发送短信验证码(前提先获取图片验证码)
// @Produce  json
// @Tags 授权
// @Param phone query string true "手机号码"
// @Param captcha_id query string true "图片验证码ID"
// @Param captcha_answer query string true "图片验证码答案"
// @Success 200 {string} json "{"code":200,"data":""}"
// @Router /verify-codes/phone [post]
func (vc VerifyCodeController) SendUsingPhone(c *gin.Context) {
	//校验表单
	request := requests.VerifyCodePhoneRequest{}
	if ok := requests.Validate(c, &request, requests.VerifyCodePhone); !ok {
		return
	}
	//发送短信
	if ok := verifycode.NewVerifyCode().SendSMS(request.Phone); !ok {
		response.Abort500(c, "发送短信失败")
	} else {
		response.Success(c)
	}
}
