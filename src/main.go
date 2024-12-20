package main

import (
	"aoc_go_2024/src/dayUtils"
)

func runAll() {
	for day := 0; true; day++ {
		args := dayUtils.GetDefaultArgs()
		args.Day = day
		dayUtils.RunSpecificDay(args)
	}
}

func main() {
	args := dayUtils.GetArgs()
	if args.Day < 1 {
		runAll()
		return
	}

	dayUtils.RunSpecificDay(args)
}
