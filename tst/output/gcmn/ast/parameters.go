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

var parametersClass = &parametersClass_{
	// Initialize class constants.
}

// Function

func Parameters() ParametersClassLike {
	return parametersClass
}

// CLASS METHODS

// Target

type parametersClass_ struct {
	// Define class constants.
}

// Constructors

func (c *parametersClass_) Make(
	parameter ParameterLike,
	additionalParameters col.ListLike[AdditionalParameterLike],
) ParametersLike {
	return &parameters_{
		// Initialize instance attributes.
		class_: c,
	}
}

// INSTANCE METHODS

// Target

type parameters_ struct {
	// Define instance attributes.
	class_ ParametersClassLike
	parameter_ ParameterLike
	additionalParameters_ col.ListLike[AdditionalParameterLike]
}

// Attributes

func (v *parameters_) GetClass() ParametersClassLike {
	return v.class_
}

func (v *parameters_) GetParameter() ParameterLike {
	return v.parameter_
}

func (v *parameters_) GetAdditionalParameters() col.ListLike[AdditionalParameterLike] {
	return v.additionalParameters_
}

// Private
