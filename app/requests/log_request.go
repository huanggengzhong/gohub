package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type LogRequest struct {
	Message string `json:"message,omitempty" valid:"message"`
}

func LogSave(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"message": []string{"required"},
	}
	messages := govalidator.MapData{
		"message": []string{
			"required:日志内容不能为空",
		},
	}
	return validate(data, rules, messages)
}
