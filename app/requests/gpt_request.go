package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type ChatGptPRequest struct {
	Content  string `valid:"content" json:"content"`
	IsStream bool   `valid:"is_stream" json:"is_stream"`
}

func ChatGpt(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"content":   []string{"required"},
		"is_stream": []string{"bool"},
	}
	messages := govalidator.MapData{
		"content": []string{
			"required:消息内容不能为空",
		},
		"is_stream": []string{
			"bool:回复类型要求bool类型",
		},
	}
	errs := validate(data, rules, messages)
	return errs
}
