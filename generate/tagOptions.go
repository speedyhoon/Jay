package generate

import (
	"github.com/speedyhoon/jay"
	"log"
	"strconv"
	"strings"
)

func (f *field) LoadTagOptions() {
	f.tag = strings.TrimSpace(f.tag)
	if f.tag == "" {
		return
	}
	for _, c := range strings.Split(f.tag, ",") {
		d := strings.Split(c, ":")
		switch g := d[0]; g {
		case max:
			loadUint(d[1], &f.tagOptions.Max)
			f.tagOptions.maxBytes = byteSize(f.tagOptions.Max)
		case min:
			loadUint(g, &f.tagOptions.Min)
		}
	}
}

func byteSize(v uint) uint {
	if v == 0 {
		return 0
	}
	if v <= jay.MaxUint8 {
		return 1
	}
	if v <= jay.MaxUint16 {
		return 2
	}
	if v <= jay.MaxUint24 {
		return 3
	}
	if v <= jay.MaxUint32 {
		return 4
	}
	if v <= jay.MaxUint40 {
		return 5
	}
	if v <= jay.MaxUint48 {
		return 6
	}
	if v <= jay.MaxUint56 {
		return 7
	}
	return 8
}

const (
	max = "max"
	min = "min"
)

func loadUint(g string, f *uint) {
	uu, err := strconv.ParseUint(g, 10, 64)
	if err != nil {
		log.Println(err)
	}
	*f = uint(uu)
}
