package main

import (
	"log"
	"os"
	"strconv"
)

func runAll() {
	var day uint = 0
	for curDay, err := GetDay(day); err != nil; day++ {
		Run(curDay, false, false)
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

	curDay, err := GetDay(uint(day))
	if err != nil {
		log.Fatal(err)
		return
	}
	Run(curDay, isTest, debug)
}
