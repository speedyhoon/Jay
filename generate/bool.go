package generate

import (
	"bytes"
	"fmt"
	"github.com/speedyhoon/jay"
	"strings"
)

var (
	marshalBoolsFuncPrefix   = strings.TrimSuffix(nameOf(jay.Bool1, nil), "1")
	unmarshalBoolsFuncPrefix = strings.TrimSuffix(nameOf(jay.ReadBool1, nil), "1")
)

func (s *structTyp) makeWriteBools(b *bytes.Buffer, byteIndex *uint, importJ *bool) {
	if len(s.bool) == 0 {
		return
	}

	hasSingles := len(s.single) != 0
	*importJ = true

	if s.returnInline {
		bufWriteF(b, "%s{", tBytes)
	}

	newList := fieldNamesArrays(s.bool)
	l := len(newList)
	for i := 0; i < l; i += 8 {
		next8Qty := min(8, len(newList[i:]))

		if s.returnInline {
			b.WriteString(boolsFunc(newList, i, next8Qty))
		} else {
			bufWriteLineF(b, "%s[%d] = %s", s.bufferName, *byteIndex, boolsFunc(newList, i, next8Qty))
		}

		if s.returnInline {
			if i+1 < l || i+1 == l && hasSingles {
				b.WriteString(",")
			}
		}
		*byteIndex++
	}

	if s.returnInline && !hasSingles {
		b.WriteString("}")
	}
}

// boolsFunc generates a function call string to one of jay.BoolX functions depending on the quantity of bools between `i` and `next8Qty`.
func boolsFunc(newList []string, i, next8Qty int) string {
	return fmt.Sprintf("%s%d(%s)", marshalBoolsFuncPrefix, next8Qty, strings.Join(newList[i:i+next8Qty], ", "))
}

func (s *structTyp) makeReadBools(b *bytes.Buffer, byteIndex *uint) {
	if len(s.bool) == 0 {
		return
	}

	newList, uList := fieldNamesArraysUnmarshalInline(s.bool)

	l := len(newList)
	for i := 0; i < l; i += 8 {
		readBools2(newList[i:], b, *byteIndex, s.bufferName, uList[i:])
		*byteIndex++
	}
}

func readBools2(bools []string, b *bytes.Buffer, byteIndex uint, bufferName string, uList []bool) {
	if len(bools) > 8 {
		bools = bools[:8]
	}

	if isUnmarshalInline2(uList) {
		bufWriteLineF(b, "%s = %s", strings.Join(bools, ", "), unmarshalBoolsInline(bufferName, byteIndex, len(bools)))
		return
	}

	bufWriteLineF(b, "%s = %s%d(%s[%d])", strings.Join(bools, ", "), unmarshalBoolsFuncPrefix, len(bools), bufferName, byteIndex)
}

func isUnmarshalInline2(bools []bool) bool {
	for _, b := range bools {
		if b {
			return true
		}
	}
	return false
}

func unmarshalBoolsInline(bufferName string, byteIndex uint, qty int) string {
	if qty <= 0 || qty > 8 {
		lg.Println("qty must be >= 1 & <= 8", qty)
		return ""
	}
	return fmt.Sprintf(strings.Join([]string{
		"%[1]s[%[2]d] >= 128",
		"%[1]s[%[2]d]&64 == 64",
		"%[1]s[%[2]d]&32 == 32",
		"%[1]s[%[2]d]&16 == 16",
		"%[1]s[%[2]d]&8 == 8",
		"%[1]s[%[2]d]&4 == 4",
		"%[1]s[%[2]d]&2 == 2",
		"%[1]s[%[2]d]&1 == 1",
	}[:qty], ", "), bufferName, byteIndex)
}

func fieldNamesArrays(fields []*field) (s []string) {
	for i := range fields {
		if fields[i].isDef {
			if fields[i].isArray() {
				for j := 0; j < fields[i].arraySize; j++ {
					s = append(s, printFunc(fields[i].typ, fmt.Sprintf("%s[%d]", fields[i].Name(), j)))
				}
			} else {
				s = append(s, printFunc(fields[i].typ, fields[i].Name()))
			}
		} else {
			if fields[i].isArray() {
				for j := 0; j < fields[i].arraySize; j++ {
					s = append(s, fmt.Sprintf("%s[%d]", fields[i].Name(), j))
				}
			} else {
				s = append(s, fields[i].Name())
			}
		}
	}
	return
}

func fieldNamesArraysUnmarshalInline(fields fieldList) (s []string, u []bool) {
	for i := range fields {
		if fields[i].isArray() {
			for j := 0; j < fields[i].arraySize; j++ {
				s = append(s, fmt.Sprintf("%s[%d]", fields[i].Name(), j))
				u = append(u, fields[i].isDef)
			}
		} else {
			s = append(s, fields[i].Name())
			u = append(u, fields[i].isDef)
		}
	}
	return
}
