package utils

type number interface {
	int | int8 | int16 | int32 | int64
}

func Abs[T number](x T) T {
	if x < 0 {
		return -x
	}
	return x
}
