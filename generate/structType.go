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
	bufferName string // Prefix assigned to the buffer variable.
	lengthName string // The variable name for storing `len(bufferName)`. Separate to the `l` prefix assigned to the length variables.
	dir        string
	imports    importList // Imports required for generated code.
	option     *Option    // Global options from commandline or API calls, different to field.tagOptions.

	// Lists of exported fields detected during AST traversal.
	bool, // Boolean fields are joined together and represented as binary.
	single, // Fields represented in one byte like int8 & uint8. These have additional optimizations.
	fixedLen, // Fixed length types like int16, uint64 and some arrays etc.
	stringSlice,
	variableLen fieldList // Variable length fields like string and all slice types. These are generated last and have the most processing overhead.

	// When true: returns the Marshal code in a single return statement, when false: MarshalJ() method contains several lines.
	returnInline          bool
	returnInlineUnmarshal canReturnInlined

	qtyBytesRequired uint

	isImportJ *bool
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
