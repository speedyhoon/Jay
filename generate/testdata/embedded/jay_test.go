package embedded

import (
	"testing"

	"github.com/go-openapi/testify/v2/require"
	"github.com/speedyhoon/rando"
)

func TestFuzz_Configuration(t *testing.T) {
	var expected, actual Configuration
	require.NoError(t, actual.UnmarshalJ(expected.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, Configuration{}, expected)
	require.Equal(t, Configuration{}, actual)

	expected = Configuration{
		Alarm: Alarm{
			SirenOn:     rando.Duration(),
			SirenOff:    rando.Duration(),
			Repeats:     rando.Uint8(),
			CheckEvery:  rando.Duration(),
			Types:       rando.Bytes(),
			Attempts:    rando.Uint8(),
			KeyAttempts: rando.Uint8(),
		},
		DataLog: DataLog{
			Version:     rando.Uint(),
			VersionRace: rando.Uint(),
			FileSizeMax: rando.Uint32(),
			Cluster:     rando.Bytes(),
			Off:         rando.Bool(),
		},
		Types: Types{
			UniMog:  UniMog(rando.Uint8()),
			Cluster: rando.Bytes(),
		},
	}
	src := expected.MarshalJ()
	require.NoError(t, actual.UnmarshalJ(src))
	// require.NotEqual(t, Configuration{}, expected)
	// require.NotEqual(t, Configuration{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_Types(t *testing.T) {
	var expected, actual Types
	require.NoError(t, actual.UnmarshalJ(expected.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, Types{}, expected)
	require.Equal(t, Types{}, actual)

	expected = Types{
		UniMog:  UniMog(rando.Uint8()),
		Cluster: rando.Bytes(),
	}
	src := expected.MarshalJ()
	require.NoError(t, actual.UnmarshalJ(src))
	// require.NotEqual(t, Types{}, expected)
	// require.NotEqual(t, Types{}, actual)
	require.Equal(t, expected, actual)
}
