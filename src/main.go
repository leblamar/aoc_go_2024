package main

import (
	"aoc_go_2024/src/day1"
	"aoc_go_2024/src/day2"
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

	just_a_test := false
	if len(args) == 3 && args[2] == "test" {
		just_a_test = true
	}

	switch day {
	case 1:
		day1.Day1(just_a_test)
	case 2:
		day2.Day2(just_a_test)
	}
}
