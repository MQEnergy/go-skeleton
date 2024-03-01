package command

import (
	_ "embed"
	"fmt"
	"os"
	"strings"
	"text/template"

	"go-skeleton/internal/bootstrap"
	"go-skeleton/internal/vars"
	"go-skeleton/pkg/command"
	"go-skeleton/pkg/helper"

	"github.com/urfave/cli/v2"
)

//go:embed tpls/gen_command.tpl
var genTpl string

//go:embed tpls/command.tpl
var cmdTpl string

var (
	commandPath = "/internal/command/"
	packageName = "command"
)

type GenCommand struct{}

func (g *GenCommand) Command() *cli.Command {
	var (
		name string
		dir  string
	)
	return &cli.Command{
		Name:  "genCommand",
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

var _ command.Interface = (*GenCommand)(nil)

// genCommand ...
func genCommand(name, dir string) error {
	cmdName := strings.ToLower(name)
	cmdDir := strings.ToLower(dir)
	moduleName := helper.GetProjectModuleName()
	fileName := fmt.Sprintf("%s.go", cmdName)
	rootPath := vars.BasePath + commandPath
	caseCmdName := strings.Title(helper.ToCamelCase(cmdName))

	// 创建目录
	if cmdDir != "" {
		rootPath += fmt.Sprintf("%s/", cmdDir)
		if err := helper.MakeMultiDir(rootPath); err != nil {
			return err
		}
		cmdDirs := strings.Split(cmdDir, "/")
		packageName = cmdDirs[len(cmdDirs)-1]
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
	t1 := template.Must(template.New("gentpl").Parse(genTpl))
	if err := t1.Execute(orPath, map[string]interface{}{
		"PkgName":       packageName,
		"CmdName":       caseCmdName,
		"ImportPackage": moduleName,
		"Name":          helper.ToCamelCase(cmdName),
		"Usage":         caseCmdName + "命令工具",
		"CmdDir":        cmdDir,
	}); err != nil {
		return err
	}
	// 修改command.go
	if err := handleComamnd(moduleName); err != nil {
		return err
	}

	fmt.Println(fmt.Sprintf("\u001B[34m%s\u001B[0m", fileName+" created successfully"))
	fmt.Println(fmt.Sprintf("\u001B[34m%s\u001B[0m", "run：go run cmd/cli/main.go "+helper.ToCamelCase(cmdName)))
	return nil
}

// handleComamnd ...
func handleComamnd(moduleName string) error {
	rootPath := vars.BasePath + commandPath
	cmdFileName := rootPath + "command.go"
	commandNames := make([]string, 0)
	importPackages := make([]string, 0)
	dirPaths, err := helper.GetFileNamesByDirPath(rootPath)
	if err != nil {
		return err
	}

	for _, item := range dirPaths {
		files := item["files"].([]string)
		path := item["path"].(string)
		filesList := make([]string, 0)
		if path == "tpls" || path == "test" || len(files) == 0 {
			continue
		}
		if path != "" {
			pathArr := strings.Split(path, "/")
			aliasName := pathArr[len(pathArr)-1] + helper.RandString(6)
			path = fmt.Sprintf("%s \"%s\"", aliasName, moduleName+commandPath+path)
			for _, s := range files {
				filesList = append(filesList, fmt.Sprintf("&%s.%s{}", aliasName, strings.Title(helper.ToCamelCase(s))))
			}
			importPackages = append(importPackages, path)
		} else {
			for _, s := range files {
				if s == "command" {
					continue
				}
				filesList = append(filesList, fmt.Sprintf("&%s{}", strings.Title(helper.ToCamelCase(s))))
			}
		}
		commandNames = append(commandNames, filesList...)
	}

	cmdFile, err := os.OpenFile(cmdFileName, os.O_CREATE|os.O_WRONLY, 0o666)
	if err != nil {
		return err
	}
	t2 := template.Must(template.New("cmdtpl").Parse(cmdTpl))
	return t2.Execute(cmdFile, map[string]interface{}{
		"ImportPackage":  moduleName,
		"Commands":       commandNames,
		"ImportPackages": importPackages,
	})
}
