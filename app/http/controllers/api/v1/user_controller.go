package v1

import (
	"github.com/gin-gonic/gin"
	"gohub/app/models/user"
	"gohub/app/requests"
	"gohub/pkg/auth"
	"gohub/pkg/response"
)

type UsersController struct {
	BaseAPIController
}

//

// @Summary 获取当前用户
// @Produce  json
// @Tags 用户
// @Success 200 {string} json "{"code":200,"data":true,"msg":"success"}"
// @Router /v1/user [get]
func (ctrl *UsersController) CurrentUser(c *gin.Context) {
	userModel := auth.CurrentUser(c)
	response.Data(c, userModel)
}

// @Summary 获取所有用户
// @Produce  json
// @Tags 用户
// @Param sort query string false "排序(id/created_at/updated_at,默认id)"
// @Param order query string false "排序规则(仅支持 asc（正序）,desc（倒序）)"
// @Param per_page query string false "每页条数(介于 2~100 之间)"
// @Param page query string false "当前页"
// @Success 200 {string} json "{"code":200,"data":true,"msg":"success"}"
// @Router /v1/users [get]
func (ctrl *UsersController) Index(c *gin.Context) {
	request := requests.PaginationRequest{}
	if ok := requests.Validate(c, &request, requests.Pagination); !ok {
		return
	}

	data, pager := user.Paginate(c, 10)
	response.JSON(c, gin.H{
		"data":  data,
		"pager": pager,
	})
}
