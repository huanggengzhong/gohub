package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type ChatGptPRequest struct {
	Content string `valid:"content" json:"content"`
}

func ChatGpt(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{

		"content": []string{"required"},
	}
	messages := govalidator.MapData{
		"content": []string{
			"required:消息内容不能为空",
		},
	}
	errs := validate(data, rules, messages)
	return errs
}
