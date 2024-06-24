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

func (c *primitiveClass_) Make(any any) PrimitiveLike {
	return &primitive_{
		// Initialize instance attributes.
		class_: c,
	}
}

// INSTANCE METHODS

// Target

type primitive_ struct {
	// Define instance attributes.
	class_ PrimitiveClassLike
	any_ any
}

// Attributes

func (v *primitive_) GetClass() PrimitiveClassLike {
	return v.class_
}

func (v *primitive_) GetAny() any {
	return v.any_
}

// Private
