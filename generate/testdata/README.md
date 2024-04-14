# testdata

### Directory Suffixes
#### Built-ins
* `*` No suffix, built in type use directly. `int16`
* `*Alias` Type alias. `type small = int16`
* `*Def` Type definition. `type small int16`

#### Slices
* `*s` A slice of a built-in type. `[]int16` used directly
* `*sAlias` Alias type used as a slice. `type small = int16` used as `[]small`
* `*sAliased` An aliased slice type. `type smalls = []int16` used as `smalls`
* `*sDef` A new slice type definition. `type smalls []int16` used as `smalls`
* `*sDefs` A new type definition used as a slice. `type small int16` used as `[]small` _(not yet supported)_
