package generate

import "unicode"

const (
	receiverNameDefault   = "b"
	receiverNameAlternate = "y"
	lengthNameDefault     = "l"
	lengthNameAlternate   = "z"
)

type structTyp struct {
	name       string
	receiver   string
	bufferName string
	lengthName string
	dir        string
	imports    importList // Imports required for generated code.
	option     *Option

	// Exported fields.
	fixedLen, // Fixed length types like int16, uint64 and some arrays etc.
	single, // Fields represented in one byte like int8 & uint8. These have additional optimisations.
	variableLen, // Variable length fields like string and all slice types. These are generated last & have the most processing overhead.
	bool, // Boolean fields are joined together and represented as binary.
	boolArray fieldList
}

func newStructTyp(dir, typeName string, o *Option) *structTyp {
	r := receiverName(typeName)
	return &structTyp{
		name:       typeName,
		receiver:   r,
		dir:        dir,
		bufferName: bufferName(r),
		lengthName: lengthName(r),
		option:     o,
	}
}

func receiverName(typeName string) string {
	return string(unicode.ToLower([]rune(typeName)[0]))
}

func bufferName(receiverName string) string {
	if receiverName == receiverNameDefault {
		return receiverNameAlternate
	}
	return receiverNameDefault
}

func lengthName(receiverName string) string {
	if receiverName == lengthNameDefault {
		return lengthNameAlternate
	}
	return lengthNameDefault
}
