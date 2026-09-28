package genjay

import (
	"errors"
	"fmt"
	"go/types"
	"log"
	"strconv"

	"golang.org/x/tools/go/packages"
)

var resolvedPackages = map[string][]*packages.Package{}

func resolveImportedTypes(importPath, identName string, f *field, parentTypes *prevTypes) (err error) {
	packs, ok := resolvedPackages[importPath]
	if !ok {
		packs, err = packages.Load(&packages.Config{Mode: packages.LoadAllSyntax}, importPath)
		if err != nil || len(packs) == 0 {
			return fmt.Errorf("could not load import package %s: %w", identName, err)
		}
		resolvedPackages[importPath] = packs
	}

	for _, pkg := range packs {
		obj := pkg.Types.Scope().Lookup(identName)
		if obj == nil {
			// Object not found in this package.
			continue
		}

		ok := typesTraverse(f, obj.Type(), parentTypes)
		if !ok {
			log.Printf("type %s not a built-in: %s", obj.Name(), f.typ)
			return fmt.Errorf("could not find type %s in %s", identName, importPath)
		}

		return nil
	}

	return errors.New("object not found")
}

func typesTraverse(f *field, obj types.Type, parentTypes *prevTypes) (ok bool) {
	switch z := obj.(type) {
	case *types.Named:
		return typesTraverse(f, z.Underlying(), parentTypes)
	case *types.Basic:
		switch z.Kind() {
		case types.Bool, types.Int, types.Int8, types.Int16, types.Int32, types.Int64, types.Uint, types.Uint8, types.Uint16, types.Uint32, types.Uint64, types.Uintptr, types.Float32, types.Float64, types.Complex64, types.Complex128, types.String:
			f.typ = z.String()
			return true
		}
	case *types.Slice:
		parentTypes.add(SLICE)
		f.arraySize = typeSlice
		ok = typesTraverse(f, z.Elem(), parentTypes)
		f.arrayType = f.typ
		f.typ = "[]" + f.typ
		f.arrayDepth++
		return f.arrayDepth == 1
	case *types.Array:
		f.arraySize = int(z.Len())
		if f.arraySize < 1 {
			return false
		}
		f.arrayDepth++
		if f.arrayDepth > 1 {
			return false
		}
		parentTypes.add(ARRAY)

		ok = typesTraverse(f, z.Elem(), parentTypes)
		f.arrayType = f.typ
		f.typ = "[]" + f.typ
		f.marshal.qtyVar = multiplier(strconv.Itoa(f.arraySize))
		f.unmarshal.qtyVar = f.marshal.qtyVar
		return ok
	}

	return false
}
