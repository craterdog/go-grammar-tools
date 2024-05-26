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
	mod "github.com/craterdog/go-model-framework/v4"
	osx "os"
	sts "strings"
)

// MAIN PROGRAM

func main() {
	var directory = retrieveArguments()
	var syntax = parseSyntax(directory)
	validateSyntax(syntax)
	var astModel = generateAST(directory, syntax)
	generateClasses(directory, astModel)
	var agentModel = generateAgent(directory, syntax)
	generateFormatter(directory, agentModel)
	generateParser(directory, agentModel)
	generateScanner(directory, agentModel)
	generateToken(directory, agentModel)
	generateValidator(directory, agentModel)
}

func retrieveArguments() (directory string) {
	if len(osx.Args) < 2 {
		fmt.Println("Usage: generate <directory>")
		return
	}
	directory = osx.Args[1]
	if !sts.HasSuffix(directory, "/") {
		directory += "/"
	}
	return directory
}

func validateSyntax(syntax gra.SyntaxLike) {
	var validator = gra.Validator()
	validator.ValidateSyntax(syntax)
}

func parseSyntax(directory string) gra.SyntaxLike {
	var syntaxFile = directory + "Syntax.cdsn"
	var bytes, err = osx.ReadFile(syntaxFile)
	if err != nil {
		panic(err)
	}
	var source = string(bytes)
	var parser = gra.Parser()
	var syntax = parser.ParseSource(source)
	return syntax
}

func generateAST(directory string, syntax gra.SyntaxLike) mod.ModelLike {
	var astDirectory = directory + "ast/"
	var err = osx.MkdirAll(astDirectory, 0755)
	if err != nil {
		panic(err)
	}
	var astFile = astDirectory + "Package.go"
	var generator = gra.Generator()
	var astModel = generator.GenerateAST(syntax)
	var formatter = mod.Formatter()
	var source = formatter.FormatModel(astModel)
	var bytes = []byte(source)
	err = osx.WriteFile(astFile, bytes, 0644)
	if err != nil {
		panic(err)
	}
	return astModel
}

func generateAgent(directory string, syntax gra.SyntaxLike) mod.ModelLike {
	var agentDirectory = directory + "agent/"
	var err = osx.MkdirAll(agentDirectory, 0755)
	if err != nil {
		panic(err)
	}
	var agentFile = agentDirectory + "Package.go"
	var generator = gra.Generator()
	var agentModel = generator.GenerateAgent(syntax)
	var formatter = mod.Formatter()
	var source = formatter.FormatModel(agentModel)
	var bytes = []byte(source)
	err = osx.WriteFile(agentFile, bytes, 0644)
	if err != nil {
		panic(err)
	}
	return agentModel
}

func generateClasses(
	directory string,
	astModel mod.ModelLike,
) {
}

func generateFormatter(
	directory string,
	agentModel mod.ModelLike,
) {
	var generator = gra.Generator()
	var source = generator.GenerateFormatter(agentModel)
	var bytes = []byte(source)
	var filename = directory + "agent/formatter.go"
	var err = osx.WriteFile(filename, bytes, 0644)
	if err != nil {
		panic(err)
	}
}

func generateParser(
	directory string,
	agentModel mod.ModelLike,
) {
	var generator = gra.Generator()
	var source = generator.GenerateParser(agentModel)
	var bytes = []byte(source)
	var filename = directory + "agent/parser.go"
	var err = osx.WriteFile(filename, bytes, 0644)
	if err != nil {
		panic(err)
	}
}

func generateScanner(
	directory string,
	agentModel mod.ModelLike,
) {
	var generator = gra.Generator()
	var source = generator.GenerateScanner(agentModel)
	var bytes = []byte(source)
	var filename = directory + "agent/scanner.go"
	var err = osx.WriteFile(filename, bytes, 0644)
	if err != nil {
		panic(err)
	}
}

func generateToken(
	directory string,
	agentModel mod.ModelLike,
) {
	var generator = gra.Generator()
	var source = generator.GenerateToken(agentModel)
	var bytes = []byte(source)
	var filename = directory + "agent/token.go"
	var err = osx.WriteFile(filename, bytes, 0644)
	if err != nil {
		panic(err)
	}
}

func generateValidator(
	directory string,
	agentModel mod.ModelLike,
) {
	var generator = gra.Generator()
	var source = generator.GenerateValidator(agentModel)
	var bytes = []byte(source)
	var filename = directory + "agent/validator.go"
	var err = osx.WriteFile(filename, bytes, 0644)
	if err != nil {
		panic(err)
	}
}
