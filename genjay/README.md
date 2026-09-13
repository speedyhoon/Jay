# genjay

[![Go Reference](https://pkg.go.dev/badge/github.com/speedyhoon/jay/genjay.svg)](https://pkg.go.dev/github.com/speedyhoon/jay/genjay)
[![Go Report Card](https://raw.githubusercontent.com/speedyhoon/speedyhoon/refs/heads/main/goReport.svg)](https://goreportcard.com/report/github.com/speedyhoon/Jay/genjay)
![license AGPL3](https://raw.githubusercontent.com/speedyhoon/speedyhoon/refs/heads/main/AGPL3.svg)

Traverses `.go` files to find exported Go `structs` to generate marshalling `.MarshalJ()` and unmarshalling `.UnmarshalJ()` methods for the [Jay serialization format](https://github.com/speedyhoon/jay).

## Field tag options
`j:-` Ignore this field.

### TODO
`max`
* Maximum value for integers and floats.
* Maximum length for strings and slices.

`min`
* Minimum value for integers and floats.
* Minimum length for strings and slices.

`req`
* Flag any string or slice field as required, indicating that it will never be empty.
* A micro optimization for marshalling, removing `length == 0` if statement before calling the function to convert into `[]byte`.

## Embedded structs
Any private struct can be embedded into an exported struct when an exported field name is included. For example,
```go
package main

type Car struct {
	Wheels          // Added - exported type
	W       Wheels  // Added - exported field name
	Gearbox gearbox // Added - exported field name

	Axel  // Ignored - no fields
	Turbo // Ignored - no exported fields

	gearbox         // Ignored - not an exported type
	gbx     gearbox // Ignored - not an exported field name
	_       Wheels  // Ignored - not an exported field name

	Engine `j:-`         // Ignored - flag present
	Gbx    gearbox `j:-` // Ignored - flag present
}

type gearbox struct {
	GearsQty uint8
}

type Axel struct{}

type Turbo struct {
	size uint16
}

type Wheels struct {
	Offset int
}

type Engine struct {
	CC uint16
}
```

## Features to add
* Compress enum integers with small value ranges.
