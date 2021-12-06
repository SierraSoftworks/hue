package spec_test

import (
	"testing"

	"github.com/sierrasoftworks/hue/spec"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBrightness(t *testing.T) {
	state, err := spec.ParseState("50%")
	require.NoError(t, err, "parsing should succeed")
	assert.Equal(t, state.Bri, uint8(128), "brightness should be 50%")
}
