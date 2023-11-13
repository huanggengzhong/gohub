package app

import (
	"gohub/pkg/config"
)

func IsLocal() bool {
	return config.Get("app.env") == "local"
}
