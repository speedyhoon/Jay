package generate

import (
	"bytes"
	"github.com/speedyhoon/jay"
	"strconv"
	"strings"
)

func (s *structTyp) writeStrSlice(buf *bytes.Buffer, vars []string, byteIndex uint) {
	qty := len(s.stringSlice)
	if qty == 0 {
		return
	}

	funcName := nameOf(s.stringSlice[0].sizeOfPick(jay.WriteStrings8, jay.WriteStrings16), nil)

	if qty == 1 {
		bufWriteF(buf, "%s(%s, %s)\n", funcName, s.stringSlice[0].sliceExpr(strconv.Itoa(qty), ""), s.stringSlice[0].Name())
		return
	}

	vars = []string{"l0", "l1", "l2"}
	for i, f := range s.stringSlice {
		switch i {
		case 0:
			bufWriteF(buf, "%s(%s[:%s], %s)\n", funcName, s.bufferName, vars[i], f.Name())
		default:
			bufWriteF(buf, "%s(%s[%s:%s], %s)\n", funcName, s.bufferName, strings.Join(vars[:i], "+"), strings.Join(vars[:i+1], "+"), f.Name())

		case qty - 1:
			bufWriteF(buf, "%s(%s[%s:], %s)\n", funcName, s.bufferName, strings.Join(vars[:i], "+"), f.Name())
		}
	}
}
