package main

import (
	"fmt"
	"github.com/speedyhoon/jay"
	"github.com/stretchr/testify/assert"
	"math"
	"math/rand"
	"testing"
)

func TestCar_MarshalUnmarshal_empty(t *testing.T) {
	empty := Car{}
	src := empty.MarshalJ()
	assert.Equal(t, make([]byte, 33), src)
	assert.Equal(t, Car{}, empty)

	err := empty.UnmarshalJ(src)
	assert.NoError(t, err)
	assert.Equal(t, Car{}, empty)
}

func TestCar_MarshalUnmarshal_small(t *testing.T) {
	c := Car{
		ID:   rand.Uint64(),
		Row:  uint(rand.Uint64()),
		Name: fmt.Sprint(rand.Uint64()),
		Auto: true,
		CC:   fmt.Sprint(rand.Uint64()),
		//Timing:  ptrStr(fmt.Sprint(rand.Uint64())),
		RedLine: uint16(rand.Intn(jay.MaxUint16)),
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
		Name: string(make([]byte, jay.MaxUint8)),
		Auto: true,
		CC:   string(make([]byte, jay.MaxUint8)),
		//Timing:  ptrStr(string(make([]byte,jay.MaxUint24))),
		RedLine: jay.MaxUint16,
		Gearbox: gearbox{},
	}
	src := c.MarshalJ()
	assert.NotEqual(t, make([]byte, 33+2*jay.MaxUint8), src)
	assert.NotEqual(t, Car{}, c)

	d := Car{}
	err := d.UnmarshalJ(src)
	assert.NoError(t, err)
	assert.Equal(t, c, d)
}

func ptrStr(s string) *string {
	return &s
}
