package generate

import (
	"strconv"
	"strings"
)

const (
	tagMax      = "max"
	tagMin      = "min"
	tagNano     = "nano"
	tagRequired = "required"
)

type tagOptions struct {
	// The maximum and minimum value expected in the variable.
	// Any value out of this range isn't guaranteed to be marshalled or unmarshaled correctly.
	Max, Min tagSize

	maxBytes uint
	TimeNano bool
	Required bool // If "required" appears in the tag, then empty checks are omitted from the generated code.
}

type tagSize uint

func (f *field) LoadTagOptions() {
	f.tag = strings.TrimSpace(f.tag)
	if f.tag == "" {
		return
	}
	for _, c := range strings.Split(f.tag, ",") {
		d := strings.Split(c, ":")
		switch g := strings.ToLower(strings.TrimSpace(d[0])); g {
		case tagMax:
			f.tagOptions.Max.set(d[1])
			f.tagOptions.maxBytes = byteSize(f.tagOptions.Max)
		case tagMin:
			f.tagOptions.Min.set(g)
		case tagNano:
			f.tagOptions.TimeNano = true
		case tagRequired:
			f.tagOptions.Required = true
		default:
			f.tagOptions.Required = isShortRequiredTag(g)
		}
	}
}

// isShortRequiredTag returns true if the tag starts with "r", "req", "require" or "required".
func isShortRequiredTag(tag string) bool {
	if len(tag) > len(tagRequired) {
		return false
	}
	for i := min(len(tag), len(tagRequired)-1); i >= 1; i-- {
		if tag == tagRequired[:i] {
			return true
		}
	}
	return false
}

const (
	MaxUint8  = 1<<8 - 1  // 255
	MaxUint16 = 1<<16 - 1 // 65535
	MaxUint24 = 1<<24 - 1 // 16777215
	MaxUint32 = 1<<32 - 1 // 4294967295
	MaxUint40 = 1<<40 - 1 // 1099511627775
	MaxUint48 = 1<<48 - 1 // 281474976710655
	MaxUint56 = 1<<56 - 1 // 72057594037927935
)

func byteSize(v tagSize) uint {
	switch {
	case v == 0:
		return 0
	case v <= MaxUint8:
		return 1
	case v <= MaxUint16:
		return 2
	case v <= MaxUint24:
		return 3
	case v <= MaxUint32:
		return 4
	case v <= MaxUint40:
		return 5
	case v <= MaxUint48:
		return 6
	case v <= MaxUint56:
		return 7
	}
	return 8
}

func (f *tagSize) set(s string) {
	uu, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		lg.Println(err)
	}
	*f = tagSize(uu)
}
