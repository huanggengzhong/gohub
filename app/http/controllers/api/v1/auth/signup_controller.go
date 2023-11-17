package auth

import (
	"github.com/gin-gonic/gin"
	v1 "gohub/app/http/controllers/api/v1"
	"gohub/app/models/user"
	"gohub/app/requests"
	"gohub/pkg/jwt"
	"gohub/pkg/logger"
	"gohub/pkg/response"
)

// 注册控制器
type SignupController struct {
	v1.BaseAPIController
}

//
//type PhoneExistReq struct {
//	Phone string `json:phone`
//}

// @Summary 校验手机号是否已注册
// @Produce  json
// @Tags 授权
// @Param phone query string true "手机号码"
// @Success 200 {string} json "{"code":200,"exist":true,"msg":"ok"}"
// @Router /v1/auth/signup/phone/exist [post]
func (sc SignupController) IsPhoneExist(c *gin.Context) {

	request := requests.SignupPhoneExistRequest{}

	if ok := requests.Validate(c, &request, requests.SignupPhoneExist); !ok {
		return
	}

	response.JSON(c, gin.H{
		"code":  200,
		"exist": user.IsPhoneExist(request.Phone),
		"msg":   "success",
	})
}

// @Summary 校验邮箱是否已注册
// @Produce  json
// @Tags 授权
// @Param email query string true "邮箱"
// @Success 200 {string} json "{"code":200,"data":true,"msg":"success"}"
// @Router /v1/auth/signup/email/exist [post]
func (sc SignupController) IsEmailExist(c *gin.Context) {
	request := requests.SignupEmailExistRequest{}
	if ok := requests.Validate(c, &request, requests.SignupEmailExist); !ok {
		return
	}
	response.JSON(c, gin.H{
		"code":  200,
		"exist": user.IsEmailExist(request.Email),
		"msg":   "success",
	})

}

// @Summary 使用手机验证码进行注册
// @Produce  json
// @Tags 授权
// @Param phone query string true "手机号"
// @Param name query string true "用户名"
// @Param password query string true "密码"
// @Param password_confirm query string true "确认密码"
// @Param verify_code query string true "短信验证码"
// @Success 200 {string} json "{"code":200,"data":true,"msg":"success"}"
// @Router /v1/auth/signup/using-phone [post]
func (sc SignupController) SignupUsingPhone(c *gin.Context) {
	//表单校验
	request := requests.SignupUsingPhoneRequest{}
	logger.Dump(request)
	if ok := requests.Validate(c, &request, requests.SignupUsingPhone); !ok {
		return
	}
	//创建用户数据
	userModel := user.User{
		Name:     request.Name,
		Phone:    request.Phone,
		Password: request.Password,
	}
	//logger.Dump(_user)
	userModel.Create()
	//根据ID判断用户是否存在
	if userModel.ID > 0 {
		//增加token
		token := jwt.NewJWT().IssueToken(userModel.GetStringID(), userModel.Name)
		response.CreatedJSON(c, gin.H{
			"code":    200,
			"token":   token,
			"data":    userModel,
			"message": "创建成功",
		})
	} else {
		response.Abort500(c, "用户创建失败")
	}
}
