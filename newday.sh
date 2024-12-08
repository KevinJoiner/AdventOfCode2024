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
    "time"

	"github.com/KevinJoiner/AdventOfCode2024/aoc"
)

func main() {
	rows, err := aoc.ReadLines("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	start := time.Now()
	output := puzzle1(rows)
	fmt.Println("Puzzle 1 output:", output)
	fmt.Println("Puzzle 1 Duration:", time.Since(start))
	start = time.Now()
	output = puzzle2(rows)
	fmt.Println("Puzzle 2 output:", output)
	fmt.Println("Puzzle 2 Duration:", time.Since(start))
}

func puzzle1(rows [][]byte) any {
	return nil
}

func puzzle2(rows [][]byte) any {
	return nil
}

EOL