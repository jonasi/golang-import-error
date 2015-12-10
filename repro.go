package main

import (
	"go/importer"
)

func main() {
	_, err := importer.Default().Import("one")
	println("ERROR", err.Error())
}
