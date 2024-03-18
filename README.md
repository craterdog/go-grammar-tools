<img src="https://craterdog.com/images/CraterDog.png" width="50%">

## Go Language Grammar Tools

### Overview
This project provides a set of Go based tools that can validate and format a
language grammar defined using Crater Dog Syntax Notation™ (aka CDSN).  For full
details on the Go Language Grammar Framework click
[here](https://github.com/craterdog/go-grammar-framework/wiki)

### Getting Started
To install the language grammar tools do the following from a terminal window:
```
$ git clone git@github.com:craterdog/go-grammar-tools.git
$ cd go-grammar-tools/
$ ./etc/build.sh
$ ls
LICENSE		bin		go.mod		src
README.md	etc		go.sum

$ls bin/
format		validate
```

### Using the Tools
The `validate` command reads in a language grammar file and ensures that it is
formatted using Crater Dog Syntax Notation™ (aka CDSN) as follows:
```
$ go-grammar-tools/bin/validate example/Grammar.cdsn
```

The `format` command reads in a language grammar file formatted using Crater Dog
Syntax Notation™ (aka CDSN) and reformats it in its canonical format as
follows:
```
$ go-grammar-tools/bin/format example/Grammar.cdsn
```

<H5 align="center"> Copyright © 2009 - 2024  Crater Dog Technologies™. All rights reserved. </H5>
