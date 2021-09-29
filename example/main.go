package main

import (
	"fmt"
	"go/importer"
)

func main() {
	pkg, err := importer.Default().Import("time")
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		return
	}

	scope := pkg.Scope()
	for _, name := range scope.Names() {
		if name == "Time" {
			obj := scope.Lookup(name)
			fmt.Printf("%#v\n", obj.Type())
		}
	}
}
