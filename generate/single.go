package generate

import (
	"bytes"
	"fmt"
)

func (s *structTyp) writeSingles(b *bytes.Buffer, byteIndex *uint, receiver string, opt Option) (isReturnInlined bool) {
	if len(s.single) == 0 {
		return false
	}

	isMake := s.useMakeFunc()

	for i, l := 0, len(s.single); i < l; i++ {
		isLast := i+1 == l
		fun, _ := opt.typeFuncs(s.single[i], isLast)
		writeSingle(s.single[i], b, *byteIndex, receiver, fun, isMake, isLast)
		*byteIndex++
	}

	return !isMake
}

func (s *structTyp) useMakeFunc() bool {
	return len(s.variableLen) >= 1 || len(s.fixedLen) >= 1
}

func writeSingle(single field, b *bytes.Buffer, byteIndex uint, receiver, fun string, isMake, isLast bool) {
	thisField := fmt.Sprintf("%s.%s", receiver, single.name)

	if isMake {
		bufWriteF(b, "b[%d]=%s\n", byteIndex, printFunc(fun, thisField))
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
		fun, _ := opt.unmarshalFuncs(s.single[i], isLast)
		readSingle(s.single[i], b, *byteIndex, receiver, fun)
		*byteIndex++
	}
}

func readSingle(single field, b *bytes.Buffer, byteIndex uint, receiver, fun string) {
	bufWriteF(b, "%s.%s=%s\n", receiver, single.name, printFunc(fun, fmt.Sprintf("b[%d]", byteIndex)))
}
