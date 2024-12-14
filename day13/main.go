package main

import (
	"fmt"
	"log"
	"math"
	"regexp"
	"strconv"
	"time"

	"github.com/KevinJoiner/AdventOfCode2024/aoc"
)

func main() {
	rows, err := aoc.ReadLines("./day13/input.txt")
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
	machs := getInput(rows)
	total := 0
	for _, mach := range machs {
		// 	if mach.a.x*101+mach.b.x*101 < mach.target.x || mach.a.y*101+mach.b.y*101 < mach.target.y {

		// 		minToks := play(&mach, mach.target, cache)
		// 		fmt.Printf("%+v\n mach min: %+v\n", mach, minToks)
		// 		continue
		// 	}
		cache := map[Loc]State{}
		minToks := play(&mach, mach.target, cache)
		if minToks.toks != math.MaxInt {
			total += minToks.toks * mach.gcd
		}
		// fmt.Printf("%+v\n mach min: %d\n", mach, minToks)
		// fmt.Println(total)
	}
	return total
}

func puzzle2(rows [][]byte) any {
	machs := getInput(rows)
	total := 0
	for _, mach := range machs {
		tx, ty := mach.other.x, mach.other.y

		det := (mach.a.y*mach.b.x - mach.a.x*mach.b.y)

		aSteps := (mach.b.x*ty - mach.b.y*tx) / det
		bSteps := (mach.a.y*tx - mach.a.x*ty) / det
		const epsilon = 1e-10
		if math.Abs(aSteps-math.Trunc(aSteps)) < epsilon &&
			math.Abs(bSteps-math.Trunc(bSteps)) < epsilon {
			total += int(aSteps)*3 + int(bSteps)
		}
		// fmt.Println(strconv.FormatFloat(cost, 'f', 10, 64))
		// total += math.Round(cost)
	}
	return total
}

//161619590963218
//69532602166545
//104015411578548

// 18 = 180 - y*z
// 18-180 = -y*z
// 18-180 / -y =z
var iter = 0

type State struct {
	toks, aPress, bPress int
}

// (A*x) + (B*y) = T
func play(mach *Machine, target Loc, cache map[Loc]State) State {
	state, ok := cache[target]
	if ok {
		return state
	}
	// if mach.a.presses == 101 || mach.b.presses == 101 {
	// 	// fmt.Printf("%+v\n", mach)
	// 	// os.Exit(1)
	// 	return State{toks: math.MaxInt}
	// }
	if target.x == 0 && target.y == 0 {
		return State{toks: 0}
	}
	// fmt.Println(target)
	if target.x < 0 || target.y < 0 {
		return State{toks: math.MaxInt}
	}

	// fmt.Println(target)
	bestA := play(mach, mach.a.Press(target), cache)
	mach.a.presses--
	if bestA.toks != math.MaxInt {
		bestA.aPress++
		bestA.toks += mach.a.cost
	}

	bestB := play(mach, mach.b.Press(target), cache)
	mach.a.presses--
	if bestB.toks != math.MaxInt {
		bestB.bPress++
		bestB.toks += mach.b.cost
	}
	best := bestA
	if bestB.toks < bestA.toks {
		best = bestB
	}
	// if target.x != 0 && target.y != 0 && mach.other.x%target.x == 0 && mach.other.y%target.y == 0 {
	// 	mach.found = true
	// } else {
	// 	fmt.Println(target.x, target.y)
	// }
	cache[target] = best
	return best
}

func greater(b Button, target Loc) bool {
	return b.Loc.x > target.x && b.Loc.y > target.y
}

type Loc struct {
	x, y float64
}

type Button struct {
	Loc
	cost    int
	presses int
}

func (b *Button) Press(l Loc) Loc {
	b.presses++
	return Loc{
		x: l.x - b.x,
		y: l.y - b.y,
	}
}

type Machine struct {
	a      Button
	b      Button
	target Loc
	other  Loc
	gcd    int
}

var numReg = regexp.MustCompile(`([0-9]+)[^0-9]*([0-9]+)`)

func getInput(rows [][]byte) []Machine {
	var machs []Machine
	for i := 0; i <= len(rows); i += 4 {
		mach := Machine{}
		mach.a.x, mach.a.y = getNums(rows[i])
		mach.a.cost = 3
		mach.b.x, mach.b.y = getNums(rows[i+1])
		mach.b.cost = 1
		mach.target.x, mach.target.y = getNums(rows[i+2])
		mach.other.x, mach.other.y = mach.target.x+10000000000000, mach.target.y+10000000000000
		machs = append(machs, mach)
	}
	return machs
}

func getNums(row []byte) (float64, float64) {
	matches := numReg.FindSubmatch(row)
	x, err := strconv.ParseFloat(string(matches[1]), 64)
	if err != nil {
		panic(err)
	}
	y, err := strconv.ParseFloat(string(matches[2]), 64)
	if err != nil {
		panic(err)
	}
	return x, y
}
func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// 19560
// 52958
