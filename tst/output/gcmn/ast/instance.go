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

var instanceClass = &instanceClass_{
	// Initialize class constants.
}

// Function

func Instance() InstanceClassLike {
	return instanceClass
}

// CLASS METHODS

// Target

type instanceClass_ struct {
	// Define class constants.
}

// Constructors

func (c *instanceClass_) MakeWithAttributes(
	declaration DeclarationLike,
	attributeses col.ListLike[AttributesLike],
	abstractionses col.ListLike[AbstractionsLike],
	methodses col.ListLike[MethodsLike],
) InstanceLike {
	return &instance_{
		// Initialize instance attributes.
		class_: c,
		declaration_: declaration,
		attributeses_: attributeses,
		abstractionses_: abstractionses,
		methodses_: methodses,
	}
}

// INSTANCE METHODS

// Target

type instance_ struct {
	// Define instance attributes.
	class_ InstanceClassLike
	declaration_ DeclarationLike
	attributeses_ col.ListLike[AttributesLike]
	abstractionses_ col.ListLike[AbstractionsLike]
	methodses_ col.ListLike[MethodsLike]
}

// Attributes

func (v *instance_) GetClass() InstanceClassLike {
	return v.class_
}

func (v *instance_) GetDeclaration() DeclarationLike {
	return v.declaration_
}

func (v *instance_) GetAttributeses() col.ListLike[AttributesLike] {
	return v.attributeses_
}

func (v *instance_) GetAbstractionses() col.ListLike[AbstractionsLike] {
	return v.abstractionses_
}

func (v *instance_) GetMethodses() col.ListLike[MethodsLike] {
	return v.methodses_
}

// Private
