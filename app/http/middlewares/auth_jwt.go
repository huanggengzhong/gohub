package middlewares

import (
	"github.com/gin-gonic/gin"
	"gohub/app/models/user"
	"gohub/pkg/jwt"
	"gohub/pkg/response"
)

// AuthJWT 中间件用在一些需要用户授权才能操作的接口，例如说创建话题、更新个人资料等。
func AuthJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从标头 Authorization:Bearer xxxxx 中获取信息，并验证 JWT 的准确性
		claims, err := jwt.NewJWT().ParserToken(c)
		if err != nil {
			response.Unauthorized(c, "身份失效,请重新登陆")
			//response.Unauthorized(c, fmt.Sprintf("token错误,请查看 %v 相关的接口文档", config.GetString("app.name")))
			return
		}
		//jwt解析成功,设置用户信息
		userModel := user.Get(claims.UserID)
		if userModel.ID == 0 {
			response.Unauthorized(c, "用户不存在,用户可能已被删除,请用新账号登录")
			return
		}
		// 将用户信息存入 gin.context 里，后续 auth 包将从这里拿到当前用户数据
		c.Set("current_user", userModel)
		c.Set("current_user_id", userModel.GetStringID())
		c.Set("current_user_name", userModel.Name)
		c.Next()
	}
}
