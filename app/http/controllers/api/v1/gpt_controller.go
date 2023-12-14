package v1

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gohub/app/requests"
	"gohub/pkg/gpt"
	"gohub/pkg/response"
)

type ChatGptController struct {
	BaseAPIController
}
type ChatCompletion struct {
	ID                string      `json:"id"`
	Object            string      `json:"object"`
	Created           int64       `json:"created"`
	Model             string      `json:"model"`
	Choices           []Choice    `json:"choices"`
	Usage             Usage       `json:"usage"`
	SystemFingerprint interface{} `json:"system_fingerprint"`
}

type Choice struct {
	Index        int     `json:"index"`
	Message      Message `json:"message"`
	FinishReason string  `json:"finish_reason"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

// @Summary 聊天
// @Produce  json
// @Tags chatgpt
// @Param content query string true "发送的消息内容"
// @Success 200 {string} json "{"code":200,"data":true,"msg":"success"}"
// @Router /v1/chatgpt/send [post]
func (t *ChatGptController) Send(c *gin.Context) {
	request := requests.ChatGptPRequest{}
	if ok := requests.Validate(c, &request, requests.ChatGpt); !ok {
		return
	}
	// 自定义请求头
	postHeaders := map[string]string{
		"Authorization": "Bearer " + "sk-ef30ytnCGI8m6RTMFd4c0a237a5647Ab877d7e842244F2Af",
		"Content-Type":  "application/json",
	}

	postResponse, err := gpt.SendPostRequest("https://openkey.cloud/v1/chat/completions", request.Content, postHeaders)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 使用json.Unmarshal进行解码
	var chatCompletion ChatCompletion
	err = json.Unmarshal([]byte(string(postResponse)), &chatCompletion)
	if err != nil {
		fmt.Println("解码JSON失败:", err)
		return
	}

	//fmt.Println("POST响应内容:", string(postResponse))
	response.Data(c, chatCompletion)
}

// @Summary 余额
// @Produce  json
// @Tags chatgpt
// @Success 200 {string} json "{"code":200,"data":true,"msg":"success"}"
// @Router /v1/chatgpt/balance [get]
func (t *ChatGptController) Balance(c *gin.Context) {

	// 自定义请求头
	postHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	postResponse, err := gpt.SendGetRequest("https://billing.openkey.cloud/api/token", "sk-ef30ytnCGI8m6RTMFd4c0a237a5647Ab877d7e842244F2Af", postHeaders)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("POST响应内容:", string(postResponse))

	type Data struct {
		Remaining float64 `json:"Remaining"`
		Status    int     `json:"Status"`
		Total     int     `json:"Total"`
		Used      float64 `json:"Used"`
	}
	var data Data
	json.Unmarshal([]byte(string(postResponse)), &data)

	response.Data(c, data)
}
