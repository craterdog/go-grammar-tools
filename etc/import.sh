input=tst/input

echo "Removing the old example language syntaxes:"
rm -r $input/*
echo

echo "Importing the new example language syntaxes:"
echo "    cdsn.cdsn"
cp ../go-grammar-framework/v4/Syntax.cdsn $input/cdsn.cdsn
echo "    gcmn.cdsn"
cp ../go-model-framework/v4/Syntax.cdsn $input/gcmn.cdsn
echo

echo "Done."
