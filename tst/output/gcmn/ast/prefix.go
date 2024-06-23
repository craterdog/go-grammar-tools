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

var prefixClass = &prefixClass_{
	// Initialize class constants.
}

// Function

func Prefix() PrefixClassLike {
	return prefixClass
}

// CLASS METHODS

// Target

type prefixClass_ struct {
	// Define class constants.
}

// Constructors

func (c *prefixClass_) MakeWithArray(array ArrayLike) PrefixLike {
	return &prefix_{
		// Initialize instance attributes.
		class_: c,
		array_: array,
	}
}

func (c *prefixClass_) MakeWithMap(map_ MapLike) PrefixLike {
	return &prefix_{
		// Initialize instance attributes.
		class_: c,
		map_: map_,
	}
}

func (c *prefixClass_) MakeWithChannel(channel ChannelLike) PrefixLike {
	return &prefix_{
		// Initialize instance attributes.
		class_: c,
		channel_: channel,
	}
}

func (c *prefixClass_) MakeWithAlias(alias AliasLike) PrefixLike {
	return &prefix_{
		// Initialize instance attributes.
		class_: c,
		alias_: alias,
	}
}

// INSTANCE METHODS

// Target

type prefix_ struct {
	// Define instance attributes.
	class_ PrefixClassLike
	any_ any
	array_ ArrayLike
	map_ MapLike
	channel_ ChannelLike
	alias_ AliasLike
}

// Attributes

func (v *prefix_) GetClass() PrefixClassLike {
	return v.class_
}

func (v *prefix_) GetAny() any {
	return v.any_
}

// Private
