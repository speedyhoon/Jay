package generate

import (
	"bytes"
	"fmt"
)

func (s *structTyp) writeSingles(b *bytes.Buffer, byteIndex *uint, importJ *bool) {
	if len(s.single) == 0 {
		return
	}

	if s.returnInline && len(s.bool) == 0 {
		bufWriteF(b, "%s{", tBytes)
	}

	for i, l := 0, len(s.single); i < l; i++ {
		isLast := i+1 == l
		fun, _ := s.single[i].MarshalFuncTemplate(importJ)
		writeSingle(s.single[i], b, *byteIndex, fun, !s.returnInline, isLast)
		*byteIndex++
	}

	if s.returnInline {
		b.WriteByte('}')
	}
}

func writeSingle(single *field, b *bytes.Buffer, byteIndex uint, fun string, isMake, isLast bool) {
	if isMake {
		bufWriteLineF(b, "%s[%d] = %s", single.structTyp.bufferName, byteIndex, printFunc(fun, single.Name()))
		return
	}

	b.WriteString(printFunc(fun, single.Name()))
	if !isLast {
		b.WriteString(",")
	}
}

func (s *structTyp) readSingles(b *bytes.Buffer, byteIndex *uint) {
	for i, l := 0, len(s.single); i < l; i++ {
		fun, _, _ := s.single[i].unmarshalFunc()
		readSingle(s.single[i], b, *byteIndex, fun)
		*byteIndex += s.single[i].typeFuncSize()
	}
}

func readSingle(single *field, b *bytes.Buffer, byteIndex uint, fun string) {
	bufWriteLineF(b, "%s=%s", single.Name(), printFunc(fun, fmt.Sprintf("%s[%d]", single.structTyp.bufferName, byteIndex)))
}
