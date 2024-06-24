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

import ()

// CLASS ACCESS

// Reference

var namedClass = &namedClass_{
	// Initialize class constants.
}

// Function

func Named() NamedClassLike {
	return namedClass
}

// CLASS METHODS

// Target

type namedClass_ struct {
	// Define class constants.
}

// Constructors

func (c *namedClass_) Make(parameters ParametersLike) NamedLike {
	return &named_{
		// Initialize instance attributes.
		class_: c,
	}
}

// INSTANCE METHODS

// Target

type named_ struct {
	// Define instance attributes.
	class_ NamedClassLike
	parameters_ ParametersLike
}

// Attributes

func (v *named_) GetClass() NamedClassLike {
	return v.class_
}

func (v *named_) GetParameters() ParametersLike {
	return v.parameters_
}

// Private
