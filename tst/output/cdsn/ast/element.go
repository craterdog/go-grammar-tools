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

var elementClass = &elementClass_{
	// Initialize class constants.
}

// Function

func Element() ElementClassLike {
	return elementClass
}

// CLASS METHODS

// Target

type elementClass_ struct {
	// Define class constants.
}

// Constructors

func (c *elementClass_) MakeWithGrouped(grouped GroupedLike) ElementLike {
	return &element_{
		// Initialize instance attributes.
		class_: c,
		grouped_: grouped,
	}
}

func (c *elementClass_) MakeWithFiltered(filtered FilteredLike) ElementLike {
	return &element_{
		// Initialize instance attributes.
		class_: c,
		filtered_: filtered,
	}
}

func (c *elementClass_) MakeWithBounded(bounded BoundedLike) ElementLike {
	return &element_{
		// Initialize instance attributes.
		class_: c,
		bounded_: bounded,
	}
}

func (c *elementClass_) MakeWithIntrinsic(intrinsic string) ElementLike {
	return &element_{
		// Initialize instance attributes.
		class_: c,
		intrinsic_: intrinsic,
	}
}

func (c *elementClass_) MakeWithLowercase(lowercase string) ElementLike {
	return &element_{
		// Initialize instance attributes.
		class_: c,
		lowercase_: lowercase,
	}
}

func (c *elementClass_) MakeWithLiteral(literal string) ElementLike {
	return &element_{
		// Initialize instance attributes.
		class_: c,
		literal_: literal,
	}
}

// INSTANCE METHODS

// Target

type element_ struct {
	// Define instance attributes.
	class_ ElementClassLike
	any_ any
	grouped_ GroupedLike
	filtered_ FilteredLike
	bounded_ BoundedLike
	intrinsic_ string
	lowercase_ string
	literal_ string
}

// Attributes

func (v *element_) GetClass() ElementClassLike {
	return v.class_
}

func (v *element_) GetAny() any {
	return v.any_
}

// Private
