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

var factorClass = &factorClass_{
	// Initialize class constants.
}

// Function

func Factor() FactorClassLike {
	return factorClass
}

// CLASS METHODS

// Target

type factorClass_ struct {
	// Define class constants.
}

// Constructors

func (c *factorClass_) MakeWithAttributes(
	predicate PredicateLike,
	cardinalitys col.ListLike[CardinalityLike],
) FactorLike {
	return &factor_{
		// Initialize instance attributes.
		class_: c,
		predicate_: predicate,
		cardinalitys_: cardinalitys,
	}
}

// INSTANCE METHODS

// Target

type factor_ struct {
	// Define instance attributes.
	class_ FactorClassLike
	predicate_ PredicateLike
	cardinalitys_ col.ListLike[CardinalityLike]
}

// Attributes

func (v *factor_) GetClass() FactorClassLike {
	return v.class_
}

func (v *factor_) GetPredicate() PredicateLike {
	return v.predicate_
}

func (v *factor_) GetCardinalitys() col.ListLike[CardinalityLike] {
	return v.cardinalitys_
}

// Private
