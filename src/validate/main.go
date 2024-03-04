package main

import (
	fmt "fmt"
	gra "github.com/craterdog/go-grammar-framework/v2"
	osx "os"
)

// MAIN PROGRAM

func main() {
	// Validate the commandline arguments.
	if len(osx.Args) < 2 {
		fmt.Println("Usage: validate <grammar-file>")
		return
	}
	var grammarFile = osx.Args[1]

	// Validate the grammar file.
	var bytes, err = osx.ReadFile(grammarFile)
	if err != nil {
		panic(err)
	}
	var source = string(bytes)
	var parser = gra.Parser().Make()
	parser.ParseSource(source)
}
