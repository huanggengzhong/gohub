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
	config.InitConfig(env)

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
