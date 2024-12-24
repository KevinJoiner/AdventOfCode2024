package main

import (
	"bytes"
	"cmp"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/KevinJoiner/AdventOfCode2024/aoc"
)

func main() {
	rows, err := aoc.ReadLines("./day24/input.txt")
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
	state, ops := getInput(rows)
	return run(state, ops)
}
func run(ogState map[string]int, ogOps []Op) int {
	state := make(map[string]int, len(ogState))
	for k, v := range ogState {
		state[k] = v
	}
	ops := make([]Op, len(ogOps))
	copy(ops, ogOps)

	i := 0
	count := 0
	for len(ops) != 0 {
		if i == len(ops) {
			if count != 0 {
				// loop detected
				return -1
			}
			count++
			i = 0
		}
		op := ops[i]
		input1Val, ok := state[op.input1]
		if !ok {
			i++
			continue
		}
		input2Val, ok := state[op.input2]
		if !ok {
			i++
			continue
		}
		count = 0
		state[op.output] = operate(input1Val, input2Val, op.gate)
		ops = slices.Delete(ops, i, i+1)
	}
	z := 0
	for k, v := range state {
		if k[0] == 'z' {
			num, err := strconv.Atoi(k[1:])
			if err != nil {
				panic(err)
			}
			z += v << num
		}
	}
	return z
}

func run2(ogState map[string]int, ogOps []Op) int {
	state := make(map[string]int, len(ogState))
	for k, v := range ogState {
		state[k] = v
	}
	ops := make([]Op, len(ogOps))
	copy(ops, ogOps)
	i := 0
	for len(ops) != 0 {
		if i == len(ops) {
			return -1
		}
		op := ops[i]
		input1Val, ok := state[op.input1]
		if !ok {
			i++
			continue
		}
		input2Val, ok := state[op.input2]
		if !ok {
			i++
			continue
		}
		if op.gate != "XOR" && strings.Contains(op.output, "z") {
			// Should not happens in full adder
			fmt.Printf("%s %s %s %s\n", op.input1, op.gate, op.input2, op.output)
		}
		if strings.Contains(op.output, "_z") {
			// should not happen in full adder
			fmt.Printf("%s %s %s %s\n", op.input1, op.gate, op.input2, op.output)
		}
		// fmt.Printf("%s %s %s %s\n", op.input1, op.gate, op.input2, op.output)
		state[op.output] = operate(input1Val, input2Val, op.gate)
		ops = slices.Delete(ops, i, i+1)
		i = 0
	}
	z := 0
	for k, v := range state {
		if k[0] == 'z' {
			num, err := strconv.Atoi(k[1:])
			if err != nil {
				panic(err)
			}
			z += v << num
		}
	}
	return z
}

// score checks how many of the bits in the target are the same as the guess
func score(target, guess int) int {
	count := 0
	for i := 0; i < 47; i++ {
		if target&(1<<i) == guess&(1<<i) {
			count++
		}
	}
	return count
}
func puzzle2(rows [][]byte) any {
	// This Prints out bad outputs in addr
	state, ops := getInput2(rows)
	slices.SortFunc(ops, func(i, j Op) int {
		val1, err := strconv.Atoi(i.input1[1:])
		if err != nil {
			return cmp.Compare(i.output, j.output)
		}
		val2, err := strconv.Atoi(j.input1[1:])
		if err != nil {
			return cmp.Compare(i.output, j.output)
		}
		return cmp.Compare(val1, val2)
	})
	run2(state, ops)

	// This checks swaps after manually reviewing the bad outputs
	state, ops = getInput(rows)
	// return "not found"
	for i, op := range ops {
		if op.output == "cdj" {
			op.output = "z08"
		} else if op.output == "z08" {
			op.output = "cdj"
		} else if op.output == "z32" {
			op.output = "gfm"
		} else if op.output == "gfm" {
			op.output = "z32"
		} else if op.output == "z16" {
			op.output = "mrb"
		} else if op.output == "mrb" {
			op.output = "z16"
		} else if op.output == "qjd" {
			op.output = "dhm"
		} else if op.output == "dhm" {
			op.output = "qjd"
		}
		ops[i] = op
	}
	target := getTarget(state)
	fmt.Println(run(state, ops), target)
	ret := []string{"cdj", "z08", "z32", "gfm", "mrb", "z16", "qjd", "dhm"}
	slices.Sort(ret)
	return strings.Join(ret, ",")

}

func getTarget(state map[string]int) int {
	x, y := 0, 0
	for k, v := range state {
		if k[0] == 'x' {
			num, err := strconv.Atoi(k[1:])
			if err != nil {
				panic(err)
			}
			x += v << num
		} else if k[0] == 'y' {
			num, err := strconv.Atoi(k[1:])
			if err != nil {
				panic(err)
			}
			y += v << num
		}
	}
	return x + y
}
func operate(in1, in2 int, op string) int {
	switch op {
	case "AND":
		return in1 & in2
	case "OR":
		return in1 | in2
	case "XOR":
		return in1 ^ in2
	}
	return -1

}

type Op struct {
	input1 string
	input2 string
	output string
	gate   string
}

func getInput(rows [][]byte) (map[string]int, []Op) {
	state := map[string]int{}
	var ops []Op
	setState := true
	for _, row := range rows {
		if len(row) == 0 {
			setState = false
			continue
		}
		if setState {
			parts := strings.Split(string(row), ":")
			val, err := strconv.Atoi(strings.TrimSpace(parts[1]))
			if err != nil {
				panic(err)
			}
			state[parts[0]] = val
			continue
		}
		parts := strings.Split(string(row), " ")
		ops = append(ops, Op{
			input1: parts[0],
			gate:   parts[1],
			input2: parts[2],
			output: parts[4],
		})
	}
	return state, ops
}

func getInput2(rows [][]byte) (map[string]int, []Op) {
	state := map[string]int{}
	reNames := map[string]string{}
	var ops []Op
	setState := true
	for _, row := range rows {
		if len(row) == 0 {
			setState = false
			continue
		}
		if setState {
			continue
		}
		parts := strings.Split(string(row), " ")
		op := Op{
			input1: parts[0],
			gate:   parts[1],
			input2: parts[2],
			output: parts[4],
		}
		if op.input1[0] == 'x' || op.input1[0] == 'y' {
			reNames[op.output] = op.gate + "_" + op.input1[1:] + "_" + op.output
		}
	}
	for ogName, newName := range reNames {
		for row := range rows {
			rows[row] = bytes.ReplaceAll(rows[row], []byte(ogName), []byte(newName))
		}
	}
	reNames = map[string]string{}
	setState = true
	for _, row := range rows {
		if len(row) == 0 {
			setState = false
			continue
		}
		if setState {
			continue
		}
		parts := strings.Split(string(row), " ")
		op := Op{
			input1: parts[0],
			gate:   parts[1],
			input2: parts[2],
			output: parts[4],
		}
		if op.input1[:3] == "AND" {
			reNames[op.output] = "COUT_" + op.input1[4:6] + "_" + op.output
		}
		if op.input2[:3] == "AND" {
			reNames[op.output] = "COUT_" + op.input2[4:6] + "_" + op.output
		}
	}
	for ogName, newName := range reNames {
		for row := range rows {
			rows[row] = bytes.ReplaceAll(rows[row], []byte(ogName), []byte(newName))
		}
	}
	setState = true
	for _, row := range rows {
		if len(row) == 0 {
			setState = false
			continue
		}
		if setState {
			parts := strings.Split(string(row), ":")
			val, err := strconv.Atoi(strings.TrimSpace(parts[1]))
			if err != nil {
				panic(err)
			}
			state[parts[0]] = val
			continue
		}
		parts := strings.Split(string(row), " ")
		in1 := parts[0]
		in2 := parts[2]
		if in1 < in2 {
			in1, in2 = in2, in1
		}
		op := Op{
			input1: in1,
			gate:   parts[1],
			input2: in2,
			output: parts[4],
		}
		ops = append(ops, op)
	}
	return state, ops
}
