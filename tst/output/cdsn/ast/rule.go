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

var ruleClass = &ruleClass_{
	// Initialize class constants.
}

// Function

func Rule() RuleClassLike {
	return ruleClass
}

// CLASS METHODS

// Target

type ruleClass_ struct {
	// Define class constants.
}

// Constructors

func (c *ruleClass_) MakeWithAttributes(
	comments col.ListLike[string],
	uppercase string,
	expression ExpressionLike,
) RuleLike {
	return &rule_{
		// Initialize instance attributes.
		class_: c,
		comments_: comments,
		uppercase_: uppercase,
		expression_: expression,
	}
}

// INSTANCE METHODS

// Target

type rule_ struct {
	// Define instance attributes.
	class_ RuleClassLike
	comments_ col.ListLike[string]
	uppercase_ string
	expression_ ExpressionLike
}

// Attributes

func (v *rule_) GetClass() RuleClassLike {
	return v.class_
}

func (v *rule_) GetComments() col.ListLike[string] {
	return v.comments_
}

func (v *rule_) GetUppercase() string {
	return v.uppercase_
}

func (v *rule_) GetExpression() ExpressionLike {
	return v.expression_
}

// Private
