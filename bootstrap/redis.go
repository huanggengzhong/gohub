package bootstrap

import (
	"fmt"
	"gohub/pkg/config"
	"gohub/pkg/redis"
)

//redis初始化

func SetupRedis() {
	redis.ConnectRedis(
		fmt.Sprintf("%v:%v", config.GetString("redis.host"), config.GetString("redis.port")),
		config.GetString("redis.username"),
		config.GetString("redis.password"),
		config.GetInt("redis.database"),
	)
}
