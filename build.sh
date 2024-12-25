#!/bin/bash

# Print starting compiling message
echo "Compiling Golang program with go build"

# Source the main project file (enable this if you don't have a main package)
sourceFile="./main.go"

# Output binary file (if you don't set sourceFile, better to comment this line)
outputFile="./modol-app"

# Check if the bin directory exists
if [ ! -d "./bin" ]; then
    mkdir bin
fi

# Compile the Golang program
if [ -z "$sourceFile" ]; then
    go build -v -o ./bin/$outputFile
else
    go build -o ./bin/$outputFile $sourceFile
fi

# Check if the compilation was successful
if [ $? -eq 0 ]; then
    echo "Compilation was successful"
else
    echo "Compilation failed"
fi

# Run the compiled program
# Check if the outputFile is set is not empty run the compiled program
if [ -n "$outputFile" ]; then
    echo "Running the compiled program"
    ./bin/$outputFile
fi