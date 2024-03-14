/*
................................................................................
.    Copyright (c) 2009-2024 Crater Dog Technologies™.  All Rights Reserved.   .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See http://opensource.org/licenses/MIT)                        .
................................................................................
*/

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
