package v1

import (
	"github.com/gin-gonic/gin"
	"gohub/app/models/topic"
	"gohub/app/policies"
	"gohub/app/requests"
	"gohub/pkg/auth"
	"gohub/pkg/response"
	"net/http"
)

type TopicsController struct {
	BaseAPIController
}

// @Summary 创建帖子
// @Produce  json
// @Tags 帖子
// @Param title query string true "标题"
// @Param body query string true "内容"
// @Param category_id query string true "分类id"
// @Success 200 {string} json "{"code":200,"data":true,"msg":"success"}"
// @Router /v1/topics [post]
func (t *TopicsController) Store(c *gin.Context) {
	request := requests.TopicRequest{}
	if ok := requests.Validate(c, &request, requests.TopicSave); !ok {
		return
	}
	topicModel := topic.Topic{
		Title:      request.Title,
		Body:       request.Body,
		CategoryID: request.CategoryID,
		UserID:     auth.CurrentUID(c),
	}
	topicModel.Create()
	if topicModel.ID > 0 {
		response.Created(c, topicModel)
	} else {
		response.Abort500(c, "创建失败")
	}
}

// @Summary 修改帖子
// @Produce  json
// @Tags 帖子
// @Success 200 {string} json "{"code":200,"data":true,"msg":"success"}"
// @Router /v1/topics/:id [put]
func (t *TopicsController) Update(c *gin.Context) {
	topicModel := topic.Get(c.Param("id"))
	if topicModel.ID == 0 {
		response.Abort404(c)
		return
	}
	//自己写的才能修改
	if ok := policies.CanModifyTopic(c, topicModel); !ok {
		response.Abort403(c)
		return
	}
	request := requests.TopicRequest{}
	if ok := requests.Validate(c, &request, requests.TopicSave); !ok {
		return
	}

	topicModel.Title = request.Title
	topicModel.Body = request.Body
	topicModel.CategoryID = request.CategoryID
	rowsAffected := topicModel.Save()

	if rowsAffected > 0 {
		response.Data(c, topicModel)
	} else {
		response.Abort500(c, "修改失败")
	}
}

// @Summary 删除帖子
// @Produce  json
// @Tags 帖子
// @Success 200 {string} json "{"code":200,"data":true,"msg":"success"}"
// @Router /v1/topics/:id [delete]
func (t *TopicsController) Delete(c *gin.Context) {
	topicModel := topic.Get(c.Param("id"))
	if topicModel.ID == 0 {
		response.Abort404(c)
		return
	}
	//自己写的才能操作
	if ok := policies.CanModifyTopic(c, topicModel); !ok {
		response.Abort403(c)
		return
	}
	request := requests.TopicRequest{}
	if ok := requests.Validate(c, &request, requests.TopicSave); !ok {
		return
	}

	rowsAffected := topicModel.Delete()

	if rowsAffected > 0 {
		response.Success(c)
	} else {
		response.Abort500(c, "删除失败")
	}
}

// @Summary 帖子列表
// @Produce  json
// @Tags 帖子
// @Param sort query string false "排序(id/created_at/updated_at,默认id)"
// @Param order query string false "排序规则(仅支持 asc（正序）,desc（倒序）)"
// @Param per_page query string false "每页条数(介于 2~100 之间)"
// @Param page query string false "当前页"
// @Success 200 {string} json "{"code":200,"data":true,"msg":"success"}"
// @Router /v1/topics [get]
func (ctrl *TopicsController) Index(c *gin.Context) {
	request := requests.PaginationRequest{}
	if ok := requests.Validate(c, &request, requests.Pagination); !ok {
		return
	}
	data, pager := topic.Paginate(c, 10)
	response.JSON(c, gin.H{
		"code":  http.StatusOK,
		"data":  data,
		"pager": pager,
	})
}

// @Summary 帖子详情
// @Produce  json
// @Tags 帖子
// @Success 200 {string} json "{"code":200,"data":true,"msg":"success"}"
// @Router /v1/topics/:id [get]
func (ctrl *TopicsController) Detail(c *gin.Context) {
	//校验参数
	topicModel := topic.Get(c.Param("id"))
	if topicModel.ID == 0 {
		response.Abort404(c)
		return
	}
	response.Data(c, topicModel)

}