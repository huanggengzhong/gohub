package middlewares

import (
	"github.com/gin-gonic/gin"
	"gohub/pkg/jwt"
	"gohub/pkg/response"
)

// GuestJWT 这个会强制使用游客身份访问,用在一些游客身份才能操作的接口上，例如说用户注册、登录接口等。
func GuestJWT() gin.HandlerFunc {
	return func(c *gin.Context) {

		if len(c.GetHeader("Authorization")) > 0 {
			_, err := jwt.NewJWT().ParserToken(c)
			// 解析 token 成功，说明登录成功了,则告知使用用户身份
			if err == nil {
				response.Unauthorized(c, "请使用游客身份,不要带正确的token")
				c.Abort() //立即停止当前请求的执行
				return
			}
		}
		c.Next()
	}
}
