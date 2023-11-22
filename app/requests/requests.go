package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"gohub/pkg/response"
)

type ValidatorFunc func(interface{}, *gin.Context) map[string][]string

func Validate(c *gin.Context, obj interface{}, handler ValidatorFunc) bool {

	// 1. 解析请求，支持 JSON 数据、表单请求和 URL Query
	// 2. 表单验证
	// 3.判断验证是否通过
	if err := c.ShouldBind(obj); err != nil {
		//解析失败和错误信息
		response.BadRequest(c, err, "请求解析错误，请确认请求格式是否正确。上传文件请使用 multipart 标头，参数请使用 JSON 格式。")
		return false
	}
	//使用表单验证
	errs := handler(obj, c)
	if len(errs) > 0 {
		response.ValidationError(c, errs)
		//c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
		//	"code": 422,
		//	"data": errs,
		//	"msg":  "参数校验失败,具体看data",
		//})
		return false
	}

	return true
}

func validate(data interface{}, rules govalidator.MapData, message govalidator.MapData) map[string][]string {
	//配置
	opts := govalidator.Options{
		Data:          data,
		Rules:         rules,
		Messages:      message,
		TagIdentifier: "valid",
	}
	return govalidator.New(opts).ValidateStruct()
}
