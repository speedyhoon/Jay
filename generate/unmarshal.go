package generate

import (
	"bytes"
	"fmt"
	"github.com/speedyhoon/jay"
	"strings"
)

// makeUnmarshal ...
func (s *structTyp) makeUnmarshal(b *bytes.Buffer) {
	var byteIndex = uint(len(s.variableLen))
	buf := bytes.NewBuffer(nil)

	s.makeReadBools(buf, &byteIndex, s.receiver)
	s.readSingles(buf, &byteIndex)

	for _, f := range s.fixedLen {
		buf.WriteString(f.unmarshalLine(&byteIndex, Utoa(byteIndex), "", ""))
		buf.WriteString("\n")
	}

	at, end := s.defineTrackingVars(buf, byteIndex)
	for i, f := range s.variableLen {
		at, end = s.tracking(buf, i, end, byteIndex, f.typ)
		buf.WriteString(f.unmarshalLine(&byteIndex, at, end, lenVariable(i)))
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

func (f *field) unmarshalLine(byteIndex *uint, at, end, lenVar string) string {
	fun, template := f.unmarshalFuncs()
	totalSize := f.typeFuncSize()

	if f.isFixedLen {
		*byteIndex += totalSize
	}
	thisField := pkgSelName(f.structTyp.receiver, f.name)
	if end == "" {
		end = Utoa(*byteIndex)
	}

	switch template {
	case tFunc:
		if f.isAliasDef && f.arraySize == typeNotArrayOrSlice {
			return fmt.Sprintf("%s = %s", thisField, printFunc(f.aliasType, printFunc(fun, f.sliceExpr(at, end))))
		}
		return fmt.Sprintf("%s = %s", thisField, printFunc(fun, f.sliceExpr(at, end)))

	case tFuncOpt:
		return fmt.Sprintf("if %s != 0 {\n%s = %s\n}", lenVar, thisField, f.sliceExpr(at, end))

	case tFuncLength:
		return fmt.Sprintf("%s = %s", thisField, printFunc(fun, f.sliceExpr(at, end), lenVar))

	case tByteConv:
		return fmt.Sprintf("%s = %s", thisField, printFunc(fun, f.sliceExpr(at, end)))

	case tByteAssign:
		return fmt.Sprintf("%s = %s", thisField, f.sliceExpr(at, end))
	}

	lg.Println("unhandled template")
	return ""
}

// unmarshalFuncs returns the function name to handle unmarshalling.
// `size` is the quantity of bytes required to represent the type.
func (f field) unmarshalFuncs() (funcName string, template uint8) {
	var c interface{}
	switch f.typ {
	case "byte", "uint8":
		if f.isAliasDef {
			return f.aliasType, tByteConv
		}
		return "", tByteAssign
	case "int8", "string":
		if f.isAliasDef {
			return f.aliasType, tByteConv
		}
		return f.typ, tByteConv
	case "int":
		if f.structTyp.option.FixedIntSize {
			if f.structTyp.option.Is32bit {
				c, template = jay.ReadIntArch32, tFunc
			}
			c, template = jay.ReadIntArch64, tFunc
			break
		}
		//c = jay.ReadIntVariable
		c, template = jay.ReadInt, tFunc
	case "int16":
		c, template = jay.ReadInt16, tFunc
	case "int32", "rune":
		c, template = jay.ReadInt32, tFunc
	case "float32":
		c, template = jay.ReadFloat32, tFunc
	case "float64":
		c, template = jay.ReadFloat64, tFunc
	case "int64":
		c, template = jay.ReadInt64, tFunc
	case "time.Duration":
		c, template = jay.ReadDuration, tFunc
	case "uint":
		if f.structTyp.option.FixedUintSize {
			if f.structTyp.option.Is32bit {
				c, template = jay.ReadUintArch32, tFunc
			}
			c, template = jay.ReadUintArch64, tFunc
			break
		}
		c, template = jay.ReadUintVariable, tFunc
	case "uint16":
		c, template = jay.ReadUint16, tFunc
	case "uint32":
		c, template = jay.ReadUint32, tFunc
	case "uint64":
		c, template = jay.ReadUint64, tFunc
	case "time.Time":
		if f.tagOptions.TimeNano {
			c, template = jay.ReadTimeNano, tFunc
		} else {
			c, template = jay.ReadTime, tFunc
		}
	case "[]int8":
		c, template = jay.ReadInt8s, tFuncLength
	case "[]bool":
		c, template = jay.ReadBools, tFuncLength

	case "[]byte", "[]uint8":
		if f.Required {
			c, template = "", tByteAssign
		} else {
			c, template = copyKeyword, tFuncOpt
		}
	case "[]float32":
		c, template = jay.ReadFloat32s, tFuncLength
	case "[]float64":
		c, template = jay.ReadFloat64s, tFuncLength
	case "[]int16":
		c, template = jay.ReadInt16s, tFuncLength
	case "[]uint":
		if f.structTyp.option.Is32bit {
			c, template = jay.ReadUintsX32, tFuncLength
		}
		c, template = jay.ReadUintsX64, tFuncLength
	case "[]uint16":
		c, template = jay.ReadUint16s, tFuncLength

	default:
		lg.Printf("no function set for type %s yet in unmarshalFuncs()", f.typ)
	}

	return nameOf(c, nil), template
}
