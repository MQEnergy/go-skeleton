package {{.PkgName}}

import (
    "fmt"

	"{{.ImportPackage}}/internal/bootstrap"
    "{{.ImportPackage}}/pkg/command"
    "{{.ImportPackage}}/pkg/config"

	"github.com/urfave/cli/v2"
)

type {{.CmdName}} struct {}

// Command ...
func (g *{{.CmdName}}) Command() *cli.Command {
	return &cli.Command{
        Name:  "{{.Name}}",
        Usage: "{{.Usage}}",
        Flags: []cli.Flag{
            &cli.StringFlag{
                Name:        "env",
                Aliases:     []string{"e"},
                Value:       "dev",
                Usage:       "请选择配置文件 [dev | test | prod]",
                Destination: &config.ConfEnv,
            },
		},
        Before: func(ctx *cli.Context) error {
            return bootstrap.InitConfig()
        },
		Action: func(ctx *cli.Context) error {
			return handle{{.CmdName}}()
		},
	}
}

var _ command.Interface = (*{{.CmdName}})(nil)

// handle{{.CmdName}} ...
func handle{{.CmdName}}() error {
    // Todo implement ...
    fmt.Println("handle{{.CmdName}}")
	return nil
}