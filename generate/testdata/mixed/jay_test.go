package mixed

import (
	"github.com/speedyhoon/rando"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLion(t *testing.T) {
	var expected, actual Lion
	require.NoError(t, actual.UnmarshalJ(expected.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, Lion{}, expected)
	require.Equal(t, Lion{}, actual)

	expected = Lion{
		B1:      rando.Bool(),
		Strings: rando.Strings(),
		B2:      rando.Bool(),
	}
	src := expected.MarshalJ()
	require.NoError(t, actual.UnmarshalJ(src))
	//require.NotEqual(t, Lion{}, expected)
	//require.NotEqual(t, Lion{}, actual)
	require.Equal(t, expected, actual)
}

func TestZebra(t *testing.T) {
	var expected, actual Zebra
	require.NoError(t, actual.UnmarshalJ(expected.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, Zebra{}, expected)
	require.Equal(t, Zebra{}, actual)

	expected = Zebra{
		B1:      rando.Bool(),
		Strings: rando.Strings(),
		B2:      rando.Bool(),
		Str:     rando.String(),
		Ints:    rando.Ints(),
		U64:     rando.Uint64(),
	}
	src := expected.MarshalJ()
	require.NoError(t, actual.UnmarshalJ(src))
	//require.NotEqual(t, Zebra{}, expected)
	//require.NotEqual(t, Zebra{}, actual)
	require.Equal(t, expected, actual)
}
