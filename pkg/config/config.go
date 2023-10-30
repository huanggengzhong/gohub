package config

import (
	"fmt"
	viperlib "github.com/spf13/viper"
	"gohub/pkg/helpers"
	"os"
)

var viper *viperlib.Viper

func init() {

	// 1. 初始化 Viper 库
	viper = viperlib.New()
	// 2. 配置类型，支持 "json", "toml", "yaml", "yml", "properties",
	//             "props", "prop", "env", "dotenv"
	viper.SetConfigType("env")
	// 3. 环境变量配置文件查找的路径，相对于 main.go
	viper.AddConfigPath(".")
	// 4. 设置环境变量前缀，用以区分 Go 的系统环境变量
	viper.SetEnvPrefix("appenv")
	// 5. 读取环境变量（支持 flags）
	viper.AutomaticEnv()

	ConfigFuncs = make(map[string]ConfigFunc)
}

type ConfigFunc func() map[string]interface{}

var ConfigFuncs map[string]ConfigFunc

func Env(envName string, defaultValue ...interface{}) interface{} {
	fmt.Println(envName, len(defaultValue), "env默认值是")
	if len(defaultValue) > 0 {
		return InternalGet(envName, defaultValue[0])
	}
	return InternalGet(envName)
}

// 初始化配置
func InitConfig(env string) {
	fmt.Println("获取命令行参数:", env)
	// 1. 加载环境变量
	loadEnv(env)
	// 2. 注册配置信息
}

func Add(name string, configFn ConfigFunc) {
	ConfigFuncs[name] = configFn
}

func InternalGet(path string, defaultValue ...interface{}) interface{} {
	if !viper.IsSet(path) || helpers.Empty(viper.Get(path)) {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return nil
	}
	return viper.Get(path)
}

func loadEnv(envSuffix string) {
	// 默认加载 .env 文件，如果有传参 --env=name 的话，加载 .env.name 文件
	envPath := ".env"
	if len(envSuffix) > 0 {
		filePath := ".env.dev" + envSuffix
		if _, err := os.Stat(filePath); err == nil {
			envPath = filePath
		}
	}
	// 加载 env
	viper.SetConfigName(envPath)
	if err := viper.ReadInConfig(); err != nil {
		//panic(err)
	}
	// 监控 .env 文件，变更时重新加载
	viper.WatchConfig()
}
