package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"sort"
	"strconv"

	"github.com/KevinJoiner/AdventOfCode2024/aoc"
)

func main() {
	rows, err := aoc.ReadLines("day1/input.txt")
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
	list1, list2, err := getList(rows)
	if err != nil {
		return nil, fmt.Errorf("could not get list: %w", err)
	}
	sort.Ints(list1)
	sort.Ints(list2)
	if len(list1) != len(list2) {
		return nil, fmt.Errorf("list lengths do not match")
	}
	total := 0
	for i := range list1 {
		diff := list2[i] - list1[i]
		if diff < 0 {
			diff = -diff
		}
		total += diff
	}

	return total, nil
}

func puzzle2(rows [][]byte) (any, error) {
	lList, rList, err := getList(rows)
	if err != nil {
		return nil, fmt.Errorf("could not get list: %w", err)
	}
	leftCount := map[int]int{}
	for _, val := range lList {
		leftCount[val] = 0
	}
	for _, rVal := range rList {
		count, ok := leftCount[rVal]
		if ok {
			leftCount[rVal] = count + 1
		}
	}
	total := 0
	for lVal, count := range leftCount {
		total += lVal * count
	}
	return total, nil
}

func getList(rows [][]byte) (leftList, rightList []int, err error) {

	list1 := []int{}
	list2 := []int{}
	for i, row := range rows {
		vals := [][]byte{}
		scanner := bufio.NewScanner(bytes.NewReader(row))
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			vals = append(vals, scanner.Bytes())
		}
		if len(vals) != 2 {
			return nil, nil, fmt.Errorf("invalid row unexpected number of values: %d row idx: %d", len(vals), i)
		}
		numVal1, err := strconv.Atoi(string(vals[0]))
		if err != nil {
			return nil, nil, fmt.Errorf("invalid number: %s", vals[0])
		}
		numVal2, err := strconv.Atoi(string(vals[1]))
		if err != nil {
			return nil, nil, fmt.Errorf("invalid number: %s", vals[1])
		}
		list1 = append(list1, numVal1)
		list2 = append(list2, numVal2)
	}
	return list1, list2, nil
}
