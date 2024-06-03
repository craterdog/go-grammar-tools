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

var exampleClass = &exampleClass_{
	// Initialize class constants.
}

// Function

func Example() ExampleClassLike {
	return exampleClass
}

// CLASS METHODS

// Target

type exampleClass_ struct {
	// Define class constants.
}

// Constructors

func (c *exampleClass_) MakeWithDefault(default_ DefaultLike) ExampleLike {
	return &example_{
		// Initialize instance attributes.
		class_: c,
		default_: default_,
	}
}

func (c *exampleClass_) MakeWithPrimitive(primitive PrimitiveLike) ExampleLike {
	return &example_{
		// Initialize instance attributes.
		class_: c,
		primitive_: primitive,
	}
}

func (c *exampleClass_) MakeWithLists(lists col.ListLike[ListLike]) ExampleLike {
	return &example_{
		// Initialize instance attributes.
		class_: c,
		lists_: lists,
	}
}

// INSTANCE METHODS

// Target

type example_ struct {
	// Define instance attributes.
	class_ ExampleClassLike
	default_ DefaultLike
	primitive_ PrimitiveLike
	lists_ col.ListLike[ListLike]
}

// Attributes

func (v *example_) GetClass() ExampleClassLike {
	return v.class_
}

func (v *example_) GetDefault() DefaultLike {
	return v.default_
}

func (v *example_) GetPrimitive() PrimitiveLike {
	return v.primitive_
}

func (v *example_) GetLists() col.ListLike[ListLike] {
	return v.lists_
}

// Private
