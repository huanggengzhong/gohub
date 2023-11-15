package auth

import (
	"github.com/gin-gonic/gin"
	v1 "gohub/app/http/controllers/api/v1"
	"gohub/pkg/captcha"
	"gohub/pkg/logger"
	"net/http"
)

type VerifyCodeController struct {
	v1.BaseAPIController
}

// ShowCaptcha 获取图片验证码

// @Summary 获取图片验证码
// @Produce  json
// @Tags 授权
// @Success 200 {string} json "{"captcha_id":1,"data":""}"
// @Router /verfy-codes/captcha [post]
func (vc VerifyCodeController) ShowCaptcha(c *gin.Context) {
	//生成验证码
	id, b64s, err := captcha.NewCaptcha().GenerateCaptcha()
	logger.LogIf(err) //日志
	c.JSON(http.StatusOK, gin.H{
		"code":       200,
		"data":       b64s,
		"captcha_id": id,
		"msg":        "success",
	})
}
