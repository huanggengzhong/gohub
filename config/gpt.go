// Package config 站点配置信息
package config

import (
	"gohub/pkg/config"
)

func init() {
	config.Add("gpt", func() map[string]interface{} {

		res := map[string]interface{}{
			"gpt_type":            config.GetString("GPT_TYPE"),
			"gpt_sk":              config.GetString("GPT_SK"),
			"gpt_completions_url": config.GetString("GPT_COMPLETIONS_URL"),
			"gpt_balance_url":     config.GetString("GPT_BALANCE_URL"),
		}
		return res
	})
}
