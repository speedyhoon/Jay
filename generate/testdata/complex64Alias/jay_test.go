package main

import (
	"testing"

	"github.com/speedyhoon/rando"
	"github.com/stretchr/testify/require"
)

func TestFuzz_1(t *testing.T) {
	var expected, actual One
	require.NoError(t, expected.UnmarshalJ(actual.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, One{}, expected)
	require.Equal(t, One{}, actual)

	actual = One{
		One: rando.Complex64(),
	}
	src := actual.MarshalJ()
	require.NoError(t, expected.UnmarshalJ(src))
	// require.NotEqual(t, One{}, expected)
	// require.NotEqual(t, One{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_2(t *testing.T) {
	var expected, actual Two
	require.NoError(t, expected.UnmarshalJ(actual.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, Two{}, expected)
	require.Equal(t, Two{}, actual)

	actual = Two{
		One: rando.Complex64(),
		Two: rando.Complex64(),
	}
	src := actual.MarshalJ()
	require.NoError(t, expected.UnmarshalJ(src))
	// require.NotEqual(t, Two{}, expected)
	// require.NotEqual(t, Two{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_3(t *testing.T) {
	var expected, actual Three
	require.NoError(t, expected.UnmarshalJ(actual.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, Three{}, expected)
	require.Equal(t, Three{}, actual)

	actual = Three{
		One:   rando.Complex64(),
		Two:   rando.Complex64(),
		Three: rando.Complex64(),
	}
	src := actual.MarshalJ()
	require.NoError(t, expected.UnmarshalJ(src))
	// require.NotEqual(t, Three{}, expected)
	// require.NotEqual(t, Three{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_4(t *testing.T) {
	var expected, actual Four
	require.NoError(t, expected.UnmarshalJ(actual.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, Four{}, expected)
	require.Equal(t, Four{}, actual)

	actual = Four{
		One:   rando.Complex64(),
		Two:   rando.Complex64(),
		Three: rando.Complex64(),
		Four:  rando.Complex64(),
	}
	src := actual.MarshalJ()
	require.NoError(t, expected.UnmarshalJ(src))
	// require.NotEqual(t, Four{}, expected)
	// require.NotEqual(t, Four{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_5(t *testing.T) {
	var expected, actual Five
	require.NoError(t, expected.UnmarshalJ(actual.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, Five{}, expected)
	require.Equal(t, Five{}, actual)

	actual = Five{
		One:   rando.Complex64(),
		Two:   rando.Complex64(),
		Three: rando.Complex64(),
		Four:  rando.Complex64(),
		Five:  rando.Complex64(),
	}
	src := actual.MarshalJ()
	require.NoError(t, expected.UnmarshalJ(src))
	// require.NotEqual(t, Five{}, expected)
	// require.NotEqual(t, Five{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_6(t *testing.T) {
	var expected, actual Six
	require.NoError(t, expected.UnmarshalJ(actual.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, Six{}, expected)
	require.Equal(t, Six{}, actual)

	actual = Six{
		One:   rando.Complex64(),
		Two:   rando.Complex64(),
		Three: rando.Complex64(),
		Four:  rando.Complex64(),
		Five:  rando.Complex64(),
		Six:   rando.Complex64(),
	}
	src := actual.MarshalJ()
	require.NoError(t, expected.UnmarshalJ(src))
	// require.NotEqual(t, Six{}, expected)
	// require.NotEqual(t, Six{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_7(t *testing.T) {
	var expected, actual Seven
	require.NoError(t, expected.UnmarshalJ(actual.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, Seven{}, expected)
	require.Equal(t, Seven{}, actual)

	actual = Seven{
		One:   rando.Complex64(),
		Two:   rando.Complex64(),
		Three: rando.Complex64(),
		Four:  rando.Complex64(),
		Five:  rando.Complex64(),
		Six:   rando.Complex64(),
		Seven: rando.Complex64(),
	}
	src := actual.MarshalJ()
	require.NoError(t, expected.UnmarshalJ(src))
	// require.NotEqual(t, Seven{}, expected)
	// require.NotEqual(t, Seven{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_8(t *testing.T) {
	var expected, actual Eight
	require.NoError(t, expected.UnmarshalJ(actual.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, Eight{}, expected)
	require.Equal(t, Eight{}, actual)

	actual = Eight{
		One:   rando.Complex64(),
		Two:   rando.Complex64(),
		Three: rando.Complex64(),
		Four:  rando.Complex64(),
		Five:  rando.Complex64(),
		Six:   rando.Complex64(),
		Seven: rando.Complex64(),
		Eight: rando.Complex64(),
	}
	src := actual.MarshalJ()
	require.NoError(t, expected.UnmarshalJ(src))
	// require.NotEqual(t, Eight{}, expected)
	// require.NotEqual(t, Eight{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_9(t *testing.T) {
	var expected, actual Nine
	require.NoError(t, expected.UnmarshalJ(actual.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, Nine{}, expected)
	require.Equal(t, Nine{}, actual)

	actual = Nine{
		One:   rando.Complex64(),
		Two:   rando.Complex64(),
		Three: rando.Complex64(),
		Four:  rando.Complex64(),
		Five:  rando.Complex64(),
		Six:   rando.Complex64(),
		Seven: rando.Complex64(),
		Eight: rando.Complex64(),
		Nine:  rando.Complex64(),
	}
	src := actual.MarshalJ()
	require.NoError(t, expected.UnmarshalJ(src))
	// require.NotEqual(t, Nine{}, expected)
	// require.NotEqual(t, Nine{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_10(t *testing.T) {
	var expected, actual Ten
	require.NoError(t, expected.UnmarshalJ(actual.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, Ten{}, expected)
	require.Equal(t, Ten{}, actual)

	actual = Ten{
		One:   rando.Complex64(),
		Two:   rando.Complex64(),
		Three: rando.Complex64(),
		Four:  rando.Complex64(),
		Five:  rando.Complex64(),
		Six:   rando.Complex64(),
		Seven: rando.Complex64(),
		Eight: rando.Complex64(),
		Nine:  rando.Complex64(),
		Ten:   rando.Complex64(),
	}
	src := actual.MarshalJ()
	require.NoError(t, expected.UnmarshalJ(src))
	// require.NotEqual(t, Ten{}, expected)
	// require.NotEqual(t, Ten{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_11(t *testing.T) {
	var expected, actual Eleven
	require.NoError(t, expected.UnmarshalJ(actual.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, Eleven{}, expected)
	require.Equal(t, Eleven{}, actual)

	actual = Eleven{
		One:    rando.Complex64(),
		Two:    rando.Complex64(),
		Three:  rando.Complex64(),
		Four:   rando.Complex64(),
		Five:   rando.Complex64(),
		Six:    rando.Complex64(),
		Seven:  rando.Complex64(),
		Eight:  rando.Complex64(),
		Nine:   rando.Complex64(),
		Ten:    rando.Complex64(),
		Eleven: rando.Complex64(),
	}
	src := actual.MarshalJ()
	require.NoError(t, expected.UnmarshalJ(src))
	// require.NotEqual(t, Eleven{}, expected)
	// require.NotEqual(t, Eleven{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_12(t *testing.T) {
	var expected, actual Twelve
	require.NoError(t, expected.UnmarshalJ(actual.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, Twelve{}, expected)
	require.Equal(t, Twelve{}, actual)

	actual = Twelve{
		One:    rando.Complex64(),
		Two:    rando.Complex64(),
		Three:  rando.Complex64(),
		Four:   rando.Complex64(),
		Five:   rando.Complex64(),
		Six:    rando.Complex64(),
		Seven:  rando.Complex64(),
		Eight:  rando.Complex64(),
		Nine:   rando.Complex64(),
		Ten:    rando.Complex64(),
		Eleven: rando.Complex64(),
		Twelve: rando.Complex64(),
	}
	src := actual.MarshalJ()
	require.NoError(t, expected.UnmarshalJ(src))
	// require.NotEqual(t, Twelve{}, expected)
	// require.NotEqual(t, Twelve{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_13(t *testing.T) {
	var expected, actual Thirteen
	require.NoError(t, expected.UnmarshalJ(actual.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, Thirteen{}, expected)
	require.Equal(t, Thirteen{}, actual)

	actual = Thirteen{
		One:      rando.Complex64(),
		Two:      rando.Complex64(),
		Three:    rando.Complex64(),
		Four:     rando.Complex64(),
		Five:     rando.Complex64(),
		Six:      rando.Complex64(),
		Seven:    rando.Complex64(),
		Eight:    rando.Complex64(),
		Nine:     rando.Complex64(),
		Ten:      rando.Complex64(),
		Eleven:   rando.Complex64(),
		Twelve:   rando.Complex64(),
		Thirteen: rando.Complex64(),
	}
	src := actual.MarshalJ()
	require.NoError(t, expected.UnmarshalJ(src))
	// require.NotEqual(t, Thirteen{}, expected)
	// require.NotEqual(t, Thirteen{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_14(t *testing.T) {
	var expected, actual Fourteen
	require.NoError(t, expected.UnmarshalJ(actual.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, Fourteen{}, expected)
	require.Equal(t, Fourteen{}, actual)

	actual = Fourteen{
		One:      rando.Complex64(),
		Two:      rando.Complex64(),
		Three:    rando.Complex64(),
		Four:     rando.Complex64(),
		Five:     rando.Complex64(),
		Six:      rando.Complex64(),
		Seven:    rando.Complex64(),
		Eight:    rando.Complex64(),
		Nine:     rando.Complex64(),
		Ten:      rando.Complex64(),
		Eleven:   rando.Complex64(),
		Twelve:   rando.Complex64(),
		Thirteen: rando.Complex64(),
		Fourteen: rando.Complex64(),
	}
	src := actual.MarshalJ()
	require.NoError(t, expected.UnmarshalJ(src))
	// require.NotEqual(t, Fourteen{}, expected)
	// require.NotEqual(t, Fourteen{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_15(t *testing.T) {
	var expected, actual Fifteen
	require.NoError(t, expected.UnmarshalJ(actual.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, Fifteen{}, expected)
	require.Equal(t, Fifteen{}, actual)

	actual = Fifteen{
		One:      rando.Complex64(),
		Two:      rando.Complex64(),
		Three:    rando.Complex64(),
		Four:     rando.Complex64(),
		Five:     rando.Complex64(),
		Six:      rando.Complex64(),
		Seven:    rando.Complex64(),
		Eight:    rando.Complex64(),
		Nine:     rando.Complex64(),
		Ten:      rando.Complex64(),
		Eleven:   rando.Complex64(),
		Twelve:   rando.Complex64(),
		Thirteen: rando.Complex64(),
		Fourteen: rando.Complex64(),
		Fifteen:  rando.Complex64(),
	}
	src := actual.MarshalJ()
	require.NoError(t, expected.UnmarshalJ(src))
	// require.NotEqual(t, Fifteen{}, expected)
	// require.NotEqual(t, Fifteen{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_16(t *testing.T) {
	var expected, actual Sixteen
	require.NoError(t, expected.UnmarshalJ(actual.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, Sixteen{}, expected)
	require.Equal(t, Sixteen{}, actual)

	actual = Sixteen{
		One:      rando.Complex64(),
		Two:      rando.Complex64(),
		Three:    rando.Complex64(),
		Four:     rando.Complex64(),
		Five:     rando.Complex64(),
		Six:      rando.Complex64(),
		Seven:    rando.Complex64(),
		Eight:    rando.Complex64(),
		Nine:     rando.Complex64(),
		Ten:      rando.Complex64(),
		Eleven:   rando.Complex64(),
		Twelve:   rando.Complex64(),
		Thirteen: rando.Complex64(),
		Fourteen: rando.Complex64(),
		Fifteen:  rando.Complex64(),
		Sixteen:  rando.Complex64(),
	}
	src := actual.MarshalJ()
	require.NoError(t, expected.UnmarshalJ(src))
	// require.NotEqual(t, Sixteen{}, expected)
	// require.NotEqual(t, Sixteen{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_17(t *testing.T) {
	var expected, actual Seventeen
	require.NoError(t, expected.UnmarshalJ(actual.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, Seventeen{}, expected)
	require.Equal(t, Seventeen{}, actual)

	actual = Seventeen{
		One:       rando.Complex64(),
		Two:       rando.Complex64(),
		Three:     rando.Complex64(),
		Four:      rando.Complex64(),
		Five:      rando.Complex64(),
		Six:       rando.Complex64(),
		Seven:     rando.Complex64(),
		Eight:     rando.Complex64(),
		Nine:      rando.Complex64(),
		Ten:       rando.Complex64(),
		Eleven:    rando.Complex64(),
		Twelve:    rando.Complex64(),
		Thirteen:  rando.Complex64(),
		Fourteen:  rando.Complex64(),
		Fifteen:   rando.Complex64(),
		Sixteen:   rando.Complex64(),
		Seventeen: rando.Complex64(),
	}
	src := actual.MarshalJ()
	require.NoError(t, expected.UnmarshalJ(src))
	// require.NotEqual(t, Seventeen{}, expected)
	// require.NotEqual(t, Seventeen{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_18(t *testing.T) {
	var expected, actual Eighteen
	require.NoError(t, expected.UnmarshalJ(actual.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, Eighteen{}, expected)
	require.Equal(t, Eighteen{}, actual)

	actual = Eighteen{
		One:       rando.Complex64(),
		Two:       rando.Complex64(),
		Three:     rando.Complex64(),
		Four:      rando.Complex64(),
		Five:      rando.Complex64(),
		Six:       rando.Complex64(),
		Seven:     rando.Complex64(),
		Eight:     rando.Complex64(),
		Nine:      rando.Complex64(),
		Ten:       rando.Complex64(),
		Eleven:    rando.Complex64(),
		Twelve:    rando.Complex64(),
		Thirteen:  rando.Complex64(),
		Fourteen:  rando.Complex64(),
		Fifteen:   rando.Complex64(),
		Sixteen:   rando.Complex64(),
		Seventeen: rando.Complex64(),
		Eighteen:  rando.Complex64(),
	}
	src := actual.MarshalJ()
	require.NoError(t, expected.UnmarshalJ(src))
	// require.NotEqual(t, Eighteen{}, expected)
	// require.NotEqual(t, Eighteen{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_19(t *testing.T) {
	var expected, actual Nineteen
	require.NoError(t, expected.UnmarshalJ(actual.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, Nineteen{}, expected)
	require.Equal(t, Nineteen{}, actual)

	actual = Nineteen{
		One:       rando.Complex64(),
		Two:       rando.Complex64(),
		Three:     rando.Complex64(),
		Four:      rando.Complex64(),
		Five:      rando.Complex64(),
		Six:       rando.Complex64(),
		Seven:     rando.Complex64(),
		Eight:     rando.Complex64(),
		Nine:      rando.Complex64(),
		Ten:       rando.Complex64(),
		Eleven:    rando.Complex64(),
		Twelve:    rando.Complex64(),
		Thirteen:  rando.Complex64(),
		Fourteen:  rando.Complex64(),
		Fifteen:   rando.Complex64(),
		Sixteen:   rando.Complex64(),
		Seventeen: rando.Complex64(),
		Eighteen:  rando.Complex64(),
		Nineteen:  rando.Complex64(),
	}
	src := actual.MarshalJ()
	require.NoError(t, expected.UnmarshalJ(src))
	// require.NotEqual(t, Nineteen{}, expected)
	// require.NotEqual(t, Nineteen{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_20(t *testing.T) {
	var expected, actual Twenty
	require.NoError(t, expected.UnmarshalJ(actual.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, Twenty{}, expected)
	require.Equal(t, Twenty{}, actual)

	actual = Twenty{
		One:       rando.Complex64(),
		Two:       rando.Complex64(),
		Three:     rando.Complex64(),
		Four:      rando.Complex64(),
		Five:      rando.Complex64(),
		Six:       rando.Complex64(),
		Seven:     rando.Complex64(),
		Eight:     rando.Complex64(),
		Nine:      rando.Complex64(),
		Ten:       rando.Complex64(),
		Eleven:    rando.Complex64(),
		Twelve:    rando.Complex64(),
		Thirteen:  rando.Complex64(),
		Fourteen:  rando.Complex64(),
		Fifteen:   rando.Complex64(),
		Sixteen:   rando.Complex64(),
		Seventeen: rando.Complex64(),
		Eighteen:  rando.Complex64(),
		Nineteen:  rando.Complex64(),
		Twenty:    rando.Complex64(),
	}
	src := actual.MarshalJ()
	require.NoError(t, expected.UnmarshalJ(src))
	// require.NotEqual(t, Twenty{}, expected)
	// require.NotEqual(t, Twenty{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_21(t *testing.T) {
	var expected, actual TwentyOne
	require.NoError(t, expected.UnmarshalJ(actual.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, TwentyOne{}, expected)
	require.Equal(t, TwentyOne{}, actual)

	actual = TwentyOne{
		One:       rando.Complex64(),
		Two:       rando.Complex64(),
		Three:     rando.Complex64(),
		Four:      rando.Complex64(),
		Five:      rando.Complex64(),
		Six:       rando.Complex64(),
		Seven:     rando.Complex64(),
		Eight:     rando.Complex64(),
		Nine:      rando.Complex64(),
		Ten:       rando.Complex64(),
		Eleven:    rando.Complex64(),
		Twelve:    rando.Complex64(),
		Thirteen:  rando.Complex64(),
		Fourteen:  rando.Complex64(),
		Fifteen:   rando.Complex64(),
		Sixteen:   rando.Complex64(),
		Seventeen: rando.Complex64(),
		Eighteen:  rando.Complex64(),
		Nineteen:  rando.Complex64(),
		Twenty:    rando.Complex64(),
		TwentyOne: rando.Complex64(),
	}
	src := actual.MarshalJ()
	require.NoError(t, expected.UnmarshalJ(src))
	// require.NotEqual(t, TwentyOne{}, expected)
	// require.NotEqual(t, TwentyOne{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_22(t *testing.T) {
	var expected, actual TwentyTwo
	require.NoError(t, expected.UnmarshalJ(actual.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, TwentyTwo{}, expected)
	require.Equal(t, TwentyTwo{}, actual)

	actual = TwentyTwo{
		One:       rando.Complex64(),
		Two:       rando.Complex64(),
		Three:     rando.Complex64(),
		Four:      rando.Complex64(),
		Five:      rando.Complex64(),
		Six:       rando.Complex64(),
		Seven:     rando.Complex64(),
		Eight:     rando.Complex64(),
		Nine:      rando.Complex64(),
		Ten:       rando.Complex64(),
		Eleven:    rando.Complex64(),
		Twelve:    rando.Complex64(),
		Thirteen:  rando.Complex64(),
		Fourteen:  rando.Complex64(),
		Fifteen:   rando.Complex64(),
		Sixteen:   rando.Complex64(),
		Seventeen: rando.Complex64(),
		Eighteen:  rando.Complex64(),
		Nineteen:  rando.Complex64(),
		Twenty:    rando.Complex64(),
		TwentyOne: rando.Complex64(),
		TwentyTwo: rando.Complex64(),
	}
	src := actual.MarshalJ()
	require.NoError(t, expected.UnmarshalJ(src))
	// require.NotEqual(t, TwentyTwo{}, expected)
	// require.NotEqual(t, TwentyTwo{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_23(t *testing.T) {
	var expected, actual TwentyThree
	require.NoError(t, expected.UnmarshalJ(actual.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, TwentyThree{}, expected)
	require.Equal(t, TwentyThree{}, actual)

	actual = TwentyThree{
		One:         rando.Complex64(),
		Two:         rando.Complex64(),
		Three:       rando.Complex64(),
		Four:        rando.Complex64(),
		Five:        rando.Complex64(),
		Six:         rando.Complex64(),
		Seven:       rando.Complex64(),
		Eight:       rando.Complex64(),
		Nine:        rando.Complex64(),
		Ten:         rando.Complex64(),
		Eleven:      rando.Complex64(),
		Twelve:      rando.Complex64(),
		Thirteen:    rando.Complex64(),
		Fourteen:    rando.Complex64(),
		Fifteen:     rando.Complex64(),
		Sixteen:     rando.Complex64(),
		Seventeen:   rando.Complex64(),
		Eighteen:    rando.Complex64(),
		Nineteen:    rando.Complex64(),
		Twenty:      rando.Complex64(),
		TwentyOne:   rando.Complex64(),
		TwentyTwo:   rando.Complex64(),
		TwentyThree: rando.Complex64(),
	}
	src := actual.MarshalJ()
	require.NoError(t, expected.UnmarshalJ(src))
	// require.NotEqual(t, TwentyThree{}, expected)
	// require.NotEqual(t, TwentyThree{}, actual)
	require.Equal(t, expected, actual)
}
