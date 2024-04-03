package generate

import (
	"bytes"
)

const marshalBoolsFuncPrefix = "Bool"

func (s *structTyp) makeWriteBools(b *bytes.Buffer, byteIndex *uint, importJ *bool) (isReturnInlined bool) {
	if len(s.bool) == 0 {
		return false
	}

	isReturnInlined = !s.useMakeFunc()
	hasSingles := len(s.single) != 0
	*importJ = true

	l := len(s.bool)
	for i := 0; i < l; i += 8 {
		writeBools(s.bool[i:], b, *byteIndex, s.receiver, s.bufferName, isReturnInlined)
		if isReturnInlined {
			if i+1 < l || i+1 == l && hasSingles {
				b.WriteString(",")
			}
		}
		*byteIndex++
	}

	return isReturnInlined
}

func boolsSliceIndex(input uint) uint {
	if input == 0 {
		return 0
	}
	return ((input-1)/8+1)*8 - 8
}

func writeBools(bools []field, b *bytes.Buffer, byteIndex uint, receiver, bufferName string, isMake bool) {
	if len(bools) > 8 {
		bools = bools[:8]
	}

	if isMake {
		bufWriteF(b, "%s.%s%d(%s)", pkgName, marshalBoolsFuncPrefix, len(bools), fieldNames(bools, receiver, true))
	} else {
		bufWriteF(b, "%s[%d] = %s.%s%d(%s)\n", bufferName, byteIndex, pkgName, marshalBoolsFuncPrefix, len(bools), fieldNames(bools, receiver, true))
	}
}

func (s *structTyp) makeReadBools(b *bytes.Buffer, byteIndex *uint, receiver string) {
	if len(s.bool) == 0 {
		return
	}

	for i := 0; i < len(s.bool); i += 8 {
		readBools(s.bool[i:], b, *byteIndex, receiver, s.bufferName)
		*byteIndex++
	}
}

func fieldNames(fields []field, receiver string, isMarshalling bool) string {
	var s []string
	for i := range fields {
		if isMarshalling && fields[i].typ != fields[i].aliasType {
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

func isUnmarshalInline(bools []field) bool {
	for _, b := range bools {
		if b.typ != b.aliasType {
			return true
		}
	}
	return false
}

func unmarshalBoolsInline(bufferName string, byteIndex uint, qty int) string {
	if qty <= 0 || qty > 8 {
		log.Println("qty must be >= 1 & <= 8", qty)
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
