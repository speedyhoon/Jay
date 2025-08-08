package main

import (
	"testing"

	"github.com/speedyhoon/rando"
	"github.com/stretchr/testify/require"
)

func TestFuzz_1(t *testing.T) {
	var expected, actual One
	require.NoError(t, actual.UnmarshalJ(expected.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, One{}, expected)
	require.Equal(t, One{}, actual)

	expected = One{
		One: rando.Uints(),
	}
	src := expected.MarshalJ()
	require.NoError(t, actual.UnmarshalJ(src))
	// require.NotEqual(t, One{}, expected)
	// require.NotEqual(t, One{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_2(t *testing.T) {
	var expected, actual Two
	require.NoError(t, actual.UnmarshalJ(expected.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, Two{}, expected)
	require.Equal(t, Two{}, actual)

	expected = Two{
		One: rando.Uints(),
		Two: rando.Uints(),
	}
	src := expected.MarshalJ()
	require.NoError(t, actual.UnmarshalJ(src))
	// require.NotEqual(t, Two{}, expected)
	// require.NotEqual(t, Two{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_3(t *testing.T) {
	var expected, actual Three
	require.NoError(t, actual.UnmarshalJ(expected.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, Three{}, expected)
	require.Equal(t, Three{}, actual)

	expected = Three{
		One:   rando.Uints(),
		Two:   rando.Uints(),
		Three: rando.Uints(),
	}
	src := expected.MarshalJ()
	require.NoError(t, actual.UnmarshalJ(src))
	// require.NotEqual(t, Three{}, expected)
	// require.NotEqual(t, Three{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_4(t *testing.T) {
	var expected, actual Four
	require.NoError(t, actual.UnmarshalJ(expected.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, Four{}, expected)
	require.Equal(t, Four{}, actual)

	expected = Four{
		One:   rando.Uints(),
		Two:   rando.Uints(),
		Three: rando.Uints(),
		Four:  rando.Uints(),
	}
	src := expected.MarshalJ()
	require.NoError(t, actual.UnmarshalJ(src))
	// require.NotEqual(t, Four{}, expected)
	// require.NotEqual(t, Four{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_5(t *testing.T) {
	var expected, actual Five
	require.NoError(t, actual.UnmarshalJ(expected.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, Five{}, expected)
	require.Equal(t, Five{}, actual)

	expected = Five{
		One:   rando.Uints(),
		Two:   rando.Uints(),
		Three: rando.Uints(),
		Four:  rando.Uints(),
		Five:  rando.Uints(),
	}
	src := expected.MarshalJ()
	require.NoError(t, actual.UnmarshalJ(src))
	// require.NotEqual(t, Five{}, expected)
	// require.NotEqual(t, Five{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_6(t *testing.T) {
	var expected, actual Six
	require.NoError(t, actual.UnmarshalJ(expected.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, Six{}, expected)
	require.Equal(t, Six{}, actual)

	expected = Six{
		One:   rando.Uints(),
		Two:   rando.Uints(),
		Three: rando.Uints(),
		Four:  rando.Uints(),
		Five:  rando.Uints(),
		Six:   rando.Uints(),
	}
	src := expected.MarshalJ()
	require.NoError(t, actual.UnmarshalJ(src))
	// require.NotEqual(t, Six{}, expected)
	// require.NotEqual(t, Six{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_7(t *testing.T) {
	var expected, actual Seven
	require.NoError(t, actual.UnmarshalJ(expected.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, Seven{}, expected)
	require.Equal(t, Seven{}, actual)

	expected = Seven{
		One:   rando.Uints(),
		Two:   rando.Uints(),
		Three: rando.Uints(),
		Four:  rando.Uints(),
		Five:  rando.Uints(),
		Six:   rando.Uints(),
		Seven: rando.Uints(),
	}
	src := expected.MarshalJ()
	require.NoError(t, actual.UnmarshalJ(src))
	// require.NotEqual(t, Seven{}, expected)
	// require.NotEqual(t, Seven{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_8(t *testing.T) {
	var expected, actual Eight
	require.NoError(t, actual.UnmarshalJ(expected.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, Eight{}, expected)
	require.Equal(t, Eight{}, actual)

	expected = Eight{
		One:   rando.Uints(),
		Two:   rando.Uints(),
		Three: rando.Uints(),
		Four:  rando.Uints(),
		Five:  rando.Uints(),
		Six:   rando.Uints(),
		Seven: rando.Uints(),
		Eight: rando.Uints(),
	}
	src := expected.MarshalJ()
	require.NoError(t, actual.UnmarshalJ(src))
	// require.NotEqual(t, Eight{}, expected)
	// require.NotEqual(t, Eight{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_9(t *testing.T) {
	var expected, actual Nine
	require.NoError(t, actual.UnmarshalJ(expected.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, Nine{}, expected)
	require.Equal(t, Nine{}, actual)

	expected = Nine{
		One:   rando.Uints(),
		Two:   rando.Uints(),
		Three: rando.Uints(),
		Four:  rando.Uints(),
		Five:  rando.Uints(),
		Six:   rando.Uints(),
		Seven: rando.Uints(),
		Eight: rando.Uints(),
		Nine:  rando.Uints(),
	}
	src := expected.MarshalJ()
	require.NoError(t, actual.UnmarshalJ(src))
	// require.NotEqual(t, Nine{}, expected)
	// require.NotEqual(t, Nine{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_10(t *testing.T) {
	var expected, actual Ten
	require.NoError(t, actual.UnmarshalJ(expected.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, Ten{}, expected)
	require.Equal(t, Ten{}, actual)

	expected = Ten{
		One:   rando.Uints(),
		Two:   rando.Uints(),
		Three: rando.Uints(),
		Four:  rando.Uints(),
		Five:  rando.Uints(),
		Six:   rando.Uints(),
		Seven: rando.Uints(),
		Eight: rando.Uints(),
		Nine:  rando.Uints(),
		Ten:   rando.Uints(),
	}
	src := expected.MarshalJ()
	require.NoError(t, actual.UnmarshalJ(src))
	// require.NotEqual(t, Ten{}, expected)
	// require.NotEqual(t, Ten{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_11(t *testing.T) {
	var expected, actual Eleven
	require.NoError(t, actual.UnmarshalJ(expected.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, Eleven{}, expected)
	require.Equal(t, Eleven{}, actual)

	expected = Eleven{
		One:    rando.Uints(),
		Two:    rando.Uints(),
		Three:  rando.Uints(),
		Four:   rando.Uints(),
		Five:   rando.Uints(),
		Six:    rando.Uints(),
		Seven:  rando.Uints(),
		Eight:  rando.Uints(),
		Nine:   rando.Uints(),
		Ten:    rando.Uints(),
		Eleven: rando.Uints(),
	}
	src := expected.MarshalJ()
	require.NoError(t, actual.UnmarshalJ(src))
	// require.NotEqual(t, Eleven{}, expected)
	// require.NotEqual(t, Eleven{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_12(t *testing.T) {
	var expected, actual Twelve
	require.NoError(t, actual.UnmarshalJ(expected.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, Twelve{}, expected)
	require.Equal(t, Twelve{}, actual)

	expected = Twelve{
		One:    rando.Uints(),
		Two:    rando.Uints(),
		Three:  rando.Uints(),
		Four:   rando.Uints(),
		Five:   rando.Uints(),
		Six:    rando.Uints(),
		Seven:  rando.Uints(),
		Eight:  rando.Uints(),
		Nine:   rando.Uints(),
		Ten:    rando.Uints(),
		Eleven: rando.Uints(),
		Twelve: rando.Uints(),
	}
	src := expected.MarshalJ()
	require.NoError(t, actual.UnmarshalJ(src))
	// require.NotEqual(t, Twelve{}, expected)
	// require.NotEqual(t, Twelve{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_13(t *testing.T) {
	var expected, actual Thirteen
	require.NoError(t, actual.UnmarshalJ(expected.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, Thirteen{}, expected)
	require.Equal(t, Thirteen{}, actual)

	expected = Thirteen{
		One:      rando.Uints(),
		Two:      rando.Uints(),
		Three:    rando.Uints(),
		Four:     rando.Uints(),
		Five:     rando.Uints(),
		Six:      rando.Uints(),
		Seven:    rando.Uints(),
		Eight:    rando.Uints(),
		Nine:     rando.Uints(),
		Ten:      rando.Uints(),
		Eleven:   rando.Uints(),
		Twelve:   rando.Uints(),
		Thirteen: rando.Uints(),
	}
	src := expected.MarshalJ()
	require.NoError(t, actual.UnmarshalJ(src))
	// require.NotEqual(t, Thirteen{}, expected)
	// require.NotEqual(t, Thirteen{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_14(t *testing.T) {
	var expected, actual Fourteen
	require.NoError(t, actual.UnmarshalJ(expected.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, Fourteen{}, expected)
	require.Equal(t, Fourteen{}, actual)

	expected = Fourteen{
		One:      rando.Uints(),
		Two:      rando.Uints(),
		Three:    rando.Uints(),
		Four:     rando.Uints(),
		Five:     rando.Uints(),
		Six:      rando.Uints(),
		Seven:    rando.Uints(),
		Eight:    rando.Uints(),
		Nine:     rando.Uints(),
		Ten:      rando.Uints(),
		Eleven:   rando.Uints(),
		Twelve:   rando.Uints(),
		Thirteen: rando.Uints(),
		Fourteen: rando.Uints(),
	}
	src := expected.MarshalJ()
	require.NoError(t, actual.UnmarshalJ(src))
	// require.NotEqual(t, Fourteen{}, expected)
	// require.NotEqual(t, Fourteen{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_15(t *testing.T) {
	var expected, actual Fifteen
	require.NoError(t, actual.UnmarshalJ(expected.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, Fifteen{}, expected)
	require.Equal(t, Fifteen{}, actual)

	expected = Fifteen{
		One:      rando.Uints(),
		Two:      rando.Uints(),
		Three:    rando.Uints(),
		Four:     rando.Uints(),
		Five:     rando.Uints(),
		Six:      rando.Uints(),
		Seven:    rando.Uints(),
		Eight:    rando.Uints(),
		Nine:     rando.Uints(),
		Ten:      rando.Uints(),
		Eleven:   rando.Uints(),
		Twelve:   rando.Uints(),
		Thirteen: rando.Uints(),
		Fourteen: rando.Uints(),
		Fifteen:  rando.Uints(),
	}
	src := expected.MarshalJ()
	require.NoError(t, actual.UnmarshalJ(src))
	// require.NotEqual(t, Fifteen{}, expected)
	// require.NotEqual(t, Fifteen{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_16(t *testing.T) {
	var expected, actual Sixteen
	require.NoError(t, actual.UnmarshalJ(expected.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, Sixteen{}, expected)
	require.Equal(t, Sixteen{}, actual)

	expected = Sixteen{
		One:      rando.Uints(),
		Two:      rando.Uints(),
		Three:    rando.Uints(),
		Four:     rando.Uints(),
		Five:     rando.Uints(),
		Six:      rando.Uints(),
		Seven:    rando.Uints(),
		Eight:    rando.Uints(),
		Nine:     rando.Uints(),
		Ten:      rando.Uints(),
		Eleven:   rando.Uints(),
		Twelve:   rando.Uints(),
		Thirteen: rando.Uints(),
		Fourteen: rando.Uints(),
		Fifteen:  rando.Uints(),
		Sixteen:  rando.Uints(),
	}
	src := expected.MarshalJ()
	require.NoError(t, actual.UnmarshalJ(src))
	// require.NotEqual(t, Sixteen{}, expected)
	// require.NotEqual(t, Sixteen{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_17(t *testing.T) {
	var expected, actual Seventeen
	require.NoError(t, actual.UnmarshalJ(expected.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, Seventeen{}, expected)
	require.Equal(t, Seventeen{}, actual)

	expected = Seventeen{
		One:       rando.Uints(),
		Two:       rando.Uints(),
		Three:     rando.Uints(),
		Four:      rando.Uints(),
		Five:      rando.Uints(),
		Six:       rando.Uints(),
		Seven:     rando.Uints(),
		Eight:     rando.Uints(),
		Nine:      rando.Uints(),
		Ten:       rando.Uints(),
		Eleven:    rando.Uints(),
		Twelve:    rando.Uints(),
		Thirteen:  rando.Uints(),
		Fourteen:  rando.Uints(),
		Fifteen:   rando.Uints(),
		Sixteen:   rando.Uints(),
		Seventeen: rando.Uints(),
	}
	src := expected.MarshalJ()
	require.NoError(t, actual.UnmarshalJ(src))
	// require.NotEqual(t, Seventeen{}, expected)
	// require.NotEqual(t, Seventeen{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_18(t *testing.T) {
	var expected, actual Eighteen
	require.NoError(t, actual.UnmarshalJ(expected.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, Eighteen{}, expected)
	require.Equal(t, Eighteen{}, actual)

	expected = Eighteen{
		One:       rando.Uints(),
		Two:       rando.Uints(),
		Three:     rando.Uints(),
		Four:      rando.Uints(),
		Five:      rando.Uints(),
		Six:       rando.Uints(),
		Seven:     rando.Uints(),
		Eight:     rando.Uints(),
		Nine:      rando.Uints(),
		Ten:       rando.Uints(),
		Eleven:    rando.Uints(),
		Twelve:    rando.Uints(),
		Thirteen:  rando.Uints(),
		Fourteen:  rando.Uints(),
		Fifteen:   rando.Uints(),
		Sixteen:   rando.Uints(),
		Seventeen: rando.Uints(),
		Eighteen:  rando.Uints(),
	}
	src := expected.MarshalJ()
	require.NoError(t, actual.UnmarshalJ(src))
	// require.NotEqual(t, Eighteen{}, expected)
	// require.NotEqual(t, Eighteen{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_19(t *testing.T) {
	var expected, actual Nineteen
	require.NoError(t, actual.UnmarshalJ(expected.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, Nineteen{}, expected)
	require.Equal(t, Nineteen{}, actual)

	expected = Nineteen{
		One:       rando.Uints(),
		Two:       rando.Uints(),
		Three:     rando.Uints(),
		Four:      rando.Uints(),
		Five:      rando.Uints(),
		Six:       rando.Uints(),
		Seven:     rando.Uints(),
		Eight:     rando.Uints(),
		Nine:      rando.Uints(),
		Ten:       rando.Uints(),
		Eleven:    rando.Uints(),
		Twelve:    rando.Uints(),
		Thirteen:  rando.Uints(),
		Fourteen:  rando.Uints(),
		Fifteen:   rando.Uints(),
		Sixteen:   rando.Uints(),
		Seventeen: rando.Uints(),
		Eighteen:  rando.Uints(),
		Nineteen:  rando.Uints(),
	}
	src := expected.MarshalJ()
	require.NoError(t, actual.UnmarshalJ(src))
	// require.NotEqual(t, Nineteen{}, expected)
	// require.NotEqual(t, Nineteen{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_20(t *testing.T) {
	var expected, actual Twenty
	require.NoError(t, actual.UnmarshalJ(expected.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, Twenty{}, expected)
	require.Equal(t, Twenty{}, actual)

	expected = Twenty{
		One:       rando.Uints(),
		Two:       rando.Uints(),
		Three:     rando.Uints(),
		Four:      rando.Uints(),
		Five:      rando.Uints(),
		Six:       rando.Uints(),
		Seven:     rando.Uints(),
		Eight:     rando.Uints(),
		Nine:      rando.Uints(),
		Ten:       rando.Uints(),
		Eleven:    rando.Uints(),
		Twelve:    rando.Uints(),
		Thirteen:  rando.Uints(),
		Fourteen:  rando.Uints(),
		Fifteen:   rando.Uints(),
		Sixteen:   rando.Uints(),
		Seventeen: rando.Uints(),
		Eighteen:  rando.Uints(),
		Nineteen:  rando.Uints(),
		Twenty:    rando.Uints(),
	}
	src := expected.MarshalJ()
	require.NoError(t, actual.UnmarshalJ(src))
	// require.NotEqual(t, Twenty{}, expected)
	// require.NotEqual(t, Twenty{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_21(t *testing.T) {
	var expected, actual TwentyOne
	require.NoError(t, actual.UnmarshalJ(expected.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, TwentyOne{}, expected)
	require.Equal(t, TwentyOne{}, actual)

	expected = TwentyOne{
		One:       rando.Uints(),
		Two:       rando.Uints(),
		Three:     rando.Uints(),
		Four:      rando.Uints(),
		Five:      rando.Uints(),
		Six:       rando.Uints(),
		Seven:     rando.Uints(),
		Eight:     rando.Uints(),
		Nine:      rando.Uints(),
		Ten:       rando.Uints(),
		Eleven:    rando.Uints(),
		Twelve:    rando.Uints(),
		Thirteen:  rando.Uints(),
		Fourteen:  rando.Uints(),
		Fifteen:   rando.Uints(),
		Sixteen:   rando.Uints(),
		Seventeen: rando.Uints(),
		Eighteen:  rando.Uints(),
		Nineteen:  rando.Uints(),
		Twenty:    rando.Uints(),
		TwentyOne: rando.Uints(),
	}
	src := expected.MarshalJ()
	require.NoError(t, actual.UnmarshalJ(src))
	// require.NotEqual(t, TwentyOne{}, expected)
	// require.NotEqual(t, TwentyOne{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_22(t *testing.T) {
	var expected, actual TwentyTwo
	require.NoError(t, actual.UnmarshalJ(expected.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, TwentyTwo{}, expected)
	require.Equal(t, TwentyTwo{}, actual)

	expected = TwentyTwo{
		One:       rando.Uints(),
		Two:       rando.Uints(),
		Three:     rando.Uints(),
		Four:      rando.Uints(),
		Five:      rando.Uints(),
		Six:       rando.Uints(),
		Seven:     rando.Uints(),
		Eight:     rando.Uints(),
		Nine:      rando.Uints(),
		Ten:       rando.Uints(),
		Eleven:    rando.Uints(),
		Twelve:    rando.Uints(),
		Thirteen:  rando.Uints(),
		Fourteen:  rando.Uints(),
		Fifteen:   rando.Uints(),
		Sixteen:   rando.Uints(),
		Seventeen: rando.Uints(),
		Eighteen:  rando.Uints(),
		Nineteen:  rando.Uints(),
		Twenty:    rando.Uints(),
		TwentyOne: rando.Uints(),
		TwentyTwo: rando.Uints(),
	}
	src := expected.MarshalJ()
	require.NoError(t, actual.UnmarshalJ(src))
	// require.NotEqual(t, TwentyTwo{}, expected)
	// require.NotEqual(t, TwentyTwo{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_23(t *testing.T) {
	var expected, actual TwentyThree
	require.NoError(t, actual.UnmarshalJ(expected.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, TwentyThree{}, expected)
	require.Equal(t, TwentyThree{}, actual)

	expected = TwentyThree{
		One:         rando.Uints(),
		Two:         rando.Uints(),
		Three:       rando.Uints(),
		Four:        rando.Uints(),
		Five:        rando.Uints(),
		Six:         rando.Uints(),
		Seven:       rando.Uints(),
		Eight:       rando.Uints(),
		Nine:        rando.Uints(),
		Ten:         rando.Uints(),
		Eleven:      rando.Uints(),
		Twelve:      rando.Uints(),
		Thirteen:    rando.Uints(),
		Fourteen:    rando.Uints(),
		Fifteen:     rando.Uints(),
		Sixteen:     rando.Uints(),
		Seventeen:   rando.Uints(),
		Eighteen:    rando.Uints(),
		Nineteen:    rando.Uints(),
		Twenty:      rando.Uints(),
		TwentyOne:   rando.Uints(),
		TwentyTwo:   rando.Uints(),
		TwentyThree: rando.Uints(),
	}
	src := expected.MarshalJ()
	require.NoError(t, actual.UnmarshalJ(src))
	// require.NotEqual(t, TwentyThree{}, expected)
	// require.NotEqual(t, TwentyThree{}, actual)
	require.Equal(t, expected, actual)
}
