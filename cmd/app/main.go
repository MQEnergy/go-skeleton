package main

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/MQEnergy/go-skeleton/internal/bootstrap"
	"github.com/MQEnergy/go-skeleton/internal/router"
	"github.com/MQEnergy/go-skeleton/pkg/config"
	"github.com/urfave/cli/v2"
)

var (
	AppName  = "go-skeleton"
	AppUsage = "基于Go语言和fiber框架的高性能高并发的Web项目骨架"
	Authors  = []*cli.Author{
		{
			Name:  "chenxi",
			Email: "bbxycx.18@163.com",
		},
	}
	AppPort string // port
	// https://patorjk.com/software/taag/#p=testall&v=1&f=ANSI%20Shadow&t=go-skeleton%20
	_UI = `
 ██████╗  ██████╗       ███████╗██╗  ██╗███████╗██╗     ███████╗████████╗ ██████╗ ███╗   ██╗
██╔════╝ ██╔═══██╗      ██╔════╝██║ ██╔╝██╔════╝██║     ██╔════╝╚══██╔══╝██╔═══██╗████╗  ██║
██║  ███╗██║   ██║█████╗███████╗█████╔╝ █████╗  ██║     █████╗     ██║   ██║   ██║██╔██╗ ██║
██║   ██║██║   ██║╚════╝╚════██║██╔═██╗ ██╔══╝  ██║     ██╔══╝     ██║   ██║   ██║██║╚██╗██║
╚██████╔╝╚██████╔╝      ███████║██║  ██╗███████╗███████╗███████╗   ██║   ╚██████╔╝██║ ╚████║
 ╚═════╝  ╚═════╝       ╚══════╝╚═╝  ╚═╝╚══════╝╚══════╝╚══════╝   ╚═╝    ╚═════╝ ╚═╝  ╚═══╝
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
		&cli.StringFlag{
			Name:        "port",
			Aliases:     []string{"p"},
			Value:       "9527",
			Usage:       "请选择启动端口",
			Destination: &AppPort,
		},
	}
	app.Action = func(ctx *cli.Context) error {
		fmt.Println(fmt.Sprintf("\u001B[34m%s\u001B[0m", _UI))
		// bootstrap service
		bootstrap.BootService()
		// register routes and listen port
		return router.Register(AppName).Listen(":" + AppPort)
	}
	return app
}

func main() {
	if err := Stack().Run(os.Args); err != nil {
		panic(err)
	}
}
