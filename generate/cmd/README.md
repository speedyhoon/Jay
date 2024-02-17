# jay

Command line tool

## Install

```shell
go install github.com/speedyhoon/jay/generate/cmd
```

## Usage

```shell
cd <<my_project_path>>
jay
```

## Flags

`-32` Force 32-bit output for ints & uints. Defaults to this systems 32-bit or 64-bit architecture.

`-fi` Fixed int size. _default: `true`_

`-fu` Fixed uint size. _default: `true`_

`-o`  Output file.  _default: `jay.go`_

`-v`  Verbose output. _default: `false`_

`-s`  Search Go test files for exported structs too. _default: `false`_

`-t`  Don't generate Go test files. _default: `false`_

`-m`  Don't generate MarshalJ() function. _default: `false`_

`-u`  Don't generate UnmarshalJ() function. _default: `false`_