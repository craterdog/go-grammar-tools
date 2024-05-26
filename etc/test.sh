input=tst/input
output=tst/output
module=github.com/craterdog/example

echo "Validating the syntax files..."
for file in `ls $input`; do
	echo "    $file"
	bin/validate $input/$file
done
echo

echo "Formatting the syntax files..."
for file in `ls $input`; do
	echo "    $file"
	model=${file%?cdsn}
	rm -rf $output/$model
	mkdir $output/$model
	cp $input/$file $output/$model/Syntax.cdsn
	bin/format $output/$model/Syntax.cdsn
done
echo

echo "Generating the AST and agent packages..."
for file in `ls $input`; do
	model=${file%?cdsn}
	echo "    $model"
	bin/generate $module/$model $output/$model
done
echo

echo "Initializing a syntax file..."
rm -rf $output/example
mkdir $output/example
bin/initialize $output/example/ example ""
echo

echo "Tidying up..."
cd $output
go mod tidy
cd - >/dev/null
echo "Done."
