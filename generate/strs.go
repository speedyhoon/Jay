package generate

import (
	"bytes"
	"strconv"
	"strings"
)

func (s *structTyp) writeStrSlice(buf *bytes.Buffer, vars []string, byteIndex uint) {
	qty := len(s.stringSlice)
	if qty == 0 {
		return
	}

	const funcName = pkgName + ".WriteStrings16"

	if qty == 1 {
		vars = []string{"l0"}
		bufWriteF(buf, "%s(%s, %s.%s)\n", funcName, s.stringSlice[0].sliceExpr(strconv.Itoa(qty), ""), s.receiver, s.stringSlice[0].name)
		return
	}

	vars = []string{"l0", "l1", "l2"}
	for i, f := range s.stringSlice {
		switch i {
		case 0:
			bufWriteF(buf, "%s(%s[:%s], %s.%s)\n", funcName, s.bufferName, vars[i], s.receiver, f.name)
		default:
			bufWriteF(buf, "%s(%s[%s:%s], %s.%s)\n", funcName, s.bufferName, strings.Join(vars[:i], "+"), strings.Join(vars[:i+1], "+"), s.receiver, f.name)

		case qty - 1:
			bufWriteF(buf, "%s(%s[%s:], %s.%s)\n", funcName, s.bufferName, strings.Join(vars[:i], "+"), s.receiver, f.name)
		}
	}
}
