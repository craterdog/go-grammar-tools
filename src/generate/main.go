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
	var module, wiki, directory = retrieveArguments()
	var syntaxFile = directory + "Syntax.cdsn"
	if !pathExists(syntaxFile) {
		fmt.Printf(
			"The syntax file %v does not exist, so aborting...",
			syntaxFile,
		)
		return
	}
	var syntax = parseSyntax(directory)
	validateSyntax(syntax)
	if !pathExists(directory + "ast") {
		var astModel = generateAst(module, wiki, directory, syntax)
		validateModel(astModel)
		generateClasses(module, directory, astModel)
	}
	if !pathExists(directory + "grammar") {
		var grammarModel = generateGrammar(module, wiki, directory, syntax)
		validateModel(grammarModel)
		generateFormatter(module, wiki, directory, syntax)
		generateParser(module, wiki, directory, syntax)
		generateScanner(module, wiki, directory, syntax)
		generateToken(module, wiki, directory, syntax)
		generateValidator(module, wiki, directory, syntax)
	}
}

func generateGrammar(
	module string,
	wiki string,
	directory string,
	syntax gra.SyntaxLike,
) mod.ModelLike {
	var grammarDirectory = directory + "grammar/"
	var err = osx.MkdirAll(grammarDirectory, 0755)
	if err != nil {
		panic(err)
	}
	var grammarFile = grammarDirectory + "Package.go"
	var generator = gra.Generator()
	var grammarModel = generator.GenerateGrammar(module, wiki, syntax)
	var validator = mod.Validator()
	validator.ValidateModel(grammarModel)
	var formatter = mod.Formatter()
	var source = formatter.FormatModel(grammarModel)
	var bytes = []byte(source)
	err = osx.WriteFile(grammarFile, bytes, 0644)
	if err != nil {
		panic(err)
	}
	return grammarModel
}

func generateAst(
	module string,
	wiki string,
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
	var astModel = generator.GenerateAst(module, wiki, syntax)
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

func generateClasses(
	module string,
	directory string,
	astModel mod.ModelLike,
) {
	var generator = mod.Generator()
	var iterator = astModel.GetClasses().GetClasses().GetIterator()
	for iterator.HasNext() {
		var class = iterator.GetNext()
		var name = sts.ToLower(sts.TrimSuffix(
			class.GetDeclaration().GetName(),
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
	wiki string,
	directory string,
	syntax gra.SyntaxLike,
) {
	var generator = gra.Generator()
	var source = generator.GenerateFormatter(module, wiki, syntax)
	var bytes = []byte(source)
	var filename = directory + "grammar/formatter.go"
	var err = osx.WriteFile(filename, bytes, 0644)
	if err != nil {
		panic(err)
	}
}

func generateParser(
	module string,
	wiki string,
	directory string,
	syntax gra.SyntaxLike,
) {
	var generator = gra.Generator()
	var source = generator.GenerateParser(module, wiki, syntax)
	var bytes = []byte(source)
	var filename = directory + "grammar/parser.go"
	var err = osx.WriteFile(filename, bytes, 0644)
	if err != nil {
		panic(err)
	}
}

func generateScanner(
	module string,
	wiki string,
	directory string,
	syntax gra.SyntaxLike,
) {
	var generator = gra.Generator()
	var source = generator.GenerateScanner(module, wiki, syntax)
	var bytes = []byte(source)
	var filename = directory + "grammar/scanner.go"
	var err = osx.WriteFile(filename, bytes, 0644)
	if err != nil {
		panic(err)
	}
}

func generateToken(
	module string,
	wiki string,
	directory string,
	syntax gra.SyntaxLike,
) {
	var generator = gra.Generator()
	var source = generator.GenerateToken(module, wiki, syntax)
	var bytes = []byte(source)
	var filename = directory + "grammar/token.go"
	var err = osx.WriteFile(filename, bytes, 0644)
	if err != nil {
		panic(err)
	}
}

func generateValidator(
	module string,
	wiki string,
	directory string,
	syntax gra.SyntaxLike,
) {
	var generator = gra.Generator()
	var source = generator.GenerateValidator(module, wiki, syntax)
	var bytes = []byte(source)
	var filename = directory + "grammar/validator.go"
	var err = osx.WriteFile(filename, bytes, 0644)
	if err != nil {
		panic(err)
	}
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
	module string,
	wiki string,
	directory string,
) {
	if len(osx.Args) < 4 {
		fmt.Println("Usage: generate <module> <wiki> <directory>")
		osx.Exit(1)
	}
	module = osx.Args[1]
	wiki = osx.Args[2]
	directory = osx.Args[3]
	if !sts.HasSuffix(directory, "/") {
		directory += "/"
	}
	return module, wiki, directory
}

func validateModel(model mod.ModelLike) {
	var validator = mod.Validator()
	validator.ValidateModel(model)
}

func validateSyntax(syntax gra.SyntaxLike) {
	var validator = gra.Validator()
	validator.ValidateSyntax(syntax)
}
