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
	require.Equal(t, One{One: []byte{}}, actual)

	expected = One{
		One: rando.Bytes(),
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
	require.Equal(t, Two{One: []byte{}, Two: []byte{}}, actual)

	expected = Two{
		One: rando.Bytes(),
		Two: rando.Bytes(),
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
	require.Equal(t, Three{One: []byte{}, Two: []byte{}, Three: []byte{}}, actual)

	expected = Three{
		One:   rando.Bytes(),
		Two:   rando.Bytes(),
		Three: rando.Bytes(),
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
	require.Equal(t, Four{One: []byte{}, Two: []byte{}, Three: []byte{}, Four: []byte{}}, actual)

	expected = Four{
		One:   rando.Bytes(),
		Two:   rando.Bytes(),
		Three: rando.Bytes(),
		Four:  rando.Bytes(),
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
	require.Equal(t, Five{One: []byte{}, Two: []byte{}, Three: []byte{}, Four: []byte{}, Five: []byte{}}, actual)

	expected = Five{
		One:   rando.Bytes(),
		Two:   rando.Bytes(),
		Three: rando.Bytes(),
		Four:  rando.Bytes(),
		Five:  rando.Bytes(),
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
	require.Equal(t, Six{One: []byte{}, Two: []byte{}, Three: []byte{}, Four: []byte{}, Five: []byte{}, Six: []byte{}}, actual)

	expected = Six{
		One:   rando.Bytes(),
		Two:   rando.Bytes(),
		Three: rando.Bytes(),
		Four:  rando.Bytes(),
		Five:  rando.Bytes(),
		Six:   rando.Bytes(),
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
	require.Equal(t, Seven{One: []byte{}, Two: []byte{}, Three: []byte{}, Four: []byte{}, Five: []byte{}, Six: []byte{}, Seven: []byte{}}, actual)

	expected = Seven{
		One:   rando.Bytes(),
		Two:   rando.Bytes(),
		Three: rando.Bytes(),
		Four:  rando.Bytes(),
		Five:  rando.Bytes(),
		Six:   rando.Bytes(),
		Seven: rando.Bytes(),
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
	require.Equal(t, Eight{One: []byte{}, Two: []byte{}, Three: []byte{}, Four: []byte{}, Five: []byte{}, Six: []byte{}, Seven: []byte{}, Eight: []byte{}}, actual)

	expected = Eight{
		One:   rando.Bytes(),
		Two:   rando.Bytes(),
		Three: rando.Bytes(),
		Four:  rando.Bytes(),
		Five:  rando.Bytes(),
		Six:   rando.Bytes(),
		Seven: rando.Bytes(),
		Eight: rando.Bytes(),
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
	require.Equal(t, Nine{One: []byte{}, Two: []byte{}, Three: []byte{}, Four: []byte{}, Five: []byte{}, Six: []byte{}, Seven: []byte{}, Eight: []byte{}, Nine: []byte{}}, actual)

	expected = Nine{
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
	require.Equal(t, Ten{One: []byte{}, Two: []byte{}, Three: []byte{}, Four: []byte{}, Five: []byte{}, Six: []byte{}, Seven: []byte{}, Eight: []byte{}, Nine: []byte{}, Ten: []byte{}}, actual)

	expected = Ten{
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
	require.Equal(t, Eleven{One: []byte{}, Two: []byte{}, Three: []byte{}, Four: []byte{}, Five: []byte{}, Six: []byte{}, Seven: []byte{}, Eight: []byte{}, Nine: []byte{}, Ten: []byte{}, Eleven: []byte{}}, actual)

	expected = Eleven{
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
	require.Equal(t, Twelve{One: []byte{}, Two: []byte{}, Three: []byte{}, Four: []byte{}, Five: []byte{}, Six: []byte{}, Seven: []byte{}, Eight: []byte{}, Nine: []byte{}, Ten: []byte{}, Eleven: []byte{}, Twelve: []byte{}}, actual)

	expected = Twelve{
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
	require.Equal(t, Thirteen{One: []byte{}, Two: []byte{}, Three: []byte{}, Four: []byte{}, Five: []byte{}, Six: []byte{}, Seven: []byte{}, Eight: []byte{}, Nine: []byte{}, Ten: []byte{}, Eleven: []byte{}, Twelve: []byte{}, Thirteen: []byte{}}, actual)

	expected = Thirteen{
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
	require.Equal(t, Fourteen{One: []byte{}, Two: []byte{}, Three: []byte{}, Four: []byte{}, Five: []byte{}, Six: []byte{}, Seven: []byte{}, Eight: []byte{}, Nine: []byte{}, Ten: []byte{}, Eleven: []byte{}, Twelve: []byte{}, Thirteen: []byte{}, Fourteen: []byte{}}, actual)

	expected = Fourteen{
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
	require.Equal(t, Fifteen{One: []byte{}, Two: []byte{}, Three: []byte{}, Four: []byte{}, Five: []byte{}, Six: []byte{}, Seven: []byte{}, Eight: []byte{}, Nine: []byte{}, Ten: []byte{}, Eleven: []byte{}, Twelve: []byte{}, Thirteen: []byte{}, Fourteen: []byte{}, Fifteen: []byte{}}, actual)

	expected = Fifteen{
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
	require.Equal(t, Sixteen{One: []byte{}, Two: []byte{}, Three: []byte{}, Four: []byte{}, Five: []byte{}, Six: []byte{}, Seven: []byte{}, Eight: []byte{}, Nine: []byte{}, Ten: []byte{}, Eleven: []byte{}, Twelve: []byte{}, Thirteen: []byte{}, Fourteen: []byte{}, Fifteen: []byte{}, Sixteen: []byte{}}, actual)

	expected = Sixteen{
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
	require.Equal(t, Seventeen{One: []byte{}, Two: []byte{}, Three: []byte{}, Four: []byte{}, Five: []byte{}, Six: []byte{}, Seven: []byte{}, Eight: []byte{}, Nine: []byte{}, Ten: []byte{}, Eleven: []byte{}, Twelve: []byte{}, Thirteen: []byte{}, Fourteen: []byte{}, Fifteen: []byte{}, Sixteen: []byte{}, Seventeen: []byte{}}, actual)

	expected = Seventeen{
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
	require.Equal(t, Eighteen{One: []byte{}, Two: []byte{}, Three: []byte{}, Four: []byte{}, Five: []byte{}, Six: []byte{}, Seven: []byte{}, Eight: []byte{}, Nine: []byte{}, Ten: []byte{}, Eleven: []byte{}, Twelve: []byte{}, Thirteen: []byte{}, Fourteen: []byte{}, Fifteen: []byte{}, Sixteen: []byte{}, Seventeen: []byte{}, Eighteen: []byte{}}, actual)

	expected = Eighteen{
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
	require.Equal(t, Nineteen{One: []byte{}, Two: []byte{}, Three: []byte{}, Four: []byte{}, Five: []byte{}, Six: []byte{}, Seven: []byte{}, Eight: []byte{}, Nine: []byte{}, Ten: []byte{}, Eleven: []byte{}, Twelve: []byte{}, Thirteen: []byte{}, Fourteen: []byte{}, Fifteen: []byte{}, Sixteen: []byte{}, Seventeen: []byte{}, Eighteen: []byte{}, Nineteen: []byte{}}, actual)

	expected = Nineteen{
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
	require.Equal(t, Twenty{One: []byte{}, Two: []byte{}, Three: []byte{}, Four: []byte{}, Five: []byte{}, Six: []byte{}, Seven: []byte{}, Eight: []byte{}, Nine: []byte{}, Ten: []byte{}, Eleven: []byte{}, Twelve: []byte{}, Thirteen: []byte{}, Fourteen: []byte{}, Fifteen: []byte{}, Sixteen: []byte{}, Seventeen: []byte{}, Eighteen: []byte{}, Nineteen: []byte{}, Twenty: []byte{}}, actual)

	expected = Twenty{
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
	require.Equal(t, TwentyOne{One: []byte{}, Two: []byte{}, Three: []byte{}, Four: []byte{}, Five: []byte{}, Six: []byte{}, Seven: []byte{}, Eight: []byte{}, Nine: []byte{}, Ten: []byte{}, Eleven: []byte{}, Twelve: []byte{}, Thirteen: []byte{}, Fourteen: []byte{}, Fifteen: []byte{}, Sixteen: []byte{}, Seventeen: []byte{}, Eighteen: []byte{}, Nineteen: []byte{}, Twenty: []byte{}, TwentyOne: []byte{}}, actual)

	expected = TwentyOne{
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
	require.Equal(t, TwentyTwo{One: []byte{}, Two: []byte{}, Three: []byte{}, Four: []byte{}, Five: []byte{}, Six: []byte{}, Seven: []byte{}, Eight: []byte{}, Nine: []byte{}, Ten: []byte{}, Eleven: []byte{}, Twelve: []byte{}, Thirteen: []byte{}, Fourteen: []byte{}, Fifteen: []byte{}, Sixteen: []byte{}, Seventeen: []byte{}, Eighteen: []byte{}, Nineteen: []byte{}, Twenty: []byte{}, TwentyOne: []byte{}, TwentyTwo: []byte{}}, actual)

	expected = TwentyTwo{
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
	require.Equal(t, TwentyThree{One: []byte{}, Two: []byte{}, Three: []byte{}, Four: []byte{}, Five: []byte{}, Six: []byte{}, Seven: []byte{}, Eight: []byte{}, Nine: []byte{}, Ten: []byte{}, Eleven: []byte{}, Twelve: []byte{}, Thirteen: []byte{}, Fourteen: []byte{}, Fifteen: []byte{}, Sixteen: []byte{}, Seventeen: []byte{}, Eighteen: []byte{}, Nineteen: []byte{}, Twenty: []byte{}, TwentyOne: []byte{}, TwentyTwo: []byte{}, TwentyThree: []byte{}}, actual)

	expected = TwentyThree{
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
	src := expected.MarshalJ()
	require.NoError(t, actual.UnmarshalJ(src))
	// require.NotEqual(t, TwentyThree{}, expected)
	// require.NotEqual(t, TwentyThree{}, actual)
	require.Equal(t, expected, actual)
}
