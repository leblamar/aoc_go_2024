package utils

import (
	"errors"
	"strconv"
)

type Position struct {
	X int
	Y int
}

// Find symetric of position
func (p1 Position) Sym(p2 Position) Position {
	xDist := p2.X - p1.X
	yDist := p2.Y - p1.Y
	newX := p2.X + xDist
	newY := p2.Y + yDist
	return Position{newX, newY}
}

func (p1 Position) Add(p2 Position) Position {
	newX := p1.X + p2.X
	newY := p1.Y + p2.Y
	return Position{newX, newY}
}

func (p Position) String() string {
	return "X:" + strconv.FormatInt(int64(p.X), 10) + ", Y:" + strconv.FormatInt(int64(p.Y), 10)
}

// True if p is inside the matrix
// matrix must be a Matrix which means that all rows have the same length
func (p Position) Inside(matrix [][]int) bool {
	if len(matrix) == 0 {
		return false
	} else if len(matrix[0]) == 0 {
		return false
	}

	if p.X < 0 || p.Y < 0 {
		return false
	}

	return p.X < len(matrix) && p.Y < len(matrix[0])
}

var cardinalToPosMap = map[rune]Position{'<': {-1, 0}, '^': {0, -1}, '>': {1, 0}, 'v': {0, 1}}
var dirToStringMap = map[Position]rune{{-1, 0}: '<', {0, -1}: '^', {1, 0}: '>', {0, 1}: 'v'}

func (d Position) ToRune() (rune, error) {
	dir, ok := dirToStringMap[d]
	if !ok {
		return dir, errors.New("you passed an unrecognized direction")
	}
	return dir, nil
}

func ParseDir(r rune) (Position, error) {
	dir, ok := cardinalToPosMap[r]
	if !ok {
		return dir, errors.New("you passed an unrecognized direction")
	}
	return dir, nil
}

var CardinalDirs = []Position{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
var WithDiagDirs = []Position{{-1, -1}, {1, 1}, {1, -1}, {-1, 1}, {0, -1}, {1, 0}, {0, 1}, {-1, 0}}
