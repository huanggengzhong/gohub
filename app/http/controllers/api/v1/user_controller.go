package v1

import (
	"github.com/gin-gonic/gin"
	"gohub/app/models/user"
	"gohub/app/requests"
	"gohub/pkg/auth"
	"gohub/pkg/config"
	"gohub/pkg/file"
	"gohub/pkg/response"
	"net/http"
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
		"code":  http.StatusOK,
		"data":  data,
		"pager": pager,
	})
}

// @Summary 修改当前用户
// @Produce  json
// @Tags 用户
// @Param name query string true "用户名"
// @Param city query string true "城市"
// @Param introduction query string true "个人简介"
// @Success 200 {string} json "{"code":200,"data":true,"msg":"success"}"
// @Router /v1/users [put]
func (ctrl *UsersController) Update(c *gin.Context) {
	request := requests.UserUpdateRequest{}
	if ok := requests.Validate(c, &request, requests.UserUpdate); !ok {
		return
	}
	userModel := auth.CurrentUser(c)
	userModel.Name = request.Name
	userModel.City = request.City
	userModel.Introduction = request.Introduction
	rowsAffected := userModel.Save()
	if rowsAffected > 0 {
		response.Data(c, userModel)
	} else {
		response.Abort500(c, "修改失败")
	}

}

// @Summary 修改用户头像
// @Produce  json
// @Tags 用户
// @description 请使用formdata格式上传,key名avatar
// @Success 200 {string} json "{"code":200,"data":true,"msg":"success"}"
// @Router /v1/users/avatar/upload [put]
func (ctrl *UsersController) UpdateAvatar(c *gin.Context) {
	request := requests.UserUpdateAvatarRequest{}
	if ok := requests.Validate(c, &request, requests.UserUpdateAvatar); !ok {
		return
	}
	//保存头像
	_, avatar, err := file.SaveUploadFile(c, request.Avatar)
	if err != nil {
		response.Abort500(c, "头像上传失败")
		return
	}
	//数据库更新
	userModel := auth.CurrentUser(c)
	userModel.Avatar = config.Get("app.url") + "/" + avatar
	userModel.Save()
	response.Data(c, userModel)
}
