package aoc

import (
	"bytes"
	"fmt"
	"os"
)

var LargeNumber = 1<<62 - 1

func ReadLines(filename string) ([][]byte, error) {
	input, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("could not read file: %w", err)
	}
	lines := bytes.Split(input, []byte("\n"))
	return lines[:len(lines)-1], nil
}

type Loc struct {
	x, y int
}

const (
	leftDir  = byte('<')
	upDir    = byte('^')
	rightDir = byte('>')
	downDir  = byte('v')
)

func (l Loc) next(dir byte) Loc {
	switch dir {
	case leftDir:
		return l.left()
	case upDir:
		return l.up()
	case rightDir:
		return l.right()
	case downDir:
		return l.down()
	}
	panic("what is even that?")
}

func (l Loc) left() Loc {
	return Loc{
		x: l.x,
		y: l.y - 1,
	}
}
func (l Loc) up() Loc {
	return Loc{
		x: l.x - 1,
		y: l.y,
	}
}
func (l Loc) right() Loc {
	return Loc{
		x: l.x,
		y: l.y + 1,
	}
}
func (l Loc) down() Loc {
	return Loc{
		x: l.x + 1,
		y: l.y,
	}
}
