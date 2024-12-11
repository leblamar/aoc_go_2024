package utils

import "strconv"

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

func (p Position) String() string {
	return "X:" + strconv.FormatInt(int64(p.X), 10) + ", Y:" + strconv.FormatInt(int64(p.Y), 10)
}

// True if p is inside the matrix
// matrix must be a Matrix which means that all rows have the same length
func (p Position) Inside(matrix [][]any) bool {
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