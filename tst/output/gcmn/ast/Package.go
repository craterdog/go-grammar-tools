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
ModelClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete model-like class.
*/
type ModelClassLike interface {
	// Constructors
	MakeWithAttributes(
		notice NoticeLike,
		header HeaderLike,
		modules ModulesLike,
		types TypesLike,
		functionals FunctionalsLike,
		aspects AspectsLike,
		classes ClassesLike,
		instances InstancesLike,
	) ModelLike
}

/*
NoticeClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete notice-like class.
*/
type NoticeClassLike interface {
	// Constructors
	MakeWithComment(comment string) NoticeLike
}

/*
HeaderClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete header-like class.
*/
type HeaderClassLike interface {
	// Constructors
	MakeWithAttributes(
		comment string,
		identifier string,
	) HeaderLike
}

/*
ModulesClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete modules-like class.
*/
type ModulesClassLike interface {
	// Constructors
	MakeWithModules(modules col.ListLike[ModuleLike]) ModulesLike
}

/*
ModuleClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete module-like class.
*/
type ModuleClassLike interface {
	// Constructors
	MakeWithAttributes(
		identifier string,
		text string,
	) ModuleLike
}

/*
TypesClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete types-like class.
*/
type TypesClassLike interface {
	// Constructors
	MakeWithTypes(types col.ListLike[TypeLike]) TypesLike
}

/*
TypeClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete type-like class.
*/
type TypeClassLike interface {
	// Constructors
	MakeWithAttributes(
		declaration DeclarationLike,
		abstraction AbstractionLike,
		enumeration EnumerationLike,
	) TypeLike
}

/*
DeclarationClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete declaration-like class.
*/
type DeclarationClassLike interface {
	// Constructors
	MakeWithAttributes(
		comment string,
		identifier string,
		parameters ParametersLike,
	) DeclarationLike
}

/*
ParametersClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete parameters-like class.
*/
type ParametersClassLike interface {
	// Constructors
	MakeWithParameters(parameters col.ListLike[ParameterLike]) ParametersLike
}

/*
ParameterClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete parameter-like class.
*/
type ParameterClassLike interface {
	// Constructors
	MakeWithAttributes(
		identifier string,
		abstraction AbstractionLike,
	) ParameterLike
}

/*
AbstractionClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete abstraction-like class.
*/
type AbstractionClassLike interface {
	// Constructors
	MakeWithAttributes(
		prefix PrefixLike,
		identifier string,
		arguments ArgumentsLike,
	) AbstractionLike
}

/*
PrefixClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete prefix-like class.
*/
type PrefixClassLike interface {
	// Constructors
	MakeWithIdentifier(identifier string) PrefixLike
}

/*
ArgumentsClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete arguments-like class.
*/
type ArgumentsClassLike interface {
	// Constructors
	MakeWithAbstractions(abstractions col.ListLike[AbstractionLike]) ArgumentsLike
}

/*
EnumerationClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete enumeration-like class.
*/
type EnumerationClassLike interface {
	// Constructors
	MakeWithAttributes(
		parameter ParameterLike,
		identifiers col.ListLike[string],
	) EnumerationLike
}

/*
FunctionalsClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete functionals-like class.
*/
type FunctionalsClassLike interface {
	// Constructors
	MakeWithFunctionals(functionals col.ListLike[FunctionalLike]) FunctionalsLike
}

/*
FunctionalClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete functional-like class.
*/
type FunctionalClassLike interface {
	// Constructors
	MakeWithAttributes(
		declaration DeclarationLike,
		parameters ParametersLike,
		result ResultLike,
	) FunctionalLike
}

/*
ResultClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete result-like class.
*/
type ResultClassLike interface {
	// Constructors
	MakeWithAbstraction(abstraction AbstractionLike) ResultLike
	MakeWithParameters(parameters ParametersLike) ResultLike
}

/*
AspectsClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete aspects-like class.
*/
type AspectsClassLike interface {
	// Constructors
	MakeWithAspects(aspects col.ListLike[AspectLike]) AspectsLike
}

/*
AspectClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete aspect-like class.
*/
type AspectClassLike interface {
	// Constructors
	MakeWithAttributes(
		declaration DeclarationLike,
		methods MethodsLike,
	) AspectLike
}

/*
ClassesClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete classes-like class.
*/
type ClassesClassLike interface {
	// Constructors
	MakeWithClasses(classes col.ListLike[ClassLike]) ClassesLike
}

/*
ClassClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete class-like class.
*/
type ClassClassLike interface {
	// Constructors
	MakeWithAttributes(
		declaration DeclarationLike,
		constants ConstantsLike,
		constructors ConstructorsLike,
		functions FunctionsLike,
	) ClassLike
}

/*
ConstantsClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete constants-like class.
*/
type ConstantsClassLike interface {
	// Constructors
	MakeWithConstants(constants col.ListLike[ConstantLike]) ConstantsLike
}

/*
ConstantClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete constant-like class.
*/
type ConstantClassLike interface {
	// Constructors
	MakeWithAttributes(
		identifier string,
		abstraction AbstractionLike,
	) ConstantLike
}

/*
ConstructorsClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete constructors-like class.
*/
type ConstructorsClassLike interface {
	// Constructors
	MakeWithConstructors(constructors col.ListLike[ConstructorLike]) ConstructorsLike
}

/*
ConstructorClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete constructor-like class.
*/
type ConstructorClassLike interface {
	// Constructors
	MakeWithAttributes(
		identifier string,
		parameters ParametersLike,
		abstraction AbstractionLike,
	) ConstructorLike
}

/*
FunctionsClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete functions-like class.
*/
type FunctionsClassLike interface {
	// Constructors
	MakeWithFunctions(functions col.ListLike[FunctionLike]) FunctionsLike
}

/*
FunctionClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete function-like class.
*/
type FunctionClassLike interface {
	// Constructors
	MakeWithAttributes(
		identifier string,
		parameters ParametersLike,
		result ResultLike,
	) FunctionLike
}

/*
InstancesClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete instances-like class.
*/
type InstancesClassLike interface {
	// Constructors
	MakeWithInstances(instances col.ListLike[InstanceLike]) InstancesLike
}

/*
InstanceClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete instance-like class.
*/
type InstanceClassLike interface {
	// Constructors
	MakeWithAttributes(
		declaration DeclarationLike,
		attributes AttributesLike,
		abstractions AbstractionsLike,
		methods MethodsLike,
	) InstanceLike
}

/*
AttributesClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete attributes-like class.
*/
type AttributesClassLike interface {
	// Constructors
	MakeWithAttributes(attributes col.ListLike[AttributeLike]) AttributesLike
}

/*
AttributeClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete attribute-like class.
*/
type AttributeClassLike interface {
	// Constructors
	MakeWithAttributes(
		identifier string,
		parameter ParameterLike,
		abstraction AbstractionLike,
	) AttributeLike
}

/*
AbstractionsClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete abstractions-like class.
*/
type AbstractionsClassLike interface {
	// Constructors
	MakeWithAbstractions(abstractions col.ListLike[AbstractionLike]) AbstractionsLike
}

/*
MethodsClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete methods-like class.
*/
type MethodsClassLike interface {
	// Constructors
	MakeWithMethods(methods col.ListLike[MethodLike]) MethodsLike
}

/*
MethodClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete method-like class.
*/
type MethodClassLike interface {
	// Constructors
	MakeWithAttributes(
		identifier string,
		parameters ParametersLike,
		result ResultLike,
	) MethodLike
}

// Instances

/*
ModelLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete model-like class.
*/
type ModelLike interface {
	// Attributes
	GetClass() ModelClassLike
	GetNotice() NoticeLike
	GetHeader() HeaderLike
	GetModules() ModulesLike
	GetTypes() TypesLike
	GetFunctionals() FunctionalsLike
	GetAspects() AspectsLike
	GetClasses() ClassesLike
	GetInstances() InstancesLike
}

/*
NoticeLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete notice-like class.
*/
type NoticeLike interface {
	// Attributes
	GetClass() NoticeClassLike
	GetComment() string
}

/*
HeaderLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete header-like class.
*/
type HeaderLike interface {
	// Attributes
	GetClass() HeaderClassLike
	GetComment() string
	GetIdentifier() string
}

/*
ModulesLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete modules-like class.
*/
type ModulesLike interface {
	// Attributes
	GetClass() ModulesClassLike
	GetModules() col.ListLike[ModuleLike]
}

/*
ModuleLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete module-like class.
*/
type ModuleLike interface {
	// Attributes
	GetClass() ModuleClassLike
	GetIdentifier() string
	GetText() string
}

/*
TypesLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete types-like class.
*/
type TypesLike interface {
	// Attributes
	GetClass() TypesClassLike
	GetTypes() col.ListLike[TypeLike]
}

/*
TypeLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete type-like class.
*/
type TypeLike interface {
	// Attributes
	GetClass() TypeClassLike
	GetDeclaration() DeclarationLike
	GetAbstraction() AbstractionLike
	GetEnumeration() EnumerationLike
}

/*
DeclarationLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete declaration-like class.
*/
type DeclarationLike interface {
	// Attributes
	GetClass() DeclarationClassLike
	GetComment() string
	GetIdentifier() string
	GetParameters() ParametersLike
}

/*
ParametersLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete parameters-like class.
*/
type ParametersLike interface {
	// Attributes
	GetClass() ParametersClassLike
	GetParameters() col.ListLike[ParameterLike]
}

/*
ParameterLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete parameter-like class.
*/
type ParameterLike interface {
	// Attributes
	GetClass() ParameterClassLike
	GetIdentifier() string
	GetAbstraction() AbstractionLike
}

/*
AbstractionLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete abstraction-like class.
*/
type AbstractionLike interface {
	// Attributes
	GetClass() AbstractionClassLike
	GetPrefix() PrefixLike
	GetIdentifier() string
	GetArguments() ArgumentsLike
}

/*
PrefixLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete prefix-like class.
*/
type PrefixLike interface {
	// Attributes
	GetClass() PrefixClassLike
	GetIdentifier() string
}

/*
ArgumentsLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete arguments-like class.
*/
type ArgumentsLike interface {
	// Attributes
	GetClass() ArgumentsClassLike
	GetAbstractions() col.ListLike[AbstractionLike]
}

/*
EnumerationLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete enumeration-like class.
*/
type EnumerationLike interface {
	// Attributes
	GetClass() EnumerationClassLike
	GetParameter() ParameterLike
	GetIdentifiers() col.ListLike[string]
}

/*
FunctionalsLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete functionals-like class.
*/
type FunctionalsLike interface {
	// Attributes
	GetClass() FunctionalsClassLike
	GetFunctionals() col.ListLike[FunctionalLike]
}

/*
FunctionalLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete functional-like class.
*/
type FunctionalLike interface {
	// Attributes
	GetClass() FunctionalClassLike
	GetDeclaration() DeclarationLike
	GetParameters() ParametersLike
	GetResult() ResultLike
}

/*
ResultLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete result-like class.
*/
type ResultLike interface {
	// Attributes
	GetClass() ResultClassLike
	GetAbstraction() AbstractionLike
	GetParameters() ParametersLike
}

/*
AspectsLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete aspects-like class.
*/
type AspectsLike interface {
	// Attributes
	GetClass() AspectsClassLike
	GetAspects() col.ListLike[AspectLike]
}

/*
AspectLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete aspect-like class.
*/
type AspectLike interface {
	// Attributes
	GetClass() AspectClassLike
	GetDeclaration() DeclarationLike
	GetMethods() MethodsLike
}

/*
ClassesLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete classes-like class.
*/
type ClassesLike interface {
	// Attributes
	GetClass() ClassesClassLike
	GetClasses() col.ListLike[ClassLike]
}

/*
ClassLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete class-like class.
*/
type ClassLike interface {
	// Attributes
	GetClass() ClassClassLike
	GetDeclaration() DeclarationLike
	GetConstants() ConstantsLike
	GetConstructors() ConstructorsLike
	GetFunctions() FunctionsLike
}

/*
ConstantsLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete constants-like class.
*/
type ConstantsLike interface {
	// Attributes
	GetClass() ConstantsClassLike
	GetConstants() col.ListLike[ConstantLike]
}

/*
ConstantLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete constant-like class.
*/
type ConstantLike interface {
	// Attributes
	GetClass() ConstantClassLike
	GetIdentifier() string
	GetAbstraction() AbstractionLike
}

/*
ConstructorsLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete constructors-like class.
*/
type ConstructorsLike interface {
	// Attributes
	GetClass() ConstructorsClassLike
	GetConstructors() col.ListLike[ConstructorLike]
}

/*
ConstructorLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete constructor-like class.
*/
type ConstructorLike interface {
	// Attributes
	GetClass() ConstructorClassLike
	GetIdentifier() string
	GetParameters() ParametersLike
	GetAbstraction() AbstractionLike
}

/*
FunctionsLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete functions-like class.
*/
type FunctionsLike interface {
	// Attributes
	GetClass() FunctionsClassLike
	GetFunctions() col.ListLike[FunctionLike]
}

/*
FunctionLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete function-like class.
*/
type FunctionLike interface {
	// Attributes
	GetClass() FunctionClassLike
	GetIdentifier() string
	GetParameters() ParametersLike
	GetResult() ResultLike
}

/*
InstancesLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete instances-like class.
*/
type InstancesLike interface {
	// Attributes
	GetClass() InstancesClassLike
	GetInstances() col.ListLike[InstanceLike]
}

/*
InstanceLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete instance-like class.
*/
type InstanceLike interface {
	// Attributes
	GetClass() InstanceClassLike
	GetDeclaration() DeclarationLike
	GetAttributes() AttributesLike
	GetAbstractions() AbstractionsLike
	GetMethods() MethodsLike
}

/*
AttributesLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete attributes-like class.
*/
type AttributesLike interface {
	// Attributes
	GetClass() AttributesClassLike
	GetAttributes() col.ListLike[AttributeLike]
}

/*
AttributeLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete attribute-like class.
*/
type AttributeLike interface {
	// Attributes
	GetClass() AttributeClassLike
	GetIdentifier() string
	GetParameter() ParameterLike
	GetAbstraction() AbstractionLike
}

/*
AbstractionsLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete abstractions-like class.
*/
type AbstractionsLike interface {
	// Attributes
	GetClass() AbstractionsClassLike
	GetAbstractions() col.ListLike[AbstractionLike]
}

/*
MethodsLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete methods-like class.
*/
type MethodsLike interface {
	// Attributes
	GetClass() MethodsClassLike
	GetMethods() col.ListLike[MethodLike]
}

/*
MethodLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete method-like class.
*/
type MethodLike interface {
	// Attributes
	GetClass() MethodClassLike
	GetIdentifier() string
	GetParameters() ParametersLike
	GetResult() ResultLike
}
