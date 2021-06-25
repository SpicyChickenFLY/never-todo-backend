<img src="./static/icon-512.png" width = "100" height = "100" div align=center />

# never-todo-backend

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
* [x] 添加新的待办、标签
* [x] 修改待办、标签内容
* [x] 删除待办和标签（软删除）

#### 项目搭建

``` bash
# 安装项目依赖
npm install

# 运行开发环境
npm run dev

# 构建对应开发平台的二进制安装包
npm run build
```
