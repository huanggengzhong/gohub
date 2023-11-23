package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type TopicRequest struct {
	Title      string `json:"title,omitempty" valid:"title"`
	Body       string `json:"body,omitempty" valid:"body"`
	CategoryID string `json:"category_id,omitempty" valid:"category_id"`
}

func TopicSave(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"title":       []string{"required", "min:3", "max:40"},
		"body":        []string{"required", "min:10", "max:50000"},
		"category_id": []string{"required", "exists:categories,id"},
	}
	messages := govalidator.MapData{
		"title": []string{
			"required:帖子标题不能为空",
			"min:标题长度需大于 3",
			"max:标题长度需小于 40",
		},
		"body": []string{
			"required:帖子内容为必填项",
			"min:长度需大于 10",
			"max:长度需小于50000",
		},
		"category_id": []string{
			"required:帖子分类为必填项",
			"exists:帖子分类未找到",
		},
	}
	return validate(data, rules, messages)
}
