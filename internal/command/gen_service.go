package command

import (
	_ "embed"
	"fmt"
	"strings"
	"text/template"

	"github.com/MQEnergy/go-skeleton/internal/vars"
	"github.com/MQEnergy/go-skeleton/pkg/helper"

	"github.com/MQEnergy/go-skeleton/internal/bootstrap"
	"github.com/MQEnergy/go-skeleton/pkg/command"
	"github.com/urfave/cli/v2"
)

//go:embed tpls/gen_service.tpl
var genServiceTpl string

var (
	servicePath    = "/internal/app/service/"
	servicePkgName = "service"
)

type GenService struct{}

// Command ...
func (g *GenService) Command() *cli.Command {
	var (
		name string
		dir  string
	)
	return &cli.Command{
		Name:  "genService",
		Usage: "Create a new service",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "name",
				Aliases:     []string{"n"},
				Value:       "",
				Usage:       "请输入命令工具名称 如：command",
				Destination: &name,
				Required:    true,
			},
			&cli.StringFlag{
				Name:        "dir",
				Aliases:     []string{"d"},
				Value:       "",
				Usage:       "请输入命令工具目录 如：test",
				Destination: &dir,
				Required:    false,
			},
		},
		Before: func(ctx *cli.Context) error {
			return bootstrap.InitConfig()
		},
		Action: func(ctx *cli.Context) error {
			return handleGenService(name, dir)
		},
	}
}

var _ command.Interface = (*GenService)(nil)

// handleGenService ...
func handleGenService(name, dir string) error {
	moduleName := helper.GetProjectModuleName()
	cmdName := strings.ToLower(name)
	cmdDir := strings.ToLower(dir)
	fileName := fmt.Sprintf("%s.go", cmdName)
	rootPath := vars.BasePath + servicePath
	caseCmdName := strings.Title(helper.ToCamelCase(cmdName))
	// 创建目录
	if cmdDir != "" {
		rootPath += fmt.Sprintf("%s/", cmdDir)
		if err := helper.MakeMultiDir(rootPath); err != nil {
			return err
		}
		cmdDirs := strings.Split(cmdDir, "/")
		servicePkgName = cmdDirs[len(cmdDirs)-1]
	}
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
	t1 := template.Must(template.New("genServiceTpl").Parse(genServiceTpl))
	if err := t1.Execute(orPath, map[string]interface{}{
		"ImportPackage": moduleName,
		"PkgName":       servicePkgName,
		"ServiceName":   caseCmdName,
		"CmdDir":        cmdDir,
	}); err != nil {
		return err
	}

	fmt.Println(fmt.Sprintf("\u001B[34m%s\u001B[0m", fileName+" created successfully"))
	return nil
}
