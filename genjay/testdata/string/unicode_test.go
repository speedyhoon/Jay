package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func TestEmoji1(t *testing.T) {
	for l := 0; l <= 63; l++ {
		t.Run(fmt.Sprintf("test[%d]", l), func(t *testing.T) {
			o := One{One: strings.Repeat("😀", l)}
			want := append([]byte{byte(len([]byte(o.One)))}, []byte(o.One)...)
			require.Equal(t, want, o.MarshalJ())
		})
	}
}

func TestEmoji2(t *testing.T) {
	for l := 0; l <= 63; l++ {
		t.Run(fmt.Sprintf("test[%d]", l), func(t *testing.T) {
			o := Two{
				One: strings.Repeat("😀", l),
				Two: strings.Repeat("😆", l),
			}
			want := []byte{
				byte(len([]byte(o.One))),
				byte(len([]byte(o.Two))),
			}
			want = append(want, []byte(o.One)...)
			want = append(want, []byte(o.Two)...)
			require.Equal(t, want, o.MarshalJ())
		})
	}
}

func TestEmoji3(t *testing.T) {
	for l := 0; l <= 63; l++ {
		t.Run(fmt.Sprintf("test[%d]", l), func(t *testing.T) {
			o := Three{
				One:   strings.Repeat("😀", l),
				Two:   strings.Repeat("😆", l),
				Three: strings.Repeat("😀", l),
			}
			want := []byte{
				byte(len([]byte(o.One))),
				byte(len([]byte(o.Two))),
				byte(len([]byte(o.Three))),
			}
			want = append(want, []byte(o.One)...)
			want = append(want, []byte(o.Two)...)
			want = append(want, []byte(o.Three)...)
			require.Equal(t, want, o.MarshalJ())
		})
	}
}

func TestEmoji4(t *testing.T) {
	for l := 0; l <= 63; l++ {
		t.Run(fmt.Sprintf("test[%d]", l), func(t *testing.T) {
			o := Four{
				One:   strings.Repeat("😀", l),
				Two:   strings.Repeat("😆", l),
				Three: strings.Repeat("😀", l),
				Four:  strings.Repeat("😆", l),
			}
			want := []byte{
				byte(len([]byte(o.One))),
				byte(len([]byte(o.Two))),
				byte(len([]byte(o.Three))),
				byte(len([]byte(o.Four))),
			}
			want = append(want, []byte(o.One)...)
			want = append(want, []byte(o.Two)...)
			want = append(want, []byte(o.Three)...)
			want = append(want, []byte(o.Four)...)
			require.Equal(t, want, o.MarshalJ())
		})
	}
}

func TestEmoji5(t *testing.T) {
	for l := 0; l <= 63; l++ {
		t.Run(fmt.Sprintf("test[%d]", l), func(t *testing.T) {
			o := Five{
				One:   strings.Repeat("😀", l),
				Two:   strings.Repeat("😆", l),
				Three: strings.Repeat("😀", l),
				Four:  strings.Repeat("😆", l),
				Five:  strings.Repeat("😀", l),
			}
			want := []byte{
				byte(len([]byte(o.One))),
				byte(len([]byte(o.Two))),
				byte(len([]byte(o.Three))),
				byte(len([]byte(o.Four))),
				byte(len([]byte(o.Five))),
			}
			want = append(want, []byte(o.One)...)
			want = append(want, []byte(o.Two)...)
			want = append(want, []byte(o.Three)...)
			want = append(want, []byte(o.Four)...)
			want = append(want, []byte(o.Five)...)
			require.Equal(t, want, o.MarshalJ())
		})
	}
}

func TestEmoji6(t *testing.T) {
	for l := 0; l <= 63; l++ {
		t.Run(fmt.Sprintf("test[%d]", l), func(t *testing.T) {
			o := Six{
				One:   strings.Repeat("😀", l),
				Two:   strings.Repeat("😆", l),
				Three: strings.Repeat("😀", l),
				Four:  strings.Repeat("😆", l),
				Five:  strings.Repeat("😀", l),
				Six:   strings.Repeat("😆", l),
			}
			want := []byte{
				byte(len([]byte(o.One))),
				byte(len([]byte(o.Two))),
				byte(len([]byte(o.Three))),
				byte(len([]byte(o.Four))),
				byte(len([]byte(o.Five))),
				byte(len([]byte(o.Six))),
			}
			want = append(want, []byte(o.One)...)
			want = append(want, []byte(o.Two)...)
			want = append(want, []byte(o.Three)...)
			want = append(want, []byte(o.Four)...)
			want = append(want, []byte(o.Five)...)
			want = append(want, []byte(o.Six)...)
			require.Equal(t, want, o.MarshalJ())
		})
	}
}

func TestEmoji7(t *testing.T) {
	for l := 0; l <= 63; l++ {
		t.Run(fmt.Sprintf("test[%d]", l), func(t *testing.T) {
			o := Seven{
				One:   strings.Repeat("😀", l),
				Two:   strings.Repeat("😆", l),
				Three: strings.Repeat("😀", l),
				Four:  strings.Repeat("😆", l),
				Five:  strings.Repeat("😀", l),
				Six:   strings.Repeat("😆", l),
				Seven: strings.Repeat("😀", l),
			}
			want := []byte{
				byte(len([]byte(o.One))),
				byte(len([]byte(o.Two))),
				byte(len([]byte(o.Three))),
				byte(len([]byte(o.Four))),
				byte(len([]byte(o.Five))),
				byte(len([]byte(o.Six))),
				byte(len([]byte(o.Seven))),
			}
			want = append(want, []byte(o.One)...)
			want = append(want, []byte(o.Two)...)
			want = append(want, []byte(o.Three)...)
			want = append(want, []byte(o.Four)...)
			want = append(want, []byte(o.Five)...)
			want = append(want, []byte(o.Six)...)
			want = append(want, []byte(o.Seven)...)
			require.Equal(t, want, o.MarshalJ())
		})
	}
}

func TestEmoji8(t *testing.T) {
	for l := 0; l <= 63; l++ {
		t.Run(fmt.Sprintf("test[%d]", l), func(t *testing.T) {
			o := Eight{
				One:   strings.Repeat("😀", l),
				Two:   strings.Repeat("😆", l),
				Three: strings.Repeat("😀", l),
				Four:  strings.Repeat("😆", l),
				Five:  strings.Repeat("😀", l),
				Six:   strings.Repeat("😆", l),
				Seven: strings.Repeat("😀", l),
				Eight: strings.Repeat("😆", l),
			}
			want := []byte{
				byte(len([]byte(o.One))),
				byte(len([]byte(o.Two))),
				byte(len([]byte(o.Three))),
				byte(len([]byte(o.Four))),
				byte(len([]byte(o.Five))),
				byte(len([]byte(o.Six))),
				byte(len([]byte(o.Seven))),
				byte(len([]byte(o.Eight))),
			}
			want = append(want, []byte(o.One)...)
			want = append(want, []byte(o.Two)...)
			want = append(want, []byte(o.Three)...)
			want = append(want, []byte(o.Four)...)
			want = append(want, []byte(o.Five)...)
			want = append(want, []byte(o.Six)...)
			want = append(want, []byte(o.Seven)...)
			want = append(want, []byte(o.Eight)...)
			require.Equal(t, want, o.MarshalJ())
		})
	}
}

func TestEmoji9(t *testing.T) {
	for l := 0; l <= 63; l++ {
		t.Run(fmt.Sprintf("test[%d]", l), func(t *testing.T) {
			o := Nine{
				One:   strings.Repeat("😀", l),
				Two:   strings.Repeat("😆", l),
				Three: strings.Repeat("😀", l),
				Four:  strings.Repeat("😆", l),
				Five:  strings.Repeat("😀", l),
				Six:   strings.Repeat("😆", l),
				Seven: strings.Repeat("😀", l),
				Eight: strings.Repeat("😆", l),
				Nine:  strings.Repeat("😀", l),
			}
			want := []byte{
				byte(len([]byte(o.One))),
				byte(len([]byte(o.Two))),
				byte(len([]byte(o.Three))),
				byte(len([]byte(o.Four))),
				byte(len([]byte(o.Five))),
				byte(len([]byte(o.Six))),
				byte(len([]byte(o.Seven))),
				byte(len([]byte(o.Eight))),
				byte(len([]byte(o.Nine))),
			}
			want = append(want, []byte(o.One)...)
			want = append(want, []byte(o.Two)...)
			want = append(want, []byte(o.Three)...)
			want = append(want, []byte(o.Four)...)
			want = append(want, []byte(o.Five)...)
			want = append(want, []byte(o.Six)...)
			want = append(want, []byte(o.Seven)...)
			want = append(want, []byte(o.Eight)...)
			want = append(want, []byte(o.Nine)...)
			require.Equal(t, want, o.MarshalJ())
		})
	}
}

func TestEmoji10(t *testing.T) {
	for l := 0; l <= 63; l++ {
		t.Run(fmt.Sprintf("test[%d]", l), func(t *testing.T) {
			o := Ten{
				One:   strings.Repeat("😀", l),
				Two:   strings.Repeat("😆", l),
				Three: strings.Repeat("😀", l),
				Four:  strings.Repeat("😆", l),
				Five:  strings.Repeat("😀", l),
				Six:   strings.Repeat("😆", l),
				Seven: strings.Repeat("😀", l),
				Eight: strings.Repeat("😆", l),
				Nine:  strings.Repeat("😀", l),
				Ten:   strings.Repeat("😆", l),
			}
			want := []byte{
				byte(len([]byte(o.One))),
				byte(len([]byte(o.Two))),
				byte(len([]byte(o.Three))),
				byte(len([]byte(o.Four))),
				byte(len([]byte(o.Five))),
				byte(len([]byte(o.Six))),
				byte(len([]byte(o.Seven))),
				byte(len([]byte(o.Eight))),
				byte(len([]byte(o.Nine))),
				byte(len([]byte(o.Ten))),
			}
			want = append(want, []byte(o.One)...)
			want = append(want, []byte(o.Two)...)
			want = append(want, []byte(o.Three)...)
			want = append(want, []byte(o.Four)...)
			want = append(want, []byte(o.Five)...)
			want = append(want, []byte(o.Six)...)
			want = append(want, []byte(o.Seven)...)
			want = append(want, []byte(o.Eight)...)
			want = append(want, []byte(o.Nine)...)
			want = append(want, []byte(o.Ten)...)
			require.Equal(t, want, o.MarshalJ())
		})
	}
}

func TestEmoji11(t *testing.T) {
	for l := 0; l <= 63; l++ {
		t.Run(fmt.Sprintf("test[%d]", l), func(t *testing.T) {
			o := Eleven{
				One:    strings.Repeat("😀", l),
				Two:    strings.Repeat("😆", l),
				Three:  strings.Repeat("😀", l),
				Four:   strings.Repeat("😆", l),
				Five:   strings.Repeat("😀", l),
				Six:    strings.Repeat("😆", l),
				Seven:  strings.Repeat("😀", l),
				Eight:  strings.Repeat("😆", l),
				Nine:   strings.Repeat("😀", l),
				Ten:    strings.Repeat("😆", l),
				Eleven: strings.Repeat("😀", l),
			}
			want := []byte{
				byte(len([]byte(o.One))),
				byte(len([]byte(o.Two))),
				byte(len([]byte(o.Three))),
				byte(len([]byte(o.Four))),
				byte(len([]byte(o.Five))),
				byte(len([]byte(o.Six))),
				byte(len([]byte(o.Seven))),
				byte(len([]byte(o.Eight))),
				byte(len([]byte(o.Nine))),
				byte(len([]byte(o.Ten))),
				byte(len([]byte(o.Eleven))),
			}
			want = append(want, []byte(o.One)...)
			want = append(want, []byte(o.Two)...)
			want = append(want, []byte(o.Three)...)
			want = append(want, []byte(o.Four)...)
			want = append(want, []byte(o.Five)...)
			want = append(want, []byte(o.Six)...)
			want = append(want, []byte(o.Seven)...)
			want = append(want, []byte(o.Eight)...)
			want = append(want, []byte(o.Nine)...)
			want = append(want, []byte(o.Ten)...)
			want = append(want, []byte(o.Eleven)...)
			require.Equal(t, want, o.MarshalJ())
		})
	}
}

func TestEmoji12(t *testing.T) {
	for l := 0; l <= 63; l++ {
		t.Run(fmt.Sprintf("test[%d]", l), func(t *testing.T) {
			o := Twelve{
				One:    strings.Repeat("😀", l),
				Two:    strings.Repeat("😆", l),
				Three:  strings.Repeat("😀", l),
				Four:   strings.Repeat("😆", l),
				Five:   strings.Repeat("😀", l),
				Six:    strings.Repeat("😆", l),
				Seven:  strings.Repeat("😀", l),
				Eight:  strings.Repeat("😆", l),
				Nine:   strings.Repeat("😀", l),
				Ten:    strings.Repeat("😆", l),
				Eleven: strings.Repeat("😀", l),
				Twelve: strings.Repeat("😆", l),
			}
			want := []byte{
				byte(len([]byte(o.One))),
				byte(len([]byte(o.Two))),
				byte(len([]byte(o.Three))),
				byte(len([]byte(o.Four))),
				byte(len([]byte(o.Five))),
				byte(len([]byte(o.Six))),
				byte(len([]byte(o.Seven))),
				byte(len([]byte(o.Eight))),
				byte(len([]byte(o.Nine))),
				byte(len([]byte(o.Ten))),
				byte(len([]byte(o.Eleven))),
				byte(len([]byte(o.Twelve))),
			}
			want = append(want, []byte(o.One)...)
			want = append(want, []byte(o.Two)...)
			want = append(want, []byte(o.Three)...)
			want = append(want, []byte(o.Four)...)
			want = append(want, []byte(o.Five)...)
			want = append(want, []byte(o.Six)...)
			want = append(want, []byte(o.Seven)...)
			want = append(want, []byte(o.Eight)...)
			want = append(want, []byte(o.Nine)...)
			want = append(want, []byte(o.Ten)...)
			want = append(want, []byte(o.Eleven)...)
			want = append(want, []byte(o.Twelve)...)
			require.Equal(t, want, o.MarshalJ())
		})
	}
}

func TestEmoji13(t *testing.T) {
	for l := 0; l <= 63; l++ {
		t.Run(fmt.Sprintf("test[%d]", l), func(t *testing.T) {
			o := Thirteen{
				One:      strings.Repeat("😀", l),
				Two:      strings.Repeat("😆", l),
				Three:    strings.Repeat("😀", l),
				Four:     strings.Repeat("😆", l),
				Five:     strings.Repeat("😀", l),
				Six:      strings.Repeat("😆", l),
				Seven:    strings.Repeat("😀", l),
				Eight:    strings.Repeat("😆", l),
				Nine:     strings.Repeat("😀", l),
				Ten:      strings.Repeat("😆", l),
				Eleven:   strings.Repeat("😀", l),
				Twelve:   strings.Repeat("😆", l),
				Thirteen: strings.Repeat("😀", l),
			}
			want := []byte{
				byte(len([]byte(o.One))),
				byte(len([]byte(o.Two))),
				byte(len([]byte(o.Three))),
				byte(len([]byte(o.Four))),
				byte(len([]byte(o.Five))),
				byte(len([]byte(o.Six))),
				byte(len([]byte(o.Seven))),
				byte(len([]byte(o.Eight))),
				byte(len([]byte(o.Nine))),
				byte(len([]byte(o.Ten))),
				byte(len([]byte(o.Eleven))),
				byte(len([]byte(o.Twelve))),
				byte(len([]byte(o.Thirteen))),
			}
			want = append(want, []byte(o.One)...)
			want = append(want, []byte(o.Two)...)
			want = append(want, []byte(o.Three)...)
			want = append(want, []byte(o.Four)...)
			want = append(want, []byte(o.Five)...)
			want = append(want, []byte(o.Six)...)
			want = append(want, []byte(o.Seven)...)
			want = append(want, []byte(o.Eight)...)
			want = append(want, []byte(o.Nine)...)
			want = append(want, []byte(o.Ten)...)
			want = append(want, []byte(o.Eleven)...)
			want = append(want, []byte(o.Twelve)...)
			want = append(want, []byte(o.Thirteen)...)
			require.Equal(t, want, o.MarshalJ())
		})
	}
}

func TestEmoji14(t *testing.T) {
	for l := 0; l <= 63; l++ {
		t.Run(fmt.Sprintf("test[%d]", l), func(t *testing.T) {
			o := Fourteen{
				One:      strings.Repeat("😀", l),
				Two:      strings.Repeat("😆", l),
				Three:    strings.Repeat("😀", l),
				Four:     strings.Repeat("😆", l),
				Five:     strings.Repeat("😀", l),
				Six:      strings.Repeat("😆", l),
				Seven:    strings.Repeat("😀", l),
				Eight:    strings.Repeat("😆", l),
				Nine:     strings.Repeat("😀", l),
				Ten:      strings.Repeat("😆", l),
				Eleven:   strings.Repeat("😀", l),
				Twelve:   strings.Repeat("😆", l),
				Thirteen: strings.Repeat("😀", l),
				Fourteen: strings.Repeat("😆", l),
			}
			want := []byte{
				byte(len([]byte(o.One))),
				byte(len([]byte(o.Two))),
				byte(len([]byte(o.Three))),
				byte(len([]byte(o.Four))),
				byte(len([]byte(o.Five))),
				byte(len([]byte(o.Six))),
				byte(len([]byte(o.Seven))),
				byte(len([]byte(o.Eight))),
				byte(len([]byte(o.Nine))),
				byte(len([]byte(o.Ten))),
				byte(len([]byte(o.Eleven))),
				byte(len([]byte(o.Twelve))),
				byte(len([]byte(o.Thirteen))),
				byte(len([]byte(o.Fourteen))),
			}
			want = append(want, []byte(o.One)...)
			want = append(want, []byte(o.Two)...)
			want = append(want, []byte(o.Three)...)
			want = append(want, []byte(o.Four)...)
			want = append(want, []byte(o.Five)...)
			want = append(want, []byte(o.Six)...)
			want = append(want, []byte(o.Seven)...)
			want = append(want, []byte(o.Eight)...)
			want = append(want, []byte(o.Nine)...)
			want = append(want, []byte(o.Ten)...)
			want = append(want, []byte(o.Eleven)...)
			want = append(want, []byte(o.Twelve)...)
			want = append(want, []byte(o.Thirteen)...)
			want = append(want, []byte(o.Fourteen)...)
			require.Equal(t, want, o.MarshalJ())
		})
	}
}

func TestEmoji15(t *testing.T) {
	for l := 0; l <= 63; l++ {
		t.Run(fmt.Sprintf("test[%d]", l), func(t *testing.T) {
			o := Fifteen{
				One:      strings.Repeat("😀", l),
				Two:      strings.Repeat("😆", l),
				Three:    strings.Repeat("😀", l),
				Four:     strings.Repeat("😆", l),
				Five:     strings.Repeat("😀", l),
				Six:      strings.Repeat("😆", l),
				Seven:    strings.Repeat("😀", l),
				Eight:    strings.Repeat("😆", l),
				Nine:     strings.Repeat("😀", l),
				Ten:      strings.Repeat("😆", l),
				Eleven:   strings.Repeat("😀", l),
				Twelve:   strings.Repeat("😆", l),
				Thirteen: strings.Repeat("😀", l),
				Fourteen: strings.Repeat("😆", l),
				Fifteen:  strings.Repeat("😀", l),
			}
			want := []byte{
				byte(len([]byte(o.One))),
				byte(len([]byte(o.Two))),
				byte(len([]byte(o.Three))),
				byte(len([]byte(o.Four))),
				byte(len([]byte(o.Five))),
				byte(len([]byte(o.Six))),
				byte(len([]byte(o.Seven))),
				byte(len([]byte(o.Eight))),
				byte(len([]byte(o.Nine))),
				byte(len([]byte(o.Ten))),
				byte(len([]byte(o.Eleven))),
				byte(len([]byte(o.Twelve))),
				byte(len([]byte(o.Thirteen))),
				byte(len([]byte(o.Fourteen))),
				byte(len([]byte(o.Fifteen))),
			}
			want = append(want, []byte(o.One)...)
			want = append(want, []byte(o.Two)...)
			want = append(want, []byte(o.Three)...)
			want = append(want, []byte(o.Four)...)
			want = append(want, []byte(o.Five)...)
			want = append(want, []byte(o.Six)...)
			want = append(want, []byte(o.Seven)...)
			want = append(want, []byte(o.Eight)...)
			want = append(want, []byte(o.Nine)...)
			want = append(want, []byte(o.Ten)...)
			want = append(want, []byte(o.Eleven)...)
			want = append(want, []byte(o.Twelve)...)
			want = append(want, []byte(o.Thirteen)...)
			want = append(want, []byte(o.Fourteen)...)
			want = append(want, []byte(o.Fifteen)...)
			require.Equal(t, want, o.MarshalJ())
		})
	}
}

func TestEmoji16(t *testing.T) {
	for l := 0; l <= 63; l++ {
		t.Run(fmt.Sprintf("test[%d]", l), func(t *testing.T) {
			o := Sixteen{
				One:      strings.Repeat("😀", l),
				Two:      strings.Repeat("😆", l),
				Three:    strings.Repeat("😀", l),
				Four:     strings.Repeat("😆", l),
				Five:     strings.Repeat("😀", l),
				Six:      strings.Repeat("😆", l),
				Seven:    strings.Repeat("😀", l),
				Eight:    strings.Repeat("😆", l),
				Nine:     strings.Repeat("😀", l),
				Ten:      strings.Repeat("😆", l),
				Eleven:   strings.Repeat("😀", l),
				Twelve:   strings.Repeat("😆", l),
				Thirteen: strings.Repeat("😀", l),
				Fourteen: strings.Repeat("😆", l),
				Fifteen:  strings.Repeat("😀", l),
				Sixteen:  strings.Repeat("😆", l),
			}
			want := []byte{
				byte(len([]byte(o.One))),
				byte(len([]byte(o.Two))),
				byte(len([]byte(o.Three))),
				byte(len([]byte(o.Four))),
				byte(len([]byte(o.Five))),
				byte(len([]byte(o.Six))),
				byte(len([]byte(o.Seven))),
				byte(len([]byte(o.Eight))),
				byte(len([]byte(o.Nine))),
				byte(len([]byte(o.Ten))),
				byte(len([]byte(o.Eleven))),
				byte(len([]byte(o.Twelve))),
				byte(len([]byte(o.Thirteen))),
				byte(len([]byte(o.Fourteen))),
				byte(len([]byte(o.Fifteen))),
				byte(len([]byte(o.Sixteen))),
			}
			want = append(want, []byte(o.One)...)
			want = append(want, []byte(o.Two)...)
			want = append(want, []byte(o.Three)...)
			want = append(want, []byte(o.Four)...)
			want = append(want, []byte(o.Five)...)
			want = append(want, []byte(o.Six)...)
			want = append(want, []byte(o.Seven)...)
			want = append(want, []byte(o.Eight)...)
			want = append(want, []byte(o.Nine)...)
			want = append(want, []byte(o.Ten)...)
			want = append(want, []byte(o.Eleven)...)
			want = append(want, []byte(o.Twelve)...)
			want = append(want, []byte(o.Thirteen)...)
			want = append(want, []byte(o.Fourteen)...)
			want = append(want, []byte(o.Fifteen)...)
			want = append(want, []byte(o.Sixteen)...)
			require.Equal(t, want, o.MarshalJ())
		})
	}
}

func TestEmoji17(t *testing.T) {
	for l := 0; l <= 63; l++ {
		t.Run(fmt.Sprintf("test[%d]", l), func(t *testing.T) {
			o := Seventeen{
				One:       strings.Repeat("😀", l),
				Two:       strings.Repeat("😆", l),
				Three:     strings.Repeat("😀", l),
				Four:      strings.Repeat("😆", l),
				Five:      strings.Repeat("😀", l),
				Six:       strings.Repeat("😆", l),
				Seven:     strings.Repeat("😀", l),
				Eight:     strings.Repeat("😆", l),
				Nine:      strings.Repeat("😀", l),
				Ten:       strings.Repeat("😆", l),
				Eleven:    strings.Repeat("😀", l),
				Twelve:    strings.Repeat("😆", l),
				Thirteen:  strings.Repeat("😀", l),
				Fourteen:  strings.Repeat("😆", l),
				Fifteen:   strings.Repeat("😀", l),
				Sixteen:   strings.Repeat("😆", l),
				Seventeen: strings.Repeat("😀", l),
			}
			want := []byte{
				byte(len([]byte(o.One))),
				byte(len([]byte(o.Two))),
				byte(len([]byte(o.Three))),
				byte(len([]byte(o.Four))),
				byte(len([]byte(o.Five))),
				byte(len([]byte(o.Six))),
				byte(len([]byte(o.Seven))),
				byte(len([]byte(o.Eight))),
				byte(len([]byte(o.Nine))),
				byte(len([]byte(o.Ten))),
				byte(len([]byte(o.Eleven))),
				byte(len([]byte(o.Twelve))),
				byte(len([]byte(o.Thirteen))),
				byte(len([]byte(o.Fourteen))),
				byte(len([]byte(o.Fifteen))),
				byte(len([]byte(o.Sixteen))),
				byte(len([]byte(o.Seventeen))),
			}
			want = append(want, []byte(o.One)...)
			want = append(want, []byte(o.Two)...)
			want = append(want, []byte(o.Three)...)
			want = append(want, []byte(o.Four)...)
			want = append(want, []byte(o.Five)...)
			want = append(want, []byte(o.Six)...)
			want = append(want, []byte(o.Seven)...)
			want = append(want, []byte(o.Eight)...)
			want = append(want, []byte(o.Nine)...)
			want = append(want, []byte(o.Ten)...)
			want = append(want, []byte(o.Eleven)...)
			want = append(want, []byte(o.Twelve)...)
			want = append(want, []byte(o.Thirteen)...)
			want = append(want, []byte(o.Fourteen)...)
			want = append(want, []byte(o.Fifteen)...)
			want = append(want, []byte(o.Sixteen)...)
			want = append(want, []byte(o.Seventeen)...)
			require.Equal(t, want, o.MarshalJ())
		})
	}
}

func TestEmoji18(t *testing.T) {
	for l := 0; l <= 63; l++ {
		t.Run(fmt.Sprintf("test[%d]", l), func(t *testing.T) {
			o := Eighteen{
				One:       strings.Repeat("😀", l),
				Two:       strings.Repeat("😆", l),
				Three:     strings.Repeat("😀", l),
				Four:      strings.Repeat("😆", l),
				Five:      strings.Repeat("😀", l),
				Six:       strings.Repeat("😆", l),
				Seven:     strings.Repeat("😀", l),
				Eight:     strings.Repeat("😆", l),
				Nine:      strings.Repeat("😀", l),
				Ten:       strings.Repeat("😆", l),
				Eleven:    strings.Repeat("😀", l),
				Twelve:    strings.Repeat("😆", l),
				Thirteen:  strings.Repeat("😀", l),
				Fourteen:  strings.Repeat("😆", l),
				Fifteen:   strings.Repeat("😀", l),
				Sixteen:   strings.Repeat("😆", l),
				Seventeen: strings.Repeat("😀", l),
				Eighteen:  strings.Repeat("😆", l),
			}
			want := []byte{
				byte(len([]byte(o.One))),
				byte(len([]byte(o.Two))),
				byte(len([]byte(o.Three))),
				byte(len([]byte(o.Four))),
				byte(len([]byte(o.Five))),
				byte(len([]byte(o.Six))),
				byte(len([]byte(o.Seven))),
				byte(len([]byte(o.Eight))),
				byte(len([]byte(o.Nine))),
				byte(len([]byte(o.Ten))),
				byte(len([]byte(o.Eleven))),
				byte(len([]byte(o.Twelve))),
				byte(len([]byte(o.Thirteen))),
				byte(len([]byte(o.Fourteen))),
				byte(len([]byte(o.Fifteen))),
				byte(len([]byte(o.Sixteen))),
				byte(len([]byte(o.Seventeen))),
				byte(len([]byte(o.Eighteen))),
			}
			want = append(want, []byte(o.One)...)
			want = append(want, []byte(o.Two)...)
			want = append(want, []byte(o.Three)...)
			want = append(want, []byte(o.Four)...)
			want = append(want, []byte(o.Five)...)
			want = append(want, []byte(o.Six)...)
			want = append(want, []byte(o.Seven)...)
			want = append(want, []byte(o.Eight)...)
			want = append(want, []byte(o.Nine)...)
			want = append(want, []byte(o.Ten)...)
			want = append(want, []byte(o.Eleven)...)
			want = append(want, []byte(o.Twelve)...)
			want = append(want, []byte(o.Thirteen)...)
			want = append(want, []byte(o.Fourteen)...)
			want = append(want, []byte(o.Fifteen)...)
			want = append(want, []byte(o.Sixteen)...)
			want = append(want, []byte(o.Seventeen)...)
			want = append(want, []byte(o.Eighteen)...)
			require.Equal(t, want, o.MarshalJ())
		})
	}
}

func TestEmoji19(t *testing.T) {
	for l := 0; l <= 63; l++ {
		t.Run(fmt.Sprintf("test[%d]", l), func(t *testing.T) {
			o := Nineteen{
				One:       strings.Repeat("😀", l),
				Two:       strings.Repeat("😆", l),
				Three:     strings.Repeat("😀", l),
				Four:      strings.Repeat("😆", l),
				Five:      strings.Repeat("😀", l),
				Six:       strings.Repeat("😆", l),
				Seven:     strings.Repeat("😀", l),
				Eight:     strings.Repeat("😆", l),
				Nine:      strings.Repeat("😀", l),
				Ten:       strings.Repeat("😆", l),
				Eleven:    strings.Repeat("😀", l),
				Twelve:    strings.Repeat("😆", l),
				Thirteen:  strings.Repeat("😀", l),
				Fourteen:  strings.Repeat("😆", l),
				Fifteen:   strings.Repeat("😀", l),
				Sixteen:   strings.Repeat("😆", l),
				Seventeen: strings.Repeat("😀", l),
				Eighteen:  strings.Repeat("😆", l),
				Nineteen:  strings.Repeat("😀", l),
			}
			want := []byte{
				byte(len([]byte(o.One))),
				byte(len([]byte(o.Two))),
				byte(len([]byte(o.Three))),
				byte(len([]byte(o.Four))),
				byte(len([]byte(o.Five))),
				byte(len([]byte(o.Six))),
				byte(len([]byte(o.Seven))),
				byte(len([]byte(o.Eight))),
				byte(len([]byte(o.Nine))),
				byte(len([]byte(o.Ten))),
				byte(len([]byte(o.Eleven))),
				byte(len([]byte(o.Twelve))),
				byte(len([]byte(o.Thirteen))),
				byte(len([]byte(o.Fourteen))),
				byte(len([]byte(o.Fifteen))),
				byte(len([]byte(o.Sixteen))),
				byte(len([]byte(o.Seventeen))),
				byte(len([]byte(o.Eighteen))),
				byte(len([]byte(o.Nineteen))),
			}
			want = append(want, []byte(o.One)...)
			want = append(want, []byte(o.Two)...)
			want = append(want, []byte(o.Three)...)
			want = append(want, []byte(o.Four)...)
			want = append(want, []byte(o.Five)...)
			want = append(want, []byte(o.Six)...)
			want = append(want, []byte(o.Seven)...)
			want = append(want, []byte(o.Eight)...)
			want = append(want, []byte(o.Nine)...)
			want = append(want, []byte(o.Ten)...)
			want = append(want, []byte(o.Eleven)...)
			want = append(want, []byte(o.Twelve)...)
			want = append(want, []byte(o.Thirteen)...)
			want = append(want, []byte(o.Fourteen)...)
			want = append(want, []byte(o.Fifteen)...)
			want = append(want, []byte(o.Sixteen)...)
			want = append(want, []byte(o.Seventeen)...)
			want = append(want, []byte(o.Eighteen)...)
			want = append(want, []byte(o.Nineteen)...)
			require.Equal(t, want, o.MarshalJ())
		})
	}
}

func TestEmoji20(t *testing.T) {
	for l := 0; l <= 63; l++ {
		t.Run(fmt.Sprintf("test[%d]", l), func(t *testing.T) {
			o := Twenty{
				One:       strings.Repeat("😀", l),
				Two:       strings.Repeat("😆", l),
				Three:     strings.Repeat("😀", l),
				Four:      strings.Repeat("😆", l),
				Five:      strings.Repeat("😀", l),
				Six:       strings.Repeat("😆", l),
				Seven:     strings.Repeat("😀", l),
				Eight:     strings.Repeat("😆", l),
				Nine:      strings.Repeat("😀", l),
				Ten:       strings.Repeat("😆", l),
				Eleven:    strings.Repeat("😀", l),
				Twelve:    strings.Repeat("😆", l),
				Thirteen:  strings.Repeat("😀", l),
				Fourteen:  strings.Repeat("😆", l),
				Fifteen:   strings.Repeat("😀", l),
				Sixteen:   strings.Repeat("😆", l),
				Seventeen: strings.Repeat("😀", l),
				Eighteen:  strings.Repeat("😆", l),
				Nineteen:  strings.Repeat("😀", l),
				Twenty:    strings.Repeat("😆", l),
			}
			want := []byte{
				byte(len([]byte(o.One))),
				byte(len([]byte(o.Two))),
				byte(len([]byte(o.Three))),
				byte(len([]byte(o.Four))),
				byte(len([]byte(o.Five))),
				byte(len([]byte(o.Six))),
				byte(len([]byte(o.Seven))),
				byte(len([]byte(o.Eight))),
				byte(len([]byte(o.Nine))),
				byte(len([]byte(o.Ten))),
				byte(len([]byte(o.Eleven))),
				byte(len([]byte(o.Twelve))),
				byte(len([]byte(o.Thirteen))),
				byte(len([]byte(o.Fourteen))),
				byte(len([]byte(o.Fifteen))),
				byte(len([]byte(o.Sixteen))),
				byte(len([]byte(o.Seventeen))),
				byte(len([]byte(o.Eighteen))),
				byte(len([]byte(o.Nineteen))),
				byte(len([]byte(o.Twenty))),
			}
			want = append(want, []byte(o.One)...)
			want = append(want, []byte(o.Two)...)
			want = append(want, []byte(o.Three)...)
			want = append(want, []byte(o.Four)...)
			want = append(want, []byte(o.Five)...)
			want = append(want, []byte(o.Six)...)
			want = append(want, []byte(o.Seven)...)
			want = append(want, []byte(o.Eight)...)
			want = append(want, []byte(o.Nine)...)
			want = append(want, []byte(o.Ten)...)
			want = append(want, []byte(o.Eleven)...)
			want = append(want, []byte(o.Twelve)...)
			want = append(want, []byte(o.Thirteen)...)
			want = append(want, []byte(o.Fourteen)...)
			want = append(want, []byte(o.Fifteen)...)
			want = append(want, []byte(o.Sixteen)...)
			want = append(want, []byte(o.Seventeen)...)
			want = append(want, []byte(o.Eighteen)...)
			want = append(want, []byte(o.Nineteen)...)
			want = append(want, []byte(o.Twenty)...)
			require.Equal(t, want, o.MarshalJ())
		})
	}
}

func TestEmoji21(t *testing.T) {
	for l := 0; l <= 63; l++ {
		t.Run(fmt.Sprintf("test[%d]", l), func(t *testing.T) {
			o := TwentyOne{
				One:       strings.Repeat("😀", l),
				Two:       strings.Repeat("😆", l),
				Three:     strings.Repeat("😀", l),
				Four:      strings.Repeat("😆", l),
				Five:      strings.Repeat("😀", l),
				Six:       strings.Repeat("😆", l),
				Seven:     strings.Repeat("😀", l),
				Eight:     strings.Repeat("😆", l),
				Nine:      strings.Repeat("😀", l),
				Ten:       strings.Repeat("😆", l),
				Eleven:    strings.Repeat("😀", l),
				Twelve:    strings.Repeat("😆", l),
				Thirteen:  strings.Repeat("😀", l),
				Fourteen:  strings.Repeat("😆", l),
				Fifteen:   strings.Repeat("😀", l),
				Sixteen:   strings.Repeat("😆", l),
				Seventeen: strings.Repeat("😀", l),
				Eighteen:  strings.Repeat("😆", l),
				Nineteen:  strings.Repeat("😀", l),
				Twenty:    strings.Repeat("😆", l),
				TwentyOne: strings.Repeat("😀", l),
			}
			want := []byte{
				byte(len([]byte(o.One))),
				byte(len([]byte(o.Two))),
				byte(len([]byte(o.Three))),
				byte(len([]byte(o.Four))),
				byte(len([]byte(o.Five))),
				byte(len([]byte(o.Six))),
				byte(len([]byte(o.Seven))),
				byte(len([]byte(o.Eight))),
				byte(len([]byte(o.Nine))),
				byte(len([]byte(o.Ten))),
				byte(len([]byte(o.Eleven))),
				byte(len([]byte(o.Twelve))),
				byte(len([]byte(o.Thirteen))),
				byte(len([]byte(o.Fourteen))),
				byte(len([]byte(o.Fifteen))),
				byte(len([]byte(o.Sixteen))),
				byte(len([]byte(o.Seventeen))),
				byte(len([]byte(o.Eighteen))),
				byte(len([]byte(o.Nineteen))),
				byte(len([]byte(o.Twenty))),
				byte(len([]byte(o.TwentyOne))),
			}
			want = append(want, []byte(o.One)...)
			want = append(want, []byte(o.Two)...)
			want = append(want, []byte(o.Three)...)
			want = append(want, []byte(o.Four)...)
			want = append(want, []byte(o.Five)...)
			want = append(want, []byte(o.Six)...)
			want = append(want, []byte(o.Seven)...)
			want = append(want, []byte(o.Eight)...)
			want = append(want, []byte(o.Nine)...)
			want = append(want, []byte(o.Ten)...)
			want = append(want, []byte(o.Eleven)...)
			want = append(want, []byte(o.Twelve)...)
			want = append(want, []byte(o.Thirteen)...)
			want = append(want, []byte(o.Fourteen)...)
			want = append(want, []byte(o.Fifteen)...)
			want = append(want, []byte(o.Sixteen)...)
			want = append(want, []byte(o.Seventeen)...)
			want = append(want, []byte(o.Eighteen)...)
			want = append(want, []byte(o.Nineteen)...)
			want = append(want, []byte(o.Twenty)...)
			want = append(want, []byte(o.TwentyOne)...)
			require.Equal(t, want, o.MarshalJ())
		})
	}
}

func TestEmoji22(t *testing.T) {
	for l := 0; l <= 63; l++ {
		t.Run(fmt.Sprintf("test[%d]", l), func(t *testing.T) {
			o := TwentyTwo{
				One:       strings.Repeat("😀", l),
				Two:       strings.Repeat("😆", l),
				Three:     strings.Repeat("😀", l),
				Four:      strings.Repeat("😆", l),
				Five:      strings.Repeat("😀", l),
				Six:       strings.Repeat("😆", l),
				Seven:     strings.Repeat("😀", l),
				Eight:     strings.Repeat("😆", l),
				Nine:      strings.Repeat("😀", l),
				Ten:       strings.Repeat("😆", l),
				Eleven:    strings.Repeat("😀", l),
				Twelve:    strings.Repeat("😆", l),
				Thirteen:  strings.Repeat("😀", l),
				Fourteen:  strings.Repeat("😆", l),
				Fifteen:   strings.Repeat("😀", l),
				Sixteen:   strings.Repeat("😆", l),
				Seventeen: strings.Repeat("😀", l),
				Eighteen:  strings.Repeat("😆", l),
				Nineteen:  strings.Repeat("😀", l),
				Twenty:    strings.Repeat("😆", l),
				TwentyOne: strings.Repeat("😀", l),
				TwentyTwo: strings.Repeat("😆", l),
			}
			want := []byte{
				byte(len([]byte(o.One))),
				byte(len([]byte(o.Two))),
				byte(len([]byte(o.Three))),
				byte(len([]byte(o.Four))),
				byte(len([]byte(o.Five))),
				byte(len([]byte(o.Six))),
				byte(len([]byte(o.Seven))),
				byte(len([]byte(o.Eight))),
				byte(len([]byte(o.Nine))),
				byte(len([]byte(o.Ten))),
				byte(len([]byte(o.Eleven))),
				byte(len([]byte(o.Twelve))),
				byte(len([]byte(o.Thirteen))),
				byte(len([]byte(o.Fourteen))),
				byte(len([]byte(o.Fifteen))),
				byte(len([]byte(o.Sixteen))),
				byte(len([]byte(o.Seventeen))),
				byte(len([]byte(o.Eighteen))),
				byte(len([]byte(o.Nineteen))),
				byte(len([]byte(o.Twenty))),
				byte(len([]byte(o.TwentyOne))),
				byte(len([]byte(o.TwentyTwo))),
			}
			want = append(want, []byte(o.One)...)
			want = append(want, []byte(o.Two)...)
			want = append(want, []byte(o.Three)...)
			want = append(want, []byte(o.Four)...)
			want = append(want, []byte(o.Five)...)
			want = append(want, []byte(o.Six)...)
			want = append(want, []byte(o.Seven)...)
			want = append(want, []byte(o.Eight)...)
			want = append(want, []byte(o.Nine)...)
			want = append(want, []byte(o.Ten)...)
			want = append(want, []byte(o.Eleven)...)
			want = append(want, []byte(o.Twelve)...)
			want = append(want, []byte(o.Thirteen)...)
			want = append(want, []byte(o.Fourteen)...)
			want = append(want, []byte(o.Fifteen)...)
			want = append(want, []byte(o.Sixteen)...)
			want = append(want, []byte(o.Seventeen)...)
			want = append(want, []byte(o.Eighteen)...)
			want = append(want, []byte(o.Nineteen)...)
			want = append(want, []byte(o.Twenty)...)
			want = append(want, []byte(o.TwentyOne)...)
			want = append(want, []byte(o.TwentyTwo)...)
			require.Equal(t, want, o.MarshalJ())
		})
	}
}

func TestEmoji23(t *testing.T) {
	for l := 0; l <= 63; l++ {
		t.Run(fmt.Sprintf("test[%d]", l), func(t *testing.T) {
			o := TwentyThree{
				One:         strings.Repeat("😀", l),
				Two:         strings.Repeat("😆", l),
				Three:       strings.Repeat("😀", l),
				Four:        strings.Repeat("😆", l),
				Five:        strings.Repeat("😀", l),
				Six:         strings.Repeat("😆", l),
				Seven:       strings.Repeat("😀", l),
				Eight:       strings.Repeat("😆", l),
				Nine:        strings.Repeat("😀", l),
				Ten:         strings.Repeat("😆", l),
				Eleven:      strings.Repeat("😀", l),
				Twelve:      strings.Repeat("😆", l),
				Thirteen:    strings.Repeat("😀", l),
				Fourteen:    strings.Repeat("😆", l),
				Fifteen:     strings.Repeat("😀", l),
				Sixteen:     strings.Repeat("😆", l),
				Seventeen:   strings.Repeat("😀", l),
				Eighteen:    strings.Repeat("😆", l),
				Nineteen:    strings.Repeat("😀", l),
				Twenty:      strings.Repeat("😆", l),
				TwentyOne:   strings.Repeat("😀", l),
				TwentyTwo:   strings.Repeat("😆", l),
				TwentyThree: strings.Repeat("😀", l),
			}
			want := []byte{
				byte(len([]byte(o.One))),
				byte(len([]byte(o.Two))),
				byte(len([]byte(o.Three))),
				byte(len([]byte(o.Four))),
				byte(len([]byte(o.Five))),
				byte(len([]byte(o.Six))),
				byte(len([]byte(o.Seven))),
				byte(len([]byte(o.Eight))),
				byte(len([]byte(o.Nine))),
				byte(len([]byte(o.Ten))),
				byte(len([]byte(o.Eleven))),
				byte(len([]byte(o.Twelve))),
				byte(len([]byte(o.Thirteen))),
				byte(len([]byte(o.Fourteen))),
				byte(len([]byte(o.Fifteen))),
				byte(len([]byte(o.Sixteen))),
				byte(len([]byte(o.Seventeen))),
				byte(len([]byte(o.Eighteen))),
				byte(len([]byte(o.Nineteen))),
				byte(len([]byte(o.Twenty))),
				byte(len([]byte(o.TwentyOne))),
				byte(len([]byte(o.TwentyTwo))),
				byte(len([]byte(o.TwentyThree))),
			}
			want = append(want, []byte(o.One)...)
			want = append(want, []byte(o.Two)...)
			want = append(want, []byte(o.Three)...)
			want = append(want, []byte(o.Four)...)
			want = append(want, []byte(o.Five)...)
			want = append(want, []byte(o.Six)...)
			want = append(want, []byte(o.Seven)...)
			want = append(want, []byte(o.Eight)...)
			want = append(want, []byte(o.Nine)...)
			want = append(want, []byte(o.Ten)...)
			want = append(want, []byte(o.Eleven)...)
			want = append(want, []byte(o.Twelve)...)
			want = append(want, []byte(o.Thirteen)...)
			want = append(want, []byte(o.Fourteen)...)
			want = append(want, []byte(o.Fifteen)...)
			want = append(want, []byte(o.Sixteen)...)
			want = append(want, []byte(o.Seventeen)...)
			want = append(want, []byte(o.Eighteen)...)
			want = append(want, []byte(o.Nineteen)...)
			want = append(want, []byte(o.Twenty)...)
			want = append(want, []byte(o.TwentyOne)...)
			want = append(want, []byte(o.TwentyTwo)...)
			want = append(want, []byte(o.TwentyThree)...)
			require.Equal(t, want, o.MarshalJ())
		})
	}
}
