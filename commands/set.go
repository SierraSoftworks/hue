package commands

import (
	"fmt"
	"strings"

	"github.com/sierrasoftworks/hue/config"
	"github.com/sierrasoftworks/hue/humanerrors"
	"github.com/sierrasoftworks/hue/spec"
	"github.com/urfave/cli/v2"
)

func Set(c *cli.Context) error {
	if c.NArg() == 0 {
		return c.App.Command("help").Run(c)
	}

	bridge, err := GetBridge()
	if err != nil {
		return err
	}

	cfg, err := config.Load()
	if err != nil {
		return err
	}

	for _, set := range c.Args().Slice() {
		parts := strings.SplitN(set, "=", 2)
		if len(parts) != 2 {
			return humanerrors.New(
				fmt.Sprintf("'%s' is not a valid target and state specifier.", set),
				"Try setting a light's state to on with '#1=on'",
				"Try setting the brightness of a light with 'bedroom=30%'",
				"Try turning all your lights off with all=off",
			)
		}

		target := parts[0]
		if aliased, ok := cfg.Aliases[target]; ok {
			target = aliased
		}

		lights, err := getTargets(strings.Split(target, ","), bridge)
		if err != nil {
			return err
		}

		state, err := spec.ParseState(parts[1])
		if err != nil {
			return err
		}

		for _, light := range lights {
			bridge.SetLightState(light, state)
		}
	}

	return nil
}
