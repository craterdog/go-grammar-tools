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

var modelClass = &modelClass_{
	// Initialize class constants.
}

// Function

func Model() ModelClassLike {
	return modelClass
}

// CLASS METHODS

// Target

type modelClass_ struct {
	// Define class constants.
}

// Constructors

func (c *modelClass_) MakeWithAttributes(
	notice NoticeLike,
	header HeaderLike,
	moduleses col.ListLike[ModulesLike],
	typeses col.ListLike[TypesLike],
	functionalses col.ListLike[FunctionalsLike],
	aspectses col.ListLike[AspectsLike],
	classeses col.ListLike[ClassesLike],
	instanceses col.ListLike[InstancesLike],
) ModelLike {
	return &model_{
		// Initialize instance attributes.
		class_: c,
		notice_: notice,
		header_: header,
		moduleses_: moduleses,
		typeses_: typeses,
		functionalses_: functionalses,
		aspectses_: aspectses,
		classeses_: classeses,
		instanceses_: instanceses,
	}
}

// INSTANCE METHODS

// Target

type model_ struct {
	// Define instance attributes.
	class_ ModelClassLike
	notice_ NoticeLike
	header_ HeaderLike
	moduleses_ col.ListLike[ModulesLike]
	typeses_ col.ListLike[TypesLike]
	functionalses_ col.ListLike[FunctionalsLike]
	aspectses_ col.ListLike[AspectsLike]
	classeses_ col.ListLike[ClassesLike]
	instanceses_ col.ListLike[InstancesLike]
}

// Attributes

func (v *model_) GetClass() ModelClassLike {
	return v.class_
}

func (v *model_) GetNotice() NoticeLike {
	return v.notice_
}

func (v *model_) GetHeader() HeaderLike {
	return v.header_
}

func (v *model_) GetModuleses() col.ListLike[ModulesLike] {
	return v.moduleses_
}

func (v *model_) GetTypeses() col.ListLike[TypesLike] {
	return v.typeses_
}

func (v *model_) GetFunctionalses() col.ListLike[FunctionalsLike] {
	return v.functionalses_
}

func (v *model_) GetAspectses() col.ListLike[AspectsLike] {
	return v.aspectses_
}

func (v *model_) GetClasseses() col.ListLike[ClassesLike] {
	return v.classeses_
}

func (v *model_) GetInstanceses() col.ListLike[InstancesLike] {
	return v.instanceses_
}

// Private
