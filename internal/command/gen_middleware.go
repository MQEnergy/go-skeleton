package command

import (
	_ "embed"
	"fmt"
	"go-skeleton/internal/vars"
	"go-skeleton/pkg/helper"
	"strings"
	"text/template"

	"github.com/urfave/cli/v2"
	"go-skeleton/internal/bootstrap"
	"go-skeleton/pkg/command"
)

//go:embed tpls/gen_middleware.tpl
var genMid string

type GenMiddleware struct{}

var (
	middlePath = "/internal/middlewares/"
)

// Command ...
func (g *GenMiddleware) Command() *cli.Command {
	var name string
	return &cli.Command{
		Name:  "genMiddleware",
		Usage: "Create a new middleware",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "name",
				Aliases:     []string{"n"},
				Value:       "",
				Usage:       "请输入命令工具名称 如：command",
				Destination: &name,
				Required:    true,
			},
		},
		Before: func(ctx *cli.Context) error {
			return bootstrap.InitConfig()
		},
		Action: func(ctx *cli.Context) error {
			return handleGenMiddleware(name)
		},
	}
}

var _ command.Interface = (*GenMiddleware)(nil)

// handleGenMiddleware ...
func handleGenMiddleware(name string) error {
	cmdName := strings.ToLower(name)
	fileName := fmt.Sprintf("%s.go", cmdName)
	rootPath := vars.BasePath + middlePath
	caseCmdName := strings.Title(helper.ToCamelCase(cmdName))
	// 判断文件是否存在
	if flag := helper.IsPathExist(rootPath + fileName); flag {
		fmt.Println(fmt.Sprintf("\x1b[31m%s\x1b[0m", cmdName+".go already existed"))
		return nil
	}
	// 创建文件
	orPath, err := helper.MakeFileOrPath(rootPath + fileName)
	if err != nil {
		return err
	}
	// 渲染模板
	t1 := template.Must(template.New("genmid").Parse(genMid))
	if err := t1.Execute(orPath, map[string]interface{}{
		"CmdName": caseCmdName,
	}); err != nil {
		return err
	}

	fmt.Println(fmt.Sprintf("\u001B[34m%s\u001B[0m", fileName+" created successfully"))
	return nil
}
