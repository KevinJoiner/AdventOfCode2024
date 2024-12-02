package aoc

import (
	"bytes"
	"fmt"
	"os"
)

func ReadLines(filename string) ([][]byte, error) {
	input, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("could not read file: %w", err)
	}
	lines := bytes.Split(input, []byte("\n"))
	return lines[:len(lines)-1], nil
}
