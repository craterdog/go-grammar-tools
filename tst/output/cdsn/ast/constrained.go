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

var constrainedClass = &constrainedClass_{
	// Initialize class constants.
}

// Function

func Constrained() ConstrainedClassLike {
	return constrainedClass
}

// CLASS METHODS

// Target

type constrainedClass_ struct {
	// Define class constants.
}

// Constructors

func (c *constrainedClass_) MakeWithAttributes(
	minimum MinimumLike,
	maximums col.ListLike[MaximumLike],
) ConstrainedLike {
	return &constrained_{
		// Initialize instance attributes.
		class_: c,
		minimum_: minimum,
		maximums_: maximums,
	}
}

// INSTANCE METHODS

// Target

type constrained_ struct {
	// Define instance attributes.
	class_ ConstrainedClassLike
	minimum_ MinimumLike
	maximums_ col.ListLike[MaximumLike]
}

// Attributes

func (v *constrained_) GetClass() ConstrainedClassLike {
	return v.class_
}

func (v *constrained_) GetMinimum() MinimumLike {
	return v.minimum_
}

func (v *constrained_) GetMaximums() col.ListLike[MaximumLike] {
	return v.maximums_
}

// Private
