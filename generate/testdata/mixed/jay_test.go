package mixed

import (
	"github.com/speedyhoon/rando"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLion(t *testing.T) {
	var expected, actual Lion
	require.NoError(t, expected.UnmarshalJ(actual.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, Lion{}, expected)
	require.Equal(t, Lion{}, actual)

	actual = Lion{
		B1:      rando.Bool(),
		Strings: rando.Strings(),
		B2:      rando.Bool(),
	}
	src := actual.MarshalJ()
	require.NoError(t, expected.UnmarshalJ(src))
	//require.NotEqual(t, Lion{}, expected)
	//require.NotEqual(t, Lion{}, actual)
	require.Equal(t, expected, actual)
}

func TestZebra(t *testing.T) {
	var expected, actual Zebra
	require.NoError(t, expected.UnmarshalJ(actual.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, Zebra{}, expected)
	require.Equal(t, Zebra{}, actual)

	actual = Zebra{
		B1:      rando.Bool(),
		Strings: rando.Strings(),
		B2:      rando.Bool(),
		Str:     rando.String(),
		Ints:    rando.Ints(),
		U64:     rando.Uint64(),
	}
	src := actual.MarshalJ()
	require.NoError(t, expected.UnmarshalJ(src))
	//require.NotEqual(t, Zebra{}, expected)
	//require.NotEqual(t, Zebra{}, actual)
	require.Equal(t, expected, actual)
}
