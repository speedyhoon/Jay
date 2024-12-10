package fuzz05

import (
	"testing"

	"github.com/speedyhoon/rando"
	"github.com/stretchr/testify/require"
)

func TestFuzz_0(t *testing.T) {
	var expected, actual S13Y8PdPX74Y7b
	require.NoError(t, expected.UnmarshalJ(actual.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, S13Y8PdPX74Y7b{}, expected)
	require.Equal(t, S13Y8PdPX74Y7b{}, actual)

	actual = S13Y8PdPX74Y7b{
		P7m4M35D81iVo4Jui88pweX: rando.Float32(),
		A7v67va1bc3o5mmXP7Q1d2YfeHHEpf482P6UcyWnRO0C675gxM5HeS7Wevi1YHeL3f42pluc5: rando.Uint(),
		Y3N: rando.Uint32(),
		E057vAuD70jsTdYLnj0waaVewYvy7y3d0UYAuJx7B2D80aq58WQQNjnA0C6811qIGl2VMr2p0400u6TQdG8e5Rwmnpu: rando.Float32(),
	}
	src := actual.MarshalJ()
	require.NoError(t, expected.UnmarshalJ(src))
	require.NotEqual(t, S13Y8PdPX74Y7b{}, expected)
	require.NotEqual(t, S13Y8PdPX74Y7b{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_1(t *testing.T) {
	var expected, actual XVSlCJMQLIo803Uwv4PYS033cc3la2yP84qq0P8aE6xhh7
	require.NoError(t, expected.UnmarshalJ(actual.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, XVSlCJMQLIo803Uwv4PYS033cc3la2yP84qq0P8aE6xhh7{}, expected)
	require.Equal(t, XVSlCJMQLIo803Uwv4PYS033cc3la2yP84qq0P8aE6xhh7{}, actual)

	actual = XVSlCJMQLIo803Uwv4PYS033cc3la2yP84qq0P8aE6xhh7{
		V5N46HInloeD0gsxmYCyP1fi32NfRMrKJp: rando.String(),
	}
	src := actual.MarshalJ()
	require.NoError(t, expected.UnmarshalJ(src))
	require.NotEqual(t, XVSlCJMQLIo803Uwv4PYS033cc3la2yP84qq0P8aE6xhh7{}, expected)
	require.NotEqual(t, XVSlCJMQLIo803Uwv4PYS033cc3la2yP84qq0P8aE6xhh7{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_2(t *testing.T) {
	var expected, actual F0bUdl0FVVI4RWe05JFPQYFeIpJ4YKY782t0Q718S2EG7543XUicNL6O03FP3S2
	require.NoError(t, expected.UnmarshalJ(actual.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, F0bUdl0FVVI4RWe05JFPQYFeIpJ4YKY782t0Q718S2EG7543XUicNL6O03FP3S2{}, expected)
	require.Equal(t, F0bUdl0FVVI4RWe05JFPQYFeIpJ4YKY782t0Q718S2EG7543XUicNL6O03FP3S2{}, actual)

	actual = F0bUdl0FVVI4RWe05JFPQYFeIpJ4YKY782t0Q718S2EG7543XUicNL6O03FP3S2{
		CW6f57keEQY07siB4JTgCAUjF27KD: rando.Int32(),
		Iot623153jl:                   rando.Int16(),
		JbnFxahdtkWJI3P0ckt068F6O5Wchnt1S41FSphGux4G1f6oJUrI7K2SLd3cp3Ifm: rando.Int(),
		HtEyGS5JY3LKv7dJ2p1K718K60GUPc4oI11Y46Al5hXP52PQFe2T3:             rando.Int8(),
		X403a1df61cro0c5K487mEQk18RvWlvakoc31De:                           rando.Uint64(),
		FNuylRn6CWFCoE850qqRsWAFrbUuU6o1tG57O7CS42:                        rando.Float32(),
		KuP5Y5Nuqwn: rando.Uint16(),
		Do87KSbiii82NO10JyMHf1U4nL66DN3vx0d6u0Kao0rtN53ftCkA025: rando.Int64(),
	}
	src := actual.MarshalJ()
	require.NoError(t, expected.UnmarshalJ(src))
	require.NotEqual(t, F0bUdl0FVVI4RWe05JFPQYFeIpJ4YKY782t0Q718S2EG7543XUicNL6O03FP3S2{}, expected)
	require.NotEqual(t, F0bUdl0FVVI4RWe05JFPQYFeIpJ4YKY782t0Q718S2EG7543XUicNL6O03FP3S2{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_3(t *testing.T) {
	var expected, actual Bt70X1y6cxJ6p4P8mmN4cTwiQM67VI65Vx5WCgla46Haf5m88maEiQ68c0s6GReiyW8oO08TFPyVYAAfd73O3TibG47Gp34q6508
	require.NoError(t, expected.UnmarshalJ(actual.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, Bt70X1y6cxJ6p4P8mmN4cTwiQM67VI65Vx5WCgla46Haf5m88maEiQ68c0s6GReiyW8oO08TFPyVYAAfd73O3TibG47Gp34q6508{}, expected)
	require.Equal(t, Bt70X1y6cxJ6p4P8mmN4cTwiQM67VI65Vx5WCgla46Haf5m88maEiQ68c0s6GReiyW8oO08TFPyVYAAfd73O3TibG47Gp34q6508{}, actual)

	actual = Bt70X1y6cxJ6p4P8mmN4cTwiQM67VI65Vx5WCgla46Haf5m88maEiQ68c0s6GReiyW8oO08TFPyVYAAfd73O3TibG47Gp34q6508{
		Tr4CRhFyrQU131P5c54NvHCsJAobs41v2e8u7J: rando.Int(),
	}
	src := actual.MarshalJ()
	require.NoError(t, expected.UnmarshalJ(src))
	require.NotEqual(t, Bt70X1y6cxJ6p4P8mmN4cTwiQM67VI65Vx5WCgla46Haf5m88maEiQ68c0s6GReiyW8oO08TFPyVYAAfd73O3TibG47Gp34q6508{}, expected)
	require.NotEqual(t, Bt70X1y6cxJ6p4P8mmN4cTwiQM67VI65Vx5WCgla46Haf5m88maEiQ68c0s6GReiyW8oO08TFPyVYAAfd73O3TibG47Gp34q6508{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_4(t *testing.T) {
	var expected, actual Y0m4GH5J1b3Pku55C03L4p17aLoBhX4WnF7QaPO1bqgum5X
	require.NoError(t, expected.UnmarshalJ(actual.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, Y0m4GH5J1b3Pku55C03L4p17aLoBhX4WnF7QaPO1bqgum5X{}, expected)
	require.Equal(t, Y0m4GH5J1b3Pku55C03L4p17aLoBhX4WnF7QaPO1bqgum5X{}, actual)

	actual = Y0m4GH5J1b3Pku55C03L4p17aLoBhX4WnF7QaPO1bqgum5X{
		CaSAAqJ5d05dd7lTSBW8P74ro8Bu: rando.Uint64(),
		VY123:                        rando.Time(),
		IC5qlktHjoAxQ85BBL11a52LLqLNa03GvsFtx6660G0Bj78LXptHo40S737W4ro0Y27s10168Xc75kpo7: rando.Bool(),
		Dxe5ORTPGA6AHajb0OQV5r2wl2lE2p6LCn85a23ysejS30:                                    rando.Byte(),
		N3L7V2OMyF3K5LqS0lRxe6IcMQPh5:                                                     rando.Time(),
		I158Luik4h616xv4cL1x1hEX082jEu42cSAbXfr0MU0phAre88mfxEkm4lr6p6r7j:                 rando.Duration(),
		C11nVJCwQXED6cgc65lS0NNld6ormyMtGY42JNPP34hOA3GegJ10ObN56e:                        rando.Bool(),
		WkMdF1S4rr0dYq0SN4TTo3h2:                                                          rando.Bools(),
	}
	src := actual.MarshalJ()
	require.NoError(t, expected.UnmarshalJ(src))
	require.NotEqual(t, Y0m4GH5J1b3Pku55C03L4p17aLoBhX4WnF7QaPO1bqgum5X{}, expected)
	require.NotEqual(t, Y0m4GH5J1b3Pku55C03L4p17aLoBhX4WnF7QaPO1bqgum5X{}, actual)
	require.Equal(t, expected, actual)
}
