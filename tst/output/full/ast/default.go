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

var defaultClass = &defaultClass_{
	// Initialize class constants.
}

// Function

func Default() DefaultClassLike {
	return defaultClass
}

// CLASS METHODS

// Target

type defaultClass_ struct {
	// Define class constants.
}

// Constructors

func (c *defaultClass_) Make() DefaultLike {
	return &default_{
		// Initialize instance attributes.
		class_: c,
	}
}

// INSTANCE METHODS

// Target

type default_ struct {
	// Define instance attributes.
	class_ DefaultClassLike
}

// Attributes

func (v *default_) GetClass() DefaultClassLike {
	return v.class_
}

// Private
