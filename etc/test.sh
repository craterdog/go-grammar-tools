input=tst/input
output=tst/output

echo "Validating the syntax files:"
for file in `ls $input`; do
	echo "    $file"
	bin/validate $input/$file
done
echo

echo "Formatting the syntax files:"
rm -r $output/*
for file in `ls $input`; do
	echo "    $file"
	model=${file%?cdsn}
	mkdir $output/$model
	cp $input/$file $output/$model/Syntax.cdsn
	bin/format $output/$model/Syntax.cdsn
done
echo

echo "Generating the AST and agent packages:"
for syntax in `ls $output`; do
	echo "    $syntax"
	bin/generate $output/$syntax
done
echo

echo "Initializing a syntax file:"
bin/initialize $output/example/ example ""
echo
