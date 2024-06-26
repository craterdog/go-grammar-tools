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
AbstractionClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete abstraction-like class.
*/
type AbstractionClassLike interface {
	// Constructors
	Make(
		prefix PrefixLike,
		identifier string,
		genericArguments GenericArgumentsLike,
	) AbstractionLike
}

/*
AbstractionsClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete abstractions-like class.
*/
type AbstractionsClassLike interface {
	// Constructors
	Make(abstractions col.ListLike[AbstractionLike]) AbstractionsLike
}

/*
AdditionalArgumentClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete additionalargument-like class.
*/
type AdditionalArgumentClassLike interface {
	// Constructors
	Make(argument ArgumentLike) AdditionalArgumentLike
}

/*
AdditionalParameterClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete additionalparameter-like class.
*/
type AdditionalParameterClassLike interface {
	// Constructors
	Make(parameter ParameterLike) AdditionalParameterLike
}

/*
AliasClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete alias-like class.
*/
type AliasClassLike interface {
	// Constructors
	Make(identifier string) AliasLike
}

/*
ArgumentClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete argument-like class.
*/
type ArgumentClassLike interface {
	// Constructors
	Make(abstraction AbstractionLike) ArgumentLike
}

/*
ArgumentsClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete arguments-like class.
*/
type ArgumentsClassLike interface {
	// Constructors
	Make(
		argument ArgumentLike,
		additionalArguments col.ListLike[AdditionalArgumentLike],
	) ArgumentsLike
}

/*
AspectClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete aspect-like class.
*/
type AspectClassLike interface {
	// Constructors
	Make(
		declaration DeclarationLike,
		methods MethodsLike,
	) AspectLike
}

/*
AspectsClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete aspects-like class.
*/
type AspectsClassLike interface {
	// Constructors
	Make(aspects col.ListLike[AspectLike]) AspectsLike
}

/*
AttributeClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete attribute-like class.
*/
type AttributeClassLike interface {
	// Constructors
	Make(
		identifier string,
		parameter ParameterLike,
		abstraction AbstractionLike,
	) AttributeLike
}

/*
AttributesClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete attributes-like class.
*/
type AttributesClassLike interface {
	// Constructors
	Make(attributes col.ListLike[AttributeLike]) AttributesLike
}

/*
ClassClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete class-like class.
*/
type ClassClassLike interface {
	// Constructors
	Make(
		declaration DeclarationLike,
		constants ConstantsLike,
		constructors ConstructorsLike,
		functions FunctionsLike,
	) ClassLike
}

/*
ClassesClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete classes-like class.
*/
type ClassesClassLike interface {
	// Constructors
	Make(classes col.ListLike[ClassLike]) ClassesLike
}

/*
ConstantClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete constant-like class.
*/
type ConstantClassLike interface {
	// Constructors
	Make(
		identifier string,
		abstraction AbstractionLike,
	) ConstantLike
}

/*
ConstantsClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete constants-like class.
*/
type ConstantsClassLike interface {
	// Constructors
	Make(constants col.ListLike[ConstantLike]) ConstantsLike
}

/*
ConstructorClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete constructor-like class.
*/
type ConstructorClassLike interface {
	// Constructors
	Make(
		identifier string,
		parameters ParametersLike,
		abstraction AbstractionLike,
	) ConstructorLike
}

/*
ConstructorsClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete constructors-like class.
*/
type ConstructorsClassLike interface {
	// Constructors
	Make(constructors col.ListLike[ConstructorLike]) ConstructorsLike
}

/*
DeclarationClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete declaration-like class.
*/
type DeclarationClassLike interface {
	// Constructors
	Make(
		comment string,
		identifier string,
		genericParameters GenericParametersLike,
	) DeclarationLike
}

/*
EnumerationClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete enumeration-like class.
*/
type EnumerationClassLike interface {
	// Constructors
	Make(
		parameter ParameterLike,
		identifiers col.ListLike[string],
	) EnumerationLike
}

/*
FunctionClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete function-like class.
*/
type FunctionClassLike interface {
	// Constructors
	Make(
		identifier string,
		parameters ParametersLike,
		result ResultLike,
	) FunctionLike
}

/*
FunctionalClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete functional-like class.
*/
type FunctionalClassLike interface {
	// Constructors
	Make(
		declaration DeclarationLike,
		parameters ParametersLike,
		result ResultLike,
	) FunctionalLike
}

/*
FunctionalsClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete functionals-like class.
*/
type FunctionalsClassLike interface {
	// Constructors
	Make(functionals col.ListLike[FunctionalLike]) FunctionalsLike
}

/*
FunctionsClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete functions-like class.
*/
type FunctionsClassLike interface {
	// Constructors
	Make(functions col.ListLike[FunctionLike]) FunctionsLike
}

/*
GenericArgumentsClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete genericarguments-like class.
*/
type GenericArgumentsClassLike interface {
	// Constructors
	Make(arguments ArgumentsLike) GenericArgumentsLike
}

/*
GenericParametersClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete genericparameters-like class.
*/
type GenericParametersClassLike interface {
	// Constructors
	Make(parameters ParametersLike) GenericParametersLike
}

/*
HeaderClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete header-like class.
*/
type HeaderClassLike interface {
	// Constructors
	Make(
		comment string,
		identifier string,
	) HeaderLike
}

/*
InstanceClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete instance-like class.
*/
type InstanceClassLike interface {
	// Constructors
	Make(
		declaration DeclarationLike,
		attributes AttributesLike,
		abstractions AbstractionsLike,
		methods MethodsLike,
	) InstanceLike
}

/*
InstancesClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete instances-like class.
*/
type InstancesClassLike interface {
	// Constructors
	Make(instances col.ListLike[InstanceLike]) InstancesLike
}

/*
MapClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete map-like class.
*/
type MapClassLike interface {
	// Constructors
	Make(identifier string) MapLike
}

/*
MethodClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete method-like class.
*/
type MethodClassLike interface {
	// Constructors
	Make(
		identifier string,
		parameters ParametersLike,
		result ResultLike,
	) MethodLike
}

/*
MethodsClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete methods-like class.
*/
type MethodsClassLike interface {
	// Constructors
	Make(methods col.ListLike[MethodLike]) MethodsLike
}

/*
ModelClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete model-like class.
*/
type ModelClassLike interface {
	// Constructors
	Make(
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
ModuleClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete module-like class.
*/
type ModuleClassLike interface {
	// Constructors
	Make(
		identifier string,
		text string,
	) ModuleLike
}

/*
ModulesClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete modules-like class.
*/
type ModulesClassLike interface {
	// Constructors
	Make(modules col.ListLike[ModuleLike]) ModulesLike
}

/*
NamedClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete named-like class.
*/
type NamedClassLike interface {
	// Constructors
	Make(parameters ParametersLike) NamedLike
}

/*
NoticeClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete notice-like class.
*/
type NoticeClassLike interface {
	// Constructors
	Make(comment string) NoticeLike
}

/*
ParameterClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete parameter-like class.
*/
type ParameterClassLike interface {
	// Constructors
	Make(
		identifier string,
		abstraction AbstractionLike,
	) ParameterLike
}

/*
ParametersClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete parameters-like class.
*/
type ParametersClassLike interface {
	// Constructors
	Make(
		parameter ParameterLike,
		additionalParameters col.ListLike[AdditionalParameterLike],
	) ParametersLike
}

/*
PrefixClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete prefix-like class.
*/
type PrefixClassLike interface {
	// Constructors
	Make(any any) PrefixLike
}

/*
ResultClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete result-like class.
*/
type ResultClassLike interface {
	// Constructors
	Make(any any) ResultLike
}

/*
TypeClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete type-like class.
*/
type TypeClassLike interface {
	// Constructors
	Make(
		declaration DeclarationLike,
		abstraction AbstractionLike,
		enumeration EnumerationLike,
	) TypeLike
}

/*
TypesClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete types-like class.
*/
type TypesClassLike interface {
	// Constructors
	Make(types col.ListLike[TypeLike]) TypesLike
}

// Instances

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
	GetGenericArguments() GenericArgumentsLike
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
AdditionalArgumentLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete additionalargument-like class.
*/
type AdditionalArgumentLike interface {
	// Attributes
	GetClass() AdditionalArgumentClassLike
	GetArgument() ArgumentLike
}

/*
AdditionalParameterLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete additionalparameter-like class.
*/
type AdditionalParameterLike interface {
	// Attributes
	GetClass() AdditionalParameterClassLike
	GetParameter() ParameterLike
}

/*
AliasLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete alias-like class.
*/
type AliasLike interface {
	// Attributes
	GetClass() AliasClassLike
	GetIdentifier() string
}

/*
ArgumentLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete argument-like class.
*/
type ArgumentLike interface {
	// Attributes
	GetClass() ArgumentClassLike
	GetAbstraction() AbstractionLike
}

/*
ArgumentsLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete arguments-like class.
*/
type ArgumentsLike interface {
	// Attributes
	GetClass() ArgumentsClassLike
	GetArgument() ArgumentLike
	GetAdditionalArguments() col.ListLike[AdditionalArgumentLike]
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
DeclarationLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete declaration-like class.
*/
type DeclarationLike interface {
	// Attributes
	GetClass() DeclarationClassLike
	GetComment() string
	GetIdentifier() string
	GetGenericParameters() GenericParametersLike
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
GenericArgumentsLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete genericarguments-like class.
*/
type GenericArgumentsLike interface {
	// Attributes
	GetClass() GenericArgumentsClassLike
	GetArguments() ArgumentsLike
}

/*
GenericParametersLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete genericparameters-like class.
*/
type GenericParametersLike interface {
	// Attributes
	GetClass() GenericParametersClassLike
	GetParameters() ParametersLike
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
MapLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete map-like class.
*/
type MapLike interface {
	// Attributes
	GetClass() MapClassLike
	GetIdentifier() string
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
NamedLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete named-like class.
*/
type NamedLike interface {
	// Attributes
	GetClass() NamedClassLike
	GetParameters() ParametersLike
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
ParametersLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete parameters-like class.
*/
type ParametersLike interface {
	// Attributes
	GetClass() ParametersClassLike
	GetParameter() ParameterLike
	GetAdditionalParameters() col.ListLike[AdditionalParameterLike]
}

/*
PrefixLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete prefix-like class.
*/
type PrefixLike interface {
	// Attributes
	GetClass() PrefixClassLike
	GetAny() any
}

/*
ResultLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete result-like class.
*/
type ResultLike interface {
	// Attributes
	GetClass() ResultClassLike
	GetAny() any
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
TypesLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete types-like class.
*/
type TypesLike interface {
	// Attributes
	GetClass() TypesClassLike
	GetTypes() col.ListLike[TypeLike]
}
