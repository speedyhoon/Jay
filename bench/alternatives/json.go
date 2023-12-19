package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	car := Car2{ID: 42, Auto: true, RedLine: 10530, Name: "Hello", gearbox: gearbox{
		Gears:        4,
		Reverse:      1,
		Sequential:   true,
		Automatic:    false,
		Model:        "H23w",
		Manufacturer: "Zebra",
		LinkageDelta: -6,
	}}

	src, err := json.Marshal(car)
	fmt.Println(err)
	fmt.Printf("%s\n", src)
}

// Car length = 8 + 8 + 1 + len.Car + 1 + 1 + len.CC + 1 + len.Timing + len.Gearbox
type Car2 struct {
	ID      uint64 `vl:"id" j:""`
	Row     uint   `vl:"id" j:"max:677"`
	Name    string `vl:"name,omitempty" json:"-"`
	Auto    bool
	CC      string `z:",omitempty"`
	Timing  string `z:",omitempty" j:"max:24"`
	RedLine uint16
	gearbox gearbox
}

// Gearbox length = 4 + 1 + 1+ 1+ 1+ len.Model + 1 + len.Make
type gearbox struct {
	Gears        int
	Reverse      uint8
	Sequential   bool
	Automatic    bool
	Model        string
	Manufacturer string
	LinkageDelta int8 // The length of the gearbox linkage cable compared to standard equipment measured in millimeters.
}
