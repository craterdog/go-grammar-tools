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
echo "    example/ast/document.go"
echo '```go' >$wiki/Example-Document-Class.md
cat $output/example/ast/document.go >>$wiki/Example-Document-Class.md
echo '```' >>$wiki/Example-Document-Class.md
echo "    example/ast/list.go"
echo '```go' >$wiki/Example-List-Class.md
cat $output/example/ast/list.go >>$wiki/Example-List-Class.md
echo '```' >>$wiki/Example-List-Class.md
echo "    example/ast/intrinsic.go"
echo '```go' >$wiki/Example-Intrinsic-Class.md
cat $output/example/ast/intrinsic.go >>$wiki/Example-Intrinsic-Class.md
echo '```' >>$wiki/Example-Intrinsic-Class.md

echo "Exporting the grammar examples:"
echo "    example/grammar/Package.go"
echo '```go' >$wiki/Example-Grammar-Package.md
cat $output/example/grammar/Package.go >>$wiki/Example-Grammar-Package.md
echo '```' >>$wiki/Example-Grammar-Package.md
echo "    example/grammar/formatter.go"
echo '```go' >$wiki/Example-Formatter-Class.md
cat $output/example/grammar/formatter.go >>$wiki/Example-Formatter-Class.md
echo '```' >>$wiki/Example-Formatter-Class.md
echo "    example/grammar/parser.go"
echo '```go' >$wiki/Example-Parser-Class.md
cat $output/example/grammar/parser.go >>$wiki/Example-Parser-Class.md
echo '```' >>$wiki/Example-Parser-Class.md
echo "    example/grammar/scanner.go"
echo '```go' >$wiki/Example-Scanner-Class.md
cat $output/example/grammar/scanner.go >>$wiki/Example-Scanner-Class.md
echo '```' >>$wiki/Example-Scanner-Class.md
echo "    example/grammar/token.go"
echo '```go' >$wiki/Example-Token-Class.md
cat $output/example/grammar/token.go >>$wiki/Example-Token-Class.md
echo '```' >>$wiki/Example-Token-Class.md
echo "    example/grammar/validator.go"
echo '```go' >$wiki/Example-Validator-Class.md
cat $output/example/grammar/validator.go >>$wiki/Example-Validator-Class.md
echo '```' >>$wiki/Example-Validator-Class.md
echo

echo "Done."
