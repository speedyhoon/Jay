# Dynamic content

`jay` doesn't encode any schema with the data like JSON. It is a strict format that once compiled can't be dynamically altered at runtime.

However, there are many ways to handle dynamic content. For example, if we have two structs `Car` and `Pet` and want to send them to the same endpoint URL in different network requests. One way to differentiate which type was received is to embed the marshalled `[]byte` into another struct like:

```go
type Message struct {
	Type uint8
	Data []byte
}

type Pet struct {
	Name    string
	Species string
	Breed   string
}

data := Message{
	Type: 2,
	Data: Pet{
		Name:    "Punchy",
		Species: "Dog",
	}.MarshalJ(),
}
fmt.Println(data.MarshalJ())
```

For a complete solution see `decode(src []byte)` in [main.go](main.go#L39) that handles multiple types.