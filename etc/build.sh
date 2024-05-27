echo "Compiling the following tools:"
mkdir -p ./bin/
echo "	version"
go build -o ./bin/ ./src/version
echo "	initialize"
go build -o ./bin/ ./src/initialize
echo "	validate"
go build -o ./bin/ ./src/validate
echo "	format"
go build -o ./bin/ ./src/format
echo "	generate"
go build -o ./bin/ ./src/generate
echo "Done."
