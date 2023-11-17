package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"gohub/app/requests/validators"
	"gohub/pkg/app"
)

type LoginByPhoneRequest struct {
	Phone      string `json:"phone,omitempty" valid:"phone"`
	VerifyCode string `json:"verify_code,omitempty" valid:"verify_code"`
}

type LoginByPasswordRequest struct {
	LoginID       string `json:"login_id" valid:"login_id"`
	Password      string `json:"password" valid:"password"`
	CaptchaID     string `json:"captcha_id,omitempty" valid:"captcha_id"`
	CaptchaAnswer string `json:"captcha_answer,omitempty" valid:"captcha_answer"`
}

// 验证表单
func LoginByPhone(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"phone":       []string{"required", "digits:11"},
		"verify_code": []string{"required", "digits:6"},
	}
	messages := govalidator.MapData{
		"phone": []string{
			"required:手机号不能为空",
			"digits:手机号长度11位",
		},
		"verify_code": []string{
			"required:短信验证码不能为空",
			"digits:短信验证码长度为6",
		},
	}
	errs := validate(data, rules, messages)
	//验证手机验证码
	//测试待删除 start
	if !app.IsLocal() {
		//测试待删除	end
		_data := data.(*LoginByPhoneRequest)
		errs = validators.ValidateVerifyCode(_data.Phone, _data.VerifyCode, errs)
	}

	return errs
}

// 验证表单
func LoginByPassword(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"login_id":       []string{"required", "min:3"},
		"password":       []string{"required", "min:6"},
		"captcha_id":     []string{"required"},
		"captcha_answer": []string{"required", "digits:6"},
	}
	messages := govalidator.MapData{
		"login_id": []string{
			"required:账号不能为空,支持手机号,email,用户名",
			"min:账号最少3位",
		},
		"password": []string{
			"required:密码不能为空",
			"min:密码最少6位",
		},
		"captcha_id": []string{
			"required:图形验证码ID不能为空",
		},
		"captcha_answer": []string{
			"required:图形验证码不能为空",
			"digits:图形验证码要求6位",
		},
	}
	errs := validate(data, rules, messages)
	//验证图形验证码
	_data := data.(*LoginByPasswordRequest)
	errs = validators.ValidateCaptcha(_data.CaptchaID, _data.CaptchaAnswer, errs)
	return errs
}
