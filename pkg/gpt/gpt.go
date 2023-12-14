package gpt

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	// 使用sync.Once确保初始化操作只执行一次
	once.Do(func() {
		// 构造消息体
		requestData := map[string]interface{}{
			"model": "gpt-3.5-turbo",
			"messages": []map[string]string{
				{"role": "user", "content": data},
			},
		}
		// 将数据转换为JSON字符串
		jsonData, e := json.Marshal(requestData)
		if e != nil {
			err = fmt.Errorf("JSON序列化失败: %s", e)
			return
		}
		// 创建一个POST请求
		request, e := http.NewRequest("POST", url, strings.NewReader(string(jsonData)))
		if e != nil {
			err = fmt.Errorf("创建POST请求失败: %s", e)
			return
		}

		// 设置请求头

		for key, value := range headers {
			request.Header.Set(key, value)
		}

		// 发送请求
		response, e := http.DefaultClient.Do(request)
		if e != nil {
			err = fmt.Errorf("HTTP POST请求失败: %s", e)
			return
		}
		defer response.Body.Close()

		// 读取响应体
		body, e = ioutil.ReadAll(response.Body)
		if e != nil {
			err = fmt.Errorf("读取响应体失败: %s", e)
			return
		}
	})

	return body, err
}

// 发送请求
func SendGetRequest(url, data string, headers map[string]string) ([]byte, error) {

	// 使用sync.Once确保初始化操作只执行一次
	once.Do(func() {
		// 构造消息体
		requestData := map[string]interface{}{
			"api_key": data,
		}
		// 将数据转换为JSON字符串
		jsonData, e := json.Marshal(requestData)
		if e != nil {
			err = fmt.Errorf("JSON序列化失败: %s", e)
			return
		}
		// 创建一个POST请求
		request, e := http.NewRequest("POST", url, strings.NewReader(string(jsonData)))
		if e != nil {
			err = fmt.Errorf("创建POST请求失败: %s", e)
			return
		}

		// 设置请求头

		for key, value := range headers {
			request.Header.Set(key, value)
		}

		// 发送请求
		response, e := http.DefaultClient.Do(request)
		if e != nil {
			err = fmt.Errorf("HTTP POST请求失败: %s", e)
			return
		}
		defer response.Body.Close()

		// 读取响应体
		body, e = ioutil.ReadAll(response.Body)
		if e != nil {
			err = fmt.Errorf("读取响应体失败: %s", e)
			return
		}
	})

	return body, err
}
