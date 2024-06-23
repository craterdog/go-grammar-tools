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

package ast

import ()

// CLASS ACCESS

// Reference

var identifierClass = &identifierClass_{
	// Initialize class constants.
}

// Function

func Identifier() IdentifierClassLike {
	return identifierClass
}

// CLASS METHODS

// Target

type identifierClass_ struct {
	// Define class constants.
}

// Constructors

func (c *identifierClass_) MakeWithLowercase(lowercase string) IdentifierLike {
	return &identifier_{
		// Initialize instance attributes.
		class_: c,
		lowercase_: lowercase,
	}
}

func (c *identifierClass_) MakeWithUppercase(uppercase string) IdentifierLike {
	return &identifier_{
		// Initialize instance attributes.
		class_: c,
		uppercase_: uppercase,
	}
}

// INSTANCE METHODS

// Target

type identifier_ struct {
	// Define instance attributes.
	class_ IdentifierClassLike
	any_ any
	lowercase_ string
	uppercase_ string
}

// Attributes

func (v *identifier_) GetClass() IdentifierClassLike {
	return v.class_
}

func (v *identifier_) GetAny() any {
	return v.any_
}

// Private
