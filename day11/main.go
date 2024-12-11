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
	rows, err := aoc.ReadLines("./day11/input.txt")
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
	startRocks := toNums(rows)
	endRocks := []int{}
	for _, startRock := range startRocks {
		nums := []int{startRock}
		for j := 0; j < 25; j++ {
			var newNums []int
			for i := 0; i < len(nums); i++ {
				num := nums[i]
				rocks := blink(num)
				newNums = append(newNums, rocks...)
			}
			nums = newNums
		}
		endRocks = append(endRocks, nums...)
	}
	return len(endRocks)
}

func puzzle2(rows [][]byte) any {
	startRocks := toNums(rows)
	cache := map[int]map[int]int{}
	size := 0
	for _, startRock := range startRocks {
		size += interate(cache, startRock, 75)
	}
	return size
}

func interate(cache map[int]map[int]int, num int, iters int) int {
	if iters < 1 {
		return 1
	}
	size := getCacheVal(cache, num, iters)
	if size != 0 {
		return size
	}
	childRocks := blink(num)

	for _, rock := range childRocks {
		size += interate(cache, rock, iters-1)
	}
	storeCacheVal(cache, num, iters, size)
	return size
}
func getCacheVal(cache map[int]map[int]int, num, iter int) int {
	vals, ok := cache[num]
	if !ok {
		return 0
	}
	val, ok := vals[iter]
	if !ok {
		return 0
	}
	return val
}
func storeCacheVal(cache map[int]map[int]int, num, iter int, val int) {
	cachedNum, ok := cache[num]
	if !ok {
		cachedNum = map[int]int{}
	}
	cachedNum[iter] = val
	cache[num] = cachedNum
}

func toNums(rows [][]byte) []int {
	ret := make([]int, 0, len(rows[0]))
	scanner := bufio.NewScanner(bytes.NewReader(rows[0]))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		ret = append(ret, num)
	}
	return ret
}

func blink(num int) []int {
	if num == 0 {
		return []int{1}
	}
	if num == 1 {
		return []int{2024}
	}
	tens := int(math.Floor(math.Log10(float64(num)))) + 1
	if tens&1 == 1 {
		return []int{num * 2024}
	}
	div := int(math.Pow10(tens / 2))
	left := num / div
	right := num % div
	return []int{left, right}
}
