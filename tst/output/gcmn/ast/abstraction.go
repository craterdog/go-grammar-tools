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

var abstractionClass = &abstractionClass_{
	// Initialize class constants.
}

// Function

func Abstraction() AbstractionClassLike {
	return abstractionClass
}

// CLASS METHODS

// Target

type abstractionClass_ struct {
	// Define class constants.
}

// Constructors

func (c *abstractionClass_) MakeWithAttributes(
	prefixs col.ListLike[PrefixLike],
	identifier string,
	genericArgumentses col.ListLike[GenericArgumentsLike],
) AbstractionLike {
	return &abstraction_{
		// Initialize instance attributes.
		class_: c,
		prefixs_: prefixs,
		identifier_: identifier,
		genericArgumentses_: genericArgumentses,
	}
}

// INSTANCE METHODS

// Target

type abstraction_ struct {
	// Define instance attributes.
	class_ AbstractionClassLike
	prefixs_ col.ListLike[PrefixLike]
	identifier_ string
	genericArgumentses_ col.ListLike[GenericArgumentsLike]
}

// Attributes

func (v *abstraction_) GetClass() AbstractionClassLike {
	return v.class_
}

func (v *abstraction_) GetPrefixs() col.ListLike[PrefixLike] {
	return v.prefixs_
}

func (v *abstraction_) GetIdentifier() string {
	return v.identifier_
}

func (v *abstraction_) GetGenericArgumentses() col.ListLike[GenericArgumentsLike] {
	return v.genericArgumentses_
}

// Private
