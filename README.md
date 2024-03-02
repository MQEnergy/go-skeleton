# go-skeleton
基于Go语言（版本：>=v1.22.0）和fiber框架的高性能高并发的Web项目骨架

### 项目结构

### 功能点

### 运行项目
```shell
go mod tidy

# web命令
go run cmd/app/main.go

# cli命令
go run cmd/cli/main.go
```

### 全局变量
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
### 数据迁移 migrate



## 二、基础功能

### model和dao生成
```shell
# 查看帮助
go run cmd/cli/main.go genModel -h

# 命令示例：-m: 数据表名称(不填是生成全部模型) -e: 环境：dev、test、prod
go run cmd/cli/main.go genModel -m=foo [-e=prod]
```
entity目录可以定义模型查询接口 可参考：internal/entity/admin
然后在entity.go中引入 参考如下：
```go
var (
	methodMaps = MethodMaps{
		"yfo_admin": { // 表名称
			func(Querier) {}, // 扩展的查询接口 可多个
			func(admin.Querier) {},
		},
		// ...
	}
)
```

### command命令
```shell
# 查看帮助
go run cmd/cli/main.go genCommand -h

# 命令示例 -n: 命令行名称 -d: 命令存放目录 支持无限极子目录 如：foo/foo
go run cmd/cli/main.go genCommand -n=foo [-d=foo]
```

### 中间件
1、通过命令创建中间件
```shell
# 查看帮助
go run cmd/cli/main.go genMiddleware -h

# 命令示例 -n: 命令行名称
go run cmd/cli/main.go genMiddleware -n=foo
```

### 日志
```go
import "log/slog"

slog.Info("Info")
slog.Error("Error")
slog.Warning("Warning")
slog.Debug("Debug")
```

### 验证器
在controller中文件中直接调用Validate方法
示例如下：
```go
package backend

import (
	"go-skeleton/pkg/response"
	"go-skeleton/internal/app/controller"
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
    params := &IndexReq{}
    if err := c.Validate(ctx, params); err != nil {
    return response.BadRequestException(ctx, err.Error())
    }
    return response.SuccessJSON(ctx, "", "index")
}
```

### 响应体
在pkg/response/response.go文件中
```go
// 基础返回
response.JSON(ctx *fiber.Ctx, status int, errcode Code, message string, data interface{})

// 成功返回
response.SuccessJSON(ctx *fiber.Ctx, message string, data interface{})
// ...
```

### 格式化代码
```shell
# install
go install mvdan.cc/gofumpt@latest

# run 
gofumpt -l -w .   
```

### 单元测试
