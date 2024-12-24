package utils

type ParseRune[T any] func(Position, rune) (T, error)
