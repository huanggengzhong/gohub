package app

import (
	"gohub/pkg/config"
	"time"
)

func IsLocal() bool {
	return config.Get("app.env") == "local"
}

func IsTesting() bool {
	return config.Get("app.env") == "testing"
}

func IsProduction() bool {
	return config.Get("app.env") == "production"
}

// 当前时间
func TimenowInTimezone() time.Time {
	chinaTimeZone, _ := time.LoadLocation(config.GetString("app.timezone"))
	return time.Now().In(chinaTimeZone)
}

func URL(path string) string {
	return config.Get("app.url") + path
}

// 拼接带v1
func V1URL(path string) string {
	return URL("/v1/" + path)
}
