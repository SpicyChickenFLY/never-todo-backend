<div align=center><img src="./static/logo.png" width = "200" height = "200" /><h1>never-todo-backend</h1></div>


> 一个基于gin框架搭建的never-todo系列产品的后端同步服务

[ Document in English ](./README.md)

## 总览
这个应用被分为五部分
* [后端数据库（开发中）](https://github.com/SpicyChickenFLY/never-todo-backend) - 使用Golang进行开发
* [前端Web页面（开发中）](https://github.com/bluepongo/never-todo-frontend) - 使用Vue进行开发，可能会用dart写Vue
* [PC端（Win/Linux/Mac）（发布v0.0.1）](https://github.com/bluepongo/never-todo-client)- 使用Electron-Vue框架搭建
* [命令行端（Win/Linux/Mac）（开发中）](https://github.com/SpicyChickenFLY/never-todo-cmd) - 使用Golang进行开发
* [移动端（Android/IOS）（尚未开发）](https://github.com/SpicyChickenFLY/never-todo-mobile) - 使用Dart/Flutter搭建

本项目为跨平台的never-todo系列产品提供了一个在线同步的后端服务功能，您可以在自己的服务器中部署该服务并在客户端中配置相应信息来实现同步功能，本项目由[SpicyChickenFLY](https://github.com/SpicyChickenFLY)与[bluepongo](https://github.com/bluepongo)合作开发

## 实现功能
* [x] 实现数据库中待办、标签的增删改查
* [x] 将待办、标签的增删改查功能暴露为RESTful的接口
* [ ] 允许用户进行注册登录
* [ ] 记录同步数据信息

## 项目搭建
### Linux
```bash
# 配置


# 构建
build/build.sh # 编译构建Linux、MacOS、WSL下的后端服务
build/build.bat # 编译构建Windows下的后端服务

# 使用
never-backend start # 后端服务启动
never-backend stop # 后端服务停止
never-backend enable # 允许后Qa端服务开机自启动
never-backend diable # 禁止后端服务自启动
```

## 暴露接口一览

