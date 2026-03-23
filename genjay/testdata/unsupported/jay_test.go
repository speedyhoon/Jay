package main

import (
	"os"
	"reflect"
	"testing"

	"github.com/go-openapi/testify/v2/require"
	"github.com/speedyhoon/jay/genjay"
)

func TestUnsupported(t *testing.T) {
	src, err := os.ReadFile(genjay.DefaultOutputFileName)
	require.Error(t, err, "Expected file %s to not exist", genjay.DefaultOutputFileName)
	require.Nil(t, src, "Expected src to be empty")

	rt := reflect.TypeOf(Unsupported{})
	m, ok := rt.MethodByName(genjay.MethodMarshalJ)
	require.False(t, ok, "expected MarshalJ to not exist")
	require.Empty(t, m, "expected MarshalJ to not exist")

	m, ok = rt.MethodByName(genjay.MethodUnmarshalJ)
	require.False(t, ok, "expected UnmarshalJ to not exist")
	require.Empty(t, m, "expected UnmarshalJ to not exist")
}
