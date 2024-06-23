/*
................................................................................
.                   Copyright (c) 2024.  All Rights Reserved.                  .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

package agent

import (
	ast "github.com/craterdog/example/example/ast"
	sts "strings"
)

// CLASS ACCESS

// Reference

var formatterClass = &formatterClass_{
	// Initialize the class constants.
	defaultMaximum_: 8,
}

// Function

func Formatter() FormatterClassLike {
	return formatterClass
}

// CLASS METHODS

// Target

type formatterClass_ struct {
	// Define the class constants.
	defaultMaximum_ uint
}

// Constants

func (c *formatterClass_) DefaultMaximum() uint {
	return c.defaultMaximum_
}

// Constructors

func (c *formatterClass_) Make() FormatterLike {
	return &formatter_{
		// Initialize the instance attributes.
		class_:   c,
		maximum_: c.defaultMaximum_,
	}
}

func (c *formatterClass_) MakeWithMaximum(maximum uint) FormatterLike {
	if maximum == 0 {
		maximum = c.defaultMaximum_
	}
	return &formatter_{
		// Initialize the instance attributes.
		class_:   c,
		maximum_: maximum,
	}
}

// INSTANCE METHODS

// Target

type formatter_ struct {
	// Define the instance attributes.
	class_   FormatterClassLike
	depth_   uint
	maximum_ uint
	result_  sts.Builder
}

// Attributes

func (v *formatter_) GetClass() FormatterClassLike {
	return v.class_
}

func (v *formatter_) GetDepth() uint {
	return v.depth_
}

func (v *formatter_) GetMaximum() uint {
	return v.maximum_
}

// Public

func (v *formatter_) FormatDocument(document ast.DocumentLike) string {
	v.formatDocument(document)
	return v.getResult()
}

// Private

func (v *formatter_) appendNewline() {
	var newline = "\n"
	var indentation = "\t"
	var level uint
	for ; level < v.depth_; level++ {
		newline += indentation
	}
	v.appendString(newline)
}

func (v *formatter_) appendString(s string) {
	v.result_.WriteString(s)
}

func (v *formatter_) formatDocument(document ast.DocumentLike) {
	// TBA - Add real method implementation.
	v.depth_++
	v.appendString("test")
	v.appendNewline()
	v.depth_--
}

func (v *formatter_) getResult() string {
	var result = v.result_.String()
	v.result_.Reset()
	return result
}
