// go test -bench=.*MVR.* . -benchmem  -benchtime=30s -shuffle=on

package byte

import (
	"math"
	"math/rand"
	"testing"
)

var mk = MakeVsReturn{Bool: true, Uint8: uint8(rand.Intn(255)), Int8: int8(rand.Intn(math.MaxInt8))}

func BenchmarkMVR_Return1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		y = mk.Return1()
	}
}
func BenchmarkMVR_Make1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		y = mk.Make1()
	}
}

func BenchmarkMVR_Return2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		y = mk.Return2()
	}
}
func BenchmarkMVR_Make2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		y = mk.Make2()
	}
}

func BenchmarkMVR_Return3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		y = mk.Return3()
	}
}
func BenchmarkMVR_Make3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		y = mk.Make3()
	}
}

var mk2 = MakeVsReturnLarge{
	Bool:       randBool(),
	Hidden:     randBool(),
	Selected:   randBool(),
	Deselected: randBool(),
	Activated:  randBool(),
	Modified:   randBool(),
	Serviced:   randBool(),
	Performed:  randBool(),
	Accessed:   randBool(),
	Updated:    randBool(),
	Uint80:     randUint8(),
	Uint81:     randUint8(),
	Uint82:     randUint8(),
	Uint83:     randUint8(),
	Uint84:     randUint8(),
	Uint85:     randUint8(),
	Uint86:     randUint8(),
	Uint87:     randUint8(),
	Uint88:     randUint8(),
	Uint89:     randUint8(),
	Uint90:     randUint8(),
	Int80:      randInt8(),
	Int81:      randInt8(),
	Int82:      randInt8(),
	Int83:      randInt8(),
	Int84:      randInt8(),
	Int85:      randInt8(),
	Int86:      randInt8(),
	Int87:      randInt8(),
	Int88:      randInt8(),
	Int89:      randInt8(),
	Int90:      randInt8(),
}

func BenchmarkMVRLarge_Make1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		y = mk2.Make1()
	}
}
func BenchmarkMVRLarge_Return1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		y = mk2.Return1()
	}
}

func randByte() byte {
	return randUint8()
}

func randUint8() uint8 {
	return uint8(rand.Intn(math.MaxUint8))
}

func randInt8() int8 {
	return int8(rand.Intn(math.MaxInt8))
}

func randBool() bool {
	return rand.Intn(1) == 1
}
