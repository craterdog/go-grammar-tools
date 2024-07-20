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

var additionalClass = &additionalClass_{
	// Initialize class constants.
}

// Function

func Additional() AdditionalClassLike {
	return additionalClass
}

// CLASS METHODS

// Target

type additionalClass_ struct {
	// Define class constants.
}

// Constructors

func (c *additionalClass_) Make(component ComponentLike) AdditionalLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(component):
		panic("The component attribute is required by this class.")
	default:
		return &additional_{
			// Initialize instance attributes.
			class_: c,
			component_: component,
		}
	}
}

// INSTANCE METHODS

// Target

type additional_ struct {
	// Define instance attributes.
	class_ AdditionalClassLike
	component_ ComponentLike
}

// Attributes

func (v *additional_) GetClass() AdditionalClassLike {
	return v.class_
}

func (v *additional_) GetComponent() ComponentLike {
	return v.component_
}

// Private
