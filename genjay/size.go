package genjay

import (
	"github.com/speedyhoon/jay"
)

func (s *structTyp) calcSize() (qty uint) {
	qty = uint(jay.SizeBools(len(s.bool)))

	qty += uint(len(s.single))

	for _, x := range s.stringSlice {
		if !x.isArray() {
			// Don't include array lengths because its size is generated as a hardcoded value.
			qty++
		}
	}

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
	if f.isSlice() {
		return f.reserveSizeOf()
	}
	switch f.BaseType() {
	case tString, tBool, tByte, tInt8:
		size = 1
	case tInt16, tUint16:
		size = 2
	case tInt32, tFloat32, tUint32:
		size = 4
	case tFloat64, tInt64, tUint64, tTime, tTimeDuration, tComplex64:
		size = 8
	case tComplex128:
		size = 16
	case tInt:
		if f.structTyp.option.VariableIntSize {
			size = 1
		}
		if f.structTyp.option.Is32bit {
			size = 4
		}
		size = 8
	case tUint:
		if f.structTyp.option.VariableUintSize {
			size = 1
		}
		if f.structTyp.option.Is32bit {
			size = 4
		}
		size = 8
	default:
		lg.Printf("type %s unhandled in typeFuncSize()", f.typ)
	}
	if f.isArray() {
		size *= uint(f.arraySize)
	}
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
