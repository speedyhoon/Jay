// Code generated by Jay; DO NOT EDIT. Tool documentation available at: https://github.com/speedyhoon/jay

package main

import "github.com/speedyhoon/jay"

func (o *One) MarshalJ() []byte {
	return []byte{byte(o.One)}
}

func (o *One) UnmarshalJ(b []byte) error {
	if len(b) != 1 {
		return jay.ErrUnexpectedEOB
	}
	o.One = tiny(b[0])
	return nil
}

func (t *Two) MarshalJ() []byte {
	return []byte{byte(t.One), byte(t.Two)}
}

func (t *Two) UnmarshalJ(b []byte) error {
	if len(b) != 2 {
		return jay.ErrUnexpectedEOB
	}
	t.One = tiny(b[0])
	t.Two = tiny(b[1])
	return nil
}

func (t *Three) MarshalJ() []byte {
	return []byte{byte(t.One), byte(t.Two), byte(t.Three)}
}

func (t *Three) UnmarshalJ(b []byte) error {
	if len(b) != 3 {
		return jay.ErrUnexpectedEOB
	}
	t.One = tiny(b[0])
	t.Two = tiny(b[1])
	t.Three = tiny(b[2])
	return nil
}

func (f *Four) MarshalJ() []byte {
	return []byte{byte(f.One), byte(f.Two), byte(f.Three), byte(f.Four)}
}

func (f *Four) UnmarshalJ(b []byte) error {
	if len(b) != 4 {
		return jay.ErrUnexpectedEOB
	}
	f.One = tiny(b[0])
	f.Two = tiny(b[1])
	f.Three = tiny(b[2])
	f.Four = tiny(b[3])
	return nil
}

func (f *Five) MarshalJ() []byte {
	return []byte{byte(f.One), byte(f.Two), byte(f.Three), byte(f.Four), byte(f.Five)}
}

func (f *Five) UnmarshalJ(b []byte) error {
	if len(b) != 5 {
		return jay.ErrUnexpectedEOB
	}
	f.One = tiny(b[0])
	f.Two = tiny(b[1])
	f.Three = tiny(b[2])
	f.Four = tiny(b[3])
	f.Five = tiny(b[4])
	return nil
}

func (s *Six) MarshalJ() []byte {
	return []byte{byte(s.One), byte(s.Two), byte(s.Three), byte(s.Four), byte(s.Five), byte(s.Six)}
}

func (s *Six) UnmarshalJ(b []byte) error {
	if len(b) != 6 {
		return jay.ErrUnexpectedEOB
	}
	s.One = tiny(b[0])
	s.Two = tiny(b[1])
	s.Three = tiny(b[2])
	s.Four = tiny(b[3])
	s.Five = tiny(b[4])
	s.Six = tiny(b[5])
	return nil
}

func (s *Seven) MarshalJ() []byte {
	return []byte{byte(s.One), byte(s.Two), byte(s.Three), byte(s.Four), byte(s.Five), byte(s.Six), byte(s.Seven)}
}

func (s *Seven) UnmarshalJ(b []byte) error {
	if len(b) != 7 {
		return jay.ErrUnexpectedEOB
	}
	s.One = tiny(b[0])
	s.Two = tiny(b[1])
	s.Three = tiny(b[2])
	s.Four = tiny(b[3])
	s.Five = tiny(b[4])
	s.Six = tiny(b[5])
	s.Seven = tiny(b[6])
	return nil
}

func (e *Eight) MarshalJ() []byte {
	return []byte{byte(e.One), byte(e.Two), byte(e.Three), byte(e.Four), byte(e.Five), byte(e.Six), byte(e.Seven), byte(e.Eight)}
}

func (e *Eight) UnmarshalJ(b []byte) error {
	if len(b) != 8 {
		return jay.ErrUnexpectedEOB
	}
	e.One = tiny(b[0])
	e.Two = tiny(b[1])
	e.Three = tiny(b[2])
	e.Four = tiny(b[3])
	e.Five = tiny(b[4])
	e.Six = tiny(b[5])
	e.Seven = tiny(b[6])
	e.Eight = tiny(b[7])
	return nil
}

func (n *Nine) MarshalJ() []byte {
	return []byte{byte(n.One), byte(n.Two), byte(n.Three), byte(n.Four), byte(n.Five), byte(n.Six), byte(n.Seven), byte(n.Eight), byte(n.Nine)}
}

func (n *Nine) UnmarshalJ(b []byte) error {
	if len(b) != 9 {
		return jay.ErrUnexpectedEOB
	}
	n.One = tiny(b[0])
	n.Two = tiny(b[1])
	n.Three = tiny(b[2])
	n.Four = tiny(b[3])
	n.Five = tiny(b[4])
	n.Six = tiny(b[5])
	n.Seven = tiny(b[6])
	n.Eight = tiny(b[7])
	n.Nine = tiny(b[8])
	return nil
}

func (t *Ten) MarshalJ() []byte {
	return []byte{byte(t.One), byte(t.Two), byte(t.Three), byte(t.Four), byte(t.Five), byte(t.Six), byte(t.Seven), byte(t.Eight), byte(t.Nine), byte(t.Ten)}
}

func (t *Ten) UnmarshalJ(b []byte) error {
	if len(b) != 10 {
		return jay.ErrUnexpectedEOB
	}
	t.One = tiny(b[0])
	t.Two = tiny(b[1])
	t.Three = tiny(b[2])
	t.Four = tiny(b[3])
	t.Five = tiny(b[4])
	t.Six = tiny(b[5])
	t.Seven = tiny(b[6])
	t.Eight = tiny(b[7])
	t.Nine = tiny(b[8])
	t.Ten = tiny(b[9])
	return nil
}

func (e *Eleven) MarshalJ() []byte {
	return []byte{byte(e.One), byte(e.Two), byte(e.Three), byte(e.Four), byte(e.Five), byte(e.Six), byte(e.Seven), byte(e.Eight), byte(e.Nine), byte(e.Ten), byte(e.Eleven)}
}

func (e *Eleven) UnmarshalJ(b []byte) error {
	if len(b) != 11 {
		return jay.ErrUnexpectedEOB
	}
	e.One = tiny(b[0])
	e.Two = tiny(b[1])
	e.Three = tiny(b[2])
	e.Four = tiny(b[3])
	e.Five = tiny(b[4])
	e.Six = tiny(b[5])
	e.Seven = tiny(b[6])
	e.Eight = tiny(b[7])
	e.Nine = tiny(b[8])
	e.Ten = tiny(b[9])
	e.Eleven = tiny(b[10])
	return nil
}

func (t *Twelve) MarshalJ() []byte {
	return []byte{byte(t.One), byte(t.Two), byte(t.Three), byte(t.Four), byte(t.Five), byte(t.Six), byte(t.Seven), byte(t.Eight), byte(t.Nine), byte(t.Ten), byte(t.Eleven), byte(t.Twelve)}
}

func (t *Twelve) UnmarshalJ(b []byte) error {
	if len(b) != 12 {
		return jay.ErrUnexpectedEOB
	}
	t.One = tiny(b[0])
	t.Two = tiny(b[1])
	t.Three = tiny(b[2])
	t.Four = tiny(b[3])
	t.Five = tiny(b[4])
	t.Six = tiny(b[5])
	t.Seven = tiny(b[6])
	t.Eight = tiny(b[7])
	t.Nine = tiny(b[8])
	t.Ten = tiny(b[9])
	t.Eleven = tiny(b[10])
	t.Twelve = tiny(b[11])
	return nil
}

func (t *Thirteen) MarshalJ() []byte {
	return []byte{byte(t.One), byte(t.Two), byte(t.Three), byte(t.Four), byte(t.Five), byte(t.Six), byte(t.Seven), byte(t.Eight), byte(t.Nine), byte(t.Ten), byte(t.Eleven), byte(t.Twelve), byte(t.Thirteen)}
}

func (t *Thirteen) UnmarshalJ(b []byte) error {
	if len(b) != 13 {
		return jay.ErrUnexpectedEOB
	}
	t.One = tiny(b[0])
	t.Two = tiny(b[1])
	t.Three = tiny(b[2])
	t.Four = tiny(b[3])
	t.Five = tiny(b[4])
	t.Six = tiny(b[5])
	t.Seven = tiny(b[6])
	t.Eight = tiny(b[7])
	t.Nine = tiny(b[8])
	t.Ten = tiny(b[9])
	t.Eleven = tiny(b[10])
	t.Twelve = tiny(b[11])
	t.Thirteen = tiny(b[12])
	return nil
}

func (f *Fourteen) MarshalJ() []byte {
	return []byte{byte(f.One), byte(f.Two), byte(f.Three), byte(f.Four), byte(f.Five), byte(f.Six), byte(f.Seven), byte(f.Eight), byte(f.Nine), byte(f.Ten), byte(f.Eleven), byte(f.Twelve), byte(f.Thirteen), byte(f.Fourteen)}
}

func (f *Fourteen) UnmarshalJ(b []byte) error {
	if len(b) != 14 {
		return jay.ErrUnexpectedEOB
	}
	f.One = tiny(b[0])
	f.Two = tiny(b[1])
	f.Three = tiny(b[2])
	f.Four = tiny(b[3])
	f.Five = tiny(b[4])
	f.Six = tiny(b[5])
	f.Seven = tiny(b[6])
	f.Eight = tiny(b[7])
	f.Nine = tiny(b[8])
	f.Ten = tiny(b[9])
	f.Eleven = tiny(b[10])
	f.Twelve = tiny(b[11])
	f.Thirteen = tiny(b[12])
	f.Fourteen = tiny(b[13])
	return nil
}

func (f *Fifteen) MarshalJ() []byte {
	return []byte{byte(f.One), byte(f.Two), byte(f.Three), byte(f.Four), byte(f.Five), byte(f.Six), byte(f.Seven), byte(f.Eight), byte(f.Nine), byte(f.Ten), byte(f.Eleven), byte(f.Twelve), byte(f.Thirteen), byte(f.Fourteen), byte(f.Fifteen)}
}

func (f *Fifteen) UnmarshalJ(b []byte) error {
	if len(b) != 15 {
		return jay.ErrUnexpectedEOB
	}
	f.One = tiny(b[0])
	f.Two = tiny(b[1])
	f.Three = tiny(b[2])
	f.Four = tiny(b[3])
	f.Five = tiny(b[4])
	f.Six = tiny(b[5])
	f.Seven = tiny(b[6])
	f.Eight = tiny(b[7])
	f.Nine = tiny(b[8])
	f.Ten = tiny(b[9])
	f.Eleven = tiny(b[10])
	f.Twelve = tiny(b[11])
	f.Thirteen = tiny(b[12])
	f.Fourteen = tiny(b[13])
	f.Fifteen = tiny(b[14])
	return nil
}

func (s *Sixteen) MarshalJ() []byte {
	return []byte{byte(s.One), byte(s.Two), byte(s.Three), byte(s.Four), byte(s.Five), byte(s.Six), byte(s.Seven), byte(s.Eight), byte(s.Nine), byte(s.Ten), byte(s.Eleven), byte(s.Twelve), byte(s.Thirteen), byte(s.Fourteen), byte(s.Fifteen), byte(s.Sixteen)}
}

func (s *Sixteen) UnmarshalJ(b []byte) error {
	if len(b) != 16 {
		return jay.ErrUnexpectedEOB
	}
	s.One = tiny(b[0])
	s.Two = tiny(b[1])
	s.Three = tiny(b[2])
	s.Four = tiny(b[3])
	s.Five = tiny(b[4])
	s.Six = tiny(b[5])
	s.Seven = tiny(b[6])
	s.Eight = tiny(b[7])
	s.Nine = tiny(b[8])
	s.Ten = tiny(b[9])
	s.Eleven = tiny(b[10])
	s.Twelve = tiny(b[11])
	s.Thirteen = tiny(b[12])
	s.Fourteen = tiny(b[13])
	s.Fifteen = tiny(b[14])
	s.Sixteen = tiny(b[15])
	return nil
}

func (s *Seventeen) MarshalJ() []byte {
	return []byte{byte(s.One), byte(s.Two), byte(s.Three), byte(s.Four), byte(s.Five), byte(s.Six), byte(s.Seven), byte(s.Eight), byte(s.Nine), byte(s.Ten), byte(s.Eleven), byte(s.Twelve), byte(s.Thirteen), byte(s.Fourteen), byte(s.Fifteen), byte(s.Sixteen), byte(s.Seventeen)}
}

func (s *Seventeen) UnmarshalJ(b []byte) error {
	if len(b) != 17 {
		return jay.ErrUnexpectedEOB
	}
	s.One = tiny(b[0])
	s.Two = tiny(b[1])
	s.Three = tiny(b[2])
	s.Four = tiny(b[3])
	s.Five = tiny(b[4])
	s.Six = tiny(b[5])
	s.Seven = tiny(b[6])
	s.Eight = tiny(b[7])
	s.Nine = tiny(b[8])
	s.Ten = tiny(b[9])
	s.Eleven = tiny(b[10])
	s.Twelve = tiny(b[11])
	s.Thirteen = tiny(b[12])
	s.Fourteen = tiny(b[13])
	s.Fifteen = tiny(b[14])
	s.Sixteen = tiny(b[15])
	s.Seventeen = tiny(b[16])
	return nil
}

func (e *Eighteen) MarshalJ() []byte {
	return []byte{byte(e.One), byte(e.Two), byte(e.Three), byte(e.Four), byte(e.Five), byte(e.Six), byte(e.Seven), byte(e.Eight), byte(e.Nine), byte(e.Ten), byte(e.Eleven), byte(e.Twelve), byte(e.Thirteen), byte(e.Fourteen), byte(e.Fifteen), byte(e.Sixteen), byte(e.Seventeen), byte(e.Eighteen)}
}

func (e *Eighteen) UnmarshalJ(b []byte) error {
	if len(b) != 18 {
		return jay.ErrUnexpectedEOB
	}
	e.One = tiny(b[0])
	e.Two = tiny(b[1])
	e.Three = tiny(b[2])
	e.Four = tiny(b[3])
	e.Five = tiny(b[4])
	e.Six = tiny(b[5])
	e.Seven = tiny(b[6])
	e.Eight = tiny(b[7])
	e.Nine = tiny(b[8])
	e.Ten = tiny(b[9])
	e.Eleven = tiny(b[10])
	e.Twelve = tiny(b[11])
	e.Thirteen = tiny(b[12])
	e.Fourteen = tiny(b[13])
	e.Fifteen = tiny(b[14])
	e.Sixteen = tiny(b[15])
	e.Seventeen = tiny(b[16])
	e.Eighteen = tiny(b[17])
	return nil
}

func (n *Nineteen) MarshalJ() []byte {
	return []byte{byte(n.One), byte(n.Two), byte(n.Three), byte(n.Four), byte(n.Five), byte(n.Six), byte(n.Seven), byte(n.Eight), byte(n.Nine), byte(n.Ten), byte(n.Eleven), byte(n.Twelve), byte(n.Thirteen), byte(n.Fourteen), byte(n.Fifteen), byte(n.Sixteen), byte(n.Seventeen), byte(n.Eighteen), byte(n.Nineteen)}
}

func (n *Nineteen) UnmarshalJ(b []byte) error {
	if len(b) != 19 {
		return jay.ErrUnexpectedEOB
	}
	n.One = tiny(b[0])
	n.Two = tiny(b[1])
	n.Three = tiny(b[2])
	n.Four = tiny(b[3])
	n.Five = tiny(b[4])
	n.Six = tiny(b[5])
	n.Seven = tiny(b[6])
	n.Eight = tiny(b[7])
	n.Nine = tiny(b[8])
	n.Ten = tiny(b[9])
	n.Eleven = tiny(b[10])
	n.Twelve = tiny(b[11])
	n.Thirteen = tiny(b[12])
	n.Fourteen = tiny(b[13])
	n.Fifteen = tiny(b[14])
	n.Sixteen = tiny(b[15])
	n.Seventeen = tiny(b[16])
	n.Eighteen = tiny(b[17])
	n.Nineteen = tiny(b[18])
	return nil
}

func (t *Twenty) MarshalJ() []byte {
	return []byte{byte(t.One), byte(t.Two), byte(t.Three), byte(t.Four), byte(t.Five), byte(t.Six), byte(t.Seven), byte(t.Eight), byte(t.Nine), byte(t.Ten), byte(t.Eleven), byte(t.Twelve), byte(t.Thirteen), byte(t.Fourteen), byte(t.Fifteen), byte(t.Sixteen), byte(t.Seventeen), byte(t.Eighteen), byte(t.Nineteen), byte(t.Twenty)}
}

func (t *Twenty) UnmarshalJ(b []byte) error {
	if len(b) != 20 {
		return jay.ErrUnexpectedEOB
	}
	t.One = tiny(b[0])
	t.Two = tiny(b[1])
	t.Three = tiny(b[2])
	t.Four = tiny(b[3])
	t.Five = tiny(b[4])
	t.Six = tiny(b[5])
	t.Seven = tiny(b[6])
	t.Eight = tiny(b[7])
	t.Nine = tiny(b[8])
	t.Ten = tiny(b[9])
	t.Eleven = tiny(b[10])
	t.Twelve = tiny(b[11])
	t.Thirteen = tiny(b[12])
	t.Fourteen = tiny(b[13])
	t.Fifteen = tiny(b[14])
	t.Sixteen = tiny(b[15])
	t.Seventeen = tiny(b[16])
	t.Eighteen = tiny(b[17])
	t.Nineteen = tiny(b[18])
	t.Twenty = tiny(b[19])
	return nil
}

func (t *TwentyOne) MarshalJ() []byte {
	return []byte{byte(t.One), byte(t.Two), byte(t.Three), byte(t.Four), byte(t.Five), byte(t.Six), byte(t.Seven), byte(t.Eight), byte(t.Nine), byte(t.Ten), byte(t.Eleven), byte(t.Twelve), byte(t.Thirteen), byte(t.Fourteen), byte(t.Fifteen), byte(t.Sixteen), byte(t.Seventeen), byte(t.Eighteen), byte(t.Nineteen), byte(t.Twenty), byte(t.TwentyOne)}
}

func (t *TwentyOne) UnmarshalJ(b []byte) error {
	if len(b) != 21 {
		return jay.ErrUnexpectedEOB
	}
	t.One = tiny(b[0])
	t.Two = tiny(b[1])
	t.Three = tiny(b[2])
	t.Four = tiny(b[3])
	t.Five = tiny(b[4])
	t.Six = tiny(b[5])
	t.Seven = tiny(b[6])
	t.Eight = tiny(b[7])
	t.Nine = tiny(b[8])
	t.Ten = tiny(b[9])
	t.Eleven = tiny(b[10])
	t.Twelve = tiny(b[11])
	t.Thirteen = tiny(b[12])
	t.Fourteen = tiny(b[13])
	t.Fifteen = tiny(b[14])
	t.Sixteen = tiny(b[15])
	t.Seventeen = tiny(b[16])
	t.Eighteen = tiny(b[17])
	t.Nineteen = tiny(b[18])
	t.Twenty = tiny(b[19])
	t.TwentyOne = tiny(b[20])
	return nil
}

func (t *TwentyTwo) MarshalJ() []byte {
	return []byte{byte(t.One), byte(t.Two), byte(t.Three), byte(t.Four), byte(t.Five), byte(t.Six), byte(t.Seven), byte(t.Eight), byte(t.Nine), byte(t.Ten), byte(t.Eleven), byte(t.Twelve), byte(t.Thirteen), byte(t.Fourteen), byte(t.Fifteen), byte(t.Sixteen), byte(t.Seventeen), byte(t.Eighteen), byte(t.Nineteen), byte(t.Twenty), byte(t.TwentyOne), byte(t.TwentyTwo)}
}

func (t *TwentyTwo) UnmarshalJ(b []byte) error {
	if len(b) != 22 {
		return jay.ErrUnexpectedEOB
	}
	t.One = tiny(b[0])
	t.Two = tiny(b[1])
	t.Three = tiny(b[2])
	t.Four = tiny(b[3])
	t.Five = tiny(b[4])
	t.Six = tiny(b[5])
	t.Seven = tiny(b[6])
	t.Eight = tiny(b[7])
	t.Nine = tiny(b[8])
	t.Ten = tiny(b[9])
	t.Eleven = tiny(b[10])
	t.Twelve = tiny(b[11])
	t.Thirteen = tiny(b[12])
	t.Fourteen = tiny(b[13])
	t.Fifteen = tiny(b[14])
	t.Sixteen = tiny(b[15])
	t.Seventeen = tiny(b[16])
	t.Eighteen = tiny(b[17])
	t.Nineteen = tiny(b[18])
	t.Twenty = tiny(b[19])
	t.TwentyOne = tiny(b[20])
	t.TwentyTwo = tiny(b[21])
	return nil
}

func (t *TwentyThree) MarshalJ() []byte {
	return []byte{byte(t.One), byte(t.Two), byte(t.Three), byte(t.Four), byte(t.Five), byte(t.Six), byte(t.Seven), byte(t.Eight), byte(t.Nine), byte(t.Ten), byte(t.Eleven), byte(t.Twelve), byte(t.Thirteen), byte(t.Fourteen), byte(t.Fifteen), byte(t.Sixteen), byte(t.Seventeen), byte(t.Eighteen), byte(t.Nineteen), byte(t.Twenty), byte(t.TwentyOne), byte(t.TwentyTwo), byte(t.TwentyThree)}
}

func (t *TwentyThree) UnmarshalJ(b []byte) error {
	if len(b) != 23 {
		return jay.ErrUnexpectedEOB
	}
	t.One = tiny(b[0])
	t.Two = tiny(b[1])
	t.Three = tiny(b[2])
	t.Four = tiny(b[3])
	t.Five = tiny(b[4])
	t.Six = tiny(b[5])
	t.Seven = tiny(b[6])
	t.Eight = tiny(b[7])
	t.Nine = tiny(b[8])
	t.Ten = tiny(b[9])
	t.Eleven = tiny(b[10])
	t.Twelve = tiny(b[11])
	t.Thirteen = tiny(b[12])
	t.Fourteen = tiny(b[13])
	t.Fifteen = tiny(b[14])
	t.Sixteen = tiny(b[15])
	t.Seventeen = tiny(b[16])
	t.Eighteen = tiny(b[17])
	t.Nineteen = tiny(b[18])
	t.Twenty = tiny(b[19])
	t.TwentyOne = tiny(b[20])
	t.TwentyTwo = tiny(b[21])
	t.TwentyThree = tiny(b[22])
	return nil
}
