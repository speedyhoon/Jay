// go test -bench=.*MVR.* . -benchmem  -benchtime=30s -shuffle=on

package byte

import (
	"github.com/speedyhoon/jay/rando"
	"testing"
)

var mk = MakeVsReturn{Bool: true, Uint8: rando.Uint8(), Int8: rando.Int8()}

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
	Bool:       rando.Bool(),
	Hidden:     rando.Bool(),
	Selected:   rando.Bool(),
	Deselected: rando.Bool(),
	Activated:  rando.Bool(),
	Modified:   rando.Bool(),
	Serviced:   rando.Bool(),
	Performed:  rando.Bool(),
	Accessed:   rando.Bool(),
	Updated:    rando.Bool(),
	Uint80:     rando.Uint8(),
	Uint81:     rando.Uint8(),
	Uint82:     rando.Uint8(),
	Uint83:     rando.Uint8(),
	Uint84:     rando.Uint8(),
	Uint85:     rando.Uint8(),
	Uint86:     rando.Uint8(),
	Uint87:     rando.Uint8(),
	Uint88:     rando.Uint8(),
	Uint89:     rando.Uint8(),
	Uint90:     rando.Uint8(),
	Int80:      rando.Int8(),
	Int81:      rando.Int8(),
	Int82:      rando.Int8(),
	Int83:      rando.Int8(),
	Int84:      rando.Int8(),
	Int85:      rando.Int8(),
	Int86:      rando.Int8(),
	Int87:      rando.Int8(),
	Int88:      rando.Int8(),
	Int89:      rando.Int8(),
	Int90:      rando.Int8(),
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
