/*
................................................................................
.    Copyright (c) 2009-2024 Crater Dog Technologies.  All Rights Reserved.    .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

package main

import (
	fmt "fmt"
	age "github.com/craterdog/go-grammar-framework/v4/cdsn/agent"
	osx "os"
)

// MAIN PROGRAM

func main() {
	// Validate the commandline arguments.
	if len(osx.Args) < 2 {
		fmt.Println("Usage: format <syntax-file>")
		return
	}
	var syntaxFile = osx.Args[1]

	// Parse the syntax file.
	var bytes, err = osx.ReadFile(syntaxFile)
	if err != nil {
		panic(err)
	}
	var source = string(bytes)
	var parser = age.Parser().Make()
	var syntax = parser.ParseSource(source)

	// Validate the syntax.
	var validator = age.Validator().Make()
	validator.ValidateSyntax(syntax)

	// Reformat the syntax file.
	var formatter = age.Formatter().Make()
	source = formatter.FormatSyntax(syntax)
	bytes = []byte(source)
	err = osx.WriteFile(syntaxFile, bytes, 0644)
	if err != nil {
		panic(err)
	}
}
