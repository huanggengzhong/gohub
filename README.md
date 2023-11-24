### 命令运行
```go
go run main.go  --env=local
```
热更新运行命令(推荐,无参数main.go默认设置local)
```go
air
```
### 接口文档

swag init 后访问:http://localhost:3000/swagger/index.html#/

### docker部署go 

根据Dockerfile打包影像,如下面的影像名gohub-docker-test
```go
docker build -t gohub-docker-test .
```
运行影像
# 容器中8000端口，自己浏览器上访问的是9001
docker run --rm -p 9001:8000 影像id


