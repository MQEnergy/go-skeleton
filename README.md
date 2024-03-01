# go-skeleton
基于Go语言和fiber框架的高性能高并发的Web项目骨架

### 项目结构

### 功能点

### 单元测试

### 运行项目

### 全局变量

### 数据迁移 migrate

### 基础功能

### model和dao生成

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

### 格式化代码
```shell
# install
go install mvdan.cc/gofumpt@latest

# run 
gofumpt -l -w .   
```