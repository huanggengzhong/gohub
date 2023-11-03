package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterAPIRoutes(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		//v1.GET("/", func(context *gin.Context) {
		//	context.JSON(http.StatusOK, gin.H{
		//		"data": "hello world",
		//	})
		//})
		authGroup := v1.Group("/auth")
		{
			authGroup.POST("/signup/phone/exist", func(context *gin.Context) {
				context.JSON(http.StatusOK, gin.H{
					"data": "hello world22",
				})
			})
		}
	}
}
