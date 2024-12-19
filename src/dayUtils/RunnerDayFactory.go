package dayUtils

import (
	"aoc_go_2024/src/day1"
	"aoc_go_2024/src/day12"
	"log"
)

func RunSpecificDay(day uint, isTest, debug bool) {
	switch day {
	case 1:
		Run(day1.Day1{}, isTest, debug)
	case 2:
		//Run(day2.Day2{}, isTest, debug)
	case 3:
		//Run(day3.Day3{}, isTest, debug)
	case 4:
		//Run(day4.Day4{}, isTest, debug)
	case 12:
		Run(day12.Day12{}, isTest, debug)
	default:
		log.Fatal("there is no such day yet")
	}
}
