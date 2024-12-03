package main

import (
	"cmp"
	"fmt"
	"log"
	"regexp"
	"slices"
	"strconv"

	"github.com/KevinJoiner/AdventOfCode2024/aoc"
)

func main() {
	rows, err := aoc.ReadLines("./day3/input.txt")
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
	mulReg := regexp.MustCompile(`mul\((\w\w?\w?),(\w\w?\w?)\)`)
	total := 0
	for _, row := range rows {
		matches := mulReg.FindAllSubmatch(row, -1)
		for _, match := range matches {
			num1, err := strconv.Atoi(string(match[1]))
			if err != nil {
				panic(err)
			}
			num2, err := strconv.Atoi(string(match[2]))
			if err != nil {
				panic(err)
			}
			total += num1 * num2
		}
	}
	return total, nil
}

const (
	mult uint8 = iota
	do
	dont
)

type statement struct {
	action uint8
	value  int
	idx    int
}

func puzzle2(rows [][]byte) (any, error) {
	mulReg := regexp.MustCompile(`mul\((\w\w?\w?),(\w\w?\w?)\)`)
	dos := regexp.MustCompile(`do\(\)`)
	donts := regexp.MustCompile(`don't\(\)`)
	var statements []statement
	totalLen := 0
	for _, row := range rows {
		matches := mulReg.FindAllSubmatchIndex(row, -1)
		for _, match := range matches {
			num1, err := strconv.Atoi(string(row[match[2]:match[3]]))
			if err != nil {
				return nil, err
			}
			num2, err := strconv.Atoi(string(row[match[4]:match[5]]))
			if err != nil {
				return nil, err
			}
			entry := statement{
				action: mult,
				value:  num1 * num2,
				idx:    totalLen + match[0],
			}
			statements = append(statements, entry)
		}
		matches = dos.FindAllIndex(row, -1)
		for _, match := range matches {
			entry := statement{
				action: do,
				idx:    totalLen + match[0],
			}
			statements = append(statements, entry)
		}
		matches = donts.FindAllIndex(row, -1)
		for _, match := range matches {
			entry := statement{
				action: dont,
				idx:    totalLen + match[0],
			}
			statements = append(statements, entry)
		}
		totalLen += len(row)
	}
	slices.SortFunc(statements, func(a, b statement) int {
		return cmp.Compare(a.idx, b.idx)
	})
	doIt := true
	total := 0
	for _, stmt := range statements {
		switch stmt.action {
		case do:
			doIt = true
		case dont:
			doIt = false
		case mult:
			if doIt {
				total += stmt.value
			}
		}
	}

	return total, nil
}
