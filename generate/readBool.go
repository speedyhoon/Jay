package generate

import (
	"bytes"
	"fmt"
)

func (s *Struct) generateReadBools(b *bytes.Buffer, byteIndex *uint, receiver string) {
	if len(s.bool) == 0 {
		return
	}

	var i, mx uint = 0, uint(len(s.bool) / 8)
	for ; i <= mx; i++ {
		readBools(s.bool[boolsSliceIndex(i):], b, *byteIndex, receiver)
		*byteIndex++
	}
}

func readBools(bools []field, b *bytes.Buffer, byteIndex uint, receiver string) {
	const marshalBoolsFuncPrefix = "ReadBool"

	if len(bools) > 8 {
		bools = bools[:8]
	}

	b.WriteString(fmt.Sprintf("%s = %s.%s%d(b[%d])\n", fieldNames(bools, receiver), pkgName, marshalBoolsFuncPrefix, len(bools), byteIndex))
}
