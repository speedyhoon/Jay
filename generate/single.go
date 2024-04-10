package generate

import (
	"bytes"
	"fmt"
)

func (s *structTyp) writeSingles(b *bytes.Buffer, byteIndex *uint, receiver string, importJ *bool) (isReturnInlined bool) {
	if len(s.single) == 0 {
		return false
	}

	isMake := s.useMakeFunc()

	for i, l := 0, len(s.single); i < l; i++ {
		isLast := i+1 == l
		fun, _ := s.single[i].MarshalFuncTemplate(importJ)
		writeSingle(s.single[i], b, *byteIndex, receiver, fun, s.bufferName, isMake, isLast)
		*byteIndex++
	}

	return !isMake
}

func (s *structTyp) useMakeFunc() bool {
	return len(s.variableLen) >= 1 || len(s.fixedLen) >= 1
}

func writeSingle(single field, b *bytes.Buffer, byteIndex uint, receiver, fun, bufferName string, isMake, isLast bool) {
	thisField := pkgSelName(receiver, single.name)

	if isMake {
		bufWriteF(b, "%s[%d]=%s\n", bufferName, byteIndex, printFunc(fun, thisField))
		return
	}

	b.WriteString(printFunc(fun, thisField))
	if !isLast {
		b.WriteString(",")
	}
}

func (s *structTyp) readSingles(b *bytes.Buffer, byteIndex *uint) {
	for i, l := 0, len(s.single); i < l; i++ {
		fun, _ := s.single[i].unmarshalFuncs()
		readSingle(s.single[i], b, *byteIndex, fun)
		*byteIndex += s.single[i].typeFuncSize()
	}
}

func readSingle(single field, b *bytes.Buffer, byteIndex uint, fun string) {
	bufWriteF(b, "%s.%s=%s\n", single.structTyp.receiver, single.name, printFunc(fun, fmt.Sprintf("%s[%d]", single.structTyp.bufferName, byteIndex)))
}
