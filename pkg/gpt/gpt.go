package gpt

import (
	"encoding/json"
	"gohub/pkg/config"
	"gohub/pkg/logger"
	"io"
	"net/http"
	"strings"
	"sync"
)

var (
	once sync.Once
	body []byte
	err  error
)

// 发送请求
func SendPostRequest(url, data string, headers map[string]string) ([]byte, error) {

	// 构造消息体
	requestData := map[string]interface{}{
		"model": config.Get("gpt.gpt_type"),
		"messages": []map[string]string{
			{"role": "user", "content": data},
		},
	}
	// 将数据转换为JSON字符串
	jsonData, e := json.Marshal(requestData)
	logger.LogIfR(e)
	// 创建一个POST请求
	request, e := http.NewRequest("POST", url, strings.NewReader(string(jsonData)))
	logger.LogIfR(err)
	// 设置请求头
	for key, value := range headers {
		request.Header.Set(key, value)
	}
	// 发送请求
	response, e := http.DefaultClient.Do(request)
	logger.LogIfR(e)
	defer response.Body.Close()
	// 读取响应体
	body, e = io.ReadAll(response.Body)
	logger.LogIfR(e)

	return body, err
}

// 发送请求
func SendGetRequest(url, data string, headers map[string]string) ([]byte, error) {

	// 构造消息体
	requestData := map[string]interface{}{
		"api_key": data,
	}
	// 将数据转换为JSON字符串
	jsonData, e := json.Marshal(requestData)
	logger.LogIfR(e)

	// 创建一个POST请求
	request, e := http.NewRequest("POST", url, strings.NewReader(string(jsonData)))
	logger.LogIfR(e)

	// 设置请求头

	for key, value := range headers {
		request.Header.Set(key, value)
	}

	// 发送请求
	response, e := http.DefaultClient.Do(request)
	logger.LogIfR(e)

	defer response.Body.Close()

	// 读取响应体
	body, e = io.ReadAll(response.Body)
	logger.LogIfR(e)

	return body, err
}
