package genjay

import (
	"fmt"

	"github.com/speedyhoon/utl"
)

type sliceExp struct {
	start, end sliceIndex
}

type sliceIndex uint

func (i sliceIndex) String(omits ...bool) string {
	if i == 0 {
		return ""
	}
	for _, omit := range omits {
		if omit {
			return ""
		}
	}
	return utl.UtoA(uint(i))
}

func newSliceExp(f *field) (x sliceExp) {
	x.start = sliceIndex(*f.indexStart)
	x.end = x.start + sliceIndex(f.elmSize)
	return
}

func (s sliceExp) Print(f *field, omits ...bool) string {
	var omit bool
	for _, o := range omits {
		if o {
			omit = true
			break
		}
	}

	return fmt.Sprintf("%s[%s:%s]",
		f.structTyp.bufferName,
		s.start.String(),
		s.end.String(f.isLast && omit))
}

func (s *sliceExp) Inc(f *field) {
	s.start, s.end = s.end, s.end+sliceIndex(f.elmSize)
}
