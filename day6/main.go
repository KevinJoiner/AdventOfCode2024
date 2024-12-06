package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/KevinJoiner/AdventOfCode2024/aoc"
)

func main() {
	rows, err := aoc.ReadLines("./day6/input.txt")
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
	for x := range rows {
		for y, val := range rows[x] {
			if val == up || val == down || val == left || val == right {
				return walk(rows, val, x, y), nil
			}
		}
	}
	return nil, errors.New("no start position found")
}

func puzzle2(rows [][]byte) (any, error) {
	for x := range rows {
		for y, val := range rows[x] {
			if val == up || val == down || val == left || val == right {
				return longWalk(rows, val, x, y), nil
			}
		}
	}
	return nil, errors.New("no start position found")
}

const (
	up    = '^'
	down  = 'v'
	left  = '<'
	right = '>'
)

func walk(grid [][]byte, dir byte, x, y int) int {
	visted := map[int]struct{}{}
	rowLen := len(grid[0])
	for {
		visted[rowLen*x+y] = struct{}{}
		nextX, nextY := next(dir, x, y)
		if !inBounds(grid, nextX, nextY) {
			break
		}
		if grid[nextX][nextY] == '#' {
			dir = turn(dir)
			continue
		}
		x = nextX
		y = nextY
	}
	return len(visted)
}

func longWalk(grid [][]byte, dir byte, startX, startY int) int {
	visted := map[int][]byte{}
	rowLen := len(grid[0])
	total := 0
	x := startX
	y := startY
	for {
		loc := rowLen*x + y
		nextX, nextY := next(dir, x, y)
		if !inBounds(grid, nextX, nextY) {
			break
		}
		if grid[nextX][nextY] == '#' {
			dir = turn(dir)
			continue
		}
		grid[nextX][nextY] = '#'
		nextIsStart := nextX == startX && nextY == startY
		_, nextHasBeenVisted := visted[nextX*rowLen+nextY]
		if !nextIsStart && !nextHasBeenVisted && shortWalk(grid, dir, x, y, visted) {
			total++
		}
		grid[nextX][nextY] = '.'
		visted[loc] = append(visted[loc], dir)
		x = nextX
		y = nextY
	}
	return total
}

func shortWalk(grid [][]byte, dir byte, x, y int, vistedOld map[int][]byte) bool {
	visted := make(map[int][]byte, len(vistedOld))
	for k, v := range vistedOld {
		visted[k] = v
	}
	rowLen := len(grid[0])
	for {
		loc := rowLen*x + y
		if prevDirs, ok := visted[loc]; ok {
			for _, prevDir := range prevDirs {
				if prevDir == dir {
					return true
				}
			}
		}
		nextX, nextY := next(dir, x, y)
		if !inBounds(grid, nextX, nextY) {
			break
		}
		if grid[nextX][nextY] == '#' {
			dir = turn(dir)
			continue
		}
		visted[loc] = append(visted[loc], dir)
		x = nextX
		y = nextY
	}
	return false
}

func turn(dir byte) byte {
	switch dir {
	case up:
		return right
	case right:
		return down
	case down:
		return left
	case left:
		return up
	default:
		return '_'
	}
}
func next(dir byte, x, y int) (int, int) {
	switch dir {
	case up:
		return x - 1, y
	case down:
		return x + 1, y
	case left:
		return x, y - 1
	case right:
		return x, y + 1
	default:
		return -99, -99
	}
}
func inBounds(grid [][]byte, i, j int) bool {
	return i != -1 && i != len(grid) && j != -1 && j != len(grid[0])
}
