package byte

import "github.com/speedyhoon/jay"

type MakeVsReturn struct {
	Bool  bool
	Uint8 uint8
	Int8  int8
}

func (m *MakeVsReturn) Make1() (b []byte) {
	b = make([]byte, 1)
	b[0] = jay.Bool1(m.Bool)
	return
}
func (m *MakeVsReturn) Return1() (b []byte) {
	return []byte{jay.Bool1(m.Bool)}
}

func (m *MakeVsReturn) Make2() (b []byte) {
	b = make([]byte, 2)
	b[0] = jay.Bool1(m.Bool)
	b[1] = m.Uint8
	return
}
func (m *MakeVsReturn) Return2() (b []byte) {
	return []byte{jay.Bool1(m.Bool), m.Uint8}
}

func (m *MakeVsReturn) Make3() (b []byte) {
	b = make([]byte, 3)
	b[0] = jay.Bool1(m.Bool)
	b[1] = m.Uint8
	b[2] = byte(m.Int8)
	return
}
func (m *MakeVsReturn) Return3() (b []byte) {
	return []byte{jay.Bool1(m.Bool), m.Uint8, byte(m.Int8)}
}

type MakeVsReturnLarge struct {
	Bool, Hidden, Selected, Deselected, Activated, Modified, Serviced, Performed, Accessed, Updated bool
	Uint80, Uint81, Uint82, Uint83, Uint84, Uint85, Uint86, Uint87, Uint88, Uint89, Uint90          uint8
	Int80, Int81, Int82, Int83, Int84, Int85, Int86, Int87, Int88, Int89, Int90                     int8
}

func (m *MakeVsReturnLarge) Make1() (b []byte) {
	b = make([]byte, 31)
	b[0] = jay.Bool8(m.Bool, m.Hidden, m.Selected, m.Deselected, m.Activated, m.Modified, m.Serviced, m.Performed)
	b[1] = jay.Bool2(m.Accessed, m.Updated)
	b[2] = m.Uint80
	b[3] = m.Uint81
	b[4] = m.Uint82
	b[5] = m.Uint83
	b[6] = m.Uint84
	b[7] = m.Uint85
	b[8] = m.Uint86
	b[9] = m.Uint87
	b[10] = m.Uint88
	b[11] = m.Uint89
	b[12] = m.Uint90
	b[13] = byte(m.Int80)
	b[14] = byte(m.Int81)
	b[15] = byte(m.Int82)
	b[16] = byte(m.Int83)
	b[17] = byte(m.Int84)
	b[18] = byte(m.Int85)
	b[19] = byte(m.Int86)
	b[20] = byte(m.Int87)
	b[21] = byte(m.Int88)
	b[22] = byte(m.Int89)
	b[23] = byte(m.Int90)
	return
}
func (m *MakeVsReturnLarge) Return1() (b []byte) {
	return []byte{
		jay.Bool8(m.Bool, m.Hidden, m.Selected, m.Deselected, m.Activated, m.Modified, m.Serviced, m.Performed),
		jay.Bool2(m.Accessed, m.Updated),
		m.Uint80,
		m.Uint81,
		m.Uint82,
		m.Uint83,
		m.Uint84,
		m.Uint85,
		m.Uint86,
		m.Uint87,
		m.Uint88,
		m.Uint89,
		m.Uint90,
		byte(m.Int80),
		byte(m.Int81),
		byte(m.Int82),
		byte(m.Int83),
		byte(m.Int84),
		byte(m.Int85),
		byte(m.Int86),
		byte(m.Int87),
		byte(m.Int88),
		byte(m.Int89),
		byte(m.Int90),
	}
}
