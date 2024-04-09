package jay_test

import (
	"fmt"
	"github.com/speedyhoon/jay"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWriteBools(t *testing.T) {
	tests := []struct {
		y    []byte
		a    []bool
		want []byte
	}{
		{y: nil, a: nil, want: nil},
		{y: []byte{0}, a: nil, want: []byte{0}},
		{y: []byte{0}, a: []bool{false}, want: []byte{0}},
		{y: []byte{0}, a: []bool{true}, want: []byte{128}},
		{y: []byte{0}, a: []bool{true, false}, want: []byte{128}},
		{y: []byte{0}, a: []bool{true, true}, want: []byte{192}},
		{y: []byte{0}, a: []bool{true, true, false}, want: []byte{192}},
		{y: []byte{0}, a: []bool{true, true, true}, want: []byte{224}},
		{y: []byte{0}, a: []bool{true, true, true, true}, want: []byte{240}},
		{y: []byte{0}, a: []bool{true, true, true, true, true, true, true, true}, want: []byte{255}},
		{y: []byte{0, 0}, a: []bool{
			true, true, true, true, true, true, true, true,
			true,
		}, want: []byte{255, 128}},
		{y: []byte{0, 0}, a: []bool{
			true, true, true, true, true, true, true, true,
			true, true, true, true, true, true, true, true,
		}, want: []byte{255, 255}},
		{y: []byte{0, 0, 0}, a: []bool{
			true, true, true, true, true, true, true, true,
			true, true, true, true, true, true, true, true,
			true,
		}, want: []byte{255, 255, 128}},
		{y: []byte{0, 0, 0}, a: []bool{
			true, true, true, true, true, true, true, true,
			true, true, true, true, true, true, true, true,
			true, true, true, true, true, true, true, true,
		}, want: []byte{255, 255, 255}},
		{y: []byte{0, 0, 0, 0}, a: []bool{
			true, true, true, true, true, true, true, true,
			true, true, true, true, true, true, true, true,
			true, true, true, true, true, true, true, true,
			true,
		}, want: []byte{255, 255, 255, 128}},
		{y: []byte{0, 0, 0, 0}, a: []bool{
			true, true, true, true, true, true, true, true,
			true, true, true, true, true, true, true, true,
			true, true, true, true, true, true, true, true,
			true, true, true, true, true, true, true, true,
		}, want: []byte{255, 255, 255, 255}},

		{y: []byte{0, 0}, a: []bool{
			false, false, false, false, false, false, false, false,
			false,
		}, want: []byte{0, 0}},
		{y: []byte{0, 0}, a: []bool{
			false, false, false, false, false, false, false, false,
			false, false, false, false, false, false, false, false,
		}, want: []byte{0, 0}},
		{y: []byte{0, 0, 0}, a: []bool{
			false, false, false, false, false, false, false, false,
			false, false, false, false, false, false, false, false,
			false,
		}, want: []byte{0, 0, 0}},
		{y: []byte{0, 0, 0}, a: []bool{
			false, false, false, false, false, false, false, false,
			false, false, false, false, false, false, false, false,
			false, false, false, false, false, false, false, false,
		}, want: []byte{0, 0, 0}},
		{y: []byte{0, 0, 0, 0}, a: []bool{
			false, false, false, false, false, false, false, false,
			false, false, false, false, false, false, false, false,
			false, false, false, false, false, false, false, false,
			false,
		}, want: []byte{0, 0, 0, 0}},
		{y: []byte{0, 0, 0, 0}, a: []bool{
			false, false, false, false, false, false, false, false,
			false, false, false, false, false, false, false, false,
			false, false, false, false, false, false, false, false,
			false, false, false, false, false, false, false, false,
		}, want: []byte{0, 0, 0, 0}},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("test: %d", i+1), func(t *testing.T) {
			jay.WriteBools(tt.y, tt.a, len(tt.a))
			assert.Equal(t, tt.want, tt.y)
		})
	}
}

func TestReadBools(t *testing.T) {
	tests := []struct {
		y    []byte
		want []bool
	}{
		{y: nil, want: nil},
		{y: []byte{0}, want: nil},
		{y: []byte{0}, want: []bool{false}},
		{y: []byte{128}, want: []bool{true}},
		{y: []byte{128}, want: []bool{true, false}},
		{y: []byte{192}, want: []bool{true, true}},
		{y: []byte{192}, want: []bool{true, true, false}},
		{y: []byte{224}, want: []bool{true, true, true}},
		{y: []byte{240}, want: []bool{true, true, true, true}},
		{y: []byte{255}, want: []bool{true, true, true, true, true, true, true, true}},
		{y: []byte{255, 128}, want: []bool{
			true, true, true, true, true, true, true, true,
			true,
		}},
		{y: []byte{255, 255}, want: []bool{
			true, true, true, true, true, true, true, true,
			true, true, true, true, true, true, true, true,
		}},
		{y: []byte{255, 255, 128}, want: []bool{
			true, true, true, true, true, true, true, true,
			true, true, true, true, true, true, true, true,
			true,
		}},
		{y: []byte{255, 255, 255}, want: []bool{
			true, true, true, true, true, true, true, true,
			true, true, true, true, true, true, true, true,
			true, true, true, true, true, true, true, true,
		}},
		{y: []byte{255, 255, 255, 128}, want: []bool{
			true, true, true, true, true, true, true, true,
			true, true, true, true, true, true, true, true,
			true, true, true, true, true, true, true, true,
			true,
		}},
		{y: []byte{255, 255, 255, 255}, want: []bool{
			true, true, true, true, true, true, true, true,
			true, true, true, true, true, true, true, true,
			true, true, true, true, true, true, true, true,
			true, true, true, true, true, true, true, true,
		}},

		{y: []byte{0, 0}, want: []bool{
			false, false, false, false, false, false, false, false,
			false,
		}},
		{y: []byte{0, 0}, want: []bool{
			false, false, false, false, false, false, false, false,
			false, false, false, false, false, false, false, false,
		}},
		{y: []byte{0, 0, 0}, want: []bool{
			false, false, false, false, false, false, false, false,
			false, false, false, false, false, false, false, false,
			false,
		}},
		{y: []byte{0, 0, 0}, want: []bool{
			false, false, false, false, false, false, false, false,
			false, false, false, false, false, false, false, false,
			false, false, false, false, false, false, false, false,
		}},
		{y: []byte{0, 0, 0, 0}, want: []bool{
			false, false, false, false, false, false, false, false,
			false, false, false, false, false, false, false, false,
			false, false, false, false, false, false, false, false,
			false,
		}},
		{y: []byte{0, 0, 0, 0}, want: []bool{
			false, false, false, false, false, false, false, false,
			false, false, false, false, false, false, false, false,
			false, false, false, false, false, false, false, false,
			false, false, false, false, false, false, false, false,
		}},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("test: %d", i+1), func(t *testing.T) {
			assert.Equal(t, tt.want, jay.ReadBools(tt.y, len(tt.want)))
		})
	}
}

func TestSizeBools(t *testing.T) {
	tests := map[int]int{
		-249: 0, -250: 0, -251: 0, -252: 0, -253: 0, -254: 0, -255: 0, -256: 0,
		-241: 0, -242: 0, -243: 0, -244: 0, -245: 0, -246: 0, -247: 0, -248: 0,
		-233: 0, -234: 0, -235: 0, -236: 0, -237: 0, -238: 0, -239: 0, -240: 0,
		-225: 0, -226: 0, -227: 0, -228: 0, -229: 0, -230: 0, -231: 0, -232: 0,
		-217: 0, -218: 0, -219: 0, -220: 0, -221: 0, -222: 0, -223: 0, -224: 0,
		-209: 0, -210: 0, -211: 0, -212: 0, -213: 0, -214: 0, -215: 0, -216: 0,
		-201: 0, -202: 0, -203: 0, -204: 0, -205: 0, -206: 0, -207: 0, -208: 0,
		-193: 0, -194: 0, -195: 0, -196: 0, -197: 0, -198: 0, -199: 0, -200: 0,
		-185: 0, -186: 0, -187: 0, -188: 0, -189: 0, -190: 0, -191: 0, -192: 0,
		-177: 0, -178: 0, -179: 0, -180: 0, -181: 0, -182: 0, -183: 0, -184: 0,
		-169: 0, -170: 0, -171: 0, -172: 0, -173: 0, -174: 0, -175: 0, -176: 0,
		-161: 0, -162: 0, -163: 0, -164: 0, -165: 0, -166: 0, -167: 0, -168: 0,
		-153: 0, -154: 0, -155: 0, -156: 0, -157: 0, -158: 0, -159: 0, -160: 0,
		-145: 0, -146: 0, -147: 0, -148: 0, -149: 0, -150: 0, -151: 0, -152: 0,
		-137: 0, -138: 0, -139: 0, -140: 0, -141: 0, -142: 0, -143: 0, -144: 0,
		-129: 0, -130: 0, -131: 0, -132: 0, -133: 0, -134: 0, -135: 0, -136: 0,
		-121: 0, -122: 0, -123: 0, -124: 0, -125: 0, -126: 0, -127: 0, -128: 0,
		-113: 0, -114: 0, -115: 0, -116: 0, -117: 0, -118: 0, -119: 0, -120: 0,
		-105: 0, -106: 0, -107: 0, -108: 0, -109: 0, -110: 0, -111: 0, -112: 0,
		-97: 0, -98: 0, -99: 0, -100: 0, -101: 0, -102: 0, -103: 0, -104: 0,
		-89: 0, -90: 0, -91: 0, -92: 0, -93: 0, -94: 0, -95: 0, -96: 0,
		-81: 0, -82: 0, -83: 0, -84: 0, -85: 0, -86: 0, -87: 0, -88: 0,
		-73: 0, -74: 0, -75: 0, -76: 0, -77: 0, -78: 0, -79: 0, -80: 0,
		-65: 0, -66: 0, -67: 0, -68: 0, -69: 0, -70: 0, -71: 0, -72: 0,
		-57: 0, -58: 0, -59: 0, -60: 0, -61: 0, -62: 0, -63: 0, -64: 0,
		-49: 0, -50: 0, -51: 0, -52: 0, -53: 0, -54: 0, -55: 0, -56: 0,
		-41: 0, -42: 0, -43: 0, -44: 0, -45: 0, -46: 0, -47: 0, -48: 0,
		-33: 0, -34: 0, -35: 0, -36: 0, -37: 0, -38: 0, -39: 0, -40: 0,
		-25: 0, -26: 0, -27: 0, -28: 0, -29: 0, -30: 0, -31: 0, -32: 0,
		-17: 0, -18: 0, -19: 0, -20: 0, -21: 0, -22: 0, -23: 0, -24: 0,
		-9: 0, -10: 0, -11: 0, -12: 0, -13: 0, -14: 0, -15: 0, -16: 0,
		-1: 0, -2: 0, -3: 0, -4: 0, -5: 0, -6: 0, -7: 0, -8: 0,
		0: 0,
		1: 1, 2: 1, 3: 1, 4: 1, 5: 1, 6: 1, 7: 1, 8: 1,
		9: 2, 10: 2, 11: 2, 12: 2, 13: 2, 14: 2, 15: 2, 16: 2,
		17: 3, 18: 3, 19: 3, 20: 3, 21: 3, 22: 3, 23: 3, 24: 3,
		25: 4, 26: 4, 27: 4, 28: 4, 29: 4, 30: 4, 31: 4, 32: 4,
		33: 5, 34: 5, 35: 5, 36: 5, 37: 5, 38: 5, 39: 5, 40: 5,
		41: 6, 42: 6, 43: 6, 44: 6, 45: 6, 46: 6, 47: 6, 48: 6,
		49: 7, 50: 7, 51: 7, 52: 7, 53: 7, 54: 7, 55: 7, 56: 7,
		57: 8, 58: 8, 59: 8, 60: 8, 61: 8, 62: 8, 63: 8, 64: 8,
		65: 9, 66: 9, 67: 9, 68: 9, 69: 9, 70: 9, 71: 9, 72: 9,
		73: 10, 74: 10, 75: 10, 76: 10, 77: 10, 78: 10, 79: 10, 80: 10,
		81: 11, 82: 11, 83: 11, 84: 11, 85: 11, 86: 11, 87: 11, 88: 11,
		89: 12, 90: 12, 91: 12, 92: 12, 93: 12, 94: 12, 95: 12, 96: 12,
		97: 13, 98: 13, 99: 13, 100: 13, 101: 13, 102: 13, 103: 13, 104: 13,
		105: 14, 106: 14, 107: 14, 108: 14, 109: 14, 110: 14, 111: 14, 112: 14,
		113: 15, 114: 15, 115: 15, 116: 15, 117: 15, 118: 15, 119: 15, 120: 15,
		121: 16, 122: 16, 123: 16, 124: 16, 125: 16, 126: 16, 127: 16, 128: 16,
		129: 17, 130: 17, 131: 17, 132: 17, 133: 17, 134: 17, 135: 17, 136: 17,
		137: 18, 138: 18, 139: 18, 140: 18, 141: 18, 142: 18, 143: 18, 144: 18,
		145: 19, 146: 19, 147: 19, 148: 19, 149: 19, 150: 19, 151: 19, 152: 19,
		153: 20, 154: 20, 155: 20, 156: 20, 157: 20, 158: 20, 159: 20, 160: 20,
		161: 21, 162: 21, 163: 21, 164: 21, 165: 21, 166: 21, 167: 21, 168: 21,
		169: 22, 170: 22, 171: 22, 172: 22, 173: 22, 174: 22, 175: 22, 176: 22,
		177: 23, 178: 23, 179: 23, 180: 23, 181: 23, 182: 23, 183: 23, 184: 23,
		185: 24, 186: 24, 187: 24, 188: 24, 189: 24, 190: 24, 191: 24, 192: 24,
		193: 25, 194: 25, 195: 25, 196: 25, 197: 25, 198: 25, 199: 25, 200: 25,
		201: 26, 202: 26, 203: 26, 204: 26, 205: 26, 206: 26, 207: 26, 208: 26,
		209: 27, 210: 27, 211: 27, 212: 27, 213: 27, 214: 27, 215: 27, 216: 27,
		217: 28, 218: 28, 219: 28, 220: 28, 221: 28, 222: 28, 223: 28, 224: 28,
		225: 29, 226: 29, 227: 29, 228: 29, 229: 29, 230: 29, 231: 29, 232: 29,
		233: 30, 234: 30, 235: 30, 236: 30, 237: 30, 238: 30, 239: 30, 240: 30,
		241: 31, 242: 31, 243: 31, 244: 31, 245: 31, 246: 31, 247: 31, 248: 31,
		249: 32, 250: 32, 251: 32, 252: 32, 253: 32, 254: 32, 255: 32, 256: 32,
	}
	for input, expected := range tests {
		assert.Equal(t, expected, jay.SizeBools(input), "input: %d", input)
	}
}
