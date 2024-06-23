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

var lineClass = &lineClass_{
	// Initialize class constants.
}

// Function

func Line() LineClassLike {
	return lineClass
}

// CLASS METHODS

// Target

type lineClass_ struct {
	// Define class constants.
}

// Constructors

func (c *lineClass_) MakeWithAttributes(
	identifier IdentifierLike,
	notes col.ListLike[string],
) LineLike {
	return &line_{
		// Initialize instance attributes.
		class_: c,
		identifier_: identifier,
		notes_: notes,
	}
}

// INSTANCE METHODS

// Target

type line_ struct {
	// Define instance attributes.
	class_ LineClassLike
	identifier_ IdentifierLike
	notes_ col.ListLike[string]
}

// Attributes

func (v *line_) GetClass() LineClassLike {
	return v.class_
}

func (v *line_) GetIdentifier() IdentifierLike {
	return v.identifier_
}

func (v *line_) GetNotes() col.ListLike[string] {
	return v.notes_
}

// Private
