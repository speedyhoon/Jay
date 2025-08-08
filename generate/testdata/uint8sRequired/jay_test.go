package main

import (
	"testing"

	"github.com/speedyhoon/rando"
	"github.com/stretchr/testify/require"
)

func TestFuzz_1(t *testing.T) {
	var expected, actual One
	require.NoError(t, actual.UnmarshalJ(expected.MarshalJ()))
	require.NotEqual(t, expected, actual)
	require.Equal(t, One{}, expected)
	require.Equal(t, One{One: []uint8{}}, actual)

	expected = One{
		One: rando.Uint8s(),
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
	require.NotEqual(t, expected, actual)
	require.Equal(t, Two{}, expected)
	require.Equal(t, Two{One: []uint8{}, Two: []uint8{}}, actual)

	expected = Two{
		One: rando.Uint8s(),
		Two: rando.Uint8s(),
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
	require.NotEqual(t, expected, actual)
	require.Equal(t, Three{}, expected)
	require.Equal(t, Three{One: []uint8{}, Two: []uint8{}, Three: []uint8{}}, actual)

	expected = Three{
		One:   rando.Uint8s(),
		Two:   rando.Uint8s(),
		Three: rando.Uint8s(),
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
	require.NotEqual(t, expected, actual)
	require.Equal(t, Four{}, expected)
	require.Equal(t, Four{One: []uint8{}, Two: []uint8{}, Three: []uint8{}, Four: []uint8{}}, actual)

	expected = Four{
		One:   rando.Uint8s(),
		Two:   rando.Uint8s(),
		Three: rando.Uint8s(),
		Four:  rando.Uint8s(),
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
	require.NotEqual(t, expected, actual)
	require.Equal(t, Five{}, expected)
	require.Equal(t, Five{One: []uint8{}, Two: []uint8{}, Three: []uint8{}, Four: []uint8{}, Five: []uint8{}}, actual)

	expected = Five{
		One:   rando.Uint8s(),
		Two:   rando.Uint8s(),
		Three: rando.Uint8s(),
		Four:  rando.Uint8s(),
		Five:  rando.Uint8s(),
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
	require.NotEqual(t, expected, actual)
	require.Equal(t, Six{}, expected)
	require.Equal(t, Six{One: []uint8{}, Two: []uint8{}, Three: []uint8{}, Four: []uint8{}, Five: []uint8{}, Six: []uint8{}}, actual)

	expected = Six{
		One:   rando.Uint8s(),
		Two:   rando.Uint8s(),
		Three: rando.Uint8s(),
		Four:  rando.Uint8s(),
		Five:  rando.Uint8s(),
		Six:   rando.Uint8s(),
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
	require.NotEqual(t, expected, actual)
	require.Equal(t, Seven{}, expected)
	require.Equal(t, Seven{One: []uint8{}, Two: []uint8{}, Three: []uint8{}, Four: []uint8{}, Five: []uint8{}, Six: []uint8{}, Seven: []uint8{}}, actual)

	expected = Seven{
		One:   rando.Uint8s(),
		Two:   rando.Uint8s(),
		Three: rando.Uint8s(),
		Four:  rando.Uint8s(),
		Five:  rando.Uint8s(),
		Six:   rando.Uint8s(),
		Seven: rando.Uint8s(),
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
	require.NotEqual(t, expected, actual)
	require.Equal(t, Eight{}, expected)
	require.Equal(t, Eight{One: []uint8{}, Two: []uint8{}, Three: []uint8{}, Four: []uint8{}, Five: []uint8{}, Six: []uint8{}, Seven: []uint8{}, Eight: []uint8{}}, actual)

	expected = Eight{
		One:   rando.Uint8s(),
		Two:   rando.Uint8s(),
		Three: rando.Uint8s(),
		Four:  rando.Uint8s(),
		Five:  rando.Uint8s(),
		Six:   rando.Uint8s(),
		Seven: rando.Uint8s(),
		Eight: rando.Uint8s(),
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
	require.NotEqual(t, expected, actual)
	require.Equal(t, Nine{}, expected)
	require.Equal(t, Nine{One: []uint8{}, Two: []uint8{}, Three: []uint8{}, Four: []uint8{}, Five: []uint8{}, Six: []uint8{}, Seven: []uint8{}, Eight: []uint8{}, Nine: []uint8{}}, actual)

	expected = Nine{
		One:   rando.Uint8s(),
		Two:   rando.Uint8s(),
		Three: rando.Uint8s(),
		Four:  rando.Uint8s(),
		Five:  rando.Uint8s(),
		Six:   rando.Uint8s(),
		Seven: rando.Uint8s(),
		Eight: rando.Uint8s(),
		Nine:  rando.Uint8s(),
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
	require.NotEqual(t, expected, actual)
	require.Equal(t, Ten{}, expected)
	require.Equal(t, Ten{One: []uint8{}, Two: []uint8{}, Three: []uint8{}, Four: []uint8{}, Five: []uint8{}, Six: []uint8{}, Seven: []uint8{}, Eight: []uint8{}, Nine: []uint8{}, Ten: []uint8{}}, actual)

	expected = Ten{
		One:   rando.Uint8s(),
		Two:   rando.Uint8s(),
		Three: rando.Uint8s(),
		Four:  rando.Uint8s(),
		Five:  rando.Uint8s(),
		Six:   rando.Uint8s(),
		Seven: rando.Uint8s(),
		Eight: rando.Uint8s(),
		Nine:  rando.Uint8s(),
		Ten:   rando.Uint8s(),
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
	require.NotEqual(t, expected, actual)
	require.Equal(t, Eleven{}, expected)
	require.Equal(t, Eleven{One: []uint8{}, Two: []uint8{}, Three: []uint8{}, Four: []uint8{}, Five: []uint8{}, Six: []uint8{}, Seven: []uint8{}, Eight: []uint8{}, Nine: []uint8{}, Ten: []uint8{}, Eleven: []uint8{}}, actual)

	expected = Eleven{
		One:    rando.Uint8s(),
		Two:    rando.Uint8s(),
		Three:  rando.Uint8s(),
		Four:   rando.Uint8s(),
		Five:   rando.Uint8s(),
		Six:    rando.Uint8s(),
		Seven:  rando.Uint8s(),
		Eight:  rando.Uint8s(),
		Nine:   rando.Uint8s(),
		Ten:    rando.Uint8s(),
		Eleven: rando.Uint8s(),
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
	require.NotEqual(t, expected, actual)
	require.Equal(t, Twelve{}, expected)
	require.Equal(t, Twelve{One: []uint8{}, Two: []uint8{}, Three: []uint8{}, Four: []uint8{}, Five: []uint8{}, Six: []uint8{}, Seven: []uint8{}, Eight: []uint8{}, Nine: []uint8{}, Ten: []uint8{}, Eleven: []uint8{}, Twelve: []uint8{}}, actual)

	expected = Twelve{
		One:    rando.Uint8s(),
		Two:    rando.Uint8s(),
		Three:  rando.Uint8s(),
		Four:   rando.Uint8s(),
		Five:   rando.Uint8s(),
		Six:    rando.Uint8s(),
		Seven:  rando.Uint8s(),
		Eight:  rando.Uint8s(),
		Nine:   rando.Uint8s(),
		Ten:    rando.Uint8s(),
		Eleven: rando.Uint8s(),
		Twelve: rando.Uint8s(),
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
	require.NotEqual(t, expected, actual)
	require.Equal(t, Thirteen{}, expected)
	require.Equal(t, Thirteen{One: []uint8{}, Two: []uint8{}, Three: []uint8{}, Four: []uint8{}, Five: []uint8{}, Six: []uint8{}, Seven: []uint8{}, Eight: []uint8{}, Nine: []uint8{}, Ten: []uint8{}, Eleven: []uint8{}, Twelve: []uint8{}, Thirteen: []uint8{}}, actual)

	expected = Thirteen{
		One:      rando.Uint8s(),
		Two:      rando.Uint8s(),
		Three:    rando.Uint8s(),
		Four:     rando.Uint8s(),
		Five:     rando.Uint8s(),
		Six:      rando.Uint8s(),
		Seven:    rando.Uint8s(),
		Eight:    rando.Uint8s(),
		Nine:     rando.Uint8s(),
		Ten:      rando.Uint8s(),
		Eleven:   rando.Uint8s(),
		Twelve:   rando.Uint8s(),
		Thirteen: rando.Uint8s(),
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
	require.NotEqual(t, expected, actual)
	require.Equal(t, Fourteen{}, expected)
	require.Equal(t, Fourteen{One: []uint8{}, Two: []uint8{}, Three: []uint8{}, Four: []uint8{}, Five: []uint8{}, Six: []uint8{}, Seven: []uint8{}, Eight: []uint8{}, Nine: []uint8{}, Ten: []uint8{}, Eleven: []uint8{}, Twelve: []uint8{}, Thirteen: []uint8{}, Fourteen: []uint8{}}, actual)

	expected = Fourteen{
		One:      rando.Uint8s(),
		Two:      rando.Uint8s(),
		Three:    rando.Uint8s(),
		Four:     rando.Uint8s(),
		Five:     rando.Uint8s(),
		Six:      rando.Uint8s(),
		Seven:    rando.Uint8s(),
		Eight:    rando.Uint8s(),
		Nine:     rando.Uint8s(),
		Ten:      rando.Uint8s(),
		Eleven:   rando.Uint8s(),
		Twelve:   rando.Uint8s(),
		Thirteen: rando.Uint8s(),
		Fourteen: rando.Uint8s(),
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
	require.NotEqual(t, expected, actual)
	require.Equal(t, Fifteen{}, expected)
	require.Equal(t, Fifteen{One: []uint8{}, Two: []uint8{}, Three: []uint8{}, Four: []uint8{}, Five: []uint8{}, Six: []uint8{}, Seven: []uint8{}, Eight: []uint8{}, Nine: []uint8{}, Ten: []uint8{}, Eleven: []uint8{}, Twelve: []uint8{}, Thirteen: []uint8{}, Fourteen: []uint8{}, Fifteen: []uint8{}}, actual)

	expected = Fifteen{
		One:      rando.Uint8s(),
		Two:      rando.Uint8s(),
		Three:    rando.Uint8s(),
		Four:     rando.Uint8s(),
		Five:     rando.Uint8s(),
		Six:      rando.Uint8s(),
		Seven:    rando.Uint8s(),
		Eight:    rando.Uint8s(),
		Nine:     rando.Uint8s(),
		Ten:      rando.Uint8s(),
		Eleven:   rando.Uint8s(),
		Twelve:   rando.Uint8s(),
		Thirteen: rando.Uint8s(),
		Fourteen: rando.Uint8s(),
		Fifteen:  rando.Uint8s(),
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
	require.NotEqual(t, expected, actual)
	require.Equal(t, Sixteen{}, expected)
	require.Equal(t, Sixteen{One: []uint8{}, Two: []uint8{}, Three: []uint8{}, Four: []uint8{}, Five: []uint8{}, Six: []uint8{}, Seven: []uint8{}, Eight: []uint8{}, Nine: []uint8{}, Ten: []uint8{}, Eleven: []uint8{}, Twelve: []uint8{}, Thirteen: []uint8{}, Fourteen: []uint8{}, Fifteen: []uint8{}, Sixteen: []uint8{}}, actual)

	expected = Sixteen{
		One:      rando.Uint8s(),
		Two:      rando.Uint8s(),
		Three:    rando.Uint8s(),
		Four:     rando.Uint8s(),
		Five:     rando.Uint8s(),
		Six:      rando.Uint8s(),
		Seven:    rando.Uint8s(),
		Eight:    rando.Uint8s(),
		Nine:     rando.Uint8s(),
		Ten:      rando.Uint8s(),
		Eleven:   rando.Uint8s(),
		Twelve:   rando.Uint8s(),
		Thirteen: rando.Uint8s(),
		Fourteen: rando.Uint8s(),
		Fifteen:  rando.Uint8s(),
		Sixteen:  rando.Uint8s(),
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
	require.NotEqual(t, expected, actual)
	require.Equal(t, Seventeen{}, expected)
	require.Equal(t, Seventeen{One: []uint8{}, Two: []uint8{}, Three: []uint8{}, Four: []uint8{}, Five: []uint8{}, Six: []uint8{}, Seven: []uint8{}, Eight: []uint8{}, Nine: []uint8{}, Ten: []uint8{}, Eleven: []uint8{}, Twelve: []uint8{}, Thirteen: []uint8{}, Fourteen: []uint8{}, Fifteen: []uint8{}, Sixteen: []uint8{}, Seventeen: []uint8{}}, actual)

	expected = Seventeen{
		One:       rando.Uint8s(),
		Two:       rando.Uint8s(),
		Three:     rando.Uint8s(),
		Four:      rando.Uint8s(),
		Five:      rando.Uint8s(),
		Six:       rando.Uint8s(),
		Seven:     rando.Uint8s(),
		Eight:     rando.Uint8s(),
		Nine:      rando.Uint8s(),
		Ten:       rando.Uint8s(),
		Eleven:    rando.Uint8s(),
		Twelve:    rando.Uint8s(),
		Thirteen:  rando.Uint8s(),
		Fourteen:  rando.Uint8s(),
		Fifteen:   rando.Uint8s(),
		Sixteen:   rando.Uint8s(),
		Seventeen: rando.Uint8s(),
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
	require.NotEqual(t, expected, actual)
	require.Equal(t, Eighteen{}, expected)
	require.Equal(t, Eighteen{One: []uint8{}, Two: []uint8{}, Three: []uint8{}, Four: []uint8{}, Five: []uint8{}, Six: []uint8{}, Seven: []uint8{}, Eight: []uint8{}, Nine: []uint8{}, Ten: []uint8{}, Eleven: []uint8{}, Twelve: []uint8{}, Thirteen: []uint8{}, Fourteen: []uint8{}, Fifteen: []uint8{}, Sixteen: []uint8{}, Seventeen: []uint8{}, Eighteen: []uint8{}}, actual)

	expected = Eighteen{
		One:       rando.Uint8s(),
		Two:       rando.Uint8s(),
		Three:     rando.Uint8s(),
		Four:      rando.Uint8s(),
		Five:      rando.Uint8s(),
		Six:       rando.Uint8s(),
		Seven:     rando.Uint8s(),
		Eight:     rando.Uint8s(),
		Nine:      rando.Uint8s(),
		Ten:       rando.Uint8s(),
		Eleven:    rando.Uint8s(),
		Twelve:    rando.Uint8s(),
		Thirteen:  rando.Uint8s(),
		Fourteen:  rando.Uint8s(),
		Fifteen:   rando.Uint8s(),
		Sixteen:   rando.Uint8s(),
		Seventeen: rando.Uint8s(),
		Eighteen:  rando.Uint8s(),
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
	require.NotEqual(t, expected, actual)
	require.Equal(t, Nineteen{}, expected)
	require.Equal(t, Nineteen{One: []uint8{}, Two: []uint8{}, Three: []uint8{}, Four: []uint8{}, Five: []uint8{}, Six: []uint8{}, Seven: []uint8{}, Eight: []uint8{}, Nine: []uint8{}, Ten: []uint8{}, Eleven: []uint8{}, Twelve: []uint8{}, Thirteen: []uint8{}, Fourteen: []uint8{}, Fifteen: []uint8{}, Sixteen: []uint8{}, Seventeen: []uint8{}, Eighteen: []uint8{}, Nineteen: []uint8{}}, actual)

	expected = Nineteen{
		One:       rando.Uint8s(),
		Two:       rando.Uint8s(),
		Three:     rando.Uint8s(),
		Four:      rando.Uint8s(),
		Five:      rando.Uint8s(),
		Six:       rando.Uint8s(),
		Seven:     rando.Uint8s(),
		Eight:     rando.Uint8s(),
		Nine:      rando.Uint8s(),
		Ten:       rando.Uint8s(),
		Eleven:    rando.Uint8s(),
		Twelve:    rando.Uint8s(),
		Thirteen:  rando.Uint8s(),
		Fourteen:  rando.Uint8s(),
		Fifteen:   rando.Uint8s(),
		Sixteen:   rando.Uint8s(),
		Seventeen: rando.Uint8s(),
		Eighteen:  rando.Uint8s(),
		Nineteen:  rando.Uint8s(),
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
	require.NotEqual(t, expected, actual)
	require.Equal(t, Twenty{}, expected)
	require.Equal(t, Twenty{One: []uint8{}, Two: []uint8{}, Three: []uint8{}, Four: []uint8{}, Five: []uint8{}, Six: []uint8{}, Seven: []uint8{}, Eight: []uint8{}, Nine: []uint8{}, Ten: []uint8{}, Eleven: []uint8{}, Twelve: []uint8{}, Thirteen: []uint8{}, Fourteen: []uint8{}, Fifteen: []uint8{}, Sixteen: []uint8{}, Seventeen: []uint8{}, Eighteen: []uint8{}, Nineteen: []uint8{}, Twenty: []uint8{}}, actual)

	expected = Twenty{
		One:       rando.Uint8s(),
		Two:       rando.Uint8s(),
		Three:     rando.Uint8s(),
		Four:      rando.Uint8s(),
		Five:      rando.Uint8s(),
		Six:       rando.Uint8s(),
		Seven:     rando.Uint8s(),
		Eight:     rando.Uint8s(),
		Nine:      rando.Uint8s(),
		Ten:       rando.Uint8s(),
		Eleven:    rando.Uint8s(),
		Twelve:    rando.Uint8s(),
		Thirteen:  rando.Uint8s(),
		Fourteen:  rando.Uint8s(),
		Fifteen:   rando.Uint8s(),
		Sixteen:   rando.Uint8s(),
		Seventeen: rando.Uint8s(),
		Eighteen:  rando.Uint8s(),
		Nineteen:  rando.Uint8s(),
		Twenty:    rando.Uint8s(),
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
	require.NotEqual(t, expected, actual)
	require.Equal(t, TwentyOne{}, expected)
	require.Equal(t, TwentyOne{One: []uint8{}, Two: []uint8{}, Three: []uint8{}, Four: []uint8{}, Five: []uint8{}, Six: []uint8{}, Seven: []uint8{}, Eight: []uint8{}, Nine: []uint8{}, Ten: []uint8{}, Eleven: []uint8{}, Twelve: []uint8{}, Thirteen: []uint8{}, Fourteen: []uint8{}, Fifteen: []uint8{}, Sixteen: []uint8{}, Seventeen: []uint8{}, Eighteen: []uint8{}, Nineteen: []uint8{}, Twenty: []uint8{}, TwentyOne: []uint8{}}, actual)

	expected = TwentyOne{
		One:       rando.Uint8s(),
		Two:       rando.Uint8s(),
		Three:     rando.Uint8s(),
		Four:      rando.Uint8s(),
		Five:      rando.Uint8s(),
		Six:       rando.Uint8s(),
		Seven:     rando.Uint8s(),
		Eight:     rando.Uint8s(),
		Nine:      rando.Uint8s(),
		Ten:       rando.Uint8s(),
		Eleven:    rando.Uint8s(),
		Twelve:    rando.Uint8s(),
		Thirteen:  rando.Uint8s(),
		Fourteen:  rando.Uint8s(),
		Fifteen:   rando.Uint8s(),
		Sixteen:   rando.Uint8s(),
		Seventeen: rando.Uint8s(),
		Eighteen:  rando.Uint8s(),
		Nineteen:  rando.Uint8s(),
		Twenty:    rando.Uint8s(),
		TwentyOne: rando.Uint8s(),
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
	require.NotEqual(t, expected, actual)
	require.Equal(t, TwentyTwo{}, expected)
	require.Equal(t, TwentyTwo{One: []uint8{}, Two: []uint8{}, Three: []uint8{}, Four: []uint8{}, Five: []uint8{}, Six: []uint8{}, Seven: []uint8{}, Eight: []uint8{}, Nine: []uint8{}, Ten: []uint8{}, Eleven: []uint8{}, Twelve: []uint8{}, Thirteen: []uint8{}, Fourteen: []uint8{}, Fifteen: []uint8{}, Sixteen: []uint8{}, Seventeen: []uint8{}, Eighteen: []uint8{}, Nineteen: []uint8{}, Twenty: []uint8{}, TwentyOne: []uint8{}, TwentyTwo: []uint8{}}, actual)

	expected = TwentyTwo{
		One:       rando.Uint8s(),
		Two:       rando.Uint8s(),
		Three:     rando.Uint8s(),
		Four:      rando.Uint8s(),
		Five:      rando.Uint8s(),
		Six:       rando.Uint8s(),
		Seven:     rando.Uint8s(),
		Eight:     rando.Uint8s(),
		Nine:      rando.Uint8s(),
		Ten:       rando.Uint8s(),
		Eleven:    rando.Uint8s(),
		Twelve:    rando.Uint8s(),
		Thirteen:  rando.Uint8s(),
		Fourteen:  rando.Uint8s(),
		Fifteen:   rando.Uint8s(),
		Sixteen:   rando.Uint8s(),
		Seventeen: rando.Uint8s(),
		Eighteen:  rando.Uint8s(),
		Nineteen:  rando.Uint8s(),
		Twenty:    rando.Uint8s(),
		TwentyOne: rando.Uint8s(),
		TwentyTwo: rando.Uint8s(),
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
	require.NotEqual(t, expected, actual)
	require.Equal(t, TwentyThree{}, expected)
	require.Equal(t, TwentyThree{One: []uint8{}, Two: []uint8{}, Three: []uint8{}, Four: []uint8{}, Five: []uint8{}, Six: []uint8{}, Seven: []uint8{}, Eight: []uint8{}, Nine: []uint8{}, Ten: []uint8{}, Eleven: []uint8{}, Twelve: []uint8{}, Thirteen: []uint8{}, Fourteen: []uint8{}, Fifteen: []uint8{}, Sixteen: []uint8{}, Seventeen: []uint8{}, Eighteen: []uint8{}, Nineteen: []uint8{}, Twenty: []uint8{}, TwentyOne: []uint8{}, TwentyTwo: []uint8{}, TwentyThree: []uint8{}}, actual)

	expected = TwentyThree{
		One:         rando.Uint8s(),
		Two:         rando.Uint8s(),
		Three:       rando.Uint8s(),
		Four:        rando.Uint8s(),
		Five:        rando.Uint8s(),
		Six:         rando.Uint8s(),
		Seven:       rando.Uint8s(),
		Eight:       rando.Uint8s(),
		Nine:        rando.Uint8s(),
		Ten:         rando.Uint8s(),
		Eleven:      rando.Uint8s(),
		Twelve:      rando.Uint8s(),
		Thirteen:    rando.Uint8s(),
		Fourteen:    rando.Uint8s(),
		Fifteen:     rando.Uint8s(),
		Sixteen:     rando.Uint8s(),
		Seventeen:   rando.Uint8s(),
		Eighteen:    rando.Uint8s(),
		Nineteen:    rando.Uint8s(),
		Twenty:      rando.Uint8s(),
		TwentyOne:   rando.Uint8s(),
		TwentyTwo:   rando.Uint8s(),
		TwentyThree: rando.Uint8s(),
	}
	src := expected.MarshalJ()
	require.NoError(t, actual.UnmarshalJ(src))
	// require.NotEqual(t, TwentyThree{}, expected)
	// require.NotEqual(t, TwentyThree{}, actual)
	require.Equal(t, expected, actual)
}
