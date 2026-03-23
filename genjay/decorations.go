package genjay

import (
	"strings"

	"github.com/dave/dst"
)

type structTag int8

const (
	TagIgnore structTag = iota - 2
	TagEmbed
	TagNone
	// TagAlways // Always include struct even when an exclusive list is given (-y flag).

	DecIgnore = "J--"
	DecEmbed  = "J-"
	// DecAlways  = "J+"
)

func (s structTag) HasFuncs() bool {
	return s >= TagNone
}

func (s structTag) HasIgnore() bool {
	return s <= TagIgnore
}

func commentTag(comments dst.Decorations) structTag {
	for _, dec := range comments {
		dec = strings.ToUpper(strings.TrimLeftFunc(dec, IsCommentOrWhitespace))

		if strings.HasPrefix(dec, "J") {
			switch {
			case strings.HasPrefix(dec, DecIgnore):
				return TagIgnore
			case strings.HasPrefix(dec, DecEmbed):
				return TagEmbed
				// case strings.HasPrefix(dec, DecAlways):
				// 	return TagAlways
			}
		}
		break
	}
	return TagNone
}

func IsCommentOrWhitespace(r rune) bool {
	return r <= ' ' || r == '/' || r == '*' || r == 0x7F || r == 0x85 || r == 0xA0
}
