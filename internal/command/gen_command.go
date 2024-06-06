package command

import (
	_ "embed"
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/MQEnergy/go-skeleton/internal/vars"
	"github.com/MQEnergy/go-skeleton/pkg/command"
	"github.com/MQEnergy/go-skeleton/pkg/helper"

	"github.com/urfave/cli/v2"
)

//go:embed tpls/gen_command.tpl
var genTpl string

//go:embed tpls/command.tpl
var cmdTpl string

var (
	commandPath = "/internal/command/"
	commandName = "command"
)

type GenCommand struct{}

func (g *GenCommand) Command() *cli.Command {
	var (
		name    string
		dir     string
		service string
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
			&cli.StringFlag{
				Name:        "service",
				Aliases:     []string{"s"},
				Value:       "",
				Usage:       "加载其他自定义服务(已经开启的服务) 如：mysql、redis ... 格式：以英文逗号相隔 如：mysql,redis",
				Destination: &service,
				Required:    false,
			},
		},
		Action: func(c *cli.Context) error {
			return genCommand(name, dir, service)
		},
	}
}

var _ command.Interface = (*GenCommand)(nil)

// genCommand ...
func genCommand(name, dir, service string) error {
	cmdName := strings.ToLower(name)
	cmdDir := strings.ToLower(dir)
	services := strings.Split(service, ",")
	serviceList := make([]string, 0)
	if len(services) > 0 {
		for _, s := range services {
			if strings.TrimSpace(s) != "" {
				serviceName := strings.Title(helper.ToCamelCase(strings.TrimSpace(s))) + "Service"
				serviceList = append(serviceList, serviceName)
			}
		}
	}
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
		if len(cmdDirs) > 0 {
			commandName = cmdDirs[len(cmdDirs)-1]
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
	t1 := template.Must(template.New("gentpl").Parse(genTpl))
	if err := t1.Execute(orPath, map[string]interface{}{
		"PkgName":       commandName,
		"CmdName":       caseCmdName,
		"ImportPackage": moduleName,
		"Name":          helper.ToCamelCase(cmdName),
		"Usage":         caseCmdName + "命令工具",
		"Services":      serviceList,
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
			aliasName := helper.RandString(6)
			if len(pathArr) > 0 {
				aliasName = pathArr[len(pathArr)-1] + aliasName
			}
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

	cmdFile, err := os.Create(cmdFileName)
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
