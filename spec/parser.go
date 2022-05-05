package spec

import (
	"fmt"
	"github.com/sierrasoftworks/humane-errors-go"
	"strings"

	"github.com/amimof/huego"
)

func ParseState(spec string) (huego.State, error) {
	state := huego.State{On: true}

	for _, s := range strings.Split(spec, ",") {
		switch s {
		case "":
			continue
		case "on":
			return huego.State{On: true}, nil
		case "off":
			return huego.State{On: false}, nil
		}

		s = resolveKnownColourAlias(s)

		if s[0] == '#' {
			err := parseHexColour(s, &state)
			if err != nil {
				return state, err
			}
		} else if strings.HasSuffix(s, "%") {
			err := parseBrightness(s, &state)
			if err != nil {
				return state, err
			}
		} else {
			return state, humane.New(
				fmt.Sprintf("'%s' is not a valid light state.", s),
				"Try 'on', 'off', or a percentage brightness between 0% and 100%.",
			)
		}
	}

	return state, nil
}
