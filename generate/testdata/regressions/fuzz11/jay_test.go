package fuzz11

import (
	"github.com/speedyhoon/rando"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFuzz35690(t *testing.T) {
	var expected, actual Fuzz35690
	require.NoError(t, actual.UnmarshalJ(expected.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, Fuzz35690{}, expected)
	require.Equal(t, Fuzz35690{}, actual)

	// Hardcoded values.
	expected = Fuzz35690{
		One:   []complex64{complex(1, 0), complex(2, 0), complex(3, 0), complex(4, 0)},
		Two:   "alphabet",
		Three: []string{"zebra", "lion", "elephant", "tiger"},
	}
	src := expected.MarshalJ()
	require.NoError(t, actual.UnmarshalJ(src))
	require.Equal(t, expected, actual, "%c\nbytes:\n%[1]v\n", src)

	// Random values.
	actual = Fuzz35690{
		One:   rando.Complex64s(),
		Two:   rando.String(),
		Three: rando.Strings(),
	}
	src = actual.MarshalJ()
	require.NoError(t, expected.UnmarshalJ(src))
	require.Equal(t, expected, actual)
}
