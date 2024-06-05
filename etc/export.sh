output=tst/output
wiki=../../wiki/go-grammar-tools.wiki

echo "Removing the old examples:"
rm -r $wiki/Example-*
echo

echo "Exporting the syntax example:"
echo "    example/Syntax.cdsn"
echo '```' >$wiki/Example-Language-Syntax.md
cat $output/example/Syntax.cdsn >>$wiki/Example-Language-Syntax.md
echo '```' >>$wiki/Example-Language-Syntax.md
echo

echo "Exporting the AST examples:"
echo "    example/ast/Package.go"
echo '```go' >$wiki/Example-AST-Package.md
cat $output/example/ast/Package.go >>$wiki/Example-AST-Package.md
echo '```' >>$wiki/Example-AST-Package.md
echo "    example/ast/component.go"
echo '```go' >$wiki/Example-Component-Class.md
cat $output/example/ast/component.go >>$wiki/Example-Component-Class.md
echo '```' >>$wiki/Example-Component-Class.md
echo "    example/ast/default.go"
echo '```go' >$wiki/Example-Default-Class.md
cat $output/example/ast/default.go >>$wiki/Example-Default-Class.md
echo '```' >>$wiki/Example-Default-Class.md
echo "    example/ast/list.go"
echo '```go' >$wiki/Example-List-Class.md
cat $output/example/ast/list.go >>$wiki/Example-List-Class.md
echo '```' >>$wiki/Example-List-Class.md
echo "    example/ast/primitive.go"
echo '```go' >$wiki/Example-Primitive-Class.md
cat $output/example/ast/primitive.go >>$wiki/Example-Primitive-Class.md
echo '```' >>$wiki/Example-Primitive-Class.md

echo "Exporting the agent examples:"
echo "    example/agent/Package.go"
echo '```go' >$wiki/Example-Agent-Package.md
cat $output/example/agent/Package.go >>$wiki/Example-Agent-Package.md
echo '```' >>$wiki/Example-Agent-Package.md
echo "    example/agent/formatter.go"
echo '```go' >$wiki/Example-Formatter-Class.md
cat $output/example/agent/formatter.go >>$wiki/Example-Formatter-Class.md
echo '```' >>$wiki/Example-Formatter-Class.md
echo "    example/agent/parser.go"
echo '```go' >$wiki/Example-Parser-Class.md
cat $output/example/agent/parser.go >>$wiki/Example-Parser-Class.md
echo '```' >>$wiki/Example-Parser-Class.md
echo "    example/agent/scanner.go"
echo '```go' >$wiki/Example-Scanner-Class.md
cat $output/example/agent/scanner.go >>$wiki/Example-Scanner-Class.md
echo '```' >>$wiki/Example-Scanner-Class.md
echo "    example/agent/token.go"
echo '```go' >$wiki/Example-Token-Class.md
cat $output/example/agent/token.go >>$wiki/Example-Token-Class.md
echo '```' >>$wiki/Example-Token-Class.md
echo "    example/agent/validator.go"
echo '```go' >$wiki/Example-Validator-Class.md
cat $output/example/agent/validator.go >>$wiki/Example-Validator-Class.md
echo '```' >>$wiki/Example-Validator-Class.md
echo

echo "Done."
