#!/bin/bash

# Source .env file
if [ -f .env ]; then
    source .env
fi

# Create folder if not exists
if [ ! -d data ]; then
    mkdir data
fi

# Set output file path
output=$(printf "data/%02d.txt" $1)

# Download data from adventofcode.com
url="${URL/DAY/$1}"
curl -o $output --cookie "session=$SESSION" $url