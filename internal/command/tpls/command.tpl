// Code generated by go-skeleton.
// Code generated by go-skeleton.
// Code generated by go-skeleton.

package command

import (

    {{- range $item := .ImportPackages}}
    {{$item -}}
    {{- end}}
    "{{.ImportPackage}}/pkg/command"

	"github.com/urfave/cli/v2"
)

type Command struct {
	cmd *cli.Command
}

func (c *Command) RegisterCmds() []command.Interface {
	return []command.Interface{
	    {{- range $item := .Commands}}
        {{$item -}},
	    {{- end}}
    }
}

func New() *Command {
	return &Command{
		cmd: &cli.Command{},
	}
}

func (c *Command) WithCommands() []*cli.Command {
	commands := c.RegisterCmds()
	var clis []*cli.Command
	for _, command := range commands {
		clis = append(clis, command.Command())
	}
	return clis
}

var _ command.CommonInterface = (*Command)(nil)
