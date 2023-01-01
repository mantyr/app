package app

import (
	"github.com/mantyr/app/commands"
	"github.com/urfave/cli/v2"
)

type Command interface {
	Init() error
	GetCommand() commands.Command
}

type Action interface {
	Action(ctx *cli.Context) error
}

type Subcommands interface {
	Subcommands() []Command
}
