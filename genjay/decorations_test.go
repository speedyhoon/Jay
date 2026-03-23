package genjay

import (
	"fmt"
	"strings"
	"testing"

	"github.com/dave/dst"
)

func Test_commentJayFlag(t *testing.T) {
	tests := []struct {
		comments dst.Decorations
		want     structTag
	}{
		{comments: dst.Decorations{"// J--"}, want: TagIgnore},
		{comments: dst.Decorations{"// j--"}, want: TagIgnore},
		{comments: dst.Decorations{"// J-- "}, want: TagIgnore},
		{comments: dst.Decorations{"// j-- "}, want: TagIgnore},
		{comments: dst.Decorations{"/* J--"}, want: TagIgnore},
		{comments: dst.Decorations{"/* j--"}, want: TagIgnore},
		{comments: dst.Decorations{"/* J-- "}, want: TagIgnore},
		{comments: dst.Decorations{"/* j-- "}, want: TagIgnore},

		{comments: dst.Decorations{"", "// J--"}, want: TagNone},
		{comments: dst.Decorations{"", "// J-"}, want: TagNone},
		{comments: dst.Decorations{"", "// J+"}, want: TagNone},
		{comments: dst.Decorations{"", "// J++"}, want: TagNone},

		{comments: dst.Decorations{"// J --"}, want: TagNone},
		{comments: dst.Decorations{"// j --"}, want: TagNone},
		{comments: dst.Decorations{"// J -- "}, want: TagNone},
		{comments: dst.Decorations{"// j -- "}, want: TagNone},
		{comments: dst.Decorations{"/* J --"}, want: TagNone},
		{comments: dst.Decorations{"/* j --"}, want: TagNone},
		{comments: dst.Decorations{"/* J -- "}, want: TagNone},
		{comments: dst.Decorations{"/* j -- "}, want: TagNone},
		{comments: dst.Decorations{"// J - -"}, want: TagNone},
		{comments: dst.Decorations{"// j - -"}, want: TagNone},
		{comments: dst.Decorations{"// J - - "}, want: TagNone},
		{comments: dst.Decorations{"// j - - "}, want: TagNone},
		{comments: dst.Decorations{"/* J - -"}, want: TagNone},
		{comments: dst.Decorations{"/* j - -"}, want: TagNone},
		{comments: dst.Decorations{"/* J - - "}, want: TagNone},
		{comments: dst.Decorations{"/* j - - "}, want: TagNone},
		{comments: dst.Decorations{"// J ++"}, want: TagNone},
		{comments: dst.Decorations{"// j ++"}, want: TagNone},
		{comments: dst.Decorations{"// J ++ "}, want: TagNone},
		{comments: dst.Decorations{"// j ++ "}, want: TagNone},
		{comments: dst.Decorations{"/* J ++"}, want: TagNone},
		{comments: dst.Decorations{"/* j ++"}, want: TagNone},
		{comments: dst.Decorations{"/* J ++ "}, want: TagNone},
		{comments: dst.Decorations{"/* j ++ "}, want: TagNone},
		{comments: dst.Decorations{"// J + +"}, want: TagNone},
		{comments: dst.Decorations{"// j + +"}, want: TagNone},
		{comments: dst.Decorations{"// J + + "}, want: TagNone},
		{comments: dst.Decorations{"// j + + "}, want: TagNone},
		{comments: dst.Decorations{"/* J + +"}, want: TagNone},
		{comments: dst.Decorations{"/* j + +"}, want: TagNone},
		{comments: dst.Decorations{"/* J + + "}, want: TagNone},
		{comments: dst.Decorations{"/* j + + "}, want: TagNone},

		{comments: dst.Decorations{"// J-"}, want: TagEmbed},
		{comments: dst.Decorations{"// j-"}, want: TagEmbed},
		{comments: dst.Decorations{"// J- "}, want: TagEmbed},
		{comments: dst.Decorations{"// j- "}, want: TagEmbed},
		{comments: dst.Decorations{"/* J-"}, want: TagEmbed},
		{comments: dst.Decorations{"/* j-"}, want: TagEmbed},
		{comments: dst.Decorations{"/* J- "}, want: TagEmbed},
		{comments: dst.Decorations{"/* j- "}, want: TagEmbed},
		{comments: dst.Decorations{"// J- -"}, want: TagEmbed},
		{comments: dst.Decorations{"// j- -"}, want: TagEmbed},
		{comments: dst.Decorations{"// J- - "}, want: TagEmbed},
		{comments: dst.Decorations{"// j- - "}, want: TagEmbed},
		{comments: dst.Decorations{"/* J- -"}, want: TagEmbed},
		{comments: dst.Decorations{"/* j- -"}, want: TagEmbed},
		{comments: dst.Decorations{"/* J- - "}, want: TagEmbed},
		{comments: dst.Decorations{"/* j- - "}, want: TagEmbed},

		{comments: dst.Decorations{}, want: TagNone},
		{comments: dst.Decorations{""}, want: TagNone},
		{comments: dst.Decorations{"//"}, want: TagNone},
		{comments: dst.Decorations{"// "}, want: TagNone},
		{comments: dst.Decorations{"j"}, want: TagNone},
		{comments: dst.Decorations{"/* j"}, want: TagNone},
		{comments: dst.Decorations{"/* jBig"}, want: TagNone},
		{comments: dst.Decorations{"/* jess "}, want: TagNone},

		{comments: dst.Decorations{"// J+"}, want: TagNone},
		{comments: dst.Decorations{"// j+"}, want: TagNone},
		{comments: dst.Decorations{"// J+ "}, want: TagNone},
		{comments: dst.Decorations{"// j+ "}, want: TagNone},
		{comments: dst.Decorations{"/* J+"}, want: TagNone},
		{comments: dst.Decorations{"/* j+"}, want: TagNone},
		{comments: dst.Decorations{"/* J+ "}, want: TagNone},
		{comments: dst.Decorations{"/* j+ "}, want: TagNone},
		{comments: dst.Decorations{"// J+ +"}, want: TagNone},
		{comments: dst.Decorations{"// j+ +"}, want: TagNone},
		{comments: dst.Decorations{"// J+ + "}, want: TagNone},
		{comments: dst.Decorations{"// j+ + "}, want: TagNone},
		{comments: dst.Decorations{"/* J+ +"}, want: TagNone},
		{comments: dst.Decorations{"/* j+ +"}, want: TagNone},
		{comments: dst.Decorations{"/* J+ + "}, want: TagNone},
		{comments: dst.Decorations{"/* j+ + "}, want: TagNone},
		{comments: dst.Decorations{"// J++"}, want: TagNone},
		{comments: dst.Decorations{"// j++"}, want: TagNone},
		{comments: dst.Decorations{"// J++ "}, want: TagNone},
		{comments: dst.Decorations{"// j++ "}, want: TagNone},
		{comments: dst.Decorations{"/* J++"}, want: TagNone},
		{comments: dst.Decorations{"/* j++"}, want: TagNone},
		{comments: dst.Decorations{"/* J++ "}, want: TagNone},
		{comments: dst.Decorations{"/* j++ "}, want: TagNone},
	}
	for n, tt := range tests {
		t.Run(fmt.Sprintf("test[%d]__`%s`", n, strings.Join(tt.comments.All(), "`__`")), func(t *testing.T) {
			if got := commentTag(tt.comments); got != tt.want {
				t.Errorf("commentTag(`%s`) = %d, want %d", strings.Join(tt.comments.All(), "`, `"), got, tt.want)
			}
		})
	}
}
