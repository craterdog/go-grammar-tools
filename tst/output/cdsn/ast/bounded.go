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

import (
	col "github.com/craterdog/go-collection-framework/v4/collection"
)

// CLASS ACCESS

// Reference

var boundedClass = &boundedClass_{
	// Initialize class constants.
}

// Function

func Bounded() BoundedClassLike {
	return boundedClass
}

// CLASS METHODS

// Target

type boundedClass_ struct {
	// Define class constants.
}

// Constructors

func (c *boundedClass_) MakeWithAttributes(
	initial InitialLike,
	extents col.ListLike[ExtentLike],
) BoundedLike {
	return &bounded_{
		// Initialize instance attributes.
		class_: c,
		initial_: initial,
		extents_: extents,
	}
}

// INSTANCE METHODS

// Target

type bounded_ struct {
	// Define instance attributes.
	class_ BoundedClassLike
	initial_ InitialLike
	extents_ col.ListLike[ExtentLike]
}

// Attributes

func (v *bounded_) GetClass() BoundedClassLike {
	return v.class_
}

func (v *bounded_) GetInitial() InitialLike {
	return v.initial_
}

func (v *bounded_) GetExtents() col.ListLike[ExtentLike] {
	return v.extents_
}

// Private
