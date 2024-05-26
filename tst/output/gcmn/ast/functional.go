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

var functionalClass = &functionalClass_{
	// Any private class constants should be initialized here.
}

// Function

func Functional() FunctionalClassLike {
	return functionalClass
}

// CLASS METHODS

// Target

type functionalClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *functionalClass_) MakeWithAttributes(
	declaration DeclarationLike,
	parameters ParametersLike,
	result ResultLike,
) FunctionalLike {
	return &functional_{
		declaration_: declaration,
		parameters_: parameters,
		result_: result,
	}
}

// Functions

// INSTANCE METHODS

// Target

type functional_ struct {
	class_ FunctionalClassLike
	declaration_ DeclarationLike
	parameters_ ParametersLike
	result_ ResultLike
}

// Attributes

func (v *functional_) GetClass() FunctionalClassLike {
	return v.class_
}

func (v *functional_) GetDeclaration() DeclarationLike {
	return v.declaration_
}

func (v *functional_) GetParameters() ParametersLike {
	return v.parameters_
}

func (v *functional_) GetResult() ResultLike {
	return v.result_
}

// Public

// Private
