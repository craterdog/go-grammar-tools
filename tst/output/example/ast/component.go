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

var componentClass = &componentClass_{
	// Initialize class constants.
}

// Function

func Component() ComponentClassLike {
	return componentClass
}

// CLASS METHODS

// Target

type componentClass_ struct {
	// Define class constants.
}

// Constructors

func (c *componentClass_) MakeWithPrimitive(primitive PrimitiveLike) ComponentLike {
	return &component_{
		// Initialize instance attributes.
		class_: c,
		primitive_: primitive,
	}
}

func (c *componentClass_) MakeWithList(list ListLike) ComponentLike {
	return &component_{
		// Initialize instance attributes.
		class_: c,
		list_: list,
	}
}

// INSTANCE METHODS

// Target

type component_ struct {
	// Define instance attributes.
	class_ ComponentClassLike
	any_ any
	primitive_ PrimitiveLike
	list_ ListLike
}

// Attributes

func (v *component_) GetClass() ComponentClassLike {
	return v.class_
}

func (v *component_) GetAny() any {
	return v.any_
}

// Private
