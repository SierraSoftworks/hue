package spec

import (
	"fmt"
	"strings"

	"github.com/amimof/huego"
	"github.com/sierrasoftworks/hue/humanerrors"
)

func ParseState(spec string) (huego.State, error) {
	state := huego.State{On: true}

	for _, s := range strings.Split(spec, ",") {
		switch s {
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
			return state, humanerrors.New(
				fmt.Sprintf("'%s' is not a valid light state.", s),
				"Try 'on', 'off', or a percentage brightness between 0% and 100%.",
			)
		}
	}

	return state, nil
}
