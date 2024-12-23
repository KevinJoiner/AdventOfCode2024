package main

import (
	"fmt"
	"log"
	"slices"
	"strings"
	"time"

	"github.com/KevinJoiner/AdventOfCode2024/aoc"
)

func main() {
	rows, err := aoc.ReadLines("./day23/input.txt")
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
	adjList := map[string][]string{}
	for _, row := range rows {
		vals := strings.Split(string(row), "-")
		first, second := vals[0], vals[1]
		adjList[first] = append(adjList[first], second)
		adjList[second] = append(adjList[second], first)
	}
	allGroups := map[string]struct{}{}
	for key := range adjList {
		groups := findCycleIn3(adjList, key)
		for _, group := range groups {
			allGroups[group] = struct{}{}
		}
	}
	return len(allGroups)
}

var cache = map[string][]string{}

func puzzle2(rows [][]byte) any {
	adjMap := map[string]map[string]bool{}
	for _, row := range rows {
		vals := strings.Split(string(row), "-")
		first, second := vals[0], vals[1]
		if _, ok := adjMap[first]; !ok {
			adjMap[first] = map[string]bool{}
		}
		if _, ok := adjMap[second]; !ok {
			adjMap[second] = map[string]bool{}
		}
		adjMap[first][second] = true
		adjMap[second][first] = true
	}

	maxSize := 0
	var maxGroup []string
	for key := range adjMap {
		localGroup := dfs2(adjMap, []string{key}, key)
		if len(localGroup) > maxSize {
			maxSize = len(localGroup)
			maxGroup = localGroup
		}
	}
	slices.Sort(maxGroup)
	return strings.Join(maxGroup, ",")
}

func findCycleIn3(adjList map[string][]string, start string) []string {
	groups := dfs(adjList, []string{start})
	retGroups := []string{}
	for _, group := range groups {
		hasT := false
		for _, pc := range group {
			if strings.HasPrefix(pc, "t") {
				hasT = true
				break
			}
		}
		if hasT {
			slices.Sort(group)
			retGroups = append(retGroups, strings.Join(group, ","))
		}
	}
	return retGroups
}

func dfs(adjList map[string][]string, group []string) [][]string {
	curr := group[len(group)-1]
	if len(group) == 3 {
		for _, child := range adjList[curr] {
			if child == group[0] {
				return [][]string{group}
			}
		}
		return nil
	}
	allGroups := [][]string{}
	for _, child := range adjList[curr] {
		if inGroup(child, group) {
			continue
		}
		gCopy := make([]string, len(group), len(group)+1)
		copy(gCopy, group)
		gCopy = append(gCopy, child)
		groups := dfs(adjList, gCopy)
		for _, group := range groups {
			if group != nil {
				allGroups = append(allGroups, group)
			}
		}
	}
	return allGroups
}
func inGroup(val string, group []string) bool {
	for _, pc := range group {
		if val == pc {
			return true
		}
	}
	return false
}

func dfs2(adjMap map[string]map[string]bool, group []string, curr string) []string {
	slices.Sort(group)
	maxGroup, ok := cache[strings.Join(group, ",")]
	if ok {
		return maxGroup
	}
	maxSize := len(group)
	maxGroup = group
childLoop:
	for child := range adjMap[curr] {
		for _, pc := range group {
			if child == pc {
				continue childLoop
			}
			if !adjMap[child][pc] {
				continue childLoop
			}
		}
		gCopy := make([]string, len(group), len(group)+1)
		copy(gCopy, group)
		gCopy = append(gCopy, child)
		localGroup := dfs2(adjMap, gCopy, child)
		if len(localGroup) > maxSize {
			maxSize = len(localGroup)
			maxGroup = localGroup
		}
	}
	cache[strings.Join(group, ",")] = maxGroup
	return maxGroup
}
