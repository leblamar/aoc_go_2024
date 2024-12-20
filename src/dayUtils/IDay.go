package dayUtils

import (
	"aoc_go_2024/src/utils"
	"fmt"
	"log"
	"time"
)

type IDay[T any] interface {
	Parse([]string) T
	Part1(bool, T) int64
	Part2(bool, T) int64
}

func Run[T any](d IDay[T], args Args) {
	if args.Day < 1 {
		log.Fatal("This should never happend")
	}
	fmt.Println("Welcome to day", args.Day, "!!!")

	lines := utils.GetLines(args.IsTest, uint(args.Day))
	input := d.Parse(lines)

	if args.OnlyP1 || !args.OnlyP2 {
		start := time.Now()
		sol1 := d.Part1(args.Debug, input)
		ellapsed := time.Since(start)
		fmt.Println("Part 1:", sol1, "in", ellapsed)
	}

	if args.OnlyP2 || !args.OnlyP1 {
		start := time.Now()
		sol2 := d.Part2(args.Debug, input)
		ellapsed := time.Since(start)
		fmt.Println("Part 2:", sol2, "in", ellapsed)
	}
}
