package command

import (
	_ "embed"
	"fmt"
	"strings"
	"text/template"

	"github.com/MQEnergy/go-skeleton/internal/bootstrap/boots"

	"github.com/MQEnergy/go-skeleton/internal/vars"
	"github.com/MQEnergy/go-skeleton/pkg/helper"

	"github.com/MQEnergy/go-skeleton/pkg/command"
	"github.com/urfave/cli/v2"
)

//go:embed tpls/gen_controller.tpl
var genCtlTpl string

var (
	ctlPath    = "/internal/app/controller/"
	ctlPkgName = "controller"
)

type GenController struct{}

// Command ...
func (g *GenController) Command() *cli.Command {
	var (
		name string
		dir  string
	)
	return &cli.Command{
		Name:  "genController",
		Usage: "Create a new controller",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "name",
				Aliases:     []string{"n"},
				Value:       "",
				Usage:       "请输入controller名称 如：user",
				Destination: &name,
				Required:    true,
			},
			&cli.StringFlag{
				Name:        "dir",
				Aliases:     []string{"d"},
				Value:       "",
				Usage:       "请输入controller目录 如：backend or backend/user",
				Destination: &dir,
				Required:    false,
			},
		},
		Before: func(ctx *cli.Context) error {
			return boots.InitConfig()
		},
		Action: func(ctx *cli.Context) error {
			return handleGenController(name, dir)
		},
	}
}

var _ command.Interface = (*GenController)(nil)

// handleGenController ...
func handleGenController(name, dir string) error {
	moduleName := helper.GetProjectModuleName()
	cmdName := strings.ToLower(name)
	cmdDir := strings.ToLower(dir)
	fileName := fmt.Sprintf("%s.go", cmdName)
	rootPath := vars.BasePath + ctlPath
	caseCmdName := strings.Title(helper.ToCamelCase(cmdName))
	// 创建目录
	if cmdDir != "" {
		rootPath += fmt.Sprintf("%s/", cmdDir)
		if err := helper.MakeMultiDir(rootPath); err != nil {
			return err
		}
		cmdDirs := strings.Split(cmdDir, "/")
		if len(cmdDirs) > 0 {
			ctlPkgName = cmdDirs[len(cmdDirs)-1]
		}
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
	t1 := template.Must(template.New("genCtlTpl").Parse(genCtlTpl))
	if err := t1.Execute(orPath, map[string]interface{}{
		"ImportPackage": moduleName,
		"PkgName":       ctlPkgName,
		"CtlName":       caseCmdName,
		"CmdDir":        cmdDir,
	}); err != nil {
		return err
	}

	fmt.Println(fmt.Sprintf("\u001B[34m%s\u001B[0m", fileName+" created successfully"))
	return nil
}
