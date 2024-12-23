package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"time"

	"github.com/KevinJoiner/AdventOfCode2024/aoc"
)

var minCache = map[[3]byte]int{}

func main() {
	rows, err := aoc.ReadLines("./day21/input.txt")
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
	sum := 0
	var stringRows []string
	for _, row := range rows {
		total := 0
		start := byte('A')
		for i := 0; i < len(row); i++ {
			total += minPaths(start, row[i], 2, true)
			start = row[i]
		}
		keyPadVal, err := strconv.Atoi(string(row[:len(row)-1]))
		if err != nil {
			panic(err)
		}
		sum += total * keyPadVal
		stringRows = append(stringRows, string(row))
	}
	return sum
}

func puzzle2(rows [][]byte) any {
	sum := 0
	var stringRows []string
	for _, row := range rows {
		total := 0
		start := byte('A')
		for i := 0; i < len(row); i++ {
			total += minPaths(start, row[i], 25, true)
			start = row[i]
		}
		keyPadVal, err := strconv.Atoi(string(row[:len(row)-1]))
		if err != nil {
			panic(err)
		}
		sum += total * keyPadVal
		stringRows = append(stringRows, string(row))
	}
	return sum
}

func minPaths(start, end byte, depth int, first bool) int {
	val, ok := minCache[[3]byte{start, end, byte(depth)}]
	if ok {
		return val
	}

	if depth == 0 {
		dirs := getAllDirs(start, end)
		minDir := math.MaxInt
		for _, dir := range dirs {
			minDir = min(minDir, len(dir))
		}
		return minDir
	}
	var allDirs [][]byte
	if first {
		allDirs = getAllKeyPadDirs(start, end)
	} else {
		allDirs = getAllDirs(start, end)
	}
	minDir := math.MaxInt
	for _, dirs := range allDirs {
		total := 0
		s := byte('A')
		for i := 0; i < len(dirs); i++ {
			total += minPaths(s, dirs[i], depth-1, false)
			s = dirs[i]
		}
		minDir = min(minDir, total)
	}
	minCache[[3]byte{start, end, byte(depth)}] = minDir
	return minDir
}

func getAllDirs(start, end byte) [][]byte {
	tar := dirPad[end]
	first := dirPad[start]
	xDiff := tar.x - first.x
	yDiff := tar.y - first.y
	if (start == '^' || start == 'A') && end == '<' {
		return [][]byte{append(append(goVert(xDiff), goHorz(yDiff)...), 'A')}
	}
	if (end == '^' || end == 'A') && start == '<' {
		return [][]byte{append(append(goHorz(yDiff), goVert(xDiff)...), 'A')}
	}

	return [][]byte{
		append(append(goVert(xDiff), goHorz(yDiff)...), 'A'),
		append(append(goHorz(yDiff), goVert(xDiff)...), 'A'),
	}
}

func getAllKeyPadDirs(start, end byte) [][]byte {
	tar := keyPad[end]
	first := keyPad[start]
	xDiff := tar.x - first.x
	yDiff := tar.y - first.y
	if (start == 'A' || start == '0') && (end == '7' || end == '4' || end == '1') {
		return [][]byte{append(append(goVert(xDiff), goHorz(yDiff)...), 'A')}
	}
	if (end == 'A' || end == '0') && (start == '7' || start == '4' || start == '1') {
		return [][]byte{append(append(goHorz(yDiff), goVert(xDiff)...), 'A')}
	}

	return [][]byte{
		append(append(goVert(xDiff), goHorz(yDiff)...), 'A'),
		append(append(goHorz(yDiff), goVert(xDiff)...), 'A'),
	}
}

// func dirsForKeyPad(key byte, start Loc) ([]byte, Loc) {
// 	tar := keyPad[key]
// 	xDiff := tar.x - start.x
// 	yDiff := tar.y - start.y
// 	ret := []byte{}

// 	if xDiff > 0 {
// 		ret = append(ret, goHorz(yDiff)...)
// 		ret = append(ret, goVert(xDiff)...)
// 	} else {
// 		ret = append(ret, goVert(xDiff)...)
// 		ret = append(ret, goHorz(yDiff)...)
// 	}
// 	ret = append(ret, 'A')
// 	return ret, tar

// }
func dirsForDirPad(key byte, start Loc, depth int) ([]byte, Loc) {
	tar := dirPad[key]
	xDiff := tar.x - start.x
	yDiff := tar.y - start.y
	ret := []byte{}

	if xDiff > 0 {
		ret = append(ret, goVert(xDiff)...)
		ret = append(ret, goHorz(yDiff)...)
	} else {
		ret = append(ret, goHorz(yDiff)...)
		ret = append(ret, goVert(xDiff)...)
	}
	ret = append(ret, 'A')
	return ret, tar

}
func goVert(xDiff int) []byte {
	var ret []byte
	if xDiff < 0 {
		for range -xDiff {
			ret = append(ret, '^')
		}
	} else {
		for range xDiff {
			ret = append(ret, 'v')
		}
	}
	return ret
}

// <A^A>^^AvvvA
// <A^A^^<A>>vvvA
func goHorz(yDiff int) []byte {
	var ret []byte
	if yDiff < 0 {
		for range -yDiff {
			ret = append(ret, '<')
		}
	} else {
		for range yDiff {
			ret = append(ret, '>')
		}
	}
	return ret
}

type Loc struct {
	x, y int
}

var keyPad = map[byte]Loc{
	'7': {0, 0}, '8': {0, 1}, '9': {0, 2},
	'4': {1, 0}, '5': {1, 1}, '6': {1, 2},
	'1': {2, 0}, '2': {2, 1}, '3': {2, 2},
	'_': {3, 0}, '0': {3, 1}, 'A': {3, 2},
}
var locToKey = func() map[Loc]byte {
	rev := map[Loc]byte{}
	for k, v := range keyPad {
		rev[v] = k
	}
	return rev
}()

var dirPad = map[byte]Loc{
	'_': {0, 0}, '^': {0, 1}, 'A': {0, 2},
	'<': {1, 0}, 'v': {1, 1}, '>': {1, 2},
}

var locToDir = func() map[Loc]byte {
	rev := map[Loc]byte{}
	for k, v := range dirPad {
		rev[v] = k
	}
	return rev
}()

// // find the diff between two dirs
// func compareDirs(a, b string) {
// 	dirs := strings.Split(a, "A")
// 	bDirs := strings.Split(b, "A")
// 	for i, dir := range dirs {
// 		aset := map[rune]int{}
// 		bset := map[rune]int{}
// 		for _, d := range dir {
// 			aset[d]++
// 		}
// 		for _, d := range bDirs[i] {
// 			bset[d]++
// 		}
// 		for k, v := range aset {
// 			if v != bset[k] {
// 				fmt.Printf("Diff at %d: %c %d %d\n", i, k, v, bset[k])
// 			}
// 			fmt.Printf("%s, %s", string)
// 		}
// 	}
// }
