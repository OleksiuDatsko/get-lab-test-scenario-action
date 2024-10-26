#!/bin/bash

# Set the source directory
SRC_DIR="src"

# Loop through each directory in src
for dir in "$SRC_DIR"/*; do
  echo "Building $dir..."
  # Check if it's a directory
  if [ -d "$dir" ]; then
    # Get the base name of the directory (e.g., "lab1-parser")
    base_name=$(basename "$dir")
    echo "Building file $base_name..."

    
    # Run go build and set output to the base directory name in the root
    go build -C "$dir" -o "../../$base_name"
    
    if [ $? -eq 0 ]; then
      echo "Built $base_name successfully."
    else
      echo "Failed to build $base_name."
    fi
  fi
done
