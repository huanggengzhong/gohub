package app

import (
	"gohub/pkg/config"
	"time"
)

func IsLocal() bool {
	return config.Get("app.env") == "local"
}

func IsProduction() bool {
	return config.Get("app.env") == "production"
}

// 当前时间
func TimenowInTimezone() time.Time {
	chinaTimeZone, _ := time.LoadLocation(config.GetString("app.timezone"))
	return time.Now().In(chinaTimeZone)
}
