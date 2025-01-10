#!/bin/bash

# Create a new day for the Advent of Code
# Usage: ./create_day.sh <year> <day>

# Check arguments
if [ "$#" -ne 2 ]; then
    echo "Usage: $0 <year> <day> 2"
    exit 1
fi

TARGET=$1/$2

# Check if the TARGET directory exists, if not, create it
if [ ! -d "$TARGET" ]; then
    mkdir -p "$TARGET"
fi

# Copy files from the template directory to TARGET
cp -r template/* "$TARGET"

# Replace the placeholder in the files
find "$TARGET/main.go.template" -type f -exec sed -i "s/##DAY##/$(basename "$TARGET")/g" {} \;

# Rename main.go.template to main.go
mv "$TARGET/main.go.template" "$TARGET/main.go"

echo "Day created at $TARGET"