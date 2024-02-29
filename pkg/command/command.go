package command

import "github.com/urfave/cli/v2"

type Interface interface {
	Command() *cli.Command
}

type CommonInterface interface {
	RegisterCmds() []Interface
}
