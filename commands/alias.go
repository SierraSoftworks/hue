package commands

import (
	"fmt"
	"strings"

	"github.com/sierrasoftworks/hue/config"
	"github.com/urfave/cli/v2"
)

var alias_command = cli.Command{
	Name:        "alias",
	Usage:       "Assign aliases for groups of lights.",
	Description: "Configures an alias for a group of lights which should be changed together.",
	Action: func(c *cli.Context) error {
		cfg, err := config.Load()
		if err != nil {
			return err
		}

		if c.NArg() == 0 {
			for alias, lights := range cfg.Aliases {
				fmt.Printf("%s=%s\n", alias, lights)
			}

			return nil
		}

		for _, arg := range c.Args().Slice() {
			parts := strings.SplitN(arg, "=", 2)
			if len(parts) != 2 {
				fmt.Printf("%s=%s\n", parts[0], cfg.Aliases[parts[0]])
				continue
			}

			cfg.Aliases[parts[0]] = parts[1]
		}

		err = config.Save(cfg)
		if err != nil {
			return err
		}

		return nil
	},
}
