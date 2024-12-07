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
		testVal, _, vals := totalsAndNumbers(row)
		if check(testVal, vals[0], vals[1:], []int{vals[0]}) {
			sum += testVal
		}
	}
	return sum, nil
}

func puzzle2(rows [][]byte) (any, error) {
	sum := 0
	for _, row := range rows {
		testVal, _, vals := totalsAndNumbers(row)
		if check2(testVal, vals[0], vals[1:], []int{vals[0]}) {
			sum += testVal
		}
	}
	return sum, nil
}

func totalsAndNumbers(row []byte) (testVal int, total int, vals []int) {
	input := bytes.Split(row, []byte(":"))
	testVal, err := strconv.Atoi(string(input[0]))
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(bytes.NewReader(input[1]))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		vals = append(vals, val)
		total += val
	}
	// slices.Reverse(vals)
	return testVal, total, vals
}

// func check(testVal int, addVal int, vals []int) bool {
// 	if addVal == testVal {
// 		return true
// 	}
// 	if len(vals) == 0 {
// 		return false
// 	}
// 	fmt.Println(vals[0], "+", vals[1:], "=", addVal, "!=", testVal)
// 	// if all add children are too big break
// 	if addVal > testVal {
// 		fmt.Println("break too high add")
// 		return false
// 	}
// 	rawVal := (addVal - vals[0])
// 	mulNum := rawVal * vals[0]

// 	if mulNum == testVal {
// 		return true
// 	}
// 	fmt.Println(vals[0], "*", vals[1:], "=", mulNum, "!=", testVal)

//		fmt.Println(addVal)
//		fmt.Println()
//		if check(testVal, addVal, vals[1:]) {
//			return true
//		}
//		fmt.Println(mulNum)
//		fmt.Println()
//		if check(testVal, mulNum, vals[1:]) {
//			return true
//		}
//		return false
//	}
func check(testVal int, startNum int, vals []int, seq []int) bool {
	if len(vals) == 0 {
		return testVal == evalSeq(seq)
	}

	prev := evalSeq(seq)
	seq = []int{prev, -1, vals[0]}
	if check(testVal, startNum, vals[1:], seq) {
		return true
	}

	seq = []int{prev, -2, vals[0]}
	if check(testVal, startNum, vals[1:], seq) {
		return true
	}
	return false
}

func check2(testVal int, startNum int, vals []int, seq []int) bool {
	if len(vals) == 0 {
		return testVal == evalSeq(seq)
	}
	prev := evalSeq(seq)
	seq = []int{prev, -1, vals[0]}
	if check2(testVal, startNum, vals[1:], seq) {
		return true
	}
	seq = []int{prev, -2, vals[0]}
	if check2(testVal, startNum, vals[1:], seq) {
		return true
	}
	seq = []int{prev, -3, vals[0]}
	if check2(testVal, startNum, vals[1:], seq) {
		return true
	}
	return false
}

func evalSeq(seq []int) int {
	total := seq[0]
	for i := 1; i < len(seq); i = i + 2 {
		val := seq[i+1]
		if seq[i] == -1 {
			total += val
		} else if seq[i] == -2 {
			total *= val
		} else {
			n := 0
			count := 0
			for val != 0 {
				r := val % 10
				n += r * int(math.Pow10(count))
				val = val / 10
				count++
			}
			total *= int(math.Pow10(count))
			total += n
		}
	}
	return total
}
