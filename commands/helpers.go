package commands

import (
	"fmt"
	"github.com/sierrasoftworks/humane-errors-go"
	"strconv"
	"strings"

	"github.com/amimof/huego"
)

func getTargets(targets []string, bridge *huego.Bridge) ([]int, error) {
	lights := map[int]struct{}{}

	var allGroups []huego.Group
	var allLights []huego.Light

	for _, arg := range targets {
		if arg == "all" {
			if allLights == nil {
				lights, err := bridge.GetLights()
				if err != nil {
					return nil, humane.Wrap(
						err,
						"Failed to retrieve lights from your Hue bridge.",
						"Make sure that your Hue bridge is online and reachable from this device.",
						"Make sure that you have setup your Hue CLI to talk to your bridge by running `hue setup`.",
						"Make sure that your Hue bridge is configured correctly.",
					)
				}

				allLights = lights
			}

			for _, light := range allLights {
				lights[light.ID] = struct{}{}
			}
		} else if strings.HasPrefix(arg, "#") {
			id, err := strconv.Atoi(arg[1:])
			if err != nil {
				return nil, err
			}
			lights[id] = struct{}{}
		} else {
			if allGroups == nil {
				groups, err := bridge.GetGroups()
				if err != nil {
					return nil, humane.Wrap(
						err,
						"Failed to retrieve light groups from your Hue bridge.",
						"Make sure that your Hue bridge is online and reachable from this device.",
						"Make sure that you have setup your Hue CLI to talk to your bridge by running `hue setup`.",
						"Make sure that your Hue bridge is configured correctly.",
					)
				}

				allGroups = groups
			}

			found := false
			for _, group := range allGroups {
				if strings.EqualFold(group.Name, arg) {
					found = true
					for _, light := range group.Lights {
						id, err := strconv.Atoi(light)
						if err != nil {
							return nil, humane.Wrap(
								err,
								fmt.Sprintf("group %s contains invalid light id %s", group.Name, light),
								"Make sure that your Hue bridge is configured correctly.",
								"Make sure that you're running the latest version of the Hue CLI.",
								"Report this error to the Hue CLI maintainers on GitHub.",
							)
						}

						lights[id] = struct{}{}
					}
					break
				}
			}

			if !found {
				return nil, humane.New(
					fmt.Sprintf("Could not find the Hue light group '%s'", arg),
					"Run `hue list` to view the full list of groups supported by your Hue bridge.",
					"If you are trying to turn on a specific light, remember to prefix its ID with the '#' symbol (i.e. #5).",
					"Ensure that group names with special characters and/or spaces are \"wrapped in quotes\".",
				)
			}
		}
	}

	out := []int{}
	for id := range lights {
		out = append(out, id)
	}
	return out, nil
}
