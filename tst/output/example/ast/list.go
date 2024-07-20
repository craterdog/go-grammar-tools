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
	abs "github.com/craterdog/go-collection-framework/v4/collection"
	col "github.com/craterdog/go-collection-framework/v4"
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

func (c *listClass_) Make(
	component ComponentLike,
	additionals abs.Sequential[AdditionalLike],
) ListLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(component):
		panic("The component attribute is required by this class.")
	case col.IsUndefined(additionals):
		panic("The additionals attribute is required by this class.")
	default:
		return &list_{
			// Initialize instance attributes.
			class_: c,
			component_: component,
			additionals_: additionals,
		}
	}
}

// INSTANCE METHODS

// Target

type list_ struct {
	// Define instance attributes.
	class_ ListClassLike
	component_ ComponentLike
	additionals_ abs.Sequential[AdditionalLike]
}

// Attributes

func (v *list_) GetClass() ListClassLike {
	return v.class_
}

func (v *list_) GetComponent() ComponentLike {
	return v.component_
}

func (v *list_) GetAdditionals() abs.Sequential[AdditionalLike] {
	return v.additionals_
}

// Private
