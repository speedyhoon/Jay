package main

import (
	"github.com/speedyhoon/jay/generate"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile)

	//instance := Car{ID: 42, Name: "Hello"}
	//byt, _ := json.Marshal(instance)
	//log.Printf("%s\n", byt)

	//src, err := generate.ProcessFile("cmd/main.go", nil)
	//log.Printf("src output:\n%s", src)
	// TODO add filename as cmdline flag
	err := generate.ProcessWrite("example/main.go", nil, generate.Option{
		FixedIntSize:  true,
		FixedUintSize: true,
	})
	if err != nil {
		log.Println(err)
	}
}

//func main() {
/*instance := Car{ID: 42, Name: "Hello"}

// Get the type of the struct
structType := reflect.TypeOf(instance)


println(gg)
// Loop through the fields of the struct
for i := 0; i < structType.NumField(); i++ {
	field := structType.Field(i)

	// Get the tags for the field
	jsonTag := field.Tag.Get("json")
	customTag := field.Tag.Get("customTag")

	// Print the tags
	fmt.Printf("Field %s: json=%s, customTag=%s\n", field.Name, jsonTag, customTag)
}*/

//return
