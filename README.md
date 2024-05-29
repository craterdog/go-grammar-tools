<img src="https://craterdog.com/images/CraterDog.png" width="50%">

## Go Language Grammar Tools

### Overview
This project provides a set of Go based command-line tools that can be used to
parse, validate and format a language grammar defined in a `Syntax.cdsn` file
using Crater Dog Syntax Notation™ (aka CDSN) using the Go Language Grammar
Framework.  It can also generate the corresponding `Package.go` files for the
abstract syntax tree (AST) and agent packages that implement the language
grammar.

### Quick Links
For more information on this project click on the following links:
 * [project documentation](https://github.com/craterdog/go-grammar-tools/wiki)
 * [grammar framework](https://github.com/craterdog/go-grammar-framework/wiki)

### Getting Started
To install the language grammar tools do the following from a terminal window:
```
$ git clone git@github.com:craterdog/go-grammar-tools.git
$ cd go-grammar-tools/
$ etc/build.sh
$ ls
LICENSE		bin		go.mod		src
README.md	etc		go.sum		tst

$ ls bin/
format		generate	initialize	validate	version

$ bin/version
bin/version: v4.6.0
```

<H5 align="center"> Copyright © 2009 - 2024  Crater Dog Technologies™. All rights reserved. </H5>
