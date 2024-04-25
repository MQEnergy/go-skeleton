# go-skeleton
基于Go语言（版本：>=v1.22.0）和fiber框架的高性能高并发的Web项目骨架

# 持续更新中...

基于go-skeleton + Reactjs + shadcn-ui 开发的面向出海的插件化电商系统 敬请期待~

项目地址：[https://github.com/MQEnergy/mqshop](https://github.com/MQEnergy/mqshop)


[![GoDoc](https://pkg.go.dev/badge/github.com/MQEnergy/go-skeleton)](https://github.com/MQEnergy/go-skeleton)
[![GitHub license](https://img.shields.io/badge/license-MIT-blue?link=MQEnergy%2Fgo-skeleton)](https://github.com/MQEnergy/go-skeleton/blob/main/LICENSE)
[![codebeat badge](https://codebeat.co/badges/09ce2b03-b0b1-40eb-9ac7-b91bccdb8c0d)](https://codebeat.co/projects/github-com-mqenergy-go-skeleton-main)

## 一、项目结构
```
├── LICENSE
├── Makefile
├── README.md
├── cmd
│   ├── app   # 接口运行命令
│   └── cli   # 命令行运行命令
├── configs         # 配置文件
├── database        # 数据表文件
├── go.mod
├── go.sum
├── internal
│   ├── app           # 模块目录
│   ├── bootstrap     # 服务启动
│   ├── command       # 命令行
│   ├── middleware    # 中间件
│   ├── request       # 请求参数绑定的结构体目录
│   ├── router        # 路由
│   └── vars          # 全局变量
└── pkg
    ├── cache         # 缓存类 redis sync.Map
    ├── command       # 命令行接口定义
    ├── config        # 配置加载类
    ├── crontab       # 定时任务
    ├── database      # 数据库类
    ├── helper        # 帮助函数
    ├── jwtauth       # jwt类
    ├── logger        # 日志类
    └── response      # 接口返回类

```
#### 目前已集成和实现：
- [x] 支持 [jwt](https://github.com/golang/jwt-go) Authorization token验证组件
- [x] 支持 [cors](https://github.com/gofiber/contrib/cors) 跨域组件
- [x] 支持 [gorm](https://gorm.io) 数据库操作组件
- [x] 支持 [redis](https://github.com/go-redis/redis) cache组件
- [x] 支持 [slog](https://github.com/samber/slog-fiber) 日志组件
- [x] 支持 [controller、service、model、middleware、command](https://github.com/MQEnergy/go-skeleton/tree/main/internal/command) 命令行方式生成代码工具
- [x] 支持 [casbin](https://github.com/casbin/casbin) rbac权限 集成于中间件中 [casbin.go](https://github.com/MQEnergy/go-skeleton/blob/main/internal/middleware/casbin.go)
- [x] 支持 [viper](https://github.com/spf13/viper) yaml、json、toml等配置文件解析组件
- [x] 支持 [validator](https://github.com/go-playground/validator) 数据字段验证器组件，同时支持中文
- [x] 支持 [snowflake](https://github.com/bwmarrin/snowflake) 生成雪花算法全局唯一ID
- [x] 实现 ip白名单配置 集成于中间件中 [whitelist.go](https://github.com/MQEnergy/go-skeleton/blob/main/internal/middleware/whitelist.go)
- [x] 实现 [code](https://github.com/MQEnergy/go-skeleton/tree/main/pkg/response/code.go) 统一定义的返回码，[exception](https://github.com/MQEnergy/go-skeleton/tree/main/pkg/response/response.go) 统一错误返回处理组件
- [x] 实现 [wecom](https://github.com/MQEnergy/go-skeleton/tree/main/pkg/wecom/wecom.go) 企业微信组件
- [x] 实现 [oss](https://github.com/MQEnergy/go-skeleton/tree/main/pkg/oss/oss.go) 阿里云oss组件

#### 下一步计划：
- [ ] 支持 cron 定时任务
- [ ] 支持 pprof 性能剖析组件
- [ ] 支持 trace 项目内部链路追踪
- [ ] 支持 [rate](https://pkg.go.dev/golang.org/x/time/rate) 接口限流组件
- [ ] 支持 [grpc](https://github.com/grpc/grpc-go) rpc组件
- [ ] 支持 [go-rabbitmq](https://github.com/MQEnergy/go-rabbitmq) 消息队列组件 基于rabbitmq官方 [amqp](https://github.com/streadway/amqp) 组件封装实现的消费者和生产者
- [ ] 实现 ticker 定时器组件

## 二、运行项目
```shell
# 安装依赖
go mod tidy

# web命令 e: 支持三种环境变量 p: 端口号（默认9527）
go run cmd/app/main.go [-e=dev|test|prod] [-p=9527...]

# 查看帮助
go run cmd/app/main.go -h
go run cmd/cli/main.go -h

# cli命令
go run cmd/cli/main.go [-e=dev|test|prod]

# 热更新
# 安装热更新
go install github.com/cosmtrek/air@latest
air

# 查看帮助
make help

# 格式化代码
make lint

# 打包成window
make windows

# 打包成linux
make linux

# 打包成macos
make darwin
```

## 三、基础功能

配置文件存在于[configs](configs)

### 1、全局变量
在internal/vars目录中可查看全局可用的参数
```go
var (
	BasePath string         // 根目录
	DB       *gorm.DB       // Mysql数据库
	Redis    *redis.Client  // redis连接池
	Router   *fiber.Router  // 路由
	Config   *config.Config // 配置
)
```

### 2、基于gorm/gen生成model和dao
```shell
# 查看帮助
go run cmd/cli/main.go genModel -h

# 命令示例：
# -m: 数据表名称(不填是生成别名为default的数据库的全部模型)
# -e: dev、test、prod(默认环境：dev) 
# -a: 数据库别名（在yaml配置文件中database.mysql.sources.alias里面配置）(默认：default)
go run cmd/cli/main.go genModel [-m=foo] [-e=prod] [-a=demo]
```

命令使用-a参数 会生成新的dao目录，

参考文档：[https://gorm.io/zh_CN/gen/dynamic_sql.html](https://gorm.io/zh_CN/gen/dynamic_sql.html)

1、在entity目录中定义模型的查询接口（按需使用）

参考：[internal/app/entity/admin/admin.go](./internal/app/entity/admin/admin.go)

代码如下：

```go
type Querier interface {
	// SELECT * FROM @@table WHERE id = @id
	GetByID(id int) (gen.T, error)

	// SELECT * FROM @@table WHERE account = @account
	GetByAccount(account string) (*gen.T, error)
}
```

2、在entity.go文件中引入数据表的相关接口，

参考：[internal/app/entity/entity.go](./internal/app/entity/entity.go)

代码如下：

```go
var methodMaps = MethodMaps{
    "yfo_admin": { // 表名称
        func(Querier) {}, // 扩展的查询接口 可多个
        func(admin.Querier) {},
    },
    // ...
}
```

### 3、创建command命令
```shell
# 查看帮助
go run cmd/cli/main.go genCommand -h

# 命令示例 
# -n: 命令行名称 
# -d: 命令存放目录 支持无限极子目录 如：foo/foo 
# -s: 加载已经存在的服务 如：mysql,redis 格式：多个服务以英文逗号相隔 如：mysql,redis
go run cmd/cli/main.go genCommand -n=foo [-d=foo] [-s=mysql,redis]
```

### 4、创建controller
```shell
# 查看帮助
go run cmd/cli/main.go genController -h

# 命令示例 -n: 命令行名称 -d: 命令存放目录 支持无限极子目录 如：foo/foo
go run cmd/cli/main.go genController -n=foo [-d=foo]
```

### 5、创建service
```shell
# 查看帮助
go run cmd/cli/main.go genService -h

# 命令示例 -n: 命令行名称 -d: 命令存放目录 支持无限极子目录 如：foo/foo
go run cmd/cli/main.go genService -n=foo [-d=foo]
```

### 6、中间件
1、通过命令创建中间件
```shell
# 查看帮助
go run cmd/cli/main.go genMiddleware -h

# 命令示例 -n: 命令行名称
go run cmd/cli/main.go genMiddleware -n=foo
```

### 7、日志
```go
import "log/slog"

slog.Info("Info")
slog.Error("Error")
slog.Warning("Warning")
slog.Debug("Debug")
```

### 8、验证器
在controller中文件中直接调用Validate方法
示例如下：
```go
package backend

import (
	"github.com/MQEnergy/go-skeleton/pkg/response"
	"github.com/MQEnergy/go-skeleton/internal/app/controller"
)
type FooController struct {
	controller.Controller
} 
// IndexReq 请求参数绑定
type IndexReq struct {
	Name string `form:"name" query:"name" json:"name" xml:"name" validate:"required"`
	Id   int    `form:"id" query:"id" xml:"id" validate:"required"`
}

// Index ...
func (c *FooController) Index(ctx *fiber.Ctx) error {
    var params IndexReq
    if err := c.Validate(ctx, &params); err != nil {
    return response.BadRequestException(ctx, err.Error())
    }
    return response.SuccessJSON(ctx, "", "index")
}
```

### 9、响应体
在[pkg/response/response.go](./pkg/response/response.go)文件中
```go
// 基础返回
response.JSON(ctx *fiber.Ctx, status int, errcode Code, message string, data interface{})

// 成功返回
response.SuccessJSON(ctx *fiber.Ctx, message string, data interface{})
// ...
```

### 10、数据迁移 migrate

### 11、上传类
参考：

1、调用
[internal/app/controller/backend/attachment.go](./internal/app/controller/backend/attachment.go)

2、组件
[pkg/upload/upload.go](./pkg/upload/upload.go)

### 12、service查询数据
查看[service/backend/auth.go](./internal/app/service/backend/auth.go)

```go
var (
    u         = dao.YfoAdmin
)
adminInfo, err = u.GetByAccount(reqParams.Account) // 这个是entity暴露的查询方法 可查看entity/admin/admin.go文件
```

```
dao.{数据模型}.{查询方法}
```


### 四、单元测试


### 五、格式化代码
```shell
# install
go install mvdan.cc/gofumpt@latest

# run 
gofumpt -l -w .   
```

### 六、检查shadow变量
```shell
# install
go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow

# run path为shadow所在目录
go vet -vettool={path}/shadow ./cmd/app/main.go 
```

### 七、注意
#### 1、air配置文件 .air.toml在不同环境下需要修改
注意查看[.air.toml](.air.toml)文件

### benchmark (Todo)
```bash
wrk -t12 -c1000 -d30s --script=benchmark/login.lua --latency http://127.0.0.1:9527/backend/auth/login
```
