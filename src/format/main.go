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
	gra "github.com/craterdog/go-grammar-framework/v4"
	osx "os"
)

// MAIN PROGRAM

func main() {
	var syntaxFile = retrieveArguments()
	var syntax = parseSyntax(syntaxFile)
	validateSyntax(syntax)
	reformatSyntax(syntaxFile, syntax)
}

func retrieveArguments() (syntaxFile string) {
	if len(osx.Args) < 2 {
		fmt.Println("Usage: generate <syntax-file>")
		return
	}
	syntaxFile = osx.Args[1]
	return syntaxFile
}

func parseSyntax(syntaxFile string) gra.SyntaxLike {
	var bytes, err = osx.ReadFile(syntaxFile)
	if err != nil {
		panic(err)
	}
	var source = string(bytes)
	var parser = gra.Parser()
	var syntax = parser.ParseSource(source)
	return syntax
}

func validateSyntax(syntax gra.SyntaxLike) {
	var validator = gra.Validator()
	validator.ValidateSyntax(syntax)
}

func reformatSyntax(
	syntaxFile string,
	syntax gra.SyntaxLike,
) {
	var formatter = gra.Formatter()
	var source = formatter.FormatSyntax(syntax)
	var bytes = []byte(source)
	var err = osx.WriteFile(syntaxFile, bytes, 0644)
	if err != nil {
		panic(err)
	}
}
