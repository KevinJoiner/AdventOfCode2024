package main

import (
	"cmp"
	"container/heap"
	"fmt"
	"log"
	"time"

	"github.com/KevinJoiner/AdventOfCode2024/aoc"
)

func main() {
	rows, err := aoc.ReadLines("/Users/kevinjoiner/dev/kevinjoiner/AdventOfCode2024/day16/input.txt")
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
	grid, start, end := getInput(rows)
	// cache := map[Loc]int{}
	lastNode := iAmAStar(grid, start, end)
	printgrid(grid, lastNode)
	return lastNode.gCost
}

func puzzle2(rows [][]byte) any {
	count := 0
	for _, row := range rows {
		for _, val := range row {
			if val == 'O' {
				count++
			}
		}
	}
	return count
}
func getInput(rows [][]byte) ([][]byte, Loc, Loc) {
	var start, end Loc
	for x, row := range rows {
		for y := range row {
			if row[y] == 'S' {
				start = Loc{x, y, rightDir}
			}
			if row[y] == 'E' {
				end = Loc{x, y, rightDir}
			}
		}
	}
	return rows, start, end
}

type Loc struct {
	x, y int
	dir  byte
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
func (l Loc) move() Loc {
	return l.next(l.dir)
}

func (l Loc) turn() Loc {
	newDir := l
	switch l.dir {
	case leftDir:
		newDir.dir = downDir
	case upDir:
		newDir.dir = leftDir
	case rightDir:
		newDir.dir = upDir
	case downDir:
		newDir.dir = rightDir
	}
	return newDir
}

func (l Loc) left() Loc {
	return Loc{
		x:   l.x,
		y:   l.y - 1,
		dir: l.dir,
	}
}
func (l Loc) up() Loc {
	return Loc{
		x:   l.x - 1,
		y:   l.y,
		dir: l.dir,
	}
}
func (l Loc) right() Loc {
	return Loc{
		x:   l.x,
		y:   l.y + 1,
		dir: l.dir,
	}
}
func (l Loc) down() Loc {
	return Loc{
		x:   l.x + 1,
		y:   l.y,
		dir: l.dir,
	}
}
func inBounds(grid [][]byte, l Loc) bool {
	// heap.Interface
	return l.x > -1 && l.x < len(grid) && l.y > -1 && l.y < len(grid[0])
}

func canVisit(grid [][]byte, l Node) bool {
	return inBounds(grid, l.Loc) && grid[l.x][l.y] != '#'
}

var i = 0

func iAmAStar(grid [][]byte, start, end Loc) *Node {
	visted := map[Loc]*Node{}
	startNode := Node{Loc: start}
	pq := &Heap[int]{values: []hasOrder[int]{&startNode}}
	heap.Init(pq)
	visted[start] = &startNode
	for pq.Len() > 0 {
		curr := heap.Pop(pq).(*Node)
		if curr.Loc.x == end.x && curr.Loc.y == end.y {
			return curr
		}
		lastNode := curr
		if curr == nil {
			panic(2)
		}
		for i := range 4 {
			var nextNode *Node
			if i == 0 {
				nextNode = &Node{
					Loc:   curr.move(),
					gCost: curr.gCost + 1,
				}
			} else {
				nextNode = &Node{
					Loc: lastNode.turn(),
				}
				if i == 2 {
					nextNode.gCost = curr.gCost + 2000
				} else {
					nextNode.gCost = curr.gCost + 1000
				}
				lastNode = nextNode
			}

			// nextNode.hDist = manhatan(nextNode.Loc, end)
			nextNode.f = nextNode.hDist + nextNode.gCost
			nextNode.parents = []*Node{curr}
			prevNode, ok := visted[nextNode.Loc]
			if ok {
				if nextNode.f < prevNode.f {
					ok = false
				} else if nextNode.f == prevNode.f {
					prevNode.parents = append(prevNode.parents, curr)
					visted[prevNode.Loc] = prevNode
				}
			}
			if !ok && canVisit(grid, *nextNode) {
				visted[nextNode.Loc] = nextNode
				heap.Push(pq, nextNode)
			} else {
				// fmt.Printf("Can Not visit %+v\n", nextNode)
			}
		}

	}
	return nil
}

func manhatan(start, end Loc) int {
	return 0
}

type Node struct {
	Loc
	f       int
	gCost   int
	hDist   int
	parents []*Node
}

func printgrid(grid [][]byte, curr *Node) {
	visit := []*Node{curr}
	for len(visit) != 0 {
		curr := visit[0]
		visit = visit[1:]
		grid[curr.x][curr.y] = 'O'
		if len(curr.parents) > 1 {
			fmt.Println(curr)
		}
		visit = append(visit, curr.parents...)
	}
	for _, row := range grid {
		fmt.Println(string(row))
	}
}
func (n *Node) order() int { return n.f }

type hasOrder[T cmp.Ordered] interface {
	order() T
}
type Heap[T cmp.Ordered] struct {
	values []hasOrder[T]
}

func (h Heap[T]) Len() int {
	return len(h.values)
}

func (h Heap[T]) Less(i, j int) bool {
	return h.values[i].order() < h.values[j].order()
}

func (h Heap[T]) Swap(i, j int) {
	h.values[i], h.values[j] = h.values[j], h.values[i]
}

func (h *Heap[T]) Push(x any) {
	h.values = append(h.values, x.(hasOrder[T]))
}
func (h *Heap[T]) Pop() any {
	n := len(h.values) - 1
	ret := h.values[n]
	h.values = h.values[:n]
	return ret
}
