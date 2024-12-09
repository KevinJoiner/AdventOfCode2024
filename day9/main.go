package main

import (
	"fmt"
	"log"
	"slices"
	"strconv"
	"time"

	"github.com/KevinJoiner/AdventOfCode2024/aoc"
)

func main() {
	rows, err := aoc.ReadLines("./day9/input.txt")
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
	items := getItems(rows)
	totalCheckSum := 0
	fileIdx := 0
	for i := 0; i < len(items); i++ {
		item := items[i]
		if item.id != -1 {
			var itemCheck int
			fileIdx, itemCheck = getCheckSum(item, fileIdx)
			totalCheckSum += itemCheck
			continue
		}
		newItems := pullFromEnd(&items, i, item.size)
		// Delete The freeSpace
		items = slices.Delete(items, i, i+1)
		items = slices.Insert(items, i, newItems...)
		// redo this idx
		i--
		continue
	}
	return totalCheckSum
}

func puzzle2(rows [][]byte) any {
	items := getItems(rows)
	totalCheckSum := 0

	var found bool
	for i := len(items) - 1; i >= 0; i-- {
		item := items[i]
		if item.id == -1 {
			continue
		}
		items, found = putInFront(items, item, i)
		if found {
			item.id = -1
		}
	}
	fileIdx := 0
	for _, item := range items {
		if item.id == -1 {
			fileIdx += item.size
			continue
		}
		for range item.size {
			totalCheckSum += item.id * fileIdx
			fileIdx++
		}
	}
	// printItems(items)
	return totalCheckSum
}

func printItems(items []*Item) {
	for _, item := range items {
		for range item.size {
			if item.id != -1 {
				fmt.Printf("%d", item.id)
			} else {
				fmt.Printf(".")
			}
		}
	}
	fmt.Println()
}

type Item struct {
	id   int
	size int
}

func getItems(rows [][]byte) []*Item {
	if len(rows) != 1 {
		panic("too manny rows")
	}
	id := 0
	items := make([]*Item, 0, len(rows[0]))
	for i, b := range rows[0] {
		size, err := strconv.Atoi(string(b))
		if err != nil {
			panic(err)
		}
		item := &Item{size: size}
		if i&1 == 0 {
			item.id = id
			id++
		} else {
			item.id = -1
		}
		items = append(items, item)
	}
	return items
}

func getCheckSum(item *Item, fileIdx int) (int, int) {
	total := 0
	for range item.size {
		total += item.id * fileIdx
		fileIdx++
	}
	return fileIdx, total
}

func pullFromEnd(items *[]*Item, currentIdx int, freeSpace int) []*Item {
	retItems := []*Item{}
	for i := len(*items) - 1; i >= 0; i-- {
		if freeSpace == 0 || i == currentIdx {
			return retItems
		}
		item := (*items)[i]
		if item.id == -1 {
			*items = slices.Delete(*items, i, i+1)
			continue
		}
		if item.size > freeSpace {
			item.size -= freeSpace
			retItems = append(retItems, &Item{
				id:   item.id,
				size: freeSpace,
			})
			return retItems
		}
		retItems = append(retItems, &Item{
			id:   item.id,
			size: item.size,
		})
		freeSpace -= item.size
		// remove this empty item
		*items = slices.Delete(*items, i, i+1)
	}
	return retItems
}

func putInFront(items []*Item, item *Item, stopIdx int) ([]*Item, bool) {
	for i := 0; i < stopIdx; i++ {
		freeItem := (items)[i]
		if freeItem.id != -1 {
			continue
		}
		if freeItem.size < item.size {
			continue
		}

		if item.size == freeItem.size {
			freeItem.id = item.id
		} else {
			items = slices.Insert(items, i, &Item{
				id:   item.id,
				size: item.size,
			})
			freeItem.size -= item.size
		}
		return items, true
	}
	return items, false
}

// for item
// if item id >0  not free space
// get checkSum
// if item if < 0 free space
// call prune(freeSpace, location) []items
// continue until we reach the end of the list

// prune
// reverse traverse item list unilt freeSpace == 0 or we reach the currentLoc
// if free space just delete
// if file create a new  file item newItem.id = file.id
// if fresSpace < file.size; file.size = file.size -freeSpace; else delete file item.size = file.size freeSpace -= file.size
