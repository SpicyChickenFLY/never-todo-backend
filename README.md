<div align=center><img src="./static/logo.png" width = "200" height = "200" /><h1>never-todo-backend</h1></div>


> 一个基于gin框架搭建的never-todo系列产品的后端同步服务

[ Document in English ](./README.md)

## 总览
这个应用被分为四部分
* [后端数据库（开发中）](https://github.com/SpicyChickenFLY/never-todo-backend) - 使用Golang进行开发
* [前端Web页面（开发中）](https://github.com/bluepongo/never-todo-frontend) - 使用Vue进行开发，可能会用dart写Vue
* [PC端（Win/Linux/Mac）（发布v0.0.1）](https://github.com/bluepongo/never-todo-client)- 使用Electron-Vue框架搭建
* [移动端（Android/IOS）（尚未开发）](https://github.com/SpicyChickenFLY/never-todo-mobile) - 使用Dart/Flutter搭建

本项目为跨平台的never-todo系列产品提供了一个在线同步的后端服务功能，您可以在自己的服务器中部署该服务并在客户端中配置相应信息来实现同步功能，本项目由[SpicyChickenFLY](https://github.com/SpicyChickenFLY)与[bluepongo](https://github.com/bluepongo)合作开发

## 实现功能
* [x] 实现数据库中待办、标签的增删改查
* [x] 将待办、标签的增删改查功能暴露为RESTful的接口
* [ ] 允许用户进行注册登录
* [ ] 记录同步数据信息

#### 项目搭建

```bash
# 编译构建后端服务
build/build.sh
# 允许后端服务开机自启动
systemctl enable never-todo
# 手动开启服务
systemctl start never-todo
# 手动关闭服务
systemctl stop never-todo
# 查看服务详情
sysctemctl status never-todo
```
