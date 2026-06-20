package main

import (
	"github.com/sierrasoftworks/humane-errors-go"
	"os"

	"github.com/sierrasoftworks/hue/commands"
	"github.com/urfave/cli/v2"
)

// version is overridden at build time via -ldflags "-X main.version=...".
var version = "dev"

func main() {
	app := &cli.App{
		Name:        "hue",
		Version:     version,
		Usage:       "Set your light states",
		Description: "Control your Philips Hue lights from your command line.",
		Authors: []*cli.Author{
			{Name: "Benjamin Pannell", Email: "admin@sierrasoftworks.com"},
		},
		ArgsUsage: "all=off bedroom=orange,30%",

		Commands: commands.GetCommands(),
		Action:   commands.Set,
	}

	err := app.Run(os.Args)
	humane.Eprint(err)
}
