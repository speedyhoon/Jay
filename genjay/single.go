package genjay

import (
	"bytes"
	"fmt"
	"strings"
)

func (s *structTyp) writeSingles(c *varCtx) {
	if len(s.bool) == 0 && len(s.single) == 0 {
		return
	}

	var lines []string
	s.writeBools(&lines)

	for i, mx := 0, len(s.single); i < mx; i++ {
		fun, _ := s.single[i].marshalFuncTemplate()
		lines = append(lines, writeSingle(s.single[i], fun, !s.returnInline))
	}

	if s.returnInline {
		c.addLineF("%s{%s}", tBytes, strings.Join(lines, ", "))
	} else {
		c.addLines(lines...)
	}
}

func writeSingle(single *field, fun string, isMake bool) string {
	if isMake {
		return fmt.Sprintf("%s[%d] = %s", single.structTyp.bufferName, *single.indexStart, printFunc(fun, single.Name()))
	}

	return printFunc(fun, single.Name())
}

func (s *structTyp) readSingles(b *bytes.Buffer) {
	for i, l := 0, len(s.single); i < l; i++ {
		fun, _, _ := s.single[i].unmarshalFunc()
		readSingle(s.single[i], b, fun)
	}
}

func readSingle(single *field, b *bytes.Buffer, fun string) {
	bufWriteLineF(b, "%s = %s", single.Name(), printFunc(fun, fmt.Sprintf("%s[%d]", single.structTyp.bufferName, *single.indexStart)))
}
