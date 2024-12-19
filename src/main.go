package main

import (
	"aoc_go_2024/src/dayUtils"
	"log"
	"os"
	"strconv"
)

func runAll() {
	for day := 0; true; day++ {
		dayUtils.RunSpecificDay(uint(day), false, false)
	}
}

func getArgs() (day int, isTest, debug bool) {
	args := os.Args
	if len(args) <= 1 {
		return -1, false, false
	}

	day, err := strconv.Atoi(args[1])
	if err != nil {
		log.Fatal(err)
		return -1, false, false
	}

	isTest = false
	debug = false
	if len(args) >= 3 {
		for _, arg := range args {
			if arg == "-t" {
				isTest = true
			} else if arg == "-d" {
				debug = true
			}
		}
	}

	return
}

func main() {
	day, isTest, debug := getArgs()
	if day < 1 {
		runAll()
		return
	}

	dayUtils.RunSpecificDay(uint(day), isTest, debug)
}
