package main

import (
	"testing"

	"github.com/go-openapi/testify/v2/require"
	"github.com/speedyhoon/jay/genjay/testdata/imports/ext"
	"github.com/speedyhoon/rando"
)

func TestSection(t *testing.T) {
	var expected, actual Section
	require.NoError(t, actual.UnmarshalJ(expected.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, Section{}, expected)
	require.Equal(t, Section{}, actual)

	expected = Section{
		Id:      nil,
		Name:    rando.String(),
		Color:   ext.C16(rando.Uint16()),
		Sectors: ext.Sectors(rando.Uint64sN(5)),
		Sections: [5]ext.Sections{
			ext.Sections(rando.Uint64()),
			ext.Sections(rando.Uint64()),
			ext.Sections(rando.Uint64()),
			ext.Sections(rando.Uint64()),
			ext.Sections(rando.Uint64()),
		},
		Project: nil,
		Order:   rando.Bytes(),
	}
	src := expected.MarshalJ()
	require.NoError(t, actual.UnmarshalJ(src))
	require.Equal(t, expected, actual)
}

func TestProject(t *testing.T) {
	var expected, actual Project
	require.NoError(t, actual.UnmarshalJ(expected.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, Project{}, expected)
	require.Equal(t, Project{}, actual)

	expected = Project{
		Id:   nil,
		Name: rando.String(),
	}
	src := expected.MarshalJ()
	require.NoError(t, actual.UnmarshalJ(src))
	require.Equal(t, expected, actual)
}
