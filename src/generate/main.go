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

func main() {
	var module, directory = retrieveArguments()
	var syntax = parseSyntax(directory)
	validateSyntax(syntax)
	var astModel = generateAST(module, directory, syntax)
	validateModel(astModel)
	generateClasses(module, directory, astModel)
	var agentModel = generateAgent(module, directory, syntax)
	validateModel(agentModel)
	generateFormatter(module, directory, syntax, agentModel)
	generateParser(module, directory, syntax, agentModel)
	generateScanner(module, directory, syntax, agentModel)
	generateToken(module, directory, syntax, agentModel)
	generateValidator(module, directory, syntax, agentModel)
}

func retrieveArguments() (
	module string,
	directory string,
) {
	if len(osx.Args) < 3 {
		fmt.Println("Usage: generate <module> <directory>")
		osx.Exit(1)
	}
	module = osx.Args[1]
	directory = osx.Args[2]
	if !sts.HasSuffix(directory, "/") {
		directory += "/"
	}
	return module, directory
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

func generateAST(
	module string,
	directory string,
	syntax gra.SyntaxLike,
) mod.ModelLike {
	var astDirectory = directory + "ast/"
	var err = osx.MkdirAll(astDirectory, 0755)
	if err != nil {
		panic(err)
	}
	var astFile = astDirectory + "Package.go"
	var generator = gra.Generator()
	var astModel = generator.GenerateAST(module, syntax)
	var validator = mod.Validator()
	validator.ValidateModel(astModel)
	var formatter = mod.Formatter()
	var source = formatter.FormatModel(astModel)
	var bytes = []byte(source)
	err = osx.WriteFile(astFile, bytes, 0644)
	if err != nil {
		panic(err)
	}
	return astModel
}

func generateAgent(
	module string,
	directory string,
	syntax gra.SyntaxLike,
) mod.ModelLike {
	var agentDirectory = directory + "agent/"
	var err = osx.MkdirAll(agentDirectory, 0755)
	if err != nil {
		panic(err)
	}
	var agentFile = agentDirectory + "Package.go"
	var generator = gra.Generator()
	var agentModel = generator.GenerateAgent(module, syntax)
	var validator = mod.Validator()
	validator.ValidateModel(agentModel)
	var formatter = mod.Formatter()
	var source = formatter.FormatModel(agentModel)
	var bytes = []byte(source)
	err = osx.WriteFile(agentFile, bytes, 0644)
	if err != nil {
		panic(err)
	}
	return agentModel
}

func validateModel(model mod.ModelLike) {
	var validator = mod.Validator()
	validator.ValidateModel(model)
}

func generateClasses(
	module string,
	directory string,
	astModel mod.ModelLike,
) {
	var generator = mod.Generator()
	var iterator = astModel.GetClasses().GetIterator()
	for iterator.HasNext() {
		var class = iterator.GetNext()
		var name = sts.ToLower(sts.TrimSuffix(
			class.GetDeclaration().GetIdentifier(),
			"ClassLike",
		))
		var source = generator.GenerateClass(astModel, name)
		var bytes = []byte(source)
		var filename = directory + "ast/" + name + ".go"
		var err = osx.WriteFile(filename, bytes, 0644)
		if err != nil {
			panic(err)
		}
	}
}

func generateFormatter(
	module string,
	directory string,
	syntax gra.SyntaxLike,
	agentModel mod.ModelLike,
) {
	var generator = gra.Generator()
	var source = generator.GenerateFormatter(module, syntax, agentModel)
	var bytes = []byte(source)
	var filename = directory + "agent/formatter.go"
	var err = osx.WriteFile(filename, bytes, 0644)
	if err != nil {
		panic(err)
	}
}

func generateParser(
	module string,
	directory string,
	syntax gra.SyntaxLike,
	agentModel mod.ModelLike,
) {
	var generator = gra.Generator()
	var source = generator.GenerateParser(module, syntax, agentModel)
	var bytes = []byte(source)
	var filename = directory + "agent/parser.go"
	var err = osx.WriteFile(filename, bytes, 0644)
	if err != nil {
		panic(err)
	}
}

func generateScanner(
	module string,
	directory string,
	syntax gra.SyntaxLike,
	agentModel mod.ModelLike,
) {
	var generator = gra.Generator()
	var source = generator.GenerateScanner(module, syntax, agentModel)
	var bytes = []byte(source)
	var filename = directory + "agent/scanner.go"
	var err = osx.WriteFile(filename, bytes, 0644)
	if err != nil {
		panic(err)
	}
}

func generateToken(
	module string,
	directory string,
	syntax gra.SyntaxLike,
	agentModel mod.ModelLike,
) {
	var generator = gra.Generator()
	var source = generator.GenerateToken(module, syntax, agentModel)
	var bytes = []byte(source)
	var filename = directory + "agent/token.go"
	var err = osx.WriteFile(filename, bytes, 0644)
	if err != nil {
		panic(err)
	}
}

func generateValidator(
	module string,
	directory string,
	syntax gra.SyntaxLike,
	agentModel mod.ModelLike,
) {
	var generator = gra.Generator()
	var source = generator.GenerateValidator(module, syntax, agentModel)
	var bytes = []byte(source)
	var filename = directory + "agent/validator.go"
	var err = osx.WriteFile(filename, bytes, 0644)
	if err != nil {
		panic(err)
	}
}
