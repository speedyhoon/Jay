package generate

import (
	"bytes"
	"fmt"
)

func (s *structTyp) writeSingles(b *bytes.Buffer, byteIndex *uint, receiver string, opt Option, importJ *bool) (isReturnInlined bool) {
	if len(s.single) == 0 {
		return false
	}

	isMake := s.useMakeFunc()

	for i, l := 0, len(s.single); i < l; i++ {
		isLast := i+1 == l
		fun, _, _ := opt.typeFuncs(s.single[i], isLast, importJ)
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

func (s *structTyp) readSingles(b *bytes.Buffer, byteIndex *uint, receiver string, opt Option) {
	for i, l := 0, len(s.single); i < l; i++ {
		isLast := i+1 == l
		fun, _, _ := opt.unmarshalFuncs(s.single[i], isLast)
		readSingle(s.single[i], b, *byteIndex, receiver, fun, s.bufferName)
		*byteIndex++
	}
}

func readSingle(single field, b *bytes.Buffer, byteIndex uint, receiver, fun, bufferName string) {
	bufWriteF(b, "%s.%s=%s\n", receiver, single.name, printFunc(fun, fmt.Sprintf("%s[%d]", bufferName, byteIndex)))
}
