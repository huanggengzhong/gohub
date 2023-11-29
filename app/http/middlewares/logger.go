// 日志中间件
package middlewares

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	"gohub/pkg/helpers"
	"gohub/pkg/logger"
	"io/ioutil"
	"strings"
	"time"
)

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

// Logger 记录请求日志
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取response内容
		w := &responseBodyWriter{body: &bytes.Buffer{}, ResponseWriter: c.Writer}
		c.Writer = w
		//获取请求数据
		var requestBody []byte
		if c.Request.Body != nil {
			// c.Request.Body 是一个 buffer 对象，只能读取一次
			requestBody, _ = ioutil.ReadAll(c.Request.Body)
			// 读取后，重新赋值 c.Request.Body ，以供后续的其他操作
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestBody))
		}
		//设置时间
		start := time.Now()
		c.Next()
		//开始记录日志
		cost := time.Since(start)
		responStatus := c.Writer.Status()

		logFields := []zap.Field{
			zap.Int("status", responStatus),
			zap.String("request", c.Request.Method+" "+c.Request.URL.String()),
			zap.String("query", c.Request.URL.RawQuery),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.String("time", helpers.MicrosecondsStr(cost)),
		}
		if c.Request.Method == "POST" || c.Request.Method == "PUT" || c.Request.Method == "DELETE" {
			//请求内容
			var allow bool = strings.Contains(c.Request.URL.Path, "upload")
			//屏蔽上传的请求日志,不然日子库文件太大了,也影响了上传速度
			if !allow {
				logFields = append(logFields, zap.String("请求体是", string(requestBody)))
			}
			//响应内容
			logFields = append(logFields, zap.String("响应体是", w.body.String()))
		}
		// 除了 StatusBadRequest 以外，warning 提示一下，常见的有 403 404，开发时都要注意
		if responStatus > 400 && responStatus <= 499 {
			logger.Warn("HTTP Warning Log"+cast.ToString(responStatus), logFields...)
		} else if responStatus >= 500 && responStatus <= 599 {
			logger.Error("HTTP Error Log"+cast.ToString(responStatus), logFields...)
		} else {

			logger.Debug("HTTP Access Log"+cast.ToString(responStatus), logFields...)
		}
	}
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}
