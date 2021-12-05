package commands

import "github.com/urfave/cli/v2"

func GetCommands() []*cli.Command {
	return []*cli.Command{
		&alias_command,
		&list_command,
		&setup_command,
	}
}
