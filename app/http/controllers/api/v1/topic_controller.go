package v1

import (
	"github.com/gin-gonic/gin"
	"gohub/app/models/topic"
	"gohub/app/requests"
	"gohub/pkg/auth"
	"gohub/pkg/response"
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
// @Param title query string true "标题"
// @Param body query string true "内容"
// @Success 200 {string} json "{"code":200,"data":true,"msg":"success"}"
// @Router /v1/topics/:id [put]
func (t *TopicsController) Update(c *gin.Context) {
	topicModel := topic.Get(c.Param("id"))
	if topicModel.ID == 0 {
		response.Abort404(c)
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
