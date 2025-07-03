package fuzz10

import (
	"github.com/speedyhoon/rando"
	"github.com/stretchr/testify/require"
	"math"
	"testing"
)

func TestFuzz710(t *testing.T) {
	var expected, actual Fuzz710
	require.NoError(t, expected.UnmarshalJ(actual.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, Fuzz710{}, expected)
	require.Equal(t, Fuzz710{}, actual)

	// Hardcoded values.
	actual = Fuzz710{
		One:   complex(math.MaxFloat32, math.MaxFloat32),
		Two:   "alphabet",
		Three: []string{"zebra", "lion", "elephant", "tiger"},
	}
	src := actual.MarshalJ()
	require.NoError(t, expected.UnmarshalJ(src))
	require.Equal(t, expected, actual)

	// Random values.
	actual = Fuzz710{
		One:   rando.Complex64(),
		Two:   rando.String(),
		Three: rando.Strings(),
	}
	src = actual.MarshalJ()
	require.NoError(t, expected.UnmarshalJ(src))
	require.Equal(t, expected, actual)
}
