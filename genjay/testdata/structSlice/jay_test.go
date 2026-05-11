package main

import (
	"github.com/go-openapi/testify/v2/require"
	"github.com/speedyhoon/rando"
	"testing"
)

func TestFuzz_1(t *testing.T) {
	var expected, actual One
	require.NoError(t, actual.UnmarshalJ(expected.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, One{}, expected)
	require.Equal(t, One{}, actual)

	expected = One{
		One: rando.Bool(),
	}
	src := expected.MarshalJ()
	require.NoError(t, actual.UnmarshalJ(src))
	require.Equal(t, expected, actual)
}

func TestFuzz_2(t *testing.T) {
	var expected, actual Two
	require.NoError(t, actual.UnmarshalJ(expected.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, Two{}, expected)
	require.Equal(t, Two{}, actual)

	expected = Two{
		Ones: nil,
		Two:  rando.BoolsN(7),
	}
	src := expected.MarshalJ()
	require.NoError(t, actual.UnmarshalJ(src))
	require.Equal(t, expected, actual)
}
