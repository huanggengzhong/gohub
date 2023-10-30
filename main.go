package main

import (
	"flag"
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
	config.InitConfig(env)

	r := gin.New()
	//初始化路由
	bootstrap.SetupRoute(r)
	r.Run(":8000")
}
