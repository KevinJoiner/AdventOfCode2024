package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/KevinJoiner/AdventOfCode2024/aoc"
)

func main() {
	rows, err := aoc.ReadLines("./day17/input.txt")
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
	machine, code := getMachine(rows)
	fmt.Printf("%+v, %v\n", machine, code)
	machine.run(code)
	strs := make([]string, len(machine.outBuf))
	for i, n := range machine.outBuf {
		strs[i] = strconv.Itoa(int(n))
	}
	return strings.Join(strs, ",")
}
func puzzle2(rows [][]byte) any {
	_, code := getMachine(rows)
	minValid := atomic.Int64{}
	minValid.Store(math.MaxInt) // 216584205979245
	wg := sync.WaitGroup{}
	for runner := range 12 {
		wg.Add(1)
		time.Sleep(time.Nanosecond)
		go func(runner int) {
			defer wg.Done()
			for i := range 100 {
				valid := evolution(code, 0)
				if valid < int(minValid.Load()) {
					minValid.Store(int64(valid))
					fmt.Println(valid, i)
				}
				if i%100 == 0 {
					fmt.Println("runner:", runner, "Tried:", i, "iterations")
				}
			}
		}(runner)
	}
	wg.Wait()
	return minValid.Load()
}
func score(code []uint8, val int) (int, []uint8) {
	n := len(code)
	i := n - 1
	score := 0
	mach := i25SantaMachine{regA: val}
	mach.run(code)
	if len(mach.outBuf) != n {
		return -1, mach.outBuf
	}
	for i > -1 && code[i] == mach.outBuf[i] {
		score += i
		i--
	}
	return score, mach.outBuf
}

// https://en.wikipedia.org/wiki/Evolutionary_algorithm
func evolution(code []uint8, seed int) int {
	max, out := score(code, seed)
	cur := seed
	var tmpScore int
	bitSize := math.MaxInt
	tries := 0
	gen := rand.New(rand.NewSource(time.Now().UnixNano()))
	for {
		bit := gen.Int() % bitSize
		tmp := cur ^ bit
		tries++
		tmpScore, out = score(code, tmp)
		if slices.Compare(code, out) == 0 {
			return tmp
		}
		if tmpScore > max {
			max = tmpScore
			cur = tmp
			tries = 0
		} else if tmpScore < max {
			bitSize = bit
		}
		if tries == 50 {
			// Pop out of any local minimums
			bitSize = math.MaxInt
			tries = 0
		}
	}
}

var numReg = regexp.MustCompile(`[0-9]+`)

func getMachine(rows [][]byte) (i25SantaMachine, []uint8) {
	mach := i25SantaMachine{}
	var err error
	mach.regA, err = strconv.Atoi(numReg.FindString(string(rows[0])))
	if err != nil {
		panic(err)
	}
	mach.regB, err = strconv.Atoi(numReg.FindString(string(rows[1])))
	if err != nil {
		panic(err)
	}
	mach.regC, err = strconv.Atoi(numReg.FindString(string(rows[2])))
	if err != nil {
		panic(err)
	}
	codeStrs := strings.Split(strings.Split(string(rows[4]), ":")[1], ",")
	code := make([]uint8, 0, len(codeStrs))
	for _, str := range codeStrs {
		num, err := strconv.Atoi(strings.TrimSpace(str))
		if err != nil {
			panic(err)
		}
		code = append(code, uint8(num))
	}
	return mach, code
}

type i25SantaMachine struct {
	regA   int
	regB   int
	regC   int
	ptr    int
	outBuf []uint8
}

func (m *i25SantaMachine) reset(code []uint8) {
	m.regA = 0
	m.regB = 0
	m.regC = 0
	m.outBuf = nil
	m.ptr = 0
}

func (m *i25SantaMachine) run(code []uint8) {
	for m.ptr < len(code) {
		m.execute(code[m.ptr], code[m.ptr+1])
	}
}

func (m *i25SantaMachine) execute(opcode, operand uint8) {
	switch opcode {
	case 0:
		m.adv(operand)
	case 1:
		m.bxl(operand)
	case 2:
		m.bst(operand)
	case 3:
		m.jnz(operand)
	case 4:
		m.bxc(operand)
	case 5:
		m.out(operand)
	case 6:
		m.bdv(operand)
	case 7:
		m.cdv(operand)
	}
}

func (m *i25SantaMachine) combo(operand uint8) int {
	switch operand {
	case 0, 1, 2, 3:
		return int(operand)
	case 4:
		return m.regA
	case 5:
		return m.regB
	case 6:
		return m.regC
	default:
		panic("What are you doing?")
	}
}
func (m *i25SantaMachine) next() {
	m.ptr += 2
}

// adv instruction (opcode 0) performs division. The numerator is the value in the A register. The denominator is found by raising 2 to the power of the instruction's combo operand. (So, an operand of 2 would divide A by 4 (2^2); an operand of 5 would divide A by 2^B.) The result of the division operation is truncated to an integer and then written to the A register.
func (m *i25SantaMachine) adv(operand uint8) {
	m.regA = m.regA / (1 << m.combo(operand))
	m.next()
}

// bxl instruction (opcode 1) calculates the bitwise XOR of register B and the instruction's literal operand, then stores the result in register B.
func (m *i25SantaMachine) bxl(operand uint8) {
	m.regB = m.regB ^ int(operand)
	m.next()
}

// bst instruction (opcode 2) calculates the value of its combo operand modulo 8 (thereby keeping only its lowest 3 bits), then writes that value to the B register.
func (m *i25SantaMachine) bst(operand uint8) {
	m.regB = m.combo(operand) % 8
	m.next()
}

// jnz instruction (opcode 3) does nothing if the A register is 0. However, if the A register is not zero, it jumps by setting the instruction pointer to the value of its literal operand; if this instruction jumps, the instruction pointer is not increased by 2 after this instruction.
func (m *i25SantaMachine) jnz(operand uint8) {
	if m.regA == 0 {
		m.next()
		return
	}
	m.ptr = int(operand)
}

// bxc instruction (opcode 4) calculates the bitwise XOR of register B and register C, then stores the result in register B. (For legacy reasons, this instruction reads an operand but ignores it.)
func (m *i25SantaMachine) bxc(_ uint8) {
	m.regB = m.regB ^ m.regC
	m.next()
}

// out instruction (opcode 5) calculates the value of its combo operand modulo 8, then outputs that value. (If a program outputs multiple values, they are separated by commas.)
func (m *i25SantaMachine) out(operand uint8) {
	m.outBuf = append(m.outBuf, uint8(m.combo(operand)%8))
	m.next()
}

// bdv instruction (opcode 6) works exactly like the adv instruction except that the result is stored in the B register. (The numerator is still read from the A register.)
func (m *i25SantaMachine) bdv(operand uint8) {
	m.regB = m.regA / (1 << m.combo(operand))
	m.next()
}

// cdv instruction (opcode 7) works exactly like the adv instruction except that the result is stored in the C register. (The numerator is still read from the A register.)
func (m *i25SantaMachine) cdv(operand uint8) {
	m.regC = m.regA / (1 << m.combo(operand))
	m.next()
}

// func bSearch(code []uint8, cmpFunc func(i int, code []uint8) int) (int, bool) {
// 	// Define cmp(x[-1], target) < 0 and cmp(x[n], target) >= 0 .
// 	// Invariant: cmp(x[i - 1], target) < 0, cmp(x[j], target) >= 0.
// 	i, j := 35184372088832, 281474976710656
// 	for i < j {
// 		h := int(uint(i+j) >> 1) // avoid overflow when computing h
// 		// i ≤ h < j

// 		if cmpFunc(h, code) < 0 {
// 			i = h + 1 // preserves cmp(x[i - 1], target) < 0
// 		} else {
// 			j = h // preserves cmp(x[j], target) >= 0
// 		}
// 	}
// 	// i == j, cmp(x[i-1], target) < 0, and cmp(x[j], target) (= cmp(x[i], target)) >= 0  =>  answer is i.
// 	return i, i < math.MaxInt32 && cmpFunc(i, code) == 0
// }
// func testLower(i int, code []uint8) int {
// 	machine := i25SantaMachine{
// 		regA: i,
// 	}
// 	machine.run(code)
// 	res := cmp.Compare(len(machine.outBuf), len(code))
// 	if res != 0 {
// 		return res
// 	}
// 	machine = i25SantaMachine{
// 		regA: i - 1,
// 	}
// 	machine.run(code)
// 	if len(code) > len(machine.outBuf) {
// 		return 0
// 	}
// 	if len(code) == len(machine.outBuf) {
// 		return 1
// 	}
// 	return 0
// }
// func testUpper(i int, code []uint8) int {
// 	machine := i25SantaMachine{
// 		regA: i,
// 	}
// 	machine.run(code)
// 	res := cmp.Compare(len(machine.outBuf), len(code))
// 	if res != 0 {
// 		return res
// 	}
// 	machine = i25SantaMachine{
// 		regA: i + 1,
// 	}
// 	machine.run(code)
// 	if len(code) < len(machine.outBuf) {
// 		return 0
// 	}
// 	if len(code) == len(machine.outBuf) {
// 		return -1
// 	}
// 	return 0
// }

// func testAll(i int, code []uint8) int {
// 	machine := i25SantaMachine{
// 		regA: i,
// 	}
// 	machine.run(code)
// 	res := cmp.Compare(len(machine.outBuf), len(code))
// 	if res != 0 {
// 		return res
// 	}
// 	return 1 * cmp.Compare(code[len(code)-5], machine.outBuf[len(code)-5])
// }
