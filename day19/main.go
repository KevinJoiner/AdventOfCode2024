package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/KevinJoiner/AdventOfCode2024/aoc"
)

func main() {
	rows, err := aoc.ReadLines("./day19/input.txt")
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
	patterns, designs := getInput(rows)
	trie := Trie{}
	for _, word := range patterns {
		trie.Add(word)
	}
	total := 0
	cache := map[string]int{}
	for _, word := range designs {
		wordTotal := getAllDesigns(word, &trie, cache)
		if wordTotal != 0 {
			total++
		}
	}
	return total
}

func puzzle2(rows [][]byte) any {
	patterns, designs := getInput(rows)
	trie := Trie{}
	for _, word := range patterns {
		trie.Add(word)
	}
	total := 0
	cache := map[string]int{}
	for _, word := range designs {
		wordTotal := getAllDesigns(word, &trie, cache)
		total += wordTotal
	}
	return total
}

func getInput(rows [][]byte) ([]string, []string) {
	var patterns, designs []string
	for i, row := range rows {
		if i == 0 {
			patterns = strings.Split(string(row), ",")
			for i := range patterns {
				patterns[i] = strings.TrimSpace(patterns[i])
			}
			continue
		}
		if i == 1 {
			continue
		}
		designs = append(designs, string(row))
	}
	return patterns, designs
}

func getAllDesigns(design string, trie *Trie, cache map[string]int) int {
	val, ok := cache[design]
	if ok {
		return val
	}
	for i := range design {
		if trie.Search(design[:i+1]) {
			if i == len(design)-1 {
				// if the entire word mataches then we are done
				val++
				continue
			}
			val += getAllDesigns(design[i+1:], trie, cache)
		}
	}

	cache[design] = val
	return val
}

type Node struct {
	c        rune
	isWord   bool
	children [26]*Node
}
type Trie struct {
	root *Node
}

func (t *Trie) Add(word string) {
	if word == "" {
		return
	}
	if t.root == nil {
		t.root = &Node{}
	}
	curr := t.root
	for _, r := range word {
		child := curr.children[r-'a']
		if child == nil {
			child = &Node{
				c: r,
			}
			curr.children[r-'a'] = child
		}
		curr = child
	}
	curr.isWord = true
}

func (t *Trie) Search(word string) bool {
	curr := t.root
	for _, r := range word {
		child := curr.children[r-'a']
		if child == nil {
			return false
		}
		curr = child
	}
	return curr.isWord
}
