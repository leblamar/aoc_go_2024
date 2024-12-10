package main

import (
	"aoc_go_2024/src/day1"
	"aoc_go_2024/src/day2"
	"aoc_go_2024/src/day3"
	"aoc_go_2024/src/day4"
	"aoc_go_2024/src/day5"
	"aoc_go_2024/src/day6"
	"aoc_go_2024/src/day7"
	"aoc_go_2024/src/day8"
	"log"
	"os"
	"strconv"
)

func main() {
	args := os.Args

	var day int
	var err error
	if len(args) == 1 {
		day1.Day1(false)
		day2.Day2(false)
		day3.Day3(false)
		day4.Day4(false)
		day5.Day5(false)
		day6.Day6(false, false)
		day7.Day7(false, false)
		day8.Day8(false)
		return
	} else if len(args) > 1 {
		day, err = strconv.Atoi(args[1])

		if err != nil {
			log.Fatal(err)
			return
		}
	}

	justATest := false
	debug := false
	if len(args) >= 3 {
		for _, arg := range args {
			if arg == "-t" {
				justATest = true
			} else if arg == "-d" {
				debug = true
			}
		}
	}

	switch day {
	case 1:
		day1.Day1(justATest)
	case 2:
		day2.Day2(justATest)
	case 3:
		day3.Day3(justATest)
	case 4:
		day4.Day4(justATest)
	case 5:
		day5.Day5(justATest)
	case 6:
		day6.Day6(justATest, debug)
	case 7:
		day7.Day7(justATest, debug)
	case 8:
		day8.Day8(justATest)
	}
}
