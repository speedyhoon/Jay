package generate

import (
	"fmt"
)

// vars stores names of variables for either marshall or unmarshalling functions.
type vars struct {
	qtyVar  multiplier // Name of variable that stores the quantity of items in the slice.
	sizeVar multiplier // Name of variable that stores how many bytes are required for the slice.
}

type multiplier string

func (v multiplier) String(elmSize uint) string {
	if elmSize <= 1 {
		return string(v)
	}
	return fmt.Sprintf("%s*%d", v, elmSize)
}

func (f *field) varsUnmarshal(index, byteIndex uint) {
	if f.isFixedLen {
		return
	}

	switch f.typ {
	case tBools:
		f.unmarshal.qtyVar = multiplier(fmt.Sprintf("%s[%d]", f.structTyp.bufferName, byteIndex))
	default:
		f.unmarshal.qtyVar = multiplier(lenVariable(int(index)))
	}

	f.unmarshal.sizeVar = multiplier(lenVariable(int(index)))
}

func (f *field) varsMarshal(index, byteIndex uint) {
	if f.isFixedLen {
		return
	}

	f.marshal.qtyVar = multiplier(lenVariable(int(index)))
	f.marshal.sizeVar = multiplier(lenVariable(int(index)))
}
