package dayUtils

import (
	"aoc_go_2024/src/utils"
	"fmt"
	"time"
)

type IDay[T any] interface {
	GetNumber() uint
	Parse([]string) T
	Part1(bool, T) int64
	Part2(bool, T) int64
}

func Run[T any](d IDay[T], isTest, debug bool) {
	fmt.Println("Welcome to day", d.GetNumber(), "!!!")

	lines := utils.GetLines(isTest, d.GetNumber())
	input := d.Parse(lines)

	start := time.Now()
	sol1 := d.Part1(debug, input)
	ellapsed := time.Since(start)
	fmt.Println("Part 1:", sol1, "in", ellapsed)

	start = time.Now()
	sol2 := d.Part2(debug, input)
	ellapsed = time.Since(start)
	fmt.Println("Part 2:", sol2, "in", ellapsed)
}
