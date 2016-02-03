package main

import (
	"os"

	"github.com/jaehoonkim/baggage/command"
	"github.com/mitchellh/cli"
)

var Commands map[string]cli.CommandFactory

func init() {
	ui := &cli.BasicUi{Writer: os.Stdout}

	Commands = map[string]cli.CommandFactory{
		"new": func() (cli.Command, error) {
			return &command.NewCommand{
				Ui: ui,
			}, nil
		},
	}

}
