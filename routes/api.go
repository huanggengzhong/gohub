package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gohub/app/http/controllers/api/v1/auth"
	_ "gohub/docs"
)

func RegisterAPIRoutes(r *gin.Engine) {
	// Swagger 文档路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/v1")
	{
		authGroup := v1.Group("/auth")
		{
			initSignupController := new(auth.SignupController)

			authGroup.POST("/signup/phone/exist", initSignupController.IsPhoneExist)
			authGroup.POST("/signup/email/exist", initSignupController.IsEmailExist)
			authGroup.POST("/signup/using-phone", initSignupController.SignupUsingPhone)

			//获取图片验证码
			vcc := new(auth.VerifyCodeController)
			authGroup.POST("/verify-codes/captcha", vcc.ShowCaptcha)
			authGroup.POST("/varify-codes/phone", vcc.SendUsingPhone)
		}
	}
}
