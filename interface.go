package app

import (
	"github.com/urfave/cli/v2"
	"github.com/mantyr/app/commands"
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
