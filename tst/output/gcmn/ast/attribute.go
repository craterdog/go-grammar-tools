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

var attributeClass = &attributeClass_{
	// Initialize class constants.
}

// Function

func Attribute() AttributeClassLike {
	return attributeClass
}

// CLASS METHODS

// Target

type attributeClass_ struct {
	// Define class constants.
}

// Constructors

func (c *attributeClass_) MakeWithAttributes(
	identifier string,
	parameters col.ListLike[ParameterLike],
	abstractions col.ListLike[AbstractionLike],
) AttributeLike {
	return &attribute_{
		// Initialize instance attributes.
		class_: c,
		identifier_: identifier,
		parameters_: parameters,
		abstractions_: abstractions,
	}
}

// INSTANCE METHODS

// Target

type attribute_ struct {
	// Define instance attributes.
	class_ AttributeClassLike
	identifier_ string
	parameters_ col.ListLike[ParameterLike]
	abstractions_ col.ListLike[AbstractionLike]
}

// Attributes

func (v *attribute_) GetClass() AttributeClassLike {
	return v.class_
}

func (v *attribute_) GetIdentifier() string {
	return v.identifier_
}

func (v *attribute_) GetParameters() col.ListLike[ParameterLike] {
	return v.parameters_
}

func (v *attribute_) GetAbstractions() col.ListLike[AbstractionLike] {
	return v.abstractions_
}

// Private
