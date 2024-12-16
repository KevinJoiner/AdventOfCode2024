package main

import (
	"fmt"
	"log"
	"time"

	"github.com/KevinJoiner/AdventOfCode2024/aoc"
)

func main() {
	rows, err := aoc.ReadLines("./day15/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	rows2, _ := aoc.ReadLines("./day15/input.txt")
	start := time.Now()
	output := puzzle1(rows)
	fmt.Println("Puzzle 1 output:", output)
	fmt.Println("Puzzle 1 Duration:", time.Since(start))
	start = time.Now()
	output = puzzle2(rows2)
	fmt.Println("Puzzle 2 output:", output)
	fmt.Println("Puzzle 2 Duration:", time.Since(start))
}

func puzzle1(rows [][]byte) any {
	grid, dirs, curr := getInput(rows)
	printgrid(grid)
	for _, dir := range dirs {
		if move(grid, dir, curr) {
			curr = curr.next(dir)
		}
	}
	score := 0
	for x := range grid {
		for y := range grid[x] {
			if grid[x][y] == 'O' {
				score += 100*x + y
			}
		}
	}
	printgrid(grid)
	return score
}

func puzzle2(rows [][]byte) any {
	grid, dirs, curr := getInput2(rows)
	printgrid(grid)
	for _, dir := range dirs {
		if move2(grid, dir, curr, true, false) {
			curr = curr.next(dir)
		}
	}
	score := 0
	for x := range grid {
		for y := range grid[x] {
			if grid[x][y] == '[' {
				score += 100*x + y
			}
		}
	}
	printgrid(grid)
	return score
}

func printgrid(grid [][]byte) {
	for i := range grid {
		for j := range grid[i] {
			fmt.Print(string(grid[i][j]))
		}
		fmt.Println()
	}
}

func move(grid [][]byte, dir byte, curr Loc) bool {
	nextSpot := curr.next(dir)
	nextItem := grid[nextSpot.x][nextSpot.y]
	if nextItem == '#' {
		return false
	}
	if nextItem == 'O' && !move(grid, dir, nextSpot) {
		return false
	}
	grid[nextSpot.x][nextSpot.y] = grid[curr.x][curr.y]
	grid[curr.x][curr.y] = '.'
	return true
}

func move2(grid [][]byte, dir byte, curr Loc, write bool, pair bool) bool {
	nextSpot := curr.next(dir)
	nextItem := grid[nextSpot.x][nextSpot.y]
	if nextItem == '#' {
		return false
	}
	currItem := grid[curr.x][curr.y]
	checkPair := (currItem == '[' || currItem == ']') && (dir == upDir || dir == downDir) && !pair
	if checkPair {
		moveDir := rightDir
		if currItem == ']' {
			moveDir = leftDir
		}
		if !move2(grid, dir, curr.next(moveDir), false, true) {
			return false
		}
		if nextItem == '[' || nextItem == ']' {
			if !move2(grid, dir, nextSpot, write, false) {
				return false
			}
		}
		if write {
			_ = move2(grid, dir, curr.next(moveDir), true, true)
		}
	} else if (nextItem == '[' || nextItem == ']') && !move2(grid, dir, nextSpot, write, false) {
		return false
	}

	if !write {
		return true
	}
	// fmt.Println()
	// fmt.Println(string(dir), string(nextItem))
	// printgrid(grid)
	grid[nextSpot.x][nextSpot.y] = grid[curr.x][curr.y]
	// printgrid(grid)
	grid[curr.x][curr.y] = '.'
	// printgrid(grid)
	return true
}
func getInput2(rows [][]byte) (grid [][]byte, dirs []byte, start Loc) {
	storeGrid := true
	for x, row := range rows {
		if len(row) == 0 {
			storeGrid = false
		}
		if storeGrid {
			newRow := make([]byte, len(row)*2)
			for y := range row {
				val := row[y]
				if val == 'O' {
					newRow[y*2] = '['
				} else {
					newRow[y*2] = val
				}
				if val == '@' {
					newRow[y*2+1] = '.'
					start = Loc{x, y * 2}
				} else if val == 'O' {
					newRow[y*2+1] = ']'
				} else {
					newRow[y*2+1] = row[y]
				}
			}
			grid = append(grid, newRow)
			continue
		}
		dirs = append(dirs, row...)
	}
	return grid, dirs, start
}
func getInput(rows [][]byte) (grid [][]byte, dirs []byte, start Loc) {
	storeGrid := true
	for x, row := range rows {
		if len(row) == 0 {
			storeGrid = false
		}
		if storeGrid {
			grid = append(grid, row)
			for y := range row {
				if row[y] == '@' {
					start = Loc{x, y}
				}
			}
			continue
		}
		dirs = append(dirs, row...)
	}
	return grid, dirs, start
}

type Loc struct {
	x, y int
}

const (
	leftDir  = byte('<')
	upDir    = byte('^')
	rightDir = byte('>')
	downDir  = byte('v')
)

func (l Loc) next(dir byte) Loc {
	switch dir {
	case leftDir:
		return l.left()
	case upDir:
		return l.up()
	case rightDir:
		return l.right()
	case downDir:
		return l.down()
	}
	panic("what is even that?")
}

func (l Loc) left() Loc {
	return Loc{
		x: l.x,
		y: l.y - 1,
	}
}
func (l Loc) up() Loc {
	return Loc{
		x: l.x - 1,
		y: l.y,
	}
}
func (l Loc) right() Loc {
	return Loc{
		x: l.x,
		y: l.y + 1,
	}
}
func (l Loc) down() Loc {
	return Loc{
		x: l.x + 1,
		y: l.y,
	}
}
