# 获取go的版本，这里也可以直接写golang:1.17.1 但是体积会比较大点
FROM golang:alpine

# 为镜像设置必须的环境变量
ENV GO115MODULE=on\
    CGO_ENABLED=0\
    GOOS=linux\
    GOARCH=amd64\
    GOPROXY="https://goproxy.cn,direct"

# 当前的工作目录
WORKDIR /app
# 复制项目中的 go.mod 和 go.sum文件并下载依赖信息
#COPY go.mod .
#COPY go.sum .
#RUN go mod download
RUN rm -rf go.mod
RUN rm -rf go.sum
RUN go mod init gohub
RUN go mod tidy



# 将别的文件拷贝到工作目录中
COPY . .

# 编译代码
RUN go build -o main main.go
# 暴露出去端口
EXPOSE 8000

# 启动命令
CMD ["./main"]