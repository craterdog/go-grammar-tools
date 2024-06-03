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

func (c *listClass_) MakeWithExamples(examples col.ListLike[ExampleLike]) ListLike {
	return &list_{
		// Initialize instance attributes.
		class_: c,
		examples_: examples,
	}
}

// INSTANCE METHODS

// Target

type list_ struct {
	// Define instance attributes.
	class_ ListClassLike
	examples_ col.ListLike[ExampleLike]
}

// Attributes

func (v *list_) GetClass() ListClassLike {
	return v.class_
}

func (v *list_) GetExamples() col.ListLike[ExampleLike] {
	return v.examples_
}

// Private
