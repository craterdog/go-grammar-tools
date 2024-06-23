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

var methodClass = &methodClass_{
	// Initialize class constants.
}

// Function

func Method() MethodClassLike {
	return methodClass
}

// CLASS METHODS

// Target

type methodClass_ struct {
	// Define class constants.
}

// Constructors

func (c *methodClass_) MakeWithAttributes(
	identifier string,
	parameterses col.ListLike[ParametersLike],
	results col.ListLike[ResultLike],
) MethodLike {
	return &method_{
		// Initialize instance attributes.
		class_: c,
		identifier_: identifier,
		parameterses_: parameterses,
		results_: results,
	}
}

// INSTANCE METHODS

// Target

type method_ struct {
	// Define instance attributes.
	class_ MethodClassLike
	identifier_ string
	parameterses_ col.ListLike[ParametersLike]
	results_ col.ListLike[ResultLike]
}

// Attributes

func (v *method_) GetClass() MethodClassLike {
	return v.class_
}

func (v *method_) GetIdentifier() string {
	return v.identifier_
}

func (v *method_) GetParameterses() col.ListLike[ParametersLike] {
	return v.parameterses_
}

func (v *method_) GetResults() col.ListLike[ResultLike] {
	return v.results_
}

// Private
