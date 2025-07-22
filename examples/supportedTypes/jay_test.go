package main

import (
	"github.com/speedyhoon/rando"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSupported(t *testing.T) {
	var expected, actual Supported
	require.NoError(t, actual.UnmarshalJ(expected.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, Supported{}, expected)
	require.Equal(t, Supported{}, actual)

	expected = Supported{
		Bool:    rando.Bool(),
		Byte:    rando.Uint8(),
		Float32: rando.Float32(),
		Float64: rando.Float64(),
		Guid:    [16]byte(rando.BytesN(16)),
		Int:     rando.Int(),
		Int8:    rando.Int8(),
		Int16:   rando.Int16(),
		Int32:   rando.Int32(),
		Int64:   rando.Int64(),
		Uint:    rando.Uint(),
		Uint8:   rando.Uint8(),
		Uint16:  rando.Uint16(),
		Uint32:  rando.Uint32(),
		Uint64:  rando.Uint64(),
		Rune:    rando.Rune(),
		String:  rando.String(),
		Time:    rando.Time(),
		Nano:    rando.TimeNano(),
		Embed: Embed{
			Bool:      rando.Bool(),
			Byte:      rando.Byte(),
			Float32:   rando.Float32(),
			Float64:   rando.Float64(),
			Int:       rando.Int(),
			Int8:      rando.Int8(),
			Int16:     rando.Int16(),
			Int32:     rando.Int32(),
			Int64:     rando.Int64(),
			Uint:      rando.Uint(),
			Uint8:     rando.Uint8(),
			Uint16:    rando.Uint16(),
			Uint32:    rando.Uint32(),
			Uint64:    rando.Uint64(),
			Rune:      rando.Rune(),
			String:    rando.String(),
			Time:      rando.Time(),
			Nano:      rando.TimeNano(),
			ByteSlice: rando.Bytes(),
		},
		SubStruct: subStruct{
			Bool:      rando.Bool(),
			Byte:      rando.Byte(),
			Float32:   rando.Float32(),
			Float64:   rando.Float64(),
			Int:       rando.Int(),
			Int8:      rando.Int8(),
			Int16:     rando.Int16(),
			Int32:     rando.Int32(),
			Int64:     rando.Int64(),
			Uint:      rando.Uint(),
			Uint8:     rando.Uint8(),
			Uint16:    rando.Uint16(),
			Uint32:    rando.Uint32(),
			Uint64:    rando.Uint64(),
			Rune:      rando.Rune(),
			String:    rando.String(),
			Time:      rando.Time(),
			Nano:      rando.TimeNano(),
			ByteSlice: rando.Bytes(),
		},
		ByteSlice: rando.Bytes(),
	}
	src := expected.MarshalJ()
	require.NoError(t, actual.UnmarshalJ(src))
	// require.NotEqual(t, Supported{}, expected)
	// require.NotEqual(t, Supported{}, actual)
	require.Equal(t, expected, actual)
}

func TestEmbed(t *testing.T) {
	var expected, actual Embed
	require.NoError(t, actual.UnmarshalJ(expected.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, Embed{}, expected)
	require.Equal(t, Embed{}, actual)

	expected = Embed{
		Bool:      rando.Bool(),
		Byte:      rando.Uint8(),
		Float32:   rando.Float32(),
		Float64:   rando.Float64(),
		Int:       rando.Int(),
		Int8:      rando.Int8(),
		Int16:     rando.Int16(),
		Int32:     rando.Int32(),
		Int64:     rando.Int64(),
		Uint:      rando.Uint(),
		Uint8:     rando.Uint8(),
		Uint16:    rando.Uint16(),
		Uint32:    rando.Uint32(),
		Uint64:    rando.Uint64(),
		Rune:      rando.Rune(),
		String:    rando.String(),
		Time:      rando.Time(),
		Nano:      rando.TimeNano(),
		ByteSlice: rando.Bytes(),
	}
	src := expected.MarshalJ()
	require.NoError(t, actual.UnmarshalJ(src))
	// require.NotEqual(t, Embed{}, expected)
	// require.NotEqual(t, Embed{}, actual)
	require.Equal(t, expected, actual)
}
