package utils

type Grid[T any] [][]T

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
