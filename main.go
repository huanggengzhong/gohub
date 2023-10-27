package main

import (
	"github.com/gin-gonic/gin"
	"gohub/bootstrap"
)

func main() {
	r := gin.New()
	////中间件
	//r.Use(gin.Logger(), gin.Recovery())
	//
	//r.GET("/", func(context *gin.Context) {
	//	context.JSON(http.StatusOK, gin.H{
	//		"data": "hello world",
	//	})
	//})
	//r.NoRoute(func(context *gin.Context) {
	//	acceptString := context.Request.Header.Get("Accept")
	//	if strings.Contains(acceptString, "text/html") {
	//		context.String(http.StatusNotFound, "404页面不存在")
	//	} else {
	//		context.JSON(http.StatusNotFound, gin.H{
	//			"code":    404,
	//			"message": "页面不存在",
	//		})
	//	}
	//})
	bootstrap.SetupRoute(r)
	r.Run(":8000")
}
