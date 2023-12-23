package generate

import (
	"bytes"
)

func (s *structTyp) makeWriteBools(b *bytes.Buffer, byteIndex *uint) (isReturnInlined bool) {
	if len(s.bool) == 0 {
		return false
	}

	isReturnInlined = !s.useMakeFunc()
	hasSingles := len(s.single) != 0

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
	const marshalBoolsFuncPrefix = "Bool"

	if len(bools) > 8 {
		bools = bools[:8]
	}

	if isMake {
		bufWriteF(b, "%s.%s%d(%s)", pkgName, marshalBoolsFuncPrefix, len(bools), fieldNames(bools, receiver))
	} else {
		bufWriteF(b, "%s[%d] = %s.%s%d(%s)\n", bufferName, byteIndex, pkgName, marshalBoolsFuncPrefix, len(bools), fieldNames(bools, receiver))
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

func readBools(bools []field, b *bytes.Buffer, byteIndex uint, receiver string, bufferName string) {
	const marshalBoolsFuncPrefix = "ReadBool"

	if len(bools) > 8 {
		bools = bools[:8]
	}

	bufWriteF(b, "%s = %s.%s%d(%s[%d])\n", fieldNames(bools, receiver), pkgName, marshalBoolsFuncPrefix, len(bools), bufferName, byteIndex)
}
