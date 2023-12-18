package generate

import (
	"github.com/speedyhoon/jay"
	"log"
	"strconv"
	"strings"
)

type tagOptions struct {
	// The maximum and minimum value expected in the variable.
	// Any value out of this range isn't guaranteed to be marshalled or unmarshaled correctly.
	Max, Min tagSize

	maxBytes uint
	TimeNano bool
}

type tagSize uint

func (f *field) LoadTagOptions() {
	const (
		tagMax  = "max"
		tagMin  = "min"
		tagTime = "nano"
	)

	f.tag = strings.TrimSpace(f.tag)
	if f.tag == "" {
		return
	}
	for _, c := range strings.Split(f.tag, ",") {
		d := strings.Split(c, ":")
		switch g := strings.ToLower(d[0]); g {
		case tagMax:
			f.tagOptions.Max.set(d[1])
			f.tagOptions.maxBytes = byteSize(f.tagOptions.Max)
		case tagMin:
			f.tagOptions.Min.set(g)
		case tagTime:
			f.tagOptions.TimeNano = true
		}
	}
}

func byteSize(v tagSize) uint {
	switch {
	case v == 0:
		return 0
	case v <= jay.MaxUint8:
		return 1
	case v <= jay.MaxUint16:
		return 2
	case v <= jay.MaxUint24:
		return 3
	case v <= jay.MaxUint32:
		return 4
	case v <= jay.MaxUint40:
		return 5
	case v <= jay.MaxUint48:
		return 6
	case v <= jay.MaxUint56:
		return 7
	}
	return 8
}

func (f *tagSize) set(s string) {
	uu, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		log.Println(err)
	}
	*f = tagSize(uu)
}
