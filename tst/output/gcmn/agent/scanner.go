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
package agent

import (
	fmt "fmt"
	cdc "github.com/craterdog/go-collection-framework/v4/cdcn"
	col "github.com/craterdog/go-collection-framework/v4/collection"
	reg "regexp"
	sts "strings"
)

// CLASS ACCESS

// Reference

var scannerClass = &scannerClass_{
	// Initialize the class constants.
	tokens_: map[TokenType]string{
		// TBA - Add additional token types.
		ErrorToken:     "error",
		DelimiterToken: "delimiter",
		EOFToken:       "EOF",
		EOLToken:       "EOL",
		SpaceToken:     "space",
	},
	matchers_: map[TokenType]*reg.Regexp{
		// TBA - Add additional token types.
		DelimiterToken: reg.MustCompile("^(?:" + delimiter_ + ")"),
		EOLToken:       reg.MustCompile("^(?:" + eol_ + ")"),
		SpaceToken:     reg.MustCompile("^(?:" + space_ + ")"),
	},
}

// Function

func Scanner() ScannerClassLike {
	return scannerClass
}

// CLASS METHODS

// Target

type scannerClass_ struct {
	// Define the class constants.
	tokens_   map[TokenType]string
	matchers_ map[TokenType]*reg.Regexp
}

// Constructors

func (c *scannerClass_) Make(
	source string,
	tokens col.QueueLike[TokenLike],
) ScannerLike {
	var scanner = &scanner_{
		// Initialize the instance attributes.
		class_:    c,
		line_:     1,
		position_: 1,
		runes_:    []rune(source),
		tokens_:   tokens,
	}
	go scanner.scanTokens() // Start scanning tokens in the background.
	return scanner
}

// Functions

func (c *scannerClass_) AsString(type_ TokenType) string {
	return c.tokens_[type_]
}

func (c *scannerClass_) FormatToken(token TokenLike) string {
	var value = token.GetValue()
	var s = fmt.Sprintf("%q", value)
	if len(s) > 40 {
		s = fmt.Sprintf("%.40q...", value)
	}
	return fmt.Sprintf(
		"Token [type: %s, line: %d, position: %d]: %s",
		c.tokens_[token.GetType()],
		token.GetLine(),
		token.GetPosition(),
		s,
	)
}

func (c *scannerClass_) MatchToken(
	type_ TokenType,
	text string,
) col.ListLike[string] {
	var matcher = c.matchers_[type_]
	var matches = matcher.FindStringSubmatch(text)
	var notation = cdc.Notation().Make()
	return col.List[string](notation).MakeFromArray(matches)
}

// INSTANCE METHODS

// Target

type scanner_ struct {
	// Define the instance attributes.
	class_    ScannerClassLike
	first_    int // A zero based index of the first possible rune in the next token.
	next_     int // A zero based index of the next possible rune in the next token.
	line_     int // The line number in the source string of the next rune.
	position_ int // The position in the current line of the next rune.
	runes_    []rune
	tokens_   col.QueueLike[TokenLike]
}

// Attributes

func (v *scanner_) GetClass() ScannerClassLike {
	return v.class_
}

// Private

func (v *scanner_) emitToken(type_ TokenType) {
	var value = string(v.runes_[v.first_:v.next_])
	switch value {
	case "\x00":
		value = "<NULL>"
	case "\a":
		value = "<BELL>"
	case "\b":
		value = "<BKSP>"
	case "\t":
		value = "<HTAB>"
	case "\f":
		value = "<FMFD>"
	case "\n":
		value = "<EOLN>"
	case "\r":
		value = "<CRTN>"
	case "\v":
		value = "<VTAB>"
	}
	var token = Token().Make(v.line_, v.position_, type_, value)
	//fmt.Println(Scanner().FormatToken(token)) // Uncomment when debugging.
	v.tokens_.AddValue(token) // This will block if the queue is full.
}

func (v *scanner_) foundEOF() {
	v.emitToken(EOFToken)
}

func (v *scanner_) foundError() {
	v.next_++
	v.emitToken(ErrorToken)
}

func (v *scanner_) foundToken(type_ TokenType) bool {
	var text = string(v.runes_[v.next_:])
	var matches = Scanner().MatchToken(type_, text)
	if !matches.IsEmpty() {
		var match = matches.GetValue(1)
		var token = []rune(match)
		var length = len(token)

		// Found the requested token type.
		v.next_ += length
		if type_ != SpaceToken {
			v.emitToken(type_)
		}
		var count = sts.Count(match, "\n")
		if count > 0 {
			v.line_ += count
			v.position_ = v.indexOfLastEOL(token)
		} else {
			v.position_ += v.next_ - v.first_
		}
		v.first_ = v.next_
		return true
	}

	// The next token is not the requested token type.
	return false
}

func (v *scanner_) indexOfLastEOL(runes []rune) int {
	var length = len(runes)
	for index := length; index > 0; index-- {
		if runes[index-1] == '\n' {
			return length - index + 1
		}
	}
	return 0
}

func (v *scanner_) scanTokens() {
loop:
	for v.next_ < len(v.runes_) {
		switch {
		// TBA - Add additional token types.
		case v.foundToken(DelimiterToken):
		case v.foundToken(EOLToken):
		case v.foundToken(SpaceToken):
		default:
			v.foundError()
			break loop
		}
	}
	v.foundEOF()
}

/*
NOTE:
These private constants define the regular expression sub-patterns that make up
all token types.  Unfortunately there is no way to make them private to the
scanner class since they must be TRUE Go constants to be initialized in this
way.  We append an underscore to each name to lessen the chance of a name
collision with other private Go class constants in this package.
*/
const (
	// TBA - Add additional token types.
	any_       = `.|` + eol_
	base16_    = `[0-9a-f]`
	control_   = `\p{Cc}`
	delimiter_ = `[:;,\.=]` // TBA - Replace with the actual delimeters.
	digit_     = `\p{Nd}`
	eof_       = `\z`
	eol_       = `\n`
	escape_    = `\\(?:(?:` + unicode_ + `)|[abfnrtv'"\\])`
	letter_    = lower_ + `|` + upper_
	lower_     = `\p{Ll}`
	number_    = `(?:` + digit_ + `)+`
	rune_      = `['][^` + control_ + `][']`
	space_     = `[ \t]+`
	string_    = `["](?:` + escape_ + `|[^"` + control_ + `])+?["]`
	unicode_   = `x` + base16_ + `{2}|u` + base16_ + `{4}|U` + base16_ + `{8}`
	upper_     = `\p{Lu}`
)
