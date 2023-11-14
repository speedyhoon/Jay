package jay

import (
	"bytes"
	"fmt"
)

// MakeUnmarshal ...
func (s *Struct) MakeUnmarshal(b *bytes.Buffer) {
	b.WriteString(fmt.Sprintf(
		"func (%s *%s) UnmarshalJ(b []byte) (err error) {\nreturn nil\n}\n",
		s.ReceiverName(),
		s.name,
	))
}
