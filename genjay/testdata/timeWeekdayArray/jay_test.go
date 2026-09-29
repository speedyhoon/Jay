package main

import (
	"testing"
	"time"

	"github.com/go-openapi/testify/v2/require"
	"github.com/speedyhoon/rando"
)

func TestFuzz_1(t *testing.T) {
	var expected, actual One
	require.NoError(t, actual.UnmarshalJ(expected.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, One{}, expected)
	require.Equal(t, One{}, actual)

	expected = One{
		One: [21]time.Weekday{
			0:  time.Weekday(rando.Int()),
			1:  time.Weekday(rando.Int()),
			2:  time.Weekday(rando.Int()),
			3:  time.Weekday(rando.Int()),
			4:  time.Weekday(rando.Int()),
			5:  time.Weekday(rando.Int()),
			6:  time.Weekday(rando.Int()),
			7:  time.Weekday(rando.Int()),
			8:  time.Weekday(rando.Int()),
			9:  time.Weekday(rando.Int()),
			10: time.Weekday(rando.Int()),
			11: time.Weekday(rando.Int()),
			12: time.Weekday(rando.Int()),
			13: time.Weekday(rando.Int()),
			14: time.Weekday(rando.Int()),
			15: time.Weekday(rando.Int()),
			16: time.Weekday(rando.Int()),
			17: time.Weekday(rando.Int()),
			18: time.Weekday(rando.Int()),
			19: time.Weekday(rando.Int()),
			20: time.Weekday(rando.Int()),
		},
	}
	src := expected.MarshalJ()
	require.NoError(t, actual.UnmarshalJ(src))
	// require.NotEqual(t, One{}, expected)
	// require.NotEqual(t, One{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_2(t *testing.T) {
	var expected, actual Two
	require.NoError(t, actual.UnmarshalJ(expected.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, Two{}, expected)
	require.Equal(t, Two{}, actual)

	expected = Two{
		One: [20]time.Weekday{
			0:  time.Weekday(rando.Int()),
			1:  time.Weekday(rando.Int()),
			2:  time.Weekday(rando.Int()),
			3:  time.Weekday(rando.Int()),
			4:  time.Weekday(rando.Int()),
			5:  time.Weekday(rando.Int()),
			6:  time.Weekday(rando.Int()),
			7:  time.Weekday(rando.Int()),
			8:  time.Weekday(rando.Int()),
			9:  time.Weekday(rando.Int()),
			10: time.Weekday(rando.Int()),
			11: time.Weekday(rando.Int()),
			12: time.Weekday(rando.Int()),
			13: time.Weekday(rando.Int()),
			14: time.Weekday(rando.Int()),
			15: time.Weekday(rando.Int()),
			16: time.Weekday(rando.Int()),
			17: time.Weekday(rando.Int()),
			18: time.Weekday(rando.Int()),
			19: time.Weekday(rando.Int()),
		},
		Two: [19]time.Weekday{
			0:  time.Weekday(rando.Int()),
			1:  time.Weekday(rando.Int()),
			2:  time.Weekday(rando.Int()),
			3:  time.Weekday(rando.Int()),
			4:  time.Weekday(rando.Int()),
			5:  time.Weekday(rando.Int()),
			6:  time.Weekday(rando.Int()),
			7:  time.Weekday(rando.Int()),
			8:  time.Weekday(rando.Int()),
			9:  time.Weekday(rando.Int()),
			10: time.Weekday(rando.Int()),
			11: time.Weekday(rando.Int()),
			12: time.Weekday(rando.Int()),
			13: time.Weekday(rando.Int()),
			14: time.Weekday(rando.Int()),
			15: time.Weekday(rando.Int()),
			16: time.Weekday(rando.Int()),
			17: time.Weekday(rando.Int()),
			18: time.Weekday(rando.Int()),
		},
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
		One: [18]time.Weekday{
			0:  time.Weekday(rando.Int()),
			1:  time.Weekday(rando.Int()),
			2:  time.Weekday(rando.Int()),
			3:  time.Weekday(rando.Int()),
			4:  time.Weekday(rando.Int()),
			5:  time.Weekday(rando.Int()),
			6:  time.Weekday(rando.Int()),
			7:  time.Weekday(rando.Int()),
			8:  time.Weekday(rando.Int()),
			9:  time.Weekday(rando.Int()),
			10: time.Weekday(rando.Int()),
			11: time.Weekday(rando.Int()),
			12: time.Weekday(rando.Int()),
			13: time.Weekday(rando.Int()),
			14: time.Weekday(rando.Int()),
			15: time.Weekday(rando.Int()),
			16: time.Weekday(rando.Int()),
			17: time.Weekday(rando.Int()),
		},
		Two: [17]time.Weekday{
			0:  time.Weekday(rando.Int()),
			1:  time.Weekday(rando.Int()),
			2:  time.Weekday(rando.Int()),
			3:  time.Weekday(rando.Int()),
			4:  time.Weekday(rando.Int()),
			5:  time.Weekday(rando.Int()),
			6:  time.Weekday(rando.Int()),
			7:  time.Weekday(rando.Int()),
			8:  time.Weekday(rando.Int()),
			9:  time.Weekday(rando.Int()),
			10: time.Weekday(rando.Int()),
			11: time.Weekday(rando.Int()),
			12: time.Weekday(rando.Int()),
			13: time.Weekday(rando.Int()),
			14: time.Weekday(rando.Int()),
			15: time.Weekday(rando.Int()),
			16: time.Weekday(rando.Int()),
		},
		Three: [16]time.Weekday{
			0:  time.Weekday(rando.Int()),
			1:  time.Weekday(rando.Int()),
			2:  time.Weekday(rando.Int()),
			3:  time.Weekday(rando.Int()),
			4:  time.Weekday(rando.Int()),
			5:  time.Weekday(rando.Int()),
			6:  time.Weekday(rando.Int()),
			7:  time.Weekday(rando.Int()),
			8:  time.Weekday(rando.Int()),
			9:  time.Weekday(rando.Int()),
			10: time.Weekday(rando.Int()),
			11: time.Weekday(rando.Int()),
			12: time.Weekday(rando.Int()),
			13: time.Weekday(rando.Int()),
			14: time.Weekday(rando.Int()),
			15: time.Weekday(rando.Int()),
		},
	}
	src := expected.MarshalJ()
	require.NoError(t, actual.UnmarshalJ(src))
	// require.NotEqual(t, Three{}, expected)
	// require.NotEqual(t, Three{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_4(t *testing.T) {
	var expected, actual Four
	require.NoError(t, actual.UnmarshalJ(expected.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, Four{}, expected)
	require.Equal(t, Four{}, actual)

	expected = Four{
		One: [15]time.Weekday{
			0:  time.Weekday(rando.Int()),
			1:  time.Weekday(rando.Int()),
			2:  time.Weekday(rando.Int()),
			3:  time.Weekday(rando.Int()),
			4:  time.Weekday(rando.Int()),
			5:  time.Weekday(rando.Int()),
			6:  time.Weekday(rando.Int()),
			7:  time.Weekday(rando.Int()),
			8:  time.Weekday(rando.Int()),
			9:  time.Weekday(rando.Int()),
			10: time.Weekday(rando.Int()),
			11: time.Weekday(rando.Int()),
			12: time.Weekday(rando.Int()),
			13: time.Weekday(rando.Int()),
			14: time.Weekday(rando.Int()),
		},
		Two: [14]time.Weekday{
			0:  time.Weekday(rando.Int()),
			1:  time.Weekday(rando.Int()),
			2:  time.Weekday(rando.Int()),
			3:  time.Weekday(rando.Int()),
			4:  time.Weekday(rando.Int()),
			5:  time.Weekday(rando.Int()),
			6:  time.Weekday(rando.Int()),
			7:  time.Weekday(rando.Int()),
			8:  time.Weekday(rando.Int()),
			9:  time.Weekday(rando.Int()),
			10: time.Weekday(rando.Int()),
			11: time.Weekday(rando.Int()),
			12: time.Weekday(rando.Int()),
			13: time.Weekday(rando.Int()),
		},
		Three: [13]time.Weekday{
			0:  time.Weekday(rando.Int()),
			1:  time.Weekday(rando.Int()),
			2:  time.Weekday(rando.Int()),
			3:  time.Weekday(rando.Int()),
			4:  time.Weekday(rando.Int()),
			5:  time.Weekday(rando.Int()),
			6:  time.Weekday(rando.Int()),
			7:  time.Weekday(rando.Int()),
			8:  time.Weekday(rando.Int()),
			9:  time.Weekday(rando.Int()),
			10: time.Weekday(rando.Int()),
			11: time.Weekday(rando.Int()),
			12: time.Weekday(rando.Int()),
		},
		Four: [12]time.Weekday{
			0:  time.Weekday(rando.Int()),
			1:  time.Weekday(rando.Int()),
			2:  time.Weekday(rando.Int()),
			3:  time.Weekday(rando.Int()),
			4:  time.Weekday(rando.Int()),
			5:  time.Weekday(rando.Int()),
			6:  time.Weekday(rando.Int()),
			7:  time.Weekday(rando.Int()),
			8:  time.Weekday(rando.Int()),
			9:  time.Weekday(rando.Int()),
			10: time.Weekday(rando.Int()),
			11: time.Weekday(rando.Int()),
		},
	}
	src := expected.MarshalJ()
	require.NoError(t, actual.UnmarshalJ(src))
	// require.NotEqual(t, Four{}, expected)
	// require.NotEqual(t, Four{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_5(t *testing.T) {
	var expected, actual Five
	require.NoError(t, actual.UnmarshalJ(expected.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, Five{}, expected)
	require.Equal(t, Five{}, actual)

	expected = Five{
		One: [11]time.Weekday{
			0:  time.Weekday(rando.Int()),
			1:  time.Weekday(rando.Int()),
			2:  time.Weekday(rando.Int()),
			3:  time.Weekday(rando.Int()),
			4:  time.Weekday(rando.Int()),
			5:  time.Weekday(rando.Int()),
			6:  time.Weekday(rando.Int()),
			7:  time.Weekday(rando.Int()),
			8:  time.Weekday(rando.Int()),
			9:  time.Weekday(rando.Int()),
			10: time.Weekday(rando.Int()),
		},
		Two: [10]time.Weekday{
			0: time.Weekday(rando.Int()),
			1: time.Weekday(rando.Int()),
			2: time.Weekday(rando.Int()),
			3: time.Weekday(rando.Int()),
			4: time.Weekday(rando.Int()),
			5: time.Weekday(rando.Int()),
			6: time.Weekday(rando.Int()),
			7: time.Weekday(rando.Int()),
			8: time.Weekday(rando.Int()),
			9: time.Weekday(rando.Int()),
		},
		Three: [9]time.Weekday{
			0: time.Weekday(rando.Int()),
			1: time.Weekday(rando.Int()),
			2: time.Weekday(rando.Int()),
			3: time.Weekday(rando.Int()),
			4: time.Weekday(rando.Int()),
			5: time.Weekday(rando.Int()),
			6: time.Weekday(rando.Int()),
			7: time.Weekday(rando.Int()),
			8: time.Weekday(rando.Int()),
		},
		Four: [8]time.Weekday{
			0: time.Weekday(rando.Int()),
			1: time.Weekday(rando.Int()),
			2: time.Weekday(rando.Int()),
			3: time.Weekday(rando.Int()),
			4: time.Weekday(rando.Int()),
			5: time.Weekday(rando.Int()),
			6: time.Weekday(rando.Int()),
			7: time.Weekday(rando.Int()),
		},
		Five: [7]time.Weekday{
			0: time.Weekday(rando.Int()),
			1: time.Weekday(rando.Int()),
			2: time.Weekday(rando.Int()),
			3: time.Weekday(rando.Int()),
			4: time.Weekday(rando.Int()),
			5: time.Weekday(rando.Int()),
			6: time.Weekday(rando.Int()),
		},
	}
	src := expected.MarshalJ()
	require.NoError(t, actual.UnmarshalJ(src))
	// require.NotEqual(t, Five{}, expected)
	// require.NotEqual(t, Five{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_6(t *testing.T) {
	var expected, actual Six
	require.NoError(t, actual.UnmarshalJ(expected.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, Six{}, expected)
	require.Equal(t, Six{}, actual)

	expected = Six{
		One: [6]time.Weekday{
			0: time.Weekday(rando.Int()),
			1: time.Weekday(rando.Int()),
			2: time.Weekday(rando.Int()),
			3: time.Weekday(rando.Int()),
			4: time.Weekday(rando.Int()),
			5: time.Weekday(rando.Int()),
		},
		Two: [5]time.Weekday{
			0: time.Weekday(rando.Int()),
			1: time.Weekday(rando.Int()),
			2: time.Weekday(rando.Int()),
			3: time.Weekday(rando.Int()),
			4: time.Weekday(rando.Int()),
		},
		Three: [4]time.Weekday{
			0: time.Weekday(rando.Int()),
			1: time.Weekday(rando.Int()),
			2: time.Weekday(rando.Int()),
			3: time.Weekday(rando.Int()),
		},
		Four: [3]time.Weekday{
			0: time.Weekday(rando.Int()),
			1: time.Weekday(rando.Int()),
			2: time.Weekday(rando.Int()),
		},
		Five: [2]time.Weekday{
			0: time.Weekday(rando.Int()),
			1: time.Weekday(rando.Int()),
		},
		Six: [1]time.Weekday{time.Weekday(rando.Int())},
	}
	src := expected.MarshalJ()
	require.NoError(t, actual.UnmarshalJ(src))
	// require.NotEqual(t, Six{}, expected)
	// require.NotEqual(t, Six{}, actual)
	require.Equal(t, expected, actual)
}
