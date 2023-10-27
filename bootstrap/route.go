package bootstrap

import (
	"github.com/gin-gonic/gin"
	"gohub/routes"
	"net/http"
	"strings"
)

func SetupRoute(r *gin.Engine) {

	registerGlobalMiddleWare(r)
	routes.RegisterAPIRoutes(r)
	setup404Router(r)
}

// 路由中间件
func registerGlobalMiddleWare(r *gin.Engine) {
	r.Use(gin.Logger(), gin.Recovery())
}

// 404
func setup404Router(r *gin.Engine) gin.HandlerFunc {
	return func(context *gin.Context) {
		acceptString := context.Request.Header.Get("accept")
		if strings.Contains(acceptString, "text/html") {
			context.String(http.StatusNotFound, "页面不存在")
		} else {
			context.JSON(http.StatusNotFound, gin.H{
				"data": "页面不存在",
			})
		}
	}
}
