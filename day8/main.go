package main

import (
	"fmt"
	"log"
	"time"

	"github.com/KevinJoiner/AdventOfCode2024/aoc"
)

func main() {
	rows, err := aoc.ReadLines("./day8/input.txt")
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
	matches := getMatches(rows)
	seen := map[Loc]struct{}{}
	for _, vals := range matches {
		for i := 0; i < len(vals); i++ {
			for j := i + 1; j < len(vals); j++ {
				anti1, anti2 := getAntinodes(vals[i], vals[j])
				if inBounds(rows, anti1) {
					seen[anti1] = struct{}{}
				}
				if inBounds(rows, anti2) {
					seen[anti2] = struct{}{}
				}
			}
		}
	}
	return len(seen)
}

func puzzle2(rows [][]byte) any {
	matches := getMatches(rows)
	seen := map[Loc]struct{}{}
	for _, vals := range matches {
		for i := 0; i < len(vals); i++ {
			for j := i + 1; j < len(vals); j++ {
				nodes := getResonantHarmonicsAntinodesfunc(rows, vals[i], vals[j])
				for _, node := range nodes {
					seen[node] = struct{}{}
				}
			}
		}
	}
	return len(seen)
}

type Loc struct {
	x int
	y int
}

func getMatches(grid [][]byte) map[byte][]Loc {
	matches := map[byte][]Loc{}
	for r := range grid {
		for c, val := range grid[r] {
			if val != '.' {
				matches[val] = append(matches[val], Loc{r, c})
			}
		}
	}
	return matches
}

// . . . #
// . . A .
// . A . .
// # . . .
// 1, 2
// 2, 1
func getAntinodes(a, b Loc) (Loc, Loc) {
	xDiff := a.x - b.x // -1
	yDiff := a.y - b.y // 1
	anti1 := Loc{
		x: a.x + xDiff,
		y: a.y + yDiff,
	}
	anti2 := Loc{
		x: b.x + -xDiff,
		y: b.y + -yDiff,
	}
	return anti1, anti2
}

func getResonantHarmonicsAntinodesfunc(grid [][]byte, a, b Loc) []Loc {
	xDiff := a.x - b.x // -1
	yDiff := a.y - b.y // 1

	// seems to work without but we should probably reduce diffs like 8/4 to 2/1 since those are also in line
	gcd := GCD(xDiff, yDiff)
	if gcd != 0 {
		gcd = abs(gcd)
		// fmt.Printf("reducing %d, %d to %d, %d\n", xDiff, yDiff, xDiff/gcd, yDiff/gcd)
		xDiff /= gcd
		yDiff /= gcd
	}

	xDiffTemp, yDiffTemp := xDiff, yDiff
	var retNodes []Loc
	anti := a
	for inBounds(grid, anti) {
		retNodes = append(retNodes, anti)
		anti = Loc{
			x: a.x + xDiffTemp,
			y: a.y + yDiffTemp,
		}
		xDiffTemp += xDiff
		yDiffTemp += yDiff
	}

	xDiffTemp, yDiffTemp = -xDiff, -yDiff
	anti = b
	for inBounds(grid, anti) {
		retNodes = append(retNodes, anti)
		anti = Loc{
			x: b.x + xDiffTemp,
			y: b.y + yDiffTemp,
		}
		xDiffTemp += -xDiff
		yDiffTemp += -yDiff
	}
	return retNodes
}
func inBounds(grid [][]byte, l Loc) bool {
	return l.x > -1 && l.x < len(grid) && l.y > -1 && l.y < len(grid[0])
}

func printGrid(grid [][]byte, seen map[Loc]struct{}) {
	fmt.Printf("\n\n")
	fmt.Printf("\n\n")
	for r, row := range grid {
		for c, val := range row {
			if val != '.' {
				fmt.Printf(string(val))
			} else if _, ok := seen[Loc{r, c}]; ok {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
