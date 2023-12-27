package rando

import (
	"math"
	"math/rand"
)

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

// Strings returns a random size string slice with random length strings.
func Strings() (s []string) {
	qty := rand.Intn(math.MaxUint8)
	for i := 0; i < qty; i++ {
		s = append(s, String())
	}
	return
}

// StringsQty returns a string slice with size `qty` with random length strings.
func StringsQty(qty int) (s []string) {
	for i := 0; i < qty; i++ {
		s = append(s, String())
	}
	return
}

// StringsQtyN returns a random size string slice up to size `qty` with random length strings.
func StringsQtyN(qty, n int) (s []string) {
	for i := 0; i < qty; i++ {
		s = append(s, StringN(n))
	}
	return
}

// StringsQN returns a random size string slice up to size `qty` with random length strings.
func StringsQN(q, n int) (s []string) {
	q = rand.Intn(q)
	for i := 0; i < q; i++ {
		s = append(s, StringN(n))
	}
	return
}

// StringsN returns a random size string slice with random strings of length `l`.
func StringsN(l int) (s []string) {
	qty := rand.Intn(math.MaxUint8)
	for i := 0; i < qty; i++ {
		s = append(s, StringN(l))
	}
	return
}
