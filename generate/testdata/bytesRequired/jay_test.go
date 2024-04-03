package main

import (
	"testing"

	"github.com/speedyhoon/rando"
	"github.com/stretchr/testify/require"
)

func TestFuzz_1(t *testing.T) {
	var expected, actual One
	require.NoError(t, expected.UnmarshalJ(actual.MarshalJ()))
	require.NotEqual(t, expected, actual)
	require.Equal(t, One{One: []byte{}}, expected)
	require.Equal(t, One{}, actual)

	actual = One{
		One: rando.Bytes(),
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
	require.NotEqual(t, expected, actual)
	require.Equal(t, Two{One: []byte{}, Two: []byte{}}, expected)
	require.Equal(t, Two{}, actual)

	actual = Two{
		One: rando.Bytes(),
		Two: rando.Bytes(),
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
	require.NotEqual(t, expected, actual)
	require.Equal(t, Three{One: []byte{}, Two: []byte{}, Three: []byte{}}, expected)
	require.Equal(t, Three{}, actual)

	actual = Three{
		One:   rando.Bytes(),
		Two:   rando.Bytes(),
		Three: rando.Bytes(),
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
	require.NotEqual(t, expected, actual)
	require.Equal(t, Four{One: []byte{}, Two: []byte{}, Three: []byte{}, Four: []byte{}}, expected)
	require.Equal(t, Four{}, actual)

	actual = Four{
		One:   rando.Bytes(),
		Two:   rando.Bytes(),
		Three: rando.Bytes(),
		Four:  rando.Bytes(),
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
	require.NotEqual(t, expected, actual)
	require.Equal(t, Five{One: []byte{}, Two: []byte{}, Three: []byte{}, Four: []byte{}, Five: []byte{}}, expected)
	require.Equal(t, Five{}, actual)

	actual = Five{
		One:   rando.Bytes(),
		Two:   rando.Bytes(),
		Three: rando.Bytes(),
		Four:  rando.Bytes(),
		Five:  rando.Bytes(),
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
	require.NotEqual(t, expected, actual)
	require.Equal(t, Six{One: []byte{}, Two: []byte{}, Three: []byte{}, Four: []byte{}, Five: []byte{}, Six: []byte{}}, expected)
	require.Equal(t, Six{}, actual)

	actual = Six{
		One:   rando.Bytes(),
		Two:   rando.Bytes(),
		Three: rando.Bytes(),
		Four:  rando.Bytes(),
		Five:  rando.Bytes(),
		Six:   rando.Bytes(),
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
	require.NotEqual(t, expected, actual)
	require.Equal(t, Seven{One: []byte{}, Two: []byte{}, Three: []byte{}, Four: []byte{}, Five: []byte{}, Six: []byte{}, Seven: []byte{}}, expected)
	require.Equal(t, Seven{}, actual)

	actual = Seven{
		One:   rando.Bytes(),
		Two:   rando.Bytes(),
		Three: rando.Bytes(),
		Four:  rando.Bytes(),
		Five:  rando.Bytes(),
		Six:   rando.Bytes(),
		Seven: rando.Bytes(),
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
	require.NotEqual(t, expected, actual)
	require.Equal(t, Eight{One: []byte{}, Two: []byte{}, Three: []byte{}, Four: []byte{}, Five: []byte{}, Six: []byte{}, Seven: []byte{}, Eight: []byte{}}, expected)
	require.Equal(t, Eight{}, actual)

	actual = Eight{
		One:   rando.Bytes(),
		Two:   rando.Bytes(),
		Three: rando.Bytes(),
		Four:  rando.Bytes(),
		Five:  rando.Bytes(),
		Six:   rando.Bytes(),
		Seven: rando.Bytes(),
		Eight: rando.Bytes(),
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
	require.NotEqual(t, expected, actual)
	require.Equal(t, Nine{One: []byte{}, Two: []byte{}, Three: []byte{}, Four: []byte{}, Five: []byte{}, Six: []byte{}, Seven: []byte{}, Eight: []byte{}, Nine: []byte{}}, expected)
	require.Equal(t, Nine{}, actual)

	actual = Nine{
		One:   rando.Bytes(),
		Two:   rando.Bytes(),
		Three: rando.Bytes(),
		Four:  rando.Bytes(),
		Five:  rando.Bytes(),
		Six:   rando.Bytes(),
		Seven: rando.Bytes(),
		Eight: rando.Bytes(),
		Nine:  rando.Bytes(),
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
	require.NotEqual(t, expected, actual)
	require.Equal(t, Ten{One: []byte{}, Two: []byte{}, Three: []byte{}, Four: []byte{}, Five: []byte{}, Six: []byte{}, Seven: []byte{}, Eight: []byte{}, Nine: []byte{}, Ten: []byte{}}, expected)
	require.Equal(t, Ten{}, actual)

	actual = Ten{
		One:   rando.Bytes(),
		Two:   rando.Bytes(),
		Three: rando.Bytes(),
		Four:  rando.Bytes(),
		Five:  rando.Bytes(),
		Six:   rando.Bytes(),
		Seven: rando.Bytes(),
		Eight: rando.Bytes(),
		Nine:  rando.Bytes(),
		Ten:   rando.Bytes(),
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
	require.NotEqual(t, expected, actual)
	require.Equal(t, Eleven{One: []byte{}, Two: []byte{}, Three: []byte{}, Four: []byte{}, Five: []byte{}, Six: []byte{}, Seven: []byte{}, Eight: []byte{}, Nine: []byte{}, Ten: []byte{}, Eleven: []byte{}}, expected)
	require.Equal(t, Eleven{}, actual)

	actual = Eleven{
		One:    rando.Bytes(),
		Two:    rando.Bytes(),
		Three:  rando.Bytes(),
		Four:   rando.Bytes(),
		Five:   rando.Bytes(),
		Six:    rando.Bytes(),
		Seven:  rando.Bytes(),
		Eight:  rando.Bytes(),
		Nine:   rando.Bytes(),
		Ten:    rando.Bytes(),
		Eleven: rando.Bytes(),
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
	require.NotEqual(t, expected, actual)
	require.Equal(t, Twelve{One: []byte{}, Two: []byte{}, Three: []byte{}, Four: []byte{}, Five: []byte{}, Six: []byte{}, Seven: []byte{}, Eight: []byte{}, Nine: []byte{}, Ten: []byte{}, Eleven: []byte{}, Twelve: []byte{}}, expected)
	require.Equal(t, Twelve{}, actual)

	actual = Twelve{
		One:    rando.Bytes(),
		Two:    rando.Bytes(),
		Three:  rando.Bytes(),
		Four:   rando.Bytes(),
		Five:   rando.Bytes(),
		Six:    rando.Bytes(),
		Seven:  rando.Bytes(),
		Eight:  rando.Bytes(),
		Nine:   rando.Bytes(),
		Ten:    rando.Bytes(),
		Eleven: rando.Bytes(),
		Twelve: rando.Bytes(),
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
	require.NotEqual(t, expected, actual)
	require.Equal(t, Thirteen{One: []byte{}, Two: []byte{}, Three: []byte{}, Four: []byte{}, Five: []byte{}, Six: []byte{}, Seven: []byte{}, Eight: []byte{}, Nine: []byte{}, Ten: []byte{}, Eleven: []byte{}, Twelve: []byte{}, Thirteen: []byte{}}, expected)
	require.Equal(t, Thirteen{}, actual)

	actual = Thirteen{
		One:      rando.Bytes(),
		Two:      rando.Bytes(),
		Three:    rando.Bytes(),
		Four:     rando.Bytes(),
		Five:     rando.Bytes(),
		Six:      rando.Bytes(),
		Seven:    rando.Bytes(),
		Eight:    rando.Bytes(),
		Nine:     rando.Bytes(),
		Ten:      rando.Bytes(),
		Eleven:   rando.Bytes(),
		Twelve:   rando.Bytes(),
		Thirteen: rando.Bytes(),
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
	require.NotEqual(t, expected, actual)
	require.Equal(t, Fourteen{One: []byte{}, Two: []byte{}, Three: []byte{}, Four: []byte{}, Five: []byte{}, Six: []byte{}, Seven: []byte{}, Eight: []byte{}, Nine: []byte{}, Ten: []byte{}, Eleven: []byte{}, Twelve: []byte{}, Thirteen: []byte{}, Fourteen: []byte{}}, expected)
	require.Equal(t, Fourteen{}, actual)

	actual = Fourteen{
		One:      rando.Bytes(),
		Two:      rando.Bytes(),
		Three:    rando.Bytes(),
		Four:     rando.Bytes(),
		Five:     rando.Bytes(),
		Six:      rando.Bytes(),
		Seven:    rando.Bytes(),
		Eight:    rando.Bytes(),
		Nine:     rando.Bytes(),
		Ten:      rando.Bytes(),
		Eleven:   rando.Bytes(),
		Twelve:   rando.Bytes(),
		Thirteen: rando.Bytes(),
		Fourteen: rando.Bytes(),
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
	require.NotEqual(t, expected, actual)
	require.Equal(t, Fifteen{One: []byte{}, Two: []byte{}, Three: []byte{}, Four: []byte{}, Five: []byte{}, Six: []byte{}, Seven: []byte{}, Eight: []byte{}, Nine: []byte{}, Ten: []byte{}, Eleven: []byte{}, Twelve: []byte{}, Thirteen: []byte{}, Fourteen: []byte{}, Fifteen: []byte{}}, expected)
	require.Equal(t, Fifteen{}, actual)

	actual = Fifteen{
		One:      rando.Bytes(),
		Two:      rando.Bytes(),
		Three:    rando.Bytes(),
		Four:     rando.Bytes(),
		Five:     rando.Bytes(),
		Six:      rando.Bytes(),
		Seven:    rando.Bytes(),
		Eight:    rando.Bytes(),
		Nine:     rando.Bytes(),
		Ten:      rando.Bytes(),
		Eleven:   rando.Bytes(),
		Twelve:   rando.Bytes(),
		Thirteen: rando.Bytes(),
		Fourteen: rando.Bytes(),
		Fifteen:  rando.Bytes(),
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
	require.NotEqual(t, expected, actual)
	require.Equal(t, Sixteen{One: []byte{}, Two: []byte{}, Three: []byte{}, Four: []byte{}, Five: []byte{}, Six: []byte{}, Seven: []byte{}, Eight: []byte{}, Nine: []byte{}, Ten: []byte{}, Eleven: []byte{}, Twelve: []byte{}, Thirteen: []byte{}, Fourteen: []byte{}, Fifteen: []byte{}, Sixteen: []byte{}}, expected)
	require.Equal(t, Sixteen{}, actual)

	actual = Sixteen{
		One:      rando.Bytes(),
		Two:      rando.Bytes(),
		Three:    rando.Bytes(),
		Four:     rando.Bytes(),
		Five:     rando.Bytes(),
		Six:      rando.Bytes(),
		Seven:    rando.Bytes(),
		Eight:    rando.Bytes(),
		Nine:     rando.Bytes(),
		Ten:      rando.Bytes(),
		Eleven:   rando.Bytes(),
		Twelve:   rando.Bytes(),
		Thirteen: rando.Bytes(),
		Fourteen: rando.Bytes(),
		Fifteen:  rando.Bytes(),
		Sixteen:  rando.Bytes(),
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
	require.NotEqual(t, expected, actual)
	require.Equal(t, Seventeen{One: []byte{}, Two: []byte{}, Three: []byte{}, Four: []byte{}, Five: []byte{}, Six: []byte{}, Seven: []byte{}, Eight: []byte{}, Nine: []byte{}, Ten: []byte{}, Eleven: []byte{}, Twelve: []byte{}, Thirteen: []byte{}, Fourteen: []byte{}, Fifteen: []byte{}, Sixteen: []byte{}, Seventeen: []byte{}}, expected)
	require.Equal(t, Seventeen{}, actual)

	actual = Seventeen{
		One:       rando.Bytes(),
		Two:       rando.Bytes(),
		Three:     rando.Bytes(),
		Four:      rando.Bytes(),
		Five:      rando.Bytes(),
		Six:       rando.Bytes(),
		Seven:     rando.Bytes(),
		Eight:     rando.Bytes(),
		Nine:      rando.Bytes(),
		Ten:       rando.Bytes(),
		Eleven:    rando.Bytes(),
		Twelve:    rando.Bytes(),
		Thirteen:  rando.Bytes(),
		Fourteen:  rando.Bytes(),
		Fifteen:   rando.Bytes(),
		Sixteen:   rando.Bytes(),
		Seventeen: rando.Bytes(),
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
	require.NotEqual(t, expected, actual)
	require.Equal(t, Eighteen{One: []byte{}, Two: []byte{}, Three: []byte{}, Four: []byte{}, Five: []byte{}, Six: []byte{}, Seven: []byte{}, Eight: []byte{}, Nine: []byte{}, Ten: []byte{}, Eleven: []byte{}, Twelve: []byte{}, Thirteen: []byte{}, Fourteen: []byte{}, Fifteen: []byte{}, Sixteen: []byte{}, Seventeen: []byte{}, Eighteen: []byte{}}, expected)
	require.Equal(t, Eighteen{}, actual)

	actual = Eighteen{
		One:       rando.Bytes(),
		Two:       rando.Bytes(),
		Three:     rando.Bytes(),
		Four:      rando.Bytes(),
		Five:      rando.Bytes(),
		Six:       rando.Bytes(),
		Seven:     rando.Bytes(),
		Eight:     rando.Bytes(),
		Nine:      rando.Bytes(),
		Ten:       rando.Bytes(),
		Eleven:    rando.Bytes(),
		Twelve:    rando.Bytes(),
		Thirteen:  rando.Bytes(),
		Fourteen:  rando.Bytes(),
		Fifteen:   rando.Bytes(),
		Sixteen:   rando.Bytes(),
		Seventeen: rando.Bytes(),
		Eighteen:  rando.Bytes(),
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
	require.NotEqual(t, expected, actual)
	require.Equal(t, Nineteen{One: []byte{}, Two: []byte{}, Three: []byte{}, Four: []byte{}, Five: []byte{}, Six: []byte{}, Seven: []byte{}, Eight: []byte{}, Nine: []byte{}, Ten: []byte{}, Eleven: []byte{}, Twelve: []byte{}, Thirteen: []byte{}, Fourteen: []byte{}, Fifteen: []byte{}, Sixteen: []byte{}, Seventeen: []byte{}, Eighteen: []byte{}, Nineteen: []byte{}}, expected)
	require.Equal(t, Nineteen{}, actual)

	actual = Nineteen{
		One:       rando.Bytes(),
		Two:       rando.Bytes(),
		Three:     rando.Bytes(),
		Four:      rando.Bytes(),
		Five:      rando.Bytes(),
		Six:       rando.Bytes(),
		Seven:     rando.Bytes(),
		Eight:     rando.Bytes(),
		Nine:      rando.Bytes(),
		Ten:       rando.Bytes(),
		Eleven:    rando.Bytes(),
		Twelve:    rando.Bytes(),
		Thirteen:  rando.Bytes(),
		Fourteen:  rando.Bytes(),
		Fifteen:   rando.Bytes(),
		Sixteen:   rando.Bytes(),
		Seventeen: rando.Bytes(),
		Eighteen:  rando.Bytes(),
		Nineteen:  rando.Bytes(),
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
	require.NotEqual(t, expected, actual)
	require.Equal(t, Twenty{One: []byte{}, Two: []byte{}, Three: []byte{}, Four: []byte{}, Five: []byte{}, Six: []byte{}, Seven: []byte{}, Eight: []byte{}, Nine: []byte{}, Ten: []byte{}, Eleven: []byte{}, Twelve: []byte{}, Thirteen: []byte{}, Fourteen: []byte{}, Fifteen: []byte{}, Sixteen: []byte{}, Seventeen: []byte{}, Eighteen: []byte{}, Nineteen: []byte{}, Twenty: []byte{}}, expected)
	require.Equal(t, Twenty{}, actual)

	actual = Twenty{
		One:       rando.Bytes(),
		Two:       rando.Bytes(),
		Three:     rando.Bytes(),
		Four:      rando.Bytes(),
		Five:      rando.Bytes(),
		Six:       rando.Bytes(),
		Seven:     rando.Bytes(),
		Eight:     rando.Bytes(),
		Nine:      rando.Bytes(),
		Ten:       rando.Bytes(),
		Eleven:    rando.Bytes(),
		Twelve:    rando.Bytes(),
		Thirteen:  rando.Bytes(),
		Fourteen:  rando.Bytes(),
		Fifteen:   rando.Bytes(),
		Sixteen:   rando.Bytes(),
		Seventeen: rando.Bytes(),
		Eighteen:  rando.Bytes(),
		Nineteen:  rando.Bytes(),
		Twenty:    rando.Bytes(),
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
	require.NotEqual(t, expected, actual)
	require.Equal(t, TwentyOne{One: []byte{}, Two: []byte{}, Three: []byte{}, Four: []byte{}, Five: []byte{}, Six: []byte{}, Seven: []byte{}, Eight: []byte{}, Nine: []byte{}, Ten: []byte{}, Eleven: []byte{}, Twelve: []byte{}, Thirteen: []byte{}, Fourteen: []byte{}, Fifteen: []byte{}, Sixteen: []byte{}, Seventeen: []byte{}, Eighteen: []byte{}, Nineteen: []byte{}, Twenty: []byte{}, TwentyOne: []byte{}}, expected)
	require.Equal(t, TwentyOne{}, actual)

	actual = TwentyOne{
		One:       rando.Bytes(),
		Two:       rando.Bytes(),
		Three:     rando.Bytes(),
		Four:      rando.Bytes(),
		Five:      rando.Bytes(),
		Six:       rando.Bytes(),
		Seven:     rando.Bytes(),
		Eight:     rando.Bytes(),
		Nine:      rando.Bytes(),
		Ten:       rando.Bytes(),
		Eleven:    rando.Bytes(),
		Twelve:    rando.Bytes(),
		Thirteen:  rando.Bytes(),
		Fourteen:  rando.Bytes(),
		Fifteen:   rando.Bytes(),
		Sixteen:   rando.Bytes(),
		Seventeen: rando.Bytes(),
		Eighteen:  rando.Bytes(),
		Nineteen:  rando.Bytes(),
		Twenty:    rando.Bytes(),
		TwentyOne: rando.Bytes(),
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
	require.NotEqual(t, expected, actual)
	require.Equal(t, TwentyTwo{One: []byte{}, Two: []byte{}, Three: []byte{}, Four: []byte{}, Five: []byte{}, Six: []byte{}, Seven: []byte{}, Eight: []byte{}, Nine: []byte{}, Ten: []byte{}, Eleven: []byte{}, Twelve: []byte{}, Thirteen: []byte{}, Fourteen: []byte{}, Fifteen: []byte{}, Sixteen: []byte{}, Seventeen: []byte{}, Eighteen: []byte{}, Nineteen: []byte{}, Twenty: []byte{}, TwentyOne: []byte{}, TwentyTwo: []byte{}}, expected)
	require.Equal(t, TwentyTwo{}, actual)

	actual = TwentyTwo{
		One:       rando.Bytes(),
		Two:       rando.Bytes(),
		Three:     rando.Bytes(),
		Four:      rando.Bytes(),
		Five:      rando.Bytes(),
		Six:       rando.Bytes(),
		Seven:     rando.Bytes(),
		Eight:     rando.Bytes(),
		Nine:      rando.Bytes(),
		Ten:       rando.Bytes(),
		Eleven:    rando.Bytes(),
		Twelve:    rando.Bytes(),
		Thirteen:  rando.Bytes(),
		Fourteen:  rando.Bytes(),
		Fifteen:   rando.Bytes(),
		Sixteen:   rando.Bytes(),
		Seventeen: rando.Bytes(),
		Eighteen:  rando.Bytes(),
		Nineteen:  rando.Bytes(),
		Twenty:    rando.Bytes(),
		TwentyOne: rando.Bytes(),
		TwentyTwo: rando.Bytes(),
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
	require.NotEqual(t, expected, actual)
	require.Equal(t, TwentyThree{One: []byte{}, Two: []byte{}, Three: []byte{}, Four: []byte{}, Five: []byte{}, Six: []byte{}, Seven: []byte{}, Eight: []byte{}, Nine: []byte{}, Ten: []byte{}, Eleven: []byte{}, Twelve: []byte{}, Thirteen: []byte{}, Fourteen: []byte{}, Fifteen: []byte{}, Sixteen: []byte{}, Seventeen: []byte{}, Eighteen: []byte{}, Nineteen: []byte{}, Twenty: []byte{}, TwentyOne: []byte{}, TwentyTwo: []byte{}, TwentyThree: []byte{}}, expected)
	require.Equal(t, TwentyThree{}, actual)

	actual = TwentyThree{
		One:         rando.Bytes(),
		Two:         rando.Bytes(),
		Three:       rando.Bytes(),
		Four:        rando.Bytes(),
		Five:        rando.Bytes(),
		Six:         rando.Bytes(),
		Seven:       rando.Bytes(),
		Eight:       rando.Bytes(),
		Nine:        rando.Bytes(),
		Ten:         rando.Bytes(),
		Eleven:      rando.Bytes(),
		Twelve:      rando.Bytes(),
		Thirteen:    rando.Bytes(),
		Fourteen:    rando.Bytes(),
		Fifteen:     rando.Bytes(),
		Sixteen:     rando.Bytes(),
		Seventeen:   rando.Bytes(),
		Eighteen:    rando.Bytes(),
		Nineteen:    rando.Bytes(),
		Twenty:      rando.Bytes(),
		TwentyOne:   rando.Bytes(),
		TwentyTwo:   rando.Bytes(),
		TwentyThree: rando.Bytes(),
	}
	src := actual.MarshalJ()
	require.NoError(t, expected.UnmarshalJ(src))
	// require.NotEqual(t, TwentyThree{}, expected)
	// require.NotEqual(t, TwentyThree{}, actual)
	require.Equal(t, expected, actual)
}
