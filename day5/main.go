package main

import (
	"bytes"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/KevinJoiner/AdventOfCode2024/aoc"
)

func main() {
	rows, err := aoc.ReadLines("./day5/input.txt")
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
	allPages := map[string]map[string]struct{}{}
	storePages := true
	total := 0
rowLoop:
	for _, row := range rows {
		if len(row) == 0 {
			storePages = false
			continue
		}
		if storePages {
			pages := bytes.Split(row, []byte{'|'})
			parent := string(pages[0])
			child := string(pages[1])
			pageDeps := allPages[parent]
			if pageDeps == nil {
				pageDeps = map[string]struct{}{}
			}
			pageDeps[child] = struct{}{}
			allPages[parent] = pageDeps
			continue
		}
		vals := strings.Split(string(row), ",")
		seen := make([]string, 0, len(vals))

		for _, val := range vals {
			deps, ok := allPages[val]
			if !ok {
				deps = map[string]struct{}{}
			}
			for _, seenVal := range seen {
				if _, ok := deps[seenVal]; ok {
					continue rowLoop
				}
			}
			seen = append(seen, val)
		}

		mid, err := strconv.Atoi(vals[len(vals)/2])
		if err != nil {
			panic(err)
		}
		total += mid
	}
	return total, nil
}

func puzzle2(rows [][]byte) (any, error) {
	allPages := map[string]map[string]struct{}{}
	storePages := true
	total := 0
rowLoop:
	for _, row := range rows {
		if len(row) == 0 {
			storePages = false
			continue
		}
		if storePages {
			pages := bytes.Split(row, []byte{'|'})
			parent := string(pages[0])
			child := string(pages[1])
			pageDeps := allPages[parent]
			if pageDeps == nil {
				pageDeps = map[string]struct{}{}
			}
			pageDeps[child] = struct{}{}
			allPages[parent] = pageDeps
			continue
		}
		vals := strings.Split(string(row), ",")
		seen := make([]string, 0, len(vals))

		for _, val := range vals {
			deps, ok := allPages[val]
			if !ok {
				deps = map[string]struct{}{}
			}
			for _, seenVal := range seen {
				if _, ok := deps[seenVal]; ok {
					newList := sortList(allPages, vals)
					mid, err := strconv.Atoi(newList[len(newList)/2])
					if err != nil {
						panic(err)
					}
					total += mid
					continue rowLoop
				}
			}
			seen = append(seen, val)
		}
	}
	return total, nil
}

func sortList(allPages map[string]map[string]struct{}, allValues []string) []string {
	newList := []string{}
	for _, val := range allValues {
		putInList(val, allPages, &newList, allValues)
	}
	return newList
}

func putInList(page string, allPages map[string]map[string]struct{}, currentList *[]string, allValues []string) {
	for _, listItem := range *currentList {
		if listItem == page {
			return
		}
	}
	deps := allPages[page]
	for dep := range deps {
		depInList := false
		for _, val := range allValues {
			if val == dep {
				depInList = true
				break
			}
		}
		if !depInList {
			continue
		}
		found := false
		for _, listItem := range *currentList {
			if listItem == dep {
				found = true
				break
			}
		}
		if !found {
			putInList(dep, allPages, currentList, allValues)
		}
	}
	*currentList = append(*currentList, page)
}
