package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"gohub/app/requests/validators"
	"gohub/pkg/app"
)

type ResetByPhoneRequest struct {
	Phone      string `json:"phone,omitempty" valid:"phone"`
	VerifyCode string `json:"verify_code,omitempty" valid:"verify_code"`
	Password   string `json:"password" valid:"password"`
}

func ResetByPhone(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"phone":       []string{"required", "digits:11"},
		"verify_code": []string{"required", "digits:6"},
		"password":    []string{"required", "min:6"},
	}
	messages := govalidator.MapData{
		"phone": []string{
			"required:手机号不能为空",
			"digits:手机号必须11位",
		},
		"verify_code": []string{
			"required:短信验证码不能为空",
			"digits:短信验证码邀请6位",
		},
		"password": []string{
			"required:新密码不能为空",
			"min:密码最少6位",
		},
	}
	errs := validate(data, rules, messages)
	//检查验证码
	if !app.IsLocal() {
		_data := data.(*ResetByPhoneRequest)
		errs = validators.ValidateVerifyCode(_data.Phone, _data.VerifyCode, errs)
	}
	return errs
}
