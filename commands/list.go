package commands

import (
	"fmt"
	"strings"

	"github.com/urfave/cli/v2"
)

var list_command = cli.Command{
	Name:  "list",
	Usage: "List your lights and light groups",
	Action: func(c *cli.Context) error {
		bridge, err := GetBridge()
		if err != nil {
			return err
		}

		groups, err := bridge.GetGroups()
		if err != nil {
			return err
		}

		if len(groups) > 0 {
			fmt.Println("Groups:")
		}
		for _, group := range groups {
			fmt.Printf("- %s: %s\n", group.Name, strings.Join(group.Lights, ", "))
		}

		fmt.Println("Lights:")
		lights, err := bridge.GetLights()
		if err != nil {
			return err
		}

		for _, light := range lights {
			fmt.Printf("- #%d: %s (%s)\n", light.ID, light.Name, ToHuman(light.State))
		}

		return nil
	},
}
