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
// @Param sort query string false "排序(id/created_at/updated_at,默认id)"
// @Param order query string false "排序规则(仅支持 asc（正序）,desc（倒序）)"
// @Param per_page query string false "每页条数(介于 2~100 之间)"
// @Param page query string false "当前页"
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

//type CustomObject struct {
//	Field1 string `json:"desc"`
//	Field2 string `json:"c_time"`
//}

// @Summary 分类详情
// @Produce  json
// @Tags 内容
// @Success 200 {string} json "{"code":200,"data":true,"msg":"success"}"
// @Router /v1/categories/:id [get]
func (ctrl *CategoryController) Detail(c *gin.Context) {
	//校验参数
	categoryModel := category.Get(c.Param("id"))
	if categoryModel.ID == 0 {
		response.Abort404(c)
		return
	}
	response.Data(c, categoryModel)
	//测试改为正常时间
	//time := categoryModel.CreatedAt.Format("2006-01-02 15:04:05")
	//response.JSON(c, gin.H{
	//	"code": 200,
	//	"data": CustomObject{
	//		Field1: categoryModel.Description,
	//		Field2: time,
	//	},
	//})
}
