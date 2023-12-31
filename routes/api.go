package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	controllers "gohub/app/http/controllers/api/v1"
	"gohub/app/http/controllers/api/v1/auth"
	"gohub/app/http/middlewares"
	_ "gohub/docs"
	"net/http"
)

// 注册路由
func RegisterAPIRoutes(r *gin.Engine) {
	//静态文件访问
	r.StaticFS("/public", http.Dir("public"))
	// Swagger 文档路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//v1 := r.Group("/v1")
	var v1 *gin.RouterGroup
	//if len(config.Get("app.api_domain")) == 0 {
	//	v1 = r.Group("/api/v1")
	//} else {
	//	v1 = r.Group("/v1")
	//}
	v1 = r.Group("/v1")
	// 全局限流中间件：每小时限流。这里是所有 API （根据 IP）请求加起来。
	// 作为参考 Github API 每小时最多 60 个请求（根据 IP）。
	// 测试时，可以调高一点。
	v1.Use(middlewares.LimitIP("200-H"))
	{
		authGroup := v1.Group("/auth")
		// 限流中间件：每小时限流，作为参考 Github API 每小时最多 60 个请求（根据 IP）
		// 测试时，可以调高一点
		authGroup.Use(middlewares.LimitIP("1000-H"))
		{
			//注册
			initSignupController := new(auth.SignupController)
			authGroup.POST("/signup/phone/exist", middlewares.GuestJWT(), middlewares.LimitPerRoute("60-H"), initSignupController.IsPhoneExist)
			authGroup.POST("/signup/email/exist", middlewares.GuestJWT(), middlewares.LimitPerRoute("60-H"), initSignupController.IsEmailExist)
			authGroup.POST("/signup/using-phone", middlewares.GuestJWT(), initSignupController.SignupUsingPhone)

			//验证
			vcc := new(auth.VerifyCodeController)
			authGroup.POST("/verify-codes/captcha", middlewares.LimitPerRoute("20-H"), vcc.ShowCaptcha)
			authGroup.POST("/varify-codes/phone", middlewares.LimitPerRoute("20-H"), vcc.SendUsingPhone)

			//登录
			loginControllerInit := new(auth.LoginController)
			authGroup.POST("/login/using-phone", middlewares.GuestJWT(), loginControllerInit.LoginByPhone)
			authGroup.POST("/login/using-password", middlewares.GuestJWT(), loginControllerInit.LoginByPassword)
			authGroup.POST("/login/refresh-token", middlewares.AuthJWT(), loginControllerInit.RefreshToken)

			//重置密码
			passwordControllerInit := new(auth.PasswordController)
			authGroup.POST("/password-reset/using-phone", middlewares.GuestJWT(), passwordControllerInit.ResetByPhone)

			//用户接口
			UsersControllerInit := new(controllers.UsersController)
			v1.GET("/user", middlewares.AuthJWT(), UsersControllerInit.CurrentUser)
			usersGroup := v1.Group("/users")
			{
				usersGroup.GET("", middlewares.AuthJWT(), UsersControllerInit.Index)
				usersGroup.PUT("", middlewares.AuthJWT(), UsersControllerInit.Update)
				usersGroup.PUT("/avatar/upload", middlewares.AuthJWT(), UsersControllerInit.UpdateAvatar)
			}
			//分类接口
			CategoryControllerInit := new(controllers.CategoryController)
			cgcGroup := v1.Group("/categories")
			{
				cgcGroup.GET("", middlewares.AuthJWT(), CategoryControllerInit.Index)
				cgcGroup.POST("", middlewares.AuthJWT(), CategoryControllerInit.Store)
				cgcGroup.PUT("/:id", middlewares.AuthJWT(), CategoryControllerInit.Update)
				cgcGroup.DELETE("/:id", middlewares.AuthJWT(), CategoryControllerInit.Delete)
				cgcGroup.GET("/:id", middlewares.AuthJWT(), CategoryControllerInit.Detail)
			}
			//帖子接口
			TopicsControllerInit := new(controllers.TopicsController)
			tpcGroup := v1.Group("/topics")
			{
				tpcGroup.POST("", middlewares.AuthJWT(), TopicsControllerInit.Store)
				tpcGroup.PUT("/:id", middlewares.AuthJWT(), TopicsControllerInit.Update)
				tpcGroup.DELETE("/:id", middlewares.AuthJWT(), TopicsControllerInit.Delete)
				tpcGroup.GET("/:id", TopicsControllerInit.Detail)
				tpcGroup.GET("", TopicsControllerInit.Index)

			}
			//通用上传
			OssControllerInit := new(controllers.OssController)
			ossGroup := v1.Group("/oss")
			{
				ossGroup.PUT("/upload", OssControllerInit.Upload)
				ossGroup.PUT("/uploadpart", OssControllerInit.UploadPart)
			}
			//chatgpt
			ChatGptControllerInit := new(controllers.ChatGptController)
			gptGroup := v1.Group("/chatgpt")
			{
				gptGroup.POST("/send", ChatGptControllerInit.Send)
				gptGroup.GET("/balance", ChatGptControllerInit.Balance)
			}
		}
	}
}
