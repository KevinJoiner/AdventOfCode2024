package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"time"

	"github.com/KevinJoiner/AdventOfCode2024/aoc"
)

func main() {
	rows, err := aoc.ReadLines("./day14/input.txt")
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
	maxX, maxY := 103, 101
	robots := getInput(rows)
	var spots []Loc
	quad1, quad2, quad3, quad4 := 0, 0, 0, 0

	for _, bot := range robots {
		spot := Loc{}
		moveX := (bot.p.x + bot.v.x*100)
		moveY := (bot.p.y + bot.v.y*100)
		spot.x = abs(moveX) % maxX
		spot.y = abs(moveY) % maxY
		if moveX < 0 && spot.x != 0 {
			spot.x = maxX - spot.x
		}
		if moveY < 0 && spot.y != 0 {
			spot.y = maxY - spot.y
		}
		if spot.x < maxX/2 {
			if spot.y < maxY/2 {
				quad1++
			} else if spot.y > maxY/2 {
				quad2++
			}
		} else if spot.x > maxX/2 {
			if spot.y < maxY/2 {
				quad3++
			} else if spot.y > maxY/2 {
				quad4++
			}
		}
		spots = append(spots, spot)
	}
	return quad1 * quad2 * quad3 * quad4
}

func puzzle2(rows [][]byte) any {
	maxX, maxY := 103, 101
	size := Loc{x: maxX, y: maxY}
	robots := getInput(rows)
	// }
	start := time.Now()
treeLoop:
	for i := range 1_000_000_000 {
		if i%1_000_0 == 0 {
			fmt.Println(i, time.Since(start))
		}
		spots := make(map[Loc]int, len(robots))
		treeSpots := 0
		for _, bot := range robots {
			spot := Loc{}
			moveX := (bot.p.x + bot.v.x*i)
			moveY := (bot.p.y + bot.v.y*i)
			spot.x = abs(moveX) % maxX
			spot.y = abs(moveY) % maxY
			if moveX < 0 && spot.x != 0 {
				spot.x = maxX - spot.x
			}
			if moveY < 0 && spot.y != 0 {
				spot.y = maxY - spot.y
			}
			if inTree(spot) {
				treeSpots++
			}
			spots[spot]++
		}
		percentInTree := float64(treeSpots) / float64(len(robots))
		if percentInTree < .8 {
			continue
		}
		fmt.Printf("%f%% %d\n", percentInTree, i)
		fmt.Printf("%f%% %d\n", percentInTree, i)
		fmt.Printf("%f%% %d\n", percentInTree, i)
		fmt.Printf("%f%% %d\n", percentInTree, i)
		fmt.Printf("%f%% %d\n", percentInTree, i)
		fmt.Printf("%f%% %d\n", percentInTree, i)
		fmt.Printf("%f%% %d\n", percentInTree, i)
		printgrid(spots, size)
		println("################################")
		println("################################")
		time.Sleep(500 * time.Millisecond)
		continue treeLoop
	}
	return "fun"
}
func printgrid(spots map[Loc]int, size Loc) {
	for i := range size.x {
		for j := range size.y {
			l := Loc{i, j}
			if num, ok := spots[l]; ok {
				fmt.Print(num)
			} else if inTree(l) {
				fmt.Print("_")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

type Loc struct {
	x, y int
}

type Robot struct {
	p Loc
	v Loc
}

func getInput(rows [][]byte) []Robot {
	var bots []Robot
	for i := 0; i < len(rows); i++ {
		bot := Robot{}
		bot.p.y, bot.p.x, bot.v.y, bot.v.x = getNums(rows[i])
		// fmt.Println(bot)
		bots = append(bots, bot)
	}
	return bots
}

var numReg = regexp.MustCompile(`(-*[0-9]+)[^0-9\-]*(-*[0-9]+)[^0-9\-]*(-*[0-9]+)[^0-9\-]*(-*[0-9]+)`)

func getNums(row []byte) (int, int, int, int) {
	matches := numReg.FindSubmatch(row)
	pX, err := strconv.Atoi(string(matches[1]))
	if err != nil {
		panic(err)
	}
	pY, err := strconv.Atoi(string(matches[2]))
	if err != nil {
		panic(err)
	}
	vX, err := strconv.Atoi(string(matches[3]))
	if err != nil {
		panic(err)
	}
	vY, err := strconv.Atoi(string(matches[4]))
	if err != nil {
		panic(err)
	}
	return pX, pY, vX, vY
}

func inTree(spot Loc) bool {
	c := 51
	leftEdge := c - (spot.x / 2)
	rightEdge := c + (spot.x / 2)
	return spot.y >= leftEdge && spot.y <= rightEdge
}
