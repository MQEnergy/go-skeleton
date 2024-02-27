package {{.PkgName}}

import (
    "fmt"

	"{{.ImportPackage}}/internal/bootstrap"
    "{{.ImportPackage}}/pkg/config"

	"github.com/urfave/cli/v2"
)

// {{.CmdName}} ...
func {{.CmdName}}() *cli.Command {
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
			return {{.FuncName}}()
		},
	}
}

// {{.FuncName}} ...
func {{.FuncName}}() error {
    // Todo implement
    fmt.Println("{{.FuncName}}")
	return nil
}