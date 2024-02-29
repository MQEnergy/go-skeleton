package main

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"go-skeleton/internal/command"
	"go-skeleton/pkg/config"

	"github.com/urfave/cli/v2"
)

var (
	AppName  = "skeleton-cli"
	AppUsage = "命令行工具集"
	Authors  = []*cli.Author{
		{
			Name:  "chenxi",
			Email: "bbxycx.18@163.com",
		},
	}
	AppPort string // port
	// https://patorjk.com/software/taag/#p=testall&v=1&f=ANSI%20Shadow&t=skeleton-cli%20
	_UI = `
███████╗██╗  ██╗███████╗██╗     ███████╗████████╗ ██████╗ ███╗   ██╗       ██████╗██╗     ██╗
██╔════╝██║ ██╔╝██╔════╝██║     ██╔════╝╚══██╔══╝██╔═══██╗████╗  ██║      ██╔════╝██║     ██║
███████╗█████╔╝ █████╗  ██║     █████╗     ██║   ██║   ██║██╔██╗ ██║█████╗██║     ██║     ██║
╚════██║██╔═██╗ ██╔══╝  ██║     ██╔══╝     ██║   ██║   ██║██║╚██╗██║╚════╝██║     ██║     ██║
███████║██║  ██╗███████╗███████╗███████╗   ██║   ╚██████╔╝██║ ╚████║      ╚██████╗███████╗██║
╚══════╝╚═╝  ╚═╝╚══════╝╚══════╝╚══════╝   ╚═╝    ╚═════╝ ╚═╝  ╚═══╝       ╚═════╝╚══════╝╚═╝
`
)

// Stack 程序运行前的处理
func Stack() *cli.App {
	buildInfo := fmt.Sprintf("%s-%s-%s", runtime.GOOS, runtime.GOARCH, time.Now().Format(time.DateTime))

	app := &cli.App{
		Name:    AppName,
		Version: buildInfo,
		Usage:   AppUsage,
		Authors: Authors,
	}
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "env",
			Aliases:     []string{"e"},
			Value:       "dev",
			Usage:       "请选择配置文件 [dev | test | prod]",
			Destination: &config.ConfEnv,
		},
	}
	app.Commands = command.New().WithCommands()

	app.Action = func(ctx *cli.Context) error {
		fmt.Println(fmt.Sprintf("\u001B[34m%s\u001B[0m", _UI))
		fmt.Println(`一、开发模式
  1、执行 go run cmd/cli/main.go -h 查看命令集

二、生产模式
  1、make build 打包命令行工具
  2、执行 ./releases/skeleton-cli -h 查看命令行
`)
		return nil
	}
	return app
}

func main() {
	if err := Stack().Run(os.Args); err != nil {
		panic(err)
	}
}
