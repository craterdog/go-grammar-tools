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
		fmt.Println("Usage: format <grammar-file>")
		return
	}
	var grammarFile = osx.Args[1]

	// Parse the grammar file.
	var bytes, err = osx.ReadFile(grammarFile)
	if err != nil {
		panic(err)
	}
	var source = string(bytes)
	var parser = gra.Parser().Make()
	var grammar = parser.ParseSource(source)

	// Reformat the grammar file.
	var formatter = gra.Formatter().Make()
	source = formatter.FormatGrammar(grammar)
	bytes = []byte(source)
	err = osx.WriteFile(grammarFile, bytes, 0644)
	if err != nil {
		panic(err)
	}
}
