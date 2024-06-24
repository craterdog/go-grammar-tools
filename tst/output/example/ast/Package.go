/*
................................................................................
.                   Copyright (c) 2024.  All Rights Reserved.                  .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

/*
Package "ast" provides...

For detailed documentation on this package refer to the wiki:
  - <wiki>

This package follows the Crater Dog Technologies™ Go Coding Conventions located
here:
  - https://github.com/craterdog/go-model-framework/wiki

Additional concrete implementations of the classes defined by this package can
be developed and used seamlessly since the interface definitions only depend on
other interfaces and primitive types—and the class implementations only depend
on interfaces, not on each other.
*/
package ast

import (
	col "github.com/craterdog/go-collection-framework/v4/collection"
)

// Classes

/*
AdditionalClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete additional-like class.
*/
type AdditionalClassLike interface {
	// Constructors
	MakeWithComponent(component ComponentLike) AdditionalLike
}

/*
ComponentClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete component-like class.
*/
type ComponentClassLike interface {
	// Constructors
	MakeWithPrimitive(primitive PrimitiveLike) ComponentLike
	MakeWithList(list ListLike) ComponentLike
}

/*
DocumentClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete document-like class.
*/
type DocumentClassLike interface {
	// Constructors
	MakeWithComponent(component ComponentLike) DocumentLike
}

/*
ListClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete list-like class.
*/
type ListClassLike interface {
	// Constructors
	MakeWithAttributes(
		component ComponentLike,
		additionals col.ListLike[AdditionalLike],
	) ListLike
}

/*
PrimitiveClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete primitive-like class.
*/
type PrimitiveClassLike interface {
	// Constructors
	MakeWithInteger(integer string) PrimitiveLike
	MakeWithRune(rune_ string) PrimitiveLike
	MakeWithText(text string) PrimitiveLike
}

// Instances

/*
AdditionalLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete additional-like class.
*/
type AdditionalLike interface {
	// Attributes
	GetClass() AdditionalClassLike
	GetComponent() ComponentLike
}

/*
ComponentLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete component-like class.
*/
type ComponentLike interface {
	// Attributes
	GetClass() ComponentClassLike
	GetAny() any
}

/*
DocumentLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete document-like class.
*/
type DocumentLike interface {
	// Attributes
	GetClass() DocumentClassLike
	GetComponent() ComponentLike
}

/*
ListLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete list-like class.
*/
type ListLike interface {
	// Attributes
	GetClass() ListClassLike
	GetComponent() ComponentLike
	GetAdditionals() col.ListLike[AdditionalLike]
}

/*
PrimitiveLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete primitive-like class.
*/
type PrimitiveLike interface {
	// Attributes
	GetClass() PrimitiveClassLike
	GetAny() any
}
