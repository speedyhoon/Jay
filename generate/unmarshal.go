package generate

import (
	"bytes"
	"fmt"
	"github.com/speedyhoon/jay"
	"github.com/speedyhoon/utl"
	"strings"
)

func (s *structTyp) ReturnInline() bool {
	return len(s.fixedLen) == 0 && len(s.single) == 0 && len(s.variableLen) == 0 && len(s.bool) == 0 && len(s.boolArray) == 0 && len(s.stringSlice) >= 1
}

// makeUnmarshal ...
func (s *structTyp) makeUnmarshal(b *bytes.Buffer) {
	var byteIndex = uint(len(s.variableLen))
	buf := bytes.NewBuffer(nil)

	s.makeReadBools(buf, &byteIndex)
	s.readSingles(buf, &byteIndex)

	for _, f := range s.fixedLen {
		buf.WriteString(f.unmarshalLine(&byteIndex, utl.UtoA(byteIndex), "", ""))
		buf.WriteString("\n")
	}

	at, end := s.defineTrackingVars(buf, byteIndex)
	for i, f := range s.variableLen {
		at, end = s.tracking(buf, i, end, byteIndex, f.typ)
		lenVar := lenVariable(i)
		buf.WriteString(f.unmarshalLine(&byteIndex, at, end, lenVar))
		buf.WriteString("\n")
	}

	for _, f := range s.stringSlice {
		//at, end = s.tracking(buf, i, end, byteIndex, f.typ)
		//lenVar := lenVariable(i)
		at = "at"
		buf.WriteString(f.unmarshalLine(&byteIndex, at, end, ""))
		buf.WriteString("\n")
	}

	code := buf.Bytes()
	if len(code) == 0 {
		return
	}

	// Prevent panic: runtime error: index out of range
	var lengthCheck string
	if len(s.variableLen) == 0 && len(s.stringSlice) == 0 {
		lengthCheck = fmt.Sprintf("if len(%s) != %d {\nreturn %s\n}", s.bufferName, byteIndex, exportedErr)
	} else if len(s.variableLen) == 0 && len(s.stringSlice) >= 1 {
		lengthCheck = fmt.Sprintf("if len(%s) < %d {\nreturn %s\n}", s.bufferName, byteIndex, exportedErr)
	} else {
		lengthCheck = fmt.Sprintf("%[1]s := len(%[2]s)\nif %[1]s < %[3]d {\nreturn %[4]s\n}", s.lengthName, s.bufferName, byteIndex, exportedErr)
	}
	variableLengthCheck := s.generateCheckSizes(exportedErr, byteIndex)

	if s.ReturnInline() {
		bufWriteF(b,
			"func (%s *%s) UnmarshalJ(%s []byte) error {\n%s\n%s\n}\n",
			s.receiver,
			s.name,
			s.bufferName,
			lengthCheck,
			code,
		)
		return
	}

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

	assignments, values := make([]string, qty), make([]string, qty)
	var conditions []string
	var sizeChecks []sizeCheck
	for i, f := range s.variableLen {
		assignments[i] = lenVariable(i)
		values[i] = fmt.Sprintf("int(%s[%d])", s.bufferName, i)
		if f.typ == tBoolS {
			conditions = append(conditions, printFunc(nameOf(jay.SizeBools, nil), assignments[i]))
		} else {
			sizeChecks = append(sizeChecks, sizeCheck{name: assignments[i], bytesNeeded: f.elmSize})
		}
	}

	return fmt.Sprintf(
		"%s := %s\nif %s != %d+%s {\nreturn %s\n}\n",
		strings.Join(assignments, ", "),
		strings.Join(values, ", "),
		s.lengthName,
		totalSize,
		groupConditionSizes(sizeChecks, conditions),
		exportedErr,
	)
}

func (f *field) unmarshalLine(byteIndex *uint, at, end, lenVar string) string {
	fun, template, canReturnInline := f.unmarshalFuncs()
	totalSize := f.typeFuncSize()

	if f.isFixedLen || f.typ == tStrings {
		*byteIndex += totalSize
	}
	thisField := pkgSelName(f.structTyp.receiver, f.name)
	if end == "" {
		end = utl.UtoA(*byteIndex)
	}

	switch template {
	case tFunc:
		if f.isDef && f.isNotArrayOrSlice() {
			return fmt.Sprintf("%s = %s", thisField, printFunc(f.aliasType, printFunc(fun, f.sliceExpr(at, end))))
		}
		return fmt.Sprintf("%s = %s", thisField, printFunc(fun, f.sliceExpr(at, end)))

	case tFuncOpt:
		return fmt.Sprintf("if %s != 0 {\n%s = %s\n}", lenVar, thisField, f.sliceExpr(at, end))

	case tFuncLength:
		return fmt.Sprintf("%s = %s", thisField, printFunc(fun, f.sliceExpr(at, end), lenVar))

	case tFuncPtr:
		return fmt.Sprintf("%s%s(%s, &%s)", canReturnInline, fun, f.sliceExpr(at, end), thisField)

	case tFuncPtrCheck:
		return fmt.Sprintf(
			"at, ok := %s(%s, &%s)\n\tif !ok {\n\t\treturn %s\n\t}",
			fun, f.structTyp.bufferName, thisField, exportedErr,
		)

	case tFuncPtrCheckAt:
		return fmt.Sprintf(
			"if !%s(%s, &%s, &at) {\n\t\treturn %s\n\t}",
			fun, f.sliceExpr(at, end), thisField, exportedErr,
		)

	case tByteConv:
		return fmt.Sprintf("%s = %s", thisField, printFunc(fun, f.sliceExpr(at, end)))

	case tByteAssign:
		return fmt.Sprintf("%s = %s", thisField, printFunc(f.convertTo(), f.sliceExpr(at, end)))
	}

	lg.Println("unhandled template")
	return ""
}

type canReturnInlined bool

func (c canReturnInlined) String() string {
	if c {
		return "return "
	}
	return ""
}

func (f *field) convertTo() string {
	if f.isArray() {
		return fmt.Sprintf("[%d]%s", f.arraySize, f.arrayType)
	}
	return ""
}

func (f *field) typeConvert() string {
	if f.pkgReq != "" {
		f.structTyp.imports.add(f.pkgReq)
		return pkgSelName(f.pkgReq, f.aliasType)
	}
	return f.aliasType
}

// unmarshalFuncs returns the function name to handle unmarshalling.
// `size` is the quantity of bytes required to represent the type.
func (f field) unmarshalFuncs() (funcName string, template uint8, canReturnInline canReturnInlined) {
	var c interface{}
	switch f.typ {
	case tByte:
		if f.isDef {
			return f.typeConvert(), tByteConv, false
		}
		return "", tByteAssign, false
	case tInt8, tString:
		if f.isDef {
			return f.typeConvert(), tByteConv, false
		}
		return f.typ, tByteConv, false
	case tInt:
		if f.structTyp.option.FixedIntSize {
			if f.structTyp.option.Is32bit {
				c, template = jay.ReadIntX32, tFunc
			}
			c, template = jay.ReadIntX64, tFunc
			break
		}
		//c = jay.ReadIntVariable
		c, template = jay.ReadInt, tFunc
	case tInt16:
		c, template = jay.ReadInt16, tFunc
	case tInt32:
		c, template = jay.ReadInt32, tFunc
	case tInt64:
		c, template = jay.ReadInt64, tFunc
	case tFloat32:
		c, template = jay.ReadFloat32, tFunc
	case tFloat64:
		c, template = jay.ReadFloat64, tFunc
	case tUint:
		if f.structTyp.option.FixedUintSize {
			if f.structTyp.option.Is32bit {
				c, template = jay.ReadUintX32, tFunc
			}
			c, template = jay.ReadUintX64, tFunc
			break
		}
		c, template = jay.ReadUintVariable, tFunc
	case tUint16:
		c, template = jay.ReadUint16, tFunc
	case tUint32:
		c, template = jay.ReadUint32, tFunc
	case tUint64:
		c, template = jay.ReadUint64, tFunc
	case tTime:
		if f.tagOptions.TimeNano {
			c, template = jay.ReadTimeNano, tFunc
		} else {
			c, template = jay.ReadTime, tFunc
		}
	case tInt8S:
		c, template = jay.ReadInt8s, tFuncLength
	case tBoolS:
		c, template = jay.ReadBools, tFuncLength

	case tByteS:
		if f.isArray() {
			c, template = "", tByteAssign
			return
		}

		if f.Required {
			c, template = "", tByteAssign
		} else {
			c, template = copyKeyword, tFuncOpt
		}
	case tFloat32S:
		c, template = jay.ReadFloat32s, tFuncLength
	case tFloat64S:
		c, template = jay.ReadFloat64s, tFuncLength
	case tInt16S:
		c, template = jay.ReadInt16s, tFuncLength
	case tUintS:
		if f.structTyp.option.Is32bit {
			c, template = jay.ReadUintsX32, tFuncLength
		}
		c, template = jay.ReadUintsX64, tFuncLength
	case tUint16S:
		c, template = jay.ReadUint16s, tFuncLength
	case tIntS:
		if f.structTyp.option.Is32bit {
			c, template = jay.ReadIntsX32, tFuncLength
		}
		c, template = jay.ReadIntsX64, tFuncLength
	case tInt32S:
		c, template = jay.ReadInt32s, tFuncLength
	case tInt64S:
		c, template = jay.ReadInt64s, tFuncLength
	case tUint32S:
		c, template = jay.ReadUint32s, tFuncLength
	case tUint64S:
		c, template = jay.ReadUint64s, tFuncLength
	case tTimeDurations:
		c, template = jay.ReadDurations, tFuncLength
	case tStrings:
		if f.isLast {
			//switch {
			//case f.tagOptions.MaxQty <= maxUint8:
			//	c, template = jay.ReadStrings8Err, tFunc
			//default:
			c, template, canReturnInline = jay.ReadStrings16Err, tFuncPtr, canReturnInlined(f.isLast)
			//}
			//break
		} else if f.isFirst {
			c, template = jay.ReadStrings16n, tFuncPtrCheck
		} else {
			//switch {
			//case f.tagOptions.MaxQty <= maxUint8:
			//	c, template = jay.ReadStrings8nErr, tFuncPtrCheck
			//default:
			c, template = jay.ReadStrings16nb, tFuncPtrCheckAt
			//}
		}

	default:
		lg.Printf("no function set for type %s yet in unmarshalFuncs()", f.typ)
	}

	return nameOf(c, nil), template, canReturnInline
}
