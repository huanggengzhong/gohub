package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type CategoryRequest struct {
	Name        string `json:"name,omitempty" valid:"name"`
	Description string `json:"description,omitempty" valid:"description"`
}

func CategorySave(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"name":        []string{"required", "not_exists:categories,name"},
		"description": []string{"required"},
	}
	messages := govalidator.MapData{
		"name": []string{
			"required:分类名称为必填项",
		},
		"description": []string{
			"required:分类描述必填",
		},
	}
	return validate(data, rules, messages)
}
