package generate

import (
	"github.com/speedyhoon/jay"
)

func (s *structTyp) calcSize() (qty uint) {
	qty = uint(jay.SizeBools(len(s.bool)))

	qty += uint(len(s.single))

	for _, x := range s.fixedLen {
		qty += x.typeFuncSize()
	}
	for _, v := range s.variableLen {
		qty += v.typeFuncSize()
	}
	return qty
}

// typeFuncSize returns the minimum quantity of bytes required to represent an empty or undefined value.
func (f field) typeFuncSize() (size uint) {
	switch {
	case f.arraySize <= typeSlice:
		return 1
	case f.arraySize >= typeArray:
		itemSize := field{typ: f.arrayType, structTyp: f.structTyp}.typeFuncSize()
		return uint(f.arraySize) * itemSize
	case f.arraySize == typeNotArrayOrSlice:
		switch f.typ {
		case tString, tBool, tByte, tInt8:
			return 1
		case tInt16, tUint16:
			return 2
		case tInt32, tFloat32, tUint32:
			return 4
		case tFloat64, tInt64, tUint64, tTime:
			return 8
		case tInt:
			if f.structTyp.option.FixedIntSize {
				if f.structTyp.option.Is32bit {
					return 4
				}
				return 8
			}
			return 1
		case tUint:
			if f.structTyp.option.FixedUintSize {
				if f.structTyp.option.Is32bit {
					return 4
				}
				return 8
			}
			return 1
		}
	}
	lg.Printf("type %s unhandled in typeFuncSize()", f.typ)
	return
}
