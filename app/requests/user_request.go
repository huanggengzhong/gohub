package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"gohub/pkg/auth"
	"mime/multipart"
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

type UserUpdateAvatarRequest struct {
	Avatar *multipart.FileHeader `valid:"avatar" form:"avatar"`
}

func UserUpdateAvatar(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		// size 的单位为 bytes
		// - 1024 bytes 为 1kb
		// - 1048576 bytes 为 1mb
		// - 20971520 bytes 为 20mb
		"file:avatar": []string{"ext:jpg,jpeg,png", "size:20971520", "required"},
	}
	messages := govalidator.MapData{
		"file:avatar": []string{
			"ext:头像文件只能上传 png, jpg, jpeg 任意一种的图片",
			"size:头像文件最大不能超过 20MB",
			"required:必须上传图片",
		},
	}

	return validateFile(c, data, rules, messages)
}
