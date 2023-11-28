package config

import (
	"gohub/pkg/config"
)

func init() {
	config.Add("jwt", func() map[string]interface{} {

		res := map[string]interface{}{
			// 允许刷新时间，单位分钟，86400 为两个月，从 Token 的签名时间算起
			"max_refresh_time": config.Env("JWT_MAX_REFRESH_TIME", 86400),
			//过期时间 单位分 一般不超过2小时
			"expire_time": config.Env("JWT_EXPIRE_TIME", 120),
			//debug模式过期时间
			"debug_expire_time": 86400,
		}
		return res
	})
}
