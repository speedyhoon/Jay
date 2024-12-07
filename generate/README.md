# generate
Traverses `.go` files to find exported Go `structs` to generate marshalling `.MarshalJ()` and unmarshalling `.UnmarshalJ()` methods for the [Jay serialisation format](https://github.com/speedyhoon/jay).

## Field tag options
`j:-` Ignore exported field.

### TODO
`max`
* maximum value for integers and floats
* maximum length for strings and slices

`min`
* minimum value for integers and floats
* minimum length for strings and slices

## Embedded structs
Any private struct can be embedded into an exported struct when an exported field name is included. For example:
```go
type Car struct {
	gearbox         // Ignored - not an exported field name
	Gearbox gearbox // Added - exported type
	gbx gearbox     // Ignored - not an exported field name

	Axel    // Ignored - no fields
	Turbo   // Ignored - no exported fields

	Wheels   // Added - exported type
	W Wheels // Added - exported field name
	w Wheels // Ignored - not an exported field name
	_ Wheels // Ignored - not an exported field name

	Wheels  `j:-`     // Ignored flag present
	Wh Wheels `j:-`   // Ignored flag present
	Gbx gearbox `j:-` // Ignored flag present
}

type gearbox struct {
	GearsQty uint8
}

type Axel struct {}

type Turbo struct {
	size uint16
}

type Wheels struct {
	Offset int
	- int
}
```

## Features to add
* Compress enum integers with small value ranges.