package app

import (
	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
)

func Convert(cmd Command) (*cli.Command, error) {
	err := cmd.Init()
	if err != nil {
		return nil, err
	}
	tmp := cmd.GetCommand()
	
	c := &cli.Command{
		Name:        tmp.Name,
		Usage:       tmp.Usage,
		Description: tmp.Description,
		Flags:       tmp.Flags,
	}
	if tmp.EnableConfigFile {
		c.Flags = append(c.Flags, &cli.StringFlag{
			Name:  "config",
			Value: "./config.yml",
			Usage: "Адрес yml файла с настройками",
		})
		c.Before = altsrc.InitInputSourceWithContext(c.Flags, altsrc.NewYamlSourceFromFlagFunc("config"))
	}
	if tmp, ok := interface{}(cmd).(Action); ok {
		c.Action = tmp.Action
	}
	if tmp, ok := interface{}(cmd).(Subcommands); ok {
		for _, item := range tmp.Subcommands() {
			command, err := Convert(item)
			if err != nil {
				return nil, err
			}
			c.Subcommands = append(c.Subcommands, command)
		}
	}
	return c, nil
}
