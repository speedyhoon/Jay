package alternatives

import (
	//"bebop/new"
	"encoding/json"
	"testing"
)

var foo = Foo{
	Make:     "Holden",
	Model:    "Commodore",
	Badge:    "Executive",
	Price:    16000,
	Wheels:   4,
	EngineCC: 3603,
}

//var foo2 = new.Foo{
//	Make:     "Holden",
//	Model:    "Commodore",
//	Badge:    "Executive",
//	Price:    16000,
//	Wheels:   4,
//	EngineCC: 3603,
//}

//func BenchmarkMarshalNew(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		_ = foo2.MarshalBebop()
//	}
//}

var bop []byte
var jy []byte
var jsn []byte
var err error
var fooEmpty Foo

func BenchmarkMarshalBebop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bop = foo.MarshalBebop()
	}
}

func BenchmarkMarshalJay(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jy = foo.MarshalJ()
	}
}

func BenchmarkMarshalJay2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jy = foo.MarshalJ2()
	}
}

func BenchmarkMarshalJson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jsn, _ = json.Marshal(foo)
	}
}

func BenchmarkUnmarshalBebop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err = fooEmpty.UnmarshalBebop(bop)
	}
}

func BenchmarkUnmarshalJay(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err = fooEmpty.UnmarshalJ(jy)
	}
}

func BenchmarkUnmarshalJay2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err = fooEmpty.UnmarshalJ2(jy)
	}
}

func BenchmarkUnmarshalJay3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err = fooEmpty.UnmarshalJ3(jy)
	}
}

func BenchmarkUnmarshalJay4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err = fooEmpty.UnmarshalJ4(jy)
	}
}

func BenchmarkUnmarshalJson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err = json.Unmarshal(jsn, &fooEmpty)
	}
}
