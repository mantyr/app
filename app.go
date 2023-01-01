package app

import (
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

type App struct {
	Name string
	Usage string
	Version string
	Description string

	app *cli.App
	err error
}

func New() *App {
	return &App{
		app: cli.NewApp(),
	}
}

func (a *App) Register(commands ...Command) {
	for _, command := range commands {
		cmd, err := Convert(command)
		if err != nil {
			a.err = err
			return
		}
		a.app.Commands = append(a.app.Commands, cmd)
	}
}

func (a *App) RunAndFatal(args []string) {
	if a.err != nil {
		log.Fatal(a.err)
		return
	}
	a.app.Name = a.Name
	a.app.Usage = a.Usage
	a.app.Version = a.Version
	a.app.Description = a.Description

	err := a.app.Run(args)
	if err != nil {
		log.Fatal(err)
	}
}
