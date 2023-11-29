// Package config 站点配置信息
package config

import (
	"gohub/pkg/config"
)

func init() {
	config.Add("oss", func() map[string]interface{} {

		res := map[string]interface{}{
			"oss_aliyun_base_url":        config.GetString("OSS_ALIYUN_BASE_URL"),
			"oss_aliyun_endpoint":        config.GetString("OSS_ALIYUN_ENDPOINT"),
			"oss_aliyun_accesskeyid":     config.GetString("OSS_ALIYUN_ACCESSKEYID"),
			"oss_aliyun_accesskeysecret": config.GetString("OSS_ALIYUN_ACCESSKEYSECRET"),
			"oss_aliyun_bucketname":      config.GetString("OSS_ALIYUN_BUCKETNAME"),
		}
		return res
	})
}
