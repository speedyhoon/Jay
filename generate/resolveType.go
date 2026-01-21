package generate

import (
	"errors"
	"fmt"
	"golang.org/x/tools/go/packages"
	"log"
)

func resolveImportedTypes(importPath, identName string) (typ string, err error) {
	packs, err := packages.Load(&packages.Config{Mode: packages.LoadAllSyntax}, importPath)
	if err != nil || len(packs) == 0 {
		return "", fmt.Errorf("could not load import package %s: %w", identName, err)
	}

	for _, pkg := range packs {
		obj := pkg.Types.Scope().Lookup(identName)
		if obj == nil {
			// Object not found in this package.
			continue
		}

		typ = obj.Type().Underlying().String()
		if isBuiltIn(typ) {
			return typ, nil
		}

		log.Printf("type %s not a built-in: %s", obj.Name(), typ)
	}

	return "", errors.New("object not found")
}
