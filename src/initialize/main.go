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
	if len(osx.Args) < 4 {
		fmt.Println(
			"Usage: initialize <package-directory> <notation-name> <copyright>",
		)
		return
	}
	var directory = osx.Args[1]
	var copyright = osx.Args[2]
	var notation = osx.Args[3]

	// Create a new syntax file.
	var generator = age.Generator().Make()
	generator.CreateSyntax(directory, notation, copyright)
}
