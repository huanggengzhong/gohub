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

### docker部署go 有问题,待定

nginx已部署:
### 编译并运行
项目目录下执行
go build main.go
nohup ./main

### 停止
lsof -i :3000
kill -9 pid


nohup用法资料链接https://zhuanlan.zhihu.com/p/490632944?utm_id=0

完整部署:(会更新最新代码)
查看影像
docker images 
删除多余的旧影像(节省内存)
docker rmi <image>

根据Dockerfile打包影像,如下面的影像名gohub-docker-test
```go
docker build -t gohub-test .
```

运行影像
# 容器中8000端口，自己浏览器上访问的是9001
docker run --rm -p 9001:8000 影像id




2023年11月24日17:40:52 备忘录
env文件APP_ENV改为production会报错