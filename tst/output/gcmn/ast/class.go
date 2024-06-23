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

var classClass = &classClass_{
	// Initialize class constants.
}

// Function

func Class() ClassClassLike {
	return classClass
}

// CLASS METHODS

// Target

type classClass_ struct {
	// Define class constants.
}

// Constructors

func (c *classClass_) MakeWithAttributes(
	declaration DeclarationLike,
	constantses col.ListLike[ConstantsLike],
	constructorses col.ListLike[ConstructorsLike],
	functionses col.ListLike[FunctionsLike],
) ClassLike {
	return &class_{
		// Initialize instance attributes.
		class_: c,
		declaration_: declaration,
		constantses_: constantses,
		constructorses_: constructorses,
		functionses_: functionses,
	}
}

// INSTANCE METHODS

// Target

type class_ struct {
	// Define instance attributes.
	class_ ClassClassLike
	declaration_ DeclarationLike
	constantses_ col.ListLike[ConstantsLike]
	constructorses_ col.ListLike[ConstructorsLike]
	functionses_ col.ListLike[FunctionsLike]
}

// Attributes

func (v *class_) GetClass() ClassClassLike {
	return v.class_
}

func (v *class_) GetDeclaration() DeclarationLike {
	return v.declaration_
}

func (v *class_) GetConstantses() col.ListLike[ConstantsLike] {
	return v.constantses_
}

func (v *class_) GetConstructorses() col.ListLike[ConstructorsLike] {
	return v.constructorses_
}

func (v *class_) GetFunctionses() col.ListLike[FunctionsLike] {
	return v.functionses_
}

// Private
