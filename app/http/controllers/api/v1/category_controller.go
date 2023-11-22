package v1

import (
	"github.com/gin-gonic/gin"
	"gohub/app/models/category"
	"gohub/app/requests"
	"gohub/pkg/response"
)

type CategoryController struct {
	BaseAPIController
}

// @Summary 分类
// @Produce  json
// @Tags 内容
// @Param name query string true "分类名"
// @Param description query string true "描述"
// @Success 200 {string} json "{"code":200,"data":true,"msg":"success"}"
// @Router /v1/categories/create [post]
func (ctrl *CategoryController) Store(c *gin.Context) {
	request := requests.CategoryRequest{}
	if ok := requests.Validate(c, &request, requests.CategorySave); !ok {
		return
	}
	categoryModel := category.Category{
		Name:        request.Name,
		Description: request.Description,
	}
	categoryModel.Create()
	if categoryModel.ID > 0 {
		response.Created(c, categoryModel)
	} else {
		response.Abort500(c, "创建失败")
	}
}
