package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/KevinJoiner/AdventOfCode2024/aoc"
)

var gens = 2000

func main() {
	rows, err := aoc.ReadLines("./day22/input.txt")
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
	total := 0
	for _, row := range rows {
		secret, err := strconv.Atoi(string(row))
		if err != nil {
			panic(err)
		}
		secret = brute(secret)
		total += secret
	}
	return total
}

func puzzle2(rows [][]byte) any {
	totals := map[[4]int]int{}
	for _, row := range rows {
		secret, err := strconv.Atoi(string(row))
		if err != nil {
			panic(err)
		}
		seqMap := info(secret)
		for seq, val := range seqMap {
			totals[seq] += val
		}
	}
	best := 0
	for _, val := range totals {
		best = max(best, val)
	}
	return best
}

const pruneVal = 16777216

func basic(secret int) int {
	secret = (secret ^ (secret * 64)) % pruneVal
	secret = (secret ^ (secret / 32)) % pruneVal
	secret = (secret ^ (secret * 2048)) % pruneVal
	return secret
}
func brute(secret int) int {
	for range gens + 1 {
		secret = basic(secret)
	}
	return secret
}

func info(secret int) map[[4]int]int {
	changes := make([]int, gens+1)
	last := 0
	ret := map[[4]int]int{}
	for i := range gens + 1 {
		val := secret % 10
		changes[i] = val - last
		last = val

		if i > 2 {
			seq := [4]int{changes[i-3], changes[i-2], changes[i-1], changes[i]}
			if _, ok := ret[seq]; !ok {
				ret[seq] = val
			}
		}
		secret = basic(secret)
	}
	return ret
}
