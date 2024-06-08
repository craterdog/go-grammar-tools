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

func main() {
	var directory, name, copyright = retrieveArguments()
	var syntaxFile = directory + "Syntax.cdsn"
	if pathExists(syntaxFile) {
		fmt.Printf(
			"The syntax file %v already exists, so aborting...",
			syntaxFile,
		)
		return
	}
	var syntax = createSyntax(name, copyright)
	saveSyntax(directory, syntax)
}

func createSyntax(
	name string,
	copyright string,
) gra.SyntaxLike {
	var generator = gra.Generator()
	var syntax = generator.CreateSyntax(name, copyright)
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

func retrieveArguments() (
	directory string,
	name string,
	copyright string,
) {
	if len(osx.Args) < 4 {
		fmt.Println(
			"Usage: initialize <directory> <name> <copyright>",
		)
		osx.Exit(1)
	}
	directory = osx.Args[1]
	if !sts.HasSuffix(directory, "/") {
		directory += "/"
	}
	name = osx.Args[2]
	copyright = osx.Args[3]
	return directory, name, copyright
}

func saveSyntax(directory string, syntax gra.SyntaxLike) {
	var err = osx.MkdirAll(directory, 0755)
	if err != nil {
		panic(err)
	}
	var syntaxFile = directory + "Syntax.cdsn"
	var formatter = gra.Formatter()
	var source = formatter.FormatSyntax(syntax)
	var bytes = []byte(source)
	err = osx.WriteFile(syntaxFile, bytes, 0644)
	if err != nil {
		panic(err)
	}
}
