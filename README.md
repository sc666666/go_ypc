# go_ypc


### 目录结构

```
├── app
│   ├── http
│   │   ├── controllers
│   │   ├── middleware
│   │   └── requests
│   └── models
├── bootstrap
├── config
├── pkg
├── routes
├── tests
└── vendor
```

- app 应用主目录，包含程序核心代码
    - http 请求处理目录
        - controllers 控制器目录
        - middleware 中间件目录
        - requests 验证器目录
    - models 模型目录
- bootstrap 引导启动目录，比如初始化路由、数据库
- config 配置文件目录
- pkg 扩展目录
- routes 路由文件目录
- tests 测试文件目录
- vendor 第三方包目录
