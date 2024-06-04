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
	col "github.com/craterdog/go-collection-framework/v4/collection"
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

func (c *componentClass_) MakeWithDefault(default_ DefaultLike) ComponentLike {
	return &component_{
		// Initialize instance attributes.
		class_: c,
		default_: default_,
	}
}

func (c *componentClass_) MakeWithPrimitive(primitive PrimitiveLike) ComponentLike {
	return &component_{
		// Initialize instance attributes.
		class_: c,
		primitive_: primitive,
	}
}

func (c *componentClass_) MakeWithLists(lists col.ListLike[ListLike]) ComponentLike {
	return &component_{
		// Initialize instance attributes.
		class_: c,
		lists_: lists,
	}
}

// INSTANCE METHODS

// Target

type component_ struct {
	// Define instance attributes.
	class_ ComponentClassLike
	default_ DefaultLike
	primitive_ PrimitiveLike
	lists_ col.ListLike[ListLike]
}

// Attributes

func (v *component_) GetClass() ComponentClassLike {
	return v.class_
}

func (v *component_) GetDefault() DefaultLike {
	return v.default_
}

func (v *component_) GetPrimitive() PrimitiveLike {
	return v.primitive_
}

func (v *component_) GetLists() col.ListLike[ListLike] {
	return v.lists_
}

// Private
