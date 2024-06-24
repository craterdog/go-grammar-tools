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

package ast

import ()

// CLASS ACCESS

// Reference

var primitiveClass = &primitiveClass_{
	// Initialize class constants.
}

// Function

func Primitive() PrimitiveClassLike {
	return primitiveClass
}

// CLASS METHODS

// Target

type primitiveClass_ struct {
	// Define class constants.
}

// Constructors

func (c *primitiveClass_) MakeWithInteger(integer string) PrimitiveLike {
	return &primitive_{
		// Initialize instance attributes.
		class_: c,
		integer_: integer,
	}
}

func (c *primitiveClass_) MakeWithRune(rune_ string) PrimitiveLike {
	return &primitive_{
		// Initialize instance attributes.
		class_: c,
		rune_: rune_,
	}
}

func (c *primitiveClass_) MakeWithText(text string) PrimitiveLike {
	return &primitive_{
		// Initialize instance attributes.
		class_: c,
		text_: text,
	}
}

// INSTANCE METHODS

// Target

type primitive_ struct {
	// Define instance attributes.
	class_ PrimitiveClassLike
	any_ any
	integer_ string
	rune_ string
	text_ string
}

// Attributes

func (v *primitive_) GetClass() PrimitiveClassLike {
	return v.class_
}

func (v *primitive_) GetAny() any {
	return v.any_
}

// Private
