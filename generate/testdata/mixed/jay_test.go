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
		One:   rando.Bool(),
		Two:   rando.Strings(),
		Three: rando.Bool(),
	}
	src := actual.MarshalJ()
	require.NoError(t, expected.UnmarshalJ(src))
	// require.NotEqual(t, Two{}, expected)
	// require.NotEqual(t, Two{}, actual)
	require.Equal(t, expected, actual)
}
