package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"math"
	"strconv"
	"time"

	"github.com/KevinJoiner/AdventOfCode2024/aoc"
)

func main() {
	rows, err := aoc.ReadLines("./day7/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	start := time.Now()
	output, err := puzzle1(rows)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Puzzle 1 output:", output)
	fmt.Println("Puzzle 1 Duration:", time.Since(start))
	start = time.Now()
	output, err = puzzle2(rows)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Puzzle 2 output:", output)
	fmt.Println("Puzzle 2 Duration:", time.Since(start))
}

func puzzle1(rows [][]byte) (any, error) {
	sum := 0
	for _, row := range rows {
		testVal, vals := totalsAndNumbers(row)
		if check(testVal, vals[0], vals[1:]) {
			sum += testVal
		}
	}
	return sum, nil
}

func puzzle2(rows [][]byte) (any, error) {
	sum := 0
	for _, row := range rows {
		testVal, vals := totalsAndNumbers(row)
		if check2(testVal, vals[0], vals[1:]) {
			sum += testVal
		}
	}
	return sum, nil
}

func totalsAndNumbers(row []byte) (int, []int) {
	input := bytes.Split(row, []byte(":"))
	testVal, err := strconv.Atoi(string(input[0]))
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(bytes.NewReader(input[1]))
	scanner.Split(bufio.ScanWords)
	var vals []int
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		vals = append(vals, val)
	}
	return testVal, vals
}

func check(testVal int, startNum int, vals []int) bool {
	if len(vals) == 0 {
		return testVal == startNum
	}

	tryVal := startNum + vals[0]
	if check(testVal, tryVal, vals[1:]) {
		return true
	}

	tryVal = startNum * vals[0]
	return check(testVal, tryVal, vals[1:])
}

func check2(testVal int, startNum int, vals []int) bool {
	if len(vals) == 0 {
		return testVal == startNum
	}
	tryVal := startNum + vals[0]
	if check2(testVal, tryVal, vals[1:]) {
		return true
	}

	tryVal = startNum * vals[0]
	if check2(testVal, tryVal, vals[1:]) {
		return true
	}
	tryVal = concat(startNum, vals[0])
	return check2(testVal, tryVal, vals[1:])

}

func concat(a, b int) int {
	n := 0
	count := 0
	for b != 0 {
		r := b % 10
		n += r * int(math.Pow10(count))
		b = b / 10
		count++
	}
	a *= int(math.Pow10(count))
	a += n
	return a
}
