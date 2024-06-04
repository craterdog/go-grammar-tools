input=tst/input
output=tst/output
module=github.com/craterdog/example

echo "Initializing the existing language syntaxes:"
rm -rf $output/*
for file in `ls $input`; do
	syntax=${file%?cdsn}
	echo "    $syntax"
	mkdir $output/$syntax
	cp $input/$file $output/$syntax/Syntax.cdsn
done
echo

echo "Creating the new language syntax:"
echo "    example"
mkdir $output/example
bin/initialize $output/example/ example ""
echo

echo "Validating the language syntaxes:"
for file in `ls $input`; do
	syntax=${file%?cdsn}
	echo "    $syntax"
	bin/validate $output/$syntax/Syntax.cdsn
done
echo "    example"
bin/validate $output/example/Syntax.cdsn
echo

echo "Formatting the language syntaxes:"
for file in `ls $input`; do
	syntax=${file%?cdsn}
	echo "    $syntax"
	bin/format $output/$syntax/Syntax.cdsn
done
echo "    example"
bin/format $output/example/Syntax.cdsn
echo

echo "Generating the AST and agent packages:"
for file in `ls $input`; do
	syntax=${file%?cdsn}
	echo "    $syntax"
	bin/generate $module/$syntax $output/$syntax
done
echo "    example"
bin/generate $module/example $output/example
echo

echo "Generating the go.mod file:"
cd $output
cat <<EOF > go.mod
module github.com/craterdog/example

go 1.22
EOF
go mod tidy
echo

echo "Running lint on the generated files:"
golangci-lint run
cd - >/dev/null
echo

echo "Done."
