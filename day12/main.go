package main

import (
	"fmt"
	"log"
	"time"

	"github.com/KevinJoiner/AdventOfCode2024/aoc"
)

func main() {
	rows, err := aoc.ReadLines("./day12/input.txt")
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
	visted := map[Loc]struct{}{}
	total := 0
	for x, row := range rows {
		for y := range row {
			loc := Loc{x, y}
			childArea, childPerm := dfs(rows, loc, visted)
			total += childArea * childPerm
		}
	}
	return total
}

func puzzle2(rows [][]byte) any {
	visted := map[Loc]Loc{}
	total := 0
	areas := map[Loc]int{}
	sides := map[Loc]int{}
	for x, row := range rows {
		for y := range row {
			loc := Loc{x, y}
			if _, ok := visted[loc]; ok {
				continue
			}
			area := dfs4(rows, loc, loc, visted)
			areas[loc] = area
		}
	}
	for x, row := range rows {
		for y := range row {
			loc := Loc{x, y}
			corners := sidesOnly(rows, loc, visted)
			sides[visted[loc]] += corners
		}
	}
	for k, v := range areas {
		s := sides[k]
		total += v * s
	}
	return total
}

func dfs(grid [][]byte, cur Loc, visted map[Loc]struct{}) (int, int) {
	if _, ok := visted[cur]; ok {
		return 0, 0
	}
	curVal := grid[cur.x][cur.y]
	visted[cur] = struct{}{}
	area := 1
	perm := 0
	next := cur.left()
	if inBounds(grid, next) && grid[next.x][next.y] == curVal {
		childArea, childPerm := dfs(grid, next, visted)
		area += childArea
		perm += childPerm
	} else {
		perm++
	}
	next = cur.right()
	if inBounds(grid, next) && grid[next.x][next.y] == curVal {
		childArea, childPerm := dfs(grid, next, visted)
		area += childArea
		perm += childPerm
	} else {
		perm++
	}
	next = cur.down()
	if inBounds(grid, next) && grid[next.x][next.y] == curVal {
		childArea, childPerm := dfs(grid, next, visted)
		area += childArea
		perm += childPerm
	} else {
		perm++
	}
	next = cur.up()
	if inBounds(grid, next) && grid[next.x][next.y] == curVal {
		childArea, childPerm := dfs(grid, next, visted)
		area += childArea
		perm += childPerm
	} else {
		perm++
	}
	return area, perm
}
func dfs4(grid [][]byte, start, cur Loc, visted map[Loc]Loc) int {
	if _, ok := visted[cur]; ok {
		return 0
	}
	curVal := grid[cur.x][cur.y]
	visted[cur] = start
	area := 1
	next := cur.right()
	if match(grid, next, curVal) {
		area += dfs4(grid, start, next, visted)
	}
	next = cur.down()
	if match(grid, next, curVal) {
		area += dfs4(grid, start, next, visted)
	}
	next = cur.left()
	if match(grid, next, curVal) {
		area += dfs4(grid, start, next, visted)
	}
	next = cur.up()
	if match(grid, next, curVal) {
		area += dfs4(grid, start, next, visted)
	}
	return area
}

func sidesOnly(grid [][]byte, cur Loc, visted map[Loc]Loc) int {
	sides := 0
	if !match2(visted, cur.left(), cur) && !match2(visted, cur.up(), cur) {
		sides++ //outside
	}
	if !match2(visted, cur.right(), cur) && !match2(visted, cur.up(), cur) {
		sides++ //outside
	}
	if !match2(visted, cur.left(), cur) && !match2(visted, cur.down(), cur) {
		sides++ //outside
	}
	if !match2(visted, cur.right(), cur) && !match2(visted, cur.down(), cur) {
		sides++ //outside
	}
	if !match2(visted, cur.up().left(), cur) && match2(visted, cur.left(), cur) && match2(visted, cur.up(), cur) {
		sides++ //inside
	}
	if !match2(visted, cur.up().right(), cur) && match2(visted, cur.right(), cur) && match2(visted, cur.up(), cur) {
		sides++ //inside
	}
	if !match2(visted, cur.down().left(), cur) && match2(visted, cur.left(), cur) && match2(visted, cur.down(), cur) {
		sides++ //inside
	}
	if !match2(visted, cur.down().right(), cur) && match2(visted, cur.right(), cur) && match2(visted, cur.down(), cur) {
		sides++ //inside
	}
	return sides

}
func match2(visted map[Loc]Loc, next, curr Loc) bool {
	start1, ok1 := visted[curr]
	start2, ok2 := visted[next]
	return ok1 == ok2 && start1 == start2
}
func match(grid [][]byte, next Loc, val byte) bool {
	return inBounds(grid, next) && grid[next.x][next.y] == val
}

// func dfs3(grid [][]byte, start *Loc, cur Loc, visted map[Loc]struct{}) (int, int) {
// 	if _, ok := visted[cur]; ok {
// 		if cur == *start {
// 			start.x = -1
// 			// fmt.Printf("%+v\n", visted)
// 		}
// 		return 0, 0
// 	}
// 	curVal := grid[cur.x][cur.y]
// 	visted[cur] = struct{}{}
// 	area := 1
// 	side := 0
// 	next := cur.right()
// 	if inBounds(grid, next) && grid[next.x][next.y] == curVal {
// 		childArea, childPerm := dfs3(grid, start, next, visted)
// 		area += childArea
// 		side += childPerm

// 	} else if start.x != -1 {

// 		side++
// 	}
// 	next = cur.down()
// 	if inBounds(grid, next) && grid[next.x][next.y] == curVal {
// 		childArea, childPerm := dfs3(grid, start, next, visted)
// 		area += childArea
// 		side += childPerm
// 	} else if start.x != -1 {
// 		// side++
// 	}
// 	next = cur.left()
// 	if inBounds(grid, next) && grid[next.x][next.y] == curVal {
// 		childArea, childPerm := dfs3(grid, start, next, visted)
// 		area += childArea
// 		side += childPerm
// 	} else if start.x != -1 {
// 		// side++
// 	}
// 	next = cur.up()
// 	if inBounds(grid, next) && grid[next.x][next.y] == curVal {
// 		childArea, childPerm := dfs3(grid, start, next, visted)
// 		area += childArea
// 		side += childPerm
// 	} else if start.x != -1 {

// 		side++
// 	}
// 	return area, side
// }
// func dfs2(grid [][]byte, start *Loc, cur Loc, prev *State, visted map[Loc]struct{}) (int, int) {
// 	curVal := grid[cur.x][cur.y]
// 	if _, ok := visted[cur]; ok {
// 		if cur == *start {
// 			start.x = -1
// 			// right := cur.right()
// 			// left := cur.down()
// 			// if (inBounds(grid, right) && grid[right.x][right.y] == curVal) && (inBounds(grid, left) && grid[left.x][left.y] == curVal) {
// 			// 	return 0, -1
// 			// }
// 		}
// 		return 0, 0
// 	}
// 	visted[cur] = struct{}{}
// 	area := 1
// 	sides := 0
// 	next := cur.left()
// 	var state State
// 	if !(inBounds(grid, next) && grid[next.x][next.y] == curVal) {
// 		state.l = true
// 		if !prev.l && start.x != -1 {
// 			sides++
// 		}
// 	}
// 	next = cur.right()
// 	if !(inBounds(grid, next) && grid[next.x][next.y] == curVal) {
// 		state.r = true
// 		if !prev.r && start.x != -1 {
// 			sides++
// 		}
// 	}
// 	next = cur.up()
// 	if !(inBounds(grid, next) && grid[next.x][next.y] == curVal) {
// 		state.u = true
// 		if !prev.u && start.x != -1 {
// 			sides++
// 		}
// 	}
// 	next = cur.down()
// 	if !(inBounds(grid, next) && grid[next.x][next.y] == curVal) {
// 		state.d = true
// 		if !prev.d && start.x != -1 {
// 			sides++
// 		}
// 	}
// 	*prev = state
// 	if !state.r {
// 		childArea, childSides := dfs2(grid, start, cur.right(), prev, visted)
// 		area += childArea
// 		sides += childSides
// 	}
// 	if !state.d {
// 		childArea, childSides := dfs2(grid, start, cur.down(), prev, visted)
// 		area += childArea
// 		sides += childSides
// 	}
// 	if !state.l {
// 		childArea, childSides := dfs2(grid, start, cur.left(), prev, visted)
// 		area += childArea
// 		sides += childSides
// 	}
// 	if !state.u {
// 		childArea, childSides := dfs2(grid, start, cur.up(), prev, visted)
// 		area += childArea
// 		sides += childSides
// 	}
// 	return area, sides
// }

//	type State struct {
//		l, r, u, d bool
//	}
type Loc struct {
	x, y int
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
	return l.x > -1 && l.x < len(grid) && l.y > -1 && l.y < len(grid[0])
}
