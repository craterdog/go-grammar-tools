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

var intrinsicClass = &intrinsicClass_{
	// Initialize class constants.
}

// Function

func Intrinsic() IntrinsicClassLike {
	return intrinsicClass
}

// CLASS METHODS

// Target

type intrinsicClass_ struct {
	// Define class constants.
}

// Constructors

func (c *intrinsicClass_) Make(any_ any) IntrinsicLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(any_):
		panic("The any attribute is required by this class.")
	default:
		return &intrinsic_{
			// Initialize instance attributes.
			class_: c,
			any_: any_,
		}
	}
}

// INSTANCE METHODS

// Target

type intrinsic_ struct {
	// Define instance attributes.
	class_ IntrinsicClassLike
	any_ any
}

// Attributes

func (v *intrinsic_) GetClass() IntrinsicClassLike {
	return v.class_
}

func (v *intrinsic_) GetAny() any {
	return v.any_
}

// Private
