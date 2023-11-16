package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"gohub/app/requests/validators"
)

//处理请求和表单验证

type SignupEmailExistRequest struct {
	Email string `json:"email,omitempty" valid:"email"`
}

type SignupPhoneExistRequest struct {
	Phone string `json:"phone,omitempty" valid:"phone"`
}

type SignupUsingPhoneRequest struct {
	Phone           string `json:"phone,omitempty" valid:"phone"`
	Name            string `json:"name" valid:"name"`
	Password        string `json:"password" valid:"password"`
	PasswordConfirm string `json:"password_confirm,omitempty" valid:"password_confirm"`
	VerifyCode      string `json:"verify_code,omitempty" valid:"verify_code"`
}

func SignupEmailExist(data interface{}, c *gin.Context) map[string][]string {
	//	自定义验证规则
	rules := govalidator.MapData{
		"email": []string{"required", "min:4", "max:30", "email"},
	}
	messages := govalidator.MapData{
		"email": []string{
			"required:邮箱不能为空,字段名email",
			"min:邮箱需大于4",
			"max:邮箱需小于30",
			"email:邮箱格式不正确",
		},
	}
	return validate(data, rules, messages)
	//options := govalidator.Options{
	//	Data:          data,
	//	Rules:         rules,
	//	TagIdentifier: "valid", //struct里的标识符
	//	Messages:      messages,
	//}
	//return govalidator.New(options).ValidateStruct()
}

func SignupPhoneExist(data interface{}, c *gin.Context) map[string][]string {

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

	//// 配置初始化
	//opts := govalidator.Options{
	//	Data:          data,
	//	Rules:         rules,
	//	TagIdentifier: "valid", // 模型中的 Struct 标签标识符
	//	Messages:      messages,
	//}
	//
	//// 开始验证
	//return govalidator.New(opts).ValidateStruct()

	return validate(data, rules, messages)
}

func SignupUsingPhone(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		//not_exists是自定义rules,假设有一个用户注册的场景，需要验证用户名是否已经存在于数据库中
		"phone": []string{"required", "digits:11", "not_exists:users,phone"},
		//alpha_num验证的字段必须完全是字母数字字符
		"name":             []string{"required", "alpha_num", "between:3,20", "not_exists:users,name"},
		"verify_code":      []string{"required", "digits:6"},
		"password":         []string{"required", "min:6"},
		"password_confirm": []string{"required", "min:6"},
	}
	messages := govalidator.MapData{
		"phone": []string{
			"required:手机号不能为空",
			"digits:手机号要求11位",
		},
		"name": []string{
			"required:用户名不能为空",
			"alpha_num:用户名格式错误,只允许输入数字和英文",
			"between:用户名长度需在3-30之间",
		},
		"password": []string{
			"required:密码不能为空",
			"min:密码不能少于6位",
		},
		"password_confirm": []string{
			"required:确认密码不能为空",
			"min:确认密码不能少于6位",
		},
		"verify_code": []string{
			"required:短信验证码不能为空",
			"digits:短信验证码要求6位",
		},
	}
	errs := validate(data, rules, messages)

	_data := data.(*SignupUsingPhoneRequest)
	//校验两个密码是否一致
	errs = validators.ValidatePasswordConfirm(_data.Password, _data.PasswordConfirm, errs)
	//测试待删除 start
	//if !app.IsLocal() {}
	//测试待删除 end
	errs = validators.ValidateVerifyCode(_data.Phone, _data.VerifyCode, errs)

	return errs

}
