package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/KevinJoiner/AdventOfCode2024/aoc"
)

func main() {
	rows, err := aoc.ReadLines("./day10/input.txt")
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

type Loc struct {
	x int
	y int
}

func puzzle1(rows [][]byte) any {
	grid := make([][]int, len(rows))
	cnts := make([][]map[Loc]struct{}, len(rows))
	for i := range rows {
		cnts[i] = make([]map[Loc]struct{}, len(rows[i]))
		grid[i] = make([]int, len(rows[i]))
		for j := range rows[i] {
			val, err := strconv.Atoi(string(rows[i][j]))
			if err != nil {
				panic(err)
			}
			grid[i][j] = val
		}
	}
	total := 0
	for x := range grid {
		for y, val := range grid[x] {
			if val == 0 {
				total += len(find9s(grid, cnts, x, y, -1))
			}
		}
	}

	return total
}

func puzzle2(rows [][]byte) any {
	grid := make([][]int, len(rows))
	cnts := make([][]int, len(rows))
	for i := range rows {
		cnts[i] = make([]int, len(rows[i]))
		grid[i] = make([]int, len(rows[i]))
		for j := range rows[i] {
			val, err := strconv.Atoi(string(rows[i][j]))
			if err != nil {
				panic(err)
			}
			grid[i][j] = val
			cnts[i][j] = -1
		}
	}
	total := 0
	for x := range grid {
		for y, val := range grid[x] {
			if val == 0 {
				total += find9s2(grid, cnts, x, y, -1)
			}
		}
	}

	return total
}

func find9s(grid [][]int, cnts [][]map[Loc]struct{}, x, y int, prev int) map[Loc]struct{} {
	if !inBounds(grid, x, y) {
		return nil
	}
	if grid[x][y] != prev+1 {
		return nil
	}
	if grid[x][y] == 9 {
		return map[Loc]struct{}{
			{x: x, y: y}: {},
		}
	}
	if cnts[x][y] != nil {
		return cnts[x][y]
	}
	cnt := map[Loc]struct{}{}
	curr := grid[x][y]
	locs := find9s(grid, cnts, x-1, y, curr)
	for loc := range locs {
		cnt[loc] = struct{}{}
	}
	locs = find9s(grid, cnts, x, y-1, curr)
	for loc := range locs {
		cnt[loc] = struct{}{}
	}
	locs = find9s(grid, cnts, x+1, y, curr)
	for loc := range locs {
		cnt[loc] = struct{}{}
	}
	locs = find9s(grid, cnts, x, y+1, curr)
	for loc := range locs {
		cnt[loc] = struct{}{}
	}
	cnts[x][y] = cnt
	return cnt
}
func find9s2(grid [][]int, cnts [][]int, x, y int, prev int) int {
	if !inBounds(grid, x, y) {
		return 0
	}
	if grid[x][y] != prev+1 {
		return 0
	}
	if grid[x][y] == 9 {
		return 1
	}
	if cnts[x][y] != -1 {
		return cnts[x][y]
	}
	cnt := 0
	curr := grid[x][y]
	cnt += find9s2(grid, cnts, x-1, y, curr)
	cnt += find9s2(grid, cnts, x, y-1, curr)
	cnt += find9s2(grid, cnts, x+1, y, curr)
	cnt += find9s2(grid, cnts, x, y+1, curr)
	cnts[x][y] = cnt
	return cnt
}
func inBounds(rows [][]int, i, j int) bool {
	return i != -1 && i != len(rows) && j != -1 && j != len(rows[0])
}
