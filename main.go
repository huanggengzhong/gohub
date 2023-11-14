package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"gohub/bootstrap"
	btsConfig "gohub/config"
	"gohub/pkg/config"
)

func init() {
	btsConfig.Initialize()
}
func main() {
	//配置初始化,获取命令行env参数
	var env string
	flag.StringVar(&env, "env", "", "加载.env.xx文件")
	flag.Parse()
	//air 运行或者无env默认设置dev
	if env == "" {
		env = "dev"
	}
	//初始化配置
	config.InitConfig(env)

	//初始化日志
	bootstrap.SetupLogger()
	// 设置 gin 的运行模式，支持 debug, release, test
	// 非 release 模式 gin 终端打印太多信息，干扰到我们程序中的 Log
	// 故此设置为 release，有特殊情况手动改为 debug 即可
	gin.SetMode(gin.ReleaseMode)
	// 初始化Gin实例
	r := gin.New()
	//初始化DB
	bootstrap.SetupDB()
	//初始化路由
	bootstrap.SetupRoute(r)
	err := r.Run(":" + config.Get("app.port"))
	if err != nil {
		fmt.Println("运行错误:", err.Error())
	}
}
