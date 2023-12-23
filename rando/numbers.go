package rando

import (
	"math"
	"math/rand"
	"time"
)

func Bool() bool {
	return rand.Intn(1) == 1
}

func Int8() int8 {
	return int8(Uint8() - math.MaxInt8 + 1)
}

func Int16() int16 {
	return int16(Uint16() - math.MaxInt16 + 1)
}

func Int32() int32 {
	return int32(Uint32() - math.MaxInt32 + 1)
}

func Rune() rune {
	return Int32()
}
func Byte() byte {
	return Uint8()
}

func Int64() int64 {
	return int64(Uint64() - math.MaxInt64 + 1)
}

func Uint8() uint8 {
	return uint8(rand.Intn(math.MaxUint8))
}

func Uint16() uint16 {
	return uint16(rand.Intn(math.MaxUint16))
}

func Uint32() uint32 {
	return rand.Uint32()
}

func Uint64() uint64 {
	return rand.Uint64()
}

const intSize = 32 << (^uint(0) >> 63) // 32 or 64
func Uint() uint {
	if intSize == 32 {
		return uint(Uint32())
	}
	return uint(Uint64())
}

func Int() int {
	if intSize == 32 {
		return int(Int32())
	}
	return int(Int64())
}

func Float32() float32 {
	return rand.Float32() * float32(math.Pow(10, float64(rand.Intn(9)+1)))
}

func Float64() float64 {
	return rand.Float64() * math.Pow(10, float64(rand.Intn(9)+1))
}

func StringN(length int) (s string) {
	if length < 0 {
		length = rand.Intn(100)
	}
	for i := 0; i < length; i++ {
		s += string(Uint8())
	}
	return
}

func String() (s string) {
	return StringN(-1)
}

func DateTime() time.Time {
	return time.Unix(Int64(), 0).UTC()
}

func TimeNano() time.Time {
	return time.Unix(0, Int64()).UTC()
}
