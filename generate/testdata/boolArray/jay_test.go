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
		One: [4]bool(rando.BoolsN(4)),
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
		One: [2]bool(rando.BoolsN(2)),
		Two: [7]bool(rando.BoolsN(7)),
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
		One:   [8]bool(rando.BoolsN(8)),
		Two:   [1]bool(rando.BoolsN(1)),
		Three: [9]bool(rando.BoolsN(9)),
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
		One:   [7]bool(rando.BoolsN(7)),
		Two:   [2]bool(rando.BoolsN(2)),
		Three: [6]bool(rando.BoolsN(6)),
		Four:  [8]bool(rando.BoolsN(8)),
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
		One:   [2]bool(rando.BoolsN(2)),
		Two:   [3]bool(rando.BoolsN(3)),
		Three: [4]bool(rando.BoolsN(4)),
		Four:  [8]bool(rando.BoolsN(8)),
		Five:  [1]bool(rando.BoolsN(1)),
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
		One:   [9]bool(rando.BoolsN(9)),
		Two:   [5]bool(rando.BoolsN(5)),
		Three: [7]bool(rando.BoolsN(7)),
		Four:  [2]bool(rando.BoolsN(2)),
		Five:  [6]bool(rando.BoolsN(6)),
		Six:   [10]bool(rando.BoolsN(10)),
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
		One:   [35]bool(rando.BoolsN(35)),
		Two:   [1]bool(rando.BoolsN(1)),
		Three: [2]bool(rando.BoolsN(2)),
		Four:  [3]bool(rando.BoolsN(3)),
		Five:  [27]bool(rando.BoolsN(27)),
		Six:   [10]bool(rando.BoolsN(10)),
		Seven: [9]bool(rando.BoolsN(9)),
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
		One:   [11]bool(rando.BoolsN(11)),
		Two:   [12]bool(rando.BoolsN(12)),
		Three: [13]bool(rando.BoolsN(13)),
		Four:  [14]bool(rando.BoolsN(14)),
		Five:  [15]bool(rando.BoolsN(15)),
		Six:   [16]bool(rando.BoolsN(16)),
		Seven: [17]bool(rando.BoolsN(17)),
		Eight: [18]bool(rando.BoolsN(18)),
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
		One:   [19]bool(rando.BoolsN(19)),
		Two:   [20]bool(rando.BoolsN(20)),
		Three: [21]bool(rando.BoolsN(21)),
		Four:  [22]bool(rando.BoolsN(22)),
		Five:  [23]bool(rando.BoolsN(23)),
		Six:   [24]bool(rando.BoolsN(24)),
		Seven: [25]bool(rando.BoolsN(25)),
		Eight: [26]bool(rando.BoolsN(26)),
		Nine:  [27]bool(rando.BoolsN(27)),
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
		One:   [28]bool(rando.BoolsN(28)),
		Two:   [29]bool(rando.BoolsN(29)),
		Three: [30]bool(rando.BoolsN(30)),
		Four:  [31]bool(rando.BoolsN(31)),
		Five:  [32]bool(rando.BoolsN(32)),
		Six:   [33]bool(rando.BoolsN(33)),
		Seven: [34]bool(rando.BoolsN(34)),
		Eight: [35]bool(rando.BoolsN(35)),
		Nine:  [36]bool(rando.BoolsN(36)),
		Ten:   [37]bool(rando.BoolsN(37)),
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
		One:    [38]bool(rando.BoolsN(38)),
		Two:    [39]bool(rando.BoolsN(39)),
		Three:  [40]bool(rando.BoolsN(40)),
		Four:   [41]bool(rando.BoolsN(41)),
		Five:   [42]bool(rando.BoolsN(42)),
		Six:    [43]bool(rando.BoolsN(43)),
		Seven:  [44]bool(rando.BoolsN(44)),
		Eight:  [45]bool(rando.BoolsN(45)),
		Nine:   [46]bool(rando.BoolsN(46)),
		Ten:    [47]bool(rando.BoolsN(47)),
		Eleven: [48]bool(rando.BoolsN(48)),
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
		One:    [49]bool(rando.BoolsN(49)),
		Two:    [50]bool(rando.BoolsN(50)),
		Three:  [51]bool(rando.BoolsN(51)),
		Four:   [52]bool(rando.BoolsN(52)),
		Five:   [53]bool(rando.BoolsN(53)),
		Six:    [54]bool(rando.BoolsN(54)),
		Seven:  [55]bool(rando.BoolsN(55)),
		Eight:  [56]bool(rando.BoolsN(56)),
		Nine:   [57]bool(rando.BoolsN(57)),
		Ten:    [58]bool(rando.BoolsN(58)),
		Eleven: [59]bool(rando.BoolsN(59)),
		Twelve: [60]bool(rando.BoolsN(60)),
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
		One:      [61]bool(rando.BoolsN(61)),
		Two:      [62]bool(rando.BoolsN(62)),
		Three:    [63]bool(rando.BoolsN(63)),
		Four:     [64]bool(rando.BoolsN(64)),
		Five:     [65]bool(rando.BoolsN(65)),
		Six:      [66]bool(rando.BoolsN(66)),
		Seven:    [67]bool(rando.BoolsN(67)),
		Eight:    [68]bool(rando.BoolsN(68)),
		Nine:     [69]bool(rando.BoolsN(69)),
		Ten:      [70]bool(rando.BoolsN(70)),
		Eleven:   [71]bool(rando.BoolsN(71)),
		Twelve:   [72]bool(rando.BoolsN(72)),
		Thirteen: [73]bool(rando.BoolsN(73)),
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
		One:      [74]bool(rando.BoolsN(74)),
		Two:      [75]bool(rando.BoolsN(75)),
		Three:    [76]bool(rando.BoolsN(76)),
		Four:     [77]bool(rando.BoolsN(77)),
		Five:     [78]bool(rando.BoolsN(78)),
		Six:      [79]bool(rando.BoolsN(79)),
		Seven:    [80]bool(rando.BoolsN(80)),
		Eight:    [81]bool(rando.BoolsN(81)),
		Nine:     [82]bool(rando.BoolsN(82)),
		Ten:      [83]bool(rando.BoolsN(83)),
		Eleven:   [84]bool(rando.BoolsN(84)),
		Twelve:   [85]bool(rando.BoolsN(85)),
		Thirteen: [86]bool(rando.BoolsN(86)),
		Fourteen: [87]bool(rando.BoolsN(87)),
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
		One:      [88]bool(rando.BoolsN(88)),
		Two:      [89]bool(rando.BoolsN(89)),
		Three:    [90]bool(rando.BoolsN(90)),
		Four:     [91]bool(rando.BoolsN(91)),
		Five:     [92]bool(rando.BoolsN(92)),
		Six:      [93]bool(rando.BoolsN(93)),
		Seven:    [94]bool(rando.BoolsN(94)),
		Eight:    [95]bool(rando.BoolsN(95)),
		Nine:     [96]bool(rando.BoolsN(96)),
		Ten:      [97]bool(rando.BoolsN(97)),
		Eleven:   [98]bool(rando.BoolsN(98)),
		Twelve:   [99]bool(rando.BoolsN(99)),
		Thirteen: [100]bool(rando.BoolsN(100)),
		Fourteen: [101]bool(rando.BoolsN(101)),
		Fifteen:  [102]bool(rando.BoolsN(102)),
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
		One:      [103]bool(rando.BoolsN(103)),
		Two:      [104]bool(rando.BoolsN(104)),
		Three:    [105]bool(rando.BoolsN(105)),
		Four:     [106]bool(rando.BoolsN(106)),
		Five:     [107]bool(rando.BoolsN(107)),
		Six:      [108]bool(rando.BoolsN(108)),
		Seven:    [109]bool(rando.BoolsN(109)),
		Eight:    [110]bool(rando.BoolsN(110)),
		Nine:     [111]bool(rando.BoolsN(111)),
		Ten:      [112]bool(rando.BoolsN(112)),
		Eleven:   [113]bool(rando.BoolsN(113)),
		Twelve:   [114]bool(rando.BoolsN(114)),
		Thirteen: [115]bool(rando.BoolsN(115)),
		Fourteen: [116]bool(rando.BoolsN(116)),
		Fifteen:  [117]bool(rando.BoolsN(117)),
		Sixteen:  [118]bool(rando.BoolsN(118)),
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
		One:       [119]bool(rando.BoolsN(119)),
		Two:       [120]bool(rando.BoolsN(120)),
		Three:     [121]bool(rando.BoolsN(121)),
		Four:      [122]bool(rando.BoolsN(122)),
		Five:      [123]bool(rando.BoolsN(123)),
		Six:       [124]bool(rando.BoolsN(124)),
		Seven:     [125]bool(rando.BoolsN(125)),
		Eight:     [126]bool(rando.BoolsN(126)),
		Nine:      [127]bool(rando.BoolsN(127)),
		Ten:       [128]bool(rando.BoolsN(128)),
		Eleven:    [129]bool(rando.BoolsN(129)),
		Twelve:    [130]bool(rando.BoolsN(130)),
		Thirteen:  [131]bool(rando.BoolsN(131)),
		Fourteen:  [132]bool(rando.BoolsN(132)),
		Fifteen:   [133]bool(rando.BoolsN(133)),
		Sixteen:   [134]bool(rando.BoolsN(134)),
		Seventeen: [135]bool(rando.BoolsN(135)),
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
		One:       [136]bool(rando.BoolsN(136)),
		Two:       [137]bool(rando.BoolsN(137)),
		Three:     [138]bool(rando.BoolsN(138)),
		Four:      [139]bool(rando.BoolsN(139)),
		Five:      [140]bool(rando.BoolsN(140)),
		Six:       [141]bool(rando.BoolsN(141)),
		Seven:     [142]bool(rando.BoolsN(142)),
		Eight:     [143]bool(rando.BoolsN(143)),
		Nine:      [144]bool(rando.BoolsN(144)),
		Ten:       [145]bool(rando.BoolsN(145)),
		Eleven:    [146]bool(rando.BoolsN(146)),
		Twelve:    [147]bool(rando.BoolsN(147)),
		Thirteen:  [148]bool(rando.BoolsN(148)),
		Fourteen:  [149]bool(rando.BoolsN(149)),
		Fifteen:   [150]bool(rando.BoolsN(150)),
		Sixteen:   [151]bool(rando.BoolsN(151)),
		Seventeen: [152]bool(rando.BoolsN(152)),
		Eighteen:  [153]bool(rando.BoolsN(153)),
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
		One:       [154]bool(rando.BoolsN(154)),
		Two:       [155]bool(rando.BoolsN(155)),
		Three:     [156]bool(rando.BoolsN(156)),
		Four:      [157]bool(rando.BoolsN(157)),
		Five:      [158]bool(rando.BoolsN(158)),
		Six:       [159]bool(rando.BoolsN(159)),
		Seven:     [160]bool(rando.BoolsN(160)),
		Eight:     [161]bool(rando.BoolsN(161)),
		Nine:      [162]bool(rando.BoolsN(162)),
		Ten:       [163]bool(rando.BoolsN(163)),
		Eleven:    [164]bool(rando.BoolsN(164)),
		Twelve:    [165]bool(rando.BoolsN(165)),
		Thirteen:  [166]bool(rando.BoolsN(166)),
		Fourteen:  [167]bool(rando.BoolsN(167)),
		Fifteen:   [168]bool(rando.BoolsN(168)),
		Sixteen:   [169]bool(rando.BoolsN(169)),
		Seventeen: [170]bool(rando.BoolsN(170)),
		Eighteen:  [171]bool(rando.BoolsN(171)),
		Nineteen:  [172]bool(rando.BoolsN(172)),
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
		One:       [173]bool(rando.BoolsN(173)),
		Two:       [174]bool(rando.BoolsN(174)),
		Three:     [175]bool(rando.BoolsN(175)),
		Four:      [176]bool(rando.BoolsN(176)),
		Five:      [177]bool(rando.BoolsN(177)),
		Six:       [178]bool(rando.BoolsN(178)),
		Seven:     [179]bool(rando.BoolsN(179)),
		Eight:     [180]bool(rando.BoolsN(180)),
		Nine:      [181]bool(rando.BoolsN(181)),
		Ten:       [182]bool(rando.BoolsN(182)),
		Eleven:    [183]bool(rando.BoolsN(183)),
		Twelve:    [184]bool(rando.BoolsN(184)),
		Thirteen:  [185]bool(rando.BoolsN(185)),
		Fourteen:  [186]bool(rando.BoolsN(186)),
		Fifteen:   [187]bool(rando.BoolsN(187)),
		Sixteen:   [188]bool(rando.BoolsN(188)),
		Seventeen: [189]bool(rando.BoolsN(189)),
		Eighteen:  [190]bool(rando.BoolsN(190)),
		Nineteen:  [191]bool(rando.BoolsN(191)),
		Twenty:    [192]bool(rando.BoolsN(192)),
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
		One:       [193]bool(rando.BoolsN(193)),
		Two:       [194]bool(rando.BoolsN(194)),
		Three:     [195]bool(rando.BoolsN(195)),
		Four:      [196]bool(rando.BoolsN(196)),
		Five:      [197]bool(rando.BoolsN(197)),
		Six:       [198]bool(rando.BoolsN(198)),
		Seven:     [199]bool(rando.BoolsN(199)),
		Eight:     [200]bool(rando.BoolsN(200)),
		Nine:      [201]bool(rando.BoolsN(201)),
		Ten:       [202]bool(rando.BoolsN(202)),
		Eleven:    [203]bool(rando.BoolsN(203)),
		Twelve:    [204]bool(rando.BoolsN(204)),
		Thirteen:  [205]bool(rando.BoolsN(205)),
		Fourteen:  [206]bool(rando.BoolsN(206)),
		Fifteen:   [207]bool(rando.BoolsN(207)),
		Sixteen:   [208]bool(rando.BoolsN(208)),
		Seventeen: [209]bool(rando.BoolsN(209)),
		Eighteen:  [210]bool(rando.BoolsN(210)),
		Nineteen:  [211]bool(rando.BoolsN(211)),
		Twenty:    [212]bool(rando.BoolsN(212)),
		TwentyOne: [213]bool(rando.BoolsN(213)),
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
		One:       [214]bool(rando.BoolsN(214)),
		Two:       [215]bool(rando.BoolsN(215)),
		Three:     [216]bool(rando.BoolsN(216)),
		Four:      [217]bool(rando.BoolsN(217)),
		Five:      [218]bool(rando.BoolsN(218)),
		Six:       [219]bool(rando.BoolsN(219)),
		Seven:     [220]bool(rando.BoolsN(220)),
		Eight:     [221]bool(rando.BoolsN(221)),
		Nine:      [222]bool(rando.BoolsN(222)),
		Ten:       [223]bool(rando.BoolsN(223)),
		Eleven:    [224]bool(rando.BoolsN(224)),
		Twelve:    [225]bool(rando.BoolsN(225)),
		Thirteen:  [226]bool(rando.BoolsN(226)),
		Fourteen:  [227]bool(rando.BoolsN(227)),
		Fifteen:   [228]bool(rando.BoolsN(228)),
		Sixteen:   [229]bool(rando.BoolsN(229)),
		Seventeen: [230]bool(rando.BoolsN(230)),
		Eighteen:  [231]bool(rando.BoolsN(231)),
		Nineteen:  [232]bool(rando.BoolsN(232)),
		Twenty:    [233]bool(rando.BoolsN(233)),
		TwentyOne: [234]bool(rando.BoolsN(234)),
		TwentyTwo: [235]bool(rando.BoolsN(235)),
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
		One:         [236]bool(rando.BoolsN(236)),
		Two:         [237]bool(rando.BoolsN(237)),
		Three:       [238]bool(rando.BoolsN(238)),
		Four:        [239]bool(rando.BoolsN(239)),
		Five:        [240]bool(rando.BoolsN(240)),
		Six:         [241]bool(rando.BoolsN(241)),
		Seven:       [242]bool(rando.BoolsN(242)),
		Eight:       [243]bool(rando.BoolsN(243)),
		Nine:        [244]bool(rando.BoolsN(244)),
		Ten:         [245]bool(rando.BoolsN(245)),
		Eleven:      [246]bool(rando.BoolsN(246)),
		Twelve:      [247]bool(rando.BoolsN(247)),
		Thirteen:    [248]bool(rando.BoolsN(248)),
		Fourteen:    [249]bool(rando.BoolsN(249)),
		Fifteen:     [250]bool(rando.BoolsN(250)),
		Sixteen:     [251]bool(rando.BoolsN(251)),
		Seventeen:   [252]bool(rando.BoolsN(252)),
		Eighteen:    [253]bool(rando.BoolsN(253)),
		Nineteen:    [254]bool(rando.BoolsN(254)),
		Twenty:      [255]bool(rando.BoolsN(255)),
		TwentyOne:   [256]bool(BoolsN(256)),
		TwentyTwo:   [257]bool(BoolsN(257)),
		TwentyThree: [258]bool(BoolsN(258)),
	}
	src := expected.MarshalJ()
	require.NoError(t, actual.UnmarshalJ(src))
	// require.NotEqual(t, TwentyThree{}, expected)
	// require.NotEqual(t, TwentyThree{}, actual)
	require.Equal(t, expected, actual)
}

// BoolsN returns a []bool with length `size`, populated with random values.
func BoolsN(size uint) (b []bool) {
	if size >= 1 {
		b = make([]bool, size)
		for i := range b {
			b[i] = rando.Bool()
		}
	}

	return
}
