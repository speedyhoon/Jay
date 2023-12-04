package main

import (
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile)
	car := Car{ID: 42, Auto: true, RedLine: 10530, Name: "Hello", Gearbox: gearbox{
		Gears:        4,
		Reverse:      1,
		Sequential:   true,
		Automatic:    false,
		Model:        "H23w",
		Manufacturer: "Zebra",
		LinkageDelta: -6,
	}}
	//byt, _ := json.Marshal(car)
	//log.Printf("%s\n", byt)

	src := car.MarshalJ()
	log.Println(src)
	log.Printf("%c", src)

	var car2 Car
	err := car2.UnmarshalJ(src)
	log.Println("err:", err)
	log.Printf("%+v", car2)
}

// Car length = 8 + 8 + 1 + len.Car + 1 + 1 + len.CC + 1 + len.Timing + len.Gearbox
type Car struct {
	ID      uint64 `vl:"id" j:""`
	Row     uint   `vl:"id" j:"max:677"`
	Name    string `vl:"name,omitempty" json:"-"`
	Auto    bool
	CC      string `z:",omitempty"`
	Timing  string `z:",omitempty" j:"max:24"`
	RedLine uint16
	Gearbox gearbox
}

// gearbox length = 4 + 1 + 1+ 1+ 1+ len.Model + 1 + len.Make
type gearbox struct {
	Gears        int
	Reverse      uint8
	Sequential   bool
	Automatic    bool
	Model        string
	Manufacturer string
	LinkageDelta int8 // The length of the gearbox linkage cable compared to standard equipment measured in millimeters.
}
