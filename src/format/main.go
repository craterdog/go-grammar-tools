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

func main() {
	var syntaxFile = retrieveArguments()
	if !pathExists(syntaxFile) {
		fmt.Printf(
			"The syntax file %v does not exist, so aborting...",
			syntaxFile,
		)
		return
	}
	var syntax = parseSyntax(syntaxFile)
	validateSyntax(syntax)
	reformatSyntax(syntaxFile, syntax)
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

func pathExists(path string) bool {
	var _, err = osx.Stat(path)
	if err == nil {
		return true
	}
	if osx.IsNotExist(err) {
		return false
	}
	panic(err)
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

func retrieveArguments() (syntaxFile string) {
	if len(osx.Args) < 2 {
		fmt.Println("Usage: format <syntax-file>")
		osx.Exit(1)
	}
	syntaxFile = osx.Args[1]
	return syntaxFile
}

func validateSyntax(syntax gra.SyntaxLike) {
	var validator = gra.Validator()
	validator.ValidateSyntax(syntax)
}
