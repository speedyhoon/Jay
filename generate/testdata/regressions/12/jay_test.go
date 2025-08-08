package fuzz12

import (
	"github.com/speedyhoon/rando"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFuzz_2(t *testing.T) {
	var expected, actual Two
	require.NoError(t, actual.UnmarshalJ(expected.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, Two{}, expected)
	require.Equal(t, Two{}, actual)

	expected = Two{
		One: [6]string{"Zero", "One", "Two", "Three", "Four", "Five"},
		Two: [6]string(rando.StringsQty(6)),
	}
	src := expected.MarshalJ()
	require.NoError(t, actual.UnmarshalJ(src))
	// require.NotEqual(t, Two{}, expected)
	// require.NotEqual(t, Two{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_3(t *testing.T) {
	var expected, actual Three
	require.NoError(t, actual.UnmarshalJ(expected.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, Three{}, expected)
	require.Equal(t, Three{}, actual)

	expected = Three{
		One: [6]string{
			// 6 + 58+49+59+11+21+10 = 214 / 566
			"abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUV",
			"abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLM",
			"abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVW",
			"abcdefghijk",
			"abcdefghijklmnopqrstu",
			"abcdefghij",
		},
		Two: [6]string{
			// 6 + 36+9+6+53+32+15	= 157 / 566
			"abcdefghijklmnopqrstuvwxyz0123456789",
			"abcdefghi",
			"abcdef",
			"abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQ",
			"abcdefghijklmnopqrstuvwxyz012345",
			"abcdefghijklmno",
		},
		Three: [6]string{
			// 6 + 53+32+45+45+14 = 195 / 566
			"abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQ",
			"abcdefghijklmnopqrstuvwxyz012345",
			"abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHI",
			"abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHI",
			"abcdefghijklmn",
		},
	}
	src := expected.MarshalJ()
	require.NoError(t, actual.UnmarshalJ(src), "actual:\n%+v\n\nsrc:\n%+v", actual, src)
	require.Equal(t, expected, actual)

	expected = Three{
		One:   [6]string(rando.StringsQty(6)),
		Two:   [6]string(rando.StringsQty(6)),
		Three: [6]string(rando.StringsQty(6)),
	}
	src = expected.MarshalJ()
	require.NoError(t, actual.UnmarshalJ(src), "actual:\n%+v\n\nsrc:\n%+v", actual, src)
	// require.NotEqual(t, Three{}, expected)
	// require.NotEqual(t, Three{}, actual)
	require.Equal(t, expected, actual)
}
