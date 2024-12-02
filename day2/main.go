package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"strconv"

	"github.com/KevinJoiner/AdventOfCode2024/aoc"
)

func main() {
	rows, err := aoc.ReadLines("./day2/input.txt")
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

func puzzle1(lines [][]byte) (any, error) {
	rows := toInts(lines)
	total := 0
	for _, row := range rows {
		if allGood(row) {
			total++
		}
	}
	return total, nil
}

func puzzle2(lines [][]byte) (any, error) {
	rows := toInts(lines)
	total := 0
	for _, row := range rows {
		if allGood(row) {
			total++
			continue
		}
		for j := range row {
			newList := make([]int, len(row[:j]))
			copy(newList, row[:j])
			newList = append(newList, row[j+1:]...)
			if allGood(newList) {
				total++
				break
			}
		}
	}
	return total, nil
}
func toInts(raw [][]byte) [][]int {
	ret := make([][]int, 0, len(raw))
	for _, line := range raw {
		var row []int
		scanner := bufio.NewScanner(bytes.NewReader(line))
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			val, err := strconv.Atoi(scanner.Text())
			if err != nil {
				panic(fmt.Sprintf("failed to convert input to int: %s", err))
			}
			row = append(row, val)
		}
		ret = append(ret, row)
	}
	return ret
}

func allGood(row []int) bool {
	dir := 0
	first := row[0]
	for i := 1; i < len(row); i++ {
		next := row[i]
		diff := first - next
		if diff < 0 {
			if dir == 0 {
				dir = -1
			} else if dir == 1 {
				return false
			}
			diff = -diff
		} else {
			if dir == 0 {
				dir = 1
			} else if dir == -1 {
				return false
			}
		}
		if diff < 1 || diff > 3 {
			return false
		}
		first = next
	}
	return true
}
