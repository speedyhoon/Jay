package main

import (
	"github.com/speedyhoon/rando"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestCar_MarshalUnmarshal_empty(t *testing.T) {
	empty := Car{}
	src := empty.MarshalJ()
	assert.Equal(t, Car{}, empty)

	err := empty.UnmarshalJ(src)
	assert.NoError(t, err)
	assert.Equal(t, Car{}, empty)
}

func TestCar_MarshalUnmarshal_small(t *testing.T) {
	c := Car{
		ID:   rando.Uint64(),
		Row:  rando.Uint(),
		Name: rando.String(),
		Auto: rando.Bool(),
		CC:   rando.String(),
		//Timing:  ptrStr(rando.String()),
		RedLine: rando.Uint16(),
		Expiry:  rando.DateTime(),
		Gearbox: gearbox{},
	}
	src := c.MarshalJ()
	assert.NotEqual(t, make([]byte, 33), src)
	assert.NotEqual(t, Car{}, c)

	d := Car{}
	err := d.UnmarshalJ(src)
	assert.NoError(t, err)
	assert.Equal(t, c, d)
}

func TestCar_MarshalUnmarshal_large(t *testing.T) {
	c := Car{
		ID:   math.MaxUint64,
		Row:  uint(math.MaxUint64),
		Name: string(make([]byte, math.MaxUint8)),
		Auto: true,
		CC:   string(make([]byte, math.MaxUint8)),
		//Timing:  ptrStr(string(make([]byte,jay.MaxUint24))),
		RedLine: math.MaxUint16,
		Gearbox: gearbox{},
	}
	src := c.MarshalJ()
	assert.NotEqual(t, make([]byte, 33+2*math.MaxUint8), src)
	assert.NotEqual(t, Car{}, c)

	d := Car{}
	err := d.UnmarshalJ(src)
	assert.NoError(t, err)
	assert.Equal(t, c, d)
}

//func ptrStr(s string) *string {
//	return &s
//}
