package main

import (
	"fmt"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/urfave/cli/v2"
	"go-skeleton/internal/app/pkg/validator"
	"go-skeleton/internal/bootstrap"
	"go-skeleton/internal/router"
	"go-skeleton/pkg/config"
	"os"
	"runtime"
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
	// https://patorjk.com/software/taag/#p=testall&v=1&f=ANSI%20Shadow&t=O2O-AMQP%20
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
	buildInfo := fmt.Sprintf("%s-%s-%s", runtime.GOOS, runtime.GOARCH, gtime.Now())

	return &cli.App{
		Name:    AppName,
		Version: buildInfo,
		Usage:   AppUsage,
		Authors: Authors,
		Flags: []cli.Flag{
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
		},
		Action: func(context *cli.Context) error {
			fmt.Println(fmt.Sprintf("\u001B[34m%s\u001B[0m", _UI))
			//	初始化服务
			bootstrap.BootService()
			//	引入验证翻译器
			validator.NewValidate()
			//	注册路由 启动程序
			return router.Register().Listen(":" + AppPort)
		},
		Commands: []*cli.Command{},
	}
}

func main() {
	if err := Stack().Run(os.Args); err != nil {
		panic(err)
	}
}
