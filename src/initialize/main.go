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
	sts "strings"
)

// MAIN PROGRAM

func main() {
	var directory, name, copyright = retrieveArguments()
	var syntax = createSyntax(name, copyright)
	saveSyntax(directory, syntax)
}

func retrieveArguments() (
	directory string,
	name string,
	copyright string,
) {
	if len(osx.Args) < 4 {
		fmt.Println(
			"Usage: initialize <directory> <name> <copyright>",
		)
		return
	}
	directory = osx.Args[1]
	if !sts.HasSuffix(directory, "/") {
		directory += "/"
	}
	name = osx.Args[2]
	copyright = osx.Args[3]
	return directory, name, copyright
}

func createSyntax(
	name string,
	copyright string,
) gra.SyntaxLike {
	var generator = gra.Generator()
	var syntax = generator.CreateSyntax(name, copyright)
	return syntax
}

func saveSyntax(directory string, syntax gra.SyntaxLike) {
	var err = osx.MkdirAll(directory, 0755)
	if err != nil {
		panic(err)
	}
	var syntaxFile = directory + "Syntax.cdsn"
	fmt.Printf(
		"The syntax file %q does not yet exist.\n\tCreating it...\n",
		syntaxFile,
	)
	var formatter = gra.Formatter()
	var source = formatter.FormatSyntax(syntax)
	var bytes = []byte(source)
	err = osx.WriteFile(syntaxFile, bytes, 0644)
	if err != nil {
		panic(err)
	}
}
