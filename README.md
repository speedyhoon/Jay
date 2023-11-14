# Jay
A fast serialisation package alternative to JSON, MessagePack and Bebop for [Go](https://go.dev). It outputs non-readable `[]byte` ideal for high throughput, using much less bandwidth than other options. 

Jay is similar to [Bebop]() where the serialisation functions are generated before compiling. This makes serialising and deserialising much faster than other options that use reflection during runtime like JSON.

Go files can be scanned by importing `jay` or using the command line tool to generate `Marshal` and `Unmarshal` functions for exported struct fields.

## Origins
Credit to the creators of [Bebop]() for creating an extremely fast serialisation format.
However, instead of using `.bop` schema files, `Jay` uses Go tags ``` `j:""` ``` to generate `Marshal` and `Unmarshal` functions for exported struct fields.

[//]: # (`Jay` differs from Bebop to do away with the)
[//]: # (Using the command line tool,)
[//]: # (to and any struct that has exported fields with a tag containing `j:""` have a `Marshal` and `Unmarshal` function generated.)
[//]: # (without the `.bop` files.)

###### Name
The name `jay` gives tribute to a 17-year-old laptop with a stuck `j` key on the keyboard. Every boot looks like this: