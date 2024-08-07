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

import (
	col "github.com/craterdog/go-collection-framework/v4"
)

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

func (c *componentClass_) Make(any_ any) ComponentLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(any_):
		panic("The any attribute is required by this class.")
	default:
		return &component_{
			// Initialize instance attributes.
			class_: c,
			any_: any_,
		}
	}
}

// INSTANCE METHODS

// Target

type component_ struct {
	// Define instance attributes.
	class_ ComponentClassLike
	any_ any
}

// Attributes

func (v *component_) GetClass() ComponentClassLike {
	return v.class_
}

func (v *component_) GetAny() any {
	return v.any_
}

// Private
