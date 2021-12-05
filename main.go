package main

import (
	"fmt"
	"os"

	"github.com/sierrasoftworks/hue/commands"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "hue",
		Usage: "bedroom=orange,30%",

		Commands: commands.GetCommands(),
		Action:   commands.Set,
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
}
