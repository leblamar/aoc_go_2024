package utils

import (
	"log"
)

type Row[T any] []T
type Grid[T any] []Row[T]

func (g Grid[T]) IsInside(p Position) bool {
	if len(g) == 0 {
		return false
	} else if len(g[0]) == 0 {
		return false
	}

	if p.X < 0 || p.Y < 0 {
		return false
	}

	return p.X < len(g) && p.Y < len(g[0])
}

func (g Grid[T]) Get(p Position) (val T, ok bool) {
	if !g.IsInside(p) {
		var defaultValue T
		return defaultValue, false
	} else {
		return g[p.X][p.Y], true
	}
}

func (g Grid[T]) Height() int {
	return len(g)
}

func (g Grid[T]) Width() int {
	if g.Height() == 0 {
		return 0
	}
	return len(g[0])
}

func Parse[T any](lines []string, subParse ParseRune[T]) (matrix Grid[T]) {
	matrix = make(Grid[T], 0, len(lines))
	for _, line := range lines {
		row := make(Row[T], 0, len(line))

		for _, elem := range line {
			val, err := subParse(elem)
			if err != nil {
				log.Fatal(err)
				continue
			}

			row = append(row, val)
		}

		matrix = append(matrix, row)
	}

	return
}
