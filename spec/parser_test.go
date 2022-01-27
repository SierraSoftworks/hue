package spec_test

import (
	"testing"

	"github.com/amimof/huego"
	"github.com/sierrasoftworks/hue/spec"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParser(t *testing.T) {
	tests := []struct {
		spec  string
		state huego.State
	}{
		{
			spec:  "",
			state: huego.State{On: true},
		},
		{
			spec:  "on",
			state: huego.State{On: true},
		},
		{
			spec:  "off",
			state: huego.State{On: false},
		},
		{
			spec:  "50%",
			state: huego.State{On: true, Bri: uint8(128)},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.spec, func(t *testing.T) {
			state, err := spec.ParseState(test.spec)
			require.NoError(t, err, "parsing should succeed")
			assert.Equal(t, test.state, state, "state should match")
		})
	}
}
