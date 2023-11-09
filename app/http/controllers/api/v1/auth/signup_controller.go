package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	v1 "gohub/app/http/controllers/api/v1"
	"gohub/app/models/user"
	"gohub/app/requests"
	"net/http"
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
	//req := PhoneExistReq{}
	request := requests.SignupPhoneExistRequest{}
	//if err := c.ShouldBindJSON(&req); err != nil {
	if err := c.ShouldBindJSON(&request); err != nil {
		//解析失败和错误信息
		//c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{ //422
		//	"code": 422,
		//	"data": err.Error(),
		//	"msg":  "缺少phone参数",
		//})
		fmt.Println(err.Error())
		return
	}
	//使用表单验证
	errs := requests.ValidateSignupPhoneExist(&request, c)
	if len(errs) > 0 {
		fmt.Println("----打印----", len(errs))
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"data": errs,
			"msg":  "参数校验失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": user.IsPhoneExist(request.Phone),
		"msg":  "success",
	})
}
