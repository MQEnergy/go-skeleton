# go-skeleton
基于Go语言和fiber框架的高性能高并发的Web项目骨架

### 项目结构

### 功能点

### 单元测试

### 运行项目

### 全局变量

### 数据迁移 migrate

### 基础功能

### 生成model和dao

### 中间件
middlewares目录定义中间件：
```go
package middlewares

func FooMiddleware() fiber.Handler {
	return func(ctx fiber.Ctx) error {
		// Todo: implement
		return ctx.Next()
    }
}
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

### 格式化代码
```shell
# install
go install mvdan.cc/gofumpt@latest

# run 
gofumpt -l -w .   
```