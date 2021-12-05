package commands

import (
	"fmt"
	"strings"

	"github.com/amimof/huego"
	"github.com/urfave/cli/v2"
	"github.com/zalando/go-keyring"
)

var setup_command = cli.Command{
	Name:  "setup",
	Usage: "Setup this app to connect to your Hue bridge.",
	Action: func(c *cli.Context) error {
		bridge, err := huego.Discover()
		if err != nil {
			return err
		}

		fmt.Printf("Discovered a Hue bridge at %s\n", bridge.Host)

		// Prompt the user to choose [Y]/n to connect to this bridge
		fmt.Printf("Connect to this bridge? [Y/n]: ")
		var answer string
		fmt.Scanln(&answer)
		if strings.ToLower(answer) == "n" {
			return nil
		}

		fmt.Println("Requesting access to your Hue bridge, please press the button on it to continue...")
		fmt.Printf("Press ENTER when you are ready.")
		fmt.Scanln(&answer)
		user, err := bridge.CreateUser("sierrasoftworks/hue")
		if err != nil {
			return err
		}

		if err := keyring.Set("com.sierrasoftworks.hue", "$default_host$", bridge.Host); err != nil {
			return err
		}

		if err := keyring.Set("com.sierrasoftworks.hue", bridge.Host, user); err != nil {
			return err
		}

		fmt.Println("Successfully connected to your Hue bridge.")

		return nil
	},
}
