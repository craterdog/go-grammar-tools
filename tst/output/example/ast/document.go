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

var documentClass = &documentClass_{
	// Initialize class constants.
}

// Function

func Document() DocumentClassLike {
	return documentClass
}

// CLASS METHODS

// Target

type documentClass_ struct {
	// Define class constants.
}

// Constructors

func (c *documentClass_) Make(component ComponentLike) DocumentLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(component):
		panic("The component attribute is required by this class.")
	default:
		return &document_{
			// Initialize instance attributes.
			class_: c,
			component_: component,
		}
	}
}

// INSTANCE METHODS

// Target

type document_ struct {
	// Define instance attributes.
	class_ DocumentClassLike
	component_ ComponentLike
}

// Attributes

func (v *document_) GetClass() DocumentClassLike {
	return v.class_
}

func (v *document_) GetComponent() ComponentLike {
	return v.component_
}

// Private
