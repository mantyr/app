# App

[![Build Status](https://travis-ci.org/mantyr/starter.svg?branch=master)](https://travis-ci.org/mantyr/app)
[![GoDoc](https://godoc.org/github.com/mantyr/starter?status.png)](http://godoc.org/github.com/mantyr/app)
[![Go Report Card](https://goreportcard.com/badge/github.com/mantyr/app?v=1)][goreport]
[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg)](LICENSE.md)

Wrapper for github.com/urfave/cli/v2

This stable version.

## main.go

```go
package main

import (
	"os"
	"github.com/mantyr/app"

	"internal/commands/cli"
	"internal/commands/server"
)

var (
	versionInfo = "dev"
)

func main() {
	app := app.New()
	app.Name = "name"
	app.Usage = "usage"
	app.Version = versionInfo
	app.Description = "description"
	app.Register(
		cli.New(),
		server.New(),
	)
	app.RunAndFatal(os.Args)
}
```

## internal/commands/cli

```go
package cli

import (
	"github.com/mantyr/app"
	"github.com/mantyr/app/commands"

	"internal/commands/cli/item1"
	"internal/commands/cli/item2"
)

type Command struct {
	commands.Command
}

func New() *Command {
	return &Command{}
}

func (c *Command) Init() error {
	c.Command.Name = "cli"
	c.Command.Usage = "Command Line Interface"
	c.Command.Description = "description"
	return nil
}

func (c *Command) Subcommands() []app.Command {
	return []app.Command{
		item1.New(),
		item2.New(),
	}
}
```

## Author

[Oleg Shevelev][mantyr]

[mantyr]: https://github.com/mantyr

[build_status]: https://travis-ci.org/mantyr/app
[godoc]:        http://godoc.org/github.com/mantyr/app
[goreport]:     https://goreportcard.com/report/github.com/mantyr/app
