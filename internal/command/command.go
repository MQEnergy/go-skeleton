package command

import (
	_ "embed"
	"fmt"
	"strings"
	"text/template"

	"go-skeleton/internal/bootstrap"
	"go-skeleton/internal/vars"

	"github.com/urfave/cli/v2"
	"go-skeleton/pkg/helper"
)

//go:embed tpls/command.tpl
var commandTemplate string

// CommandCmd 创建command工具
func CommandCmd() *cli.Command {
	var (
		name string
		dir  string
	)
	return &cli.Command{
		Name:  "gen:command",
		Usage: "Create a new command",
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
		Before: func(c *cli.Context) error {
			return bootstrap.InitConfig()
		},
		Action: func(c *cli.Context) error {
			return genCommand(name, dir)
		},
	}
}

func genCommand(name, dir string) error {
	cmdName := strings.ToLower(name)
	cmdDir := strings.ToLower(dir)
	projectModuleName := helper.GetProjectModuleName()
	fileName := fmt.Sprintf("%s.go", cmdName) // 生成的文件名
	filePath := vars.BasePath + "/internal/command/"
	packageName := "command"

	t1 := template.Must(template.New("commandTemplate").Parse(commandTemplate))

	// 创建目录
	if cmdDir != "" {
		filePath += fmt.Sprintf("%s/", cmdDir)
		if err := helper.MakeMultiDir(filePath); err != nil {
			return err
		}
		cmdDirs := strings.Split(cmdDir, "/")
		packageName = cmdDirs[len(cmdDirs)-1]
	}
	// 判断文件是否存在
	if flag := helper.IsPathExist(filePath + fileName); flag {
		fmt.Println(fmt.Sprintf("\x1b[31m%s\x1b[0m", cmdName+".go already existed"))
		return nil
	}
	// 创建文件
	orPath, err := helper.MakeFileOrPath(filePath + fileName)
	if err != nil {
		return err
	}
	defer orPath.Close()
	if err := t1.Execute(orPath, map[string]interface{}{
		"PkgName":       packageName,
		"CmdName":       strings.Title(cmdName + "Cmd"),
		"ImportPackage": projectModuleName,
		"Name":          cmdName,
		"Usage":         strings.Title(cmdName) + "命令工具",
		"FuncName":      "gen" + strings.Title(cmdName),
	}); err != nil {
		return err
	}
	fmt.Println(fmt.Sprintf("\u001B[34m%s\u001B[0m", cmdName+".go create success"))
	fmt.Println(fmt.Sprintf("\u001B[34m%s\u001B[0m", "1、需要在cmd/cli/main.go的app.Commands中引用如下方法："+packageName+"."+strings.Title(cmdName)+"Cmd()"))
	fmt.Println(fmt.Sprintf("\u001B[34m%s\u001B[0m", "2、查看帮助：go run cmd/cli/main.go "+cmdName+" -h"))
	return nil
}
