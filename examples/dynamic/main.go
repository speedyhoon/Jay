package main

import "fmt"

const (
	tCar = iota + 1
	tPet
)

func main() {
	data := Message{
		Type: tPet,
		Data: Pet{
			Name:    "Giggles",
			Species: "Bird",
			Breed:   "Kookaburra",
		}.MarshalJ(),
	}

	err := decode(data.MarshalJ())
	if err != nil {
		fmt.Println("err:", err)
	}

	data = Message{
		Type: tCar,
		Data: Car{
			Make:  "Ford",
			Model: "Escort",
			Year:  1993,
		}.MarshalJ(),
	}
	err = decode(data.MarshalJ())
	if err != nil {
		fmt.Println("err:", err)
	}
}

func decode(src []byte) (err error) {
	fmt.Println(src)

	var message Message
	err = message.UnmarshalJ(src)
	if err != nil {
		return
	}

	switch message.Type {
	case tCar:
		var car Car
		err = car.UnmarshalJ(message.Data)
		fmt.Printf("%+v\n", car)
	case tPet:
		var pet Pet
		err = pet.UnmarshalJ(message.Data)
		fmt.Printf("%+v\n", pet)
	default:
		return fmt.Errorf("unknown message type %d", message.Type)
	}
	return
}

type (
	Message struct {
		Type uint8
		Data []byte
	}

	Car struct {
		Make  string
		Model string
		Year  uint16
	}

	Pet struct {
		Name    string
		Species string
		Breed   string
	}
)
