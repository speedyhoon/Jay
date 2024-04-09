package generate

import (
	"bytes"
	"fmt"
	"github.com/speedyhoon/jay"
	"strings"
)

// makeUnmarshal ...
func (s *structTyp) makeUnmarshal(b *bytes.Buffer, o Option) {
	var byteIndex = uint(len(s.variableLen))
	buf := bytes.NewBuffer(nil)

	s.makeReadBools(buf, &byteIndex, s.receiver)
	s.readSingles(buf, &byteIndex, s.receiver, o)

	for _, f := range s.fixedLen {
		buf.WriteString(o.unmarshalLine(s, f, &byteIndex, "", "", ""))
		buf.WriteString("\n")
	}

	at, end := s.defineTrackingVars(buf, byteIndex)
	for i, f := range s.variableLen {
		at, end = s.tracking(buf, i, end, byteIndex, f.typ)
		buf.WriteString(o.unmarshalLine(s, f, &byteIndex, at, end, lenVariable(i)))
		buf.WriteString("\n")
	}

	code := buf.Bytes()
	if len(code) == 0 {
		return
	}

	const exportedErr = "jay.ErrUnexpectedEOB"
	// Prevent panic: runtime error: index out of range
	var lengthCheck string
	if len(s.variableLen) == 0 {
		lengthCheck = fmt.Sprintf("if len(%s) < %d {\nreturn %s\n}", s.bufferName, byteIndex, exportedErr)
	} else {
		lengthCheck = fmt.Sprintf("%[1]s := len(%[2]s)\nif %[1]s < %[3]d {\nreturn %[4]s\n}", s.lengthName, s.bufferName, byteIndex, exportedErr)
	}
	variableLengthCheck := s.generateCheckSizes(exportedErr, byteIndex)

	bufWriteF(b,
		"func (%s *%s) UnmarshalJ(%s []byte) error {\n%s\n%s%sreturn nil\n}\n",
		s.receiver,
		s.name,
		s.bufferName,
		lengthCheck,
		variableLengthCheck,
		code,
	)
}

func (s *structTyp) generateCheckSizes(exportedErr string, totalSize uint) string {
	qty := len(s.variableLen)
	if qty == 0 {
		return ""
	}

	assignments, values, conditions := make([]string, qty), make([]string, qty), make([]string, qty)
	for i, f := range s.variableLen {
		assignments[i] = lenVariable(i)
		values[i] = fmt.Sprintf("int(%s[%d])", s.bufferName, i)
		if f.typ == "[]bool" {
			conditions[i] = fmt.Sprintf("%s(%s)", nameOf(jay.SizeBools, nil), assignments[i])
		} else {
			conditions[i] = assignments[i]
		}
	}

	return fmt.Sprintf(
		"%s := %s\nif %s < %d+%s {\nreturn %s\n}\n",
		strings.Join(assignments, ", "),
		strings.Join(values, ", "),
		s.lengthName,
		totalSize,
		strings.Join(conditions, "+"),
		exportedErr,
	)
}

func (o Option) unmarshalLine(s *structTyp, f field, byteIndex *uint, at, end, lenVar string) string {
	fun, totalSize := o.unmarshalFuncs(f), f.typeFuncSize(o)
	if fun == "" {
		// Unknown type, not supported yet.
		lg.Printf("no generateLine for type `%s` yet in unmarshalLine()", f.typ)
		return ""
	}

	start := *byteIndex
	if f.isFixedLen {
		*byteIndex += totalSize
	}
	thisField := pkgSelName(s.receiver, f.name)

	switch f.typ {
	case "string":
		if f.isAliasDef {
			fun = f.aliasType
		}
		return fmt.Sprintf("%s = %s(%s)", thisField, fun, sliceExpr(s, f, at, end))
	case "[]byte", "[]uint8":
		if f.isAliasDef {
			fun = f.aliasType
		}
		if f.Required {
			return fmt.Sprintf("%s = %s", thisField, sliceExpr(s, f, at, end))
		}
		return fmt.Sprintf("if %s != 0 {\n%s = %s\n}", lenVar, thisField, sliceExpr(s, f, at, end))
	case "[]int8", "[]bool":
		return fmt.Sprintf("%s = %s(%s, %s)", thisField, fun, sliceExpr(s, f, at, end), lenVar)
	}

	if f.isArray() && f.arrayType == "int8" {
		values := make([]string, f.arraySize)
		for i := 0; i < f.arraySize; i++ {
			values[i] = fmt.Sprintf("%s(%s[%d])", f.arrayType, s.bufferName, i)
		}
		return fmt.Sprintf("%s = %s{%s}", thisField, f.typ, strings.Join(values, ", "))
	}

	fun = printFunc(fun, sliceExprU(s, f, start, *byteIndex))
	if f.isAliasDef {
		fun = printFunc(f.aliasType, fun)
	}
	return fmt.Sprintf("%s = %s", thisField, fun)
}

// unmarshalFuncs returns the function name to handle unmarshalling.
// `size` is the quantity of bytes required to represent the type.
func (o Option) unmarshalFuncs(f field) (funcName string) {
	var c interface{}
	switch f.typ {
	case "byte", "uint8":
		if f.isAliasDef {
			return f.aliasType
		}
		return ""
	case "int8":
		if f.isAliasDef {
			return f.aliasType
		}
		return "int8"
	case "string":
		return "string"
	case "int":
		if o.FixedIntSize {
			if o.Is32bit {
				c = jay.ReadIntArch32
			}
			c = jay.ReadIntArch64
			break
		}
		//c = jay.ReadIntVariable
		c = jay.ReadInt
	case "int16":
		c = jay.ReadInt16
	case "int32", "rune":
		c = jay.ReadInt32
	case "float32":
		c = jay.ReadFloat32
	case "float64":
		c = jay.ReadFloat64
	case "int64":
		c = jay.ReadInt64
	case "time.Duration":
		c = jay.ReadDuration
	case "uint":
		if o.FixedUintSize {
			if o.Is32bit {
				c = jay.ReadUintArch32
			}
			c = jay.ReadUintArch64
			break
		}
		c = jay.ReadUintVariable
	case "uint16":
		c = jay.ReadUint16
	case "uint32":
		c = jay.ReadUint32
	case "uint64":
		c = jay.ReadUint64
	case "time.Time":
		if f.tagOptions.TimeNano {
			c = jay.ReadTimeNano
		} else {
			c = jay.ReadTime
		}
	case "[]int8":
		c = jay.ReadInt8s
	case "[]bool":
		c = jay.ReadBools

	case "[]byte", "[]uint8":
		c = copyKeyword
	case "[]float32":
		c = jay.ReadFloat32s

	default:
		lg.Printf("no function set for type %s yet in unmarshalFuncs()", f.typ)
	}

	return nameOf(c, nil)
}
