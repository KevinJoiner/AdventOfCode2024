package main

import (
	"cmp"
	"container/heap"
	"fmt"
	"log"
	"regexp"
	"sort"
	"strconv"
	"time"

	"github.com/KevinJoiner/AdventOfCode2024/aoc"
)

func main() {
	rows, err := aoc.ReadLines("./day18/input.txt")
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
	n := 71
	faults := getInput(rows)
	ns := 1024
	grid := make([][]byte, n)
	for x := range n {
		grid[x] = make([]byte, n)
		for y := range n {
			grid[x][y] = '.'
		}
	}
	for i := range ns {
		grid[faults[i].x][faults[i].y] = '#'
	}
	start := Loc{0, 0}
	end := Loc{n - 1, n - 1}
	endNode := iAmAStar(grid, start, end)
	printgrid(grid, endNode)
	return endNode.gCost
}

func puzzle2(rows [][]byte) any {
	faults := getInput(rows)
	minNs := sort.Search(len(faults), cmpFunc(faults))
	return string(rows[minNs-1])
}
func cmpFunc(faults []Fault) func(i int) bool {
	return func(ns int) bool {
		n := 71
		grid := make([][]byte, n)
		for x := range n {
			grid[x] = make([]byte, n)
			for y := range n {
				grid[x][y] = '.'
			}
		}
		for i := range ns {
			grid[faults[i].x][faults[i].y] = '#'
		}
		start := Loc{0, 0}
		end := Loc{n - 1, n - 1}
		endNode := iAmAStar(grid, start, end)
		if endNode == nil {
			return true
		}
		return false
	}
}

var cordReg = regexp.MustCompile(`([0-9]+),([0-9]+)`)

func getInput(rows [][]byte) []Fault {
	var faults []Fault
	for i, row := range rows {
		matches := cordReg.FindSubmatch(row)
		x, err := strconv.Atoi(string(matches[1]))
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(string(matches[2]))
		if err != nil {
			panic(err)
		}
		f := Fault{
			Loc: Loc{
				x: x,
				y: y,
			},
			ns: i,
		}
		faults = append(faults, f)
	}
	return faults
}
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
		if curr == nil {
			panic(2)
		}
		for _, dir := range []byte{upDir, rightDir, leftDir, downDir} {
			nextNode := &Node{
				Loc:   curr.next(dir),
				gCost: curr.gCost + 1,
			}

			// nextNode.hDist = manhatan(nextNode.Loc, end)
			nextNode.f = nextNode.hDist + nextNode.gCost
			nextNode.parents = []*Node{curr}
			prevNode, ok := visted[nextNode.Loc]
			if ok {
				if nextNode.f < prevNode.f {
					ok = false
				}
				// }else if nextNode.f == prevNode.f {
				// 	prevNode.parents = append(prevNode.parents, curr)
				// 	visted[prevNode.Loc] = prevNode
				// }
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
	return abs(start.x-end.x) + abs(start.y-end.y)
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
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
		if curr == nil {
			continue
		}
		grid[curr.x][curr.y] = 'O'
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

type Loc struct {
	x, y int
}
type Fault struct {
	Loc
	ns int
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

func canVisit(grid [][]byte, l Node) bool {
	return inBounds(grid, l.Loc) && grid[l.x][l.y] != '#'
}
