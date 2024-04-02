package bool

import (
	"github.com/speedyhoon/jay"
	"testing"
)

var eight Eight
var err error
var twenty3 TwentyThree

func BenchmarkUnmarshalJ(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err = eight.UnmarshalJ([]byte{215})
	}
}

func BenchmarkUnmarshalJ2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err = eight.UnmarshalJ2([]byte{215})
	}
}

func BenchmarkUnmarshalJ3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err = eight.UnmarshalJ3([]byte{215})
	}
}

type Eight struct {
	One   bool
	Two   bool
	Three bool
	Four  bool
	Five  bool
	Six   bool
	Seven bool
	Eight bool
}

func (e *Eight) UnmarshalJ(b []byte) error {
	if len(b) < 1 {
		return jay.ErrUnexpectedEOB
	}
	e.One, e.Two, e.Three, e.Four, e.Five, e.Six, e.Seven, e.Eight = jay.ReadBool8(b[0])
	return nil
}

func (e *Eight) UnmarshalJ2(b []byte) error {
	if len(b) < 1 {
		return jay.ErrUnexpectedEOB
	}
	e.One, e.Two, e.Three, e.Four, e.Five, e.Six, e.Seven, e.Eight = b[0] >= 128, b[0]&64 == 64, b[0]&32 == 32, b[0]&16 == 16, b[0]&8 == 8, b[0]&4 == 4, b[0]&2 == 2, b[0]&1 == 1
	return nil
}

func (e *Eight) UnmarshalJ3(b []byte) error {
	if len(b) < 1 {
		return jay.ErrUnexpectedEOB
	}
	c := b[0]
	e.One, e.Two, e.Three, e.Four, e.Five, e.Six, e.Seven, e.Eight = c >= 128, c&64 == 64, c&32 == 32, c&16 == 16, c&8 == 8, c&4 == 4, c&2 == 2, c&1 == 1
	return nil
}

type TwentyThree struct {
	One         bool
	Two         bool
	Three       bool
	Four        bool
	Five        bool
	Six         bool
	Seven       bool
	Eight       bool
	Nine        bool
	Ten         bool
	Eleven      bool
	Twelve      bool
	Thirteen    bool
	Fourteen    bool
	Fifteen     bool
	Sixteen     bool
	Seventeen   bool
	Eighteen    bool
	Nineteen    bool
	Twenty      bool
	TwentyOne   bool
	TwentyTwo   bool
	TwentyThree bool
}

func BenchmarkUnmarshalJ23(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err = twenty3.UnmarshalJ([]byte{215, 49, 51})
	}
}

func BenchmarkUnmarshalJ23_2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err = twenty3.UnmarshalJ2([]byte{215, 49, 51})
	}
}

func BenchmarkUnmarshalJ23_3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err = twenty3.UnmarshalJ3([]byte{215, 49, 51})
	}
}

func (t *TwentyThree) UnmarshalJ(b []byte) error {
	if len(b) < 3 {
		return jay.ErrUnexpectedEOB
	}
	t.One, t.Two, t.Three, t.Four, t.Five, t.Six, t.Seven, t.Eight = jay.ReadBool8(b[0])
	t.Nine, t.Ten, t.Eleven, t.Twelve, t.Thirteen, t.Fourteen, t.Fifteen, t.Sixteen = jay.ReadBool8(b[1])
	t.Seventeen, t.Eighteen, t.Nineteen, t.Twenty, t.TwentyOne, t.TwentyTwo, t.TwentyThree = jay.ReadBool7(b[2])
	return nil
}

func (t *TwentyThree) UnmarshalJ2(b []byte) error {
	if len(b) < 3 {
		return jay.ErrUnexpectedEOB
	}
	t.One, t.Two, t.Three, t.Four, t.Five, t.Six, t.Seven, t.Eight = b[0] >= 128, b[0]&64 == 64, b[0]&32 == 32, b[0]&16 == 16, b[0]&8 == 8, b[0]&4 == 4, b[0]&2 == 2, b[0]&1 == 1
	t.Nine, t.Ten, t.Eleven, t.Twelve, t.Thirteen, t.Fourteen, t.Fifteen, t.Sixteen = b[1] >= 128, b[1]&64 == 64, b[1]&32 == 32, b[1]&16 == 16, b[1]&8 == 8, b[1]&4 == 4, b[1]&2 == 2, b[1]&1 == 1
	t.Seventeen, t.Eighteen, t.Nineteen, t.Twenty, t.TwentyOne, t.TwentyTwo, t.TwentyThree = b[2] >= 128, b[2]&64 == 64, b[2]&32 == 32, b[2]&16 == 16, b[2]&8 == 8, b[2]&4 == 4, b[2]&2 == 2
	return nil
}

func (t *TwentyThree) UnmarshalJ3(b []byte) error {
	if len(b) < 3 {
		return jay.ErrUnexpectedEOB
	}
	c := b[0]
	t.One, t.Two, t.Three, t.Four, t.Five, t.Six, t.Seven, t.Eight = c >= 128, c&64 == 64, c&32 == 32, c&16 == 16, c&8 == 8, c&4 == 4, c&2 == 2, c&1 == 1
	c = b[1]
	t.Nine, t.Ten, t.Eleven, t.Twelve, t.Thirteen, t.Fourteen, t.Fifteen, t.Sixteen = c >= 128, c&64 == 64, c&32 == 32, c&16 == 16, c&8 == 8, c&4 == 4, c&2 == 2, c&1 == 1
	c = b[2]
	t.Seventeen, t.Eighteen, t.Nineteen, t.Twenty, t.TwentyOne, t.TwentyTwo, t.TwentyThree = c >= 128, c&64 == 64, c&32 == 32, c&16 == 16, c&8 == 8, c&4 == 4, c&2 == 2
	return nil
}
