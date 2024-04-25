package {{.PkgName}}

import (
    "fmt"

    {{- if eq (len .Services) 0}}
	"{{.ImportPackage}}/internal/bootstrap/boots"
	{{- else}}
    "{{.ImportPackage}}/internal/bootstrap"
	{{- end}}
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
            {{- if eq (len .Services) 0}}
            return boots.InitConfig()
            {{- else}}
            bootstrap.BootService(
            {{- range $item := .Services}}
                bootstrap.{{$item}},
            {{- end}}
            )
            return nil
            {{- end}}
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