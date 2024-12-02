#! /bin/bash

# Set up dir for a new day
# Usage: ./newday.sh <day>

set -e
if [ $# -ne 1 ]; then
    echo "Usage: $0 <day>"
    exit 1
fi

mkdir -p day$1

# create main.go
cat <<EOL > day$1/main.go
package main

import (
    "fmt"
    "log"

	"github.com/KevinJoiner/AdventOfCode2024/aoc"
)

func main() {
	rows, err := aoc.ReadLines("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	output, err := puzzle1(rows)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Puzzle 1 output:", output)
	output, err = puzzle2(rows)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Puzzle 2 output:", output)
}

func puzzle1(rows [][]byte) (any, error) {
	return nil, nil
}

func puzzle2(rows [][]byte) (any, error) {
	return nil, nil
}

EOL