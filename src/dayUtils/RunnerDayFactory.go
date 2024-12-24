package dayUtils

import (
	"aoc_go_2024/src/day1"
	"aoc_go_2024/src/day12"
	"aoc_go_2024/src/day13"
	"aoc_go_2024/src/day14"
	"aoc_go_2024/src/day15"
	"aoc_go_2024/src/day2"
	"aoc_go_2024/src/day3"
	"aoc_go_2024/src/day4"
	"log"
)

func RunSpecificDay(args Args) {
	switch args.Day {
	case 1:
		Run(day1.Day1{}, args)
	case 2:
		Run(day2.Day2{}, args)
	case 3:
		Run(day3.Day3{}, args)
	case 4:
		Run(day4.Day4{}, args)
	case 12:
		Run(day12.Day12{}, args)
	case 13:
		Run(day13.Day13{}, args)
	case 14:
		Run(day14.Day14{IsTest: args.IsTest}, args)
	case 15:
		Run(day15.Day15{}, args)
	default:
		log.Fatal("there is no such day yet")
	}
}
