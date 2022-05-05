package spec

import (
	"fmt"
	"github.com/sierrasoftworks/humane-errors-go"
	"strconv"

	"github.com/amimof/huego"
)

func parseBrightness(s string, state *huego.State) error {
	pct, err := strconv.ParseFloat(s[:len(s)-1], 64)
	if err != nil {
		return humane.Wrap(
			err,
			fmt.Sprintf("Could not parse '%s' as a percentage.", s),
			"Make sure you have entered a valid percentage brightness in the form '57%'.",
		)
	}

	if pct < 0 || pct > 100 {
		return humane.New(
			fmt.Sprintf("'%s' is not a valid brightness between 0%% and 100%%.", s),
			"Make sure your brightness figure falls between 0% and 100%.",
		)
	}

	state.Bri = uint8(pct * 256 / 100)

	return nil
}
