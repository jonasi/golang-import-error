package main

import (
	"go/importer"
)

func main() {
	imp("two")
	imp("one")
}

func imp(pkg string) {
	println("IMPORTING", pkg)
	_, err := importer.Default().Import(pkg)

	if err == nil {
		println("SUCCESS")
	} else {
		println("ERROR:", err.Error())
	}
}
