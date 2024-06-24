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

var listClass = &listClass_{
	// Initialize class constants.
}

// Function

func List() ListClassLike {
	return listClass
}

// CLASS METHODS

// Target

type listClass_ struct {
	// Define class constants.
}

// Constructors

func (c *listClass_) MakeWithAttributes(
	component ComponentLike,
	additionals col.ListLike[AdditionalLike],
) ListLike {
	return &list_{
		// Initialize instance attributes.
		class_: c,
		component_: component,
		additionals_: additionals,
	}
}

// INSTANCE METHODS

// Target

type list_ struct {
	// Define instance attributes.
	class_ ListClassLike
	component_ ComponentLike
	additionals_ col.ListLike[AdditionalLike]
}

// Attributes

func (v *list_) GetClass() ListClassLike {
	return v.class_
}

func (v *list_) GetComponent() ComponentLike {
	return v.component_
}

func (v *list_) GetAdditionals() col.ListLike[AdditionalLike] {
	return v.additionals_
}

// Private
