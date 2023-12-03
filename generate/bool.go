package generate

import (
	"bytes"
	"fmt"
)

func (s *Struct) generateBools(b *bytes.Buffer, byteIndex *uint, receiver string) {
	if len(s.bool) == 0 {
		return
	}

	var i, mx uint = 0, uint(len(s.bool) / 8)
	for ; i <= mx; i++ {
		writeBools(s.bool[boolsSliceIndex(i):], b, *byteIndex, receiver)
		*byteIndex++
	}
}

func boolsSliceIndex(input uint) uint {
	if input == 0 {
		return 0
	}
	return ((input-1)/8+1)*8 - 8
}

func writeBools(bools []field, b *bytes.Buffer, byteIndex uint, receiver string) {
	const marshalBoolsFuncPrefix = "Bool"

	if len(bools) > 8 {
		bools = bools[:8]
	}

	b.WriteString(fmt.Sprintf("b[%d] = %s.%s%d(%s)\n", byteIndex, pkgName, marshalBoolsFuncPrefix, len(bools), fieldNames(bools, receiver)))
}
