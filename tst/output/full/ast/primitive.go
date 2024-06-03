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

func (c *primitiveClass_) MakeWithCharacter(character string) PrimitiveLike {
	return &primitive_{
		// Initialize instance attributes.
		class_: c,
		character_: character,
	}
}

func (c *primitiveClass_) MakeWithText(text string) PrimitiveLike {
	return &primitive_{
		// Initialize instance attributes.
		class_: c,
		text_: text,
	}
}

func (c *primitiveClass_) MakeWithInteger(integer string) PrimitiveLike {
	return &primitive_{
		// Initialize instance attributes.
		class_: c,
		integer_: integer,
	}
}

func (c *primitiveClass_) MakeWithAnything(anything string) PrimitiveLike {
	return &primitive_{
		// Initialize instance attributes.
		class_: c,
		anything_: anything,
	}
}

// INSTANCE METHODS

// Target

type primitive_ struct {
	// Define instance attributes.
	class_ PrimitiveClassLike
	character_ string
	text_ string
	integer_ string
	anything_ string
}

// Attributes

func (v *primitive_) GetClass() PrimitiveClassLike {
	return v.class_
}

func (v *primitive_) GetCharacter() string {
	return v.character_
}

func (v *primitive_) GetText() string {
	return v.text_
}

func (v *primitive_) GetInteger() string {
	return v.integer_
}

func (v *primitive_) GetAnything() string {
	return v.anything_
}

// Private
