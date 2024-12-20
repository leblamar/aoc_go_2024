package dayUtils

import (
	"aoc_go_2024/src/day1"
	"aoc_go_2024/src/day12"
	"log"
)

func RunSpecificDay(args Args) {
	switch args.Day {
	case 1:
		Run(day1.Day1{}, args)
	case 2:
		//Run(day2.Day2{}, args)
	case 3:
		//Run(day3.Day3{}, args)
	case 4:
		//Run(day4.Day4{}, args)
	case 12:
		Run(day12.Day12{}, args)
	default:
		log.Fatal("there is no such day yet")
	}
}
