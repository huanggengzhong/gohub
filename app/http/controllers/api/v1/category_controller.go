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

// @Summary 分类列表
// @Produce  json
// @Tags 内容
// @Param name query string true "分类名"
// @Param description query string true "描述"
// @Success 200 {string} json "{"code":200,"data":true,"msg":"success"}"
// @Router /v1/categories [get]
func (ctrl *CategoryController) Index(c *gin.Context) {
	request := requests.PaginationRequest{}
	if ok := requests.Validate(c, &request, requests.Pagination); !ok {
		return
	}
	data, pager := category.Paginate(c, 10)
	response.JSON(c, gin.H{
		"data":  data,
		"pager": pager,
	})
}

// @Summary 创建分类
// @Produce  json
// @Tags 内容
// @Param name query string true "分类名"
// @Param description query string true "描述"
// @Success 200 {string} json "{"code":200,"data":true,"msg":"success"}"
// @Router /v1/categories [post]
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

// @Summary 修改分类
// @Produce  json
// @Tags 内容
// @Param name query string true "分类名"
// @Param description query string true "描述"
// @Success 200 {string} json "{"code":200,"data":true,"msg":"success"}"
// @Router /v1/categories/:id [put]
func (ctrl *CategoryController) Update(c *gin.Context) {
	//校验参数
	categoryModel := category.Get(c.Param("id"))
	if categoryModel.ID == 0 {
		response.Abort404(c)
		return
	}
	//校验表单
	request := requests.CategoryRequest{}
	if ok := requests.Validate(c, &request, requests.CategorySave); !ok {
		return
	}

	//修改
	categoryModel.Name = request.Name
	categoryModel.Description = request.Description
	rowsEffected := categoryModel.Save()
	if rowsEffected > 0 {
		response.Data(c, categoryModel)
	} else {
		response.Abort500(c, "更新失败")
	}
}

// @Summary 删除分类
// @Produce  json
// @Tags 内容
// @Param name query string true "分类名"
// @Param description query string true "描述"
// @Success 200 {string} json "{"code":200,"data":true,"msg":"success"}"
// @Router /v1/categories/:id [delete]
func (ctrl *CategoryController) Delete(c *gin.Context) {
	//校验参数
	categoryModel := category.Get(c.Param("id"))
	if categoryModel.ID == 0 {
		response.Abort404(c)
		return
	}
	//删除
	rowsEffected := categoryModel.Delete()
	if rowsEffected > 0 {
		response.Success(c)
	} else {
		response.Abort500(c, "删除失败")
	}
}
