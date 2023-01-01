package commands

import (
	"github.com/urfave/cli/v2"
)

// Command это объект описания команды для встраивания
type Command struct {
	Name        string
	Usage       string
	Description string
	Flags       []cli.Flag

	// EnableConfigFile включает --config флаг с адресом файла в формате YML
	EnableConfigFile bool
}

func (c *Command) GetCommand() Command {
	return *c
}

func (c *Command) AddFlags(flags ...cli.Flag) {
	c.Flags = append(c.Flags, flags...)
}
