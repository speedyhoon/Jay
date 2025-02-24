# `jay`

The `jay` command line tool scans `.go` files for exported structs and generates methods `.MarshalJ()` and `.UnmarshalJ()` to serialize structs into the [Jay serialization format](https://github.com/speedyhoon/jay).
`jay` is a wrapper around the [jay/generate](../../README.md) package, 
providing optional command line flags and traverses directories in search for `.go` files.

## Install

```shell
go install github.com/speedyhoon/jay/generate/cmd/jay
```

## Usage
Whole directory:
```shell
cd my_project_path
jay
```

Single file:
```shell
cd my_project_path
jay my_file.go
```

## Flags

`-32` Force 32-bit output for `ints` & `uints`. _Defaults to your system's architecture._

`-d` Debug mode, always write to disk. _Default: `false`_

`-fi` Fixed int size. _Default: `true`_

`-fu` Fixed uint size. _Default: `true`_

`-o` Output file.  _Default: `jay.go`_

`-v` Verbose output. _Default: `false`_

`-p` Generates pointer `.MarshalJ()` methods instead of functions. _Default: `false`_

`-s` Search Go test files for exported structs too. _Default: `false`_

`-t` Don't generate Go test files. _Default: `false`_

`-m` Don't generate `MarshalJ()` functions or methods. _Default: `false`_

`-u` Don't generate `UnmarshalJ()` function. _Default: `false`_

`-y` Exclusive list of comma separated types to generate marshalling and/or unmarshalling for. _Default is to process all exported types._ <br>
         For example: `-y Vet,animal.Cat,animal.Cow` will process locally defined `Vet struct` along with `Cat` & `Cow` in imported package `animal`.

## When to regenerate code
How often does `jay` need to be executed?
* Regenerate when:
  * Any exported struct is added, deleted or renamed,
  * Any embedded struct with an exported field name is added, deleted or renamed,
  * Any exported field is added, deleted or renamed,
  * Any exported field type is changed,
  * Any exported field tag is added, deleted or altered,
  * Any field is toggled between exported or private.
* Regenerate isn't required when:
  * Struct field order has changed,
  * Unexported structs or fields are added, deleted or renamed,
  * Unexported field types are modified.