package main

import (
	"fmt"
	"log"

	"github.com/KevinJoiner/AdventOfCode2024/aoc"
)

func main() {
	rows, err := aoc.ReadLines("./day4/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	output, err := puzzle1(rows)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Puzzle 1 output:", output)
	output, err = puzzle2(rows)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Puzzle 2 output:", output)
}

func puzzle1(rows [][]byte) (any, error) {
	count := 0
	seq := []byte{'A', 'S'}
	for i := range rows {
		for j := range rows[i] {
			if rows[i][j] != 'X' {
				continue
			}
			if left(rows, i, j, seq) {
				count++
			}
			if upLeft(rows, i, j, seq) {
				count++
			}
			if up(rows, i, j, seq) {
				count++
			}
			if upRight(rows, i, j, seq) {
				count++
			}
			if right(rows, i, j, seq) {
				count++
			}
			if downRight(rows, i, j, seq) {
				count++
			}
			if down(rows, i, j, seq) {
				count++
			}
			if downLeft(rows, i, j, seq) {
				count++
			}
		}
	}
	return count, nil
}

func puzzle2(rows [][]byte) (any, error) {
	count := 0
	m := []byte{'M'}
	s := []byte{'S'}
	for i := range rows {
		for j := range rows[i] {
			if rows[i][j] != 'A' {
				continue
			}
			if !(upLeft(rows, i, j, m) && downRight(rows, i, j, s) || upLeft(rows, i, j, s) && downRight(rows, i, j, m)) {
				continue
			}
			if !(downLeft(rows, i, j, m) && upRight(rows, i, j, s) || downLeft(rows, i, j, s) && upRight(rows, i, j, m)) {
				continue
			}
			count++
		}
	}
	return count, nil
}

func left(row [][]byte, i, j int, seq []byte) bool {
	if len(seq) == 0 {
		return true
	}
	j--
	return inBounds(row, i, j) && row[i][j] == seq[0] && left(row, i, j, seq[1:])
}
func upLeft(row [][]byte, i, j int, seq []byte) bool {
	if len(seq) == 0 {
		return true
	}
	i--
	j--
	return inBounds(row, i, j) && row[i][j] == seq[0] && upLeft(row, i, j, seq[1:])
}
func up(row [][]byte, i, j int, seq []byte) bool {
	if len(seq) == 0 {
		return true
	}
	i--
	return inBounds(row, i, j) && row[i][j] == seq[0] && up(row, i, j, seq[1:])
}
func upRight(row [][]byte, i, j int, seq []byte) bool {
	if len(seq) == 0 {
		return true
	}
	i--
	j++
	return inBounds(row, i, j) && row[i][j] == seq[0] && upRight(row, i, j, seq[1:])
}
func right(row [][]byte, i, j int, seq []byte) bool {
	if len(seq) == 0 {
		return true
	}
	j++
	return inBounds(row, i, j) && row[i][j] == seq[0] && right(row, i, j, seq[1:])
}
func downRight(row [][]byte, i, j int, seq []byte) bool {
	if len(seq) == 0 {
		return true
	}
	i++
	j++
	return inBounds(row, i, j) && row[i][j] == seq[0] && downRight(row, i, j, seq[1:])
}
func down(row [][]byte, i, j int, seq []byte) bool {
	if len(seq) == 0 {
		return true
	}
	i++
	return inBounds(row, i, j) && row[i][j] == seq[0] && down(row, i, j, seq[1:])
}
func downLeft(row [][]byte, i, j int, seq []byte) bool {
	if len(seq) == 0 {
		return true
	}
	i++
	j--
	return inBounds(row, i, j) && row[i][j] == seq[0] && downLeft(row, i, j, seq[1:])
}
func inBounds(rows [][]byte, i, j int) bool {
	return i != -1 && i != len(rows) && j != -1 && j != len(rows[0])
}
