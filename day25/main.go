package main

import (
	"fmt"
	"log"
	"time"

	"github.com/KevinJoiner/AdventOfCode2024/aoc"
)

func main() {
	rows, err := aoc.ReadLines("./day25/input.txt")
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
	locks, keys := getInput(rows)
	pairs := 0
	for _, lock := range locks {
		for _, key := range keys {
			if fits(lock, key) {
				pairs++
			}
		}
	}
	return pairs
}

func puzzle2(rows [][]byte) any {
	return "Feels Good"
}

func getInput(rows [][]byte) ([][5]int, [][5]int) {
	var locks, keys [][5]int
	isLock := false
	start := -1
	vals := [5]int{}
	for _, row := range rows {
		if len(row) == 0 {
			if isLock {
				locks = append(locks, vals)
			} else {
				keys = append(keys, vals)
			}
			start = -1
			vals = [5]int{}
			continue
		}
		if start == -1 {
			if row[0] == '#' {
				isLock = true
			} else {
				isLock = false
			}
			start = 0
			continue
		}
		start++
		if !isLock && start == 6 {
			continue
		}
		for j := range row {
			if row[j] == '#' {
				vals[j]++
			}
		}
	}

	if isLock {
		locks = append(locks, vals)
	} else {
		keys = append(keys, vals)
	}
	return locks, keys
}

func fits(lock, key [5]int) bool {
	for i := range 5 {
		if lock[i]+key[i] > 5 {
			return false
		}
	}
	return true
}
