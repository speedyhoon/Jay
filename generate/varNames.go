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

func (s *structTyp) varsUnmarshal() {
	for i, f := range append(s.stringSlice, s.variableLen...) {
		switch f.typ {
		case tBools:
			f.unmarshal.qtyVar = multiplier(fmt.Sprintf("%s[%d]", f.structTyp.bufferName, uint(i)))
		default:
			f.unmarshal.qtyVar = multiplier(lenVariable(uint(i)))
		}

		f.unmarshal.sizeVar = multiplier(lenVariable(uint(i)))
	}
}

// varsMarshal sets the generated variable names for the quantity and size variables if needed.
func (s *structTyp) varsMarshal() {
	for i, f := range append(s.stringSlice, s.variableLen...) {
		f.marshal.qtyVar = multiplier(lenVariable(uint(i)))
		f.marshal.sizeVar = f.marshal.qtyVar
	}
}
