#!/bin/bash

# List files and directories ordered by access time, ignoring hidden files
files=$(ls -At --time=atime --ignore='.*' --indicator-style=slash)

# Replace newlines with commas
result=$(echo "$files" | paste -sd, -)

# Print the result
echo "$result"
