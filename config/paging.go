package config

import "gohub/pkg/config"

func init() {
	config.Add("paging", func() map[string]interface{} {
		return map[string]interface{}{
			// 默认每页条数
			"perpage": 10,
			// URL 中用以分辨多少页的参数
			"url_query_page": "page",
			// URL 中用以分辨排序的参数（使用 id 或者其他）
			"url_query_sort": "sort",
			// URL 中用以分辨排序规则的参数（辨别是正序还是倒序）
			"url_query_order": "order",
			// URL 中用以分辨每页条数的参数
			"url_query_per_page": "per_page",
		}
	})
}
