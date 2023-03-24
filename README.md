# gin-layout
`gin-layout`是基于`gin`+`wire`的脚手架。项目结构类似与微服务框架[kratos](https://github.com/go-kratos/kratos)

Only supports `Go1.18+`

### 用到的相关包有：

[gin](https://github.com/gin-gonic/gin) 

[wire](https://github.com/google/wire)

[gorm](https://gorm.io/gorm)

[go-redis](https://github.com/go-redis/redis)

[logrus](https://github.com/sirupsen/logrus)

[fasthttp](https://github.com/valyala/fasthttp)

### 编译及运行可查看Makefile文件

### 结构如下：
```
.
├── Makefile
├── README.md
├── bin // 放可执行文件
├── cmd // 整个项目启动的入口文件
│         └── app
│             ├── main.go
│             ├── wire.go // 用wire依赖注入
│             └── wire_gen.go
├── conf // 配置文件
│         └── app.yml
├── go.mod
├── go.sum
├── internal // 该服务所有不对外暴露的代码，通常的业务逻辑都在这下面，使用internal避免错误引用
│         ├── biz // 业务逻辑的组装层，类似 DDD 的 domain 层，data 类似 DDD 的 repo，而 repo 接口在这里定义，使用依赖倒置的原则。
│         │         └── biz.go
│         ├── conf // config的结构定义
│         │         └── config.go
│         ├── data // 业务数据访问，包含 cache、db 等封装，实现了 biz 的 repo 接口。
│         │         ├── data.go
│         │         ├── model // 数据模型
│         │         │         └── model.go
│         │         └── redis.go
│         ├── pkg // 内部使用的一些公共代码以及错误码
│         ├── router // api的路由以及中间件，可以称之为server层
│         │         ├── middleware.go
│         │         └── router.go
│         └── service // 实现了 api 定义的服务层，类似 DDD 的 application 层。
│             └─── service.go
└── pkg // 对一些公共包的简单封装

```

### 待完善

