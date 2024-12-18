package utils

type ParseRune[T any] func(rune) (T, error)
