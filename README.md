# nginx-demo-backend
Nginx介绍后端Golang项目，最简单的服务端项目（比Java启动SpringBoot项目轻量得多）
## 快速开始
```shell
git clone https://github.com/xlxingRun/nginx-demo-backend.git
```
进入文件夹 
```shell
cd nginx-demo-backend
```
运行
```shell
go run main.go
```
## 介绍
本项目监听端口号9000，包含了两个后端请求
- `/hello`：第一条消息
- `/api/hello`：第二条消息

如果在本地启动，在浏览器中输入  
http://localhost:9000/hello  
http://localhost:9000/api/hello
