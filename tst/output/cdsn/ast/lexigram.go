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

var lexigramClass = &lexigramClass_{
	// Initialize class constants.
}

// Function

func Lexigram() LexigramClassLike {
	return lexigramClass
}

// CLASS METHODS

// Target

type lexigramClass_ struct {
	// Define class constants.
}

// Constructors

func (c *lexigramClass_) MakeWithAttributes(
	comments col.ListLike[string],
	lowercase string,
	pattern PatternLike,
	notes col.ListLike[string],
) LexigramLike {
	return &lexigram_{
		// Initialize instance attributes.
		class_: c,
		comments_: comments,
		lowercase_: lowercase,
		pattern_: pattern,
		notes_: notes,
	}
}

// INSTANCE METHODS

// Target

type lexigram_ struct {
	// Define instance attributes.
	class_ LexigramClassLike
	comments_ col.ListLike[string]
	lowercase_ string
	pattern_ PatternLike
	notes_ col.ListLike[string]
}

// Attributes

func (v *lexigram_) GetClass() LexigramClassLike {
	return v.class_
}

func (v *lexigram_) GetComments() col.ListLike[string] {
	return v.comments_
}

func (v *lexigram_) GetLowercase() string {
	return v.lowercase_
}

func (v *lexigram_) GetPattern() PatternLike {
	return v.pattern_
}

func (v *lexigram_) GetNotes() col.ListLike[string] {
	return v.notes_
}

// Private
