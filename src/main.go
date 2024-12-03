package main

import (
	"aoc_go_2024/src/day1"
	"aoc_go_2024/src/day2"
	"aoc_go_2024/src/day3"
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
		return
	} else if len(args) > 1 {
		day, err = strconv.Atoi(args[1])

		if err != nil {
			log.Fatal(err)
			return
		}
	}

	justATest := false
	if len(args) == 3 && args[2] == "test" {
		justATest = true
	}

	switch day {
	case 1:
		day1.Day1(justATest)
	case 2:
		day2.Day2(justATest)
	case 3:
		day3.Day3(justATest)
	}
}
