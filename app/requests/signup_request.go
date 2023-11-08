package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

//处理请求和表单验证

//type SignupPhoneExistRequest struct {
//	Phone string `json:"phone,omitempty" valid:"phone"`
//}
//
//func ValiadateSignupPhoneExist(data interface{}, c *gin.Context) map[string][]string {
//	//	自定义验证规则
//	rules := govalidator.MapData{
//		"phone": []string{"requied", "digits:11"},
//	}
//	messages := govalidator.MapData{
//		"phone": []string{
//			"requied:手机号不能为空,字段名phone",
//			"digits:手机号长度要求11位",
//		},
//	}
//	options := govalidator.Options{
//		Data:          data,
//		Rules:         rules,
//		TagIdentifier: "valid", //struct里的标识符
//		Messages:      messages,
//	}
//	return govalidator.New(options).ValidateStruct()
//}

type SignupPhoneExistRequest struct {
	Phone string `json:"phone,omitempty" valid:"phone"`
}

func ValidateSignupPhoneExist(data interface{}, c *gin.Context) map[string][]string {

	// 自定义验证规则
	rules := govalidator.MapData{
		"phone": []string{"required", "digits:11"},
	}

	// 自定义验证出错时的提示
	messages := govalidator.MapData{
		"phone": []string{
			"required:手机号为必填项，参数名称 phone",
			"digits:手机号长度必须为 11 位的数字",
		},
	}

	// 配置初始化
	opts := govalidator.Options{
		Data:          data,
		Rules:         rules,
		TagIdentifier: "valid", // 模型中的 Struct 标签标识符
		Messages:      messages,
	}

	// 开始验证
	return govalidator.New(opts).ValidateStruct()
}
