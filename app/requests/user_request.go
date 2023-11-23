package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"gohub/pkg/auth"
)

type UserUpdateRequest struct {
	Name         string `json:"name" valid:"name"`
	City         string `json:"city" valid:"city"`
	Introduction string `json:"introduction" valid:"introduction"`
}

func UserUpdate(data interface{}, c *gin.Context) map[string][]string {
	uid := auth.CurrentUID(c) //如果name被别人使用,校验不通过
	rules := govalidator.MapData{
		"name":         []string{"required", "alpha_num", "between:3,20", "not_exists:users,name," + uid},
		"city":         []string{"required", "max:20"},
		"introduction": []string{"required", "max:240"},
	}
	messages := govalidator.MapData{
		"name": []string{
			"required:用户名不能为空",
			"between:用户名长度需 3-20 之间",
			"alpha_num:用户名只能填数字和字母",
			"not_exists:此用户名已被别人使用,请修改后再提交",
		},
		"city": []string{
			"required:城市为必填项",
			"max:城市需小于20",
		},
		"introduction": []string{
			"required:个人简介为必填项",
			"max:个人简介需小于240",
		},
	}
	return validate(data, rules, messages)
}
