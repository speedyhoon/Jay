package generate

import (
	"bytes"
)

func (s *structTyp) makeWriteBools(b *bytes.Buffer, byteIndex *uint, receiver string, onlyOneByteUsed bool) (isReturnInlined bool) {
	if len(s.bool) == 0 {
		return false
	}

	isReturnInlined = onlyOneByteUsed && len(s.bool) <= 8

	for i := 0; i <= len(s.bool); i += 8 {
		writeBools(s.bool[i:], b, *byteIndex, receiver, isReturnInlined)
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

func writeBools(bools []field, b *bytes.Buffer, byteIndex uint, receiver string, onlyOneByteUsed bool) {
	const marshalBoolsFuncPrefix = "Bool"

	if len(bools) > 8 {
		bools = bools[:8]
	}

	if onlyOneByteUsed {
		bufWriteF(b, "return []byte{%s.%s%d(%s)}\n", pkgName, marshalBoolsFuncPrefix, len(bools), fieldNames(bools, receiver))
	} else {
		bufWriteF(b, "b[%d] = %s.%s%d(%s)\n", byteIndex, pkgName, marshalBoolsFuncPrefix, len(bools), fieldNames(bools, receiver))
	}
}

func (s *structTyp) makeReadBools(b *bytes.Buffer, byteIndex *uint, receiver string) {
	if len(s.bool) == 0 {
		return
	}

	for i := 0; i <= len(s.bool); i += 8 {
		readBools(s.bool[i:], b, *byteIndex, receiver)
		*byteIndex++
	}
}

func readBools(bools []field, b *bytes.Buffer, byteIndex uint, receiver string) {
	const marshalBoolsFuncPrefix = "ReadBool"

	if len(bools) > 8 {
		bools = bools[:8]
	}

	bufWriteF(b, "%s = %s.%s%d(b[%d])\n", fieldNames(bools, receiver), pkgName, marshalBoolsFuncPrefix, len(bools), byteIndex)
}
