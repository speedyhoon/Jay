package generate

import (
	"bytes"
	"fmt"
	"strings"
)

const marshalBoolsFuncPrefix = "Bool"

func (s *structTyp) makeWriteBools(b *bytes.Buffer, byteIndex *uint, importJ *bool) {
	if len(s.bool) == 0 {
		return
	}

	hasSingles := len(s.single) != 0
	*importJ = true

	if s.returnInline {
		b.WriteString("[]byte{")
	}

	newList := fieldNamesArrays(s.bool, s.receiver)
	l := len(newList)
	for i := 0; i < l; i += 8 {
		next8 := min(8, len(newList[i:]))

		if s.returnInline {
			b.WriteString(boolsFunc(newList, i, next8))
			//bufWriteF(b, "%s.%s%d(%s)", pkgName, marshalBoolsFuncPrefix, next8, strings.Join(newList[i:i+next8], ", "))
		} else {
			bufWriteF(b, "%s[%d] = %s\n", s.bufferName, *byteIndex, boolsFunc(newList, i, next8)) // pkgName, marshalBoolsFuncPrefix, next8, strings.Join(newList[i:i+next8], ", "))
		}

		if s.returnInline {
			if i+1 < l || i+1 == l && hasSingles {
				b.WriteString(",")
			}
		}
		*byteIndex++
	}

	if s.returnInline {
		b.WriteString("}")
	}
}

func boolsFunc(newList []string, i, next8 int) string {
	return fmt.Sprintf("%s.%s%d(%s)", pkgName, marshalBoolsFuncPrefix, next8, strings.Join(newList[i:i+next8], ", "))
}

func (s *structTyp) makeReadBools(b *bytes.Buffer, byteIndex *uint, receiver string) {
	if len(s.bool) == 0 {
		return
	}

	newList, uList := fieldNamesArraysUnmarshalInline(s.bool, s.receiver)

	l := len(newList)
	for i := 0; i < l; i += 8 {
		readBools2(newList[i:], b, *byteIndex, receiver, s.bufferName, uList[i:])
		*byteIndex++
	}
}

func fieldNames(fields []field, receiver string, isMarshalling bool) string {
	var s []string
	for i := range fields {
		if isMarshalling && fields[i].isDef {
			s = append(s, printFunc(fields[i].typ, pkgSelName(receiver, fields[i].name)))
		} else {
			s = append(s, pkgSelName(receiver, fields[i].name))
		}
	}
	return strings.Join(s, ", ")
}

func readBools(bools []field, b *bytes.Buffer, byteIndex uint, receiver string, bufferName string) {
	const marshalBoolsFuncPrefix = "ReadBool"

	if len(bools) > 8 {
		bools = bools[:8]
	}

	if isUnmarshalInline(bools) {
		bufWriteF(b, "%s = %s\n", fieldNames(bools, receiver, false), unmarshalBoolsInline(bufferName, byteIndex, len(bools)))
		return
	}

	bufWriteF(b, "%s = %s.%s%d(%s[%d])\n", fieldNames(bools, receiver, false), pkgName, marshalBoolsFuncPrefix, len(bools), bufferName, byteIndex)
}

func readBools2(bools []string, b *bytes.Buffer, byteIndex uint, receiver string, bufferName string, uList []bool) {
	const marshalBoolsFuncPrefix = "ReadBool"

	if len(bools) > 8 {
		bools = bools[:8]
	}

	if isUnmarshalInline2(uList) {
		bufWriteF(b, "%s = %s\n", strings.Join(bools, ", ") /*fieldNames2(bools, receiver, false)*/, unmarshalBoolsInline(bufferName, byteIndex, len(bools)))
		return
	}

	bufWriteF(b, "%s = %s.%s%d(%s[%d])\n", strings.Join(bools, ", ") /*fieldNames2(bools, receiver, false)*/, pkgName, marshalBoolsFuncPrefix, len(bools), bufferName, byteIndex)
}

func isUnmarshalInline(bools []field) bool {
	for _, b := range bools {
		if b.isDef {
			return true
		}
	}
	return false
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

func fieldNamesArrays(fields []field, receiver string) (s []string) {
	for i := range fields {
		if fields[i].isDef {
			if fields[i].arraySize >= typeArray {
				for j := 0; j < fields[i].arraySize; j++ {
					s = append(s, printFunc(fields[i].typ, pkgSelName(receiver, fmt.Sprintf("%s[%d]", fields[i].name, j))))
				}
			} else {
				s = append(s, printFunc(fields[i].typ, pkgSelName(receiver, fields[i].name)))
			}
		} else {
			if fields[i].arraySize >= typeArray {
				for j := 0; j < fields[i].arraySize; j++ {
					s = append(s, pkgSelName(receiver, fmt.Sprintf("%s[%d]", fields[i].name, j)))
				}
			} else {
				s = append(s, pkgSelName(receiver, fields[i].name))
			}
		}
	}
	return
}

func fieldNamesArraysUnmarshalInline(fields []field, receiver string) (s []string, u []bool) {
	for i := range fields {
		if fields[i].arraySize >= typeArray {
			for j := 0; j < fields[i].arraySize; j++ {
				s = append(s, pkgSelName(receiver, fmt.Sprintf("%s[%d]", fields[i].name, j)))
				u = append(u, fields[i].isDef)
			}
		} else {
			s = append(s, pkgSelName(receiver, fields[i].name))
			u = append(u, fields[i].isDef)
		}
	}
	return
}
