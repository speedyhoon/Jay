package generate

import (
	"github.com/speedyhoon/jay"
)

func (s *structTyp) calcSize() (qty uint) {
	qty = uint(jay.SizeBools(len(s.bool)))

	qty += uint(len(s.single))
	qty += uint(len(s.stringSlice))

	for _, x := range s.fixedLen {
		qty += x.typeFuncSize()
	}
	for _, v := range s.variableLen {
		qty += v.typeFuncSize()
	}
	return qty
}

// typeFuncSize returns the minimum quantity of bytes required to represent an empty or undefined value.
func (f *field) typeFuncSize() (size uint) {
	switch {
	case f.isSlice():
		return f.reserveSizeOf()
	case f.isArray():
		fe := field{typ: f.arrayType, structTyp: f.structTyp}
		return uint(f.arraySize) * fe.typeFuncSize()
	default:
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

// reserveSizeOf returns how many bytes are prefixed to a slice to describe its length.
// Either 1-byte for 0-255 items or 2-bytes for 0-65,535 items.
func (f *field) reserveSizeOf() (size uint) {
	if f.tagOptions.MaxQty <= maxUint8 {
		return 1
	}
	return 2
}

func (f *field) sizeOfPick(small, large any) any {
	if f.tagOptions.MaxQty <= maxUint8 {
		return small
	}
	return large
}
