package main

import (
	"fmt"
	"log"
	"time"

	"github.com/KevinJoiner/AdventOfCode2024/aoc"
)

var maxShortCuts = 1

func main() {
	rows, err := aoc.ReadLines("/Users/kevinjoiner/dev/kevinjoiner/AdventOfCode2024/day20/input.txt")
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
	rows, start := getInput(rows)
	n := len(rows)
	grid := make([][]byte, n)
	for x := range n {
		grid[x] = make([]byte, n)
		for y := range n {
			grid[x][y] = rows[x][y]
		}
	}
	maxPaths := cheatersGoingToCheat(grid, start, 2)
	return maxPaths
}

func puzzle2(rows [][]byte) any {
	rows, start := getInput(rows)
	n := len(rows)
	grid := make([][]byte, n)
	for x := range n {
		grid[x] = make([]byte, n)
		for y := range n {
			grid[x][y] = rows[x][y]
		}
	}
	maxPaths := cheatersGoingToCheat(grid, start, 20)
	return maxPaths
}
func getInput(rows [][]byte) ([][]byte, Loc) {
	var start Loc
	for x, row := range rows {
		for y := range row {
			if row[y] == 'S' {
				start = Loc{x, y}
			}
		}
	}
	return rows, start
}

func cheatersGoingToCheat(grid [][]byte, start Loc, cuts int) int {
	maxPaths := 0
	betterbar := 100
	path := toPath(grid, start)
	maxDist := len(path)
	for i, start := range path {
		// minDist := math.MaxInt
		for j := i + 1; j < len(path); j++ {
			end := path[j]
			dist := manhatan(start, end)
			if dist <= cuts {
				// costFromEnd := (maxDist - end.gCost)
				distToEnd := (len(path) - j)
				totalDist := i + dist + distToEnd
				if totalDist+betterbar <= maxDist {
					maxPaths++
				}
			}
		}
		// if minDist+76 == maxDist {
		// 	maxPaths++
		// }
	}
	return maxPaths
}
func manhatan(start, end Loc) int {
	return abs(start.x-end.x) + abs(start.y-end.y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func toPath(grid [][]byte, start Loc) []Loc {
	locs := []Loc{start}
	curr := start
	for {
	checkLoop:
		for _, dir := range []byte{upDir, rightDir, leftDir, downDir} {
			next := curr.next(dir)
			if inBounds(grid, next) && (grid[next.x][next.y] == '.' || grid[next.x][next.y] == 'E') {
				locs = append(locs, next)
				grid[curr.x][curr.y] = ','
				curr = next
				goto checkLoop
			}
		}
		return locs
	}
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
	panic(fmt.Sprintf("what is even that? %s", string(dir)))
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
func inBounds(grid [][]byte, l Loc) bool {
	// heap.Interface
	return l.x > -1 && l.x < len(grid) && l.y > -1 && l.y < len(grid[0])
}
